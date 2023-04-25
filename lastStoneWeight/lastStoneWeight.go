package laststoneweight

// Max Heap
type heap []int

func (h heap) up(i int) {
	for {
		parent := (i - 1) / 2
		if i == parent || h[parent] > h[i] {
			break
		}
		h.swap(&h[i], &h[parent])
		i = parent
	}
}

func (h heap) down(i0 int) {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= h.len() || j1 < 0 {
			break
		}
		j := j1
		if j2 := 2*i + 2; j2 < h.len() && h[j2] > h[j1] {
			j = j2
		}
		if h[i] > h[j] {
			break
		}
		h.swap(&h[i], &h[j])
		i = j
	}
}

func (h heap) swap(a, b *int) {
	*a, *b = *b, *a
}

func (h *heap) push(num int) {
	*h = append(*h, num)
	h.up(h.len() - 1)
}

func (h heap) len() int {
	return len(h)
}

func (h *heap) pop() int {
	n := (*h)[0]
	h.swap(&(*h)[0], &(*h)[h.len()-1])
	*h = (*h)[:len(*h)-1]
	if h.len() > 1 {
		h.down(0)
	}
	return n
}

func LastStoneWeight(stones []int) int {
	mh := &heap{}
	for _, v := range stones {
		mh.push(v)
	}

	for mh.len() > 1 {
		x := mh.pop()
		y := mh.pop()
		if x != y {
			mh.push(x - y)
		}
	}

	if mh.len() == 1 {
		return mh.pop()
	}

	return 0
}
