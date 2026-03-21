package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loadObject(path string) ([]Vertex, []Triangle, error) {
	file, err := os.Open(path)
    if err != nil {
        return nil, nil, err
    }
	defer file.Close()

	var vertex []Vertex
	var triangles []Triangle
	
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "v") {
			coordinate := strings.Fields(line)

			x,_ := strconv.ParseFloat(coordinate[1], 64)
			y,_ := strconv.ParseFloat(coordinate[2], 64)
			z,_ := strconv.ParseFloat(coordinate[3], 64)

			vertex = append(vertex, Vertex{x, y, z})
		} else if strings.HasPrefix(line, "f") {
			index := strings.Fields(line)

			i,_ := strconv.Atoi(index[1])
			j,_ := strconv.Atoi(index[2])
			k,_ := strconv.Atoi(index[3])

			// take vertex from vertex[], arr index starts from 0
			v1 := vertex[i-1]
			v2 := vertex[j-1]
			v3 := vertex[k-1]

			triangles = append(triangles, Triangle{v1, v2, v3})

		}
	}

	if err := scanner.Err(); err != nil {
        return nil, nil, err
    }

	return vertex, triangles, nil;
}

func saveObject (path string, vertex []Vertex, triangles []Triangle) {}