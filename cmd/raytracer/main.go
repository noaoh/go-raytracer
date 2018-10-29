package main 

import (
        "log"
        "math"
        "os"
        r "github.com/noaoh/raytracer"
)

func main() {
        floor := r.DefaultSphere()
        floor.Transform = r.ScalingMatrix(10, .01, 10)
        floor.Material = r.DefaultMaterial()
        floor.Material.Color = r.Color{R: 1, G: .9, B: .9}
        floor.Material.Specular = 0.0

        leftWall := r.DefaultSphere()
        a, _ := r.TranslationMatrix(0, 0, 5).MultiplyMatrix(r.YAxisRotationMatrix(-math.Pi/4))
        b, _ := a.MultiplyMatrix(r.XAxisRotationMatrix(math.Pi/2))
        c, _ := b.MultiplyMatrix(r.ScalingMatrix(10, .01, 10))
        leftWall.Transform = c
        leftWall.Material = floor.Material

        middle := r.DefaultSphere()
        middle.Transform = r.TranslationMatrix(-.5, 1, .5)
        middle.Material = r.DefaultMaterial()
        middle.Material.Color = r.Color{R: .1, G: 1, B: .5}
        middle.Material.Diffuse = .7
        middle.Material.Specular = .3

        right := r.DefaultSphere()
        right.Transform, _ = r.TranslationMatrix(1.5, .5, -.5).MultiplyMatrix(r.ScalingMatrix(.5, .5, .5))
        right.Material = r.DefaultMaterial()
        right.Material.Color = r.Color{R: .5, G: 1, B: 0.1}
        right.Material.Diffuse = .7
        right.Material.Specular = .3
         
        world := r.DefaultWorld()
        world.Shapes = []r.Sphere {floor, leftWall, middle, right}

        camera := r.CreateCamera(1000.0, 500.0, math.Pi/3)
        var err error
        camera.Transform, err = r.ViewTransform( r.Tuple {X: 0, Y: 1.5, Z: -5, W: 1}, r.Tuple {X: 0, Y: 1, Z: 0, W: 1}, r.Tuple {X: 0, Y: 1, Z: 0, W: 0})

        if err != nil {
                log.Print("error from r.ViewTransform")
                log.Print(err)
                os.Exit(1)
        }

        var canvas r.Canvas
        canvas, err = camera.Render(world); if err != nil {
                log.Print("error from camera.Render")
                log.Print(err)
                os.Exit(1)
        }

        canvas.WriteFile("test.ppm")
}
