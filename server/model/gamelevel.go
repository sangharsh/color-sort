package model

func NewGameLevel(level int32, tubes []*Testtube) *GameLevel {
	tubes2 := append(tubes, &Testtube{Size: 4, Colors: []Color{}}, &Testtube{Size: 4, Colors: []Color{}})
	gameLevel := &GameLevel{
		Level: level,
		Tubes: tubes2,
	}
	return gameLevel
}

func (level *GameLevel) Pour(srcidx, dstidx int) (bool, error) {
	src := level.Tubes[srcidx]
	dst := level.Tubes[dstidx]
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

func (level *GameLevel) Won() bool {
	for _, tt := range level.GetTubes() {
		if !tt.IsComplete() {
			return false
		}
	}
	return true
}
