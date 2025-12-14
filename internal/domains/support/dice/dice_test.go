package dice

import (
	"fmt"
	"testing"
)

func TestRoll(t *testing.T) {
	fmt.Println(Roll(2))
	fmt.Println(Roll(3))
	fmt.Println(Roll(4))
}
