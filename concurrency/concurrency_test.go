package concurrency

import (
	"reflect"
	"testing"
	"time"
)

// Let's use a benchmark to test the speed of CheckWebsites,
// so we can see the effect of our changes when we implement concurrency

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	// b.N is benchmarks max iteration
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}

}

func mockWebsiteChecker(url string) bool {
	if url == "not-a-website.com" {
		return false
	}

	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"not-a-website.com",
		"google.com",
		"youtube.com",
	}

	want := map[string]bool{
		"not-a-website.com": false,
		"google.com":        true,
		"youtube.com":       true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}
