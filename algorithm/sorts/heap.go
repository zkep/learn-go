package sorts

type MaxHeap struct {
	slice    []Comparable
	headSize int
	indices  map[int]int
}

func (h *MaxHeap) Init(slice []Comparable) {
	if slice == nil {
		slice = make([]Comparable, 0)
	}
	h.slice = slice
	h.headSize = len(slice)
	h.indices = make(map[int]int)
	h.Heapify()
}

func (h *MaxHeap) Pop() Comparable {
	if h.headSize == 0 {
		return nil
	}
	i := h.slice[0]
	h.headSize--
	h.slice[0] = h.slice[h.headSize]
	h.updateidx(0)
	h.heapifyDown(0)
	h.slice = h.slice[0:h.headSize]
	return i
}

func (h *MaxHeap) Push(i Comparable) {
	h.slice = append(h.slice, i)
	h.updateidx(h.headSize)
	h.heapifyUp(h.headSize)
	h.headSize++
}

func (h MaxHeap) Size() int {
	return h.headSize
}

func (h MaxHeap) Heapify() {
	for i, v := range h.slice {
		h.indices[v.Idx()] = i
	}
	for i := h.headSize / 2; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func (h *MaxHeap) heapifyDown(i int) {
	l, r := 2*i+1, 2*i+2
	max := i
	if l < h.headSize && h.slice[l].More(h.slice[max]) {
		max = l
	}
	if r < h.headSize && h.slice[r].More(h.slice[max]) {
		max = r
	}
	if max != i {
		h.slice[i], h.slice[max] = h.slice[max], h.slice[i]
		h.updateidx(i)
		h.updateidx(max)
		h.heapifyDown(max)
	}
}

func (h *MaxHeap) heapifyUp(i int) {
	if i == 0 {
		return
	}
	p := i / 2
	if h.slice[i].More(h.slice[p]) {
		h.slice[i], h.slice[p] = h.slice[p], h.slice[i]
		h.updateidx(i)
		h.updateidx(p)
		h.heapifyUp(p)
	}
}

func (h MaxHeap) updateidx(i int) {
	h.indices[h.slice[i].Idx()] = i
}

type Comparable interface {
	Idx() int
	More(interface{}) bool
}

type Int int

func (a Int) More(b interface{}) bool {
	return a > b.(Int)
}

func (a Int) Idx() int {
	return int(a)
}

func buildMaxHeap(array []int) MaxHeap {
	var slice []Comparable
	for _, i := range array {
		slice = append(slice, Int(i))
	}
	h := MaxHeap{}
	h.Init(slice)
	return h
}

func HeapSort(slice []int) []int {
	h := buildMaxHeap(slice)
	for i := len(h.slice) - 1; i >= 1; i-- {
		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
		h.headSize--
		h.heapifyDown(0)
	}
	res := []int{}
	for _, i := range h.slice {
		res = append(res, int(i.(Int)))
	}
	return res
}
