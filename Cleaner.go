package commons

import (
	"fmt"
	"github.com/fatih/color"
)

type Cleaner struct {
	name string
	callback func(bool)
}

var cleanerStack []Cleaner

func RegisterCleaner(name string, callback func(bool)) {
	if cleanerStack == nil {
		cleanerStack = make([]Cleaner, 0)
	}
	cleaner := Cleaner{name, callback}
	cleanerStack = append([]Cleaner{cleaner}, cleanerStack...)
}

func Cleanup(interrupted bool)  {
	color.Unset()
	Log("Cleaning up")
	for _, cleaner := range cleanerStack {
		LogInfo("Closing: %s", cleaner.name)
		cleaner.callback(interrupted)
	}
	fmt.Println("-- All cleaners done -- ")
}