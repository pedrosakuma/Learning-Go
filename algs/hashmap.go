package hashmap

var defaultCapacity uint64 = 1 << 10

type node struct {
	key   any
	value any
	next  *node
}

type HashMap struct {
	capacity uint64
	size     uint64
	table    []*node
}

func New() *HashMap {
	return &HashMap{
		capacity: defaultCapacity,
		table:    make([]*node, defaultCapacity),
	}
}

func Make(size, capacity uint64) HashMap {
	return HashMap{
		size:     size,
		capacity: capacity,
		table:    make([]*node, capacity),
	}
}

func (hm *HashMap) Get(key any) any {
	node := hm.getNodeByHash(hm.hash(key))

	if node != nil {
		return node.value
	}

	return nil
}

func (hm *HashMap) Contains(key any) bool {
	return hm.getNodeByHash(hm.hash(key)) != nil
}

func (hm *HashMap) Put(key any, value any) any {
	return hm.putValue(hm.hash(key), key, value)
}

func (hm *HashMap) putValue(hash uint64, key any, value any) any {
	if hm.capacity == 0 {
		hm.capacity = defaultCapacity
		hm.table = make([]*node, defaultCapacity)
	}

	node := hm.getNodeByHash(hash)

	if node == nil {
		hm.table[hash] = newNode(key, value)
	} else if node.key == key {
		hm.table[hash] = newNodeWithNext(key, value, node)
		return value
	} else {
		hm.resize()
		return hm.putValue(hash, key, value)
	}

	hm.size++
	return value
}

func (hm *HashMap) resize() {
	hm.capacity <<= 1

	temTable := hm.table

	for i := 0; i < len(temTable); i++ {
		node := temTable[i]
		if node == nil {
			continue
		}

		hm.table[hm.hash(node.key)] = node
	}
}

func (hm *HashMap) getNodeByHash(hash uint64) *node {
	return hm.table[hash]
}

func newNode(key any, value any) *node {
	return &node{
		key:   key,
		value: value,
	}
}

func newNodeWithNext(key any, value any, next *node) *node {
	return &node{
		key:   key,
		value: value,
		next:  next,
	}
}

func (hm *HashMap) hash(key any) uint64 {
	h := fv.New64a()

	_, _ = h.Write([]byte(ftm.Sprintf("%v", key)))

	hashValue := h.Sum64()

	return (hm.capacity - 1) & (hashValue ^ (hashValue >> 16))
}
