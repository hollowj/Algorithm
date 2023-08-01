package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Expression string

func (this Expression) IsEqual() bool {
	return this[1:3] == "=="
}
func (this Expression) Left() uint8 {
	return this[0]
}
func (this Expression) Right() uint8 {
	return this[3]
}

type UnionFindSet struct {
	set []uint8
}

func NewUnionFindSet() *UnionFindSet {
	set := &UnionFindSet{set: make([]uint8, 26)}
	for i := 0; i < 26; i++ {
		set.set[i] = uint8(i)
	}
	return set
}
func (this *UnionFindSet) GetOffsetByChar(a uint8) uint8 {
	return a - 97
}
func (this *UnionFindSet) Find(a uint8) uint8 {
	if a > 26 {
		a = this.GetOffsetByChar(a)
	}
	b := this.set[a]
	if b == a {
		return b
	} else {
		this.set[a] = this.Find(this.set[a])
		return this.set[a]
	}
}
func (this *UnionFindSet) Merge(a, b uint8) {
	a1 := this.Find(a)
	b1 := this.Find(b)
	this.set[b1] = a1
}
func equationsPossible(equations []string) bool {
	set := NewUnionFindSet()
	var notEqualExpressions []Expression

	for _, equation := range equations {
		expression := Expression(equation)
		if expression.IsEqual() {
			set.Merge(expression.Left(), expression.Right())
		} else {
			notEqualExpressions = append(notEqualExpressions, expression)
		}
	}
	for _, expression := range notEqualExpressions {
		if set.Find(expression.Left()) == set.Find(expression.Right()) {
			return false
		}
	}
	return true
}
func TestQ990(t *testing.T) {
	arr := []string{"a==b", "b!=a"}
	assert.False(t, equationsPossible(arr))
	arr = []string{"a==b", "b!=c", "c==a"}
	assert.False(t, equationsPossible(arr))
	arr = []string{"c==c", "b==d", "x!=z"}
	assert.True(t, equationsPossible(arr))
}
