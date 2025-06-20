package basic

import "github.com/Galdoba/t5/internal/cosmology/star"

type StarSystem struct {
	composition []string
	t5ssData    string
	Stars       map[string]*star.Star
}
