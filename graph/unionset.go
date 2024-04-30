package graph

type UnionSet []int

func NewUnionSet(n int) *UnionSet {
	unionSet := make(UnionSet, n)
	for i := 0; i < n; i++ {
		(unionSet)[i] = i
	}
	return &unionSet
}
func (s *UnionSet) Find(from int) int {
	to := (*s)[from]
	if from == to {
		return from
	}
	(*s)[from] = s.Find(to)
	return (*s)[from]
}
func (s *UnionSet) Join(from, to int) {
	u := s.Find(from)
	v := s.Find(to)
	if u == v {
		return
	}
	(*s)[u] = v
}
func (s *UnionSet) IsSame(from, to int) bool {
	return s.Find(from) == s.Find(to)

}
