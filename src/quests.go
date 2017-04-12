// Copyright 2017, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type Quests []Quest

type Quest struct {
	isDone bool
	name   string
}

func (q *Quests) StartQuests() {
	*q = Quests{
		{name: "собрать рюкзак"},
		{name: "идти в универ"},
	}
}
