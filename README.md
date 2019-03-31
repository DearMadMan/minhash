[Minhashing](https://en.wikipedia.org/wiki/MinHash) is an efficient similarity estimation technique that is often used to identify near-duplicate documents in large text collections. This package offers a Golang implementation of the minhash algorithm.

## Usage

```go
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
```

## Refer
[duhaime/minhash](https://github.com/duhaime/minhash)