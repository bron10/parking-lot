package services

func GetNearestLot(status string) int {
	lots := GetLots()
	lot := findByStatus(lots, "open")
	return lot.id
}

func CreateLot(id int, date int64) {
	lot := Lot{id: id, status: "open", created_date: date, updated_date: date}
	LotStore = append(LotStore, lot)
}

func GetLots() []Lot {
	return LotStore
}

func GetStatusByLot(lotId int) string {
	var found Lot
	for _, v := range LotStore {
		if v.id == lotId {
			found = v
			break
		}
	}

	return found.status
}

func findByStatus(lots []Lot, status string) Lot {
	var found Lot
	for _, v := range lots {
		if v.status == status {
			found = v
			break
		}
	}

	return found
}

func findIndexById(lots []Lot, id int) int {
	var index = len(lots) + 1
	for i, v := range lots {
		if v.id == id {
			index = i
			break
		}
	}

	return index
}

func CreateVehicleEntry(lotId int, regNum string, driverAge int, date int64) int {
	// nearestLotId := GetNearestLot("open")
	var vehicleIndex int = len(VehicleStore) + 1
	vehicleData := Vehicle{
		id:           vehicleIndex,
		age:          driverAge,
		reg_num:      regNum,
		lot_id:       lotId,
		inn:          true,
		created_date: date,
		updated_date: date}
	VehicleStore = append(VehicleStore, vehicleData)
	return vehicleIndex
}

func ChangeLotStatus(lotId int, status string) {
	lotIndex := findIndexById(LotStore, lotId)
	lot := LotStore[lotIndex]
	lot.status = status
	LotStore[lotIndex] = lot
}

func ByAge(item Vehicle, age int) bool {
	if item.age == age && item.inn {
		return true
	} else {
		return false
	}
}

func GetLotsByAge(age int) []int {
	lots := make([]int, 0)
	for _, v := range VehicleStore {
		if ByAge(v, age) {
			lots = append(lots, v.lot_id)
		}
	}
	return lots
}

func GetLotByRegNum(regNum string) int {
	lotId := 0
	for _, v := range VehicleStore {
		if v.reg_num == regNum && v.inn {
			return v.lot_id
		}
	}
	return lotId
}

func RemoveVehicle(lot int) int {
	vehicleId := 0
	for k, v := range VehicleStore {
		if v.lot_id == lot && v.inn {
			vehicleId = v.id
			VehicleStore[k].inn = false
			break
		}
	}

	return vehicleId
}

func GetVehicleInfoById(id int) (int, string) {
	age := 0
	regNum := ""
	for _, v := range VehicleStore {
		if v.id == id {
			age = v.age
			regNum = v.reg_num
			break
		}
	}

	return age, regNum
}

func GetRegNumByAge(age int) []string {
	regNums := make([]string, 0)

	for _, v := range VehicleStore {
		if ByAge(v, age) {
			regNums = append(regNums, v.reg_num)
		}
	}
	return regNums
}
