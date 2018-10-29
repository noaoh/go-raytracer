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
        canvasPixels := 500.0
        pixelSize := wallSize / canvasPixels
        half := wallSize / 2.0
        rayOrigin := r.Tuple{X: 0.0, Y: 0.0, Z: -5.0, W: 1.0}
        canvas := r.CreateCanvas(int(canvasPixels), int(canvasPixels))
        noHit := r.Intersection{T: math.MaxFloat64, Obj: r.Sphere {}}
        sphere := r.DefaultSphere()
        sphere.Transform = r.IdentityMatrix(4)
        sphere.Material.Color = r.Color{R: 1.0, G: 0.2, B: 1.0}
        lightPos := r.Tuple{X: -10.0, Y: 10.0, Z: -10.0, W: 1.0}
        lightColor := r.Color{R: 1.0, G: 1.0, B: 1.0}
        light := r.Light{Intensity: lightColor, Position: lightPos}

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

                        ray.Direction, err = ray.Direction.Normalize()
                        if err != nil {
                                log.Print(err)
                                os.Exit(1)
                        }

                        xs, err := sphere.Intersect(ray); if err != nil {
                                log.Print(err)
                                os.Exit(1)
                        }

                        hit := r.Hit(xs)
                        if !r.IntersectionEqual(hit, noHit) {
                                point, err := ray.Position(hit.T); if err != nil {
                                        log.Print(err)
                                        os.Exit(1)
                                }

                                normal, err := hit.Obj.NormalAt(point); if err != nil {
                                        log.Print(err)
                                        os.Exit(1)
                                }
                                eye := ray.Direction
                                color, err := r.Lighting(hit.Obj.Material, light, point, eye, normal)
                                if err != nil {
                                        log.Print(err)
                                        os.Exit(1)
                                }

                                canvas.Write(x, y, color)
                        }
                }
        }
        canvas.WriteFile("test.ppm")
}
