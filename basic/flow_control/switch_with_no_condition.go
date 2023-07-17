package class

import (
	"time"
)

func SwitchWithNoCondition() string {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		return "Good morning!"
	case t.Hour() < 17:
		return "Good afternoon!"
	default:
		return "Good evening!"
	}
}
