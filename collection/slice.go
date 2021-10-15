package collection

import (
	"reflect"
	"strconv"
	"strings"
)

// Contains checks whether slice/array `s` contains element `elem`
func Contains(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)

	if arrV.Kind() == reflect.Slice || arrV.Kind() == reflect.Array {
		for i := 0; i < arrV.Len(); i++ {

			// XXX - panics if slice element points to an unexported struct field
			// see https://golang.org/pkg/reflect/#Value.Interface
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}

// JoinInt64 join int64 array by specific string
func JoinInt64(pieces []int64, glue string) string {
	if len(pieces) == 0 {
		return ""
	}
	stringPieces := make([]string, len(pieces))
	for i, v := range pieces {
		stringPieces[i] = strconv.FormatInt(v, 10)
	}
	return strings.Join(stringPieces, glue)
}

// IntersectionInt64 get duplicated elements distincted from two arrays, 交集
func IntersectionInt64(a1, a2 []int64) []int64 {
	pending := make(map[int64]int)

	for i := 0; i < len(a1); i++ {
		pending[a1[i]] = 1
	}
	for i := 0; i < len(a2); i++ {
		pending[a2[i]] |= 2
	}
	intersection := make([]int64, 0, len(pending))
	for keys, status := range pending {
		if status == 3 {
			intersection = append(intersection, keys)
		}
	}

	return intersection
}

// GroupInt64 returns a new slice contains slices with maximum length each
func GroupInt64(from []int64, each int) [][]int64 {
	result := make([][]int64, 0, len(from)/each+1)
	for n := 0; ; n += each {
		if n+each >= len(from) {
			result = append(result, from[n:])
			break
		}
		result = append(result, from[n:n+each])
	}

	return result
}

// DiffInt64 returns diff elem from two slice
// 传入 [1,2,3], [2,3,4] 返回 [1,4]
func DiffInt64(s1, s2 []int64) []int64 {
	checksMap := make(map[int64]int)
	for i := 0; i < len(s1); i++ {
		checksMap[s1[i]] = 1
	}
	for i := 0; i < len(s2); i++ {
		checksMap[s2[i]] |= 2
	}

	res := make([]int64, 0, len(checksMap))
	for key, status := range checksMap {
		if status != 3 {
			res = append(res, key)
		}
	}
	return res
}

// FindDuplicateInt64 检测并返回重复数字
func FindDuplicateInt64(arr []int64) []int64 {
	duplicates := make([]int64, 0, 8) //
	m := make(map[int64]int, len(arr))
	for _, v := range arr {
		m[v]++
	}
	for k, v := range m {
		if v > 1 {
			duplicates = append(duplicates, k)
		}
	}
	return duplicates
}

// SplitToInt64 split str with sep, return int64 array
func SplitToInt64(s, sep string) ([]int64, error) {
	if s == "" {
		return []int64{}, nil
	}

	splitList := strings.Split(s, sep)
	result := make([]int64, len(splitList))
	for i, v := range splitList {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		result[i] = id
	}
	return result, nil
}
