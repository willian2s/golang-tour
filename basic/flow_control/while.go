package flow_control

func While() int {
	sum := 1

	for sum < 1000 {
		sum += sum
	}

	return sum
}
