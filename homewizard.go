//package homewizard implements basic API calls for reading HomeWizard data
package homewizard

import (
	"fmt"
	"net"
	"net/http"
)

// HomeWizard is the main struct for calling all methods
type HomeWizard struct {
	Name       string
	IP         net.IP
	Password   string
	Verbose    bool
	HTTPClient *http.Client
}

func (hw *HomeWizard) GetSensors() (*GetSensorsOutput, error) {
	in := new(GetSensorsInput)
	out := new(GetSensorsOutput)
	err := hw.Do(in, out)
	if err != nil {
		return nil, err
	}
	return out, err
}

func (hw *HomeWizard) GetSwitches() (*GetSwitchesOutput, error) {
	in := new(GetSwitchesInput)
	out := new(GetSwitchesOutput)
	err := hw.Do(in, out)
	if err != nil {
		return nil, err
	}
	return out, err
}

func (hw *HomeWizard) OperateSwitch(in *OperateSwitchInput) (*OperateSwitchOutput, error) {
	out := new(OperateSwitchOutput)
	err := hw.Do(in, out)
	if err != nil {
		return nil, err
	}
	return out, err
}

func (hw *HomeWizard) GetThermometers() (*GetThermometersOutput, error) {
	in := new(GetThermometersInput)
	out := new(GetThermometersOutput)
	err := hw.Do(in, out)
	if err != nil {
		return nil, err
	}
	return out, err
}

func (hw *HomeWizard) GetThermometerGraph(id int, period TimePeriod) (ThermGraph, error) {
	in := new(GetThermometerGraphsInput)
	in.Id = id
	in.Period = period

	var out HWOutput

	switch period {
	case TimePeriodDay, TimePeriodWeek:
		out = new(GetThermometerGraphsDWOutput)
	case TimePeriodMonth, TimePeriodYear:
		out = new(GetThermometerGraphsMYOutput)
	default:
		return nil, fmt.Errorf("period must be day, week, month or year")
	}

	err := hw.Do(in, out)
	if err != nil {
		return nil, err
	}

	// Cast to ThermGraph
	if tg, ok := out.(ThermGraph); ok {
		return tg, err
	}

	return nil, fmt.Errorf("internal datatype error")
}

func (hw *HomeWizard) GetEnergyMeters() (*GetEnergyMetersOutput, error) {
	in := new(GetEnergyMetersInput)
	out := new(GetEnergyMetersOutput)
	err := hw.Do(in, out)
	if err != nil {
		return nil, err
	}
	return out, err
}
