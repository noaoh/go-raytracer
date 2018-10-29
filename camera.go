package raytracer

import (
        "math"
)

type Camera struct {
        HSize float64
        VSize float64
        FOV float64
        HalfWidth float64
        HalfHeight float64
        PixelSize float64
        Transform Matrix
}

func CreateCamera(width, height, fov float64) Camera {
        c := Camera{HSize: width, VSize: height, FOV: fov, Transform: IdentityMatrix(4)}
        c.updatePixelSize()
        return c

}

func (c *Camera) updatePixelSize() {
        halfView := math.Tan(c.FOV / 2)
        aspect := c.HSize / c.VSize

        if aspect >= 1 {
                c.HalfWidth = halfView
                c.HalfHeight = halfView / aspect
        } else {
                c.HalfWidth = halfView * aspect
                c.HalfHeight = halfView
        }

        c.PixelSize = (c.HalfWidth * 2) / c.HSize
}

func (c Camera) RayForPixel(px, py int) (Ray, error) {
        xoffset := (float64(px) + .5) * c.PixelSize
        yoffset := (float64(py) + .5) * c.PixelSize

        worldX := c.HalfWidth - xoffset
        worldY := c.HalfHeight - yoffset

        invCamera, err := c.Transform.Inverse(); if err != nil {
                return Ray{}, err
        }

        pixel, err := invCamera.MultiplyTuple(Tuple{X: worldX, Y: worldY, Z: -1, W: 1})
        if err != nil {
                return Ray{}, err
        }

        origin, err := invCamera.MultiplyTuple(Tuple{X: 0, Y: 0, Z: 0, W: 1})
        if err != nil {
                return Ray{}, err
        }

        sub, err := pixel.Subtract(origin); if err != nil {
                return Ray{}, err
        }

        direction, err := sub.Normalize(); if err != nil {
                return Ray{}, err
        }

        return Ray{Origin: origin, Direction: direction}, nil
}

func (c Camera) Render(world World) (Canvas, error) {
        canvas := CreateCanvas(int(c.HSize), int(c.VSize))
        for y := 0; y < int(c.VSize) - 1; y++ {
                for x := 0; x < int(c.HSize) - 1; x++ {
                        ray, err := c.RayForPixel(x, y); if err != nil {
                                return canvas, err
                        }

                        color, err := world.ColorAt(ray); if err != nil {
                                return canvas, err
                        }

                        (&canvas).Write(x, y, color)
                }
        }
        return canvas, nil
}
