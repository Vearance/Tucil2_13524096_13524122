package main

// in .obj files Vertex is written as 'v' or 'vertex'
type Vertex struct {
    X, Y, Z float64
}

// in .obj files Triangle is written as 'f' or 'face'
type Triangle struct {
    V1, V2, V3 Vertex
}