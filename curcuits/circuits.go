package circuits

import (
	"encoding/csv"
	"errors"
	"fmt"
	dipoles "goEl/lib"
	"math"
	"math/cmplx"
	"os"
	"os/exec"
	"time"
)

type Circuit interface {
	FreqResponse(float64) complex128
}

func Nyquist_plot(circuit Circuit, filename string, min_logF float64, max_logF float64, ppdec int) error {
	if max_logF < 0 {
		return errors.New("Upper bound cannot be nagative")
	}
	var decade_interval float64 = 1 / float64(ppdec)
	fmt.Println(ppdec, decade_interval)

	file, err := os.Create(filename)
	if err != nil {
		return errors.New("Cannot create output file ")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	Computation_start := time.Now()

	var logf float64
	for logf = min_logF; logf < max_logF; logf += decade_interval {
		freq := math.Pow(10, logf)
		r, i := dipoles.Nyquist(circuit.FreqResponse(freq))

		writer.Write([]string{fmt.Sprintf("%.9f", math.Pow(10, logf)), fmt.Sprintf("%.9f", r), fmt.Sprintf("%.9f", -i)})
	}

	fmt.Println("Nyquist plot computation time: ", time.Since(Computation_start))

	command := exec.Command("gnuplot", "-p", "Recipes/recipeBodeNyq.gp")
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
	fmt.Println(ppdec, decade_interval)

	file, err := os.Create(filename)
	if err != nil {
		return errors.New("Cannot create output file ")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	Computation_start := time.Now()

	var logf float64
	for logf = min_logF; logf <= max_logF; logf += decade_interval {
		freq := math.Pow(10, logf)

		H := circuit.FreqResponse(freq)
		mag := math.Log10(cmplx.Abs(H))
		phase := -cmplx.Phase(H) * 180 / math.Pi

		writer.Write([]string{fmt.Sprintf("%.9f", math.Pow(10, logf)), fmt.Sprintf("%.9f", mag), fmt.Sprintf("%.9f", phase)})
	}

	fmt.Println("BodePlot plot computation time: ", time.Since(Computation_start))

	command := exec.Command("gnuplot", "-p", "Recipes/boderecipe.gp")
	err = command.Run()

	if err != nil {
		fmt.Println("Plotting Error, ", err)
		return err
	}
	return nil
}

func Lasajous(circuit Circuit, freq float64) error {
	file, err := os.Create("OutputFiles/Lasajous.csv")
	if err != nil {
		return errors.New("Cannot create output file ")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	Z := circuit.FreqResponse(freq)
	//amplitude := cmplx.Abs(Z)
	phase := cmplx.Phase(Z)

	var t float64
	for t = 0.0; t < 1.0; t += 0.001 {
		//V=ZI
		V := math.Sin(2 * math.Pi * t)
		I := math.Sin(2*math.Pi*t + phase)
		writer.Write(
			[]string{fmt.Sprintf("%.9e", V), fmt.Sprintf("%.9e", I)},
		)
	}

	return nil
}

type IdealNPElectrode struct {
	SolutionResistance dipoles.Resistor
	InterphaseCapacity dipoles.Capacitor
}

func (parts IdealNPElectrode) FreqResponse(freq float64) complex128 {
	return dipoles.Series(parts.InterphaseCapacity, parts.SolutionResistance, freq)
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

type RandlesTB struct {
	Solution_rasistance   dipoles.Resistor
	Reaction_resistance   dipoles.Resistor
	Double_layer_capacity dipoles.Capacitor
	Diffusion_impedance   dipoles.WarburgTrasmissive
}

func (parts RandlesTB) FreqResponse(freq float64) complex128 {
	z1 := parts.Reaction_resistance.Impedance(freq) + parts.Diffusion_impedance.Impedance(freq)
	z2 := 1 / (1/z1 + 1/parts.Double_layer_capacity.Impedance(freq))

	return z2 + parts.Solution_rasistance.Impedance(freq)
}

type RandlesRB struct {
	Solution_rasistance   dipoles.Resistor
	Reaction_resistance   dipoles.Resistor
	Double_layer_capacity dipoles.Capacitor
	Diffusion_impedance   dipoles.WarburgReflective
}

func (parts RandlesRB) FreqResponse(freq float64) complex128 {
	z1 := parts.Reaction_resistance.Impedance(freq) + parts.Diffusion_impedance.Impedance(freq)
	z2 := 1 / (1/z1 + 1/parts.Double_layer_capacity.Impedance(freq))

	return z2 + parts.Solution_rasistance.Impedance(freq)
}
