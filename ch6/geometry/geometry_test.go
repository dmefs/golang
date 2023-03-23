package geometry_test

import (
	"fmt"
	"testing"
	"time"

	"gopl.io/ch6/geometry"
)

func TestGeo(t *testing.T) {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}
	distance := geometry.Point.Distance
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*geometry.Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)
	const day = 24 * time.Hour
	fmt.Println(day.Seconds()) // "86400"
}
