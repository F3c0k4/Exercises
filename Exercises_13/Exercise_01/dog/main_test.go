package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	v := Years(5)

	if v != 35 {
		t.Error()
	}

}

func ExampleYears() {
	fmt.Println(Years(5))
	// Output: 35
}
