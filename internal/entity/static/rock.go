package static

type Rock struct {
	*StaticEntity
}

func NewRock() *Rock {
	return &Rock{
		StaticEntity: NewStaticEntity(),
	}
}

func (r *Rock) String() string {
	//todo: move constant in utility class
	return "rock"
}
