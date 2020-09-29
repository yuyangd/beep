package beep

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Policy struct {
	App             App              `mapstructure:"app"`
	Conditions      []Condition      `mapstructure:"conditions"`
	KeyTransactions []KeyTransaction `mapstructure:"keyTransactions"`
}

type App struct {
	Slo                  string   `mapstructure:"slo"`
	SupportHours         int8     `mapstructure:"supportHours"`
	ExcludedTransactions []string `mapstructure:"excludedTransactions"`
}

type Condition struct {
	BudgetConsumed string `mapstructure:"budgetConsumed"`
	ConsumedHours  int8   `mapstructure:"consumedHours"`
}

type KeyTransaction struct {
	Name     string `mapstructure:"name"`
	Policies []struct {
		ErrRateAbove string `mapstructure:"errRateAbove"`
		InMins       int8   `mapstructure:"inMins"`
	} `mapstructure:"policies"`
}

func p2f(percent string) float64 {
	s := strings.TrimSuffix(percent, "%")
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic("Failed to convert percentage string to float")
	}
	return math.Round(f*100) / 10000
}

func (p *Policy) Rate() {
	for _, c := range p.Conditions {
		fmt.Printf("%v%% of Error Budget consumed in %v support hours", p2f(c.BudgetConsumed)*100, p.App.SupportHours)
		fmt.Println()
		budget := 1 - p2f(p.App.Slo)

		burnRate := p2f(c.BudgetConsumed) * float64(p.App.SupportHours/c.ConsumedHours) * 30

		fmt.Printf("Error Budget Burn Rate: %v", burnRate)
		fmt.Println()
		errorRate := burnRate * budget

		fmt.Printf("Error Rate > %.2f%% in %v hours, ", errorRate*100, c.ConsumedHours)

		fmt.Printf("Full Outage TTD: %.2f mins", float64(c.ConsumedHours)*60/(1/errorRate))
		fmt.Println()
	}

}
