package main

import (
	"fmt"
	 "math"
	)

//Movable - object with this interface can move, returns distance or -1 if cannot move
type Movable interface{
	Move(int ,int) int
}
//Talkable - object can talk
type Talkable interface{
	Talk(string)
}
//Object - base object
type Object struct{
	x,y int
}
type ObjectCircle struct{
	x,y int
	r int
}

//Move object to new place
func (u *Object) Move(x, y int) int{
	s := int(math.Sqrt((float64)( (x-u.x)*(x-u.x)+(y-u.y)*(y-u.y))))
	u.x = x
	u.y = y
	return s
}
func (u *ObjectCircle) Move(x, y int) int{
	s := int(math.Sqrt((float64)( (x-u.x)*(x-u.x)+(y-u.y)*(y-u.y))))
	u.x = x
	u.y = y
	return s
}
func (u Object) String() string{
	return fmt.Sprintf("%d;%d", u.x, u.y)
}
//MoveWorld00 - moves all objects to position 00
func MoveWorld00(objects ...Movable){
	for _, obj :=range objects{
		obj.Move(0,0)
	}
}



func main(){
	fmt.Println("interfaces")
	u1 := &Object{x:10, y:10}
	u2 := &Object{x:15, y:25}
	fmt.Println(u1, u2)
	
	u1.Move(44,11)
	fmt.Println(u1, u2)
	
	MoveWorld00(u1,u2)
	fmt.Println(u1, u2)

}
