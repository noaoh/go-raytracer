package tuple

import (
        "math"
        "testing"
)

func TestIsVector(t *testing.T) {
        expected := {true, true, false, false}
        actual := {
                Tuple {X: 0, Y: 0, Z: 0, W: 1}, 
                Tuple {X: 1, Y: 1, Z: 1, W: 1}, 
                Tuple {X: 0, Y: 0, Z: 0, W: 0}, 
                Tuple {X: 1, Y: 1, Z: 1, W: 0}, 
        }

        for i, x := range actual {
                t.Run(string(i), func(t *testing.T) {
                        if x.isVector() != expected {
                                t.Logf("Failed isVector for: %+v\n", x)
                                t.Fail()
                        }
                }
        }
}

func TestIsPoint(t *testing.T) {
        expected := {true, true, false, false}
        actual := {
                Tuple {X: 0, Y: 0, Z: 0, W: 0}, 
                Tuple {X: 1, Y: 1, Z: 1, W: 0}, 
                Tuple {X: 0, Y: 0, Z: 0, W: 1}, 
                Tuple {X: 1, Y: 1, Z: 1, W: 1}, 
        }

        for i, x := range actual {
                t.Run(string(i), func(t *testing.T) {
                        if x.isPoint() != expected {
                                t.Logf("Failed isPoint for: %+v\n", x)
                                t.Fail()
                        }
                }
        }
}

func TestAdd(t *testing.T) {
        x := Tuple {X: 3, Y: -2, Z: 5, W: 1}
        y := Tuple {X: -2, Y: 3, Z: 1, W: 0}
        a := Add(a1, a2)
        e := Tuple {X: 1, Y: 1, Z: 6, W: 1}
        if a != e {
                t.Logf("Add(%+v, %+v) != %+v\n", x, y, e)
                t.Fail()
        }
}

func TestSubtract(t *testing.T) {
        expected := {
                Tuple {X: -2, Y: -4, Z: -6, W: 0},
                Tuple {X: -2, Y: -4, Z: -6, W: 1},
                Tuple {X: -2, Y: -4, Z: -6, W: 0},
        }

        actual := {
                { Tuple {X: 3, Y: 2, Z: 1, W: 1}, Tuple {X: 5, Y: 6, Z: 7, W: 1} },
                { Tuple {X: 3, Y: 2, Z: 1, W: 1}, Tuple {X: 5, Y: 6, Z: 7, W: 0} },
                { Tuple {X: 3, Y: 2, Z: 1, W: 0}, Tuple {X: 5, Y: 6, Z: 7, W: 0} },
        }

        for i, e := range expected {
                t.Run(string(i), func (t *testing.T) {
                        x := actual[i][0]
                        y := actual[i][1]
                        s, _ := Subtract(x, y)
                        if s != e {
                                t.Logf("Subtract(%+v,%+v) != %+v\n", x, y, e)
                                t.Fail()
                        }
                }
        }
}

func TestNegate(t *testing.T) {
        e := Tuple {X: -1, Y: 2, Z: -3, W: 4}
        a := Tuple {X: 1, Y: -2, Z: 3, W: -4}
        if Negate(a) != e {
                t.Logf("Negate(%+v) != %+v", a, e)
        }
}

func TestMultiply(t *testing.T) {
        f := {3.5, 0.5}
        a := Tuple {X: 1, Y: -2, Z: 3, W: -4}
        e := { Tuple {X: 3.5, Y: -7, Z: 10.5, W: -14}, 
        Tuple {X: 0.5, Y: -1.0, 1.5, -2.0},
        }

        for i, x := range e {
                t.Run(string(i), func(t *testing.T) {
                        if Multiply(a, f[i]) != x {
                                t.Logf("Multiply(%+v, %f) != %+v\n", a, f[i], x)
                                t.Fail()
                        }
                }
        }
}

func TestMagnitude(t *testing.T) {
        a := {
                Tuple {X: 1, Y: 0, Z: 0, W: 0},
                Tuple {X: 0, Y: 1, Z: 0, W: 0},
                Tuple {X: 0, Y: 0, Z: 1, W: 0},
                Tuple {X: 1, Y: 2, Z: 3, W: 0},
                Tuple {X: -1, Y: -2, Z: -3, W: 0},
        }
        e := {1, 1, 1, math.Sqrt(14), math.Sqrt(14)}

        for i, x := range e {
                t.Run(string(i), func (t *testing.T) {
                        if Magnitude(a[i]) != x {
                                t.Logf("Magnitude(%+v) != %f\n", a[i], x)
                                t.Fail()
                        }
                }
        }
}

func TestNormalize(t *testing.T) {
        a := Tuple {X: 1, Y: 2, Z: 3, W: 0}
        e := Tuple {X: 1/math.Sqrt(14), Y: 2/math.Sqrt(14), Z: 3/math.Sqrt(14), W: 0}
        n, _ := Normalize(a)
        if n != e {
                t.Logf("Normalize(%+v) != %+v", a, e)
                t.Fail()
        }
}

func TestDot(t *testing.T) {
        e := 20.0
        x := Tuple {X: 1, Y: 2, Z: 3, W: 0}
        y := Tuple {X: 2, Y: 3, Z: 4, W: 0}
        d, _ := Dot(x, y)
        if d != e {
                t.Logf("Dot(%+v, %+v) != %f", x, y, e)
                t.Fail()
        }
}

