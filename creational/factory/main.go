package main

import "fmt"

type iCar interface {
	setName(name string)
	SetHP(hp int)
	getName() string
	getHP() int
}

type car struct {
	name string
	hp   int
}

func (c *car) setName(name string) {
	c.name = name
}

func (c *car) getName() string {
	return c.name
}

func (c *car) SetHP(hp int) {
	c.hp = hp
}

func (c *car) getHP() int {
	return c.hp
}

type car1 struct {
	car
}

func newCar1() iCar {
	return &car1{
		car: car{
			name: "Car 1",
			hp:   500,
		},
	}
}

type car2 struct {
	car
}

func newCar2() iCar {
	return &car2{
		car: car{
			name: "Car 2",
			hp:   750,
		},
	}
}

func getCar(carNo string) (iCar, error) {

	if carNo == "car1" {
		return newCar1(), nil
	}
	if carNo == "car2" {
		return newCar2(), nil
	}
	return nil, fmt.Errorf("Wrong car")
}

func main() {
	car1, _ := getCar("car1")
	car2, _ := getCar("car2")

	printDetails(car1)
	printDetails(car2)

}

func printDetails(c iCar) {
	fmt.Printf("Car: %s", c.getName())
	fmt.Println()
	fmt.Printf("Horsepower: %d", c.getHP())
	fmt.Println()
	fmt.Println()
}
