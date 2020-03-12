# Roman Numeral Encoder

This is a programming Kata taken from [Codewars](https://www.codewars.com/kata/51b62bf6a9c58071c600001b) and implemented in _Golang_.

## Solution

Split an integer to its sum factors. Then convert each factor to a roman numeral. Finally, concatenate numerals to get the result.

**Example 1: `2020` to Roman:**

```go
n := 2020
=   2000 + 20
=   1000 + 1000 + 20
=   1000 + 1000 + 10 + 10
=   {M} + 1000 + 10 + 10
=   {M} + {M} + 10 + 10
=   {M} + {M} + {X} + 10
=   {M, M} + {X, X}
=   MM + XX
=   MMXX
```

**Example 1: `8192` to Roman:**

Following example skips some of the trivial steps for brevity's sake.

```go
n := 8192
=   8000 + 100 + 90 + 2
=   {M, M, M, M} + 100 + 90 + 2
=   {M, M, M, M} + {C} + 90 + 2
=   {M, M, M, M} + {C} + {XC} + 2
=   {M, M, M, M} + {C} + {XC} + {II}
=   MMMM + C + XC + II
=   MMMMMMMMCXCII
```

See [here](encoder/encoder.go) for my complete solution.

## Testing

Use `./test.sh` to run both unit tests and performance benchmarks.
