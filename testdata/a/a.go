package a

type IA interface  {
	Add()
	Add2(i int) int
	Add3(i int, j,k int) (int,int)
	Add4(i int) (int)
}

type SA struct {
}

func (this * SA) Add(){}

func (this * SA) Add2(int)(int){
	return 0
}

func (this * SA) Add3(i int, j int, k int)(int, int){
	return 0,0
}

func (this * SA) Add4(i int)int{
	return 0
}

