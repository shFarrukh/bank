package types

//money is kept in minimum currencies (cent kopeck diram)
type Money int64

//payment categories (foods , clothes...)
type Category string

type Status string

//Предопределенный статус платежей
const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInPrigress Status = "INPROGRESS"
)

// Paymant information
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}




