package main

import "testing"

func Test_part01(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "testinput", args: args{input: `$ cd /
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
7214296 k
`}, want: 95437},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part01(tt.args.input); got != tt.want {
				t.Errorf("part01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isCommand(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "$ ls", args: args{input: "$ ls"}, want: true},
		{name: "dir a", args: args{input: "dir a"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCommand(tt.args.input); got != tt.want {
				t.Errorf("isCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_directory_getDirectorySize(t *testing.T) {
	type fields struct {
		name           string
		files          map[string]*file
		subdirectories map[string]*directory
		parent         *directory
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "test 1 file", fields: fields{
				name: "a",
				files: map[string]*file{"a": {
					name: "a",
					size: 10,
				}},
				subdirectories: nil,
				parent:         nil,
			}, want: 10,
		},
		{
			name: "test 2 files", fields: fields{
				name: "a",
				files: map[string]*file{"a": {
					name: "a",
					size: 10,
				}, "b": {
					name: "b",
					size: 10,
				}},
				subdirectories: nil,
				parent:         nil,
			}, want: 20,
		},
		{
			name: "test 1 subdirectory", fields: fields{
				name:  "root",
				files: nil,
				subdirectories: map[string]*directory{
					"a": {files: map[string]*file{"a": {
						name: "a",
						size: 10,
					}}},
				},
				parent: nil,
			},
			want: 10,
		},
		{
			name: "test 2 subdirectory", fields: fields{
				name:  "root",
				files: nil,
				subdirectories: map[string]*directory{
					"a": {files: map[string]*file{"a": {
						name: "a",
						size: 10,
					}}},
					"b": {files: map[string]*file{"b": {
						name: "b",
						size: 10,
					}}},
				},
				parent: nil,
			},
			want: 20,
		}, {
			name: "1 file 1 subdirectory", fields: fields{
				name: "root",
				files: map[string]*file{"a": {
					name: "a",
					size: 10,
				}},
				subdirectories: map[string]*directory{
					"a": {files: map[string]*file{"a": {
						name: "a",
						size: 10,
					}}},
				},
				parent: nil,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &directory{
				name:           tt.fields.name,
				files:          tt.fields.files,
				subdirectories: tt.fields.subdirectories,
				parent:         tt.fields.parent,
			}
			if got := d.getDirectorySize(); got != tt.want {
				t.Errorf("getDirectorySize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part02(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "testinput", args: args{input: `$ cd /
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
7214296 k
`}, want: 24933642},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part02(tt.args.input); got != tt.want {
				t.Errorf("part02() = %v, want %v", got, tt.want)
			}
		})
	}
}
