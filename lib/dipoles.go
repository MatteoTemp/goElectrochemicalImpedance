package dipoles

import (
	"math"
	"math/cmplx"
)

func Nyquist(Z complex128) (float64, float64) {
	Re := real(Z)
	Im := imag(Z)
	return Re, Im
}

type Dipole interface {
	Impedance(float64) complex128
}

func Series(A Dipole, B Dipole, freq float64) complex128 {
	return A.Impedance(freq) + B.Impedance(freq)
}

func Parallel(A Dipole, B Dipole, freq float64) complex128 {
	return 1 / (1/A.Impedance(freq) + 1/(B.Impedance(freq)))
}

func VoltageDivider(A Dipole, B Dipole, freq float64) complex128 {
	return B.Impedance(freq) / Series(A, B, freq)
}

type Generic_impedace struct {
	Z complex128
}

type Resistor struct {
	Resistance float64
}

func (r Resistor) Impedance(freq float64) complex128 {
	return complex(r.Resistance, 0.0)
}

type Capacitor struct {
	Capacitance float64
}

func (r Capacitor) Impedance(freq float64) complex128 {
	return complex(0.0, -1/(2*math.Pi*freq*r.Capacitance))
}

type Inductor struct {
	Inductance float64
}

func (r Inductor) Impedance(freq float64) complex128 {
	return complex(0.0, (2 * math.Pi * freq * r.Inductance))
}

type Warburg struct {
	W_sigma float64
}

func (r Warburg) Impedance(freq float64) complex128 {
	return complex(r.W_sigma/math.Sqrt(2*math.Pi*freq), -r.W_sigma/math.Sqrt(2*math.Pi*freq))
}

type WarburgTrasmissive struct {
	Sigma          float64
	Delta          float64
	DiffusionCoeff float64
}

func (r WarburgTrasmissive) Impedance(freq float64) complex128 {

	tan_factor := cmplx.Tanh(cmplx.Sqrt(complex(0.0, freq*r.Delta/r.DiffusionCoeff)))

	return complex(r.Sigma/math.Sqrt(freq), 0.0) * tan_factor * complex(1.0, -1.0)
}

type WarburgReflective struct {
	Sigma          float64
	Delta          float64
	DiffusionCoeff float64
}

func (r WarburgReflective) Impedance(freq float64) complex128 {

	tan_factor := 1 / cmplx.Tanh(cmplx.Sqrt(complex(0.0, freq*r.Delta/r.DiffusionCoeff)))

	return complex(r.Sigma/math.Sqrt(freq), 0.0) * tan_factor * complex(1.0, -1.0)
}

type Constant_phase_Element struct {
	TParameter float64
	Phi        float64
}

func (r Constant_phase_Element) Impedance(freq float64) complex128 {
	return complex(
		math.Cos(0.5*math.Pi*r.Phi)/(r.TParameter*math.Pow(freq, r.Phi)),
		-math.Sin(0.5*math.Pi*r.Phi)/(r.TParameter*math.Pow(freq, r.Phi)),
	)

}
