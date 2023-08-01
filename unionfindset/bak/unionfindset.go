package bak

type UnionFindSet struct {
	arr []int
}

func NewUnionFindSet(len int) *UnionFindSet {
	set := &UnionFindSet{arr: make([]int, len)}
	set.Init()
	return set
}
func (this *UnionFindSet) Init() {
	for i := 0; i < len(this.arr); i++ {
		this.arr[i] = i
	}
}
func (this *UnionFindSet) Find(a int) int {
	if this.arr[a] == a {
		return a
	}
	this.arr[a] = this.Find(this.arr[a])
	return this.arr[a]
}
func (this *UnionFindSet) Merge(a, b int) {
	a1 := this.Find(a)
	b1 := this.Find(b)
	this.arr[b1] = a1
	//this.arr[a1] = b1
}
func (this *UnionFindSet) Sum() int {
	var sum int
	for i, _ := range this.arr {
		if i == this.Find(i) {
			sum++
		}
	}
	return sum
}
