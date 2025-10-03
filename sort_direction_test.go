package textfmt

import (
	"testing"
)

func TestSortDirectionValues(t *testing.T) {
	tests := []struct {
		name string
		got  SortDirection
		want SortDirection
	}{
		{"Ascending should be 0", Ascending, 0},
		{"Descending should be 1", Descending, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("got %d, want %d", tt.got, tt.want)
			}
		})
	}
}
