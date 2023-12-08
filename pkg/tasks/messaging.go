package tasks

import "fmt"
import "time"

func Announce(message string) {
	fmt.Println("****************************************************************************************")
	fmt.Println(message)
	fmt.Println("----------------------------------------------------------------------------------------")
}

func Highlight(message string) {
	fmt.Println("-------  " + message + "  -------  ")
}

func Communicate(message string, args ...any) {
	fmt.Printf(message+"\n", args...)
}

func ElapsedTime(message string, start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("%s took: %s\n", message, elapsed)
}
