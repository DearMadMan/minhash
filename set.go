package minhash

type Set struct {
	minhash    *Minhash
	Elements   []string
	Signatures map[int]float64
}

func (s *Set) init() {
	n := s.minhash.permutationsNumber
	s.Signatures = make(map[int]float64)
	for i := 0; i < n; i++ {
		s.Signatures[i] = MaxHash
	}

	for _, el := range s.Elements {
		s.update(el)
	}
}

func (s *Set) update(str string) {
	for i, v := range s.Signatures {
		a := s.minhash.permArgs[0][i]
		b := s.minhash.permArgs[1][i]

		h := float64((uint64(a)*uint64(hash(str)) + uint64(b)) % Prime)
		if h < v {
			s.Signatures[i] = h
		}
	}
}

func (s Set) Jaccard(other Set) float64 {
	var intersections float64
	union := float64(len(s.Signatures))
	for i, v := range other.Signatures {
		if v == s.Signatures[i] {
			intersections++
		}
	}
	return intersections / union
}
