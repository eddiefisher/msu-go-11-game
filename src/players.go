// Copyright 2017, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type Players []Player

type Player struct {
	id           int
	room         Room
	backpack     Backpack
	haveBackpack bool
	quests       Quests
}

func (p *Players) NewPlayer(room Room) Player {
	pl := Player{id: len(*p) + 1, room: room}
	pl.quests.StartQuests()
	*p = append(*p, pl)
	return pl
}
