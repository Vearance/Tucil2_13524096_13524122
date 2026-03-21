package main

func getBoundingBox(vertex []Vertex) BoundingBox {
	var min Vertex = vertex[0]
	var max Vertex = vertex[0]

	for _, coor := range vertex {
		if coor.X < min.X {
			min.X = coor.X
		}
		if coor.Y < min.Y {
			min.Y = coor.Y
		}
		if coor.Z < min.Z {
			min.Z = coor.Z
		}

		if coor.X > min.X {
			max.X = coor.X
		}
		if coor.Y > min.Y {
			max.Y = coor.Y
		}
		if coor.Z > min.Z {
			max.Z = coor.Z
		}
	}

	return BoundingBox{min, max}
}

func getBoundingCube(bounding BoundingBox) Cube {
	x := bounding.Max.X - bounding.Min.X
	y := bounding.Max.Y - bounding.Min.Y
	z := bounding.Max.Z - bounding.Min.Z

	max := x
	if y > max {
		max = y
	}
	if z > max {
		max = z
	}

	return Cube{bounding.Min, max}
}