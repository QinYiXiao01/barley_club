package utils

import (
	"math/rand"
	"time"
)

func RandomName(n int) string {
	var name = []byte("diuqwuibcunoewncmenwphfiwehonnCIOEWNOICNWIEH")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())

	for i := range result {
		result[i] = name[rand.Intn(len(name))]
	}
	return string(result)
}
