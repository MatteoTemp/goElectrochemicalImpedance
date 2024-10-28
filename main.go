package main

import (
	circuits "goEl/curcuits"
	dipoles "goEl/lib"
	//cmath "math/cmplx"
)

func main() {

	equivalent := circuits.RCHighpass{
		Resistor:  dipoles.Resistor{Resistance: 100},
		Capacitor: dipoles.Capacitor{Capacitance: 1e-6},
	}

	circuits.Nyquist_plot(equivalent, "Nyq.csv", 8.0, 100)
	circuits.Bode(equivalent, "Bode.csv", 8.0, 100)
}
