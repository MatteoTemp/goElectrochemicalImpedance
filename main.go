package main

import (
	circuits "goEl/curcuits"
	dipoles "goEl/lib"
	//cmath "math/cmplx"
)

func main() {

	equivalent1 := circuits.Randles{
		Solution_rasistance:   dipoles.Resistor{Resistance: 100},
		Reaction_resistance:   dipoles.Resistor{Resistance: 1000},
		Double_layer_capacity: dipoles.Capacitor{Capacitance: 1e-6},
		Diffusion_impedance:   dipoles.Warburg{W_sigma: 1000},
	}

	circuits.Nyquist_plot(equivalent1, "OutputFiles/Nyq.csv", -2, 3.0, 1000)
	circuits.Bode(equivalent1, "OutputFiles/Bode.csv", -6, 6.0, 1000)

}
