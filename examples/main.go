package main

import (
	"log"

	"github.com/vishal1132/gset"
)

func main() {
	a := gset.New()
	b := gset.New()
	a.Append("abcd")
	b.Append("efgh")
	b = gset.Union(a, b)
	a = b
	log.Println(a.GetSet())
	log.Println(b.GetSet())
}
