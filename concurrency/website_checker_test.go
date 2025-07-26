package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "https://ok.pluto.dev"
}

func TestWebChecker(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://beta-api.yelo.red",
		"https://ok.pluto.dev",
		"https://pluto-food.freelancer.jungleworks.me",
	}

	want := map[string]bool{
		"https://google.com":                           true,
		"https://beta-api.yelo.red":                    true,
		"https://ok.pluto.dev":                         false,
		"https://pluto-food.freelancer.jungleworks.me": true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkWebChecker(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "some url"
	}

	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
