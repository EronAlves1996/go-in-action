package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

// A job to be processed. In this case, a URL to fetch.
type Job struct {
	URL string
}

// A result of a job. Contains the URL and its title, or an error.
type Result struct {
	URL   string
	Title string
	Err   error
}

// Fetches the HTML body of a URL and extracts the title tag.
// Returns an error if any step fails.
func fetchTitle(url string) (string, error) {
	if strings.Contains(url, "panic") {
		panic("should panic")
	}
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code %d", resp.StatusCode)
	}

	// A very simple regex to find the <title> tag.
	// Note: Using a proper HTML parser (like goquery) is better for real applications.
	re := regexp.MustCompile(`<title>(.*?)</title>`)
	body := make([]byte, 1024*1024) // Read up to 1MB
	n, _ := resp.Body.Read(body)
	body = body[:n]

	matches := re.FindSubmatch(body)
	if len(matches) < 2 {
		return "", fmt.Errorf("no title tag found")
	}
	return string(matches[1]), nil
}

// The worker function. This will be run by multiple goroutines.
// 1. It receives a Job from the `jobs` channel.
// 2. It processes the job (calls `fetchTitle`).
// 3. It sends the Result to the `results` channel.
func worker(wg *sync.WaitGroup, jobs <-chan Job, results chan<- Result) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in worker", r)
		}
	}()
	defer wg.Done()
	// TODO: Implement the worker loop.
	// Use a for-range loop on the `jobs` channel.
	for j := range jobs {
		// Call fetchTitle for each job.
		r, err := fetchTitle(j.URL)
		// Send the result (with the URL and title or error) back to the `results` channel.
		results <- Result{
			URL:   j.URL,
			Title: r,
			Err:   err,
		}
	}
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://golang.org/doc",
		"https://github.com",
		"https://stackoverflow.com",
		"https://example.com",
		"http://localhost",
		"http://panic",
		"https://gobyexample.com/worker-pools", // Bonus!
	}

	// TODO:
	// 1. Create channels for jobs and results.
	jobs := make(chan Job, len(urls))
	results := make(chan Result, len(urls))

	// 2. Create a WaitGroup for the workers.
	var wg sync.WaitGroup

	// 3. Launch a fixed number of workers (e.g., 3).
	numWorkers := 3
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(&wg, jobs, results)
	}

	// 4. Send all the URLs as Jobs to the jobs channel.
	for _, url := range urls {
		jobs <- Job{URL: url}
	}
	close(jobs) // Important: Close the channel to signal that no more jobs are coming.

	// 5. Wait for all workers to finish, then close the results channel.
	// (Hint: Use a goroutine to wait and then close the channel).
	go func() {
		wg.Wait()
		close(results)
	}()

	// 6. Collect and print all the results from the results channel.
	for result := range results {
		if result.Err != nil {
			fmt.Printf("Error fetching %s: %v\n", result.URL, result.Err)
		} else {
			fmt.Printf("%s: %s\n", result.URL, result.Title)
		}
	}
}
