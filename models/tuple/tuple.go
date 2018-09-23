package tuple 

import (
        "math"
        "errors"
)

type Tuple struct  {
        X float64 
        Y float64
        Z float64
        W float64
}

func (t Tuple) isPoint() bool {
        return t.W == 1.0;
}

func (t Tuple) isVector() bool {
        return t.W == 0.0;
}

func Add(a, b Tuple) (Tuple, error) {
        if (a.isPoint() && b.isPoint()) {
                return Tuple {}, errors.New("Can not add a point to a point\n")
        }

        return Tuple{X: a.X + b.X, Y: a.Y + b.Y, Z: a.Z + b.Z, W: a.W + b.W}, nil
}

func Subtract(a, b Tuple) (Tuple, error) {
        if (a.isVector() && b.isPoint()) {
                return Tuple{}, errors.New("Can not subtract a vector from a point\n")
        }

        return Tuple{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z, W: a.W - b.W}, nil
}

func Negate(a Tuple) Tuple {
        return Tuple{X: a.X * -1, Y: a.Y * -1, Z: a.Z * -1, W: a.W * -1}
}

func Multiply(a Tuple, f float64) Tuple {
        return Tuple{X: a.X * f, Y: a.Y * f, Z: a.Z * f, W: a.W * f}
}

func Magnitude(a Tuple) float64 {
        sum := math.Pow(a.X, 2) + math.Pow(a.Y, 2) + math.Pow(a.Z, 2) + math.Pow(a.W, 2)
        return math.Sqrt(sum)
}

func Normalize(a Tuple) (Tuple, error) {
        m := Magnitude(a)
        if m == 0 {
                return Tuple {}, errors.New("Can not divide bY Zero\n")
        }
        return Tuple { X: a.X / m, Y: a.Y / m, Z: a.Z / m, W: a.W / m }, nil
}

func Dot(t1, t2 Tuple) (float64, error) {
        if (t1.isPoint() || t2.isPoint()) {
                return 1.0, errors.New("Can not perform dot product on points\n")
        }

        return (t1.X * t2.X) + (t1.Y * t2.Y) + (t1.Z * t2.Z) + (t1.W * t2.W), nil
}

func Cross(t1, t2 Tuple) (Tuple, error) {
        if (t1.isPoint() || t2.isPoint()) {
                return Tuple {}, errors.New("Can not perform cross product on points\n")
        }

        zy := t1.Y * t2.Z - t2.Z * t1.Y
        zx := t1.Z * t2.X - t1.X * t1.Z
        xy := t1.X * t2.Y - t1.Y * t2.X
        return Tuple { X: zy, Y: zx, Z: xy, W: t1.W}, nil
}


