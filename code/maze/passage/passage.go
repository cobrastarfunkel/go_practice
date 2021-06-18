package passage

// Passage in a maze
type Passage struct {
	IsOpen bool
	Key    string
}

// Check if the Passage Requires a key
func (p *Passage) RequiresKey() bool {
	return len(p.Key) > 0
}

// Try to Open the Passage, if the Passage requires no key but is closed
// It will open the passage
func (p *Passage) Open(key string) {
	if p.RequiresKey() {
		p.IsOpen = key == p.Key
	} else {
		p.IsOpen = true
	}
}
