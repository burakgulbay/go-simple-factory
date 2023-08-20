package main

type ICar interface {
	setBrand(brand string)
	getBrand() string

	setModel(model string)
	getModel() string

	setMotorType(motorType string)
	getMotorType() string

	setMotorPower(motorPower int)
	getMotorPower() int

	setMotorPowerUnit(motorPowerUnit string)
	getMotorPowerUnit() string

	setGearBox(gearBox string)
	getGearBox() string

	setRangeInKm(rangeInKm int)
	getRangeInKm() int

	setMaxSpeed(maxSpeed int)
	getMaxSpeed() int

	setAcceleration(acceleration float64)
	getAcceleration() float64

	printCar() string
}
