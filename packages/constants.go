package packages

//Commands
const CREATE_LOTS string = "Create_parking_lot"
const PARK string = "Park"
const GET_SLOT_NUM_DRIVER_AGE string = "Slot_numbers_for_driver_of_age"
const GET_SLOT_NUM_CAR_NUM string = "Slot_number_for_car_with_number"
const LEAVE_LOT string = "Leave"
const GET_VEHICLE_REG_NUM_DRIVER_AGE string = "Vehicle_registration_number_for_driver_of_age"

//Responses
const RES_CREATE_LOTS string = "Created parking of <slotNum> slots"
const RES_PARK string = "Car with vehicle registration number <regNum> has been parked at slot number <slotNum>"
const RES_NO_MATCH string = "No parked car matches the query"
const RES_LEAVE string = "Slot number <slotNum> vacated, the car with vehicle registration number <regNum> left the space, the driver of the car was of age <age>"
