package golang_united_school_homework

import "errors"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return errors.New("Can't add shape. Out of shapes capacity")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("Out of the range.")
	}
	// May be we never meet nil value, but i decide to add this here.
	if b.shapes[i] == nil {
		return nil, errors.New("No shape with this index")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	s, err := b.GetByIndex(i)
	if err != nil {
		return s, err
	}
	// Truncate slice for 1 element that we extract
	copy(b.shapes[i:], b.shapes[i+1:])
	b.shapes = b.shapes[:len(b.shapes)-1]
	return s, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	s, err := b.GetByIndex(i)
	if err != nil {
		return s, err
	}
	b.shapes[i] = shape
	return s, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var res float64
	for _, s := range b.shapes {
		res += s.CalcPerimeter()
	}
	return res
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var res float64
	for _, s := range b.shapes {
		res += s.CalcArea() 
	}
	return res
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	// Go through slice of shapes and Extract circle only, then shift index to the left by 1
	var exist bool = false
	for i := 0; i < len(b.shapes); i++ {
	     if _, ok := b.shapes[i].(*Circle); ok {
		     b.ExtractByIndex(i)
		     i -= 1
		     exist = true
	     }
        }
	if exist {
		return nil
	}
	return errors.New("There are no circles in box.")
}
