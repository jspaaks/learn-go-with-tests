package concurrency

type UrlChecker func(string) bool
type StatusMap map[string]bool

func CheckUrls(check UrlChecker, urls []string) StatusMap {
	// make the results channel
	type result struct {
		url    string
		status bool
	}
	channel := make(chan result)

	// sending
	for i := 0; i < len(urls); i++ {
		go func() {
			channel <- result{url: urls[i], status: check(urls[i])}
		}()
	}

	// receiving
	results := make(StatusMap)
	for i := 0; i < len(urls); i++ {
		result := <-channel
		results[result.url] = result.status
	}

	return results
}
