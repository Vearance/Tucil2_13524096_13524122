package main

func getCubeVertex(cube Cube) []Vertex {
	min := cube.Min
	max := Vertex{min.X + cube.Size, min.Y + cube.Size, min.Z + cube.Size}

	return []Vertex{
		{min.X, min.Y, min.Z},
		{max.X, min.Y, min.Z},
		{max.X, max.Y, min.Z},
		{min.X, max.Y, min.Z},
		{min.X, min.Y, max.Z},
		{max.X, min.Y, max.Z},
		{max.X, max.Y, max.Z},
		{min.X, max.Y, max.Z},
	}
}

func cubeMesh(vertex *[]Vertex, faces *[][3]int, cube Cube) {
	cubeVertex := getCubeVertex(cube)
	base := len(*vertex) + 1
	*vertex = append(*vertex, cubeVertex...)

	cubeFace := [][3]int{
		{0, 1, 2}, {0, 2, 3},
		{4, 6, 5}, {4, 7, 6},
		{0, 5, 1}, {0, 4, 5},
		{3, 2, 6}, {3, 6, 7},
		{0, 3, 7}, {0, 7, 4},
		{1, 5, 6}, {1, 6, 2},
	}

	for _, face := range cubeFace {
		*faces = append(*faces, [3]int{base + face[0], base + face[1], base + face[2]})
	}
}

func voxelToMesh(voxels []Voxel) ([]Vertex, [][3]int) {
	vertex := make([]Vertex, 0, len(voxels)*8)
	faces := make([][3]int, 0, len(voxels)*12)

	for _, voxel := range voxels {
		cubeMesh(&vertex, &faces, voxel.Cube)
	}

	return vertex, faces
}
