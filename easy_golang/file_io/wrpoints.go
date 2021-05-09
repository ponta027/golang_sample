package main

import (
	"fmt"
	"log"
	"os"
)

type Point struct{ X, Y int }

const fname = "points.data"

func Save(points []Point) int {
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, p := range points {
		fmt.Fprintf(file, "%d %d\n", p.X, p.Y)
	}
	return len(points)
}

/**
 */
func Load() []Point {
	var points = []Point{}
	file, err := os.OpenFile(fname, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var x, y int
	for {
		n, err := fmt.Fscanf(file, "%d %d", &x, &y)
		if n == 0 || err != nil {
			break
		} else {
			points = append(points, Point{x, y})
		}
	}
	return points

}
func main() {
	var points = []Point{}
	points = append(points, Point{5, 5})
	points = append(points, Point{7, 8})
	points = append(points, Point{4, 3})
	points = append(points, Point{7, 4})
	points = append(points, Point{9, 5})
	fmt.Println(" writing data..")
	for i, p := range points {
		fmt.Printf("point [%d]=(%d,%d)\n", i+1, p.X, p.Y)
	}
	len := Save(points)
	fmt.Printf("%d counts.\n", len)
	data := Load()
	fmt.Printf("read data\n")
	for i, p := range data {
		fmt.Printf("point [%d]=(%d,%d)\n", i+1, p.X, p.Y)
	}

}
