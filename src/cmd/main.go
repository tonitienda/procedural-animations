package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Size struct {
	Width  int
	Height int
}

type Game struct {
	Animal Animal
	Size   Size
}

type Part struct {
	Position Position
	Size     float32
}

type Animal struct {
	Parts []Part
}

func (a *Animal) Head() Part {
	return a.Parts[0]
}

func (a *Animal) MoveBy(x, y float32) {
	a.Parts[0].Position.x += x
	a.Parts[0].Position.y += y
}

func (a *Animal) MoveTo(x, y float32) {
	a.Parts[0].Position.x = x
	a.Parts[0].Position.y = y
}

func (a *Animal) FollowMouse(mx, my float32, width, height int) {
	// Move the head
	// FIXME - Current functionality calculates movement separately from x and y axis
	// Improve it to calculate the movement vector and normalize it
	maxSpeed := float64(3.0)

	fmt.Println("Mouse position: ", mx, my)
	fmt.Println("Animal position: ", a.Head().Position.x, a.Head().Position.y)

	diffX, diffY := float32(mx)-a.Head().Position.x, float32(my)-a.Head().Position.y

	fmt.Println("Difference: ", diffX, diffY)

	stepX := float32(0.0)
	stepY := float32(0.0)

	if diffX > 0 {
		stepX = float32(math.Min(maxSpeed, float64(diffX)))
	} else {
		stepX = float32(math.Max(-maxSpeed, float64(diffX)))
	}

	if diffY > 0 {
		stepY = float32(math.Min(maxSpeed, float64(diffY)))
	} else {
		stepY = float32(math.Max(-maxSpeed, float64(diffY)))
	}

	fmt.Println("Step: ", stepY, stepX)
	a.MoveBy(stepX, stepY)

	// Control that the Animal does not move out of the screen
	a.MoveTo(float32(math.Max(0, math.Min(float64(width), float64(a.Head().Position.x)))), float32(math.Max(0, math.Min(float64(height), float64(a.Head().Position.y)))))

	fmt.Println("-------------------")

}

func (g *Game) Update() error {

	mx, my := ebiten.CursorPosition()
	g.Animal.FollowMouse(float32(mx), float32(my), g.Size.Width, g.Size.Height)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	circleColor := color.RGBA{255, 255, 255, 255} // Red color with full opacity

	for _, part := range g.Animal.Parts {
		vector.StrokeCircle(screen, float32(part.Position.x), float32(part.Position.y), float32(part.Size), float32(1), circleColor, true)
		vector.DrawFilledCircle(screen, float32(part.Position.x), float32(part.Position.y), float32(5), circleColor, true)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Size.Width, g.Size.Height
}

type Position struct {
	x float32
	y float32
}

func main() {

	width, height := 800, 400

	animal := Animal{
		Parts: []Part{
			{
				Position: Position{
					x: float32(width / 2),
					y: float32(height / 2),
				},
				Size: 50,
			},
		}}

	game := &Game{
		Size:   Size{Width: width, Height: height},
		Animal: animal,
	}
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Procedural animations")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
