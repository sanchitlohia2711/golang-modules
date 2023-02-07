package main

type ParkingRateFactory struct {
	parkingRateMap map[VehicleType]ParkingRate
}

func (this *ParkingRateFactory) getParkingRateInstanceByVehicleType(vType VehicleType) ParkingRate {
	if this.parkingRateMap[vType] != nil {
		return this.parkingRateMap[vType]
	}
	if vType == car {
		this.parkingRateMap[vType] = &CarParkingRate{
			ParkingRateBase{perHourCharges: 2},
		}
		return this.parkingRateMap[vType]
	}
	if vType == truck {
		this.parkingRateMap[vType] = &TruckParkingRate{
			ParkingRateBase{perHourCharges: 3},
		}
		return this.parkingRateMap[vType]
	}
	if vType == motorcycle {
		this.parkingRateMap[vType] = &MotorCycleParkingRate{
			ParkingRateBase{perHourCharges: 1},
		}
		return this.parkingRateMap[vType]
	}
	return nil
}

type ParkingRateBase struct {
	perHourCharges int
}

type ParkingRate interface {
	amountToPay(int) int
}

type CarParkingRate struct {
	ParkingRateBase
}

func (this *CarParkingRate) amountToPay(hours int) int {
	return this.perHourCharges * hours
}

type TruckParkingRate struct {
	ParkingRateBase
}

func (this *TruckParkingRate) amountToPay(hours int) int {
	return this.perHourCharges * hours
}

type MotorCycleParkingRate struct {
	ParkingRateBase
}

func (this *MotorCycleParkingRate) amountToPay(hours int) int {
	return this.perHourCharges * hours
}
