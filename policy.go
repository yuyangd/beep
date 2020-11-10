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
	Interval             int8     `mapstructure:"interval"`
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

type ErrorBudget struct {
	budget   float64
	burnRate float64
}

func p2f(percent string) float64 {
	s := strings.TrimSuffix(percent, "%")
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic("Failed to convert percentage string to float")
	}
	return math.Round(f*100) / 10000
}

func (p *Policy) AvailabilityRate(c *Condition) {
	fmt.Printf("%v%% of Error Budget consumed in %v support hours", p2f(c.BudgetConsumed)*100, c.ConsumedHours)
	fmt.Println()

	errBudget := &ErrorBudget{
		budget:   1 - p2f(p.App.Slo),
		burnRate: p2f(c.BudgetConsumed) * float64(p.App.SupportHours/c.ConsumedHours) * float64(p.App.Interval),
	}

	errorRate := errBudget.burnRate * errBudget.budget
	ttd := float64(c.ConsumedHours) * 60 / (1 / errorRate)

	fmt.Printf("Error Budget Burn Rate: %v", errBudget.burnRate)
	fmt.Println()
	fmt.Printf("Error Rate > %.2f%% in %v hours, ", errorRate*100, c.ConsumedHours)
	fmt.Println()
	fmt.Printf("Health Check failure > %.2f mins in %v hours, ", ttd, c.ConsumedHours)
	fmt.Println()
	fmt.Printf("Full Outage TTD: %.2f mins", ttd)
	fmt.Println()

}

func (p *Policy) AlertCfg() {

	for _, c := range p.Conditions {
		p.AvailabilityRate(&c)
		fmt.Println()
	}

}
