package raytracer 

type World struct {
    Shapes []Sphere
    Sources []Light
}

func DefaultWorld() World {
        s1 := DefaultSphere()
        s1.Material.Color = Color{R: 0.8, G: 1.0, B: 0.6}
        s1.Material.Diffuse = 0.7
        s1.Material.Specular = 0.2

        s2 := DefaultSphere()
        s2.Transform = ScalingMatrix(0.5, 0.5, 0.5)

        l := Light {Position: Tuple{-10, 10, -10, 1}, Intensity: Color{R: 1.0, G: 1.0, B: 1.0}}

        return World { Shapes: []Sphere{s1, s2}, Sources: []Light{l} }
}

func WorldEqual(w1, w2 World) bool {
        if len(w1.Shapes) != len(w2.Shapes) || len(w1.Sources) != len(w2.Sources) {
                return false
        }

        for i, s := range w1.Shapes {
                if !SphereEqual(s, w2.Shapes[i]) {
                        return false
                }
        }

        for i, l := range w1.Sources {
                if !LightEqual(l, w1.Sources[i]) {
                        return false
                }
        }

        return true
}
