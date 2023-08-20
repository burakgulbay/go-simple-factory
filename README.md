
# go-simple-factory
Implementation of simple factory pattern using Go

## Background
Simple Factory is one of creational patterns in OOP design.

## Motivation
Since Go is not exactly an object oriented programming language, It is not possible to implement a **Factory Method** pattern to create objects. <br>

We do not have classes and inheritance in Go. <br>

Though we can still implement **Simple Factory** pattern.

## Structure
  

**<i>icar.go<i>**
Contains interface defination of ICar.

**<i>car.go<i>**
Contains base struct of Car with getters and setters for base struct's properties.

**<i>electric_car.go<i>**
Concrete struct of electric car. Implements actual creation of electric car.

**<i>petrol_car.go<i>**
Concrete struct of petrol car. Implements actual creation of petrol car.

**<i>car_factory.go<i>**
Contains an abstract createCar function to call actual create functions in concrete structsa. Returns a struct of type ICar.

**<i>main.go<i>**
Contains sample client code to create electric and petrol cars.