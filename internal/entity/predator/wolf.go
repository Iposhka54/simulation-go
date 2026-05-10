package predator

type Wolf struct {
	*Predator
}

func NewWolf(hp, maxHp, speed, damage int) *Wolf {
	return &Wolf{
		Predator: New(hp, maxHp, speed, damage),
	}
}
