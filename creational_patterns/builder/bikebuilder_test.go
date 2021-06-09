package builder

import "testing"

func TestBikeBuilder(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	car := bikeBuilder.GetVehicle()

	if car.Wheels != 2 {
		t.Errorf("Wheels on the Bike must be 4 and they were %d\n", car.Wheels)
	}
	if car.Structure != "MotorBike" {
		t.Errorf("Structure on the Bike must be 'MotorBike' and they were %s\n", car.Structure)
	}

	if car.Seats != 1 {
		t.Errorf("Seats on the Bike must be 1 and they were %d\n", car.Seats)
	}
}
