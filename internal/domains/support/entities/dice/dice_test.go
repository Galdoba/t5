package roll

import (
	"fmt"
	"testing"
)

func TestRoll(t *testing.T) {
	fmt.Println(Roll(2))
	fmt.Println(Roll(3))
	rr := Roll(5)
	fmt.Println(rr.Sum(), rr.Results())
}
