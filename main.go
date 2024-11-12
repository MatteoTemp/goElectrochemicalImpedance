package main

import (
	circuits "goEl/curcuits"
	dipoles "goEl/lib"
	//cmath "math/cmplx"
)

func main() {

	ew := circuits.RCLBandpass{
		R: dipoles.Resistor{Resistance: 10},
		C: dipoles.Capacitor{Capacitance: 1e-9},
		L: dipoles.Inductor{Inductance: 1e-1},
	}

	circuits.Nyquist_plot(ew, "OutputFiles/Nyq.csv", -0, 6, 100)
	circuits.Bode(ew, "OutputFiles/Bode.csv", -3, 12, 50)

}
