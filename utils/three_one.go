package utils

import "sync"

type ThreeOne struct {
	sync.Once
	v float64
}

func (t *ThreeOne) GetThree() float64 {
	t.Do(func() {
		t.v = 3.0
	})
	return t.v
}
