package sprint

type Circle struct {
	Radius   float32
	Diameter float32
	Area     float32
	Perimeter float32
}
//test
func GetCircle(r float32) Circle {
	const pi = 3.14

	diameter := 2 * r
	area := pi * r * r
	perimeter := 2 * pi * r

	return Circle{
		Radius:   r,
		Diameter: diameter,
		Area:     area,
		Perimeter: perimeter,
	}
}