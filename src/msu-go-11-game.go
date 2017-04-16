// Copyright 2017, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// msu-go-11-game.go [created: Mon, 10 Apr 2017]

package main

import "fmt"

var players Players
var currentPlayer Player
var rooms Rooms

func main() {
	initGame()

	currentPlayer = players[0]
	// fmt.Println(currentPlayer)
	fmt.Println("У меня есть рюкзак?", YesNo(currentPlayer.haveBackpack))
	fmt.Println("положить ключ?", currentPlayer.TakeItem("ключ"))
	fmt.Println("Что в рюкзаке?", currentPlayer.backpack.ItemsToString())
	fmt.Println("Я одел рюкзак?", currentPlayer.ClotheBackpack())
	fmt.Println("Что в рюкзаке?", currentPlayer.backpack.ItemsToString())
	fmt.Println("У меня есть рюкзак?", YesNo(currentPlayer.haveBackpack))
	fmt.Println(currentPlayer.room.Lookup())
	fmt.Println(currentPlayer.Go("комната"))
	fmt.Println(currentPlayer.room.Lookup())
	fmt.Println(currentPlayer.Go("коридор"))
	fmt.Println(currentPlayer.room.Lookup())
	fmt.Println(currentPlayer.Go("комната"))
	fmt.Println("положить ключи?", currentPlayer.TakeItem("ключи"))
	fmt.Println("Что в рюкзаке?", currentPlayer.backpack.ItemsToString())
	fmt.Println(currentPlayer.room.Lookup())
}

func initGame() {
	rooms = Rooms{
		{id: 1, name: "кухня", associated: map[int]string{2: "коридор"}, def: true},
		{id: 2, name: "коридор", associated: map[int]string{1: "кухня", 3: "комната"}},
		{id: 3, name: "комната", associated: map[int]string{2: "коридор"},
			items: []Item{{"ключи"}, {"конспекты"}},
		},
	}

	players.NewPlayer(rooms.GetByName("кухня"))
}
