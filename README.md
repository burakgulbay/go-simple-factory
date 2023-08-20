# go-simple-factory
Implementation of simple factory pattern using Go

# Background

Simple Factory is one of creational patterns in OOP design. 

# Motivation

Since Go is not exactly an object oriented programming language, It is not possible to implement a <b>Factory Method<b> pattern to create objects. <br>
We do not have classes and inheritance in Go. <br>
Though we can still implement <b>Simple Factory<b> pattern.

# Structure
[Title](icar.go)
Contains interface defination of ICar. 
[Title](car.go)
Contains base struct of Car with getters and setters for base struct's properties 
[Title](electric_car.go)   
Concrete struct of electric car. Implements actual creation of electric car. 
[Title](petrol_car.go)
Concrete struct of petrol car. Implements actual creation of petrol car. 
[Title](car_factory.go)
Contains an abstract createCar function to call actual create functions in concrete structsa. Returns a struct of type ICar.
[Title](main.go)
Contains sample client code to create electric and petrol cars. 
