package main 

import (
        "fmt"
        "math"
        "log"
        t "github.com/noaoh/raytracer/models/tuple"
        c "github.com/noaoh/raytracer/canvas"
)

type Projectile struct {
        position t.Tuple
        velocity t.Tuple
}

type World struct {
        gravity t.Tuple
        wind t.Tuple
}

func tick(w World, p Projectile) (Projectile, error) {

        pos, err := t.Add(p.position, p.velocity); if err != nil {
                return Projectile{}, err
        }

        world_vel, err := t.Add(w.gravity, w.wind); if err != nil {
                return Projectile{}, err
        }

        vel, err := t.Add(p.velocity, world_vel); if err != nil {
               return Projectile{}, err
        }

        return Projectile{position: pos, velocity: vel}, nil
}

func main() {
        pos := t.Tuple{X: 0, Y: 1, Z: 0, W: 1}

        norm, err := t.Normalize(t.Tuple{X: 1, Y: 1.8, Z: 0, W: 0}) 

        if err != nil {
                log.Println(err)
        }

        v := t.Multiply(norm, 11.25)
        v.W = 0

        p := Projectile{position: pos, velocity: v}
        grav := t.Tuple{X: 0, Y: -0.1, Z: 0, W: 0}
        wind := t.Tuple{X: -0.01, Y: 0, Z: 0, W: 0}
        world := World{gravity: grav, wind: wind}
        canvas := c.Canvas(550, 900)
        red := c.Color{R: 1.0, G: 0.0, B: 0.0}
        numTicks := 0
        for {
                fmt.Printf("%+v\n", p)
                x := int(math.Round(math.Abs(p.position.X)))
                y := int(math.Round(math.Abs(p.position.Y)))
                canvas.Write(x, y, red)
                if (p.position.Y <= 0) {
                        break
                }
        
                p, err = tick(world, p); if err != nil {
                        log.Println(err)
                }

                numTicks += 1
        }
        fmt.Printf("%d ticks to hit the ground\n", numTicks)
        canvas.WriteFile("test.ppm")
}
