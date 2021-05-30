# Golang tutorial notes

## Commands used in the project

Enable dependency tracking:
`go mod init example.com/hello`

Add module requirements and sums:
`go mod tidy`

Run a file:
`go run hello.go`

Run the main package:
`go run .`

Import local module:
`go mod edit -replace=example.com/greetings=../greetings`

Run tests:
`go test`

Compile the packages, along with their dependencies, but don't install the results:
`go build`

Discover the Go install path, where the `go` command will install the current package:
`go list -f '{{.Target}}'`

## Notes

1. In Go, the `:=` operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type).
2. `Sprintf` substitutes the name parameter's value for the `%v` format verb.
3. Common error handling in Go: Return an error as a value so the caller can check for it.
4. A slice is like an array, except that its size changes dynamically as you add and remove items. The slice is one of Go's most useful types.
5. Go executes `init` functions automatically at program startup, after global variables have been initialized.
6. When declaring a slice, you omit its size in the brackets, like this: `[]string`. This tells Go that the size of the array underlying the slice can be dynamically changed.
7. You initialize a map with the following syntax: `make(map[key-type]value-type)`.
8. `range` returns two values: the index of the current item in the loop and a copy of the item's value.
9. Test function names have the form Test*Name*, where _Name_ says something about the specific test.
10. Test functions take a pointer to the `testing` package's [`testing.T`](https://pkg.go.dev/testing/#T) type as a parameter.
11. By convention, the package name is the same as the last element of the import path.
12. In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package. When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.
13. When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
14. A function can return any number of results.
15. Go's return values may be named. If so, they are treated as variables defined at the top of the function.
    T`hese names should be used to document the meaning of the return values.
    A return statement without arguments returns the named return values. This is known as a "naked" return.
    Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
16. The var statement declares a list of variables; as in function argument lists, the type is last.
    A var statement can be at package or function level.
17. A var declaration can include initializers, one per variable.
    If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
18. Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with implicit type.
    Outside a function, every statement begins with a keyword (`var`, `func`, and so on) and so the `:=` construct is not available.
19. Go's basic types are:

```go
bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32; represents a Unicode code point
float32 float64
complex64 complex128
```

20. Variable declarations may be "factored" into blocks, as with import statements.
21. The `int`, `uint`, and `uintptr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
    When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
22. Variables declared without an explicit initial value are given their zero value. The zero value is:
    - `0` for numeric types,
    - `false` for the boolean type, and
    - `""` (the empty string) for strings.
23. The expression `T(v)` converts the value `v` to the type `T`. E.g. `i := 42; f := float64(i)`
24. In Go assignment between items of different type requires an explicit conversion.
25. Constants are declared like variables, but with the const keyword. Constants cannot be declared using the := syntax.
26. An untyped constant takes the type needed by its context.
27. Go has only one looping construct, the for loop.
    The init and post statements are optional. At that point you can drop the semicolons: C's `while` is spelled `for` in Go, e.g. `for i < 10 {}`
28. If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
29. Like `for`, the `if` statement can start with a short statement to execute before the condition.
    Variables declared inside an `if` short statement are also available inside any of the `else` blocks.
30. Go's `switch` is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the `break` statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

```go
switch i {
case 0:
case f():
}
```

31. Switch without a condition is the same as `switch true`.
    This construct can be a clean way to write long if-then-else chains.
32. A defer statement defers the execution of a function until the surrounding function returns.
    The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
33. Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
34. The type `*T` is a pointer to a `T` value. Its zero value is `nil`. E.g. `var p *int`
    The `&` operator generates a pointer to its operand. E.g. `i := 42; p = &i;`
    The `*` operator denotes the pointer's underlying value. E.g. `fmt.Println(*p); *p = 43; fmt.Println(i)`
35. A `struct` is a collection of fields.

```go
type Vertex struct {
	X int
	Y int
}
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)
func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Println(v.X)
}
// anonymous structs are a thing
s := []struct {
    i int
    b bool
}{
    {2, true},
}

m := make(map[string]Vertex)
m["Bell Labs"] = Vertex{
    40.68433, -74.39967,
}
```

36. Go has no pointer arithmetic.
37. Struct fields can be accessed through a struct pointer.
    To access the field `X` of a struct when we have the struct pointer p we could write `(*p).X`. However, that notation is cumbersome, so the language permits us instead to write just `p.X`, without the explicit dereference.
38. The type `[n]T` is an array of `n` values of type `T`.
39. The type `[]T` is a slice with elements of type `T`.
40. We can create a slice from an array `a[low : high]`, which will select from index `low` until `high-1`.
41. Slices are like references to arrays.
    A slice does not store any data, it just describes a section of an underlying array.
    Changing the elements of a slice modifies the corresponding elements of its underlying array.
    Other slices that share the same underlying array will see those changes.
42. A slice literal is like an array literal without the length.
43. When slicing, you may omit the high or low bounds to use their defaults instead.
44. A slice has both a _length_ and a _capacity_.
    The length of a slice is the number of elements it contains.
    The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
    The length and capacity of a slice `s` can be obtained using the expressions `len(s)` and `cap(s)`.
    You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
45. The zero value of a slice is `nil`.
    A `nil` slice has a length and capacity of 0 and has no underlying array.
46. Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
    `b := make([]int, 0, 5) // len(b)=0, cap(b)=5`. Capacity argument is optional. If not provided, array is filled with zero value.
47. Use the built-in `append` function to insert new items into a slice.
    If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.
48. The `range` form of the `for` loop iterates over a slice or map.
    When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
    If you only want the index, you can omit the second variable.
49. The zero value of a `map` is `nil`. A `nil` map has no keys, nor can keys be added.
    The `make` function returns a map of the given type, initialized and ready for use.
50. - Insert or update an element in map m: `m[key] = elem`
    - Retrieve an element: `elem = m[key]`
    - Delete an element: `delete(m, key)`
    - Test that a key is present with a two-value assignment: `elem, ok = m[key]`
      If `key` is in m, `ok` is true. If not, `ok` is `false`.
      If `key` is not in the map, then `elem` is the zero value for the map's element type.
51. If `elem` or `ok` have not yet been declared you could use a short declaration form: `elem, ok := m[key]`
52. Functions are values too. They can be passed around just like other values.
53. Go functions may be closures. A closure is a function value that references variables from outside its body.
    The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
54. Go does not have classes. However, you can define methods on types.
    A method is a function with a special receiver argument.
    The receiver appears in its own argument list between the `func` keyword and the method name.
55. You can declare a method on non-struct types, too, e.g. `type MyFloat float64` and use `MyFloat` for method type.
56. You can only declare a method with a receiver whose type is defined in the same package as the method.
    You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as `int`).
57. Structs are passed by value to functions if not using a pointer.
58. You can declare methods with pointer receivers.
    This means the receiver type has the literal syntax `*T` for some type `T`. (Also, `T` cannot itself be a pointer such as `*int`.)
    Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
59. There are two reasons to use a pointer receiver.
    - The first is so that the method can modify the value that its receiver points to.
    - The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
      In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.
60. An `interface` type is defined as a set of method signatures.
    A value of interface type can hold any value that implements those methods.
61. A type implements an interface by implementing its methods.
    There is no explicit declaration of intent, no "implements" keyword.
    Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.
62. Under the hood, interface values can be thought of as a tuple of a value and a concrete type.
    Calling a method on an interface value executes the method of the same name on its underlying type.
63. If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
    In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver.
64. A nil interface value holds neither value nor concrete type.
65. The interface type that specifies zero methods is known as the empty interface: `interface{}`.
    An empty interface may hold values of any type. (Every type implements at least zero methods.)
    Empty interfaces are used by code that handles values of unknown type.
66. A _type assertion_ provides access to an interface value's underlying concrete value. `t := i.(T)`.
    This statement asserts that the interface value `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`.
    If `i` does not hold a `T`, the statement will trigger a panic.
67. A type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded. `t, ok := i.(T)`
    If `i` holds a `T`, then `t` will be the underlying value and `ok` will be `true`.
    If not, `ok` will be `false` and `t` will be the zero value of type `T`, and no panic occurs.
68. A _type switch_ is a construct that permits several type assertions in series.

```go
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
```

69. A `Stringer` is a type that can describe itself as a string. The `fmt` package (and many others) look for this interface to print values.
    The `Stringer` type has a single `String() string` method.
70. Go programs express error state with `error` values.
    The `error` type is a built-in interface similar to `fmt.Stringer`. `type error interface { Error() string }`
    Functions often return an `error` value, and calling code should handle errors by testing whether the error equals `nil`.
71. The `io` package specifies the `io.Reader` interface, which represents the read end of a stream of data.
    The `io.Reader` interface has a `Read` method: `func (T) Read(b []byte) (n int, err error)`
    `Read` populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an `io.EOF` error when the stream ends.
72. A _goroutine_ is a lightweight thread managed by the Go runtime.
    `go f(x, y, z)` starts a goroutine running `f(x, y, z)`.
    The evaluation of `f`, `x`, `y`, and `z` happens in the current goroutine and the execution of `f` happens in the new goroutine.
    Goroutines run in the same address space, so access to shared memory must be synchronized.
    The `sync` package provides useful primitives, although you won't need them much in Go as there are other primitives.
73. Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`.
    channels must be created before use.

```go
ch := make(chan int)
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v.
```

74. By default, sends and receives block until the other side is ready.
    This allows goroutines to synchronize without explicit locks or condition variables.
75. Channels can be buffered.
    Provide the buffer length as the second argument to make to initialize a buffered channel: `ch := make(chan int, 100)`
    Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
76. A sender can `close` a channel to indicate that no more values will be sent.
    Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: `i, ok := <-ch`.
    `ok` is `false` if there are no more values to receive and the channel is closed.
77. The loop `for i := range c` receives values from the channel repeatedly until it is closed.
78. Only the sender should close a channel, never the receiver.
    Sending on a closed channel will cause a panic.
79. Channels aren't like files; you don't usually need to close them.
    Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a `range` loop.
80. The `select` statement lets a goroutine wait on multiple communication operations.
    A `select` blocks until one of its cases can run, then it executes that case.
    It chooses one at random if multiple are ready.
81. The `default` case in a `select` is run if no other case is ready.
    Use a `default` case to try a send or receive without blocking.
82. Use a _mutex_ to make sure only one goroutine can access a variable at a time to avoid conflicts.
    Go's standard library provides mutual exclusion with `sync.Mutex` and its two methods: `Lock` and `Unlock`.
83. Use `defer` to make sure the mutex will be unlocked.

## Links

1. [Effective Go](https://golang.org/doc/effective_go.html#multiple-returns)
2. [Go maps in action](https://blog.golang.org/maps)
3. [The blank identifier](https://golang.org/doc/effective_go.html#blank)
4. [Keeping your modules compatible](https://blog.golang.org/module-compatibility)
5. [`testing.T`](https://pkg.go.dev/testing/#T)
6. [t parameter's Fatalf method](https://pkg.go.dev/testing/#T.Fatalf)
7. [Go format parameters](https://golang.org/pkg/fmt/)
8. [Managing dependencies](https://golang.org/doc/modules/managing-dependencies)
9. [Developing and publishing modules](https://golang.org/doc/modules/developing)
10. [Published Go packages](https://pkg.go.dev/)
11. [Tour of Go](https://tour.golang.org/welcome/1)
12. [Go's declaration syntax](https://blog.golang.org/gos-declaration-syntax)
13. [Go's defer statements](https://blog.golang.org/defer-panic-and-recover)
14. [`append` documentation](https://golang.org/pkg/builtin/#append)
15. [Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals)
16. [`strings.Fields`](https://golang.org/pkg/strings/#Fields)
17. [More on default(zero) values](https://stackoverflow.com/a/28625828)
18. [Interface values with nil underlying values](https://tour.golang.org/methods/12)
19. [Recursive behaviour with Error() method](https://stackoverflow.com/questions/27474907/why-would-a-call-to-fmt-sprinte-inside-the-error-method-result-in-an-infinit)
20. [`io.Reader` implementations](<https://cs.opensource.google/search?q=Read%5C(.*%5Cs%5C%5B%5C%5Dbyte%5C)&ss=go%2Fgo>)
21. [More on Go](https://tour.golang.org/concurrency/11)
    - [Go docs](https://golang.org/doc/)
    - [Organize and work with Go code](https://www.youtube.com/watch?v=XCsL89YtqCs)
    - [How to write Go code](https://golang.org/doc/code.html)
    - [std lib reference](https://golang.org/pkg/)
    - [Language spec](https://golang.org/ref/spec)
    - [Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)
    - [Go Concurrency Patterns (slides)](https://talks.golang.org/2012/concurrency.slide)
    - [Advanced Go Concurrency Patterns](https://www.youtube.com/watch?v=QDDwwePbDtw)
    - [Advanced Go Concurrency Patterns (slides)](https://talks.golang.org/2013/advconc.slide)
    - [Share Memory by Communicating](https://golang.org/doc/codewalk/sharemem/)
    - [A simple programming environment](https://vimeo.com/53221558)
    - [A simple programming environment (slides)](https://talks.golang.org/2012/simple.slide)
    - [Writing Web Applications](https://golang.org/doc/articles/wiki/)
    - [First Class Functions in Go](https://golang.org/doc/codewalk/functions/)
