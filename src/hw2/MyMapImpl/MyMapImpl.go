package MyMapImpl

import (
    "bytes"
    "fmt"
)

type entry struct {
	key   *string
	value *int
        next  *entry
}

type MyMap struct {
        // ~.25 load
	table  [100000]*entry
	length int            
}

func NewMyMap() *MyMap {
    return &MyMap{ length: 0}
}

func hashString(s string) uint32 {
	var h uint32
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h%100000
}

func (m *MyMap) Remove(key string) bool {
	if m != nil {
		hash := hashString(key)
		bucketHead := m.table[hash]
                if (*(*bucketHead).key) == key {
                    tmp := bucketHead
                    (*tmp).key   = nil
                    (*tmp).value = nil
                    (*tmp).next  = nil
                    tmp          = nil
                    bucketHead = bucketHead.next
                }
		for {
                    if (*bucketHead).next == nil {
                        break
                    }
                    if (*(*(*bucketHead).next).key) == key {
                        tmp := (*bucketHead).next
                        (*tmp).key   = nil
                        (*tmp).value = nil
                        (*tmp).next  = nil
                        tmp          = nil

                        (*bucketHead).next = (*(*bucketHead).next).next
                        m.length--
                        return true
                    }
		    bucketHead = (*bucketHead).next
 
		}
	}
	return false
}

func (m *MyMap) At(key string) *int {
    result := 0
    if m != nil {
        bucketHead := m.table[hashString(key)] 
        if bucketHead == nil{
            return &result
        }
        if (*(*bucketHead).key) == key {
            return (*bucketHead).value
        }
        for {
            if (*(*bucketHead).key) == key {
                return (*bucketHead).value
            } 
            if (*bucketHead).next == nil {
                break
            }
            bucketHead = (*bucketHead).next
        }
    }
	return &result
}

func (m *MyMap) Add(key string, value int) int {
    comparisons := 0
    hash := hashString(key)
    bucketHead := m.table[hash]
    comparisons++
    if bucketHead == nil {
        bucketHead = &entry{ key: &key, value: &value, next: nil }
        m.table[hash] = bucketHead
        m.length ++
        return comparisons
    }
    for {
         comparisons++
         if (*(*bucketHead).key) == key {
             (*(*bucketHead).value) = value
             return comparisons
         }
         if (*bucketHead).next == nil {
             (*bucketHead).next = &entry{ key: &key, value: &value, next: nil }
             m.length++
             return comparisons
         }
         bucketHead = (*bucketHead).next
    }
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
            for {
                if bucket != nil {
                    f((*(*bucket).key), (*(*bucket).value))
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

