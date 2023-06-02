package minimumspanningtree

type UnionSet struct {
	set []int
}

func NewUnionSet(vertexNum int) *UnionSet {
	return &UnionSet{
		set: make([]int, vertexNum),
	}
}

func (u *UnionSet) Find(x int) int {
	if u.set[x] == 0 {
		return x
	}
	return u.Find(u.set[x])
}

func (u *UnionSet) Union(x, y int) {
	xRoot := u.Find(x)
	yRoot := u.Find(y)
	if xRoot != yRoot {
		u.set[xRoot] = yRoot
	}
}
