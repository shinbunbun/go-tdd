package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	str := Repeat("a", 10)
	expected := "aaaaaaaaaa"
	if str != expected {
		t.Errorf("expected '%q' but got '%q'", expected, str)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	res := Repeat("b", 10)
	fmt.Println(res)
	// Output: bbbbbbbbbb
}
