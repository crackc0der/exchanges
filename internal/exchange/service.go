package exchange

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ServiceExchange(exc *Exchange) *ExchangeResult {
	result := s.findChange(exc.Amount, exc.Banknotes)

	return &ExchangeResult{result}
}

func (s *Service) findChange(mount int, banknotes []int) [][]int {
	if mount == 0 {
		return [][]int{{}}
	}

	if len(banknotes) == 0 || mount < 0 {
		return [][]int{}
	}

	var result [][]int

	for i, banknote := range banknotes {
		if banknote <= mount {
			ways := s.findChange(mount-banknote, banknotes[i:])
			for _, way := range ways {
				result = append(result, append([]int{banknote}, way...))
			}
		}
	}

	return result
}
