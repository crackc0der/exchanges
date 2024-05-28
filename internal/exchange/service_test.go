package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExchange(t *testing.T) {
	t.Parallel()

	service := NewService()

	tests := []struct {
		name  string
		setup func()
		want  *ExchangeResult
	}{
		{
			name: "1",
			setup: func() {
				service.ServiceExchange(
					&Exchange{
						Amount:    400,
						Banknotes: []int{5000, 2000, 1000, 500, 200, 100, 50},
					},
				)
			},
			want: &ExchangeResult{
				Exchanges: [][]int{
					{200, 200},
					{200, 100, 100},
					{200, 100, 50, 50},
					{200, 50, 50, 50, 50},
					{100, 100, 100, 100},
					{100, 100, 100, 50, 50},
					{100, 100, 50, 50, 50, 50},
					{100, 50, 50, 50, 50, 50, 50},
					{50, 50, 50, 50, 50, 50, 50, 50},
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			testCase.setup()

			got := service.ServiceExchange(&Exchange{
				Amount:    400,
				Banknotes: []int{5000, 2000, 1000, 500, 200, 100, 50},
			})

			assert.Equal(t, testCase.want, got)
		})
	}
}
