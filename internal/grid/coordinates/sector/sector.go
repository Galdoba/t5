package sector

import (
	"fmt"
	"strings"
)

func Name(sx, sy int) string {
	return "Undefined"
}

func Abb(name string) string {
	letters := strings.Split(name, "")
	return strings.Join(letters[:4], "")
}

func Hex(lx, ly int) string {
	w := fmt.Sprintf("%v", ly)
	for len(w) < 2 {
		w = "0" + w
	}
	h := fmt.Sprintf("%v", lx)
	for len(h) < 2 {
		h = "0" + h
	}
	return h + w
}
