# Resources

[A Tour of GO](https://tour.golang.org)  
[Inside the Go Playground - Architecture](https://blog.golang.org/playground)

# ToC

- [Resources](#resources)
- [ToC](#toc)
- [Basics](#basics)
  - [Package, variables and Functions](#package-variables-and-functions)
    - [Imports](#imports)
    - [Exported names](#exported-names)
    - [Functions](#functions)
    - [Multiple Results](#multiple-results)
    - [Named return values](#named-return-values)
    - [Variables](#variables)
    - [Variables with initializers](#variables-with-initializers)
    - [Short variable declarations](#short-variable-declarations)
    - [Basic types](#basic-types)
    - [Zero values](#zero-values)
    - [Type conversions](#type-conversions)
    - [Type inference](#type-inference)
    - [Constants](#constants)
    - [Numeric Constants](#numeric-constants)
  - [Flow control statements: for, if, else, switch and defer](#flow-control-statements-for-if-else-switch-and-defer)
    - [For loop](#for-loop)
    - [For is Go's "while"](#for-is-gos-while)
    - [Forever (infinite loop)](#forever-infinite-loop)
    - [If condition](#if-condition)
    - [If with a short statement](#if-with-a-short-statement)
    - [If and else](#if-and-else)
    - [Exercise: Loops and Functions](#exercise-loops-and-functions)
    - [Switch](#switch)
    - [Switch evaluation order](#switch-evaluation-order)
    - [Switch with no expression](#switch-with-no-expression)
    - [Defer](#defer)
    - [Stacking defers](#stacking-defers)
  - [More types: structs, slices, and maps](#more-types-structs-slices-and-maps)
    - [Pointers](#pointers)
    - [Structs](#structs)
    - [Struct Fields](#struct-fields)
    - [Pointers to structs](#pointers-to-structs)
    - [Struct Literals](#struct-literals)
    - [Arrays](#arrays)
    - [Slices](#slices)
    - [Slices are like references to arrays](#slices-are-like-references-to-arrays)
    - [Slice literals](#slice-literals)
    - [Slice defaults](#slice-defaults)
    - [Slice length and capacity](#slice-length-and-capacity)
    - [Nil Slices](#nil-slices)
    - [Creating a slice with make](#creating-a-slice-with-make)
    - [Slices of slices](#slices-of-slices)
    - [Appending to a slice](#appending-to-a-slice)
    - [Range](#range)
    - [Range continued](#range-continued)
    - [Exercise: Slices](#exercise-slices)
    - [Maps](#maps)
    - [Map Literals](#map-literals)
    - [Map literals continued](#map-literals-continued)
    - [Mutating Maps](#mutating-maps)
    - [Exercise: Maps](#exercise-maps)
    - [Function values](#function-values)
    - [Function closures](#function-closures)
    - [Relevant Functional Programming Concepts](#relevant-functional-programming-concepts)
    - [Exercise: Fibonacci closure](#exercise-fibonacci-closure)
- [Methods and interfaces](#methods-and-interfaces)
- [Concurrency](#concurrency)

# Basics

## Package, variables and Functions

Learn the basic components of any Go program.

### Imports

Multiple import statements

```go
import "fmt"
import "math"
```

2."Factored" import statement (preferred way)

```go
import(
    "fmt",
    "math/rand"
)
```

### Exported names

In go, a name is exported if it begins with a capital letter.  
For examples, in `math.Pi` the `Pi` is exported.  
When importing a package, you can only refer to its exported names. Any "unexported" names are not accessible from outside the package.  
> What is this Exported thing?  
> A: seems similar to export from js modules, except for the title casing part

### Functions

Notable Difference: The return type is **after** the function arguments.  
More info at [Go's Declaration Syntax](https://blog.golang.org/gos-declaration-syntax)

```go
package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}
```

When two or more consecutive named function parameters share a type, you can omit the type from all but the last one.  
i.e. `(x int, y int)` can be shortened to `x, y int` in the example above.  

```go
func add(x, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}
```

### Multiple Results

A function can return multiple values. Here, the `swap` function returns 2 strings.

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}
```

### Named return values

Go's return values may be named. If so, they are treated as variables defined at the top of the function.
These names should be used to document the meaning of the return values.  
A return statement without arguments returns the named return values. This is known as a "naked" return.  
Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.  

```go
package main

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17))
}
```

### Variables

The `var` statement declares a list of variables. As in function argument lists, the type is at last.  
A `var` statement can be at package or function level. We see both in this example.

```go
package main

import "fmt"

var c, python, java bool

func main() {
    var i int
    fmt.Println(i, c, python, java)
}
```

### Variables with initializers

A var declaration can include initializers, one per variable.  
If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
    var c, python, java = true, false, "no!"
    fmt.Println(i, j, c, python, java)
}
```

### Short variable declarations

Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with implicit type.  
**Outside a function, every statement begins with a keyword** (`var`, `func`, and so on) and so the `:=` construct is not available.

```go
package main

import "fmt"

func main() {
    var i, j int = 1, 2
    k := 3
    c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java)
}
```

> What is short assignment?  
> [= vs :=](https://www.godesignpatterns.com/2014/04/assignment-vs-short-variable-declaration.html)

### Basic types

Go's basic types are

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

The following example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

```go
package main

import (
    "fmt"
    "math/cmplx"
)

var (
    ToBe   bool       = false
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
    fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

The `int`, `uint`, and `uintptr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use `int` unless you have a specific reason to use a sized or unsigned integer type.

More detailed types in [type.go](https://golang.org/src/go/types/type.go)  

### Zero values

Variable declared without an explicit initial value are given their *zero value*
The zero value is:

- `0` for numeric types,
- `false` for the boolean type, and
- `""` (the empty string) for strings.

```go
package main

import "fmt"

func main() {
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```

### Type conversions

The expression `T(v)` converts the value `v` to the type `T`.

Some numeric conversions:

```go
var i int = 42
var f float64 = float64(i)
var u unit = uint(f)
```

In simpler *short assignment* notation:

```go
i := 42
f := float64(i)
u := uint(f)
```

Unlike in C, in Go assignment between items of different type requires an explicit conversion i.e **Explicit Typecasting**. Try removing the float64 or uint conversions in the example and see what happens.

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(x*x + y*y))
    var z uint = uint(f)
    fmt.Println(x, y, z)
}
```

### Type inference

When declaring a variable without specifying an explicit type (either by using the `:=` syntax or `var = expression` syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:  

```go
var i int
j := i // j is an int
```

But when the right hand side contains an untyped numeric constant, the new variable may be an `int`, `float64`, or `complex128` depending on the precision of the constant:

```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

Try changing the initial value of `v` in the example code and observe how its type is affected.

```go
package main

import "fmt"

func main() {
    v := 42 // change me!
    // v := 42.0    // float64
    fmt.Printf("v is of type %T\n", v)
}
```

### Constants

Constants are declared like variables, but with the `const` keyword.  
Constants can be character, string, boolean, or numeric values.  
Constants cannot be declared using the `:=` syntax.  

```go
package main

import "fmt"

const Pi = 3.14

func main() {
    const World = "世界"
    fmt.Println("Hello", World)     // Hello 世界
    fmt.Println("Happy", Pi, "Day") // Happy 3.14 Day

    const Truth = true
    fmt.Println("Go rules?", Truth) // Go rules? true
}
```

### Numeric Constants

Numeric constants are high-precision *values*.  
An untyped constant takes the type needed by its context.  
Try printing `needInt(Big)` too.  
(An `int` can store at maximum a 64-bit integer, and sometimes less.)  

```go
package main

import "fmt"

const (
    // Create a huge number by shifting a 1 bit left 100 places.
    // In other words, the binary number that is 1 followed by 100 zeroes.
    Big = 1 << 100
    // Shift it right again 99 places, so we end up with 1<<1, or 2.
    Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    fmt.Println(needInt(Small))     // 21
    fmt.Println(needFloat(Small))   // 0.2
    fmt.Println(needFloat(Big))     // 1.2676506002282295e+29
    fmt.Println(needInt(Big))       // constant 1267650600228229401496703205376 overflows int
}
```

## Flow control statements: for, if, else, switch and defer

Learn how to control the flow of your code with conditionals, loops, switches and defers.

### For loop

Go has only one looping construct, the `for` loop.

The basic for loop has three components separated by semicolons:

- the init statement: executed before the first iteration
- the condition expression: evaluated before every iteration
- the post statement: executed at the end of every iteration

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the `for` statement.

The loop will stop iterating once the boolean condition evaluates to `false`.

Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the `for` statement and the braces `{ }` are always required.

```go
package main

import "fmt"

func main() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)    // 45
}
```

The init and post statements are optional.

```go
func main() {
    sum := 1
    for ; sum < 1000; {
        sum += sum
    }
    fmt.Println(sum)    // 1024
}
```

### For is Go's "while"

At that point you can drop the semicolons: C's `while` is spelled `for` in Go.

```go
package main

import "fmt"

func main() {
    sum := 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)
}
```

### Forever (infinite loop)

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

```go
package main

func main() {
    for {

    }
}
```

### If condition

Go's `if` statements are like its `for` loops; the expression need not be surrounded by parentheses `( )` but the braces `{ }` are required.

```go
package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func main() {
    fmt.Println(sqrt(2), sqrt(-4))  // 1.4142135623730951 2i
}
```

### If with a short statement

Like `for`, the `if` statement can start with a short statement to execute before the condition.  
Variables declared by the statement are only in scope until the end of the `if`.  
(Try using `v` in the last return statement.)

```go
package main

import (
    "fmt"
    "math"
)

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
    // return v    // undefined: v
}

func main() {
    fmt.Println(
        pow(3, 2, 10),  // 9
        pow(3, 3, 20),  // 20
    )
}
```

### If and else

Variables declared inside an `if` short statement are also available inside any of the `else` blocks.  
(Both calls to `pow` return their results before the call to `fmt.Println` in `main` begins.)

```go
package main

import (
    "fmt"
    "math"
)

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)    // 27 >= 20
    }
    // can't use v here, though
    return lim
}

func main() {
    fmt.Println(
        pow(3, 2, 10),  // 9
        pow(3, 3, 20),  // 20
    )
}
```

### Exercise: Loops and Functions

As a way to play with functions and loops, let's implement a square root function: given a number x, we want to find the number z for which z² is most nearly x.

Computers typically compute the square root of x using a loop. Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:

`z -= (z*z - x) / (2*z)`

Repeating this adjustment makes the guess better and better until we reach an answer that is as close to the actual square root as can be.

Implement this in the `func Sqrt` provided. A decent starting guess for z is 1, no matter what the input. To begin with, repeat the calculation 10 times and print each z along the way. See how close you get to the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves.

Hint: To declare and initialize a floating point value, give it floating point syntax or use a conversion:

```go
z := 1.0
z := float64(1)
```

Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small amount). See if that's more or fewer than 10 iterations. Try other initial guesses for z, like x, or x/2. How close are your function's results to the [math.Sqrt](https://golang.org/pkg/math/#Sqrt) in the standard library?

```go
package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    /* function to be implemented */
    z := 1.0    // guess
    // z := (x/2.0)    // guess 2 -> works
    // z := x    // guess 3 -> works
    for i := 0; i < 10; i++ {
        z -= (z*z - x) / (2*z)
        fmt.Printf("Sqrt: %v",z)
        fmt.Printf("\tDifference: %v\n", z - math.Sqrt(2))    // Difference from actual Sqrt
        if diff := z - math.Sqrt(2); diff < 1.0e-4 {
            break
        }
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(3))
    fmt.Println(Sqrt(5))
}
```

Output:

```go

Sqrt: 1.5    Difference: 0.08578643762690485
Sqrt: 1.4166666666666667    Difference: 0.002453104293571595
Sqrt: 1.4142156862745099    Difference: 2.1239014147411694e-06
1.4142156862745099
Sqrt: 2    Difference: 0.5857864376269049
Sqrt: 1.75    Difference: 0.33578643762690485
Sqrt: 1.7321428571428572    Difference: 0.31792929476976206
Sqrt: 1.7320508100147276    Difference: 0.31783724764163246
Sqrt: 1.7320508075688772    Difference: 0.31783724519578205
Sqrt: 1.7320508075688774    Difference: 0.31783724519578227
Sqrt: 1.7320508075688772    Difference: 0.31783724519578205
Sqrt: 1.7320508075688774    Difference: 0.31783724519578227
Sqrt: 1.7320508075688772    Difference: 0.31783724519578205
Sqrt: 1.7320508075688774    Difference: 0.31783724519578227
1.7320508075688774
Sqrt: 3    Difference: 1.5857864376269049
Sqrt: 2.3333333333333335    Difference: 0.9191197709602383
Sqrt: 2.238095238095238    Difference: 0.823881675722143
Sqrt: 2.2360688956433634    Difference: 0.8218553332702683
Sqrt: 2.236067977499978    Difference: 0.821854415126883
Sqrt: 2.23606797749979    Difference: 0.8218544151266947
Sqrt: 2.23606797749979    Difference: 0.8218544151266947
Sqrt: 2.23606797749979    Difference: 0.8218544151266947
Sqrt: 2.23606797749979    Difference: 0.8218544151266947
Sqrt: 2.23606797749979    Difference: 0.8218544151266947
2.23606797749979
```

**Note**: If you are interested in the details of the algorithm, the z² − x above is how far away z² is from where it needs to be (x), and the division by 2z is the derivative of z², to scale how much we adjust z by how quickly z² is changing. This general approach is called [Newton's method](https://en.wikipedia.org/wiki/Newton%27s_method). It works well for many functions but especially well for square root.

More info on Newton's method:  

- [MIT 18.335J - Introduction to Numerical Methods](https://ocw.mit.edu/courses/mathematics/18-335j-introduction-to-numerical-methods-spring-2019/index.htm)
- SICP 1.1.7

### Switch

A `switch` statement is a shorter way to write a sequence of `if - else` statements. It runs the first case whose value is equal to the condition expression.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the `break` statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Print("Go runs on ")
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        // freebsd, openbsd,
        // plan9, windows...
        fmt.Printf("%s.\n", os)
    }
}
```

### Switch evaluation order

Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

(For example,

```go
switch i {
case 0:
case f():
}
```

does not call `f` if `i==0`.)

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("When's Saturday?")
    today := time.Now().Weekday()
    switch time.Saturday {
    case today + 0:
        fmt.Println("Today.")
    case today + 1:
        fmt.Println("Tomorrow.")
    case today + 2:
        fmt.Println("In two days.")
    default:
        fmt.Println("Too far away.")
    }
}
```

### Switch with no expression

Switch without an expression is the same as `switch true`.

> What is a switch expression?
> The case statements are excuted when they match the switch expression value. In absence of value, the case statements are compared to `true` i.e. any condition can be used in the case statements instead of writing a long `if-else` chain

This construct can be a clean way to write long if-then-else chains.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
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
```

[Go wiki Switch](https://github.com/golang/go/wiki/Switch)

### Defer

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```go
package main

import "fmt"

func main() {
    defer fmt.Println("world")

    fmt.Println("hello")
}
```

### Stacking defers

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

To learn more about defer statements read this [blog post](https://blog.golang.org/defer-panic-and-recover).

```go
package main

import "fmt"

func main() {
    fmt.Println("counting")

    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }

    fmt.Println("done")
}
```

Output:

```go
counting
done
9
8
7
6
5
4
3
2
1
0
```

## More types: structs, slices, and maps

Learn how to define types based on existing ones: this lesson covers structs, arrays, slices, and maps.

### Pointers

Go has pointers. A pointer holds the memory address of a value.

The type `*T` is a pointer to a `T` value. Its zero value is `nil`.

`var p *int`

The & operator generates a pointer to its operand.

```go
i := 42
p = &i
```

The `*` operator denotes the pointer's underlying value.

```go
fmt.Println(*p) // read (value of) i through the pointer p
*p = 21         // set (value of) i through the pointer p
```

This is known as "dereferencing" or "indirecting".

Unlike C, Go has **no pointer arithmetic**.

```go
package main

import "fmt"

func main() {
    i, j := 42, 2701

    p := &i         // point to i
    fmt.Println(p)  // pointer address
    fmt.Println(*p) // read i through the pointer
    *p = 21         // set i through the pointer
    fmt.Println(i)  // see the new value of i

    p = &j         // point to j
    fmt.Println(p) // pointer address
    *p = *p / 37   // divide j through the pointer
    fmt.Println(j) // see the new value of j
}
```

Output:

```go
0x40e020
42
21
0x40e024
73
```

### Structs

A `struct` is a collection of fields.

```go
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    fmt.Println(Vertex{1, 2})   // {1 2}
}
```

### Struct Fields

Struct fields are accessed using a dot.

```go
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)    // 4
}
```

### Pointers to structs

Struct fields can be accessed through a struct pointer.

To access the field `X` of a struct when we have the struct pointer `p` we could write `(*p).X`. However, that notation is cumbersome, so the language permits us instead to write just `p.X`, without the explicit dereference.

```go
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    p := &v         // address of v
    p.X = 1e9       // p.X is works instead of (*p).X, but NOT *p.X
    fmt.Println(v)  // {1000000000 2}
}
```

### Struct Literals

A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the `Name:` syntax. (And the order of named fields is irrelevant.)

The special prefix `&` returns a pointer to the struct value.

```go
package main

import "fmt"

type Vertex struct {
    X, Y int
}

var (
    v1 = Vertex{1, 2}  // has type Vertex
    v2 = Vertex{X: 1}  // Y:0 is implicit
    v3 = Vertex{}      // X:0 and Y:0
    p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
    fmt.Println(v1, p, v2, v3)  // {1 2} &{1 2} {1 0} {0 0}
}
```

### Arrays

The type `[n]T` is an array of `n` values of type `T`.

The expression  
`var a [10]int`  
declares a variable a as an array of ten integers.

An array's length is part of its type, **so arrays cannot be resized**. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.

```go
package main

import "fmt"

func main() {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1]) // Hello World
    fmt.Println(a)          // [Hello World]

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)     // [2 3 5 7 11 13]
}
```

### Slices

An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type `[]T` is a slice with elements of type `T`.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:  
`a[low : high]`  
This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of `a`:  
`a[1:4]`

```go
package main

import "fmt"

func main() {
    primes := [6]int{2, 3, 5, 7, 11, 13}

    var s []int = primes[1:4]
    fmt.Println(s)  // [3 5 7]
}
```

### Slices are like references to arrays

A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

**Other slices that share the same underlying array will see those changes.**

```go
package main

import "fmt"

func main() {
    names := [4]string{
        "John",
        "Paul",
        "George",
        "Ringo",
    }
    fmt.Println(names)  // [John Paul George Ringo]

    a := names[0:2]
    b := names[1:3]
    fmt.Println(a, b)   // [John Paul] [Paul George]

    b[0] = "XXX"
    fmt.Println(a, b)   // [John XXX] [XXX George]
    fmt.Println(names)  // [John XXX George Ringo]
}
```

### Slice literals

A slice literal is like an array literal without the length.

This is an array literal:  
`[3]bool{true, true, false}`  
And this creates the same array as above, then builds a slice that references it:  
`[]bool{true, true, false}`

```go
package main

import "fmt"

func main() {
    q := []int{2, 3, 5, 7, 11, 13}
    fmt.Println(q)  // [2 3 5 7 11 13]

    r := []bool{true, false, true, true, false, true}
    fmt.Println(r)  // [true false true true false true]

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
    fmt.Println(s)  // [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
}
```

### Slice defaults

When slicing an existing array, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

For the array defined as  
`var a [10]int`  
these slice expressions are equivalent:

```go
a[0:10]
a[:10]
a[0:]
a[:]
```

```go
package main

import "fmt"

func main() {
    s := []int{2, 3, 5, 7, 11, 13}

    fmt.Println(s[1:])  // [3 5 7 11 13]

    s = s[1:4]
    fmt.Println(s)      // [3 5 7]

    s = s[:2]
    fmt.Println(s)      // [3 5]

    s = s[1:]
    fmt.Println(s)      // [5]
}
```

### Slice length and capacity

A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, **_counting from the first element in the slice_**.

The length and capacity of a slice `s` can be obtained using the expressions `len(s)` and `cap(s)`.

You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

```go
package main

import "fmt"

func main() {
    s := []int{2, 3, 5, 7, 11, 13}
    printSlice(s)

    // Slice the slice to give it zero length.
    s = s[:0]
    printSlice(s)

    // Extend its length.
    s = s[:6]
    printSlice(s)

    // Limit its length
    s = s[:4]
    printSlice(s)

    // Drop its first two values. (reduce capacity)
    s = s[2:]
    printSlice(s)

    // ERROR -> panic: runtime error: slice bounds out of range [:7] with capacity 4
    s = s[:6]
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

Output:

```go
len=6 cap=6 [2 3 5 7 11 13]
len=0 cap=6 []
len=6 cap=6 [2 3 5 7 11 13]
len=4 cap=6 [2 3 5 7]
len=2 cap=4 [5 7]
panic: runtime error: slice bounds out of range [:6] with capacity 4

goroutine 1 [running]:
main.main()
    /tmp/sandbox829483878/prog.go:25 +0x140
```

### Nil Slices

The zero value of a slice is `nil`.

A nil slice has a length and capacity of 0 and has no underlying array.

```go
package main

import "fmt"

func main() {
    var s []int
    fmt.Println(s, len(s), cap(s))  // [] 0 0
    if s == nil {
        fmt.Println("nil!")         // nil!
    }
}
```

### Creating a slice with make

Slices can be created with the built-in `make` function; this is how you create **_dynamically-sized arrays_**.

The make function allocates a zeroed array and returns a slice that refers to that array:

```go
a := make([]int, 5)  // len(a)=5
```

To specify a capacity, pass a third argument to `make`:

```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

```go
package main

import "fmt"

func main() {
    a := make([]int, 5)
    printSlice("a", a)  // a len=5 cap=5 [0 0 0 0 0]

    b := make([]int, 0, 5)
    printSlice("b", b)  // b len=0 cap=5 []

    c := b[:2]
    printSlice("c", c)  // c len=2 cap=5 [0 0]

    d := c[2:5]
    printSlice("d", d)  // d len=3 cap=3 [0 0 0]
}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n",
        s, len(x), cap(x), x)
}
```

### Slices of slices

Slices can contain any type, including other slices.

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    // Create a tic-tac-toe board.
    board := [][]string{
        []string{"_", "_", "_"},
        []string{"_", "_", "_"},
        []string{"_", "_", "_"},
    }

    // The players take turns.
    board[0][0] = "X"
    board[2][2] = "O"
    board[1][2] = "X"
    board[1][0] = "O"
    board[0][2] = "X"

    for i := 0; i < len(board); i++ {
        fmt.Printf("%s\n", strings.Join(board[i], " "))
    }
}

```

Output:

```go
X _ X
O _ X
_ _ O
```

### Appending to a slice

It is common to append new elements to a slice, and so Go provides a built-in `append` function. The [documentation](https://golang.org/pkg/builtin/#append) of the built-in package describes `append`.  
`func append(s []T, vs ...T) []T`  
The first parameter `s` of `append` is a slice of type `T`, and the rest are `T` values to append to the slice.

The resulting value of `append` is a slice containing all the elements of the original slice plus the provided values.

If the backing array of `s` is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

(To learn more about slices, read the [Slices: usage and internals](https://blog.golang.org/slices-intro) article.)

```go
package main

import "fmt"

func main() {
    var s []int
    printSlice(s)       // len=0 cap=0 []

    // append works on nil slices.
    s = append(s, 0)
    printSlice(s)       // len=1 cap=1 [0]

    // The slice grows as needed.
    s = append(s, 1)
    printSlice(s)       // len=2 cap=2 [0 1]

    // We can add more than one element at a time.
    s = append(s, 2, 3, 4)
    printSlice(s)       // len=5 cap=6 [0 1 2 3 4]
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

### Range

The `range` form of the f`or loop iterates over a slice or map.(this is like a forEach loop)

When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }
}
```

Output:

```text
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128
```

### Range continued

You can skip the index or value by assigning to `_`.

```go
for i, _ := range pow
for _, value := range pow
```

If you only want the index, you can omit the second variable.

```go
for i := range pow
```

```go
package main

import "fmt"

func main() {
    pow := make([]int, 10)
    for i := range pow {
        pow[i] = 1 << uint(i) // == 2**i
    }
    for _, value := range pow {
        fmt.Printf("%d\n", value)
    }
}
```

Output:

```text
1
2
4
8
16
32
64
128
256
512
```

Article:
[4 basic range loop (for-each) patterns](https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/)

### Exercise: Slices

Implement `Pic`. It should return a slice of length `dy`, each element of which is a slice of `dx` 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include `(x+y)/2`, `x*y`, `x^y`, `(x^y)*(x^y)`, and `x*x+y*y` .

(You need to use a loop to allocate each `[]uint8` inside the `[][]uint8`.)

(Use `uint8(intValue)` to convert between types.)

[Explanation of the Exercise](https://stackoverflow.com/questions/25459474/go-tour-slices-exercise-logic)

```go
package main

import (
    "golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
    pic := make([][]uint8, dy)
    /* this works too
    for i:=0; i<len(pic); i++ {
       pic[i] = make([]uint8, dx)
       for j:=0; j<len(pic[i]); j++ {
           pic[i][j] = uint8(i*j)
       }
    }
     */
    for y:=range pic {
        pic[y] = make([]uint8, dx)
        for x:=range pic[y] {
            pic[y][x] = uint8((x ^ y) * (x ^ y))
        }
    }
    return pic
}

func main() {
    pic.Show(Pic)
}
```

Output:

![Output 1](./img/slices-pattern.png "(x ^ y) * (x ^ y)")

Use [this site](https://codebeautify.org/base64-to-image-converter) to covert base-64 text to image if needed.

Documentation link of the [pic](https://pkg.go.dev/golang.org/x/tour/pic?tab=doc) package that generates the above image. Code available [here](https://github.com/golang/tour/tree/0608babe047d/pic)  

[stackoverflow: iterating over a 2D slice in Go](https://stackoverflow.com/questions/37668224/iterating-over-over-a-2d-slice-in-go)  

### Maps

A map maps keys to values.

The zero value of a map is `nil`. A `nil` map has no keys, nor can keys be added.

The `make` function returns a map of the given type, initialized and ready for use.

```go
package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }
    fmt.Println(m["Bell Labs"])     // {40.68433 -74.39967}
}
```

### Map Literals

Map literals are like struct literals, but the keys are required.

```go
package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    "Bell Labs": Vertex{
        40.68433, -74.39967,
    },
    "Google": Vertex{
        37.42202, -122.08408,
    },
}

func main() {
    fmt.Println(m)  // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
}
```

### Map literals continued

If the top-level type is just a type name, you can omit it from the elements of the literal.

```go
package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967},
    "Google":    {37.42202, -122.08408},
}

func main() {
    fmt.Println(m)  // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
}
```

### Mutating Maps

Insert or update an element in map `m`:  
`m[key] = elem`

Retrieve an element:  
`elem = m[key]`

Delete an element:  
`delete(m, key)`

Test that a key is present with a two-value assignment:  
`elem, ok = m[key]`

If `key` is in `m`, `ok` is `true`. If not, `ok` is `false`.

If `key` is not in the map, then `elem` is the zero value for the map's element type.

**Note**: If `elem` or `ok` have not yet been declared you could use a short declaration form:  
`elem, ok := m[key]`

```go
package main

import "fmt"

func main() {
    m := make(map[string]int)

    m["Answer"] = 42
    fmt.Println("The value:", m["Answer"])          // The value: 42

    m["Answer"] = 48
    fmt.Println("The value:", m["Answer"])          // The value: 48

    delete(m, "Answer")
    fmt.Println("The value:", m["Answer"])          // The value: 0

    v, ok := m["Answer"]
    fmt.Println("The value:", v, "Present?", ok)    // The value: 0 Present? false
}
```

### Exercise: Maps

Implement `WordCount`. It should return a map of the counts of each “word” in the string `s`. The `wc.Test` function runs a test suite against the provided function and prints success or failure.

You might find [strings.Fields](https://golang.org/pkg/strings/#Fields) helpful.

```go
package main

import (
    "golang.org/x/tour/wc"
    "strings"
    // "fmt"
)

func WordCount(s string) map[string]int {
    words := strings.Fields(s)
    m := make(map[string]int)
    for _, word := range words {
        // fmt.Println(word)
        v, ok := m[word]
        if ok {
            m[word] = v+1
        } else {
            m[word] = 1
        }
    }
    // return map[string]int{"x": 1}
    return m
}

func main() {
    wc.Test(WordCount)
}
```

Output:

```text
PASS
 f("I am learning Go!") =
  map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
PASS
 f("The quick brown fox jumped over the lazy dog.") =
  map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}
PASS
 f("I ate a donut. Then I ate another donut.") =
  map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
PASS
 f("A man a plan a canal panama.") =
  map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}
```

### Function values

Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values.

```go
package main

import (
    "fmt"
    "math"
)

func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func main() {
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(hypot(5, 12))       // 13

    fmt.Println(compute(hypot))     // 5
    fmt.Println(compute(math.Pow))  // 81
}
```

### Function closures

Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

For example, the `adder` function returns a closure. Each closure is bound to its own `sum` variable.

```go
package main

import "fmt"

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }
}
```

Output:

```text
0 0
1 -2
3 -6
6 -12
10 -20
15 -30
21 -42
28 -56
36 -72
45 -90
```

### Relevant Functional Programming Concepts

[Closure: Wikipedia](https://en.wikipedia.org/wiki/Closure_(computer_programming))  
[What is a Closure? - Stackoverflow](https://stackoverflow.com/questions/36636/what-is-a-closure)  
[Closure Because of What it Can Do or Because it Does - Stackoverflow](https://stackoverflow.com/questions/4103750/closure-because-of-what-it-can-do-or-because-it-does/4103834#4103834)  
[What is 'Currying'? - Stackoverflow](https://stackoverflow.com/questions/36314/what-is-currying)  
[Currying: Wikipedia](https://en.wikipedia.org/wiki/Currying)  
[Higher Order Functions and Currying: Geeksforgeeks](https://www.geeksforgeeks.org/higher-order-functions-currying/)  

### Exercise: Fibonacci closure

Let's have some fun with functions.

Implement a `fibonacci` function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

```go
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    f1, f2 := 0, 1
    return func() int {
        f := f1
        f1, f2 = f2, f1+f2
        return f
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}

```

Output:

```text
0
1
1
2
3
5
8
13
21
34
```


# Methods and interfaces

This lesson covers methods and interfaces, the constructs that define objects and their behavior.

# Concurrency

Go provides concurrency features as part of the core language.

This module goes over goroutines and channels, and how they are used to implement different concurrency patterns.
