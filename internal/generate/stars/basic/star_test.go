package basic

import (
	"fmt"
	"testing"
	"time"

	"github.com/Galdoba/t5/pkg/dice"
)

func TestNewStar(t *testing.T) {

	for i := 0; i < 40; i++ {
		time.Sleep(time.Millisecond * 5)
		st, err := NewStar(dice.NewDicepool())
		switch err == nil {
		case true:
			fmt.Printf("%v:%v\n", st.String(), st.DebugText())
		case false:
			fmt.Println(err)
		}
	}
}
