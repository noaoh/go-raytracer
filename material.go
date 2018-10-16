package raytracer

import (
	"fmt"
	"math"
)

type Material struct {
	Color     Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func DefaultMaterial() Material {
	return Material{
		Color:     Color{1, 1, 1},
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200,
	}

}

func MaterialEqual(m1, m2 Material) bool {
	return ColorEqual(m1.Color, m2.Color) &&
		FloatEqual(m1.Ambient, m2.Ambient) &&
		FloatEqual(m1.Diffuse, m2.Diffuse) &&
		FloatEqual(m1.Specular, m2.Specular) &&
		FloatEqual(m1.Shininess, m2.Shininess)
}

func Lighting(m Material, l Light, pos, eyeV, normalV Tuple) (Color, error) {
	errValue := Color{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64}

	if !pos.IsPoint() {
		return errValue, fmt.Errorf("invalid input: pos: %+v", pos)
	}

	if eyeV.IsPoint() || normalV.IsPoint() {
		return errValue, fmt.Errorf("invalid input: eyeV: %+v, normalV: %+v", eyeV, normalV)
	}

	effectiveColor := Hadamard(m.Color, l.Intensity)

	lightV, err := l.Position.Subtract(pos)
	if err != nil {
		return errValue, err
	}

	lightVNorm, err := lightV.Normalize()
	if err != nil {
		return errValue, err
	}

	ambient := effectiveColor.MultiplyScalar(m.Ambient)

	lightDotNormal, err := Dot(lightVNorm, normalV)
	if err != nil {
		return errValue, err
	}

	diffuse := Color{}
	specular := Color{}
	if lightDotNormal < 0 {
		diffuse = Color{0, 0, 0}
		specular = Color{0, 0, 0}
	} else {
		diffuse = effectiveColor.MultiplyScalar(m.Diffuse).MultiplyScalar(lightDotNormal)
		reflectV, err := Reflect(lightVNorm.Negate(), normalV)
		if err != nil {
			return errValue, err
		}

		reflectDot, err := Dot(reflectV, eyeV)
		if err != nil {
			return errValue, err
		}

		reflectDotEye := math.Pow(reflectDot, m.Shininess)

		if reflectDotEye <= 0 {
			specular = Color{0, 0, 0}
		} else {
			specular = l.Intensity.MultiplyScalar(m.Specular).MultiplyScalar(reflectDotEye)
		}
	}

	return ambient.Add(diffuse).Add(specular), nil
}
