package main

import (
	circuits "goEl/curcuits"
	dipoles "goEl/lib"
)

func main() {

	circuit := circuits.IdealNPElectrode{
		SolutionResistance: dipoles.Resistor{Resistance: 1000},
		InterphaseCapacity: dipoles.Capacitor{Capacitance: 1e-6},
	}
	circuits.Nyquist_plot(circuit, "OutputFiles/Nyq.csv", -3, 9, 100)
	circuits.Bode(circuit, "OutputFiles/Bode.csv", -3, 9, 100)

}
