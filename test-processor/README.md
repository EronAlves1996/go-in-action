### Text Processor Kata

A simple implementation of a text processor to practice Go testing and benchmarking from "Go in Action" Chapter 9.

**What I Built**

A `TextProcessor` with two methods: `ReverseWords` and `CountVowels`. The main focus was on writing table-driven tests and benchmarks rather than the logic itself.

**Implementation Notes**

The `ReverseWords` method uses `strings.Split` on spaces, which means it normalizes multiple spaces into single ones. If you need to preserve exact whitespace formatting, you'd need a different approach using `strings.Fields` or regex.

For `CountVowels`, I used `slices.Contains` which is clean but might be slower than a map lookup for benchmarks. A map[rune]bool would probably perform better under heavy load.

**Key Learnings**

- Table tests make testing multiple cases much cleaner than individual test functions
- `strings.Builder` with pre-allocated space (`Grow(len(input))`) performs better than `strings.Join`
- The modern benchmark pattern uses `b.Loop()` instead of traditional for loops
- Your tests define the contract - my implementation passes the tests but has specific behavior around whitespace

**Benchmark Results**
The current implementation shows good performance due to `strings.Builder` and pre-allocation. The vowel counting could be optimized further with a map if needed.

**Questions for Further Exploration**

- How would the performance change if I used a map for vowel lookup instead of `slices.Contains`?
- What would be the most efficient way to preserve original whitespace in `ReverseWords`?
- How does this compare to using `strings.Fields` which handles whitespace differently?

The code is straightforward - the value here is in the test patterns and benchmarking approach.
