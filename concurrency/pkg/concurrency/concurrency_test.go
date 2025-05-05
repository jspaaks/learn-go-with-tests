package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func TestUrlChecker(t *testing.T) {
	const url0 string = "http://google.com"
	const url1 string = "http://blog.gypsydave5.com"
	const url2 string = "waat://furhurterwe.geds"
	urls := []string{
		url0,
		url1,
		url2,
	}
	mockUrlChecker := func (url string) bool {
		return url != url2
	}
	got := CheckUrls(mockUrlChecker, urls)
	want := StatusMap{
		url0: true,
		url1: true,
		url2: false,
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func BenchmarkUrlChecker(b * testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "some url"
	}
	mockUrlChecker := func (_ string) bool {
		time.Sleep(20 * time.Millisecond)
		return true
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckUrls(mockUrlChecker, urls)
	}
}
