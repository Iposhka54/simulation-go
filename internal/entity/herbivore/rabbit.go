package herbivore

type Rabbit struct {
	*Herbivore
}

func NewRabbit(hp, maxHp, speed int) *Rabbit {
	return &Rabbit{
		Herbivore: New(hp, maxHp, speed),
	}
}
