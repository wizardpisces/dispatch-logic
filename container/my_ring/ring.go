package my_ring

type LinkedRing struct {
	next, prev *LinkedRing
	Value      any
}

func (r *LinkedRing) Next() *LinkedRing {
	return r.next
}
