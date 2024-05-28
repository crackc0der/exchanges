package exchange

type Exchange struct {
	Amount    int   `json:"amount"`
	Banknotes []int `json:"banknotes"`
}

type ExchangeResult struct {
	Exchanges [][]int `json:"exchanges"`
}
