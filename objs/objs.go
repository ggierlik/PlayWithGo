package objs

type Rect struct {
	W, H int
}

//metoda rect
func (r *Rect) Area() int {
	return r.W * r.H
}
