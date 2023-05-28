package temperature

type Thermal struct {
	Temperature *Temperature
	options     ThermalOptions
}

type ThermalOptions struct {
	RelativeAirSpeed     *float64
	RelativeHumidity     *float64
	MetabolicRate        *float64
	NetRadiationAbsorbed *float64
}

func NewThermal(t *Temperature) *Thermal {
	return &Thermal{
		Temperature: t,
		options:     ThermalOptions{},
	}
}

func (t *Thermal) Humidity(h float64) *Thermal {
	t.options.RelativeHumidity = &h
	return t
}

func (t *Thermal) RelativeHumidity(h float64) *Thermal {
	t.options.RelativeHumidity = &h
	return t
}

func (t *Thermal) RelativeAirSpeed(v float64) *Thermal {
	vr := v
	if t.options.MetabolicRate != nil && *t.options.MetabolicRate > 1 {
		vr = v + 0.3*(*t.options.MetabolicRate-1)
	}

	t.options.RelativeAirSpeed = &vr
	return t
}

func (t *Thermal) NetRadiationAbsorbed(r float64) *Thermal {
	t.options.NetRadiationAbsorbed = &r
	return t
}
