// +build js,wasm

package dom

import (
	"syscall/js"
)

type Document struct {
	value js.Value
}

func (doc *Document) GetElementByID(id string) (*Element, error) {
	jsEl := doc.value.Call("getElementById", id)
	if jsEl.IsNull() {
		return nil, newError("element with id " + id + " does not exist")
	}
	return &Element{
		value: jsEl,
	}, nil
}

func (doc *Document) CreateCanvasElement() (*Canvas, error) {
	jsCanvas := doc.value.Call("createElement", "canvas")
	if jsCanvas.IsNull() {
		return nil, newError("could not create canvas element")
	}
	return &Canvas{
		value: jsCanvas,
	}, nil
}

func (doc *Document) CreateImageElement(src string) (*Image, error) {
	jsImg := doc.value.Call("createElement", "img")
	if jsImg.IsNull() {
		return nil, newError("could not create image element")
	}
	jsImg.Set("src", src)
	img := &Image{
		value: jsImg,
	}
	return img, nil
}