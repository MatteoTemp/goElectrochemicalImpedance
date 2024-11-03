package main

import (
	circuits "goEl/curcuits"
	dipoles "goEl/lib"
	//cmath "math/cmplx"
)

func main() {

	equivalent := circuits.RealLCSeries{
		R: dipoles.Resistor{Resistance: 10},
		C: dipoles.Capacitor{Capacitance: 1e-6},
		L: dipoles.Inductor{Inductance: 1e-3},
	}

	circuits.Nyquist_plot(equivalent, "OutputFiles/Nyq.csv", 0, 6.0, 100)
	circuits.Bode(equivalent, "OutputFiles/Bode.csv", 2, 6.0, 100)
}
