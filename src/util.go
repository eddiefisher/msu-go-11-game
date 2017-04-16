// Copyright 2017, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func DefaultForEmptyString(s, defaultString string) string {
	if len(s) == 0 {
		return defaultString
	}
	return s
}

func YesNo(b bool) string {
	if b {
		return "Да"
	}
	return "Нет"
}
