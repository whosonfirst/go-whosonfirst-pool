package main

import (
       "fmt"
       pool "github.com/whosonfirst/go-whosonfirst-pool"
)

type Foo struct {
     Value int64
}

func main() {

     p := pool.NewPool()

     f := Foo{123}

     p.Push(f)

     fmt.Printf("%v", p.Pop())
}
