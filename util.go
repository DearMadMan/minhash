package minhash

import (
	"hash/fnv"
	"math/rand"
	"time"
)

func hash(str string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return h.Sum32()
}

func randInt() uint32 {
	rand.NewSource(int64(time.Now().Nanosecond()))
	return rand.Uint32()
}
