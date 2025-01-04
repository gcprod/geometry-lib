package shape

type Rectangle struct {
	Width  float64
	Length float64
}

func (rectangle *Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Length
}
func (rectangle *Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Width + rectangle.Length)
}
