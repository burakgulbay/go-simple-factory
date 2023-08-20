package main

import (
	"fmt"
	"log"
)

func main() {

	dieselCar, err := createCar("diesel")
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println(dieselCar.printCar()) // ♥️

	electricCar, err := createCar("electric")
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println(electricCar.printCar())

}
