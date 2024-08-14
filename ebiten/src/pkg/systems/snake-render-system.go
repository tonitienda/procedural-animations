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
	positions        map[entities.Entity]*components.Position
	initialPositions map[entities.Entity]*components.InitialPosition
	orientations     map[entities.Entity]*components.Orientation
	circles          map[entities.Entity]*components.Circle
	snakes           map[entities.Entity]*components.Snake
	perlin           *perlin.Perlin
	textture         *ebiten.Image
}

func NewSnakeRenderSystem(world *world.World) *SnakeRenderSystem {

	return &SnakeRenderSystem{
		perlin:           perlin.NewPerlin(2, 2, 3, 100),
		positions:        world.Positions,
		initialPositions: world.InitialPositions,
		orientations:     world.Orientations,
		circles:          world.Circles,
		snakes:           world.Snakes,
	}
}

func generateTexture(width, height int, scale float64, perlin *perlin.Perlin) *ebiten.Image {
	img := ebiten.NewImage(width, height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Scale the coordinates
			nx := float64(x) / scale
			ny := float64(y) / scale

			// Generate Perlin noise value
			noise := perlin.Noise2D(nx, ny)

			// Map the noise value to a color (e.g., green to brown)
			green := uint8((noise+1)/2*64 + 64) // Scale to [0, 255]
			red := uint8((noise+1)/2*32 + 64)   // Scale to [0, 255]

			// Set the color based on Perlin noise
			img.Set(x, y, color.RGBA{R: red, G: green, B: 0, A: 255})
		}
	}

	return img
}

func (s *SnakeRenderSystem) Draw(screen *ebiten.Image, op *world.DrawOptions) {
	if s.textture == nil {
		s.textture = generateTexture(screen.Bounds().Dx(), screen.Bounds().Dy(), 100, s.perlin)
	}

	// Calculate the points for the skin
	for _, snake := range s.snakes {

		snakeContour := make([][2]float32, len(snake.Segments)*2+2)
		head := snake.Segments[0]
		headOrientation := s.orientations[head].Radians

		second := snake.Segments[1]
		secondOrientation := s.orientations[second].Radians

		rightEye := [2]float32{
			float32(s.positions[second].X) + float32(math.Cos(float64(secondOrientation)+math.Pi/3)*float64(s.circles[second].Radius)),
			float32(s.positions[second].Y) + float32(math.Sin(float64(secondOrientation)+math.Pi/3)*float64(s.circles[second].Radius)),
		}

		leftEye := [2]float32{
			float32(s.positions[second].X) + float32(math.Cos(float64(secondOrientation)-math.Pi/3)*float64(s.circles[second].Radius)),
			float32(s.positions[second].Y) + float32(math.Sin(float64(secondOrientation)-math.Pi/3)*float64(s.circles[second].Radius)),
		}

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

		for idx, segment := range snake.Segments {
			pos := s.positions[segment]
			orientation := s.orientations[segment]
			c := s.circles[segment]
			rightPoint, leftPoint := calculatePerpendicularPoints(pos, orientation, float64(c.Radius))

			snakeContour[idx+1] = leftPoint
			snakeContour[snakeContourLastIndex-idx-1] = rightPoint
		}

		smoothedContour := interpolateCatmullRom(snakeContour, 10)

		//initialPointCount := len(snakeContour)
		//contourPointsFactor := len(smoothedContour) / initialPointCount

		textureScale := 1
		segmentLength := len(smoothedContour) / (len(snakeContour) - 3) // -3 accounts for the first, last, and Catmull-Rom needing 4 points

		// Create vertices from the snake contour
		vertices := make([]ebiten.Vertex, len(smoothedContour))

		for i, p := range smoothedContour {
			var srcX, srcY float32

			if i%segmentLength == 0 {
				// This is a point corresponding to an original point
				originalIndex := i / segmentLength
				srcX = snakeContour[originalIndex][0] / float32(textureScale)
				srcY = snakeContour[originalIndex][1] / float32(textureScale)
			} else {
				// Optionally interpolate between known points

				// previousOriginalIndex := int(math.Floor(float64(i) / float64(segmentLength)))
				// nextOriginalIndex := previousOriginalIndex + 1
				// if nextOriginalIndex >= len(snakeContour) {
				// 	nextOriginalIndex = previousOriginalIndex
				// }
				// weight := float32(i%(segmentLength)) / float32(segmentLength)

				// srcX = snakeContour[previousOriginalIndex][0]*(1-weight)/float32(textureScale) +
				// 	snakeContour[nextOriginalIndex][0]*weight/float32(textureScale)
				// srcY = snakeContour[previousOriginalIndex][1]*(1-weight)/float32(textureScale) +
				// 	snakeContour[nextOriginalIndex][1]*weight/float32(textureScale)
				originalIndex := i / segmentLength
				srcX = snakeContour[originalIndex][0] / float32(textureScale)
				srcY = snakeContour[originalIndex][1] / float32(textureScale)

			}

			fmt.Println("X", p[0], "Y", p[1], "srcX", srcX, "srcY", srcY)

			vertices[i] = ebiten.Vertex{
				DstX:   p[0],
				DstY:   p[1],
				SrcX:   srcX,
				SrcY:   srcY,
				ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1, // Green color
			}
		}

		// Create triangle indices
		indices := []uint16{}
		for i := 0; i < len(smoothedContour)/2-1; i++ {
			// Connecting the current pair of points with the next pair
			indices = append(indices,
				uint16(i),                        // Top point of current segment
				uint16(len(smoothedContour)-1-i), // Bottom point of current segment
				uint16(i+1),                      // Top point of next segment

				uint16(len(smoothedContour)-1-i), // Bottom point of current segment
				uint16(len(smoothedContour)-2-i), // Bottom point of next segment
				uint16(i+1),                      // Top point of next segment
			)
		}

		screen.DrawTriangles(vertices, indices, s.textture, nil)

		vector.DrawFilledCircle(screen, float32(leftEye[0]), float32(leftEye[1]), 10, color.RGBA{100, 100, 220, 255}, true)
		vector.DrawFilledCircle(screen, float32(rightEye[0]), float32(rightEye[1]), 10, color.RGBA{100, 100, 220, 255}, true)

		if op.Debug {
			for _, snake := range s.snakes {
				for _, segment := range snake.Segments {
					pos := s.positions[segment]
					circle := s.circles[segment]
					orientation := s.orientations[segment]

					vector.StrokeLine(screen,
						float32(pos.X),
						float32(pos.Y),
						float32(pos.X)+float32(math.Cos(orientation.Radians)*float64(circle.Radius)),
						float32(pos.Y)+float32(math.Sin(orientation.Radians)*float64(circle.Radius)),
						1,
						color.RGBA{255, 0, 255, 255}, true)
					vector.DrawFilledCircle(screen, float32(pos.X), float32(pos.Y), 3, color.RGBA{255, 0, 255, 255}, true)
					vector.StrokeCircle(screen, float32(pos.X), float32(pos.Y), float32(circle.Radius), 1, color.RGBA{255, 0, 255, 255}, true)
				}
			}
		}

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

func interpolateCatmullRom(points [][2]float32, segmentsPerCurve int) [][2]float32 {
	if len(points) < 4 {
		return points // Catmull-Rom splines require at least 4 points
	}

	var smoothedPoints [][2]float32

	// Ensure the first point is included
	smoothedPoints = append(smoothedPoints, points[0])

	for i := 0; i < len(points)-3; i++ {
		for t := 1; t <= segmentsPerCurve; t++ {
			tt := float32(t) / float32(segmentsPerCurve)
			smoothedPoints = append(smoothedPoints, catmullRom(points[i], points[i+1], points[i+2], points[i+3], tt))
		}
	}

	// Ensure the last point is included
	smoothedPoints = append(smoothedPoints, points[len(points)-1])

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
