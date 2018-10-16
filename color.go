package raytracer

type Color struct {
	R float64
	G float64
	B float64
}

func (x Color) Add(y Color) Color {
	return Color{R: x.R + y.R, G: x.G + y.G, B: x.B + y.B}
}

func (x Color) Subtract(y Color) Color {
	return Color{R: x.R - y.R, G: x.G - y.G, B: x.B - y.B}
}

func (x Color) MultiplyScalar(f float64) Color {
	return Color{R: x.R * f, G: x.G * f, B: x.B * f}
}

func Hadamard(x, y Color) Color {
	return Color{R: x.R * y.R, G: x.G * y.G, B: x.B * y.B}
}

func ColorEqual(c1, c2 Color) bool {
	return FloatEqual(c1.R, c2.R) &&
		FloatEqual(c1.G, c2.G) &&
		FloatEqual(c1.B, c2.B)
}
