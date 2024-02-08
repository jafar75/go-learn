package main

import (
	"fmt"
	"math"
	"os"
	_ "embed"
)

//go:embed logs1.txt
var embeddedLogFile string

type Logger interface {
	log(lg string)
}

type FileLogger struct {
	fileName string
}

type StdoutLogger struct {}

func (fileLogger FileLogger) log(lg string) {
	// open a file with filename and write lg into it
	err := os.WriteFile(fileLogger.fileName, []byte(lg), 0644);
	if err != nil {
		panic(err);
	}
}

func (stdoutLogger StdoutLogger) log (lg string) {
	fmt.Println("lg: ", lg);
}

// User
type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

type Vertex struct {
	X float64
	Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y);
}

func (v *Vertex) AbsByRef() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y);
}


// Closures, lexically scoped: Functions can access values that were
// in scope when defining the function
func scope() func() int{
    outer_var := 2
    foo := func() int { return outer_var}
    return foo
}

func another_scope() func() int{
	foo := scope();
	return foo;
    // won't compile because outer_var and foo not defined in this scope
    // outer_var = 444
    // return foo
}


// Closures
func outer() (func() int, int) {
    outer_var := 2
    inner := func() int {
        outer_var += 99 // outer_var from outer scope is mutated.
        return outer_var
    }
    inner()
    return inner, outer_var // return inner func and mutated outer_var 101
}

func adderWithOffset(offset int, args ...int) int {
	total := offset;
	for _, v := range args {
		total = total + v;
	}
	return total;
}

func main() {
	fmt.Println("foo(): ", another_scope()());
    // assign a function to a name
    add := func(a, b int) int {
        return a + b
    }
    // use the name to call the function
    fmt.Println(add(3, 4))

	fmt.Println("1 arg: ", adderWithOffset(12, 1));
	fmt.Println("1 4 3 1 arg: ", adderWithOffset(12, 1, 4, 3, 1));

	x := [3]int{1, 2, 3};
	s := x[:];
	s = append(s, 5);
	t := []int{};
	t = append(t, 4);
	t = append(t,s...);
	fmt.Println("concat: ", t);

	vertices := []Vertex{{1, 2}, {3, 4}};
	fmt.Println("vertices: ", vertices);
	fmt.Println("abs by value: ", vertices[0].Abs());
	fmt.Println("abs by ref: ", (&vertices[0]).AbsByRef());

	doLog := func(logger Logger, lg string) {
		logger.log(lg);
	}
	stdLogger := StdoutLogger{};
	fileLogger := FileLogger{"logs1.txt"};

	doLog(stdLogger, "log for std");
	doLog(fileLogger, "log for file");



	fmt.Println("contents of logfile: ", embeddedLogFile);
}

// func main() {
// 	// var foo int // declaration without initialization
// 	// var foo int = 42 // declaration with initialization
// 	// var foo, bar int = 42, 1302 // declare and init multiple vars at once
// 	// var foo = 42 // type omitted, will be inferred
// 	// foo := 42 // shorthand, only in func bodies, omit var keyword, type is always implicit
// 	const constant = "This is a constant"

// 	// iota can be used for incrementing numbers, starting from 0
// 	const (
// 	    _ = iota
// 	    a
// 	    b
// 	    c = 1 << iota
// 	    d
// 	)
// 	fmt.Println(a, b) // 1 2 (0 is skipped)
// 	fmt.Println(c, d) // 8 16 (2^3, 2^4)
// }