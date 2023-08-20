package main

import "errors"

// guess what
func createCar(motorType string) (ICar, error) {
	if motorType == "electric" {
		return newElectricCar(), nil
	} else if motorType == "diesel" {
		return newPetrolCar(), nil
	}
	return nil, errors.New("invalid motor type")
}
