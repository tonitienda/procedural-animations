package main

import (
	"image/color"
	"time"

	"github.com/aquilax/go-perlin"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	noiseImage *ebiten.Image
	offsetX    float64
	offsetY    float64
	perlin     *perlin.Perlin
}

func NewGame() *Game {
	noiseImage := ebiten.NewImage(800, 600)
	p := perlin.NewPerlin(2, 3, 4, time.Now().UnixNano()) // persistence, frequency, octaves, seed
	return &Game{
		noiseImage: noiseImage,
		offsetX:    0,
		offsetY:    0,
		perlin:     p,
	}
}

func (g *Game) Update() error {
	g.offsetX += 0.001 // Slower and smoother movement
	g.offsetY += 0.001
	g.generateNoiseImage()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{50, 150, 255, 255}) // Full blue with full opacity

	opts := &ebiten.DrawImageOptions{}
	opts.ColorScale.Scale(0.4, 0.8, 1.0, 0.5) // Blue tint, 50% transparency
	screen.DrawImage(g.noiseImage, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func (g *Game) generateNoiseImage() {
	width, height := g.noiseImage.Size()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			nx := g.offsetX + float64(x)/float64(width)
			ny := g.offsetY + float64(y)/float64(height)
			noiseValue := g.perlin.Noise2D(nx, ny)
			c := uint8((noiseValue + 1) * 128) // Map from [-1, 1] to [0, 255]
			g.noiseImage.Set(x, y, color.RGBA{c, c, c, 255})
		}
	}
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Water Effect Example")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
