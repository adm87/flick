package shapes

type Polygon struct {
	xy       [2]float32
	bounds   [4]float32
	vertices [][2]float32
}

func NewPolygon() *Polygon {
	return &Polygon{
		vertices: make([][2]float32, 0),
	}
}

// Vertices returns the vertices of the polygon in local space.
func (p *Polygon) Vertices() [][2]float32 {
	return p.vertices
}

// SetVertices sets the vertices of the polygon in local space.
func (p *Polygon) SetVertices(vertices [][2]float32) {
	if !isCCW(vertices) {
		vertices = reverseVertices(vertices)
	}

	p.bounds = calculateBounds(vertices)
	p.vertices = vertices
}

func (p *Polygon) Position() (float32, float32) {
	return p.xy[0], p.xy[1]
}

func (p *Polygon) SetPosition(x, y float32) Shape {
	p.xy[0] = x
	p.xy[1] = y
	return p
}

func (p *Polygon) Bounds(x, y float32) [4]float32 {
	return [4]float32{
		x + p.xy[0] + p.bounds[0],
		y + p.xy[1] + p.bounds[1],
		x + p.xy[0] + p.bounds[2],
		y + p.xy[1] + p.bounds[3],
	}
}

func (p *Polygon) Type() ShapeType {
	return ShapeTypePolygon
}

// =============== Polygon Helpers ==================

func GroupVertices(points []float32) [][2]float32 {
	if len(points)%2 != 0 {
		return nil
	}
	n := len(points) / 2
	vertices := make([][2]float32, n)
	for i := range n {
		vertices[i] = [2]float32{points[2*i], points[2*i+1]}
	}
	return vertices
}

func calculateBounds(vertices [][2]float32) [4]float32 {
	if len(vertices) == 0 {
		return [4]float32{0, 0, 0, 0}
	}

	left := vertices[0][0]
	right := vertices[0][0]
	top := vertices[0][1]
	bottom := vertices[0][1]

	for _, v := range vertices[1:] {
		if v[0] < left {
			left = v[0]
		}
		if v[0] > right {
			right = v[0]
		}
		if v[1] < top {
			top = v[1]
		}
		if v[1] > bottom {
			bottom = v[1]
		}
	}

	return [4]float32{left, top, right, bottom}
}

func isCCW(vertices [][2]float32) bool {
	if len(vertices) < 3 {
		return false
	}

	sum := float32(0)
	for i := range vertices {
		j := (i + 1) % len(vertices)
		sum += (vertices[j][0] - vertices[i][0]) * (vertices[j][1] + vertices[i][1])
	}

	return sum > 0
}

func reverseVertices(vertices [][2]float32) [][2]float32 {
	n := len(vertices)
	reversed := make([][2]float32, n)
	for i := range n {
		reversed[i] = vertices[n-1-i]
	}
	return reversed
}
