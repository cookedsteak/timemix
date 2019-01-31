package timemix

import (
	"math/rand"
	"testing"
	"time"
)

// Make @param:t pieces of fake time sets
func makeTimes(t int) TSlice {
	var set TSlice
	rand.Seed(time.Now().Unix())

	for i := 0; i < t; i++ {
		randStart := rand.Int63n(50)
		randEnd := randStart + rand.Int63n(25) + 1
		set = append(set, T{Start: randStart, End: randEnd})
	}

	return set
}

func Test_Union(t *testing.T) {
	testSet := makeTimes(10)

	t.Logf("Input set: %+v\n", testSet)

	res := testSet.Union()

	t.Logf("Output set: %+v\n", res)
	t.Logf("Sum of time set: %d", res.Sum())
}

func Test_Intersect(t *testing.T) {
	testSet := makeTimes(2)

	t.Logf("Input set: %+v\n", testSet)

	res := testSet.Intersect()

	t.Logf("Output set: %+v\n", res)
	t.Logf("Sum of time set: %d", res.Sum())
}
