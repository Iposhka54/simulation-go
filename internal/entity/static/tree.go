package static

type Tree struct {
	*StaticEntity
}

func NewTree() *Tree {
	return &Tree{
		StaticEntity: NewStaticEntity(),
	}
}

func (t *Tree) String() string {
	//todo: move constant in utility class
	return "tree"
}
