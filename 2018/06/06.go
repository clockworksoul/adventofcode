package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Point struct {
	X, Y float64
}

// AngleTo returns the angle in degrees from p to another point, where 0
// is returned if Point{0,0}.AngleTo(Point{2, 0}).
func (p *Point) AngleTo(p2 *Point) float64 {
	if p == p2 {
		return 360.0
	}

	deltax := float64(p2.X - p.X)
	deltay := float64(p2.Y - p.Y)
	radians := math.Atan(deltay / deltax)
	degrees := radians * (180.0 / math.Pi)

	if deltax < 0 {
		degrees += 180.0
	} else if degrees < 0 {
		degrees += 360.0
	}

	return degrees
}

// DistanceTo returns the Manhattan distance from this to another Point.
func (p *Point) DistanceTo(p2 *Point) float64 {
	return math.Abs(p.X-p2.X) + math.Abs(p.Y-p2.Y)
}

func (p *Point) String() string {
	return fmt.Sprintf("(%.0f,%.0f)", p.X, p.Y)
}

// FindMargins finds all of the points on the outside of the the cluster. That
// is, those which will have an infinite number of squares closest to them.
// This is SO brute force.
func FindMargins(points []*Point) []*Point {
	minX, minY := 999999.0, 999999.0
	maxX, maxY := -999999.0, -999999.0
	marginMap := make(map[Point]*Point)

	for _, p := range points {
		if p.X > maxX {
			maxX = p.X
		}
		if p.X < minX {
			minX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
		if p.Y < minY {
			minY = p.Y
		}
	}

	minX -= maxX
	maxX *= 2
	minY -= maxY
	maxY *= 2

	var p *Point

	for x := minX; x < maxX; x++ {
		p = FindClosestTo(points, &Point{x, minY})
		marginMap[*p] = p
	}
	for x := minX; x < maxX; x++ {
		p = FindClosestTo(points, &Point{x, maxY})
		marginMap[*p] = p
	}
	for y := minY; y < maxY; y++ {
		p = FindClosestTo(points, &Point{minX, y})
		marginMap[*p] = p
	}
	for y := minY; y < maxY; y++ {
		p = FindClosestTo(points, &Point{maxX, y})
		marginMap[*p] = p
	}

	margins := make([]*Point, 0)
	for _, v := range marginMap {
		margins = append(margins, v)
	}

	return margins
}

func FindClosestTo(points []*Point, point *Point) *Point {
	mind := 999999.9
	var minp *Point

	for _, p := range points {
		d := point.DistanceTo(p)
		if d < mind {
			mind = d
			minp = p
		}
	}

	return minp
}

// IngestPoints reads all the points from a text file and returns a slice.
func IngestPoints(filename string) ([]*Point, error) {
	points := make([]*Point, 0)

	file, err := os.Open(filename)
	if err != nil {
		return points, err
	}

	scanner := bufio.NewScanner(file)
	counter := 1

	for ; scanner.Scan(); counter++ {
		p := Point{}
		line := scanner.Text()
		_, err := fmt.Sscanf(line, "%f, %f", &p.X, &p.Y)
		if err != nil {
			return points, err
		}

		points = append(points, &p)
	}

	return points, nil
}

func SizeOfClosestAreas(points []*Point) map[Point]int {
	areas := make(map[Point]int)

	minX, minY := 999999.0, 999999.0
	maxX, maxY := -999999.0, -999999.0

	for _, p := range points {
		areas[*p] = 0

		if p.X > maxX {
			maxX = p.X
		}
		if p.X < minX {
			minX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
		if p.Y < minY {
			minY = p.Y
		}
	}

	// Scan all points to find the closest ones.
	for x := 0.0; x < maxX; x += 1.0 {
		for y := 0.0; y < maxY; y += 1.0 {
			qp := &Point{X: x, Y: y}

			sort.Slice(points, func(i, j int) bool {
				disti := qp.DistanceTo(points[i])
				distj := qp.DistanceTo(points[j])

				return disti < distj
			})

			if qp.DistanceTo(points[0]) != qp.DistanceTo(points[1]) {
				areas[*points[0]]++
			}
		}
	}

	// Remove margin points, since they'll have infinite areas
	for _, p := range FindMargins(points) {
		areas[*p] = -1
	}

	return areas
}

func MaxClosestAreas(points []*Point) (Point, int) {
	maxp := Point{X: -1, Y: -1}
	maxa := -1

	for p, a := range SizeOfClosestAreas(points) {
		if a > maxa {
			maxp = p
			maxa = a
		}
	}

	return maxp, maxa
}

func FindSafe(points []*Point) int {
	minX, minY := 999999.0, 999999.0
	maxX, maxY := -999999.0, -999999.0

	for _, p := range points {
		if p.X > maxX {
			maxX = p.X
		}
		if p.X < minX {
			minX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
		if p.Y < minY {
			minY = p.Y
		}
	}

	minX -= 1000
	maxX += 1000
	minY -= 1000
	maxY += 1000

	count := 0

	for y := minY; y < maxY; y++ {

	xLoop:
		for x := minX; x < maxX; x++ {
			qp := &Point{x, y}
			sum := 0

			for _, p := range points {
				sum += int(qp.DistanceTo(p))

				if sum >= 10000 {
					continue xLoop
				}
			}

			// fmt.Printf("%s --> %d\n", qp, sum)
			count++
		}
	}

	return count
}

func main() {
	points, err := IngestPoints("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// GeneratePlot(points)

	// maxp, maxa := MaxClosestAreas(points)
	// fmt.Printf("Biggest %s with area of %d\n", &maxp, maxa)

	fmt.Println(FindSafe(points))
}
