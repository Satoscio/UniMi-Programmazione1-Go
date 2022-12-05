package main

import (
	"fmt"
	"time"
)

type Clock struct {
	hour, min, sec int
}

func countdown(timer *Clock) {
	if timer.sec < 60 && timer.sec > 0 {
		timer.sec--
	} else {
		timer.sec = 59
		updateMin(timer)
	}
}

func updateMin(timer *Clock) {
	if timer.min < 60 && timer.min > 0 {
		timer.min--
	} else {
		timer.min = 59
		updateHour(timer)
	}
}

func updateHour(timer *Clock) {
	timer.hour--
}

func main() {
	var h, m, s int
	var timer Clock

inputFor:
	for {
		fmt.Print("time (hh mm ss): ")
		fmt.Scan(&h, &m, &s)
		if (m > 60 || m < 0) || (s > 60 || s < 0) {
			fmt.Println("Valori errati!")
		} else {
			break inputFor
		}
	}

	timer.hour, timer.min, timer.sec = h, m, s
	for {
		countdown(&timer)
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println(timer)
		if timer.hour == 0 && timer.min == 0 && timer.sec == 0 {
			break
		}
	}
}
