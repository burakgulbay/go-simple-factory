package main

import "fmt"

type PetrolCar struct {
	Car
	fuelConsumption float64
}

func newPetrolCar() ICar {
	return &PetrolCar{
		Car: Car{
			brand:          "Mercedes",
			model:          "Cls400d", // ♥️ oh my lovely dear
			motorType:      "diesel",
			motorPower:     330,
			motorPowerUnit: "Hp",
			gearBox:        "6 Speed",
			rangeInKm:      1200,
			maxSpeed:       250,
			acceleration:   5.2,
		},
		fuelConsumption: 7.8,
	}
}
func (pc *PetrolCar) printCar() string {
	return fmt.Sprintf("Brand: %v\nModel: %v\nMotor Type: %v\nMotor Power: %d\nMotor Power Unit: %v\nGear Box: %v\nRange in Km: %d\nMaximum Speed: %d\nAcceleration: %.1f seconds (0-100 km/s)\nFuel Consumption: %.1f\n", pc.getBrand(), pc.getModel(), pc.getMotorType(), pc.getMotorPower(), pc.getMotorPowerUnit(), pc.getGearBox(), pc.getRangeInKm(), pc.getMaxSpeed(), pc.getAcceleration(), pc.getFuelConsumption())
}

func (pc *PetrolCar) getFuelConsumption() float64 {
	return pc.fuelConsumption
}
func (pc *PetrolCar) setFuelConsumption(fuelConsumption float64) {
	pc.fuelConsumption = fuelConsumption
}
