package main

import (
	circuits "goEl/curcuits"
	dipoles "goEl/lib"
	//cmath "math/cmplx"
)

func main() {

	batteria := circuits.LiIon{
		WireR:     dipoles.Resistor{Resistance: 10},
		Csei:      dipoles.Capacitor{Capacitance: 10e-6},
		Rsei:      dipoles.Resistor{Resistance: 100},
		Cct:       dipoles.Capacitor{Capacitance: 1e-6},
		Rct:       dipoles.Resistor{Resistance: 100},
		Diffusion: dipoles.Warburg{W_sigma: 1000},
	}

	circuits.Nyquist_plot(batteria, "OutputFiles/Nyq.csv", -0, 6, 100)
	circuits.Bode(batteria, "OutputFiles/Bode.csv", -3, 12, 50)

}
