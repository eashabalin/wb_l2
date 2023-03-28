package pattern

type Motorbike struct {
	Make         string
	Model        string
	Color        string
	EngineVolume int
	WheelSize    int
}

func NewMotorbike(make, model, color string, engineVolume, wheelSize int) *Motorbike {
	return &Motorbike{make, model, color, engineVolume, wheelSize}
}

type MotorbikeBuilder interface {
	GetResult() Motorbike
	SetMake() MotorbikeBuilder
	SetModel() MotorbikeBuilder
	SetColor() MotorbikeBuilder
	SetEngineVolume() MotorbikeBuilder
	SetWheelSize() MotorbikeBuilder
}

type MinskBuilder struct {
	m Motorbike
}

func (mb *MinskBuilder) SetMake() MotorbikeBuilder {
	mb.m.Make = "Minsk"
	return mb
}

func (mb *MinskBuilder) SetModel() MotorbikeBuilder {
	mb.m.Model = "125"
	return mb
}

func (mb *MinskBuilder) SetColor() MotorbikeBuilder {
	mb.m.Color = "orange"
	return mb
}

func (mb *MinskBuilder) SetEngineVolume() MotorbikeBuilder {
	mb.m.EngineVolume = 125
	return mb
}

func (mb *MinskBuilder) SetWheelSize() MotorbikeBuilder {
	mb.m.WheelSize = 19
	return mb
}

func (mb *MinskBuilder) GetResult() Motorbike {
	return mb.m
}

type MotorbikeBuildDirector struct {
	builder MotorbikeBuilder
}

func NewMotorbikeBuildDirector(builder MotorbikeBuilder) *MotorbikeBuildDirector {
	return &MotorbikeBuildDirector{builder: builder}
}

func (d *MotorbikeBuildDirector) Construct() {
	d.builder.SetMake().SetModel().SetEngineVolume().SetWheelSize().SetColor()
}

func (d *MotorbikeBuildDirector) GetResult() Motorbike {
	return d.builder.GetResult()
}
