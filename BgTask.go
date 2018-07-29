package commons

// thanks https://medium.com/@matryer/stopping-goroutines-golang-1bf28799c1cb

type Task struct {
	OnStop        func(*Task)
	task          func(*Task)
	stopchan      chan struct{}
	running       bool
	stopRequested bool
}

func NewTask(task func(*Task)) *Task {
	t := new(Task)
	t.task = task
	return t
}

func (t *Task) Start() {
	t.stopchan = make(chan struct{})
	t.running = true
	go func() { // work in background
		// TODO: do setup work
		defer func() {
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

func (t *Task) RequestStop() {
	if !t.stopRequested {
		t.stopRequested = true
		close(t.stopchan) // tell it to stop
	}
}

func (t *Task) StopRequested() bool {
	return t.stopRequested
}
func (t *Task) IsRunning() bool {
	return t.running
}
