package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// Learnings: 
// 1. goroutines using go ... speed up the things that can be done async 
// 2. anon functions -> can take in whatevers already in its scope 
// 3. channels for synchronisation, for the parts that have to be done sync 
// 4. bench tests / race tests / make the code right and work before making it fast 

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}