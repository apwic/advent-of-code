package main

type Button struct {
	pad  string
	idx  int
	path []string
}

type PriorityQueue []*Button

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].idx == pq[j].idx {
		return len(pq[i].path) < len(pq[j].path)
	}

	return pq[i].idx > pq[j].idx
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Button))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
