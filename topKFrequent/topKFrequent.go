package topkfrequent

//Max maxheap
type ky struct {
	K int
	V int
}
type maxheap []ky

func (h maxheap) up(i int) {
	for {
		parent := (i - 1) / 2
		if parent == i || h[parent].V > h[i].V {
			break
		}
		h.swap(&h[parent], &h[i])
		i = parent
	}
}
func (h maxheap) down(i0 int) {
	i := i0
	for {
		j1 := i*2 + 1
		if j1 < 0 || j1 >= h.len() {
			break
		}
		j := j1
		if j2 := i*2 + 2; j2 < h.len() && h[j1].V < h[j2].V {
			j = j2
		}
		if h[i].V > h[j].V {
			break
		}
		h.swap(&h[i], &h[j])
		i = j
	}
}
func (h *maxheap) push(n ky) {
	*h = append(*h, n)
	h.up(h.len() - 1)
}
func (h *maxheap) pop() int {
	n := (*h)[0].K
	h.swap(&(*h)[0], &(*h)[h.len()-1])
	*h = (*h)[:h.len()-1]
	if h.len() > 1 {
		h.down(0)
	}
	return n
}

func (h maxheap) swap(a, b *ky) {
	*a, *b = *b, *a
}

func (h maxheap) len() int {
	return len(h)
}

func TopKFrequent(nums []int, k int) []int {
	var result []int
	m := make(map[int]int, 0)
	h := &maxheap{}

	for _, v := range nums {
		if _, found := m[v]; found {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	for k, v := range m {
		ky := &ky{
			K: k,
			V: v,
		}
		h.push(*ky)
	}

	for i := 0; i < k; i++ {
		result = append(result, h.pop())
	}

	return result
}
