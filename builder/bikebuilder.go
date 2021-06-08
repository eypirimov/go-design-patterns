package builder

type BikeBuilder struct {
	v VehicleProduce
}

func (c *BikeBuilder) SetSeats() BuildProcess {
	c.v.Seats = 1
	return c
}

func (c *BikeBuilder) SetStructure() BuildProcess {
	c.v.Structure = "MotorBike"
	return c
}

func (c *BikeBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 2
	return c
}

func (c *BikeBuilder) GetVehicle() VehicleProduce {
	return c.v
}
