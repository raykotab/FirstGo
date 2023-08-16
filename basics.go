
package main

import (
	"fmt"
	"math/cmplx"
	"reflect"
	"strconv"
	"math"
	"time"
)

func main() {
	good()
	printTypesOf()
	kindOfZeroes()
	printSwap()
	printSplit()
	printVariables()
	printVarTypes()
	printConst()
	printIfs()
	printLoop()
	printDeference()
	printThreeTimes()
	printAnonFunc()
	blanckIdentifier()
	evenOnly(9)
	loopea()
	lopeaNoArgs()
	fmt.Println(Sqrt(2))
	fmt.Println(pow(2, 6, 5))
	deference()
	printPointersDemo()
	vertex()
	printFirstStructs()
	printArray()
	printSliceDemo()
	printSliceAccessesArray()
	sliceLiterals()

	
	// fmt.Println()
}


func printTypesOf() {
	a := 3.14
	b := int(a)
	fmt.Println(b, reflect.TypeOf(b))

	c := "12.34"
	d, _ := strconv.ParseFloat(c, 64)
	fmt.Println(d, reflect.TypeOf(d))

	e := false
	f := fmt.Sprint(e)
	fmt.Println(f, reflect.TypeOf(f))
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

func kindOfZeroes() {
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

func printSwap() {
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

func printSplit() {
	fmt.Println(split(17))
}

/*
  >>> OUTPUT <<<
  7 10

  Naked return statements should be used only in short functions,
  as with the example shown here. They can harm readability in longer functions.
*/

// ## Short variable declarations ##

func printVariables() {
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

func printVarTypes() {
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

func printConst() {
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

func printIfs() {
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

func printLoop() {
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

func printDeference() {
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

func printThreeTimes() {
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

func printAnonFunc() {
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

func lopeaNoArgs() {
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


/* Switch without a condition is the same as switch true.

This construct can be a clean way to write long if-then-else chains. */

func good() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

/*Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order. */

func deference() {
	defer fmt.Println("world")

	fmt.Println("hello")
}



//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//POINTERS * POINTERS & POINTERS * POINTERS & POINTERS * POINTERS & POINTERS * POINTERS & POINTERS * POINTERS & POINT
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func printPointersDemo() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
	fmt.Println(2701 / 37)
}

//*   output 
//	42
//	21
//	73
//	73


//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//A struct is a collection of fields. A struct is a collection of fields. A struct is a collection of fields. A struct
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


type Vertex struct {
	X int
	Y int
}

func vertex() {
	fmt.Println(Vertex{1, 2})
}

/*
Struct fields are accessed using a dot. 
v.X = 4

Struct fields can be accessed through a struct pointer. 
v := Vertex{1, 2}
	p := &v
	p.X = 1e9

	output 
	{1000000000 2}
*/

/*
 A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.) 

*/

type firstStruct struct {
	X, Y int
}

var (
	v1 = firstStruct{1, 2}  // has type Vertex
	v2 = firstStruct{X: 1}  // Y:0 is implicit
	v3 = firstStruct{}      // X:0 and Y:0
	p  = &v1 // has type *Vertex
)

func printFirstStructs() {
	fmt.Println(v1, p, v2, v3)
}

/*
The type [n]T is an array of n values of type T. 
An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays
*/

func printArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

/*
	OUTPUT
Hello World
[Hello World]
[2 3 5 7 11 13]
*/

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// SLICE · SLICE · SLICE · SLICE · SLICE · SLICE · SLICE · SLICE · SLICE · SLICE · SLICE · SLICE · SLICE · SLICE · 
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
 An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type []T is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

a[low : high]

This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:

a[1:4] = [3 5 7]

*/

func printSliceDemo() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

/*

 A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes. 

*/

func printSliceAccessesArray() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	a[0] = "yyy"
	b[1] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

/*
	OUTPUT

[John Paul George Ringo]
[John Paul] [Paul George]
[yyy Paul] [Paul XXX]
[yyy Paul XXX Ringo]
*/


/*
 A slice literal is like an array literal without the length.

This is an array literal:

[3]bool{true, true, false}

And this creates the same array as above, then builds a slice that references it:

[]bool{true, true, false}
*/

func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

/*
	OUTPUT
[2 3 5 7 11 13]
[true false true true false true]
[{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
*/