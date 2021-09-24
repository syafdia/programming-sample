package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Case 001",
			args: args{
				words: []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"},
			},
			want: [][]string{
				[]string{"kita", "atik", "tika"},
				[]string{"aku", "kua"},
				[]string{"makan"},
				[]string{"kia"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := GroupAnagram(tt.args.words)
			sort.Slice(got, func(i int, j int) bool {
				if len(got[i]) > 0 && len(got[i]) == len(got[j]) {
					a := got[i][0]
					b := got[j][0]

					return len(a) > len(b)
				}
				return len(got[i]) > len(got[j])
			})
			sort.Slice(tt.want, func(i int, j int) bool {
				w1 := tt.want[i]
				w2 := tt.want[j]
				if len(w1) > 0 && len(w1) == len(w2) {
					a := w1[0]
					b := w2[0]

					return len(a) > len(b)
				}

				return len(w1) > len(w2)
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
