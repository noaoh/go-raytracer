package main 

import (
        "fmt"
        "log"
)

type Projectile struct {
        position Tuple
        velocity Tuple
}

type World struct {
        gravity Tuple
        wind Tuple
}

func tick(w World, p Projectile) (Projectile, error) {
        log.Printf("Projectile before tick: +%v\n", p)

        pos, err := Add(p.position, p.velocity); if err != nil {
                return Projectile{}, err
        }

        world_vel, err := Add(w.gravity, w.wind); if err != nil {
                return Projectile{}, err
        }

        vel, err := Add(p.velocity, world_vel); if err != nil {
               return Projectile{}, err
        }

        return Projectile{position: pos, velocity: vel}, nil
}

func main() {
        pos := Tuple{x: 0, y: 1, z: 0, w: 1}

        v, err := Normalize(Tuple{x: 1, y: 1, z: 0, w: 0}); if err != nil {
                log.Println(err)
        }

        p := Projectile{position: pos, velocity: v}
        grav := Tuple{x: 0, y: -0.1, z: 0, w: 0}
        wind := Tuple{x: -0.01, y: 0, z: 0, w: 0}
        world := World{gravity: grav, wind: wind}
        numTicks := 0
        for {
                if (p.position.y <= 0) {
                        break
                }
        
                p, err = tick(world, p); if err != nil {
                        log.Println(err)
                }

                numTicks += 1
        }
        fmt.Printf("%d ticks to hit the ground\n", numTicks)
}
