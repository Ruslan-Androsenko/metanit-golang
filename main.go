package main

import "fmt"

type Vehicle interface {
	move()
}

func drive(vehicle Vehicle) {
	vehicle.move()
}

// структура "Автомобиль"
type Car struct{}

// струкрура "Самолет"
type Aircraft struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}

func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func main() {
	tesla := Car{}
	boeing := Aircraft{}

	drive(tesla)
	drive(boeing)
}
