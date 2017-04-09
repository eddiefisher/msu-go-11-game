// Copyright 2017, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// msu-go-11-game.go [created: Mon, 10 Apr 2017]

package main

import "fmt"

var (
  Version string
  Build   string
)

func main() {
  fmt.Println("Version: ", Version)
  fmt.Println("Built Time: ", Build)
  fmt.Println("Hello, Eddie Fisher!")
}
