package tasks

import (
	"github.com/volodymyr-mykhailyk/advent-of-code-go-tools/pkg/streams"
	"testing"
)

func TestReadLines(t *testing.T) {
	content := ReadLines("LICENSE")
	t.Run("Length", func(t *testing.T) {
		got := len(content)
		want := 25
		if got != want {
			t.Errorf("len(ReadLines()) = %v, want %v", got, want)
		}
	})
	t.Run("First Line", func(t *testing.T) {
		got := content[0]
		want := "This is free and unencumbered software released into the public domain."
		if got != want {
			t.Errorf("ReadLines()[0] = %v, want %v", got, want)
		}
	})
	t.Run("Empty Line", func(t *testing.T) {
		got := content[1]
		want := ""
		if got != want {
			t.Errorf("ReadLines()[1] = %v, want %v", got, want)
		}
	})
}

func TestStreamLines(t *testing.T) {
	lines := make(chan string)
	go StreamLines("LICENSE", lines)
	content := streams.CollectValues(lines).([]string)
	t.Run("Length", func(t *testing.T) {
		got := len(content)
		want := 25
		if got != want {
			t.Errorf("len(ReadLines()) = %v, want %v", got, want)
		}
	})
	t.Run("First Line", func(t *testing.T) {
		got := content[0]
		want := "This is free and unencumbered software released into the public domain."
		if got != want {
			t.Errorf("ReadLines()[0] = %v, want %v", got, want)
		}
	})
	t.Run("Empty Line", func(t *testing.T) {
		got := content[1]
		want := ""
		if got != want {
			t.Errorf("ReadLines()[1] = %v, want %v", got, want)
		}
	})
}
