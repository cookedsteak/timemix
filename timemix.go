package timemix

import (
	"fmt"
	"sort"
)

type T struct {
	Start int64
	End   int64
}

type TSlice []T

func (t TSlice) Len() int { return len(t) }

func (t TSlice) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

func (t TSlice) Less(i, j int) bool { return t[i].Start < t[j].Start }

func (t TSlice) Union() TSlice {
	var s TSlice

	if len(t) > 1 {
		sort.Stable(t)
fmt.Println(t)
		s = append(s, t[0])

		for k, v := range t {
			if v.Start > v.End {
				return s
			}
			if k == 0 {
				continue
			}

			if v.Start >= s[len(s)-1].Start && v.Start <= s[len(s)-1].End {
				// combine
				if v.End > s[len(s)-1].End {
					s[len(s)-1].End = v.End
				}
			} else if v.Start > s[len(s)-1].End {
				// split
				inner := T{Start: v.Start, End: v.End}
				s = append(s, inner)
			}
		}
	}

	return s
}

func (t TSlice) Intersect() TSlice {
	var s TSlice

	if len(t) > 1 {
		sort.Stable(t)
		s = append(s, t[0])

		for k, v := range t {
			if v.Start > v.End {
				return s
			}

			if k == 0 {
				continue
			}

			if v.Start >= s[0].Start && v.Start <= s[0].End {
				s[0].Start = v.Start
				if v.End <= s[0].End {
					s[0].End = v.End
				}
			} else {
				return s[:0]
			}
		}
	}

	return s
}

func (t TSlice) Sum() (sum int64) {
	for i := 0; i < len(t); i++ {
		sum += t[i].End - t[i].Start
	}

	return sum
}
