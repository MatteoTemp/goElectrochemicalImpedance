package main

import (
	circuits "goEl/curcuits"
	dipoles "goEl/lib"
)

func main() {

	circuit1 := circuits.NonIdealNPElectrode{
		SolutionResistance: dipoles.Resistor{Resistance: 100},
		ImperfectDL:        dipoles.Constant_phase_Element{TParameter: 0.01, Phi: 0.9999},
	}
	//circuits.Nyquist_plot(circuit1, "OutputFiles/Nyq.csv", 0, 6, 100)
	circuits.Bode(circuit1, "OutputFiles/Bode.csv", -3, 3, 100)
	circuits.Nyquist_plot(circuit1, "OutputFiles/Nyq.csv", -0.3, 3, 1000, false)
}
