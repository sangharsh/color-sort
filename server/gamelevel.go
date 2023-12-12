package main

type GameLevel struct {
	level int
	tubes []Testtube
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
