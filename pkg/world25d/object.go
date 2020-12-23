package world25d

type Object struct {
	X        float64
	Y        float64
	Z        float64
	Rotation float64
}

func ViewObjects(cam Camera, objs ...Object) PerceivedObjects {
	result := make(PerceivedObjects, len(objs))
	for i := range objs {
		result[i] = cam.View(objs[i])
	}
	return result
}

type PerceivedObject struct {
	X        float64
	Y        float64
	Rotation float64
}

type PerceivedObjects []PerceivedObject

func (po PerceivedObjects) Len() int {
	return len(po)
}

func (po PerceivedObjects) Less(i, j int) bool {
	return po[i].Y < po[j].Y
}

func (po PerceivedObjects) Swap(i, j int) {
	po[i], po[j] = po[j], po[i]
}
