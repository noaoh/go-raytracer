package raytracer

type Sphere struct {
	Radius    float64
	Origin    Tuple
	Transform Matrix
	Material  Material
}

func SphereEqual(s1, s2 Sphere) bool {
	return FloatEqual(s1.Radius, s2.Radius) &&
		TupleEqual(s1.Origin, s2.Origin) &&
		MatrixEqual(s1.Transform, s2.Transform)
}

func DefaultSphere() Sphere {
	return Sphere{
		Radius:    1.0,
		Origin:    Tuple{0, 0, 0, 1},
		Transform: IdentityMatrix(4),
		Material:  DefaultMaterial(),
	}
}

func (s Sphere) NormalAt(t Tuple) (Tuple, error) {
	inv, err := s.Transform.Inverse()
	if err != nil {
		return Tuple{}, err
	}

	objPoint, err := inv.MultiplyTuple(t)
	if err != nil {
		return Tuple{}, err
	}

	objNorm, err := objPoint.Subtract(s.Origin)
	if err != nil {
		return Tuple{}, err
	}

	worldPoint, err := inv.Transpose().MultiplyTuple(objNorm)
	if err != nil {
		return Tuple{}, err
	}

	worldPoint.W = 0

	worldNorm, err := worldPoint.Normalize()
	if err != nil {
		return Tuple{}, err
	}

	return worldNorm, nil
}
