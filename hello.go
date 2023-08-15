package main

import (
  "fmt"
  "rsc.io/quote"
)

func amaine() {
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}

// $ go run bye_asteroid.go // build and run
// Bye, asteroid!

// $ go build -o bin test.go // build
// $ ./bin // and run
// Bye, asteroid!
//ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRUnes(s string) string {
	r := []runes(s)
	for i, j:=0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r) 
}

//////////////////////////////////////////



func mainet() {
  a := 1    // var a int
  b := 3.14 // var b float
  c := "hi" // var c string
  d := true // var d bool
  fmt.Println(a, b, c, d)

  e := []int{1, 2, 3} // slice
  e = append(e, 4)
  fmt.Println(e, len(e), e[0], e[1:3], e[1:], e[:2])
  
  f := make(map[string]int) // map
  f["one"] = 1
  f["two"] = 2
  fmt.Println(f, len(f), f["one"], f["three"])
}

/*
>>> OUTPUT <<<
1 3.14 hi true
[1 2 3 4] 4 1 [2 3] [2 3 4] [1 2]
map[one:1 two:2] 2 1 0
*/

type Person struct{
	Name    string `json:"name"`
	Age     int    `json:"age"`
	isAdmin bool
  }
  
  func mainete() {
	p := Person{
	  Name:    "Mike",
	  Age:     16,
	  isAdmin: false,
	}
	fmt.Println(p, p.Name, p.Age, p.isAdmin)
  }
  
  /*
  >>> OUTPUT <<<
  {Mike 16 false} Mike 16 false
  */