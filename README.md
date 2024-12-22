# Go Race Condition Example

This repository demonstrates a race condition in a Go program that calculates the sum of numbers using multiple goroutines.  The program uses a `sync.WaitGroup` to ensure all goroutines complete, but it lacks proper synchronization mechanisms for accessing and modifying the shared `count` variable.

The `bug.go` file contains the code exhibiting the race condition. The solution is demonstrated in `bugSolution.go`.

## How to Reproduce

1. Clone this repository.
2. Run `go run bug.go`.
3. Observe the varying results of the summation due to the race condition.
4. Run `go run bugSolution.go` to see the corrected version.