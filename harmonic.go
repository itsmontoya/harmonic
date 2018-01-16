package harmonic

// New will return a Harmonic with the requested size
func New(sz int) *Harmonic {
	var h Harmonic
	h.s = make([]*seekItem, 0, sz)
	h.end = -1
	return &h
}

// Harmonic is a sorted in-memory data store
type Harmonic struct {
	s   []*seekItem
	end int
}

// getIndex will get the index of a provided key
func (h *Harmonic) getIndex(key string) (idx int, found bool) {
	// Fast track for keys which occur AFTER our tail
	if h.end == -1 || key > h.s[h.end].key {
		idx = h.end + 1
		return
	}

	if key < h.s[0].key {
		idx = -1
		return
	}

	return h.seek(key, 0, h.end)
}

func (h *Harmonic) seek(key string, start, end int) (idx int, found bool) {
	if end-start < 3 {
		return h.iterFind(key, start)
	}

	idx = (start + end) / 2

	if key > h.s[idx].key {
		return h.seek(key, idx, end)
	} else if key < h.s[idx].key {
		return h.seek(key, start, idx)
	}

	found = true
	return
}

func (h *Harmonic) iterFind(key string, start int) (idx int, found bool) {
	for idx = start; idx <= h.end; idx++ {
		if key > h.s[idx].key {
			continue
		} else if key < h.s[idx].key {
			return
		} else {
			found = true
			return
		}
	}

	return
}

// Get will retrieve a value for a provided key
func (h *Harmonic) Get(key string) (val int, ok bool) {
	var idx int
	if idx, ok = h.getIndex(key); ok {
		val = h.s[idx].val
	}

	return
}

// Put will insert a value for a provided key
func (h *Harmonic) Put(key string, val int) {
	idx, found := h.getIndex(key)
	if found {
		h.s[idx].val = val
		return
	}

	if idx > h.end {
		h.append(key, val)
	} else if idx == -1 {
		h.prepend(key, val)
	} else {
		h.insert(idx, key, val)
	}

	h.end++
}
func (h *Harmonic) prepend(key string, val int) {
	var ip *seekItem
	for i, item := range h.s {
		h.s[i] = ip
		ip = item
	}

	h.s = append(h.s, ip)
	h.s[0] = &seekItem{key, val}
}

func (h *Harmonic) append(key string, val int) {
	h.s = append(h.s, &seekItem{key, val})
}

func (h *Harmonic) insert(idx int, key string, val int) {
	h.s = append(h.s[:idx], h.s[idx-1:]...)
	h.s[idx] = &seekItem{key, val}
}

// ForEach will iterate through each key and value within the data store
func (h *Harmonic) ForEach(fn func(key string, val int)) {
	for _, item := range h.s {
		fn(item.key, item.val)
	}
}

type seekItem struct {
	key string
	val int
}
