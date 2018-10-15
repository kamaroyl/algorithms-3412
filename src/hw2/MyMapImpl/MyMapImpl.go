package MyMapImpl

import (
    "bytes"
    "fmt"
)

type MyMap struct {
	table  map[uint32][]entry // maps hash to bucket; entry.key==nil means unused
	length int                // number of map entries
}

func NewMyMap() *MyMap {
    return &MyMap{ table: make(map[uint32][]entry), length: 0}
}

func hashString(s string) uint32 {
	var h uint32
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h
}


type entry struct {
	key   *string
	value *int
}

func (m *MyMap) Remove(key string) bool {
	if m != nil && m.table != nil {
		hash := hashString(key)
		bucket := m.table[hash]
		for i, e := range bucket {
			if e.key != nil && key == *e.key {
				// We can't compact the bucket as it
				// would disturb iterators.
				bucket[i] = entry{}
				m.length--
				return true
			}
		}
	}
	return false
}

func (m *MyMap) At(key string) int {
	if m != nil && m.table != nil {
		for _, e := range m.table[hashString(key)] {
			if e.key != nil && key == *e.key {
				return *e.value
			}
		}
	}
	return 0
}

func (m *MyMap) Add(key string, value int) int {
        comparisons := 0
	if m.table != nil {
		hash := hashString(key)
		bucket := m.table[hash]
		var hole *entry
		for i, e := range bucket {
                        comparisons++
			if e.key == nil {
				hole = &bucket[i]
			} else if key == (*e.key) {
				*bucket[i].value = value
				return comparisons
			}
		}

		if hole != nil {
			*hole = entry{&key, &value} // overwrite deleted entry
		} else {
			m.table[hash] = append(bucket, entry{&key, &value})
		}
	} else {
		hash := hashString(key)
		m.table = map[uint32][]entry{hash: {entry{&key, &value}}}
	}

	m.length++
	return comparisons
}

func (m *MyMap) Len() int {
    if m != nil {
        return m.length
    }
    return 0
}

func (m *MyMap) Iterate(f func(key string, value int)) {
	if m != nil {
		for _, bucket := range m.table {
			for _, e := range bucket {
				if e.key != nil {
					f(*e.key, *e.value)
				}
			}
		}
	}
}

// Keys returns a new slice containing the set of map keys.
// The order is unspecified.
func (m *MyMap) Keys() []string {
	keys := make([]string, 0, m.Len())
	m.Iterate(func(key string, _ int) {
		keys = append(keys, key)
	})
	return keys
}

func (m *MyMap) toString(values bool) string {
	if m == nil {
		return "{}"
	}
	var buf bytes.Buffer
	fmt.Fprint(&buf, "{")
	sep := ""
	m.Iterate(func(key string, value int) {
		fmt.Fprint(&buf, sep)
		sep = ", "
		fmt.Fprint(&buf, key)
		if values {
			fmt.Fprintf(&buf, ": %q", value)
		}
	})
	fmt.Fprint(&buf, "}")
	return buf.String()
}

func (m *MyMap) String() string {
	return m.toString(true)
}

func (m *MyMap) KeysString() string {
	return m.toString(false)
}

