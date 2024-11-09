package transaction

type OperationType int

const (
	Normal OperationType = iota + 1
	Installment
	Withdraw
	CreditVoucher
)
