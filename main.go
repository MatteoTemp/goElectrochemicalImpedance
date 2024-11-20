package main

import (
	circuits "goEl/curcuits"
	dipoles "goEl/lib"
)

func main() {

	circuit1 := circuits.RandlesTB{
		Solution_rasistance:   dipoles.Resistor{Resistance: 100},
		Reaction_resistance:   dipoles.Resistor{Resistance: 500},
		Double_layer_capacity: dipoles.Capacitor{Capacitance: 1e-6},
		Diffusion_impedance:   dipoles.WarburgTrasmissive{Sigma: 300, Delta: 1e-6, DiffusionCoeff: 1e-6},
	}
	circuits.Nyquist_plot(circuit1, "OutputFiles/Nyq.csv", 0, 6, 100)
	circuits.Bode(circuit1, "OutputFiles/Bode.csv", -3, 6, 100)

	circuit2 := circuits.RandlesRB{
		Solution_rasistance:   dipoles.Resistor{Resistance: 100},
		Reaction_resistance:   dipoles.Resistor{Resistance: 500},
		Double_layer_capacity: dipoles.Capacitor{Capacitance: 1e-6},
		Diffusion_impedance:   dipoles.WarburgReflective{Sigma: 300, Delta: 1e-6, DiffusionCoeff: 1e-6},
	}
	circuits.Nyquist_plot(circuit2, "OutputFiles/Nyq.csv", 0, 6, 100)
	circuits.Bode(circuit2, "OutputFiles/Bode.csv", -3, 6, 100)

	circuit3 := circuits.Randles{
		Solution_rasistance:   dipoles.Resistor{Resistance: 100},
		Reaction_resistance:   dipoles.Resistor{Resistance: 500},
		Double_layer_capacity: dipoles.Capacitor{Capacitance: 1e-6},
		Diffusion_impedance:   dipoles.Warburg{W_sigma: 300},
	}
	circuits.Nyquist_plot(circuit3, "OutputFiles/Nyq.csv", 0, 6, 100)
	circuits.Bode(circuit3, "OutputFiles/Bode.csv", -3, 6, 100)

}
