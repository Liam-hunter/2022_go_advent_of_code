package calories

type Max struct {
	calories int
}

func NewMax() *Max {
	return &Max{calories: 0}
}

func (m *Max)SetMax(i int) {
	if m.calories < i {
		m.calories = i
	}
}

func (m *Max)GetMax() int {
	return m.calories
}
