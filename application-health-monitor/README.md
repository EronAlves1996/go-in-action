# Health Monitor Kata

This is a simple demonstration of Go's standard library as a code kata, based on Chapter 8 of "Go in Action".

## Proposal

How do the `log`, `encoding/json`, and `os` packages work together? The focus was to experiment in a practical way with customized logging, marshaling, and file I/O.

## What I learned

I solidified the use of the `log.New()` function to create customized loggers with different outputs and prefixes. The `io.Writer` interface is key here, as both the file and the standard output satisfy it.

The JSON marshaling with `json.Marshal` is straightforward, but the struct tags (like `` `json:"field_name"` ``) are essential for controlling the output. I also used `json.MarshalIndent` for the final report, which is useful for human-readable files.

Error handling is mandatory for every operation that can fail. I used `log.Fatal` for critical errors and simple logging for others. The `defer` statement for `file.Close()` is the idiomatic way to manage resources.

## Questions for further exploration

A thing that I tried, but had to think about, was the random status generation. I used an array of status strings indexed by an enum, but my first implementation with `rand.Intn(2)` only generated two values, so the `WARNING` state was never reached. The range have to be adjusted to `rand.Intn(3)` to include all possibilities.

Another point: is it better to use `log.Println` or `fmt.Println` for the final JSON report? The logger adds a timestamp prefix, while `fmt` gives you the raw JSON. I kept the logger for consistency, but the choice depends on the use case.
