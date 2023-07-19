package packages

import (
	"fmt"
	"time"
)

type MyError struct {
	When string
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func Run() error {
	time := time.Now().Format("2006-01-02 15:04:05")
	return &MyError{
		time,
		"it didn't work",
	}
}
