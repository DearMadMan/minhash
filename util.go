package minhash

import (
	"hash/fnv"
)

func hash(str string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return h.Sum32()
}
