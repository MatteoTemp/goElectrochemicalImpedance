package main

import (
	circuits "goEl/curcuits"
	dipoles "goEl/lib"
)

func main() {

	circuit1 := circuits.NIRandles{
		Solution_rasistance: dipoles.Resistor{Resistance: 10},
		Reaction_resistance: dipoles.Resistor{Resistance: 500},
		Diffusion_impedance: dipoles.Warburg{W_sigma: 500},
		Double_layer_capacity: dipoles.Constant_phase_Element{
			Phi:        0.99,
			TParameter: 1e-6,
		},
	}
	//circuits.Nyquist_plot(circuit1, "OutputFiles/Nyq.csv", 0, 6, 100)
	circuits.Bode(circuit1, "OutputFiles/Bode.csv", -3, 6, 100)
	circuits.Nyquist_plot(circuit1, "OutputFiles/Nyq.csv", -0.3, 6, 1000, false)
}
