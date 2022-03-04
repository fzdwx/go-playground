package main

import "time"

func main() {

	i := 12
	switch i {
	case 1:
		println("one")
	case 2:
		println("two")
	case 4:
		println("four")
	case 12:
		println("twelve")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("weekend")
	default:
		println("weekday")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			println("bool")
		case int:
			println("int")
		default:
			println("unknown")
		}
	}

	whatAmI(1)
	whatAmI(false)
	whatAmI("hello")
}
