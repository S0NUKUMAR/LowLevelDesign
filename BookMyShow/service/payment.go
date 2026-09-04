package service

type PaymentGateway interface {
	Charge (userID string, amount float64) (transactionID string, err error)
}

type MockPaymentGateway struct {
	AlwaysFail bool
}

func (m *MockPaymentGateway) Charge(userID string, amount float64)(string,error){
	if m.AlwaysFail {
		return "",ErrPaymentDeclined 
	}
	return "txn_" + userID + "_mock", nil
}

