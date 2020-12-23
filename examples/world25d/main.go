package main

import (
	"math"
	"math/rand"

	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/dominit"
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/world25d"

	m "github.com/GodsBoss/go-pkg/affinematrix2d"
)

func main() {
	win, _ := dom.GlobalWindow()
	doc, _ := win.Document()
	sprite, _ := doc.CreateImageElement("../assets/small_square.png")

	dominit.Run(
		&game{
			sprite:  sprite,
			cam:     world25d.NewCamera(),
			objects: createRandomObjects(100, -800, -600, 1600, 1200),
		},
	)
	<-make(chan struct{}, 0)
}

func createRandomObjects(count int, minX, minY, maxX, maxY int) []world25d.Object {
	objects := make([]world25d.Object, count)
	for i := 0; i < count; i++ {
		objects[i] = world25d.Object{
			X:        float64(rand.Intn(maxX-minX) + minX),
			Y:        float64(rand.Intn(maxY-minY) + minY),
			Z:        float64(rand.Intn(100)),
			Rotation: rand.Float64() * math.Pi,
		}
	}
	return objects
}

type game struct {
	sprite *dom.Image
	output *dom.Context2D

	cam     world25d.Camera
	objects []world25d.Object

	left  int
	right int
	down  int
	up    int

	rotLeft  int
	rotRight int
}

func (g *game) TicksPerSecond() int {
	return 40
}

func (g *game) Logic() dominit.Logic {
	return g
}

func (g *game) Tick(ms int) {
	camSpeed := 10.0

	x, y := g.cam.Position()

	rotSpeed := 0.05
	r := g.cam.Rotation() + rotSpeed*float64(g.rotRight-g.rotLeft)
	g.cam.SetRotation(r)

	transform := m.Combine(
		m.Translation(x, y),
		m.Rotation(r),
		m.Translation(float64(g.right-g.left)*camSpeed, float64(g.down-g.up)*camSpeed),
	)

	pos := transform.Transform(m.VectorFromCartesian(0, 0))
	g.cam.SetPosition(pos.X(), pos.Y())
}

func (g *game) ReceiveKeyEvent(event interaction.KeyEvent) {
	if event.Type == interaction.KeyDown {
		switch event.Key {
		case "w":
			g.up = 1
		case "a":
			g.left = 1
		case "s":
			g.down = 1
		case "d":
			g.right = 1
		case "q":
			g.rotLeft = 1
		case "e":
			g.rotRight = 1
		}
	}
	if event.Type == interaction.KeyUp {
		switch event.Key {
		case "w":
			g.up = 0
		case "a":
			g.left = 0
		case "s":
			g.down = 0
		case "d":
			g.right = 0
		case "q":
			g.rotLeft = 0
		case "e":
			g.rotRight = 0
		}
	}
}

func (g *game) ReceiveMouseEvent(event interaction.MouseEvent) {}

func (g *game) Renderer() dominit.Renderer {
	return g
}

func (g *game) SetOutput(ctx2d *dom.Context2D) {
	g.output = ctx2d
}

func (g *game) Render() {
	g.output.ClearRect(0, 0, 800, 600)

	for i := range g.objects {
		pObj := g.cam.View(g.objects[i])

		// We add (400, 300) here to have (0, 0) be the center of the viewport.
		g.output.DrawImage(g.sprite, 0, 0, 20, 20, int(pObj.X)+400, int(pObj.Y)+300, 20, 20)
	}
}

func (g *game) Scale(availableWidth, availableHeight int) (realWidth, realHeight int, xScale, yScale float64) {
	return 800, 600, 1, 1
}
