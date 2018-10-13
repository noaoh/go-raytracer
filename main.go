package main 

import (
        "log"
        "math"
        "os"
        m "github.com/noaoh/raytracer/models"
        c "github.com/noaoh/raytracer/canvas"
)


func main() {
        wallZ := 10.0
        wallSize := 7.0
        canvasPixels := 100.0
        pixelSize := wallSize / canvasPixels
        half := wallSize / 2.0
        rayOrigin := m.Tuple{X: 0.0, Y: 0.0, Z: -5.0, W: 1.0}
        canvas := c.CreateCanvas(int(canvasPixels), int(canvasPixels))
        red := c.Color {R: 1.0}
        noHit := m.Intersection{T: math.MaxFloat64, Obj: m.Sphere {}}
        sphereTransform, _ := m.ShearingMatrix(1, 0, 0, 1, 0, 0).MultiplyMatrix(m.ScalingMatrix(0.5, 1, 1))
        sphere := m.Sphere{
                Radius: 1.0,
                Origin: m.Tuple{X: 0.0, Y: 0.0, Z: 0.0, W: 1.0},
                Transform: sphereTransform,
        }

        for y := 0; y < canvas.Height; y++ {
                worldY := half - pixelSize * float64(y)
                for x := 0; x < canvas.Width; x++ {
                        worldX := -1 * half + pixelSize * float64(x)
                        pos := m.Tuple{X: worldX, Y: worldY, Z: wallZ, W: 1.0}

                        wallPos, err := m.Add(pos, rayOrigin.MultiplyFloat(-1))
                        if err != nil {
                                log.Print(err)
                                os.Exit(1)
                        }

                        normWallPos, err := m.Normalize(wallPos); if err != nil {
                                log.Print(err)
                                os.Exit(1)
                        }

                        r, err := m.CreateRay(rayOrigin, normWallPos) 
                        if err != nil {
                                log.Print(err)
                                os.Exit(1)
                        }

                        xs, err := m.Intersect(sphere, r); if err != nil {
                                log.Print(err)
                                os.Exit(1)
                        }

                        hit := m.Hit(xs)
                        if !m.IntersectionEqual(hit, noHit) {
                                canvas.Write(x, y, red)
                        }
                }
        }
        canvas.WriteFile("test.ppm")
}
