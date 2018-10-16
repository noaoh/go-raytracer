package raytracer

type Light struct {
	Intensity Color
	Position  Tuple
}

func LightEqual(l1, l2 Light) bool {
	return ColorEqual(l1.Intensity, l2.Intensity) &&
		TupleEqual(l1.Position, l2.Position)
}
