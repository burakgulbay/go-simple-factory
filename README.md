# go-simple-factory
Implementation of simple factory pattern using Go

# Background

Simple Factory is one of creational patterns in OOP design. 

# Motivation

Since Go is not exactly an object oriented programming language, It is not possible to implement a <b>Factory Method<b> pattern to create objects. <br>
We do not have classes and inheritance in Go. <br>
Though we can still implement <b>Simple Factory<b> pattern.

# Structure

<i>icar.go<i><br>
Contains interface defination of ICar. <br><br>
<i>car.go<i><br>
Contains base struct of Car with getters and setters for base struct's properties <br><br>
<i>electric_car.go<i> <br>
Concrete struct of electric car. Implements actual creation of electric car. <br><br>
<i>petrol_car.go<i><br>
Concrete struct of petrol car. Implements actual creation of petrol car. <br><br>
<i>car_factory.go<i><br>
Contains an abstract createCar function to call actual create functions in concrete structsa. Returns a struct of type ICar.<br><br>
<i>main.go<i><br>
Contains sample client code to create electric and petrol cars.
