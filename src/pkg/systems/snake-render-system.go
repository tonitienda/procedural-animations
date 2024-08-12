package systems

import (
	"image/color"
	"math"

	"github.com/aquilax/go-perlin"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonitienda/procedural-animations-go/src/pkg/components"
	"github.com/tonitienda/procedural-animations-go/src/pkg/entities"
	"github.com/tonitienda/procedural-animations-go/src/pkg/world"
)

type SnakeRenderSystem struct {
	positions    map[entities.Entity]*components.Position
	orientations map[entities.Entity]*components.Orientation
	circles      map[entities.Entity]*components.Circle
	snakes       map[entities.Entity]*components.Snake
	perlin       *perlin.Perlin
}

func NewSnakeRenderSystem(world *world.World) *SnakeRenderSystem {
	return &SnakeRenderSystem{
		perlin:       perlin.NewPerlin(2, 2, 3, 100),
		positions:    world.Positions,
		orientations: world.Orientations,
		circles:      world.Circles,
		snakes:       world.Snakes,
	}
}

func (s *SnakeRenderSystem) Draw(screen *ebiten.Image) {
	//fmt.Printf("Screen: %#v\n", screen)

	// Get the dimensions of the screen
	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	//fmt.Printf("Screen dimensions: %d x %d\n", w, h)

	img := ebiten.NewImage(w, h)

	// Clear the image with a background color
	img.Fill(color.RGBA{64, 200, 0, 255}) // Green background

	// Calculate the points for the skin
	for _, snake := range s.snakes {

		snakeContour := make([][2]float32, len(snake.Circles)*2+2)

		//rightSidePoints := [][2]float32{}
		//leftSidePoints := [][2]float32{}
		// Rename Circles to Parts or segments
		// Head and tail are special. We need to close the body

		head := snake.Circles[0]
		headOrientation := s.orientations[head].Radians

		// End with a point in the head that is almost parallel to the orientation
		// but a little bit to the right side
		rightHeadPoint := [2]float32{
			float32(s.positions[head].X) + float32(math.Cos(float64(headOrientation)+math.Pi/8)*float64(s.circles[head].Radius)),
			float32(s.positions[head].Y) + float32(math.Sin(float64(headOrientation)+math.Pi/8)*float64(s.circles[head].Radius)),
		}

		// Start with a point in the head that is almost parallel to the orientation
		// but a little bit to the left side
		leftHeadPoint := [2]float32{
			float32(s.positions[head].X) + float32(math.Cos(float64(headOrientation)-math.Pi/8)*float64(s.circles[head].Radius)),
			float32(s.positions[head].Y) + float32(math.Sin(float64(headOrientation)-math.Pi/8)*float64(s.circles[head].Radius)),
		}

		snakeContourLastIndex := len(snakeContour) - 1
		snakeContour[0] = leftHeadPoint
		snakeContour[snakeContourLastIndex] = rightHeadPoint

		for idx, part := range snake.Circles {
			pos := s.positions[part]
			orientation := s.orientations[part]
			c := s.circles[part]
			rightPoint, leftPoint := calculatePerpendicularPoints(pos, orientation, float64(c.Radius))

			snakeContour[idx+1] = leftPoint
			snakeContour[snakeContourLastIndex-idx-1] = rightPoint
		}

		//snakeContour = interpolateCatmullRom(snakeContour)

		// Create vertices from the snake contour
		vertices := make([]ebiten.Vertex, len(snakeContour))
		for i, p := range snakeContour {
			vertices[i] = ebiten.Vertex{
				DstX:   p[0],
				DstY:   p[1],
				ColorR: 0, ColorG: 1, ColorB: 0, ColorA: 1, // Green color
			}
		}

		// Create triangle indices
		indices := []uint16{}
		for i := 0; i < len(snakeContour)/2-1; i++ {
			// Connecting the current pair of points with the next pair
			indices = append(indices,
				uint16(i),                     // Top point of current segment
				uint16(len(snakeContour)-1-i), // Bottom point of current segment
				uint16(i+1),                   // Top point of next segment

				uint16(len(snakeContour)-1-i), // Bottom point of current segment
				uint16(len(snakeContour)-2-i), // Bottom point of next segment
				uint16(i+1),                   // Top point of next segment
			)
		}

		screen.DrawTriangles(vertices, indices, img, nil)

		// for _, point := range snakeContour {
		// 	vector.DrawFilledCircle(screen, float32(point[0]), float32(point[1]), 2, color.RGBA{255, 0, 0, 255}, true)
		// }

	}

}

func calculatePerpendicularPoints(pos *components.Position, orientation *components.Orientation, radius float64) (point1, point2 [2]float32) {
	normVX := float32(math.Cos(orientation.Radians))
	normVY := float32(math.Sin(orientation.Radians))

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
