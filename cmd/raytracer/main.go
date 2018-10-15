package main 

import (
        "log"
        "math"
        "os"
        r "github.com/noaoh/raytracer"
)


func main() {
        wallZ := 10.0
        wallSize := 7.0
        canvasPixels := 100.0
        pixelSize := wallSize / canvasPixels
        half := wallSize / 2.0
        rayOrigin := r.Tuple{X: 0.0, Y: 0.0, Z: -5.0, W: 1.0}
        canvas := r.CreateCanvas(int(canvasPixels), int(canvasPixels))
        red := r.Color {R: 1.0}
        noHit := r.Intersection{T: math.MaxFloat64, Obj: r.Sphere {}}
        sphereTransform, _ := r.ShearingMatrix(1, 0, 0, 1, 0, 0).MultiplyMatrix(r.ScalingMatrix(0.5, 1, 1))
        sphere := r.Sphere{
                Radius: 1.0,
                Origin: r.Tuple{X: 0.0, Y: 0.0, Z: 0.0, W: 1.0},
                Transform: sphereTransform,
        }

        for y := 0; y < canvas.Height; y++ {
                worldY := half - pixelSize * float64(y)
                for x := 0; x < canvas.Width; x++ {
                        worldX := -1 * half + pixelSize * float64(x)
                        pos := r.Tuple{X: worldX, Y: worldY, Z: wallZ, W: 1.0}

                        wallPos, err := pos.Add(rayOrigin.MultiplyFloat(-1))
                        if err != nil {
                                log.Print(err)
                                os.Exit(1)
                        }

                        normWallPos, err := wallPos.Normalize(); if err != nil {
                                log.Print(err)
                                os.Exit(1)
                        }

                        ray, err := r.CreateRay(rayOrigin, normWallPos) 
                        if err != nil {
                                log.Print(err)
                                os.Exit(1)
                        }

                        xs, err := r.Intersect(sphere, ray); if err != nil {
                                log.Print(err)
                                os.Exit(1)
                        }

                        hit := r.Hit(xs)
                        if !r.IntersectionEqual(hit, noHit) {
                                canvas.Write(x, y, red)
                        }
                }
        }
        canvas.WriteFile("test.ppm")
}
