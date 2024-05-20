package main

type RefillStrategy interface {
	Refill() uint32
	GetUntilNextRefill(timeUnit string)
}

type RefillStrategyImpl struct {
	numTokensPerPeriod uint32
	periodDurationInNanos uint32
	lastRefillTime uint32
	nextRefillTime uint32
}

func (rs RefillStrategyImpl) Refill() uint32 {
	return 0
}

func (rs RefillStrategyImpl) GetUntilNextRefill(timeUnit string) {	

}
