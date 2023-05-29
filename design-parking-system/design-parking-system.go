package designparkingsystem

type ParkingSystem struct {
	Big int
	Mid int
	Sml int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		Big: big,
		Mid: medium,
		Sml: small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.Big > 0 {
			this.Big--
			return true
		}
	case 2:
		if this.Mid > 0 {
			this.Mid--
			return true
		}
	case 3:
		if this.Sml > 0 {
			this.Sml--
			return true
		}
	}
	return false
}
