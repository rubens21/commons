package commons

import (
	"errors"
	"time"
	"github.com/rafaeljesus/rabbus"
	"context"
	"log"
	"net/url"
)

type Comunicator struct {
	onMessage   func(msg *rabbus.ConsumerMessage)
	speaker     *rabbus.Rabbus
	queue       string
	exchange    string
	url    		url.URL
	stackCloses []func()
}

func CreateListener(urlConn url.URL, exchange, chat string, onMessage func(msg *rabbus.ConsumerMessage)) *Comunicator {
	c := new(Comunicator)
	c.onMessage = onMessage
	c.queue = chat
	c.url = urlConn
	c.exchange = exchange
	c.startListener()
	return c
}

func CreateSpeaker(urlConn url.URL, exchange, chat string) *Comunicator {
	c := new(Comunicator)
	c.queue = chat
	c.exchange = exchange
	c.url = urlConn
	c.startSpeaker()
	return c
}
func CreateCommunicator(urlConn url.URL, exchange, chat string, onMessage func(msg *rabbus.ConsumerMessage)) *Comunicator {
	c := CreateListener(urlConn, exchange, chat, onMessage)
	c.startSpeaker()
	return c
}

func (c *Comunicator) startSpeaker() {
	cbStateChangeFunc := func(name, from, to string) {
		// do something when state is changed
	}
	var err error
	c.speaker, err = rabbus.New(
		c.url.String(),
		rabbus.Durable(true),
		rabbus.Attempts(5),
		rabbus.Sleep(time.Second*2),
		rabbus.Threshold(3),
		rabbus.OnStateChange(cbStateChangeFunc),
	)

	if err != nil {
		panic(errors.New("Fail on AMQP connection: " + err.Error()))
	}

	c.deferCloser(func() {
		if err := c.speaker.Close(); err != nil {
			log.Printf("Fail on closing producer: %s", err)
		}
	})

	ctx, cancel := context.WithCancel(context.Background())
	c.deferCloser(func() {
		cancel()
	})

	go c.speaker.Run(ctx)
}

func (c *Comunicator) Send(amsg []byte) {
	rabbusMsg := rabbus.Message{
		Exchange:     c.exchange,
		Kind:         "direct",
		Key:          c.queue,
		Payload:      amsg,
		DeliveryMode: rabbus.Persistent,
	}

	c.speaker.EmitAsync() <- rabbusMsg
	outer:
	for {
		select {
		case <-c.speaker.EmitOk():
			break outer
		case err := <-c.speaker.EmitErr():
			LogError("Failed to send message: %s", err)
			break outer
		}
	}
}

func (c *Comunicator) startListener() {
	cbStateChangeFunc := func(name, from, to string) {
		// do something when state is changed
	}
	r, err := rabbus.New(
		c.url.String(),
		rabbus.Durable(true),
		rabbus.Attempts(5),
		rabbus.Sleep(time.Second*2),
		rabbus.Threshold(3),
		rabbus.OnStateChange(cbStateChangeFunc),
	)
	if err != nil {
		LogError("Failed to init rabbus connection %s", err)
		return
	}

	c.deferCloser(func() {
		if err := r.Close(); err != nil {
			LogError("Failed to close rabbus connection %s", err)
		}
	})

	ctx, cancel := context.WithCancel(context.Background())
	c.deferCloser(func() {
		cancel()
	})
	go r.Run(ctx)

	messages, err := r.Listen(rabbus.ListenConfig{
		Exchange: c.exchange,
		Kind:     "direct",
		Key:      c.queue,
		Queue:    c.queue,
	})
	if err != nil {
		LogError("Failed to create listener %s", err)
		return
	}

	c.deferCloser(func() {
		close(messages)
	})

	go func(messages chan rabbus.ConsumerMessage) {
		for m := range messages {
			c.onMessage(&m)
		}
	}(messages)
}

func (c *Comunicator) deferCloser(closer func()) {
	if c.stackCloses == nil {
		c.stackCloses = []func(){}
	}
	c.stackCloses = append([]func(){closer}, c.stackCloses...)
}

func (c *Comunicator) Close() {
	for _, function := range c.stackCloses {
		function()
	}
}
