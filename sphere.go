package raytracer

type Sphere struct {
	Radius float64
	Origin Tuple
        Transform Matrix
}

func SphereEqual(s1, s2 Sphere) bool {
	return FloatEqual(s1.Radius, s2.Radius) && 
        TupleEqual(s1.Origin, s2.Origin) && 
        MatrixEqual(s1.Transform, s2.Transform)
}
