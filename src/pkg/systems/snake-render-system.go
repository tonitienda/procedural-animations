package systems

import (
	"fmt"
	"image/color"
	"math"

	"github.com/aquilax/go-perlin"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tonitienda/procedural-animations-go/src/pkg/components"
	"github.com/tonitienda/procedural-animations-go/src/pkg/entities"
	"github.com/tonitienda/procedural-animations-go/src/pkg/world"
)

type SnakeRenderSystem struct {
	positions  map[entities.Entity]*components.Position
	velocities map[entities.Entity]*components.Velocity
	circles    map[entities.Entity]*components.Circle
	snakes     map[entities.Entity]*components.Snake
	perlin     *perlin.Perlin
}

func NewSnakeRenderSystem(world *world.World) *SnakeRenderSystem {
	return &SnakeRenderSystem{
		perlin:     perlin.NewPerlin(2, 2, 3, 100),
		positions:  world.Positions,
		velocities: world.Velocities,
		circles:    world.Circles,
		snakes:     world.Snakes,
	}
}
func (s *SnakeRenderSystem) Draw(screen *ebiten.Image) {
	//buffer := ebiten.NewImage(screen.Size()) // Create an intermediate buffer

	rightSidePoints := [][]float32{}
	leftSidePoints := [][]float32{}

	// Calculate the points for the skin
	for _, snake := range s.snakes {
		// Rename Circles to Parts or segments
		for _, part := range snake.Circles {
			pos := s.positions[part]
			vel := s.velocities[part]
			c := s.circles[part]
			rightPoint, leftPoint := calculatePerpendicularPoints(pos, vel, float64(c.Radius))

			rightSidePoints = append(rightSidePoints, rightPoint[:])
			leftSidePoints = append(leftSidePoints, leftPoint[:])
		}

		// Smooth the points using Catmull-Rom splines
		smoothRightSide := interpolateCatmullRom(rightSidePoints)
		smoothLeftSide := interpolateCatmullRom(leftSidePoints)

		// Generate the polygon from the smoothed points
		//vertices, indices := generatePolygon(smoothRightSide, smoothLeftSide)

		// Draw the polygon
		//buffer.DrawTriangles(vertices, indices, nil, nil)
		//screen.DrawTriangles(vertices, indices, nil, nil)
		for _, point := range smoothRightSide {
			vector.DrawFilledCircle(screen, point[0], point[1], 2, color.RGBA{32, 128, 0, 255}, true)
		}
		for _, point := range smoothLeftSide {
			vector.DrawFilledCircle(screen, point[0], point[1], 2, color.RGBA{32, 128, 0, 255}, true)
		}

	}

	//screen.DrawImage(buffer, nil)
}

// calculatePerpendicularPoints calculates two points perpendicular to the motion
func calculatePerpendicularPoints(pos *components.Position, vel *components.Velocity, radius float64) (point1, point2 [2]float32) {
	length := float32(math.Sqrt(float64(vel.X*vel.X + vel.Y*vel.Y)))

	if length == 0 {
		length = 1 // Avoid division by zero
	}

	normalizedVX := float32(vel.X) / length
	normalizedVY := float32(vel.Y) / length

	perpX := -normalizedVY
	perpY := normalizedVX

	point1[0] = float32(pos.X) + perpX*float32(radius)
	point1[1] = float32(pos.Y) + perpY*float32(radius)
	point2[0] = float32(pos.X) - perpX*float32(radius)
	point2[1] = float32(pos.Y) - perpY*float32(radius)

	return point1, point2
}

// Catmull-Rom interpolation function
func interpolateCatmullRom(points [][]float32) [][]float32 {
	if len(points) < 4 {
		return points // Catmull-Rom splines require at least 4 points
	}

	var smoothedPoints [][]float32
	for i := 0; i < len(points)-3; i++ {
		for t := 0; t <= 10; t++ {
			tt := float32(t) / 10.0
			smoothedPoints = append(smoothedPoints, catmullRom(points[i], points[i+1], points[i+2], points[i+3], tt))
		}
	}

	return smoothedPoints
}

// Catmull-Rom interpolation function
func catmullRom(p0, p1, p2, p3 []float32, t float32) []float32 {
	t2 := t * t
	t3 := t2 * t

	x := 0.5 * (2*p1[0] +
		(-p0[0]+p2[0])*t +
		(2*p0[0]-5*p1[0]+4*p2[0]-p3[0])*t2 +
		(-p0[0]+3*p1[0]-3*p2[0]+p3[0])*t3)

	y := 0.5 * (2*p1[1] +
		(-p0[1]+p2[1])*t +
		(2*p0[1]-5*p1[1]+4*p2[1]-p3[1])*t2 +
		(-p0[1]+3*p1[1]-3*p2[1]+p3[1])*t3)

	return []float32{x, y}
}

// generatePolygon creates the vertices and indices for drawing the snake
func generatePolygon(rightSide, leftSide [][]float32) ([]ebiten.Vertex, []uint16) {
	var vertices []ebiten.Vertex
	var indices []uint16

	for i := 0; i < len(rightSide)-1; i++ {
		// Right side vertices
		vertices = append(vertices, ebiten.Vertex{
			DstX:   float32(rightSide[i][0]),
			DstY:   float32(rightSide[i][1]),
			ColorR: 0.0,
			ColorG: 1.0,
			ColorB: 0.0,
			ColorA: 1.0,
		})
		vertices = append(vertices, ebiten.Vertex{
			DstX:   float32(rightSide[i+1][0]),
			DstY:   float32(rightSide[i+1][1]),
			ColorR: 0.0,
			ColorG: 1.0,
			ColorB: 0.0,
			ColorA: 1.0,
		})

		// Left side vertices
		vertices = append(vertices, ebiten.Vertex{
			DstX:   float32(leftSide[i][0]),
			DstY:   float32(leftSide[i][1]),
			ColorR: 0.0,
			ColorG: 1.0,
			ColorB: 0.0,
			ColorA: 1.0,
		})
		vertices = append(vertices, ebiten.Vertex{
			DstX:   float32(leftSide[i+1][0]),
			DstY:   float32(leftSide[i+1][1]),
			ColorR: 0.0,
			ColorG: 1.0,
			ColorB: 0.0,
			ColorA: 1.0,
		})

		// Create two triangles for each segment
		indices = append(indices, uint16(4*i), uint16(4*i+2), uint16(4*i+1))
		indices = append(indices, uint16(4*i+1), uint16(4*i+2), uint16(4*i+3))
	}
	fmt.Println("vertices", vertices, (len(vertices)))
	fmt.Println("indices", indices)
	return vertices, indices
}
