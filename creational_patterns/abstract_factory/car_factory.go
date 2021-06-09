package abstract_factory

import "errors"

const (
	LuxuryCarType=1
	FamilyCarType=2
)

type CarFactory struct {}

func (c *CarFactory) GetVehicle(v int) (Vehicle,error)  {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New("vehicle of type was not recognized")
	}
}
