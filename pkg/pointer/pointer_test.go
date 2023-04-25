package pointer

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				s: "123",
			},
		},
		{
			name: "",
			args: args{
				s: "456",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.s); got == nil {
				t.Errorf("String() = %v", got)
			} else {
				fmt.Println(got)
			}
		})
	}
}
