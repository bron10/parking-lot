package services

type Lot struct {
	id           int
	status       string //booked, open, allocated
	created_date int64
	updated_date int64
}

var LotStore []Lot

type Vehicle struct {
	id           int
	age          int
	reg_num      string
	lot_id       int
	inn          bool //default is false
	created_date int64
	updated_date int64
}

var VehicleStore []Vehicle
