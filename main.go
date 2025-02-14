package main

import (
	"fmt"
	"goEl/kinetics"
)

func main() {
	system := kinetics.ButlerVolmer{
		Alpha:                 0.5,
		K0:                    1e4,
		Equilibrium_potential: 0.0,
		Cbulk:                 1e-3,
		CatodicLimitCurrent:   1e-5,
		AnodicLimitCurrent:    1e-5,
	}

	fmt.Println(system.EquilibriumCurrent())

	fmt.Println(system.CurrentUnlimited(1.0))
	fmt.Println(system.CurrentUnlimited(-1.0))
	fmt.Println(system.CurrentLimited(2.0))
	fmt.Println(system.CurrentLimited(-2.0))

}
