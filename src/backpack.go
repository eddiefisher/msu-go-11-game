// Copyright 2017, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type Backpack []Item

func (b Backpack) ItemsToString() string {
	var res string
	for _, v := range b {
		res += v.name
	}
	res = DefaultForEmptyString(res, "Ничего нет")
	return res
}
