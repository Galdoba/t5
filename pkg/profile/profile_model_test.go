package profile

import (
	"fmt"
	"testing"

	"github.com/Galdoba/t5/pkg/ehex"
	"github.com/Galdoba/t5/pkg/wrld"
)

func TestProfile(t *testing.T) {
	w := &wrld.World{}
	w.S = ehex.FromInt(10)
	w.A = ehex.FromInt(9)
	w.H = ehex.FromInt(11)
	pr := Profile(w)
	fmt.Println(pr.String())
	err := InjectString(w, "CBD")
	fmt.Println(err)
	pr = Profile(w)
	fmt.Println(pr.String())
}
