package main

import "testing"

func TestRacer(t *testing.T) {
	slowUrl := "https://www.youtube.com"
	fastUrl := "https://www.ozon.ru"

	want := fastUrl
	got := Racer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
