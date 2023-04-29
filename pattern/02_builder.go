package pattern

// Порождающий паттерн Строитель предлагает вынести конструирование объекта за пределы его собственного класса,
// поручив это дело отдельным объектам, называемым строителями.

type House struct {
}

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

// MinskBuilder – конкретный строитель мотоцикла Минск
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

// Вы можете пойти дальше и выделить вызовы методов строителя в отдельный класс, называемый директором.
// В этом случае директор будет задавать порядок шагов строительства, а строитель — выполнять их.

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

// UralBuilder – конкретный строитель мотоцикла Минск
type UralBuilder struct {
	m Motorbike
}

func (mb *UralBuilder) SetMake() MotorbikeBuilder {
	mb.m.Make = "Ural"
	return mb
}

func (mb *UralBuilder) SetModel() MotorbikeBuilder {
	mb.m.Model = "650"
	return mb
}

func (mb *UralBuilder) SetColor() MotorbikeBuilder {
	mb.m.Color = "blue"
	return mb
}

func (mb *UralBuilder) SetEngineVolume() MotorbikeBuilder {
	mb.m.EngineVolume = 650
	return mb
}

func (mb *UralBuilder) SetWheelSize() MotorbikeBuilder {
	mb.m.WheelSize = 19
	return mb
}

func (mb *UralBuilder) GetResult() Motorbike {
	return mb.m
}
