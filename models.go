package homewizard

import (
	"fmt"
	"time"
)

type TimePeriod string

const (
	TimePeriodDay   TimePeriod = "day"
	TimePeriodWeek  TimePeriod = "week"
	TimePeriodMonth TimePeriod = "month"
	TimePeriodYear  TimePeriod = "year"
)

type SwitchStatus string

func (st SwitchStatus) Revert() SwitchStatus {
	if st == SwitchStatusOn {
		return SwitchStatusOff
	}
	return SwitchStatusOn
}

const (
	SwitchStatusOn  SwitchStatus = "on"
	SwitchStatusOff SwitchStatus = "off"
)

type SwitchType string

const (
	SwitchTypeSwitch   SwitchType = "switch"
	SwitchTypeRadiator SwitchType = "radiator"
	SwitchTypeDimmer   SwitchType = "dimmer"
)

type Handshake struct {
	Homewizard              string `json:"homewizard"`
	Version                 string `json:"version"`
	Firmwareupdateavailable string `json:"firmwareupdateavailable"`
	Appupdaterequired       string `json:"appupdaterequired"`
	Serial                  string `json:"serial"`
}

type Camera struct {
	ID       int      `json:"id"`
	IP       string   `json:"ip"`
	Mode     int      `json:"mode"`
	Model    int      `json:"model"`
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Port     string   `json:"port"`
	Presets  []Preset `json:"presets"`
	URL      struct {
		Auth  string `json:"auth"`
		Path  string `json:"path"`
		Query string `json:"query"`
	} `json:"url"`
	Username string `json:"username"`
}

type Kakusensor struct {
	Cameraid  int    `json:"cameraid"`
	Favorite  string `json:"favorite"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
}

type Switch struct {
	ID       int          `json:"id"`
	Name     string       `json:"name"`
	Type     SwitchType   `json:"type"`
	Status   SwitchStatus `json:"status"`
	Favorite string       `json:"favorite"`
	Tte      float64      `json:"tte"`
}

type Switches []Switch

func (sws Switches) ByName(name string) *Switch {
	for _, sw := range sws {
		if sw.Name == name {
			return &sw
		}
	}
	return nil
}

func (sws Switches) ByID(id int) *Switch {
	for _, sw := range sws {
		if sw.ID == id {
			return &sw
		}
	}
	return nil
}

func (sws Switches) ByType(switchType SwitchType) []Switch {
	out := make([]Switch, 0)
	for _, sw := range sws {
		if sw.Type == switchType {
			out = append(out, sw)
		}
	}
	return out
}

type Thermometer struct {
	Code       string  `json:"code"`
	Favorite   string  `json:"favorite"`
	Hu         int     `json:"hu"`
	HuMax      int     `json:"hu+"`
	HuMaxT     string  `json:"hu+t"`
	HuMin      int     `json:"hu-"`
	HuMinT     string  `json:"hu-t"`
	ID         int     `json:"id"`
	LowBattery string  `json:"lowBattery"`
	Model      int     `json:"model"`
	Name       string  `json:"name"`
	Outside    string  `json:"outside"`
	Te         float64 `json:"te"`
	TeMax      float64 `json:"te+"`
	TeMaxT     string  `json:"te+t"`
	TeMin      float64 `json:"te-"`
	TeMinT     string  `json:"te-t"`
	Version    float64 `json:"version"`
}

type Rainmeter struct {
	Mm float64 `json:"mm"`
	H3 float64 `json:"3h"`
}

type Windmeter struct {
	Ws     float64 `json:"ws"`
	Dir    string  `json:"dir"`
	Gu     float64 `json:"gu"`
	Wc     float64 `json:"wc"`
	WsMax  float64 `json:"ws+"`
	WsMaxT string  `json:"ws+t"`
	WsMin  float64 `json:"ws-"`
	WsMinT string  `json:"ws-t"`
}

type Energymeter struct {
	Code       string  `json:"code"`
	DayTotal   float64 `json:"dayTotal"`
	Favorite   string  `json:"favorite"`
	ID         int     `json:"id"`
	Key        string  `json:"key"`
	LowBattery string  `json:"lowBattery"`
	Name       string  `json:"name"`
	Po         int     `json:"po"`
	PoMax      int     `json:"po+"`
	PoMaxT     string  `json:"po+t"`
	PoMin      int     `json:"po-"`
	PoMinT     string  `json:"po-t"`
}

type Energylink struct{}     //TODO
type Heatlink struct{}       //TODO
type Hue struct{}            //TODO
type Preset struct{}         //TODO
type Scene struct{}          //TODO
type Uvmeter struct{}        //TODO
type Weatherdisplay struct{} //TODO

type Time struct {
	time.Time
}

// Format: time: "2015-07-05 22:18"
func (t *Time) UnmarshalJSON(b []byte) error {
	// TODO Set Location dynamically (get it from HW?)
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}
	t.Time, err = time.ParseInLocation(`"2006-01-02 15:04"`, string(b), loc)
	return err
}

type ResponseElems struct {
	Request struct {
		Route string `json:"route"`
	} `json:"request"`
	Status  string `json:"status"`
	Version string `json:"version"`
}

func (r *ResponseElems) GetVersion() string {
	return r.Version
}

func (r *ResponseElems) GetStatus() string {
	return r.Status
}

type HWInput interface {
	Route() *RouteInfo
}

type HWOutput interface {
	GetStatus() string
	GetVersion() string
}

type RouteInfo struct {
	Method string
	Route  string
}

type GetSensorsInput struct{}

func (i *GetSensorsInput) Route() *RouteInfo {
	return &RouteInfo{"GET", "/get-sensors"}
}

type GetSensorsOutput struct {
	ResponseElems
	Response struct {
		Cameras         []Camera         `json:"cameras"`
		Energylinks     []Energylink     `json:"energylinks"`
		Energymeters    []Energymeter    `json:"energymeters"`
		Heatlinks       []Heatlink       `json:"heatlinks"`
		Hues            []Hue            `json:"hues"`
		Kakusensors     []Kakusensor     `json:"kakusensors"`
		Preset          int              `json:"preset"`
		Rainmeters      []Rainmeter      `json:"rainmeters"`
		Scenes          []Scene          `json:"scenes"`
		Switches        []Switch         `json:"switches"`
		Thermometers    []Thermometer    `json:"thermometers"`
		Time            Time             `json:"time"`
		Uvmeters        []Uvmeter        `json:"uvmeters"`
		Weatherdisplays []Weatherdisplay `json:"weatherdisplays"`
		Windmeters      []Windmeter      `json:"windmeters"`
	} `json:"response"`
}

type GetSwitchesInput struct{}

func (i *GetSwitchesInput) Route() *RouteInfo {
	return &RouteInfo{"GET", "/swlist"}
}

type GetSwitchesOutput struct {
	ResponseElems
	Response Switches `json:"response"`
}

type OperateSwitchInput struct {
	ID            int
	Type          SwitchType
	DesiredStatus SwitchStatus
	DesiredTte    float64
}

func (i *OperateSwitchInput) Route() *RouteInfo {
	switch i.Type {
	case SwitchTypeSwitch:
		return &RouteInfo{"GET", fmt.Sprintf("/sw/%d/%s", i.ID, i.DesiredStatus)}
	case SwitchTypeDimmer:
		return &RouteInfo{"GET", fmt.Sprintf("/sw/dim/%d/%.2f", i.ID, i.DesiredTte)}
	case SwitchTypeRadiator:
		return &RouteInfo{"GET", fmt.Sprintf("/sw/%d/settarget/%.2f", i.ID, i.DesiredTte)}
	}
	return &RouteInfo{"GET", "/swlist"}
}

type OperateSwitchOutput struct {
	ResponseElems
	//Response []Switch `json:"response"`
}

type GetThermometersInput struct{}

func (i *GetThermometersInput) Route() *RouteInfo {
	return &RouteInfo{"GET", "/telist"}
}

type GetThermometersOutput struct {
	ResponseElems
	Response []Thermometer `json:"response"`
}

type GetThermometerGraphsInput struct {
	Id     int
	Period TimePeriod //day-week-month-year
}

func (i *GetThermometerGraphsInput) Route() *RouteInfo {
	return &RouteInfo{"GET", fmt.Sprintf("/te/graph/%d/%s", i.Id, i.Period)}
}

//For day and week graphs
type GetThermometerGraphsDWOutput struct {
	ResponseElems
	Response []ThermGraphT `json:"response"`
}

func (x *GetThermometerGraphsDWOutput) Len() uint {
	return uint(len(x.Response))
}

func (x *GetThermometerGraphsDWOutput) Elem(i uint) ThermGraphPoint {
	return &(x.Response[i])
}

//For month and year graphs
type GetThermometerGraphsMYOutput struct {
	ResponseElems
	Response []ThermGraphMinMax `json:"response"`
}

func (x *GetThermometerGraphsMYOutput) Len() uint {
	return uint(len(x.Response))
}

func (x *GetThermometerGraphsMYOutput) Elem(i uint) ThermGraphPoint {
	return &(x.Response[i])
}

// A ThermGraph is a slice of datapoints (graph data)
// Because HomeWizard returns a different JSON format for specific graphs
// we can use this interface to read all data
type ThermGraph interface {
	Len() uint
	Elem(uint) ThermGraphPoint
}

// A ThermGraphPoint is a single datapoint in a slice of datapoints in graph data
// Because HomeWizard returns a different JSON format for specific graphs
// we can use this interface to read all data
// For Temp() and Time() you pass:
// - 0 : to get the current temp (or temp of that datapoint)
// - >0 : to get the maximum temp
// - <0 : to get the minimum temp
type ThermGraphPoint interface {
	Time() *Time
	Temp(int8) float64
	Hum(int8) int
}

type ThermGraphT struct {
	T  Time    `json:"t"`
	Te float64 `json:"te"`
	Hu int     `json:"hu"`
}

func (g *ThermGraphT) Time() *Time {
	return &(g.T)
}

func (g *ThermGraphT) Temp(p int8) float64 {
	if p == 0 {
		return g.Te
	}
	return 0
}

func (g *ThermGraphT) Hum(p int8) int {
	if p == 0 {
		return g.Hu
	}
	return 0
}

type ThermGraphMinMax struct {
	T     Time    `json:"t"`
	TeMax float64 `json:"te+"`
	TeMin float64 `json:"te-"`
	HuMax int     `json:"hu+"`
	HuMin int     `json:"hu-"`
}

func (g *ThermGraphMinMax) Time() *Time {
	return &(g.T)
}

func (g *ThermGraphMinMax) Temp(p int8) float64 {
	switch {
	case p < 0:
		return g.TeMin
	case p > 0:
		return g.TeMax
	default:
		return 0
	}
}

func (g *ThermGraphMinMax) Hum(p int8) int {
	switch {
	case p < 0:
		return g.HuMin
	case p > 0:
		return g.HuMax
	default:
		return 0
	}
}

type GetEnergyMetersInput struct{}

func (i *GetEnergyMetersInput) Route() *RouteInfo {
	return &RouteInfo{"GET", "/enlist"}
}

type GetEnergyMetersOutput struct {
	ResponseElems
	Response []Energymeter `json:"response"`
}

type HandshakeInput struct{}

func (i *HandshakeInput) Route() *RouteInfo {
	return &RouteInfo{"GET", "/handshake"}
}

type HandshakeOutput struct {
	ResponseElems
	Response Handshake `json:"response"`
}
