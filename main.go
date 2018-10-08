package main 

import (
        "math"
        "log"
        m "github.com/noaoh/raytracer/models"
        c "github.com/noaoh/raytracer/canvas"
)


func main() {
        startPos := m.Tuple{X: 0.0, Y: 0.0, Z: 9.955, W: 1.0}
        canvas := c.CreateCanvas(100, 100)
        i := 0.0
        red := c.Color {R: 1.0}
        r := float64(canvas.Width) * .375
        center := m.Tuple{X: 50, Y: 50, Z: 50, W: 0}
        for {
             if i == 12.0 {
                     break
             }
             log.Println(i)

             t := m.YAxisRotationMatrix(i * (math.Pi/6.0))
             log.Printf("Y Rotation Matrix: %+v\n", t)

             clockPos, err := t.MultiplyTuple(startPos); if err != nil {
                     log.Println(err)
             }

             clockPos.X *= r
             clockPos.Z *= r

             realClockPos, err:= m.Add(center, clockPos); if err != nil {
                     log.Println(err)
             }

             x := int(realClockPos.X) % 100
             z := int(realClockPos.Z) % 100

             if x < 0 {
                     x += 100
             }

             if z < 0 {
                     z += 100
             }

             log.Printf("Position: (%d, %d)\n", x, z)
             canvas.Write(x, z, red)
             i += 1.0
        }
        canvas.WriteFile("test.ppm")
}
