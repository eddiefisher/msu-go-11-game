// Copyright 2017, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type RoomPlace struct {
	name  string
	items []Item
}

func (rp RoomPlace) HaveItem(s string) (Item, bool) {
	for _, v := range rp.items {
		if v.name == s {
			return v, true
		}
	}
	return Item{}, false
}
