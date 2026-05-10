package static

type Grass struct {
	*StaticEntity
}

func NewGrass() *Grass {
	return &Grass{
		StaticEntity: NewStaticEntity(),
	}
}

func (g *Grass) String() string {
	//todo: move constant in utility class
	return "grass"
}
