package main

import (
	"runtime"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

var (
	spritesize = 16
)

var res *Resources

var menubar *MenuBar

//TODO: Tiles
//var tilemap [][]*Tile

//TODO: Scrollbars
//var scrollbarRight
//var scrollbarBottom

func main() {
	runtime.LockOSThread()

	res = LoadResources()

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "JSON Tilemap Editor", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	menubar = NewMenuBar()

	for window.IsOpen() {
		if event := window.PollEvent(); event != nil {
			switch event.Type {
			case sf.EventClosed:
				window.Close()
			}
		}

		window.Clear(sf.ColorWhite)

		window.Draw(menubar)
		for _, item := range menubar.items {
			window.Draw(item.Button.RectangleShape)
			window.Draw(item.Button.text)
		}

		window.Display()
	}
}
