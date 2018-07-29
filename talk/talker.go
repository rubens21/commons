package talk

import (
	"github.com/gorilla/websocket"
	"net/url"
	"github.com/makeitplay/commons/BasicTypes"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/makeitplay/commons"
	"sync"
)

type Channel struct {
	ws             *websocket.Conn
	playerSpec     BasicTypes.PlayerSpecifications
	urlConnection  url.URL
	listenerTask   *commons.Task
	onMessage      func(bytes []byte)
	connectionOpen bool
	mu             sync.Mutex
}

func NewTalkChannel(url url.URL, playerSpec BasicTypes.PlayerSpecifications) *Channel {
	c := Channel{}
	c.playerSpec = playerSpec
	c.urlConnection = url
	return &c
}

func (c *Channel) Send(data []byte) error {
	return c.ws.WriteMessage(websocket.TextMessage, data)
}

func (c *Channel) OpenConnection(onMessage func(bytes []byte)) error {
	c.onMessage = onMessage
	if err := c.dial(); err != nil {
		return err
	}
	c.connectionOpen = true // please, let me know when gorilla brings a better way to figure out whether the conn is open or not
	c.defineListenerTask()
	c.defineWebsocketCloseHandler()
	c.listenerTask.Start()
	return nil
}

func (c *Channel) dial() error {
	connectHeader := http.Header{}
	specJson, err := json.Marshal(c.playerSpec)
	if err != nil {
		return fmt.Errorf("fail on bulding the player spec header: %s", err.Error())
	}
	connectHeader.Add("X-Player-Specs", string(specJson))

	c.ws, _, err = websocket.DefaultDialer.Dial(c.urlConnection.String(), connectHeader)
	if err != nil {
		return fmt.Errorf("fail on dialing to ws server: %s", err.Error())
	}
	return nil
}

func (c *Channel) defineListenerTask() {
	c.listenerTask = commons.NewTask(func(task *commons.Task) {
		defer func() {
			if err := recover(); err != nil {
				commons.LogWarning("Connection lost: %s", err)
			}
		}()

		c.mu.Lock()
		defer c.mu.Unlock()
		msgType, message, err := c.ws.ReadMessage()
		if msgType == -1 {
			commons.LogError("Msg error: %s %s", msgType, err)
			task.RequestStop()
			return
		} else if err != nil {
			commons.LogError("Fail reading websocket message (%d): %s", msgType, err)
			task.RequestStop()
			return
		} else {
			c.onMessage(message)
		}
	})
	c.listenerTask.OnStop = func(task *commons.Task) {
		if task.IsRunning() {
			//@todo implement a recover method to try to connect again when an error is detected
			commons.LogError("Here is a nice place to implement a recover method")
			commons.Cleanup(true)
		}
	}
	c.listenerTask.Start()
}
func (c *Channel) CloseConnection() {
	c.listenerTask.RequestStop()
	if c.connectionOpen { // trying to avoid panic on writing in a closed connection
		err := c.ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			commons.LogError("Fail on closing ws connection: %s", err.Error())
		}

	}
	c.connectionOpen = false
	c.ws.Close()
}

func (c *Channel) defineWebsocketCloseHandler() {
	c.ws.SetCloseHandler(func(code int, text string) error {
		c.connectionOpen = false
		c.listenerTask.RequestStop()
		if code == websocket.CloseNormalClosure {
			commons.Log("Connection closed by the server")
		} else {
			commons.LogError("Connection abnormal closed: %d-%s", code, text)
		}
		return nil
	})
}
