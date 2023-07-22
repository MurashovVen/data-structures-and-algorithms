package collections

type (
	HeapMax struct {
		tree []int
	}

	HeapMaxOption func(*HeapMax)
)

func NewHeapMax(opts ...HeapMaxOption) *HeapMax {
	h := &HeapMax{}
	for _, opt := range opts {
		opt(h)
	}

	return h
}

func HeapMaxFromSlice(sl []int) HeapMaxOption {
	return func(heap *HeapMax) {
		heap.tree = sl
	}
}

// Insert -
// worst - O (log n)
// average - 0 (1)
func (t *HeapMax) Insert(value int) {
	t.tree = append(t.tree, value)

	nodeNewIdx := len(t.tree) - 1

	if len(t.tree) == 1 {
		return
	}

	currNodeIdx := nodeNewIdx
loop:
	for {
		if currNodeIdx == 0 {
			break loop
		}
		parentIdx := (currNodeIdx - 1) / 2

		switch {
		case t.tree[parentIdx] >= t.tree[currNodeIdx]:
			break loop

		case t.tree[parentIdx] < t.tree[currNodeIdx]:
			t.tree[parentIdx], t.tree[currNodeIdx] = t.tree[currNodeIdx], t.tree[parentIdx]
			currNodeIdx = parentIdx
		}
	}
}

// Pop -
// O (log n)
func (t *HeapMax) Pop() (int, error) {
	if len(t.tree) == 0 {
		return 0, ErrEmpty
	}

	previousMax := t.tree[0]

	t.tree[0], t.tree[len(t.tree)-1] = t.tree[len(t.tree)-1], t.tree[0]

	t.tree = t.tree[0 : len(t.tree)-1]

	var (
		currIdx = 0
	)
loop:
	for {
		var (
			left  = currIdx*2 + 1
			right = currIdx*2 + 2
		)
		if left >= len(t.tree) {
			break loop
		}

		maxIdx := left
		if right < len(t.tree) && t.tree[left] < t.tree[right] {
			maxIdx = right
		}

		if t.tree[currIdx] < t.tree[maxIdx] {
			t.tree[currIdx], t.tree[maxIdx] = t.tree[maxIdx], t.tree[currIdx]

			currIdx = maxIdx
		} else {
			break
		}
	}

	return previousMax, nil
}

// Max -
// O (1)
func (t *HeapMax) Max() (int, error) {
	if len(t.tree) == 0 {
		return 0, ErrEmpty
	}

	return t.tree[0], nil
}
