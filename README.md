# go-examples

Example projects for Golang

### Project 1

Hello World

### Project 2

Values types in Go

### Project 3

Variables in Golang

### Project 4

Constants

### Project 5

Only looping construct in Go "for"

### Project 6

if/else

### Project 7

Switch

### Project 8

Arrays

### Project 9

Slices

### Project 10

Maps: `Maps are Go’s built-in associative data type
(sometimes called hashes or dicts in other languages).`

### Project 11

Functions

### Project 12

Multi Return

### Project 13

Variadic Functions: `Variadic functions can be called with any number of trailing arguments.`

### Project 14

Closures: `Define a function inline without having to name it.`

### Project 15

Range

### Project 16

Pointers : `Allowing you to pass references to values and records within your program.`

### Project 17

Runes: `In other languages, strings are made of “characters”. In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point. `

### Project 18

Structs: `Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.`

### Project 19

Methods: `Go supports methods defined on struct types.`

### Project 20

Interfaces: `Interfaces are named collections of method signatures.`

### Project 21

Enums

### Project 22

Struct Embedding

### Project 23

Errors

### Project 24

Goroutines: `A goroutine is a lightweight thread of execution.`

### Project 25

Channels: `Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.`

### Project 26

Channel Buffering: `By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.`

### Project 27

Channel Synchronization: `We can use channels to synchronize execution across goroutines.`

### Project 28

Channel Directions

### Project 29

Select: `Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.`

### Project 30

Timeouts

### Project 31

Timers

### Project 32

Tickers: `Timers are for when you want to do something once in the future - tickers are for when you want to do something repeatedly at regular intervals. `

### Project 33

Worker Pools: `Example for goroutines and channels`

### Project 34

Wait Groups: `To wait for multiple goroutines to finish, we can use a wait group.`

### Project 35

Rate Limiting: `Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service.`

### Project 36

Atomic Counters: `The primary mechanism for managing state in Go is communication over channels.  Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.`

### Project 37

Stateful Goroutines: `Turkish explanation is inside of the folder.`

### Project 38

Sorting

### Project 39

Sorting by Functions: `Sometimes we’ll want to sort a collection by something other than its natural order. For example, suppose we wanted to sort strings by their length instead of alphabetically. Here’s an example of custom sorts in Go.`

### Project 40

Panic: `A panic typically means something went unexpectedly wrong.
Mostly we use it to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.`

### Project 41

Defer: `
Defer is used to ensure that a function call is performed later in a program’s execution`

### Project 42

Recover: `
Go makes it possible to recover from a panic, by using the recover built-in function. A recover can stop a panic from aborting the program and let it continue with execution instead.`

### Project 43

String Functions

### Project 44

String Formatting

### Project 45

Text Templates

### Project 46

JSON

### Project 47

XML

### Project 48

Time

### Project 49

Number Parsing

### Project 50

URL Parsing

### Project 51

SHA256 Hashing

### Project 52

Reading Files

### Project 53

Writing Files

### Project 54

Line Filters

### Project 55

Command-Line Arguments: `
Command-line arguments are a common way to parameterize execution of programs. For example, go run hello.go uses run and hello.go arguments to the go program.`

### Project 56

Command-Line Flags: `
Command-line flags are a common way to specify options for command-line programs. For example, in wc -l the -l is a command-line flag.`

### Project 57

Command-Line Subcommands: `
Some command-line tools, like the go tool or git have many subcommands, each with its own set of flags. For example, go build and go get are two different subcommands of the go tool. The flag package lets us easily define simple subcommands that have their own flags.`

### Project 58

Logging

### Project 59

HTTP Client: `
The Go standard library comes with excellent support for HTTP clients and servers in the net/http package. In this example we’ll use it to issue simple HTTP requests.`

### Project 60

HTTP Server: `
Writing a basic HTTP server is easy using the net/http package.`

### Project 61

Context
