package seeker

// New will return a Seeker with the requested size
func New(sz int) *Seeker {
	var s Seeker
	s.s = make([]*seekItem, 0, sz)
	s.end = -1
	return &s
}

// Seeker is a sorted in-memory data store
type Seeker struct {
	s   []*seekItem
	end int
}

// getIndex will get the index of a provided key
func (s *Seeker) getIndex(key string) (idx int, found bool) {
	// Fast track for keys which occur AFTER our tail
	if s.end == -1 || key > s.s[s.end].key {
		idx = s.end + 1
		return
	}

	if key < s.s[0].key {
		idx = -1
		return
	}

	return s.seek(key, 0, s.end)
}

func (s *Seeker) seek(key string, start, end int) (idx int, found bool) {
	if end-start < 3 {
		return s.iterFind(key, start)
	}

	idx = (start + end) / 2

	if key > s.s[idx].key {
		return s.seek(key, idx, end)
	} else if key < s.s[idx].key {
		return s.seek(key, start, idx)
	} else {
		found = true
		return
	}
}

func (s *Seeker) iterFind(key string, start int) (idx int, found bool) {
	for idx = start; idx <= s.end; idx++ {
		if key > s.s[idx].key {
			continue
		} else if key < s.s[idx].key {
			return
		} else {
			found = true
			return
		}
	}

	return
}

// Get will retrieve a value for a provided key
func (s *Seeker) Get(key string) (val int, ok bool) {
	var idx int
	if idx, ok = s.getIndex(key); ok {
		val = s.s[idx].val
	}

	return
}

// Put will insert a value for a provided key
func (s *Seeker) Put(key string, val int) {
	idx, found := s.getIndex(key)
	if found {
		s.s[idx].val = val
		return
	}

	if idx > s.end {
		s.s = append(s.s, &seekItem{key, val})
	} else if idx == -1 {
		var ip *seekItem
		for i, item := range s.s {
			s.s[i] = ip
			ip = item
		}

		s.s = append(s.s, ip)
		s.s[0] = &seekItem{key, val}
	} else {
		s.s = append(s.s[:idx], s.s[idx-1:]...)
		s.s[idx] = &seekItem{key, val}
	}

	s.end++
}

// ForEach will iterate through each key and value within the data store
func (s *Seeker) ForEach(fn func(key string, val int)) {
	for _, item := range s.s {
		fn(item.key, item.val)
	}
}

type seekItem struct {
	key string
	val int
}
