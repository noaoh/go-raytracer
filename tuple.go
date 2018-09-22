package main 

import (
        "math"
        "errors"
)

type Tuple struct  {
        x float64 
        y float64
        z float64
        w float64
}

func (t Tuple) isPoint() bool {
        return t.w == 1.0;
}

func (t Tuple) isVector() bool {
        return t.w == 0.0;
}

func Add(a, b Tuple) (Tuple, error) {
        if (a.isPoint() && b.isPoint()) {
                return Tuple {}, errors.New("Can not add a point to a point\n")
        }

        return Tuple{x: a.x + b.x, y: a.y + b.y, z: a.z + b.z, w: a.w + b.w}, nil
}

func Subtract(a, b Tuple) (Tuple, error) {
        if (a.isVector() && b.isPoint()) {
                return Tuple{}, errors.New("Can not subtract a vector from a point\n")
        }


        return Tuple{x: a.x - b.x, y: a.y - b.y, z: a.z - b.z, w: a.w - b.w}, nil

}

func Negate(a Tuple) Tuple {
        return Tuple{x: a.x * -1, y: a.y * -1, z: a.z * -1, w: a.w * -1}
}

func Multiply(a Tuple, f float64) Tuple {
        return Tuple{x: a.x * f, y: a.y * f, z: a.z * f, w: a.w * f}
}

func Magnitude(a Tuple) float64 {
        sum := math.Pow(a.x, 2) + math.Pow(a.y, 2) + math.Pow(a.z, 2) + math.Pow(a.w, 2)
        return math.Sqrt(sum)
}

func Normalize(a Tuple) (Tuple, error) {
        m := Magnitude(a)
        if m == 0 {
                return Tuple {}, errors.New("Can not divide by zero\n")
        }
        return Tuple { x: a.x / m, y: a.y / m, z: a.z / m, w: a.w / m }, nil
}

func Dot(t1, t2 Tuple) (float64, error) {
        if (t1.isPoint() || t2.isPoint()) {
                return 1.0, errors.New("Can not perform dot product on points\n")
        }

        return (t1.x * t2.x) + (t1.y * t2.y) + (t1.z * t2.z) + (t1.w * t2.w), nil
}

func Cross(t1, t2 Tuple) (Tuple, error) {
        if (t1.isPoint() || t2.isPoint()) {
                return Tuple {}, errors.New("Can not perform cross product on points\n")
        }

        zy := t1.y * t2.z - t2.z * t1.y
        zx := t1.z * t2.x - t1.x * t1.z
        xy := t1.x * t2.y - t1.y * t2.x
        return Tuple { x: zy, y: zx, z: xy, w: t1.w}, nil
}


