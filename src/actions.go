// Copyright 2017, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strings"
)

func handleCommand(cmd string) string {
	actions := strings.Split(cmd, " ")
	for _, v := range actions {
		fmt.Println(translate(v))
	}
	return ""
}

func translate(cmd string) string {
	list := map[string]string{
		"осмотреться": "lookup",
		"идти":        "go",
		"взять":       "take",
		"применить":   "use",
		"одеть":       "clothe",
	}
	return list[cmd]
}
