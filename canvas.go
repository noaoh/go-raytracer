package raytracer

import (
	"fmt"
	"os"
)

type Canvas struct {
	Width  int
	Height int
	Pixels [][]Color
}

func CreateCanvas(w, h int) Canvas {
	p := make([][]Color, h)
	for x := range p {
		p[x] = make([]Color, w)
	}

	return Canvas{Width: w, Height: h, Pixels: p}
}

func (c *Canvas) Write(w, h int, col Color) {
	c.Pixels[h][w] = col
}

func (c *Canvas) WriteFile(fh string) error {
	f, err := os.Create(fh)
	if err != nil {
		os.Remove(fh)
		return err
	}
	defer f.Close()

	_, err = f.WriteString("P3\n")
	if err != nil {
		return err
	}

	dim := fmt.Sprintf("%d %d\n", c.Width, c.Height)
	_, err = f.WriteString(dim)
	if err != nil {
		return err
	}

	_, err = f.WriteString("256\n")
	if err != nil {
		return err
	}

	for _, row := range c.Pixels {
		for _, col := range row {
			// There's probably a better place to do
			// this, I'm going to try and figure what
			// dot products I should apply this to
			// in the Lighting function
			r := int(Clamp(256*col.R, 0, 256))
			g := int(Clamp(256*col.G, 0, 256))
			b := int(Clamp(256*col.B, 0, 256))
			// This code returns the wrong value
			// for the specular reflection
			/*
			   r := int(256 * col.R)
			   g := int(256 * col.G)
			   b := int(256 * col.B)
			*/
			rgb := fmt.Sprintf("%d %d %d\n", r, g, b)
			_, err := f.WriteString(rgb)
			if err != nil {
				return err
			}
		}
	}
	f.WriteString("\n")
	return nil
}
