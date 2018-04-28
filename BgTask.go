package commons

// thanks https://medium.com/@matryer/stopping-goroutines-golang-1bf28799c1cb

type Task struct {
	OnStop      func(*Task)
	task        func(*Task)
	stopchan    chan struct{}
	running 	bool
}

func NewTask(task func(*Task)) *Task {
	t := new(Task)
	t.task = task
	return t
}

func (t *Task) Start()  {
	t.stopchan = make(chan struct{})
	t.running = true
	go func(){ // work in background
		// TODO: do setup work
		defer func(){
			//bug! It does not work!
			if t.OnStop != nil {
				t.OnStop(t)
			}
			t.running = false
		}()
		for {
			select {
			default:
				t.task(t)
			case <-t.stopchan:
				return
			}
		}
	}()
}

func (t *Task) Stop()  {
	close(t.stopchan)  // tell it to stop
}
