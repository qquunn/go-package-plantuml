package a

import "sync"

type IA interface  {
	Add()
}

type SA struct {
	a int
	b sync.Mutex
}

func (this * SA) Add(){}
