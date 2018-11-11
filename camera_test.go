package raytracer

import (
	"math"
	"testing"
)

func TestPixelSize(t *testing.T) {
	cameras := []Camera{
		CreateCamera(200.0, 125.0, math.Pi/2),
		CreateCamera(125.0, 200.0, math.Pi/2),
	}

	expected := []float64{.01, .01}

	for i, x := range expected {
		t.Run(string(i), func(t *testing.T) {
			cameras[i].updatePixelSize()
			if !FloatEqual(cameras[i].PixelSize, x) {
				t.Fail()
				t.Logf("%f != %f", cameras[i].PixelSize, x)
			}
		})
	}
}

func TestRayForPixel(t *testing.T) {
	cameras := []Camera{
		CreateCamera(201, 101, math.Pi/2),
		CreateCamera(201, 101, math.Pi/2),
		CreateCamera(201, 101, math.Pi/2),
	}

	mult, _ := YAxisRotationMatrix(math.Pi / 4).MultiplyMatrix(TranslationMatrix(0, -2, 5))
	cameras[2].Transform = mult

	pxs := []int{100, 0, 100}
	pys := []int{50, 0, 50}
	sqrt2 := math.Sqrt(2) / 2
	expected := []Ray{
		Ray{
			Origin:    Tuple{X: 0, Y: 0, Z: 0, W: 1},
			Direction: Tuple{X: 0, Y: 0, Z: -1.0, W: 0},
		},
		Ray{
			Origin:    Tuple{X: 0, Y: 0, Z: 0, W: 1},
			Direction: Tuple{X: .66519, Y: .33259, Z: -.66851, W: 0},
		},
		Ray{
			Origin:    Tuple{X: 0, Y: 2, Z: -5, W: 1},
			Direction: Tuple{X: sqrt2, Y: 0, Z: -sqrt2, W: 0},
		},
	}

	for i, x := range expected {
		t.Run(string(i), func(t *testing.T) {
			ray, err := cameras[i].RayForPixel(pxs[i], pys[i])
			if err != nil {
				t.Log(err)
				t.Fail()
			}

			if !RayEqual(ray, x) {
				t.Logf("%+v != %+v", ray, x)
				t.Fail()
			}
		})
	}
}

func TestRender(t *testing.T) {
	w := DefaultWorld()
	c := CreateCamera(11, 11, math.Pi/2)
	from := Tuple{X: 0, Y: 0, Z: -5, W: 1}
	to := Tuple{X: 0, Y: 0, Z: 0, W: 1}
	up := Tuple{X: 0, Y: 1, Z: 0, W: 0}
	vt, err := ViewTransform(from, to, up)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	c.Transform = vt

	image, err := c.Render(w)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if !ColorEqual(image.Pixels[5][5], Color{R: .38066, G: .47583, B: .2855}) {
		t.Logf("%+v != %+v", image.Pixels[5][5], Color{R: .38066, G: .47583, B: .2855})
		t.Fail()
	}
}
