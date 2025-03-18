package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// * Implementing concurrency
		go func() {
			// * Sending result struct to the resultChannel with a send statement (<-)
			// * Taking a channel on the left and a value on the right
			resultChannel <- result{url, wc(url)}
		}()
	}
	for i := 0; i < len(urls); i++ {
		// * Receive expression, ehich assigns a value received from a channel to a variable
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
