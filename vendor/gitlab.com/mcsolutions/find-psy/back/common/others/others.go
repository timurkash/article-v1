package others

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func Contains[T comparable](slice []T, item T) bool {
	sliceMap := SliceToMap(slice)
	_, ok := sliceMap[item]
	return ok
}

func SliceToMap[T comparable](slice []T) map[T]struct{} {
	sliceMap := make(map[T]struct{}, len(slice))
	for _, s := range slice {
		sliceMap[s] = struct{}{}
	}
	return sliceMap
}

func Intersect[T comparable](slice1, slice2 []T) []T {
	slice1Map := SliceToMap(slice1)
	slice2Map := SliceToMap(slice2)
	var res []T
	for k := range slice1Map {
		if _, ok := slice2Map[k]; ok {
			res = append(res, k)
		}
	}
	return res
}

func Unique[T comparable](slice []T) (res []T) {
	keys := make(map[T]struct{})
	for _, s := range slice {
		if _, ok := keys[s]; !ok {
			keys[s] = struct{}{}
			res = append(res, s)
		}
	}
	return
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func DoEvery(d time.Duration, f func(time.Time)) {
	go func() {
		for x := range time.Tick(d) {
			f(x)
		}
	}()
}

func GetWeek(offset int) (*time.Time, *time.Time) {
	now := time.Now()
	weekBegin := time.Time{}
	switch now.Weekday() {
	case time.Monday:
		weekBegin = now
	case time.Tuesday:
		weekBegin = now.Add(-24 * time.Hour)
	case time.Wednesday:
		weekBegin = now.Add(-48 * time.Hour)
	case time.Thursday:
		weekBegin = now.Add(-72 * time.Hour)
	case time.Friday:
		weekBegin = now.Add(-96 * time.Hour)
	case time.Saturday:
		weekBegin = now.Add(-120 * time.Hour)
	case time.Sunday:
		weekBegin = now.Add(-144 * time.Hour)
	}
	weekBegin = weekBegin.Truncate(24 * time.Hour).Add(time.Duration(offset*168) * time.Hour)
	weekEnd := weekBegin.Add(168 * time.Hour)
	return &weekBegin, &weekEnd
}

func Shorten(str string, limit int) string {
	if len(str) < limit {
		return str
	}
	str = FirstN(str, limit-3) + "..."
	return str
}

func FirstN(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}

func GetNameExt(filename string) (string, string) {
	parts := strings.Split(filename, ".")
	l := len(parts)
	if l == 1 {
		return filename, ""
	}
	return strings.Join(parts[:l-1], "."), parts[l-1]
}

func Min(value int, values ...int) int {
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

var re = regexp.MustCompile(`\W+`)

func GetLatinized(title string) string {
	t := strings.ToLower(title)
	t = strings.ReplaceAll(t, " ", "_")
	t = strings.ReplaceAll(t, "а", "a")
	t = strings.ReplaceAll(t, "б", "b")
	t = strings.ReplaceAll(t, "в", "v")
	t = strings.ReplaceAll(t, "г", "g")
	t = strings.ReplaceAll(t, "д", "d")
	t = strings.ReplaceAll(t, "е", "e")
	t = strings.ReplaceAll(t, "ё", "yo")
	t = strings.ReplaceAll(t, "ж", "zh")
	t = strings.ReplaceAll(t, "з", "z")
	t = strings.ReplaceAll(t, "и", "i")
	t = strings.ReplaceAll(t, "й", "y")
	t = strings.ReplaceAll(t, "к", "k")
	t = strings.ReplaceAll(t, "л", "l")
	t = strings.ReplaceAll(t, "м", "m")
	t = strings.ReplaceAll(t, "н", "n")
	t = strings.ReplaceAll(t, "о", "o")
	t = strings.ReplaceAll(t, "п", "p")
	t = strings.ReplaceAll(t, "р", "r")
	t = strings.ReplaceAll(t, "с", "s")
	t = strings.ReplaceAll(t, "т", "t")
	t = strings.ReplaceAll(t, "у", "u")
	t = strings.ReplaceAll(t, "ф", "f")
	t = strings.ReplaceAll(t, "х", "kh")
	t = strings.ReplaceAll(t, "ц", "ts")
	t = strings.ReplaceAll(t, "ч", "ch")
	t = strings.ReplaceAll(t, "ш", "sh")
	t = strings.ReplaceAll(t, "щ", "sch")
	t = strings.ReplaceAll(t, "ъ", "y")
	t = strings.ReplaceAll(t, "ы", "y")
	t = strings.ReplaceAll(t, "ь", "")
	t = strings.ReplaceAll(t, "э", "e")
	t = strings.ReplaceAll(t, "ю", "yu")
	t = strings.ReplaceAll(t, "я", "ya")
	t = strings.ReplaceAll(t, "-", "_")
	t = strings.ReplaceAll(t, "--", "_")
	t = re.ReplaceAllString(t, "")
	return t
}

func GetExecutionDuration() func() {
	start := time.Now()
	return func() {
		fmt.Println(time.Since(start))
	}
}
