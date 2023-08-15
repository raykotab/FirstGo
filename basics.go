// ## Type conversion ##
package main

import (
	"fmt"
	"math/cmplx"
	// "reflect"
	// "strconv"
	"math"
)

func main() {
	// a := 3.14
	// b := int(a)
	// fmt.Println(b, reflect.TypeOf(b))

	// c := "12.34"
	// d, _ := strconv.ParseFloat(c, 64)
	// fmt.Println(d, reflect.TypeOf(d))

	// e := false
	// f := fmt.Sprint(e)
	// fmt.Println(f, reflect.TypeOf(f))

	fmt.Println(Sqrt(2))
}

/*
>>> OUTPUT <<<
3 int
12.34 float64
false string

Some pairs of type can convert by casting directly (int(1.23), float(3), uint32(5)),
but for some require an extra library to do a conversion such as
string-int, int-string, bool-string, and string-bool.
*/

//## Zero Values ##

func moine() {
	var a int
	var b float64
	var c bool
	var d string
	fmt.Println(a, b, c, d)

	var e []int
	var f map[string]float64
	fmt.Println(e, f)

	type person struct {
		name string
		age  int
	}
	var g person
	var h *person
	fmt.Println(g, h)
}

/*
  >>> OUTPUT <<<
  0 0 false
  [] map[]
  { 0} <nil>
*/

// ## Multiple Returns ##

func swap(x, y string) (string, string) {
	return y, x
}

func amain() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

/*
  >>> OUTPUT <<<
  world hello
*/

// ## Named Returns ##

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func namain() {
	fmt.Println(split(17))
}

/*
  >>> OUTPUT <<<
  7 10

  Naked return statements should be used only in short functions,
  as with the example shown here. They can harm readability in longer functions.
*/

// ## Short variable declarations ##

func anemain() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

/*
  >>> OUTPUT <<<
  1 2 3 true false no!

   Inside a function, the := short assignment statement can be used
   in place of a var declaration with implicit type.
	Outside a function, every statement begins with a keyword
	(var, func, and so on) and so the := construct is not available.
*/

// ## Types ##

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func antemain() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

/*
  >>> OUTPUT <<<
  	Type: bool Value: false
	Type: uint64 Value: 18446744073709551615
	Type: complex128 Value: (2+3i)

 	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // alias for uint8

	rune // alias for int32
		// represents a Unicode code point

	float32 float64

	complex64 complex128

	Unlike in C, in Go assignment between items of
	different type requires an explicit conversion.
	./prog.go:11:15: cannot use f (variable of type float64)
	as uint value in variable declaration
*/

// ## Constants ##

const Pi = 3.14

func anetemain() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

/*

>>> OUTPUT <<<

Hello 世界
Happy 3.14 Day
Go rules? true

Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.

*/

//## Conditions ##

func moinet() {
	// if, else
	a := 5
	if a > 3 {
		a = a - 3
	} else if a == 3 {
		a = 0
	} else {
		a = a + 3
	}
	fmt.Println(a)

	// switch, case
	b := "NO"
	switch b {
	case "YES":
		b = "Y"
	case "NO":
		b = "N"
	default:
		b = "X"
	}
	fmt.Println(b)
}

/*
  >>> OUTPUT <<<
  2
  N
*/

//## Loop ##

/*In Go, there are control flow statements like other languages:
  if, else, switch, case, but there is only one looping statement in Go which is “for”.
  Because you can replace “while” statement with “for” statement like the example */

func moinete() {
	// for
	c := 3
	for i := 0; i < c; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// replace while with for
	for c > 0 {
		fmt.Print(c, " ")
		c--
	}
	fmt.Println()

	// for with range
	d := []int{0, 1, 1, 2, 3, 5, 8}
	for _, i := range d {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

/*
  >>> OUTPUT <<<
  0 1 2
  3 2 1
  0 1 1 2 3 5 8
*/
/*Moreover, you can iterate on some iterable variable such as
array, slice, and map by using “range”, it return two values
which are index/key and value.*/

// ## Defer ##

func moineter() {
	f()
}

func f() (int, error) {
	defer fmt.Println("DEFER 1")
	defer fmt.Println("DEFER 2")
	defer fmt.Println("DEFER 3")
	fmt.Println("BODY")
	return fmt.Println("RETURN")
}

/*
  >>> OUTPUT <<<
  BODY
  RETURN
  DEFER 3
  DEFER 2
  DEFER 1
*/

/*
func myFunc (a,b string, c int)(string, string, int) {
	x := b
	y := a
	z := c * c
	return x, y, z
}
    Name: Should be named in camelCase / CamelCase.
    Arguments: A function can take zero or more arguments.
		For two or more consecutive arguments with the same type,
		you can define the type on the back of the last argument
		(like “string” in the example).
    Return Types: A function can return zero or more value.
		If returns more than one values,
		you need to cover them with parentheses.
    Body: It is a logic of a function.
*/

func moinetern() {
	fmt.Println(threeTimes("Thank You"))
}

func threeTimes(msg string) (tMsg string) {
	tMsg = msg + ", " + msg + ", " + msg
	return
}

/*
  >>> OUTPUT <<<
  Thank You, Thank You, Thank You

  You also can name result values of a function.
  So, you don’t need to declare returned variables and define what the function returns
  in the return statement. In this case, you need to put parentheses
  into the return types although there is only one argument.
*/

// ## Anonymous Functions ##

func maineterne() {
	a := 1
	b := 1
	c := func(x int) int {
		b *= 2
		return x * 2
	}(a)
	fmt.Println(a, b, c)
}

/*
  >>> OUTPUT <<<
  1 2 2

  You can declare a function inside another function
  and execute immediately after declaring it. It is called
  Anonymous Function that can access the scope of its parent.
  In this case, the anonymous function can access variable b
  which in the main function’s scope
*/

func anonFunc() {
	myFunc := func(x int) int {
		return x * x
	}
	fmt.Println(myFunc(2), myFunc(3))
}

/*
  >>> OUTPUT <<<
  4 9
*/

// ## Blank Identifiers ##

func blanckIdentifier() {
	x, _ := evenOnly(10)
	fmt.Println(x)
}

func evenOnly(n int) (int, error) {
	if n%2 == 0 {
		return n / 2, nil
	}
	return 0, fmt.Errorf("not even")
}

/*
  >>> OUTPUT <<<
  5
*/

func loopea() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func lopeaSinArgs() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

//  IF

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}


func Sqrt(x float64) (float64, int) {
	z := 1.0
	epsilon := 1e-10 // Small value for convergence
	control := 0

	for {
		delta := (z*z - x) / (2 * z)
		z -= delta
		control++
		if math.Abs(delta) < epsilon {
			break // Converged
		}
	}

	return z, control
}
