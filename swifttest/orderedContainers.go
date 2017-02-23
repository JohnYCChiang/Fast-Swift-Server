package swifttest

// orderedContainers holds a slice of containers that can be sorted
// by name.
type orderedContainers []*Container

func (s orderedContainers) Len() int {
	return len(s)
}
func (s orderedContainers) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s orderedContainers) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}
