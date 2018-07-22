package hashtable

// HashTable represents a hashtable with the expect get, put, and delete methods.
type HashTable struct {
	arr         [][]string
	bucketCount int
	Size        int64
}

// CreateHashTable creates a new HashTable.
func CreateHashTable(bucketCount int) *HashTable {
	var ht HashTable
	ht.arr = make([][]string, bucketCount)
	ht.bucketCount = bucketCount
	ht.Size = 0
	return &ht
}

func hash(val string) int {
	tot := 0
	for _, c := range val {
		tot += int(c) * 27
	}
	return tot
}

// Put inserts the value s into the HashTabel ht.
func (ht *HashTable) Put(s string) error {
	index := hash(s) % ht.bucketCount
	if ht.Get(s) == s {
		return nil
	}
	ht.arr[index] = append(ht.arr[index], s)
	ht.Size++
	return nil
}

// Get retrieves the value s if it exists, otherwise it returns an empty string
func (ht *HashTable) Get(s string) string {
	index := hash(s) % ht.bucketCount
	for _, item := range ht.arr[index] {
		if item == s {
			return s
		}
	}
	return ""
}
