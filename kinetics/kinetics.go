package kinetics

import "math"

const f = 38.9 //Volt

type ButlerVolmer struct {
	Alpha                 float64 //Fattore di asimmetria
	K0                    float64 //Costante cinetica base
	Equilibrium_potential float64
	Cbulk                 float64 //Concentrazione Bulk

	CatodicLimitCurrent float64
	AnodicLimitCurrent  float64
}

func (m ButlerVolmer) EquilibriumCurrent() float64 {
	return m.K0 * m.Cbulk * math.Exp(-m.Alpha*f*m.Equilibrium_potential)
}

func (m ButlerVolmer) CurrentUnlimited(overpotential float64) (float64, float64) {

	catodic := math.Exp(-m.Alpha * f * overpotential)
	anodic := math.Exp((1.0 - m.Alpha) * f * overpotential)

	net_current := anodic - catodic

	return net_current, net_current * m.EquilibriumCurrent()

}

func (m ButlerVolmer) CurrentLimited(overpotential float64) float64 {

	catodic := math.Exp(-m.Alpha * f * overpotential)
	anodic := math.Exp((1.0 - m.Alpha) * f * overpotential)

	j0 := m.EquilibriumCurrent()

	rcat := j0 / m.CatodicLimitCurrent
	ran := j0 / m.AnodicLimitCurrent

	return (anodic - catodic) / (1 + rcat*catodic + ran*anodic)

}
