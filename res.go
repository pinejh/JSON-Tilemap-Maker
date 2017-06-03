package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Resources struct {
	images map[string]*sf.Texture
	fonts  map[string]*sf.Font
}

func LoadResources() (r *Resources) {
	r = new(Resources)
	r.images = make(map[string]*sf.Texture)
	r.fonts = make(map[string]*sf.Font)
	r.loadImages("./assets/images")
	r.loadFonts("./assets/fonts")
	return
}

func (r *Resources) loadImages(dir string) {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			r.loadImages(dir + "/" + f.Name())
		} else if filepath.Ext(f.Name()) == ".png" {
			texture := sf.NewTexture(dir + "/" + f.Name())
			r.images[f.Name()] = texture
		}
	}

}

func (r *Resources) loadFonts(dir string) {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			r.loadFonts(dir + "/" + f.Name())
		} else if filepath.Ext(f.Name()) == ".ttf" {
			font := sf.NewFont(dir + "/" + f.Name())
			r.fonts[f.Name()] = font
		}
	}
}
