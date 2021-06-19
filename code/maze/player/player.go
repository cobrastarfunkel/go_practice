package player

type Player struct {
	row, col, moves int
	items           map[string]bool
}

func (p *Player) HasItem(item string) bool {
	return p.items[item]
}

func (p *Player) AddItem(item string) {
	p.items[item] = true
}

func (p *Player) UseItem(item string) {
	delete(p.items, item)
}

func (p *Player) SetPosition(row int, col int) {
	p.row, p.col = row, col
}

func (p *Player) IncrementMoves() {
	p.moves++
}

func (p *Player) GetMoves() int {
	return p.moves
}

func (p *Player) GetPosition() (row int, col int) {
	return p.row, p.col
}

func NewPlayer() *Player {
	player := Player{}
	player.items = make(map[string]bool)
	return &player
}
