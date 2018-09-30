package canvas

import (
        "testing"
)

func TestAdd(t *testing.T) {
        c1 := Color {R: 0.9, G: 0.6, B: 0.75}
        c2 := Color {R: 0.7, G: 0.1, B: 0.25}
        e := Color {R: 1.6, G: 0.7, B: 1.0}
        a := Add(c1, c2); if a != e {
                t.Logf("%+v != %+v", a, e)
                t.Fail()
        }
}

func TestSubtract(t *testing.T) {
        c1 := Color {R: 0.9, G: 0.6, B: 0.75}
        c2 := Color {R: 0.7, G: 0.1, B: 0.25}
        e := Color {R: 0.2, G: 0.5, B: 0.50}
        a := Subtract(c1, c2); if a != e {
                t.Logf("%+v != %+v", a, e)
                t.Fail()
        }
}

func TestMultiplyScalar(t *testing.T) {
        i := 2.0
        c := Color {R: 0.2, G: 0.3, B: 0.4}
        e := Color {R: 0.4, G: 0.6, B: 0.8}
        a := MultiplyScalar(c, i); if a != e {
                t.Logf("%+v != %+v", a, e)
                t.Fail()
        }
}

func TestHadamard(t *testing.T) {
        c1 := Color {R: 1.0, G: 0.2, B: 0.4}
        c2 := Color {R: 0.9, G: 1.0, B: 0.1}
        e := Color {R: 0.9, G: 0.2, B: 0.04}
        a := Hadamard(c1, c2); if a != e {
                t.Logf("%+v != %+v", a, e)
                t.Fail()
        }
}


