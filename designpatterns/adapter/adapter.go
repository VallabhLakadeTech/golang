package main

type ExchangeAdapter struct {
	client *IndianClient
}

func (adapter *ExchangeAdapter) payInDollars(amount int) int {

	return adapter.client.payInRupees(amount) / 83

}
