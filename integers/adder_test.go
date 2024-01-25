package integers

import (
	"fmt"
	"testing"
	"time"
)

// Add takes two integers and returns sum of them.
func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	excepted := 4

	if sum != excepted {
		t.Errorf("expected '%d' but got '%d'", excepted, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	time.Sleep(1 * time.Second)
	fmt.Println(sum)
	// Output: 6
}

func ExampleAdd_second() {
	sum := Add(1, 5)
	time.Sleep(1 * time.Second)
	fmt.Println(sum)
	// Output: 6
}
