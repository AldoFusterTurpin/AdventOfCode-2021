package device_test

import (
	"reflect"
	"testing"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day7/device"
)

func TestConcatenatePaths(t *testing.T) {
	tests := []struct {
		path []string
		want string
	}{
		{
			path: []string{"/"},
			want: "/",
		},
		{
			path: []string{"/", "a"},
			want: "//a",
		},
		{
			path: []string{"/", "a", "b", "c"},
			want: "//a/b/c",
		},
		{
			path: []string{"/", "a", "d"},
			want: "//a/d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := device.ConcatenatePaths(tt.path, "/"); got != tt.want {
				t.Errorf("got %v,but want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyseLsCommand(t *testing.T) {
	tests := []struct {
		name                    string
		currentDirName          string
		dirsInfo                map[string]int
		lines                   []string
		dirsStack               []device.Dir
		wantNItemsInDir         int
		wantDirectoriesStackOut []device.Dir
		wantDirsInfo            map[string]int
	}{
		{
			name:            "ls when / is the cwd, contains 2 files but not subfolders",
			currentDirName:  "/",
			lines:           []string{"$ ls", "14848514 b.txt", "8504156 c.dat"},
			wantNItemsInDir: 2,
			wantDirsInfo:    map[string]int{"/": 23352670},
		},
		{
			name:            "ls when / is the cwd, contains 4 files but not subfolders",
			currentDirName:  "/",
			lines:           []string{"$ ls", "4060174 j", "8033020 d.log", "5626152 d.ext", "7214296 k"},
			wantNItemsInDir: 4,
			wantDirsInfo:    map[string]int{"/": 24933642},
		},
		{
			name:           "ls when / is the cwd and contains 2 files and 1 subfolder",
			currentDirName: "/",
			lines:          []string{"$ ls", "14848514 b.txt", "8504156 c.dat", "dir d"},
			wantDirectoriesStackOut: []device.Dir{
				{
					Path:    "/",
					Content: []string{"14848514 b.txt", "8504156 c.dat", "dir d"},
				},
			},
			wantNItemsInDir: 3,
			wantDirsInfo:    map[string]int{},
		},
		{
			name:           "ls when / is the cwd and contains 3 files and 3 subfolder",
			currentDirName: "/",
			lines:          []string{"$ ls", "dir a", "73823 a.txt", "dir b", "14848514 b.txt", "8504156 c.dat", "dir d"},
			wantDirectoriesStackOut: []device.Dir{
				{
					Path:    "/",
					Content: []string{"dir a", "73823 a.txt", "dir b", "14848514 b.txt", "8504156 c.dat", "dir d"},
				},
			},
			wantNItemsInDir: 6,
			wantDirsInfo:    map[string]int{},
		},
		{
			name:           "ls when //a/b is the cwd and contains 1 files and 2 subfolders",
			currentDirName: "//a/b",
			lines:          []string{"$ ls", "dir a", "73823 a.txt", "dir b", "$ cd ..", "$ cd .."},
			wantDirectoriesStackOut: []device.Dir{
				{
					Path:    "//a/b",
					Content: []string{"dir a", "73823 a.txt", "dir b"},
				},
			},
			wantNItemsInDir: 3,
			wantDirsInfo:    map[string]int{},
		},
		{
			name:            "ls when //a/b is the cwd and is empty",
			currentDirName:  "//a/b",
			lines:           []string{"$ ls", "$ cd ..", "$ cd b", "$ ls", "$ cd ..", "$ cd .."},
			wantNItemsInDir: 0,
			wantDirsInfo:    map[string]int{"//a/b": 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNItemsInDir, gotDirectoriesStackOut, gotDirsInfo := device.AnalyseLsCommand(tt.currentDirName, tt.dirsInfo, tt.lines, tt.dirsStack)
			if gotNItemsInDir != tt.wantNItemsInDir {
				t.Errorf("got %v, but wanted %v", gotNItemsInDir, tt.wantNItemsInDir)
			}

			if !reflect.DeepEqual(gotDirectoriesStackOut, tt.wantDirectoriesStackOut) {
				t.Errorf("got %v, but wanted %v", gotDirectoriesStackOut, tt.wantDirectoriesStackOut)
			}

			if !reflect.DeepEqual(gotDirsInfo, tt.wantDirsInfo) {
				t.Errorf("got %v, but wanted %v", gotDirsInfo, tt.wantDirsInfo)
			}
		})
	}
}

func TestGetDirsInfoAndDirsStack(t *testing.T) {
	type TestData struct {
		input             string
		expectedDirsInfo  map[string]int
		expectedDirsStack []device.Dir
	}

	tests := map[string]TestData{
		"root / folder": {
			input: `$ cd /
$ ls
14848514 b.txt
8504156 c.dat`,
			expectedDirsInfo: map[string]int{
				"/": 23352670,
			},
			expectedDirsStack: nil,
		},
		"leaf d folder": {
			input: `$ cd /
$ cd a
$ cd ..
$ cd a
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`,
			expectedDirsInfo: map[string]int{
				"//d": 24933642,
			},
		},
		"leaf b folder inside d": {
			input: `$ cd /
$ cd a
$ cd ..
$ cd a
$ cd ..
$ cd a
$ cd d
$ cd ..
$ cd d
$ cd s
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`,
			expectedDirsInfo: map[string]int{
				"//a/d/s": 24933642,
			},
		},
		"simple_case": {
			input: `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
29116 f
2557 g
62596 h.lst`,
			expectedDirsInfo: map[string]int{
				"//a": 94269,
			},
			expectedDirsStack: []device.Dir{
				{
					Path:    "/",
					Content: []string{"dir a", "14848514 b.txt", "8504156 c.dat", "dir d"},
				},
			},
		},
		"statement_case": {
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
			expectedDirsInfo: map[string]int{
				"//a/e": 584,
				"//d":   24933642,
			},
			expectedDirsStack: []device.Dir{
				{
					Path:    "/",
					Content: []string{"dir a", "14848514 b.txt", "8504156 c.dat", "dir d"},
				},
				{
					Path:    "//a",
					Content: []string{"dir e", "29116 f", "2557 g", "62596 h.lst"},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			dirsInfo, dirsStack := device.GetDirsInfoAndDirsStack(tc.input)
			if !reflect.DeepEqual(dirsInfo, tc.expectedDirsInfo) {
				t.Fatalf("expected %v, but got %v", tc.expectedDirsInfo, dirsInfo)
			}

			if !reflect.DeepEqual(dirsStack, tc.expectedDirsStack) {
				t.Fatalf("expected %v, but got %v", tc.expectedDirsStack, dirsStack)
			}
		})
	}
}

func TestGetSizeOfDirs(t *testing.T) {
	type TestData struct {
		input          string
		expectedResult map[string]int //key:name, value: size
	}

	tests := map[string]TestData{
		"leaf root folder": {
			input: `$ cd /
$ ls
14848514 b.txt
8504156 c.dat`,
			expectedResult: map[string]int{
				"/": 23352670,
			},
		},
		"leaf d folder": {
			input: `$ cd /
$ cd a
$ cd ..
$ cd a
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`,
			expectedResult: map[string]int{
				"//d": 24933642,
			},
		},
		"leaf b folder inside d": {
			input: `$ cd /
$ cd a
$ cd ..
$ cd a
$ cd ..
$ cd a
$ cd d
$ cd ..
$ cd d
$ cd s
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`,
			expectedResult: map[string]int{
				"//a/d/s": 24933642,
			},
		},
		"sample_input": {
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
			expectedResult: map[string]int{
				"/":     48381165,
				"//a":   94853,
				"//a/e": 584,
				"//d":   24933642,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := device.GetSizeOfDirs(tc.input)
			if !reflect.DeepEqual(got, tc.expectedResult) {
				t.Fatalf("expected %v, but got %v", tc.expectedResult, got)
			}
		})
	}
}

func TestTraverseStackAndUpdateDirsInfo(t *testing.T) {
	type TestData struct {
		dirsStack        []device.Dir
		dirsInfo         map[string]int
		expectedDirsInfo map[string]int
	}

	tests := map[string]TestData{
		"statement_case": {
			dirsInfo: map[string]int{
				"//a/e": 584,
				"//d":   24933642,
			},
			dirsStack: []device.Dir{
				{
					Path:    "/",
					Content: []string{"dir a", "14848514 b.txt", "8504156 c.dat", "dir d"},
				},
				{
					Path:    "//a",
					Content: []string{"dir e", "29116 f", "2557 g", "62596 h.lst"},
				},
			},
			expectedDirsInfo: map[string]int{
				"//a/e": 584,
				"//d":   24933642,
				"/":     48381165,
				"//a":   94853,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			device.TraverseStackAndUpdateDirsInfo(tc.dirsInfo, tc.dirsStack)
			if !reflect.DeepEqual(tc.dirsInfo, tc.expectedDirsInfo) {
				t.Fatalf("expected %v, but got %v", tc.expectedDirsInfo, tc.dirsInfo)
			}
		})
	}
}

func TestSolveProblem(t *testing.T) {
	type TestData struct {
		input          string
		maxSizePerDir  int
		expectedResult int
	}

	tests := map[string]TestData{
		"sample_input": {
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
			expectedResult: 95437,
			maxSizePerDir:  100000,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := device.SolveProblem(tc.input, tc.maxSizePerDir)
			if got != tc.expectedResult {
				t.Fatalf("expected %v, but got %v", tc.expectedResult, got)
			}
		})
	}
}
