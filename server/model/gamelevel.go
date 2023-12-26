package model

type GameLevel struct {
	level int
	tubes []Testtube
}


func NewGameLevel(level int, tubes []Testtube) (*GameLevel) {
	tubes2 := append(tubes, Testtube{4, []string{}}, Testtube{4, []string{}})
	return &GameLevel{level, tubes2}
}

func (level *GameLevel) Pour(srcidx, dstidx int) (bool, error) {
	src := &level.tubes[srcidx]
	dst := &level.tubes[dstidx]
	color, err := src.Peek()
	if err != nil {
		return false, err
	}
	err = dst.AddColor(color)
	if err != nil {
		return false, err
	}
	_, err = src.Pop()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (level GameLevel) Won() bool {
	for _, tt := range level.tubes {
		if !tt.IsComplete() {
			return false
		}
	}
	return true
}
