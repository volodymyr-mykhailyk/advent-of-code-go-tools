package streams

import (
	"reflect"
	"testing"
)

func TestCollectValues(t *testing.T) {
	channel := make(chan int)
	want := []int{0, 1, 2, 3, 4}
	go func() {
		for i := 0; i < 5; i++ {
			channel <- i
		}
		close(channel)
	}()

	if got := CollectValues(channel).([]int); !reflect.DeepEqual(got, want) {
		t.Errorf("CollectValues() = %v, want %v", got, want)
	}
}
