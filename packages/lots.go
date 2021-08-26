package packages

import (
	"github.com/bron10/parking-lot/services"
	"time"
)

func GetData() string {
	return "hello in lots"
}

func CreateLots(slots int) int {
	now := time.Now() // current local time
	sec := now.Unix()
	for i := 1; i <= slots; i++ { // counter-controlled loop
		// lot := {id: i, status: "open", created_date: sec, updated_date: sec}
		services.CreateLot(i, sec)
	}
	return slots
}

func AllocateLot(regNum string, driverAge int) int {
	nearestLotId := services.GetNearestLot("open")
	services.ChangeLotStatus(nearestLotId, "allocated")
	return nearestLotId
}

func bookLot(lotId int) {
	services.ChangeLotStatus(lotId, "booked")
}

func GetLotsByAge(age int) string {
	lots := services.GetLotsByAge(age)
	return ArrayToString(lots, ",")
}

func GetLotByRegNum(regNum string) int {
	lot := services.GetLotByRegNum(regNum)
	return lot
}

func LeaveLot(lot int) bool {
	findStatus := "open"
	services.ChangeLotStatus(lot, findStatus)
	status := services.GetStatusByLot(lot)
	return status == findStatus
}
