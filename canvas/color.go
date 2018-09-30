package canvas 

type Color struct {
        R float64
        G float64
        B float64
}

func Add(x, y Color) Color {
        return Color {R: x.R + y.R, G: x.G + y.G, B: x.B + y.B}
}

func Subtract(x, y Color) Color {
        return Color {R: x.R - y.R, G: x.G - y.G, B: x.B - y.B}
}

func MultiplyScalar(x Color, f float64) Color {
        return Color {R: x.R * f, G: x.G * f, B: x.B * f}
}

func Hadamard(x, y Color) Color {
        return Color { R: x.R * y.R, G: x.G * y.G, B: x.B * y.B }
}
