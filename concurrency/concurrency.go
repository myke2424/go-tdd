package concurrency

type WebsiteChecker func(string) bool

// anonymous fields in the struct
// useful when you don't know a name for the values
type result struct {
	string
	bool
}

// Returns a map of each URL checked to a boolean
// true for good response / false for bad response

// In go and most programming languages, when a function is executing,
// we're are waiting for the execution to finish. This is called a "blocking" call.
// An operation that does NOT block in Go will run in a seperate process called a goroutine.
// To tell Go to start a new routine; we turn a function into a 'go' statement
// by putting the keyword 'go' in front of the function: go doSomething()
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	// creating a channel of result, our new type
	// this type is made to associate the return value of website checker
	resultChannel := make(chan result)

	// Since the only way to start a goroutine is to put 'go' in front of the function call,
	// we often use anonymous functions when we want to start a goroutine.
	// anon functions share the same scope as the function
	// Here we are passing url value to our anon func

	// Maps in Go don't like it when more that one thing tries to write to them (concurrent writing)
	// This is a race condition.
	// A bug that occurs when the output of our software is dependent on the timing and sequence of events that we
	// have no control over.
	// Go has a built in race detector, type go test -race
	// We can solve this race condition by coordinating our goroutines using channels.
	// Channels are a Go data structure that can both send/recieve values.
	// They allow communication between different processes.

	for _, url := range urls {
		go func(u string) {
            // Passing the result struct values to the resultChannel
            // <- operator is the send operator, taking a channel on the left and a value on the right
			resultChannel <- result{u, wc(u)}
		}(url)
	}

    for i := 0; i < len(urls); i++ {
        // Recieve data from the channel
        // <-resultChannel is the recieve expressions, which assigns a value recieved from the channel to a variable
        // The two operands are now reversed, the channel on the right, and variable assigned on left
        r := <-resultChannel
        results[r.string] = r.bool
    }

	return results
}
