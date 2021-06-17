package gset

import (
	"fmt"
	"testing"
)

var s Gset

func init() {
	s = New()
}

func TestAppend(t *testing.T) {
	// s.Append("abcd")
	cases := []struct {
		name    string
		in      string
		desired []string
	}{
		{"empty string", "", []string{""}},
		{"normal value", "some", []string{"", "some"}},
	}
	for i := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			s.Append(cases[i].in)
			got := s.GetSet()
			if len(got) != len(cases[i].desired) {
				t.Fatalf("expected a slice of length %d and got %d ", len(cases[i].desired), len(got))
			}
		})
	}
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Append(fmt.Sprintf("%d", i))
	}
}

func BenchmarkGetSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.GetSet()
	}
}
