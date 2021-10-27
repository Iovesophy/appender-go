# appender-go
basic append ops lib by golang

Usage

```Go
package main

import (
	"fmt"
	"log"

	"github.com/Iovesophy/appender-go"
)

type Demo struct {
	t string
	a []interface{}
}

func main() {
	demo := Demo{}

	demo.t = "declare"
	demo.a = []interface{}{"hoge", "huga", 100, true}
	fmt.Println(demo)

	demo.t = "push"
	demo.a = appender.Push(demo.a, "pushtest")
	fmt.Println(demo)

	demo.t = "pop"
	demo.a = appender.Pop(demo.a)
	fmt.Println(demo)

	demo.t = "unshift"
	demo.a = appender.Unshift(demo.a, "unshifttest")
	fmt.Println(demo)

	demo.t = "shift"
	demo.a = appender.Shift(demo.a)
	fmt.Println(demo)

	demo.t = "insert"
	demo.a = appender.Insert(demo.a, "inserttest", 3)
	fmt.Println(demo)

	demo.t = "remove"
	demo.a = appender.Remove(demo.a, 3)
	fmt.Println(demo)

	demo.t = "concat"
	demo.a = appender.Concat(demo.a, demo.a)
	fmt.Println(demo)

	demo.t = "edit"
	demo.a = appender.Edit(demo.a, "edittest", 3)
	fmt.Println(demo)

	var err error

	demo.t = "picker(flag)"
	demo.a, err = appender.Picker(demo.a, []interface{}{true, false, true, false, false, false, false, false})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(demo)

	demo.t = "picker(index)"
	demo.a, err = appender.Picker(demo.a, []interface{}{0, 1})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(demo)
}

```

## License
Copyright (c) 2021 Kazuya yuda.
This software is released under the MIT License, see LICENSE.
https://opensource.org/licenses/mit-license.php
