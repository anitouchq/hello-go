package greeting

import (
	"strings"
	"testing"
)

func TestMessage(t *testing.T) {
	got := Message()
	if !strings.Contains(got, "Go") {
		t.Fatalf("Message() = %q, want it to contain Go", got)
	}
}

func TestDetails(t *testing.T) {
	got := Details("linux", "amd64")
	want := "OS: linux\nArch: amd64"
	if got != want {
		t.Fatalf("Details() = %q, want %q", got, want)
	}
}
