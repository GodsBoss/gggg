//go:build js && wasm

package canvas2drendering

import (
	"github.com/GodsBoss/gggg/v2/pkg/dom"
)

type sprite struct {
	// sourceImage is the image this sprite is taken from.
	sourceImage *dom.Image

	// width is both the width of the source and the target, but may be scaled.
	width int

	// height is both the height of the source and the target, but may be scaled.
	height int

	// sourceX is the left horizontal start of the sprite at the source.
	sourceX int

	// sourceY is the top vertical start of the sprite at the source.
	sourceY int

	destinationX int
	destinationY int

	// destinationScale is the scale the sprite will be scaled to.
	destinationScale int
}

func (s sprite) Render(ctx *dom.Context2D) {
	ctx.DrawImage(
		s.sourceImage,
		s.sourceX,
		s.sourceY,
		s.width,
		s.height,
		s.destinationX,
		s.destinationY,
		s.width*s.destinationScale,
		s.height*s.destinationScale,
	)
}

type SpriteMap interface {
	// AddSpriteSpecification adds a sprite specification and returns a sprite key which
	// acts as an identifier for that specification. It is then passed to CreateSprite
	// and CreateScaledSprite.
	AddSpriteSpecification(SpriteSpecification) SpriteKey

	// CreateSprite creates a scales sprite. The scaling affects the resulting
	// sprite size, but not its position.
	CreateSprite(key SpriteKey, attr SpriteAttributes, x int, y int, scale int) Renderable
}

// NewSpriteMap creates a new sprite map using the source image as the source for
// sprites. Afterwards, define sprites.
func NewSpriteMap(sourceImage *dom.Image) SpriteMap {
	return &spriteMap{
		sourceImage: sourceImage,
		sprites:     make(map[SpriteKey]SpriteSpecification),
	}
}

type spriteMap struct {
	currentSpriteKey spriteKey
	sourceImage      *dom.Image
	sprites          map[SpriteKey]SpriteSpecification
}

func (sm *spriteMap) AddSpriteSpecification(specs SpriteSpecification) SpriteKey {
	sKey := spriteKey(sm.currentSpriteKey)
	sm.currentSpriteKey++
	sm.sprites[sKey] = specs
	return sKey
}

func (sm *spriteMap) CreateSprite(key SpriteKey, attr SpriteAttributes, x int, y int, scale int) Renderable {
	dataByAttr, ok := sm.sprites[key]
	if !ok {
		return NopRenderable()
	}
	data, ok := dataByAttr[attr]
	if !ok {
		return NopRenderable()
	}
	return sprite{
		sourceImage:      sm.sourceImage,
		width:            data.Width,
		height:           data.Height,
		sourceX:          data.X,
		sourceY:          data.Y,
		destinationX:     x,
		destinationY:     y,
		destinationScale: scale,
	}
}

// SpriteKey is the key returned by SpriteMap.AddSpriteSpecification to identify
// a given sprite specification. It cannot be implemented by outside packages.
type SpriteKey interface {
	spriteKey()
}

type spriteKey int

func (k spriteKey) spriteKey() {}

// SpriteAttributes define scale and rotation of a sprite.
type SpriteAttributes struct {
	Scale    int
	Rotation int
}

type SpriteData struct {
	X      int
	Y      int
	Width  int
	Height int
}

// SpriteSpecification serves mostly as a name.
type SpriteSpecification map[SpriteAttributes]SpriteData

// MergeSpriteSpecifications takes several sprite specifications and returns
// their union. If specifications contain the same attributes as previous
// specifications, they overwrite their values.
func MergeSpriteSpecifications(specss ...SpriteSpecification) SpriteSpecification {
	result := make(SpriteSpecification)
	for i := range specss {
		for attr := range specss[i] {
			result[attr] = specss[i][attr]
		}
	}
	return result
}

// CreateSpriteSpecs creates a partial sprite specification with a fixed scale,
// width and height. rotations is the number of attributes, with Rotation ranging
// from 0 to rotations-1. Subsequent sprite data is shifted by offset.
func CreateSpriteSpecs(
	scale int,
	width, height int,
	rotations int,
	startX, startY int,
	offsetX, offsetY int,
) SpriteSpecification {
	result := make(SpriteSpecification)
	for rotation := 0; rotation < rotations; rotation++ {
		result[SpriteAttributes{
			Scale:    scale,
			Rotation: rotation,
		}] = SpriteData{
			X:      startX + rotation*offsetX,
			Y:      startY + rotation*offsetY,
			Width:  width,
			Height: height,
		}
	}
	return result
}
