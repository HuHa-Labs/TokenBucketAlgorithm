package main

import "testing"

func TestRefillStrategyImplIsRefillStrategy(t *testing.T) {
	var _ RefillStrategy = &RefillStrategyImpl{}
}