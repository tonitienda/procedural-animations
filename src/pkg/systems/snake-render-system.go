package systems

import (
	"fmt"
	"image/color"
	"math"

	"github.com/aquilax/go-perlin"
	"github.com/hajimehoshi/ebiten/v2"
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
	fmt.Printf("Screen: %#v\n", screen)

	// Get the dimensions of the screen
	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	fmt.Printf("Screen dimensions: %d x %d\n", w, h)

	img := ebiten.NewImage(w, h)

	// Clear the image with a background color
	img.Fill(color.RGBA{64, 200, 0, 255}) // Green background

	// vs := []ebiten.Vertex{
	// 	{DstX: 100, DstY: 100, ColorR: 1, ColorG: 0, ColorB: 0, ColorA: 1},
	// 	{DstX: 200, DstY: 100, ColorR: 0, ColorG: 1, ColorB: 0, ColorA: 1},
	// 	{DstX: 150, DstY: 200, ColorR: 0, ColorG: 0, ColorB: 1, ColorA: 1},
	// }
	// is := []uint16{0, 1, 2}
	// path := vector.Path{}
	// path.MoveTo(100, 100)
	// path.LineTo(200, 100)
	// path.LineTo(150, 200)
	// path.Close()

	//vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)

	op := &ebiten.DrawTrianglesOptions{}
	op.AntiAlias = true

	op.FillRule = ebiten.NonZero

	// Calculate the points for the skin
	for _, snake := range s.snakes {

		rightSidePoints := [][2]float32{}
		leftSidePoints := [][2]float32{}
		// Rename Circles to Parts or segments
		// Head and tail are special. We need to close the body
		// head := snake.Circles[0]

		// headOrientationX, headOrientationY := calculateOrientation(s.velocities[head])

		// // See how to calculate the extra points for the head
		// rightSidePoints = append(rightSidePoints, [2]float32{float32(s.positions[head].X) + headOrientationX*float32(s.circles[head].Radius/2), float32(s.positions[head].Y) + headOrientationY*float32(s.circles[head].Radius*2)})
		// leftSidePoints = append(leftSidePoints, [2]float32{float32(s.positions[head].X) - headOrientationX*float32(s.circles[head].Radius/2), float32(s.positions[head].Y) + headOrientationY*float32(s.circles[head].Radius*2)})

		for _, part := range snake.Circles {
			pos := s.positions[part]
			vel := s.velocities[part]
			c := s.circles[part]
			rightPoint, leftPoint := calculatePerpendicularPoints(pos, vel, float64(c.Radius))

			rightSidePoints = append(rightSidePoints, rightPoint)
			leftSidePoints = append(leftSidePoints, leftPoint)
		}

		// See how to calculate the extra points for the tail
		//tail := snake.Circles[len(snake.Circles)-1]
		//tailOrientationX, tailOrientationY := calculateOrientation(s.velocities[head])
		//rightSidePoints = append(rightSidePoints, [2]float32{float32(s.positions[tail].X) + tailOrientationX*float32(s.circles[tail].Radius), float32(s.positions[tail].Y) + tailOrientationY*float32(s.circles[tail].Radius)})
		//leftSidePoints = append(leftSidePoints, [2]float32{float32(s.positions[tail].X) + tailOrientationX*float32(s.circles[tail].Radius), float32(s.positions[tail].Y) + tailOrientationY*float32(s.circles[tail].Radius)})

		//smoothRightSide := interpolateCatmullRom(rightSidePoints)
		//smoothLeftSide := interpolateCatmullRom(leftSidePoints)

		//vs, is := s.generatePolygon(smoothLeftSide, smoothRightSide)

		vs, is := s.generatePolygon(leftSidePoints, rightSidePoints)

		screen.DrawTriangles(vs, is, img, op)
	}

	// 	// Smooth the points using Catmull-Rom splines
	// 	//smoothRightSide := interpolateCatmullRom(rightSidePoints)
	// 	//smoothLeftSide := interpolateCatmullRom(leftSidePoints)

	// 	// Generate the polygon from the smoothed points
	// 	//vertices, indices := generatePolygon(smoothRightSide, smoothLeftSide)
	// 	// vertices, indices := generatePolygon(rightSidePoints, leftSidePoints)

	// 	// if len(vertices) == 0 || len(indices) == 0 {
	// 	// 	fmt.Println("Empty vertices or indices")
	// 	// 	return
	// 	// }
	// 	// if len(indices)%3 != 0 {
	// 	// 	fmt.Println("Invalid number of indices")
	// 	// 	return
	// 	// }

	// 	// // Draw the polygon
	// 	// screen.DrawTriangles(vertices, indices, nil, nil)

	// 	vertices := make([]ebiten.Vertex, 0, 3)
	// 	indices := []uint16{0, 1, 2}

	// 	// Create vertices for the first triangle
	// 	vertices = append(vertices, ebiten.Vertex{
	// 		DstX:   rightSidePoints[0][0],
	// 		DstY:   rightSidePoints[0][1],
	// 		ColorR: 0,
	// 		ColorG: 1,
	// 		ColorB: 0,
	// 		ColorA: 1,
	// 	})
	// 	vertices = append(vertices, ebiten.Vertex{
	// 		DstX:   rightSidePoints[1][0],
	// 		DstY:   rightSidePoints[1][1],
	// 		ColorR: 0,
	// 		ColorG: 1,
	// 		ColorB: 0,
	// 		ColorA: 1,
	// 	})
	// 	vertices = append(vertices, ebiten.Vertex{
	// 		DstX:   leftSidePoints[0][0],
	// 		DstY:   leftSidePoints[0][1],
	// 		ColorR: 0,
	// 		ColorG: 1,
	// 		ColorB: 0,
	// 		ColorA: 1,
	// 	})

	// 	// Print debug information
	// 	fmt.Printf("Vertices: %+v\n", vertices)
	// 	fmt.Printf("Indices: %v\n", indices)

	// 	// Check for NaN or Inf values
	// 	for _, v := range vertices {
	// 		if math.IsNaN(float64(v.DstX)) || math.IsInf(float64(v.DstX), 0) ||
	// 			math.IsNaN(float64(v.DstY)) || math.IsInf(float64(v.DstY), 0) {
	// 			fmt.Println("Warning: Invalid coordinate detected")
	// 			return
	// 		}
	// 	}

	// 	if screen == nil {
	// 		fmt.Println("Error: screen is nil")
	// 		return
	// 	}
	// 	//	screen.DrawTriangles(vertices, indices, nil, &ebiten.DrawTrianglesOptions{})

	//	}

}

// calculatePerpendicularPoints calculates two points perpendicular to the motion
func calculateOrientation(vel *components.Velocity) (float32, float32) {
	length := float32(math.Sqrt(float64(vel.X*vel.X + vel.Y*vel.Y)))

	if length == 0 {
		length = 1 // Avoid division by zero
	}

	normalizedVX := float32(vel.X) / length
	normalizedVY := float32(vel.Y) / length

	return normalizedVX, normalizedVY
}

// calculatePerpendicularPoints calculates two points perpendicular to the motion
func calculatePerpendicularPoints(pos *components.Position, vel *components.Velocity, radius float64) (point1, point2 [2]float32) {
	normVX, normVY := calculateOrientation(vel)

	perpX := -normVY
	perpY := normVX

	point1[0] = float32(pos.X) + perpX*float32(radius)
	point1[1] = float32(pos.Y) + perpY*float32(radius)
	point2[0] = float32(pos.X) - perpX*float32(radius)
	point2[1] = float32(pos.Y) - perpY*float32(radius)

	return point1, point2
}

// Catmull-Rom interpolation function
func interpolateCatmullRom(points [][2]float32) [][2]float32 {
	if len(points) < 4 {
		return points // Catmull-Rom splines require at least 4 points
	}

	var smoothedPoints [][2]float32
	for i := 0; i < len(points)-3; i++ {
		for t := 0; t <= 10; t++ {
			tt := float32(t) / 10.0
			smoothedPoints = append(smoothedPoints, catmullRom(points[i], points[i+1], points[i+2], points[i+3], tt))
		}
	}

	return smoothedPoints
}

// Catmull-Rom interpolation function
func catmullRom(p0, p1, p2, p3 [2]float32, t float32) [2]float32 {
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

	return [2]float32{x, y}
}

// generatePolygon creates the vertices and indices for drawing the snake
// FIXME - Noise location should be calculated from the HEAD perspective.
// So the body of the snake is stable and does not fluctuate based on the movement thought the board
func (s *SnakeRenderSystem) generatePolygon(rightSide, leftSide [][2]float32) ([]ebiten.Vertex, []uint16) {
	vertices := make([]ebiten.Vertex, 0, 4*(len(rightSide)-1))
	indices := make([]uint16, 0, 6*(len(rightSide)-1))

	colorBase := float32(0.3)
	noiseFactor := float32(0.7)

	for i := 0; i < len(rightSide)-1; i++ {

		riNoise := float32(s.perlin.Noise2D(float64(i*10), float64(i*10)))
		rip1Noise := float32(s.perlin.Noise2D(float64(i*10+1), float64(i*10+1)))
		liNoise := float32(s.perlin.Noise2D(float64(i*10), float64(i*10)))
		lip1Noise := float32(s.perlin.Noise2D(float64(i*10+1), float64(i*10+1)))

		riFactor := colorBase + noiseFactor*riNoise
		rip1Factor := colorBase + noiseFactor*rip1Noise
		liFactor := colorBase + noiseFactor*liNoise
		lip1Factor := colorBase + noiseFactor*lip1Noise

		// Right side vertices
		vertices = append(vertices, ebiten.Vertex{
			DstX:   float32(rightSide[i][0]),
			DstY:   float32(rightSide[i][1]),
			ColorR: riFactor,
			ColorG: riFactor,
			ColorB: riFactor,
			ColorA: 1.0,
		})
		vertices = append(vertices, ebiten.Vertex{
			DstX:   float32(rightSide[i+1][0]),
			DstY:   float32(rightSide[i+1][1]),
			ColorR: rip1Factor,
			ColorG: rip1Factor,
			ColorB: rip1Factor,
			ColorA: 1.0,
		})

		// Left side vertices
		vertices = append(vertices, ebiten.Vertex{
			DstX:   float32(leftSide[i][0]),
			DstY:   float32(leftSide[i][1]),
			ColorR: liFactor,
			ColorG: liFactor,
			ColorB: liFactor,
			ColorA: 1.0,
		})
		vertices = append(vertices, ebiten.Vertex{
			DstX:   float32(leftSide[i+1][0]),
			DstY:   float32(leftSide[i+1][1]),
			ColorR: lip1Factor,
			ColorG: lip1Factor,
			ColorB: lip1Factor,
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
