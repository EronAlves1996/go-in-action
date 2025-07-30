# Library Management System

This is a simple demonstration of Go Type System as a code kata

## Proposal

How the Go's type system works? Structs and interfaces and primitive types?
The focus here is only experiment in a practical way how is the behavior of go type system.

## What I learned

I don't learned so much things on this regard, because the go type system works
the same way as much type systems out there.

Here we have some pointer manipulations, but only to a extend, where I have more freedom
to define the program memory layout. In Java, as example, the primitive types
lives in stack, and are immutable, and all the other types lives in heap.

In Go, you can choose where your use defined types lives. But we still have some
primitive types like string (here is primitive), numeric types, boolean, and the reference types.
For the reference types, there's a difference: you can choose where the header lives, but internaly
it points to the content of data structure. Slices, maps and channels obeys this contract, as example.

## Questions for further exploration

A thing that I tried, but get some doubt about is the following pattern:

```go
func (u *User) Borrow(b *Book) {
  ub := u.Books
  ub = append(ub, b)
}
```

The new slice will be assigned to u.Books directly? Or should I assign
`ub` to u.Books right after the append operation?
