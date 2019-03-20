package algo

import (
	"fmt"
	"testing"
)

type Data struct {
	value int
}

type Comparator struct {

}

func (r *Comparator) compare(c interface{}, d interface{}) int {
	field, ok := c.(*Data)
	if !ok {
		panic("parameter not Data instance")
	}
	field2, ok2 := d.(*Data)
	if !ok2 {
		panic("parameter not Data instance")
	}
	if field.value < field2.value {
		return -1
	} else if field.value == field2.value {
		return 0
	} else {
		return 1
	}
}

func TestMinHeap(t *testing.T) {
	var heap = MinHeap{10, 0, nil, &Comparator{}}

	heap.addData(&Data{20})
	heap.addData(&Data{10})
	heap.addData(&Data{15})
	heap.addData(&Data{8})
	heap.addData(&Data{19})
	heap.addData(&Data{4})
	heap.addData(&Data{12})
	heap.addData(&Data{10})
	heap.addData(&Data{16})
	heap.addData(&Data{7})

	fmt.Println(heap.data[0].(*Data).value)
	for i:=0;i<heap.occupied;i++ {
		fmt.Println(heap.data[i].(*Data).value)
	}
	heap.addData(&Data{17})

	fmt.Println()
	for i:=0;i<heap.occupied;i++ {
		fmt.Println(heap.data[i].(*Data).value)
	}

	heap.addData(&Data{13})

	fmt.Println()
	for i:=0;i<heap.occupied;i++ {
		fmt.Println(heap.data[i].(*Data).value)
	}

}
