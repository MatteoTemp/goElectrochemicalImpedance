package main

import (
	circuits "goEl/curcuits"
	dipoles "goEl/lib"
)

func main() {

	circuit1 := circuits.Randles{
		Solution_rasistance:   dipoles.Resistor{Resistance: 0},
		Reaction_resistance:   dipoles.Resistor{Resistance: 500},
		Double_layer_capacity: dipoles.Capacitor{Capacitance: 1e-6},
		Diffusion_impedance:   dipoles.Warburg{W_sigma: 500},
	}
	//circuits.Nyquist_plot(circuit1, "OutputFiles/Nyq.csv", 0, 6, 100)
	circuits.Bode(circuit1, "OutputFiles/Bode.csv", -3, 6, 100)
	circuits.Lasajous(circuit1, 1)
	circuits.Lasajous(circuit1, 10)
	circuits.Lasajous(circuit1, 100)
	circuits.Lasajous(circuit1, 10000)

}
