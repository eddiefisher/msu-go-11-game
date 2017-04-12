// Copyright 2017, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type Rooms []Room

func (r Rooms) GetByName(name string) Room {
	for _, v := range r {
		if v.name == name {
			return v
		}
	}
	panic("Room Not Found")
}

type Room struct {
	id         int
	def        bool
	name       string
	items      []Item
	associated map[int]string
}

func (r Room) Lookup() string {
	if r.def {
		return "ты находишься на кухне, . можно пройти - " + r.AssociatedRooms()
	}
	return "пустая комната. можно пройти - " + r.AssociatedRooms()
}

func (r Room) AssociatedRooms() string {
	var res string
	for _, v := range r.associated {
		res += v
	}
	return res
}

func (r Room) CanIGoToRoom(s string) bool {
	room := rooms.GetByName(s)
	_, ok := r.associated[room.id]
	return ok
}
