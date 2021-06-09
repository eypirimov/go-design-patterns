package abstract_factory

import "errors"

const (
	SportMotorbikeType=1
	CruiseMotorbikeType=2
)

type MotorbikeFactory struct {}

func (c *MotorbikeFactory) GetVehicle(v int) (Vehicle,error)  {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New("vehicle of type was not recognized")
	}
}
