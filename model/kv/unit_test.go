package kv

import (
	"testing"
)

func TestOrderedMap(t *testing.T) {
	testCases := []struct {
		name      string
		addKey    []int
		addValue  []int
		removeKey []int
		want      func(int, int) bool
	}{
		{
			name:      "CommonCase-1",
			addKey:    []int{1, 2, 3, 4, 5},
			addValue:  []int{1, 2, 3, 4, 5},
			removeKey: []int{},
			want: func(k int, v int) bool {
				if k != v {
					return false
				} else {
					return true
				}
			},
		},
		{
			name:      "CommonCase-2",
			addKey:    []int{1, 2, 3, 4, 5},
			addValue:  []int{1, 2, 3, 4, 5},
			removeKey: []int{3, 4, 5, 6},
			want: func(k int, v int) bool {
				if k != v {
					return false
				} else {
					return true
				}
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			m := NewOrderedMap[int, int]()
			length := len(tt.addKey)

			// 测试OrderedMap.Set
			for i := 0; i < length; i++ {
				m.Set(tt.addKey[i], tt.addValue[i])
			}

			// 测试OrderedMap.Length
			if l := m.Length(); l != uint(length) {
				t.Errorf("incorrect length for OrderedMap, expected %d, got %d", length, l)
			}

			// 测试OrderedMap.Remove
			for _, i := range tt.removeKey {
				m.Remove(i)
			}

			// 测试OrderedMap.Range
			m.Range(func(k int, v int) bool {
				if judge := tt.want(k, v); !judge {
					t.Errorf("incorrect judge result")
				}
				return true
			})
		})
	}
}
