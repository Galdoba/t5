package basic

import (
	"fmt"
	"testing"

	"github.com/Galdoba/t5/pkg/dice"
)

func TestNew(t *testing.T) {
	dp := dice.NewDicepool()
	w := New(MainWorld(true))
	w.Generate(dp)
	fmt.Println("Mainworld")
	fmt.Println(w.String())
	fmt.Println("Non Mainworld")
	w2 := New(MainWorld(false))
	w2.Generate(dp)
	fmt.Println(w2.String())
}
