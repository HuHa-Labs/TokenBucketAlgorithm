package main

type RefillStrategy interface {
	Refill() uint32
	GetUntilNextRefill(timeUnit string)
}

type RefillStrategyImpl struct {
}

func (rs RefillStrategyImpl) Refill() uint32 {
	return 0
}

func (rs RefillStrategyImpl) GetUntilNextRefill(timeUnit string) {	

}
