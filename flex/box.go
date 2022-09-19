package flex

type Box interface {
	SetSize(int, int)
	View() string
}

type boxConfig struct {
	ratio int
}

func NewBoxStyle(ratio int) boxConfig {
	return boxConfig{
		ratio: ratio,
	}
}
