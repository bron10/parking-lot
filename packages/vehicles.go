package packages

import (
	"github.com/bron10/parking-lot/services"
	"strings"
	"time"
)

func UnparkVehicle(lot int) (int, string) {
	vehicleId := services.RemoveVehicle(lot)
	age, regNum := services.GetVehicleInfoById(vehicleId)
	return age, regNum
}

func CreateVehicleEntry(allocatedLot int, regNum string, driverAge int) {
	now := time.Now() // current local time
	timeStamp := now.Unix()
	services.CreateVehicleEntry(allocatedLot, regNum, driverAge, timeStamp)
}

func GetRegNumByAge(age int) string {
	regNums := services.GetRegNumByAge(age)
	return strings.Join(regNums, ",")
}
