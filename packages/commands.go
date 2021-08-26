package packages

import (
	"errors"
	"strconv"
	"strings"
)

func Serve(commands []string) (string, error) {
	switch commands[0] {
	case CREATE_LOTS:
		opt := commands[1]
		slots, _ := strconv.Atoi(opt)
		CreateLots(slots)
		res := map[string]string{"<slotNum>": opt}
		return createResponse(RES_CREATE_LOTS, res), nil
	case PARK:
		regNum := commands[1]
		driverAge, _ := strconv.Atoi(commands[3])
		allocatedLot := AllocateLot(regNum, driverAge)
		CreateVehicleEntry(allocatedLot, regNum, driverAge)
		bookLot(allocatedLot)
		allocatedLotStr := strconv.Itoa(allocatedLot)
		res := map[string]string{"<regNum>": regNum, "<slotNum>": allocatedLotStr}
		return createResponse(RES_PARK, res), nil
	case GET_SLOT_NUM_DRIVER_AGE:
		driverAge, _ := strconv.Atoi(commands[1])
		lots := GetLotsByAge(driverAge)
		if len(lots) == 0 {
			return RES_NO_MATCH, nil
		}
		return lots, nil
	case GET_SLOT_NUM_CAR_NUM:
		regNum := commands[1]
		lot := GetLotByRegNum(regNum)
		if lot == 0 {
			return RES_NO_MATCH, nil
		}
		lotStr := strconv.Itoa(lot)
		return lotStr, nil
		// return RES_PARK, nil
	case LEAVE_LOT:
		lotStr := commands[1]
		lot, _ := strconv.Atoi(lotStr)
		age, regNum := UnparkVehicle(lot)
		left := LeaveLot(lot)
		ageStr := strconv.Itoa(age)
		if left {
			res := map[string]string{"<regNum>": regNum, "<slotNum>": lotStr, "<age>": ageStr}
			return createResponse(RES_LEAVE, res), nil
		}

		return RES_NO_MATCH, nil
	case GET_VEHICLE_REG_NUM_DRIVER_AGE:
		ageStr := commands[1]
		age, _ := strconv.Atoi(ageStr)
		regNums := GetRegNumByAge(age)
		if len(regNums) != 0 {
			return regNums, nil
		}
		return RES_NO_MATCH, nil
	default:
		return "", errors.New("please try with different command")
	}
}

func createResponse(str string, res map[string]string) string {
	for key, value := range res {
		str = strings.Replace(str, key, value, 1)
	}
	return str
}
