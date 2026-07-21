// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond provides a function to add a gigasecond to a time.
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond returns the time one gigasecond (1e9 seconds) after t.
func AddGigasecond(t time.Time) time.Time {
	gigasec := 1000000000

	return t.Add(time.Second * time.Duration(gigasec))
}
