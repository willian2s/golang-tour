package class

import "time"

func SwitchEvaluationOrder(today time.Weekday) string {
	switch time.Monday {
	case today + 0:
		return "Today"
	case today + 1:
		return "Tomorrow"
	case today + 2:
		return "In two days"
	default:
		return "Too far away"
	}
}
