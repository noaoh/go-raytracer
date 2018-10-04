package main 

import (
        "fmt"
        "math"
        "log"
        m "github.com/noaoh/raytracer/models"
        c "github.com/noaoh/raytracer/canvas"
)

type Projectile struct {
        position m.Tuple
        velocity m.Tuple
}

type World struct {
        gravity m.Tuple
        wind m.Tuple
}

func tick(w World, p Projectile) (Projectile, error) {

        pos, err := m.Add(p.position, p.velocity); if err != nil {
                return Projectile{}, err
        }

        world_vel, err := m.Add(w.gravity, w.wind); if err != nil {
                return Projectile{}, err
        }

        vel, err := m.Add(p.velocity, world_vel); if err != nil {
               return Projectile{}, err
        }

        return Projectile{position: pos, velocity: vel}, nil
}

func main() {
        pos := m.Tuple{X: 0, Y: 1, Z: 0, W: 1}

        norm, err := m.Normalize(m.Tuple{X: 1, Y: 1.8, Z: 0, W: 0}) 

        if err != nil {
                log.Println(err)
        }

        v := norm.MultiplyFloat(11.25)
        v.W = 0

        p := Projectile{position: pos, velocity: v}
        grav := m.Tuple{X: 0, Y: -0.1, Z: 0, W: 0}
        wind := m.Tuple{X: -0.01, Y: 0, Z: 0, W: 0}
        world := World{gravity: grav, wind: wind}
        canvas := c.CreateCanvas(900, 550)
        red := c.Color{R: 1.0, G: 0.0, B: 0.0}
        numTicks := 0
        for {
                if p.position.Y <= 0 {
                        break
                }
                fmt.Printf("%+v\n", p)
                x := int(math.Round(math.Abs(p.position.X)))
                y := int(math.Round(math.Abs(float64(canvas.Height) - p.position.Y)))
                canvas.Write(x, y, red)
        
                p, err = tick(world, p); if err != nil {
                        log.Println(err)
                }

                numTicks += 1
        }
        fmt.Printf("%d ticks to hit the ground\n", numTicks)
        canvas.WriteFile("test.ppm")
}
