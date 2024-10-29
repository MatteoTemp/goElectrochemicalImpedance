package circuits

import (
	"encoding/csv"
	"errors"
	"fmt"
	dipoles "goEl/lib"
	"math"
	"os"
	"os/exec"
)

type Circuit interface {
	FreqResponse(float64) complex128
}

func Nyquist_plot(circuit Circuit, filename string, min_logF float64, max_logF float64, ppdec int) error {
	if max_logF < 0 {
		return errors.New("Upper bound cannot be nagative")
	}
	var decade_interval float64 = 1 / float64(ppdec)

	file, err := os.Create(filename)
	if err != nil {
		return errors.New("Cannot create output file ")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var logf float64
	for logf = min_logF; logf <= max_logF; logf += decade_interval {
		freq := math.Pow(10, logf)
		r, i := dipoles.Nyquist(circuit.FreqResponse(freq))

		writer.Write([]string{fmt.Sprintf("%.5f", math.Pow(10, logf)), fmt.Sprintf("%.5f", r), fmt.Sprintf("%.5f", i)})
	}

	command := exec.Command("gnuplot", "-p", "recipeBodeNyq.gp")
	err = command.Run()

	if err != nil {
		fmt.Println("Plotting Error, ", err)
		return err
	}

	return nil
}

func Bode(circuit Circuit, filename string, min_logF float64, max_logF float64, ppdec int) error {

	if max_logF < 0 {
		return errors.New("Upper bound cannot be nagative")
	}
	var decade_interval float64 = 1 / float64(ppdec)

	file, err := os.Create(filename)
	if err != nil {
		return errors.New("Cannot create output file ")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var logf float64
	for logf = min_logF; logf <= max_logF; logf += decade_interval {
		freq := math.Pow(10, logf)
		r, i := dipoles.Nyquist(circuit.FreqResponse(freq))
		mag := math.Log10(math.Hypot(r, i))
		phase := -math.Atan(i/r) * 180 / math.Pi

		writer.Write([]string{fmt.Sprintf("%.5f", math.Pow(10, logf)), fmt.Sprintf("%.5f", mag), fmt.Sprintf("%.5f", phase)})
	}

	command := exec.Command("gnuplot", "-p", "boderecipe.gp")
	err = command.Run()

	if err != nil {
		fmt.Println("Plotting Error, ", err)
		return err
	}

	return nil
}

type Randles struct {
	Solution_rasistance   dipoles.Resistor
	Reaction_resistance   dipoles.Resistor
	Double_layer_capacity dipoles.Capacitor
	Diffusion_impedance   dipoles.Warburg
}

func (parts Randles) FreqResponse(freq float64) complex128 {
	z1 := parts.Reaction_resistance.Impedance(freq) + parts.Diffusion_impedance.Impedance(freq)
	z2 := 1 / (1/z1 + 1/parts.Double_layer_capacity.Impedance(freq))

	return z2 + parts.Solution_rasistance.Impedance(freq)
}

type RCLowpass struct {
	Resistor  dipoles.Resistor
	Capacitor dipoles.Capacitor
}

func (parts RCLowpass) FreqResponse(freq float64) complex128 {

	return parts.Capacitor.Impedance(freq) / (parts.Resistor.Impedance(freq) + parts.Capacitor.Impedance(freq))

}

type RCHighpass struct {
	Resistor  dipoles.Resistor
	Capacitor dipoles.Capacitor
}

func (parts RCHighpass) FreqResponse(freq float64) complex128 {
	z_tot := (parts.Resistor.Impedance(freq) + parts.Capacitor.Impedance(freq))
	return parts.Resistor.Impedance(freq) / z_tot
}

type RCBandPass struct {
	R1 dipoles.Resistor
	C1 dipoles.Capacitor

	R2 dipoles.Resistor
	C2 dipoles.Capacitor
}

func (parts RCBandPass) FreqResponse(freq float64) complex128 {

	HPResponse := parts.C1.Impedance(freq) / (parts.R1.Impedance(freq) + parts.C1.Impedance(freq))
	LPResponse := parts.R2.Impedance(freq) / (parts.R2.Impedance(freq) + parts.C2.Impedance(freq))

	return HPResponse * LPResponse
}
