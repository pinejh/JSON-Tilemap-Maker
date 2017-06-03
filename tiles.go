package main

import (
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Tile struct {
	*sf.Sprite

	x float32
	y float32
}
