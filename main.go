package main

import (
	circuits "goEl/curcuits"
	dipoles "goEl/lib"
)

func main() {

	circuit1 := circuits.RandlesRB{
		Solution_rasistance:   dipoles.Resistor{Resistance: 10},
		Reaction_resistance:   dipoles.Resistor{Resistance: 50},
		Double_layer_capacity: dipoles.Capacitor{Capacitance: 20e-6},
		Diffusion_impedance: dipoles.WarburgReflective{
			Sigma:          30,
			Delta:          0.0001,
			DiffusionCoeff: 1.0e-5,
		},
	}
	//circuits.Nyquist_plot(circuit1, "OutputFiles/Nyq.csv", 0, 6, 100)
	circuits.Bode(circuit1, "OutputFiles/Bode.csv", -3, 6, 100)
	circuits.Nyquist_plot(circuit1, "OutputFiles/Nyq.csv", -0.7, 6, 1000)

	circuit2 := circuits.Randles{
		Solution_rasistance:   dipoles.Resistor{Resistance: 10},
		Reaction_resistance:   dipoles.Resistor{Resistance: 50},
		Double_layer_capacity: dipoles.Capacitor{Capacitance: 20e-6},
		Diffusion_impedance: dipoles.Warburg{
			W_sigma: 30,
		},
	}
	//circuits.Nyquist_plot(circuit1, "OutputFiles/Nyq.csv", 0, 6, 100)
	circuits.Bode(circuit2, "OutputFiles/Bode.csv", -3, 6, 100)
	circuits.Nyquist_plot(circuit2, "OutputFiles/Nyq.csv", -0.7, 6, 1000)

}
