package easy

type OrderedStream struct {
	ptr  int
	strs []string
}

func Constructor_(n int) OrderedStream {
	return OrderedStream{ptr: 1, strs: make([]string, n+1)}
}

func (o *OrderedStream) Insert(idKey int, value string) []string {
	index := o.ptr
	o.strs[idKey] = value
	for o.ptr < len(o.strs) && o.strs[o.ptr] != "" {
		o.ptr++
	}
	return o.strs[index:o.ptr]
}
