package main

type Car struct {
	brand          string  // toyota, mercedes
	model          string  // bz4x, cls400
	motorType      string  // electric, diesel
	motorPower     int     // 65, 330
	motorPowerUnit string  // kw , hp
	gearBox        string  // 1 speed, 6 speed
	rangeInKm      int     // 257, 1200
	maxSpeed       int     // 167,  250
	acceleration   float64 // 18.5, 5.2
}

func (c *Car) setBrand(brand string) {
	c.brand = brand
}
func (c *Car) getBrand() string {
	return c.brand
}

func (c *Car) setModel(model string) {
	c.model = model
}
func (c *Car) getModel() string {
	return c.model
}

func (c *Car) setMotorType(motorType string) {
	c.motorType = motorType
}
func (c *Car) getMotorType() string {
	return c.motorType
}

func (c *Car) setMotorPower(motorPower int) {
	c.motorPower = motorPower
}
func (c *Car) getMotorPower() int {
	return c.motorPower
}

func (c *Car) setMotorPowerUnit(motorPowerUnit string) {
	c.motorPowerUnit = motorPowerUnit
}

func (c *Car) getMotorPowerUnit() string {
	return c.motorPowerUnit
}

func (c *Car) setGearBox(gearBox string) {
	c.gearBox = gearBox
}
func (c *Car) getGearBox() string {
	return c.gearBox
}

func (c *Car) setRangeInKm(rangeInKm int) {
	c.rangeInKm = rangeInKm
}
func (c *Car) getRangeInKm() int {
	return c.rangeInKm
}

func (c *Car) setMaxSpeed(maxSpeed int) {
	c.maxSpeed = maxSpeed
}
func (c *Car) getMaxSpeed() int {
	return c.maxSpeed
}

func (c *Car) setAcceleration(acceleration float64) {
	c.acceleration = acceleration
}

func (c *Car) getAcceleration() float64 {
	return c.acceleration
}
