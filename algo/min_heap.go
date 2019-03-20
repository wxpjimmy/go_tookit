package algo

type Order interface {
	compare(a interface{}, b interface{})  int
}

type MinHeap struct {
	size		int
	occupied	int
	data	[]interface{}
	sortFunc	Order
}

func (heap *MinHeap) addData(item interface{}) {
	if len(heap.data) == 0 {
		heap.data = make([]interface{}, heap.size)
		heap.data[0] = item
		heap.occupied = 1
		return
	}
	if heap.occupied < heap.size {
		heap.data[heap.occupied] = item
		heap.occupied += 1
		heap.adjustHeapUp()
	} else {
		if heap.sortFunc.compare(heap.data[0], item) < 0 {
			heap.data[0] = item
			heap.adjustHeapDown()
		}
	}
}

func (heap *MinHeap) adjustHeapUp() {
	idx := heap.occupied - 1
	if idx == 0 {
		return
	}
	for idx > 0 {
		var parent= -1
		if idx%2 == 0 {
			parent = (idx - 2) / 2
		} else {
			parent = idx / 2
		}

		if heap.sortFunc.compare(heap.data[idx], heap.data[parent]) < 0 {
			heap.swap(idx, parent)
			idx = parent
		} else {
			break
		}
	}

	heap.adjustHeapDown()
}

func (heap *MinHeap) swap(from, to int) {
	tmp := heap.data[from]
	heap.data[from] = heap.data[to]
	heap.data[to] = tmp
}

func (heap *MinHeap) adjustHeapDown() {
	idx := 0
	for idx < heap.occupied -1 {
		leftChild := 2*idx + 1
		rightChild := 2*idx + 2
		if heap.occupied-1 <= leftChild {
			return
		}

		minIdx := leftChild
		if heap.occupied-1 >= rightChild {
			if heap.sortFunc.compare(heap.data[leftChild], heap.data[rightChild]) > 0 {
				minIdx = rightChild
			}
		}
		if heap.sortFunc.compare(heap.data[minIdx], heap.data[idx]) < 0 {
			heap.swap(minIdx, idx)
			idx = minIdx
		} else {
			break
		}
	}
}
