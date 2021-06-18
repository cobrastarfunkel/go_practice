package room

import (
	"maze/passage"
)

type Room struct {
	Name                                                 string
	Items                                                []string
	NorthPassage, EastPassage, SouthPassage, WestPassage *passage.Passage
}

func (r *Room) AddItem(item string) {
	r.Items = append(r.Items, item)
}

func (r *Room) AquireNextItem() string {
	if len(r.Items) > 0 {
		temp := r.Items[0]
		r.Items = r.Items[1:]
		return temp
	}
	return ""
}
