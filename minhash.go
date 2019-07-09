package minhash

import (
	"math"
	"math/rand"
)

const (
	// Prime is the smallest prime larger than the largest
	// possible hash value (max hash = 32 bit int)
	Prime   uint64  = 4294967311
	MaxHash float64 = 1<<32 - 1
)

type Minhash struct {
	permutationsNumber int
	permArgs           map[int][]uint32
	rand               *rand.Rand
}

func New(permNumber int) Minhash {
	m := Minhash{permutationsNumber: permNumber, rand: rand.New(rand.NewSource(math.MaxInt64))}
	m.init()
	return m
}

func NewWithSeed(permNumber int, seed int64) Minhash {
	m := Minhash{permutationsNumber: permNumber, rand: rand.New(rand.NewSource(seed))}
	m.init()
	return m
}

func (m *Minhash) init() {
	used := make(map[uint32]bool)
	m.permArgs = make(map[int][]uint32)
	for i := 0; i < 2; i++ {
		numbers := []uint32{}
		for j := 0; j < m.permutationsNumber; j++ {
			ok := true
			for ok {
				n := m.randInt()
				_, ok = used[n]
				if !ok {
					used[n] = true
					numbers = append(numbers, n)
				}
			}
		}
		m.permArgs[i] = numbers
	}
}

func (m *Minhash) NewSet(els []string) Set {
	s := Set{Elements: els, minhash: m}
	s.init()
	return s
}

func (m *Minhash) randInt() uint32 {
	return m.rand.Uint32()
}
