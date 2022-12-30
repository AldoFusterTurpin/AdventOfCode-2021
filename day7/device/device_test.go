package device_test

import (
	"reflect"
	"testing"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day7/device"
)

func TestGetDirectorySizes(t *testing.T) {
	type TestData struct {
		input          string
		expectedResult []device.FileInfo
		expectedErr    error
	}

	tests := map[string]TestData{
		"sample_input": TestData{
			input: `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`,
			expectedResult: []device.FileInfo{
				{
					Name: "e",
					Size: 584,
				},
				{
					Name: "a",
					Size: 94853,
				},
				{
					Name: "d",
					Size: 24933642,
				},
				{
					Name: "/",
					Size: 48381165,
				},
			},
			expectedErr: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := device.GetDirectorySizes(tc.input)
			// if err != nil {
			// 	t.Fatalf("unexpected error: %v", err)
			// }

			if !reflect.DeepEqual(got, tc.expectedResult) {
				t.Fatalf("expected %v, but got %v", tc.expectedResult, got)
			}
		})
	}

}
