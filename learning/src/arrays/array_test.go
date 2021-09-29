package arrays

import "testing"

func TestArray(t *testing.T) {
	var a [2]int
	b := [...]int{1, 2, 3, 4}
	t.Log(a[0], a[1])
	t.Log(b)

	for index, el := range b {
		t.Log(index, el)
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	t.Log(arr[1:3])
	t.Log(arr[:2])
}

func TestSlice(t *testing.T) {
	var s1 []int
	t.Log(len(s1), cap(s1))

	s2 := []int{1, 2, 3, 4, 5}
	t.Log(len(s2), cap(s2))

	s3 := make([]int, 2, 5)
	s3 = append(s3, 1)
	t.Log(s3[0], s3[2], len(s3), cap(s3))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceSharing(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	q2 := year[3:6]
	t.Log(q2, len(q2), cap(q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "六月"
	t.Log(q2)
	t.Log(year)
}
