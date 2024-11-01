package main

import (
	circuits "goEl/curcuits"
	dipoles "goEl/lib"
	//cmath "math/cmplx"
)

func main() {

	equivalent := circuits.RCBandPass{
		R1: dipoles.Resistor{Resistance: 10000},
		C1: dipoles.Capacitor{Capacitance: 10e-9},

		R2: dipoles.Resistor{Resistance: 500},
		C2: dipoles.Capacitor{Capacitance: 1e-9},
	}

	circuits.Nyquist_plot(equivalent, "OutputFiles/Nyq.csv", 0.0, 8.0, 100)
	circuits.Bode(equivalent, "OutputFiles/Bode.csv", 0.0, 8.0, 100)
}
