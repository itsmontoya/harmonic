package seeker

func New(sz int) *Seeker {
	var s Seeker
	s.s = make([]*seekItem, 0, sz)
	s.end = -1
	return &s
}

type Seeker struct {
	s   []*seekItem
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

func (s *Seeker) ForEach(fn func(key string, val int)) {
	for _, item := range s.s {
		fn(item.key, item.val)
	}
}

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

	var state int8
	idx = s.end / 2
	delta := idx / 2

	for {
		item := s.s[idx]
		if delta /= 2; delta == 0 {
			delta = 1
		}

		if item.key < key {
			if delta == 1 {
				if state == 1 {
					idx++
					return
				}
			}

			idx += delta
			state = -1
		} else if item.key > key {
			if delta == 1 {
				if state == -1 {
					return
				}
			}

			idx -= delta
			state = 1
		} else {
			found = true
			return
		}

		//		if half < 1 {
		//			half = 1
		//		}
	}

	return
}

type seekItem struct {
	key string
	val int
}
