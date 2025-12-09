package main

import "fmt"

// The Prototype Design Pattern is a creational design pattern used to create new
// objects by copying (cloning) an existing object, instead of creating them from scratch using new.

type ShapType int

const (
	SquareType ShapType = 1
	RectType   ShapType = 2
)

type Shape interface {
	GetID() ShapType
	Area()
	Clone() Shape
}

type Square struct {
	ID   ShapType
	Side int
}

func NewSquare(side int) *Square {
	return &Square{
		ID:   SquareType,
		Side: side,
	}
}

func (sq *Square) Area() {
	fmt.Printf("Area of circler %d:%d\n", sq.Side, sq.Side*sq.Side)
}

func (sq *Square) GetID() ShapType {
	return sq.ID
}

func (sq *Square) Clone() Shape {
	return NewSquare(sq.Side)
}

type Rectangle struct {
	ID      ShapType
	Breadth int
	Width   int
}

func NewRectangle(b, w int) *Rectangle {
	return &Rectangle{
		ID:      RectType,
		Breadth: b,
		Width:   w,
	}
}

func (r *Rectangle) Area() {
	fmt.Printf("Area of rectangle %dX%d is %d\n", r.Breadth, r.Width, r.Breadth*r.Width)
}

func (r *Rectangle) GetID() ShapType {
	return r.ID
}

func (r *Rectangle) Clone() Shape {
	return NewRectangle(r.Breadth, r.Width)
}



func PrototypeRegistry()map[ShapType]Shape {

    registry:=make(map[ShapType]Shape)
	sq := NewSquare(4)
	registry[SquareType] = sq
	rc := NewRectangle(2, 2)
	registry[RectType] = rc
   return registry
}

func main() {
    registry:=PrototypeRegistry()
	defSq:=registry[SquareType].(*Square)
	defSq.Area()
	clone:=defSq.Clone()
    clone.Area()

}
