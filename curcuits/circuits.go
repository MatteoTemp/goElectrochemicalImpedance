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
		r, i := dipoles.Nyquist(circuit.FreqResponse(freq))

		writer.Write([]string{fmt.Sprintf("%f", math.Pow(10, logf)), fmt.Sprintf("%f", r), fmt.Sprintf("%f", -i)})
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

		writer.Write([]string{fmt.Sprintf("%f", math.Pow(10, logf)), fmt.Sprintf("%f", mag), fmt.Sprintf("%f", phase)})
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

type LiIon struct {
	WireR     dipoles.Resistor
	Rsei      dipoles.Resistor
	Csei      dipoles.Capacitor
	Rct       dipoles.Resistor
	Cct       dipoles.Capacitor
	Diffusion dipoles.Warburg
}

func (parts LiIon) FreqResponse(freq float64) complex128 {
	total_impedance := parts.WireR.Impedance(freq)
	total_impedance += dipoles.Parallel(parts.Csei, parts.Rsei, freq)

	total_impedance += dipoles.Parallel(parts.Cct, parts.Rct, freq) + parts.Diffusion.Impedance(freq)

	return total_impedance
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

	//Da implementare correttamente
	return 0
}

type IdealLCSeries struct {
	C dipoles.Capacitor
	L dipoles.Inductor
}

func (parts IdealLCSeries) FreqResponse(freq float64) complex128 {

	series_impedance := parts.L.Impedance(freq) + parts.C.Impedance(freq)

	return parts.C.Impedance(freq) / series_impedance

}

type RealLCSeries struct {
	R dipoles.Resistor
	C dipoles.Capacitor
	L dipoles.Inductor
}

func (parts RealLCSeries) FreqResponse(freq float64) complex128 {

	series_impedance := parts.L.Impedance(freq) + parts.C.Impedance(freq) + parts.R.Impedance(freq)

	return parts.C.Impedance(freq) / series_impedance

}

type RCLBandpass struct {
	R dipoles.Resistor
	C dipoles.Capacitor
	L dipoles.Inductor
}

func (parts RCLBandpass) FreqResponse(freq float64) complex128 {

	LC_tank := dipoles.Parallel(parts.L, parts.C, freq)

	return LC_tank / (parts.R.Impedance(freq) + LC_tank)
}
