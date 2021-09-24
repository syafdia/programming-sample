package main

import "testing"

type args struct {
	str string
}

type testCase struct {
	name string
	args args
	want string
}

var testCases []testCase = []testCase{
	{
		name: "Case 001",
		args: args{str: "abcd"},
		want: "",
	},
	{
		name: "Case 002",
		args: args{str: "abcd (x)"},
		want: "",
	},
	{
		name: "Case 003",
		args: args{str: "abcd (xyz)"},
		want: "xy",
	},
	{
		name: "Case 004",
		args: args{str: "abcd (xy"},
		want: "",
	},
}

func Test_findFirstStringInBracketV1(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFirstStringInBracketV1(tt.args.str); got != tt.want {
				t.Errorf("findFirstStringInBracketOld() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findFirstStringInBracketV2(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFirstStringInBracketV2(tt.args.str); got != tt.want {
				t.Errorf("findFirstStringInBracketOld() = %v, want %v", got, tt.want)
			}
		})
	}
}
