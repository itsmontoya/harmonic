package seeker

func New(sz int) *Seeker {
	var s Seeker
	s.s = make([]seekItem, 0, sz)
	s.end = -1
	return &s
}

type Seeker struct {
	s   []seekItem
	end int
}

func (s *Seeker) Get(key string) (val int, ok bool) {
	var idx int
	if idx, ok = s.seek(key); ok {
		val = s.s[idx].val
	}

	return
}

func (s *Seeker) Put(key string, val int) {
	idx, found := s.seek(key)
	if found {
		s.s[idx].val = val
		return
	}

	if idx > s.end {
		s.s = append(s.s, seekItem{key, val})
	} else if idx == -1 {
		var ip seekItem
		for i, item := range s.s {
			s.s[i] = ip
			ip = item
		}

		s.s = append(s.s, ip)
		s.s[0] = seekItem{key, val}
	} else {
		s.s = append(s.s[:idx], s.s[idx-1:]...)
		s.s[idx] = seekItem{key, val}
	}

	s.end++
}

func (s *Seeker) ForEach(fn func(key string, val int)) {
	for _, item := range s.s {
		fn(item.key, item.val)
	}
}

// I'm aware my naming is awful, just getting the proof of concept working. I'll clean this all up shortly, I promise
func (s *Seeker) seek(key string) (idx int, found bool) {
	// Fast track for keys which occur AFTER our tail
	if s.end == -1 || key > s.s[s.end].key {
		idx = s.end + 1
		return
	}

	if key < s.s[0].key {
		idx = -1
		return
	}

	return s.seek2(key, 0, s.end)
}

func (s *Seeker) seek2(key string, start, end int) (idx int, found bool) {
	if end-start < 3 {
		return s.seek3(key, start)
	}

	idx = (start + end) / 2

	it := &s.s[idx]
	if key > it.key {
		return s.seek2(key, idx, end)
	} else if key < it.key {
		return s.seek2(key, start, idx)
	} else {
		found = true
		return
	}
}

func (s *Seeker) seek3(key string, start int) (idx int, found bool) {
	for idx = start; idx <= s.end; idx++ {
		it := &s.s[idx]
		if key > it.key {
			continue
		} else if key < it.key {
			return
		} else {
			found = true
			return
		}
	}

	return
}

type seekItem struct {
	key string
	val int
}
