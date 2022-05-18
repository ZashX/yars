package src

// Index keeps the mapping between names (string) and indices (integer).
type Index interface {
	Len() int
	Add(name string)
	ToNumber(name string) int
	ToName(index int) string
	GetNames() []string
}

// MapIndex manages the map between sparse Names and dense indices. A sparse ID is
// a user ID or item ID. The dense index is the internal user index or item index
// optimized for faster parameter access and less memory usage.
type MapIndex struct {
	Numbers map[string]int // sparse ID -> dense index
	Names   []string       // dense index -> sparse ID
}

// NotId represents an ID doesn't exist.
const NotId = -1

// NewMapIndex creates a MapIndex.
func NewMapIndex() *MapIndex {
	set := new(MapIndex)
	set.Numbers = make(map[string]int)
	set.Names = make([]string, 0)
	return set
}

// Len returns the number of indexed Names.
func (idx *MapIndex) Len() int {
	if idx == nil {
		return 0
	}
	return len(idx.Names)
}

// Add adds a new ID to the indexer.
func (idx *MapIndex) Add(name string) {
	if _, exist := idx.Numbers[name]; !exist {
		idx.Numbers[name] = len(idx.Names)
		idx.Names = append(idx.Names, name)
	}
}

// ToNumber converts a sparse ID to a dense index.
func (idx *MapIndex) ToNumber(name string) int {
	if denseId, exist := idx.Numbers[name]; exist {
		return denseId
	}
	return NotId
}

// ToName converts a dense index to a sparse ID.
func (idx *MapIndex) ToName(index int) string {
	return idx.Names[index]
}

// GetNames returns all names in current index.
func (idx *MapIndex) GetNames() []string {
	return idx.Names
}
