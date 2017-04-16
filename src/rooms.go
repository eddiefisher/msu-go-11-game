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
	panic("Комната не найдена")
}
