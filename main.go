package main

import (
	circuits "goEl/curcuits"
	dipoles "goEl/lib"
)

func main() {

	circuit1 := circuits.ButterworthVanDyke{
		Dissipation_resistance: dipoles.Resistor{Resistance: 10},
		Crystal_inertia:        dipoles.Inductor{Inductance: 1e-2},
		Crystal_Capacity:       dipoles.Capacitor{Capacitance: 1e-5},
		Parassitic_capacitance: dipoles.Capacitor{Capacitance: 1e-5},
	}

	//circuits.Nyquist_plot(circuit1, "OutputFiles/Nyq.csv", 0, 6, 100)
	circuits.Bode(circuit1, "OutputFiles/Bode.csv", -3, 6, 100)
	circuits.Nyquist_plot(circuit1, "OutputFiles/Nyq.csv", -0.3, 6, 1000, false)

}
