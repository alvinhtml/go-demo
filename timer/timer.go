package timer

import (
	"fmt"
	"time"
)

// DoWithTimer
func DoWithTimer() {
	fmt.Println(time.Now())

	timer := time.NewTimer(time.Second * 3)

	ch := <-timer.C

	fmt.Println(ch)
}

// DoWithSleep
func DoWithSleep() {
	fmt.Println(time.Now())
	time.Sleep(time.Second * 3)
	fmt.Println(time.Now())
}

// DoWithAfter
func DoWithAfter() {
	fmt.Println(time.Now())
	fmt.Println(<-time.After(time.Second * 3))
}

func init() {
	DoWithTimer()
}
