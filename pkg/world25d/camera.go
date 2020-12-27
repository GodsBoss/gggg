package world25d

import (
	"fmt"
	"math"

	m "github.com/GodsBoss/go-pkg/affinematrix2d"
)

type Camera interface {
	// View returns an object as seen through the camera.
	View(obj Object) PerceivedObject

	Position() (x, y float64)
	SetPosition(x, y float64)

	Height() float64

	// SetHeight returns an error if a negative height is to be set.
	SetHeight(h float64) error

	Rotation() float64
	SetRotation(r float64)

	Angle() float64

	// SetAngle returns an error if angle is beyond PI/2 or less than 0.
	SetAngle(angle float64) error
}

// NewCamera creates a new camera with somewhat decent default values.
// See CameraDefault* constants.
func NewCamera() Camera {
	cam := &camera{
		x:        CameraDefaultX,
		y:        CameraDefaultY,
		z:        CameraDefaultHeight,
		rotation: CameraDefaultRotation,
		angle:    CameraDefaultAngle,
	}
	cam.calculateTransformation()
	return cam
}

const (
	CameraDefaultX        float64 = 0
	CameraDefaultY        float64 = 0
	CameraDefaultHeight   float64 = 100
	CameraDefaultRotation float64 = 0
	CameraDefaultAngle    float64 = 0
)

type camera struct {
	// x is the camera's X position in the world.
	x float64

	// y is the camera's Y position in the world.
	y float64

	// z is the distance of the camera to the "ground".
	z float64

	// rotation is the camera's rotation. 0 is looking straight into negative Y ("up"),
	// PI/2 is looking straight into positive X ("right"), and so on. Basically,
	// it looks clockwise.
	rotation float64

	// angle is the angle of the camera. Looking straight down is 0. Looking parallel
	// to the ground would be PI/2, but that is not possible.
	angle float64

	// m is the matrix used for transformations.
	m m.Transformation
}

func (cam *camera) View(obj Object) PerceivedObject {
	pos := cam.m.Transform(m.VectorFromCartesian(obj.X, obj.Y))

	zOffset := -math.Sin(cam.angle) * obj.Z * 100.0 / cam.z

	return PerceivedObject{
		X:             pos.X(),
		Y:             pos.Y(),
		YHeightOffset: zOffset,
		Rotation:      obj.Rotation + cam.rotation,
		Scale:         100.0 / cam.z,
	}
}

func (cam *camera) Position() (x, y float64) {
	return cam.x, cam.y
}

func (cam *camera) SetPosition(x, y float64) {
	cam.x = x
	cam.y = y
	cam.calculateTransformation()
}

func (cam *camera) Height() float64 {
	return cam.z
}

func (cam *camera) SetHeight(h float64) error {
	if h <= 0 {
		return fmt.Errorf("height must not be < 0")
	}
	cam.z = h
	cam.calculateTransformation()
	return nil
}

func (cam *camera) Rotation() float64 {
	return cam.rotation
}

func (cam *camera) SetRotation(r float64) {
	cam.rotation = r
	cam.calculateTransformation()
}

func (cam *camera) Angle() float64 {
	return cam.angle
}

func (cam *camera) SetAngle(angle float64) error {
	if angle < 0 || angle >= math.Pi/2.0 {
		return fmt.Errorf("angle must be in the interval [0, PI/2)")
	}
	cam.angle = angle
	cam.calculateTransformation()
	return nil
}

func (cam *camera) calculateTransformation() {
	cam.m = m.Combine(
		m.Scale(100.0/cam.z, 100.0/cam.z),
		m.Scale(1, math.Cos(cam.angle)),
		m.Rotation(-cam.rotation),
		m.Translation(-cam.x, -cam.y),
	)
}
