# Go 101

Go is a general purpose programming language that was created by Google in 2007 and open sourced in 2009. Version 1.0 was released in 2012. The last two major Go releases are supported. 

Go 1 is the stable platform. Major releases are identified by the second major number. [All releases under Go 1 are backwards compatible.](https://go.dev/doc/go1compat)

**Good Resources**
- [Go.dev](https://go.dev/)
- [Effective Go](https://go.dev/doc/effective_go)
- [The Go Blog](https://go.dev/blog/)

## Seven Languages in Seven Weeks style
In the style of *Seven Languages in Seven Weeks*, let's breakdown some aspects of Go (as I try to teach myself the language)

- [Day 2](Go02.md)
- [Day 3](Go03.md)

## Installing Go
Download and install Go from https://go.dev/dl/
After it completes, that can be checked as follows:

```shell
$ go version
go version go1.26.0
```

### Ecosystem and project setup
Every language has an ecosystem that is often harder to learn than the language syntax itself. As part of that ecosystem are also best practices on how to organize and setup a project. 

The management of dependencies (importing other packages) is done in the `go.mod` file. 

We can create one easily to get started

```shell
cd hello_go
go mod init example/hello
```

The project setup for a typical Go project:

```
hello_go/
  ├── go.mod
  ├── main.go
  ├── internal/ (optional)
  └── *_test.go (tests live alongside the code they test)
```

For example, `foo.go` typically has a matching `foo_test.go` in the same folder.

### Hello World
After creating a simple Hello Word 
We can build it...

```shell
$ go build
```

Unlike `mvn` in Java the compiled files are NOT placed in a `target` folder. 

We can run it from the executable file
```shell
$ ./hello
Hello, World!
```

Or we can combine the build and the run commands

```shell
$ go run .
Hello, World!
```