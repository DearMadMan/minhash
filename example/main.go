package main

import (
	"fmt"
	"time"

	"github.com/DearMadMan/minhash"
)

func main() {
	m := minhash.New(128)
	set1 := m.NewSet([]string{
		"minhash", "is", "a", "probabilistic", "data", "structure", "for",
		"estimating", "the", "similarity", "between", "datasets",
	})

	set2 := m.NewSet([]string{
		"minhash", "is", "a", "probability", "data", "structure", "for",
		"estimating", "the", "similarity", "between", "documents",
	})

	set3 := m.NewSet([]string{
		"cats", "are", "tall", "and", "have", "been",
		"known", "to", "sing", "quite", "loudly",
	})

	fmt.Printf("set1 & set2: %f\n", set1.Jaccard(set2))
	fmt.Printf("set1 & set3: %f\n", set1.Jaccard(set3))
	fmt.Printf("set2 & set3: %f\n", set2.Jaccard(set3))

	seed := int64(time.Now().Nanosecond())
	fmt.Printf("\nnew min hash with seed: %d\n\n", seed)
	m2 := minhash.NewWithSeed(128, seed)
	set4 := m2.NewSet([]string{
		"minhash", "is", "a", "probabilistic", "data", "structure", "for",
		"estimating", "the", "similarity", "between", "datasets",
	})

	set5 := m2.NewSet([]string{
		"minhash", "is", "a", "probability", "data", "structure", "for",
		"estimating", "the", "similarity", "between", "documents",
	})

	set6 := m.NewSet([]string{
		"cats", "are", "tall", "and", "have", "been",
		"known", "to", "sing", "quite", "loudly",
	})

	fmt.Printf("set4 & set5: %f\n", set4.Jaccard(set5))
	fmt.Printf("set4 & set6: %f\n", set4.Jaccard(set6))
	fmt.Printf("set5 & set6: %f\n", set5.Jaccard(set6))

}
