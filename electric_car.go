package main

import "fmt"

type ElectricCar struct {
	Car
}

func newElectricCar() ICar {
	return &ElectricCar{
		Car: Car{
			brand:          "Toyota",
			model:          "bz4x", // 😏 who cares
			motorType:      "electric",
			motorPower:     65,
			motorPowerUnit: "Kw",
			gearBox:        "1 Speed",
			rangeInKm:      257,
			maxSpeed:       167,
			acceleration:   18.5,
		},
	}
}

func (ea *ElectricCar) printCar() string {
	return fmt.Sprintf("Brand: %v\nModel: %v\nMotor Type: %v\nMotor Power: %d\nMotor Power Unit: %v\nGear Box: %v\nRange in Km: %d\nMaximum Speed: %d\nAcceleration: %.1f seconds (0-100 km/s)\n", ea.getBrand(), ea.getModel(), ea.getMotorType(), ea.getMotorPower(), ea.getMotorPowerUnit(), ea.getGearBox(), ea.getRangeInKm(), ea.getMaxSpeed(), ea.getAcceleration())
}
