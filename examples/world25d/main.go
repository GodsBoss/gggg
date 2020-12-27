package main

import (
	"math/rand"
	"sort"

	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/dominit"
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
	"github.com/GodsBoss/gggg/pkg/world25d"

	m "github.com/GodsBoss/go-pkg/affinematrix2d"
)

func main() {
	win, _ := dom.GlobalWindow()
	doc, _ := win.Document()
	sprite, _ := doc.CreateImageElement("../assets/facesprite.png")
	shadow, _ := doc.CreateImageElement("../assets/shadow.png")

	dominit.Run(
		&game{
			sprite:  sprite,
			shadow:  shadow,
			cam:     world25d.NewCamera(),
			objects: createRandomObjects(100, -800, -600, 1600, 1200),
		},
	)
	<-make(chan struct{}, 0)
}

func createRandomObjects(count int, minX, minY, maxX, maxY int) []object {
	objects := make(objects, count)
	for i := 0; i < count; i++ {
		objects[i] = object{
			X: float64(rand.Intn(maxX-minX) + minX),
			Y: float64(rand.Intn(maxY-minY) + minY),
			Z: float64(rand.Intn(20)),
		}
	}
	return objects
}

type game struct {
	sprite *dom.Image
	shadow *dom.Image
	output *dom.Context2D

	cam     world25d.Camera
	objects objects

	left  int
	right int
	down  int
	up    int

	rotLeft  int
	rotRight int

	higher int
	lower  int

	angleHorizon int
	angleDown    int
}

func (g *game) TicksPerSecond() int {
	return 40
}

func (g *game) Logic() dominit.Logic {
	return g
}

func (g *game) Tick(ms int) {
	for i := range g.objects {
		g.objects[i].Tick(ms)
	}

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

	zSpeed := 2.5
	_ = g.cam.SetHeight(g.cam.Height() + float64(g.higher-g.lower)*zSpeed)

	angleSpeed := 0.1
	_ = g.cam.SetAngle(g.cam.Angle() + float64(g.angleHorizon-g.angleDown)*angleSpeed)
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
		case "r":
			g.higher = 1
		case "f":
			g.lower = 1
		case "g":
			g.angleHorizon = 1
		case "t":
			g.angleDown = 1
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
		case "r":
			g.higher = 0
		case "f":
			g.lower = 0
		case "g":
			g.angleHorizon = 0
		case "t":
			g.angleDown = 0
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

	bg, _ := canvas2drendering.MonochromeBackground("#efe")
	bg.Render(g.output)

	pObjs := world25d.ViewObjects(g.cam, g.objects.Shadows()...)
	sort.Sort(pObjs)

	for i := range pObjs {
		// We add (400, 300) here to have (0, 0) be the center of the viewport.
		// We add (-10, -5) here, because that is the bottom center of the objects.
		g.output.DrawImage(g.shadow, 0, 0, 20, 5, int(pObjs[i].X)+400-10, int(pObjs[i].ComputedY())+300-5, 20, 5)
	}

	pObjs = world25d.ViewObjects(g.cam, g.objects.ToWorld25dObjects()...)
	sort.Sort(pObjs)

	for i := range pObjs {
		// We add (400, 300) here to have (0, 0) be the center of the viewport.
		// We add (-10, -20) here, because that is the bottom center of the objects.
		g.output.DrawImage(g.sprite, 1, 1, 20, 20, int(pObjs[i].X)+400-10, int(pObjs[i].ComputedY())+300-20, 20, 20)
	}
}

func (g *game) Scale(availableWidth, availableHeight int) (realWidth, realHeight int, xScale, yScale float64) {
	return 800, 600, 1, 1
}

type object struct {
	X      float64
	Y      float64
	Z      float64
	ZSpeed float64
}

func (obj *object) Tick(ms int) {
	obj.ZSpeed += gravity
	obj.Z += obj.ZSpeed
	if obj.Z < 0 {
		obj.Z = 0
		obj.ZSpeed = jumpSpeed
	}
}

func (obj *object) ToWorld25dObject() world25d.Object {
	return world25d.Object{
		X: obj.X,
		Y: obj.Y,
		Z: obj.Z,
	}
}

func (obj *object) Shadow() world25d.Object {
	return world25d.Object{
		X: obj.X,
		Y: obj.Y,
		Z: 0,
	}
}

const gravity float64 = -0.05
const jumpSpeed float64 = 3.0

type objects []object

func (objs objects) ToWorld25dObjects() []world25d.Object {
	result := make([]world25d.Object, len(objs))
	for i := range objs {
		result[i] = objs[i].ToWorld25dObject()
	}
	return result
}

func (objs objects) Shadows() []world25d.Object {
	result := make([]world25d.Object, len(objs))
	for i := range objs {
		result[i] = objs[i].Shadow()
	}
	return result
}
