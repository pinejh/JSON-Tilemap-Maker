package main

import (
	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	menubarHeight = 25
	fontsize      = 15
	buttonpadding = 20
)

var (
	menucolor       = sf.Color{230, 230, 230, 255}
	menuselectcolor = sf.Color{215, 215, 215, 255}
)

type MenuBar struct {
	*sf.RectangleShape

	items []*MenuBarItem
}

type MenuBarItem struct {
	*Button

	showdropdown  bool
	dropdown      *sf.RectangleShape
	dropdownitems []*Button
}

type Button struct {
	*sf.RectangleShape

	text     *sf.Text
	canClick bool
	active   bool
}

func NewMenuBar() (mb *MenuBar) {
	mb = new(MenuBar)
	mb.RectangleShape = sf.NewRectangleShape(sf.Vector2f{screenWidth, menubarHeight})
	mb.SetFillColor(menucolor)

	mb.items = []*MenuBarItem{}
	mb.AddMenuBarItem("File")
	mb.AddMenuBarItem("Edit")
	mb.AddMenuBarItem("Spritesheet")
	//mb.items[2].AddDropdownItem("Load Spritesheet...")

	return
}

func (mb *MenuBar) AddMenuBarItem(text string) {
	mbi := new(MenuBarItem)

	mbi.Button = NewButton(text)

	if len(mb.items) > 0 {
		last := mb.items[len(mb.items)-1].Button.RectangleShape.GetGlobalBounds()
		mbi.Button.SetPosition(sf.Vector2f{last.Left + last.Width, 0})
		mbi.Button.text.SetPosition(sf.Vector2f{last.Left + last.Width + mbi.Button.GetGlobalBounds().Width/2, 3})
	} else {
		mbi.Button.SetPosition(sf.Vector2f{0, 0})
		mbi.Button.text.SetPosition(sf.Vector2f{mbi.Button.GetGlobalBounds().Width / 2, 3})
	}

	mb.items = append(mb.items, mbi)
}

func NewButton(text string) (b *Button) {
	b = new(Button)

	b.text = sf.NewText(text, res.fonts["Arial.ttf"], fontsize)
	b.text.SetColor(sf.ColorBlack)
	tbox := b.text.GetGlobalBounds()
	b.text.SetOrigin(sf.Vector2f{tbox.Width / 2, 0})

	b.RectangleShape = sf.NewRectangleShape(sf.Vector2f{tbox.Width + buttonpadding, menubarHeight})
	b.SetFillColor(menucolor)
	b.SetOutlineColor(sf.Color{0, 0, 0, 25})
	b.SetOutlineThickness(-1)

	return
}
