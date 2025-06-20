package generate

import (
	"fmt"
	"testing"

	"github.com/Galdoba/t5/internal/cosmology/star"
)

func TestNewGenerator(t *testing.T) {
	gn := NewGenerator(WithSeed(4253622), WithRule(Rule_Method, Method_BasicT5))
	st, err := gn.GenerateStar(star.WithPosition(star.Primary))
	fmt.Println(err)
	fmt.Println(st)
}
