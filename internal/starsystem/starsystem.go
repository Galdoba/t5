package starsystem

import "github.com/Galdoba/t5/internal/grid/stellarhex"

/*
starsystem
 is mutable
 has profile
 can be fully restored from profile
 can be partialy restored from profile

 has 0-8 bright stars
 has 0-2 dim stars
 has 0-6 rogue GG
 has 0-14 rogue planetoids

*/

type StarSystem struct {
	composition []string
	t5ssData    string
}

func NewStarSystem(hex stellarhex.StellarHex) {

}
