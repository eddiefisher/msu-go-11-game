// Copyright 2017, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// msu-go-11-game.go [created: Mon, 10 Apr 2017]

package main

import "fmt"

var players Players
var rooms Rooms

func main() {
	initGame()

	currentPlayer := players[0]
	currentRoom := currentPlayer.room
	fmt.Println(currentPlayer)
	fmt.Println(currentPlayer.id)
	fmt.Println("Have Backpack:", currentPlayer.haveBackpack)
	fmt.Println(currentRoom.Lookup())
	fmt.Println(currentRoom.CanIGoToRoom("коридор"), "могу пойти в коридор?")
	fmt.Println(currentRoom.CanIGoToRoom("комната"), "могу пойти в комната?")
	// fmt.Println(currentRoom.Lookup(), "перешел в:", currentRoom.name)
}

func initGame() {
	rooms = Rooms{
		{id: 1, name: "кухня", associated: map[int]string{2: "коридор"}, def: true},
		{id: 2, name: "коридор", associated: map[int]string{1: "кухня", 3: "комната"}},
		{id: 3, name: "комната", associated: map[int]string{2: "коридор"}},
	}

	players.NewPlayer(rooms.GetByName("кухня"))
}
