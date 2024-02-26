# Go!

## What is it?

Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.

Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.

The language is often referred to as Golang because of its domain name, golang.org, but the proper name is Go.

![](img%5Cgoinaday1.png)

## Whut?

It’s a programming language, designed by Google, primarily motivated by their dislike of C++ but keeping some characteristics they liked in other languages:

- static typing and run-time efficiency (like C),

- readability and usability (like Python or JavaScript),

- high-performance networking and multiprocessing.

---

Statically typed is a programming language characteristic in which variable types are explicitly declared and thus are determined at compile time. 

This lets the compiler decide whether a given variable can perform the actions requested from it or not. *Static typing associates types with variables, not with values*.

## Why use Go?

> “Golang is very useful for writing light-weight microservices. We currently use it for generating APIs that interact with our front-end applications. If you want to build a small functional microservice quickly, then Golang is a great tool to use.”

> “The advantages of using Go over other similar coding languages is that it’s an easy language for developers to learn quickly, and it has several built-in features to assist in asynchronous development.2

> “If you’re looking to build a back end service that isn’t data-science heavy, use Golang. Otherwise, use Python.”

> “The simplicity of the language makes for easier code review and debugging, while the philosophy of leanness in libraries also eases debugging. Meanwhile, the inclusion of all dependencies in the compiled binary makes for simple deployments.”

---

There are other areas where Go excels. For example, there are no dependencies when running a compiled Go program. You don't have to worry if your users have Ruby or the JVM installed, and if so, what version. 

For this reason, Go is becoming increasingly popular as a language for command-line interface programs and other types of utility programs you need to distribute (e.g., a log collector).

Put plainly, learning Go is an efficient use of your time. You won't have to spend long hours learning or even mastering Go, and you'll end up with something practical from your effort.


# Part 1

## Go Fundamentals

We're going to be learning the Go Fundamentals in the Go Playground, which means we don't have to install Go just yet.

NOTE: We used the old Go playground and are updating the pictures, but the contents will be the same.

https://goplay.tools/ 

![](img%5Cgoinaday3.png)

---

## Hello Learner!

![](img%5Cgoinaday4.png)

[Click here to see the code](https://goplay.tools/snippet/UTfp9e_aAbd)

Click the `Run` button to execute the code.

> fmt is the format package

---

Hopefully, the code that we just executed is understandable. We've created a function and printed out a string with the built-in `Println` function. 

Saying that a language has a C-like syntax means that if you're used to any other C-like languages such as C, C++, Java, JavaScript and C#, then you're going to find Go familiar -- superficially, at least. For example, it means:
- `&&` is used as a boolean AND, 
- `==` is used to compare equality, 
- `{` and `}` start and end a scope, 
- and array indexes start at 0.

C-like syntax also tends to mean *semi-colon terminated* lines and parentheses around conditions. Go does away with both of these, though parentheses are still used to control precedence. 

In Go, the entry point to a program has to be a function called `main` within a `package main`.

We'll talk more about packages later. For now, while we focus on understanding the basics of Go, we'll always write our code within the `main` function.

## Imports


Go has a number of built-in functions, such as `Println`, which can be used without reference. We can't get very far though, without making use of Go's standard library and eventually using third-party libraries. In Go, the `import` keyword is used to declare the packages that are used by the code in the file.

![](img%5Cgoinaday5.png)

[Click here to see the code in the Go Playground](https://goplay.tools/snippet/fairSSz2PXW)


You've probably noticed we prefix the function name with the package, e.g., `fmt.Println`. 

This is different from many other languages. We'll learn more about packages in later chapters. For now, knowing how to import and use a package is a good start.

Go is strict about importing packages. It will **not** compile if you import a package but don't use it. 

Over time, you'll get used to it (it'll still be annoying though). Go is strict about this because unused imports can slow compilation; admittedly a problem most of us don't have to this degree.

Another thing to note is that Go's standard library is well documented. You can head over to https://golang.org/pkg/fmt/#Println to learn more about the `Println` function that we used. 

## Variables & Declarations

It'd be nice to begin and end our look at variables by saying you declare and assign to a variable by doing `x = 4`.

Unfortunately, things are more complicated in Go. We'll begin our conversation by looking at simple examples. Then, in the next chapter, we'll expand this when we look at creating and using structures. Still, it'll probably take some time before you truly feel comfortable with it.

You might be thinking *Woah! What can be so complicated about this?* 

Let's start looking at some examples:

![](img%5Cgoinaday6.png)

[Code here](https://goplay.tools/snippet/8oQwThFogpS)

The most explicit way to deal with variable declaration and assignment in Go is also the most verbose.

Go only has one keyword for declaring a variable - `var`. 

You'll notice that we can declare variables in multiple ways, and the first one infers the type from the assigned value, so 

`var a = "initial"`

Creates a variable named `a` infers the type as `string` from the value `"initial"`.

> NOTE : If you haven't worked with typed languages before, here's something you can do in Javascript:
>
> ```js
> let a = "initial"
> a = 1
> ```
>
> You CANNOT do this in Go. Once you have declared a variable as a type - even if the type has been inferred - you cannot give it another value that is of another type.

The last example is the one you will see most often used in Go, and is known as the short assignment operator `:=` :

```go
f := apple
```

You don't have to use the `var` keyword when using it.

### Variable Scope

It's important that you remember that `:=` is used to *declare* the variable as well as *assign* a value to it. Why? Because a variable can't be declared twice (not in the same scope anyway). If you try to run the following, you'll get an error.

![](img%5Cgoinaday7.png)

[Code here](https://goplay.tools/snippet/DxwHGSSZWpZ )

The compiler will complain with `no new variables on left side of :=`. This means that when we first declare a variable, we use `:=` but on subsequent assignment, we use the assignment operator `=`. This makes a lot of sense, but it can be tricky for your muscle memory to remember when to switch between the two.

If you read the error message closely, you'll notice that variables is plural. That's because Go lets you assign multiple variables (using either `=` or `:=`):

For now the last thing to know is that, like imports, Go **won't** let you have unused variables.

#### A Quick Note

The short assignment operator `:=` can only be used inside a function!

It’s semantically the same as

`var d int = 1`

Or

`var d = 1`

## For Loops

Go has only one looping construct, the `for` loop.
The basic `for` loop has three components separated by semicolons:
- **the init statement**: executed before the first iteration
- **the condition expression**: evaluated before every iteration
- **the post statement**: executed at the end of every iteration

![](img%5Cgoinaday8.png)

[Code here](https://goplay.tools/snippet/TeaO0Yxt95g )

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the `for` statement.

The loop will stop iterating once the boolean condition evaluates to false.

> Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces `{ }` are always required.

## Flow Control

Go's `if` statements are like its `for` loops; the expression need not be surrounded by parentheses `( )` but the braces `{ }` **are** required.

![](img%5Cgoinaday9.png)

[Code here](https://goplay.tools/snippet/UA6GTDP13YG )

## Arrays

If you come from Python, Ruby, Perl, JavaScript or PHP (and more), you're probably used to programming with dynamic arrays. These are arrays that resize themselves as data is added to them. 

In Go, like many other languages, arrays are *fixed*. Declaring an array requires that we specify the size, and once the size is specified, it cannot grow.

![](img%5Cgoinaday10.png)

[Code here](https://goplay.tools/snippet/NIMcD_lX8q0 )

**Block 1** - Here we create an array a that will hold exactly 5 integers. The type of elements and length are both part of the array’s type. By default an array is zero-valued, which for integers means 0s.

**Block 2** - We can set a value at an index using the `array[index] = value` syntax, and get a value with `array[index]`.

**Block 3** - The builtin `len` returns the length of an array.

**Block 4** - Use this syntax to declare and initialize an array in one line:
> ```go
> b := [5]int{1,2,3,4,5}
>```

**Block 5** - We can use `len` to get the length of the array. `range` can be used to iterate over it. We'll discuss `range` later.

## Slices



In Go, you rarely, if ever, use arrays directly. Instead, you use *slices*. A slice is a lightweight structure that wraps and represents a portion of an array. There are a few ways to create a slice, and we'll go over when to use which later on. The first is a slight variation on how we created an array:

```go
  scores1 := []int{1, 4, 293, 4, 9}
```
 Unlike the array declaration, our slice isn't declared with a length within the square brackets.

![](img%5Cgoinaday11.png)

[Code here](https://goplay.tools/snippet/xcMbUbQTx8B)

 To understand how arrays and slices are different, let's see another way to create a slice, using `make`. We use `make` instead of `new` because there's more to creating a slice than just allocating the memory (which is what `new` does). Specifically, we have to allocate the memory for the underlying array and also initialize the slice. 

 In this case, using make to create the slice creates and output for `scores2` of:

 `[0 0 0 0 0 0 0 0 0 0]`

 So we have a dynamic array with a capacity of 10, with ten integers already assigned as default values. This is known as the length of the slice (we initialize a slice with a length of 10 and a capacity of 10. The length is the size of the slice, the capacity is the size of the underlying array.). We can declare it to have zero length:

 ```go
 scores3 := make(([]int, 0, 10))
 ```

This creates a slice with a length of 0 but with a capacity of 10.

Appending to a slice of length 0 will set the first element, as the slice we mad had a length of zero (but can hold up to 10).

## Reslice & Append

How large can we resize a slice? Up to its capacity which, in this case, is 10. You might be thinking this doesn't actually solve the fixed-length issue of arrays. It turns out that `append` is pretty special. If the underlying array is full, it will create a new larger array and copy the values over (this is exactly how dynamic arrays work in PHP, Python, Ruby, JavaScript, ...). 

![](img%5Cgoinaday12.png)

[Code here](https://goplay.tools/snippet/KANWcaYws90)

This is why, in the example above that used `append`, we had to re-assign the value returned by `append` to our scores variable: `append` might have created a new value if the original had no more space.

**What will block 2 do?**

Here, the output is going to be `[0, 0, 0, 0, 0, 9332]`. Maybe you thought it would be `[9332, 0, 0, 0, 0]`? To a human, that might seem logical. To a compiler, you're telling it to append a value to a slice that already holds 5 values.

## Functions

Functions are very similar to other languages. There is a keyword `func`, the name of the function `plus`, and parameters `(a int, b int)` and the return type `int`:

```go
func plus(a int, b int) int {
  //..code here
}
```

![](img%5Cgoinaday13.png)

[Code here](https://goplay.tools/snippet/lqPNF6DBVgX )

You call a function exactly like you'd expect with `functionName(arguments)`.

### Functions Can Return Multiple Values

This is a good time to point out that functions can return multiple values in Go:

![](img%5Cgoinaday14.png)

When returning multiple values in Go, the return types must be enclosed in parenthesis `()`;

![](img%5Cgoinaday15.png)

### I only care about one return type

Sometimes, you only care about one of the return values. In these cases, you assign the other values to. This is more than a convention. `_` the *blank identifier* is special in that the return value isn't actually assigned. This lets you use `_` over and over again regardless of the returned type.

![](img%5Cgoinaday16.png)

[Code here](https://goplay.tools/snippet/MRJP8ntz47F)

---

# Lab #1

* Given an array of integers
  * Create a function that accepts the list
  * Returns the product (or the sum, I don’t care) of all the numbers

* Given a slice of strings like `[“1”, “3”, “77”, “hello”, “sunshine”]`
  * Create a function that accepts the slice
  * Returns the sum of all the EVEN integers  (`strconv.Atoi` may be needed here)
  * And returns the strings concatenated together

[Skeleton Code For Lab](https://goplay.tools/snippet/rTtHTqDSY3J)

[Go Cheatsheet](https://devhints.io/go)

---

# Part 2

## Structures, Methods & Pointers

### The Basics

Although Go doesn't do OO like you may be used to, you'll notice a lot of similarities between the definition of a *structure* and that of a *class*. A simple example is the following `Source` structure:

![](img%5Cgoinaday17.png)


Go isn't an object-oriented (OO) language like C++, Java, Ruby and C#. It doesn't have objects nor inheritance and thus, doesn't have the many concepts associated with OO such as polymorphism and overloading.

What Go does have are structures, which can be associated with methods. Go also supports a simple but effective form of composition. Overall, it results in simpler code, but there'll be occasions where you'll miss some of what OO has to offer. (It's worth pointing out that composition over inheritance is an old battle cry and Go is the first language I've used that takes a firm stand on the issue.)

We'll soon see how to add a method to this structure, much like you'd have methods as part of a class. Before we do that, we have to dive back into declarations.

## Declarations
When we first looked at variables and declarations, we looked only at built-in types, like integers and strings. Now that we're talking about structures, we need to know how wer use them.

The simplest way to create a value of our structure is below.

![](img%5Cgoinaday18.png)

[Code here](https://go.dev/play/p/xOQ8J-jzB9m)

Note: The trailing `,` in the above structure when we give it values on intialisation is required. Without it, the compiler will give an error. You'll appreciate the required consistency, especially if you've used a language or format that enforces the opposite.

## Declarations – Part 2

We don't have to set all or even any of the fields. Both of these are valid:

![](img%5Cgoinaday19.png)

[Code here](https://go.dev/play/p/R_DWHlrQCsd)

Just like unassigned variables have a zero value, so do fields.
Furthermore, you can skip the field name and rely on the order of the field declarations (though for the sake of clarity, you should only do this for structures with few fields):


### Passing By Value

What all of the above examples do is declare a variable `gas` and assign a value to it.

Many times though, we don't want a variable that is directly associated with our value but rather a variable that has a pointer to our value. A pointer is a memory address; it's the location of where to find the actual value. It's a level of indirection. Loosely, it's the difference between being at a house and having directions to the house.

> If you have done C, then pointers in Go *are NOT exactly the same*.

Why do we want a pointer to the value, rather than the actual value? It comes down to the way Go passes arguments to a function: as copies. 

Knowing this, what does the following print:

![](img%5Cgoinaday20.png)

The answer is 9000, not 19000. Why? Because Super made changes to a *copy* of our original source value and thus, changes made in Super weren't reflected in the caller. To make this work as you probably expect, we need to pass a pointer to our value.

## Pointers

![](img%5Cgoinaday21.png)

We made two changes. The first is the use of the `&` operator to get the address of our value (it's called the *address of* operator). Next, we changed the type of parameter `Super` expects. It used to expect a value of type `Source` but now expects an address of type `*Source`, where `*X` means pointer to value of type X. There's obviously some relation between the types `Source` and `*Source`, but they are two distinct types.

Note that we're still passing a copy of `gas`’s value to `Super` it just so happens that `gas`'s value has become an address. That copy is the same address as the original, which is what that indirection buys us. Think of it as copying the directions to a restaurant. What you have is a copy, but it still points to the same restaurant as the original.

# Pointers vs Values

As you write Go code, it's natural to ask yourself should this be a value, or a pointer to a value? 

1. **Use Pointers when you need to modify the original value**

If you want a function to be able to modify the original value of a variable, you should pass a pointer to that variable. This is because when you pass a value to a function in Go, a copy of the value is created, and modifications made to the copy do not affect the original.

```go
func modifyValue(v *int) {
    *v = *v * 2
}

func main() {
    x := 10
    modifyValue(&x)
    fmt.Println(x) // Output: 20
}

```

2. **Use Values for immutability and simplicity**

If your data is immutable or if you want to work with a copy of the data and NOT modify the original, then passing by value is appropriate.

```go
func calculateSquare(num int) int {
    return num * num
}

func main() {
    y := 5
    result := calculateSquare(y)
    fmt.Println(result) // Output: 25
    fmt.Println(y)      // Output: 5 (original value is unchanged)
}

```

Here is [an example](https://goplay.tools/snippet/wwt4TpEXYn1) of passing a value by Value and as a Pointer and how the value is or isn't updated.

## Constructors

Structures don't have constructors. Instead, you create a function that returns an instance of the desired type (like a factory):

![](img%5Cgoinaday24.png)

Our factory doesn't have to return a pointer; this is absolutely valid:

![](img%5Cgoinaday23.png)

## Methods

Go supports *methods* defined on struct types:

![](img%5Cgoinaday25.png)

The area method has a *receiver type* of `*rect`:

```go
func (r *rect) area() int {
  return r.width * r.height
}
```

Keyword `func`, followed the the receiver type `(r *rect)` which indicates that it is a *method on the `rect` struct*, the name of the method `area()` and the return type `int`.

Methods can be defined for either pointer or value receiver types. The `perim` method is an example of Value type receiver.

Here we call the 2 methods defined for our struct.

![](img%5Cgoinaday26.png)

Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.

[Code here](https://go.dev/play/p/Ep45Ns3yTlN)

## Quick Lab

Create a  __struct__  for a Circle that has  __diameter__  and  __radius__  as properties of type `int`

Implement  __methods__  to return the  __circumference__  and the  __area__  of the circle

## STRUCTS & METHODS - Extended Lab

Create a  __struct__  for a Circle that has  __diameter__  and  __radius__  as properties of type int​

Implement  __methods__  to return the  __circumference__  and the  __area__  of the circle​

Create a  __function__  that accepts a slice of  _ints_  that represent the radius of the Circle, and return a slice of Circle, ordered by the circumference of the Circle​

Example input might be: `[10, 5, 3,100,99, 12, 9, 1, 4, 35, 21]​`

[Skeleton code can be found here](https://goplay.tools/snippet/4e5CL43RNSY) 

---

# Part 3

## Maps, Ranges & Errors

## Maps

Maps in Go are what other languages call hashtables or dictionaries. They work as you expect: you define a key and value, and can get, set and delete values from it.
Maps, like slices, can be created with the `make` function. 

![](img%5Cgoinaday27.png)

[Code here](https://go.dev/play/p/yRUAZAjpCr7)

To get the number of keys, we use len. To remove a value based on its key, we use delete

Maps grow dynamically. However, we can supply a second argument to make to set an initial size., `lookup := make(map[string]int, 100)`. If you have some idea of how many keys your map will have, defining an initial size can help with performance.

You can also declare maps *without* using `make`:

```go
lookup := map[string]int{}
```


## Range

We can iterate over a map using a for loop combined with the `range` keyword:

![](img%5Cgoinaday28.png)

[Code here](https://go.dev/play/p/cvwftQosOvM)

So we have the `for` keyword, then the `key` (or index) and `value` variables that wil be assigned to on each iteration of the loop, the short assignment operator `:=`, and then the keyword `range` and our iterable (array, slice, map etc) variable `lookup`.

> NOTE! Iteration over maps isn't ordered. Each iteration will return the key value pair in a random order.

`range` iterates over elements in a variety of data structures. Let’s see how to use `range` with some of the data structures we’ve already learned:

![](img%5Cgoinaday29.png)

> Here we use range to sum the numbers in a slice. Arrays work like this too.

`range` on arrays and slices provides both the `index` and `value` (the value returned being the `num` variable) for each entry. 

Above we didn’t need the `index`, so we ignored it with the blank identifier `_`. 

Sometimes we actually want the indexes though!

[More examples of range here](https://go.dev/play/p/CzyQe5jWaC-)

## Errors

In Go it’s idiomatic to communicate errors via an explicit, separate return value. This contrasts with the exceptions used in languages like Java and Ruby and the overloaded single result / error value sometimes used in C. Go’s approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for any other, non-error tasks.

![](img%5Cgoinaday30.png)

[Code here](https://go.dev/play/p/vcpKA8tPKFl)

https://pkg.go.dev/strconv

https://go.dev/blog/error-handling-and-go


## Custom Errors

We can create our own errors by importing the `errors` package and using it as a return type on a function:

![](img%5Cgoinaday31.png)

[Code here](https://go.dev/play/p/2C2BmAlqt02)

---

# LAB #2

Create an in-memory key value store.

- For now assume keys and values will be strings

- HINT – you will need a `map`…

**Stretch goals**

1. Cater for arbitrary data types

2. Modify to return error values

> Skeleton code to be implemented - https://goplay.tools/snippet/QYVuuQYIceh

---

# Part 4

## Installing Go

## Install

* [https://go.dev/doc/install](https://go.dev/doc/install)
* Open the MSI file you downloaded and follow the prompts to install Go. By default, the installer will install Go to Program Files or Program Files (x86). You can change the location as needed. After installing, you will need to close and reopen any open command prompts so that changes to the environment made by the installer are reflected at the command prompt.
* Verify that you've installed Go.
  * In  __Windows__ , click the  __Start__  menu.
  * In the menu's search box, type  __cmd__ , then press the  __Enter__  key.
  * In the Command Prompt window that appears, type the following command: `go version`
  * Confirm that the command prints the installed version of Go.

## VS Code Extensions

Install the VS Code Go extension:

[https://code.visualstudio.com/docs/languages/go](https://code.visualstudio.com/docs/languages/go)

## Running and Building

Create `main.go`. For now, you can save it anywhere you want; we don't need to live inside Go's workspace for trivial examples.

![](img%5Cgoinaday33.png)

> `println` is not `Println` and needs no import – that’s because println is part of the runtime of Go, and may be removed in future. 

Next, open a shell/command prompt and change the directory to where you saved the file.

`go run main.go`

`go build main.go`

## Formatting

Most programs written in Go follow the same formatting rules, namely a tab is used to indent and braces go on the same line as their statement.

When you're inside a project, you can apply the formatting rule to it and all sub-projects via:

![](img%5Cgoinaday34.png)

Give it a try. It does more than indent your code; it also aligns field declarations and alphabetically orders imports.

---

# Part 5

## Packages

## How do we use our own packages?

1. Create a new folder (I’ve used `GoInADay` in my documents folder)

2. Open this folder up via VS Code

3. Open a new Terminal and run `go mod init store`

![](img%5Cgoinaday35.png)


To keep more complicated libraries and systems organized, we need to learn about packages. 

A note on go mod and its history – sometimes you’ll look up things about Go and it’ll tell you all about GOPATH or GOROOT. This applies to older versions of Go. For our purposes we’ve ignored it and are only using modules.

When you do this, you will have a `go.mod` file in your root which is like `package.json` in node – it’s really a list of all your packages and where you get them. Maven, Nuget – all similar things.

In the store directory, open the `go.mod` file using nano, or your favorite text editor:

The first line, the module directive, tells Go the name of your module so that when it’s looking at import paths in a package, it knows not to look elsewhere for store. The store value comes from the parameter you passed to go mod init.

This will become very important shortly about how Go finds its packages

## Create a DB package

Create a new folder under your base folder  (for me _GoInADay_ referenced as the module  _store_  in `go.mod`) called `shopping` and a subfolder within it called `db`:

![](img%5Cgoinaday37.png)

Inside of  __shopping/__  __db__ , create a file called  __db.go__  __ __ and add the following code:

![](img%5Cgoinaday36.png)

In Go, package names follow the directory structure of your Go workspace. If we were building a shopping system, we'd probably start with a package name "shopping" and put our source files in `/shopping/`.

We don't want to put everything inside this folder though. For example, maybe we want to isolate some database logic inside its own folder. To achieve this, we create a subfolder at `/shopping/db`. The package name of the files within this subfolder is simply `db`, but to access it from another package, including the shopping package, we need to `import shopping/db`.

In other words, when you name a package, via the package keyword, you provide a single value, not a complete hierarchy (e.g., "shopping" or "db"). When you import a package, you specify the complete path.

Notice that the name of the package is the same as the name of the folder. Also, obviously, we aren't actually accessing the database. We're just using this as an example to show how to organize code.


## Create a Shopping Package

Now, create a file called `pricecheck.go` inside of the main shopping folder.

![](img%5Cgoinaday38.png)


You’ll see here that we need to refer to the fully-qualified path from the module root (GoInADay folder) up to the `db` package. This is what you have to remember about modules!

It's tempting to think that importing `shopping/db` is somehow special because we're inside the shopping package/folder already. In reality, you're importing `store/shopping/db`, which means you could just as easily import test/db so long as you had a package named db inside of your workspace’s modulename/test folder.

## You still need a `main()`…

![](img%5Cgoinaday39.png)

If you're building a package, you don't need anything more than what we've seen. To build an executable, you still need a main. I added mine to the root, but other prefer to create a folder – for instance named shopping – and put it in there.

You can now run your code by going into your shopping project and typing:

`go run main/main.go`

## Visibility – or public vs private

Go uses a simple rule to define what types and functions are visible outside of a package. If the name of the type or function starts with an uppercase letter, it's visible. If it starts with a lowercase letter, it isn’t.

This also applies to structure fields. If a structure field name starts with a lowercase letter, only code within the same package will be able to access them.

For example, if our `db.go` file had a function that looked like:

![](img%5Cgoinaday40.png)

it could be called via `db.NewItem()`. But if the function was named `newItem`, we wouldn't be able to access it from a different package.

Go ahead and change the name of the various functions, types and fields from the shopping code. For example, if you rename the Item's `Price` field `to price`, you should get an error.

## Package Management

* The go command we've been using to run and build has a get subcommand which is used to fetch third-party libraries. go get supports various protocols but for this example, we'll be getting a library from Github.
* From a shell/command prompt, enter:
  * go get github.com/mattn/go-sqlite3
* __go get __ fetches the remote files and stores them in your workspace. Go ahead and check your $GOPATH/src. In addition to the shopping project that we created, you'll now see a  __github.com __ folder. Within, you'll see a  __mattn__  folder which contains a  __go-sqlite3 __ folder.
* We just talked about how to import packages that live in our workspace. To use our newly gotten go-sqlite3 package, we'd import it like so:
* I know this looks like a URL but in reality, it'll simply import the go-sqlite3 package which it expects to find in $module/github.com/mattn/go-sqlite3.

![](img%5Cgoinaday41.png)

# Go Mod Tidy

* This command will basically match the  __go.mod__  file with the dependencies required in the source files.
  * Download all the dependencies that are required in your source files and update  __go.mod__  file with that dependency.
  * Remove all dependencies from the go.mod file which are not required in the source files.

![](img%5Cgoinaday42.png)

---

# Part 6

# Variadic, Anonymous Functions, Currying

# Variadic Parameters

![](img%5Cgoinaday43.png)

---

Rather than accepting a slice – in this case of int - we might want to accept any amount of ints directly

Must be last in the argument list. Cannot have variadic first, must be last and there can be only one!



# Unfurling Slices

![](img%5Cgoinaday44.png)

---

If you’ve done some JS, you will have seen this!

## Anonymous Functions

Anonymous functions can be declared inside other functions

![](img%5Cgoinaday45.png)

Note in the first one we do not have to provide a param to it.

In the second, we do.

## Currying

Function named wtf which has no parameters, returns an anonymous function which has no parameters that returns a string.

![](img%5Cgoinaday46.png)

The important thing about currying is the function returns a function *and* a return type


## Implementing Interfaces

![](img%5Cgoinaday47.png)

Okay, first thing you’ll notice is again….this isn’t Java

Implementing interfaces is not done the same way – there appears to be no relation between our structs and our interface…

![](img%5Cgoinaday48.png)

We have to create a function with a receiver of the type of struct for each with the EXACT name and return type declared in our interface. 
Go doesn’t care, this is just how it is….

## Structural vs Nominative Polymorphism



Okay, first thing you’ll notice is again….this isn’t Java

![](img%5Cgoinaday49.png)

Implementing interfaces is not done the same way – there appears to be no relation between our structs and our interface….. (click)

We have to create a function with a receiver of the type of struct for each with the EXACT name and return type declared in our interface. Go doesn’t care, this is just how it is….

![](img%5Cgoinaday50.png)


Now we have that, we can utilise a single function to accept the interface (Employee) and operate on ONLY interface defined functions

![](img%5Cgoinaday51.png)

We create a list of Employees in main, let’s follow along as to what happens….

![](img%5Cgoinaday52.png)

## Redo!

A hospital has five departments: Cardiology, Neurology, Orthopaedics, Gynaecology and Oncology. There are N patients, numbered from 0 to N-1, and the K-th of them is in department represented by a string A[K].

Write a function that, given an array A consisting of N strings, returns the maximum number of patients in one department.

Examples:

1. Given A = ["Cardiology", "Orthopaedics", "Neurology", "Cardiology", "Orthopaedics", "Cardiology"], the function should return 3. The department of Cardiology is occupied by three patients.

2. Given A = ["Oncology", "Gynaecology", "Orthopaedics", "Oncology", "Gynaecology", "Orthopaedics"], the function should return 2.

3. Given A = ["Neurology", "Cardiology", "Oncology"], the function should return 1

Each element of array A is a string that can have one of the following values: Cardiology, Neurology, Orthopaedics, Gynaecology, Oncology.

---

# Part 7

## CONCURRENCY SUPPORT

## Goroutines

A goroutine is similar to a thread, but it is scheduled by Go, not the OS. Code that runs in a goroutine can run concurrently with other code. Let's look at an example.

![](img%5Cgoinaday53.png)

There are a few interesting things going on here, but the most important is how we start a goroutine. We simply use the go keyword followed by the function we want to execute. If we just want to run a bit of code, such as the above, we can use an anonymous function. Do note that anonymous functions aren't only used with goroutines, however.

Goroutines are easy to create and have little overhead. Multiple goroutines will end up running on the same underlying OS thread. This is often called an M:N threading model because we have M application threads (goroutines) running on N OS threads. The result is that a goroutine has a fraction of overhead (a few KB) than OS threads. On modern hardware, it's possible to have millions of goroutines.

Furthermore, the complexity of mapping and scheduling is hidden. We just say this code should run concurrently and let Go worry about making it happen.
If we go back to our example, you'll notice that we had to Sleep for a few milliseconds. That's because the main process exits before the goroutine gets a chance to execute (the process doesn't wait until all goroutines are finished before exiting). To solve this, we need to coordinate our code.


## Channels

The challenge with concurrent programming stems from sharing data. If your goroutines share no data, you needn't worry about synchronizing them. That isn't an option for all systems, however. In fact, many systems are built with the exact opposite goal in mind: to share data across multiple requests. An in-memory cache or a database, are good examples of this. This is becoming an increasingly common reality.

Channels help make concurrent programming saner by taking shared data out of the picture. A  __channel__  is a  __communication__  pipe between  __goroutines__  which is used to pass  __data__ . In other words, a goroutine that has data can pass it to another goroutine via a channel. The result is that, at any point in time, only one goroutine has access to the data.

The final thing to know before we move on is that receiving and sending to and from a channel is  __blocking__ . That is, when we receive from a channel, execution of the goroutine won't continue until data is available. Similarly, when we send to a channel, execution won't continue until the data is received.

![](img%5Cgoinaday54.png)

---

Creating goroutines is trivial, and they are so cheap that we can start many; however, concurrent code needs to be coordinated. To help with this problem, Go provides channels. Before we look at channels, I think it's important to understand a little bit about the basics of concurrent programming.
Writing concurrent code requires that you pay specific attention to where and how you read and write values. In some ways, it's like programming without a garbage collector -- it requires that you think about your data from a new angle, always watchful for possible danger.



## CHANNELS BASICS Lab

* Ordering experiment
  * Create 10 goroutines. Make each one print a single, different word. Run all 10. Do this a few times
  * What do you notice about the order?
* Channel experiment
  * Create a goroutine that sends the integers 1 to 10 over a channel. Create another go routine that reads them off and prints them out
* Worker pool experiment
  * Modify Channel Experiment. Make three ‘worker’ goroutines read numbers of the channel and print
  * Run this a few times. Is the order consistent? When might this matter? When might it not matter?

# Part 8: Handling Concurrency

## Actor model

## Actor pattern

## Lab 3

* __OH NOES!__  It turns out that map[string]interface{} is not concurrency safe!
* So our key value store will corrupt data if we call it concurrently
* Create an Actor to add concurrency safety to our key-value store which
  * __put__ ( key, value )
  * __delete __ ( key )
  * __read__ ( key ) returns value
  * __shutdown __ – without losing any work in flight
* Do not change the key-value store we already have; let the actor manage all calls to it

