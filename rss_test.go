package main

import "testing"

func TestStripPodcastEpisodePrefix(t *testing.T) {
	for _, tt := range []struct{ in, want string }{
		{"Better Commit Messages", "Better Commit Messages"},
		{"10: Minisode 0.1.1", "Minisode 0.1.1"},
		{"", ""},
	} {
		got := stripPodcastEpisodePrefix(tt.in)

		if got != tt.want {
			t.Errorf("stripPodcastEpisodePrefix(%q) = %q; want %q", tt.in, got, tt.want)
		}
	}
}
