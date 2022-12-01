package tasks

import "fmt"

func Announce(message string) {
	fmt.Println("--------------------------------------------")
	fmt.Println(message)
	fmt.Println("--------------------------------------------")
}

func Communicate(message string, args ...any) {
	fmt.Printf(message+"\n", args...)
}
