package temperature

import (
	"fmt"
	"math"
)

// ApparentTemperature is from the Australian Bureau of Meteorology
// wind velocity is taken into account but for indoor 0.0 fits well
// Closely matches Steadman 1994 and works well for indoor temperatures
// http://www.bom.gov.au/info/thermal_stress/
// If airspeed is not set assumes still air
func (t *Thermal) ApparentTemperature() (*Temperature, error) {
	var relativeAirSpeed float64
	if t.options.RelativeAirSpeed != nil {
		relativeAirSpeed = *t.options.RelativeAirSpeed
	}

	// dividing Pressure of water vapor by 100 since the at equation requires pVap to be in hPa
	pVap := t.PressureVapor() / 100

	var at float64
	if t.options.NetRadiationAbsorbed != nil {
		at = t.Temperature.Celsius + 0.348*pVap - 0.7*+0.7**t.options.NetRadiationAbsorbed/(relativeAirSpeed+10) - 4.25
	} else {
		at = t.Temperature.Celsius + 0.33*pVap - 0.7*relativeAirSpeed - 4.00
	}
	return Celsius(at), nil
}

// PressureVapor partial pressure of water vapor in moist air
// If humidity has not been set assumes a 50% humidity
func (t *Thermal) PressureVapor() float64 {
	if t.options.RelativeHumidity == nil || *t.options.RelativeHumidity == 0.0 {
		*t.options.RelativeHumidity = 50.0
	}
	return *t.options.RelativeHumidity / 100 * t.Temperature.PressureSaturation()
}

// Humidex Calculates the humidex (short for "humidity index"). It has been  developed by the Canadian Meteorological
// service. It was introduced in 1965 and then it was revised by Masterson and Richardson (1979). . It aims  to
// describe how hot, humid weather is felt by the average person. The Humidex differs from the heat index in being
// related to the dew point rather than relative humidity.
func (t *Thermal) Humidex() (*Temperature, error) {
	if t.options.RelativeHumidity == nil {
		return nil, fmt.Errorf("relative humidity needed for humidex")
	}

	hd := t.Temperature.Celsius + float64(5)/float64(9)*(6.112*math.Pow(10, 7.5*float64(t.Temperature.Celsius)/(237.7+float64(t.Temperature.Celsius)))**t.options.RelativeHumidity/float64(100)-10)

	return Celsius(hd), nil
}

// Net calculates the Normal Effective Temperature from Missenard (1993) equation.
func (t *Thermal) Net() (*Temperature, error) {
	if t.options.RelativeHumidity == nil {
		return nil, fmt.Errorf("relative humidity needed for humidex")
	}
	v := 0.0
	frac := 1.0 / (1.76 + 1.4*math.Pow(v, 0.75))
	n := round(37 - (37-t.Temperature.Celsius)/(0.68-0.0014**t.options.RelativeHumidity+frac) - 0.29*t.Temperature.Celsius*(1-0.01**t.options.RelativeHumidity))
	return Celsius(n), nil
}
