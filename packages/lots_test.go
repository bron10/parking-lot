package packages

import (
	"github.com/bron10/parking-lot/services"
	"testing"
)

func TestCreateLots(t *testing.T) {
	lots := 6
	resultLots := CreateLots(lots)
	if resultLots == lots {
		t.Log("OK: Input lots same as output lots")
	} else {
		t.Fail()
	}

	lots = 0
	resultLots = CreateLots(lots)
	if resultLots == lots {
		t.Log("OK: Input lots should be same output lots")
	} else {
		t.Fail()
	}
}

func TestAllocateLot(t *testing.T) {

	AllocateLot := AllocateLot("KA-01-HH-1234", 21)
	if AllocateLot > 0 {
		t.Log("OK: Lot is allocated")
	} else {
		t.Fail()
	}
}

func TestbookLot(t *testing.T) {
	lots := 2
	CreateLots(lots)
	lot := 1
	bookLot(1)
	status := services.GetStatusByLot(lot)
	if status == "booked" {
		t.Log("OK: Lot is allocated")
	} else {
		t.Fail()
	}
}
