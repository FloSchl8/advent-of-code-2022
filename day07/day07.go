package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type directory struct {
	name           string
	files          map[string]*file
	subdirectories map[string]*directory
	parent         *directory
	size           int
}

type file struct {
	name string
	size int
}

func (d *directory) addFile(newFile *file) {
	if d.files != nil {
		d.files[newFile.name] = newFile
	} else {
		d.files = make(map[string]*file)
		d.files[newFile.name] = newFile
	}
}

func (d *directory) getDirectorySize() int {
	sum := 0
	for _, f := range d.files {
		sum += f.size
	}
	for _, subdirectory := range d.subdirectories {
		sum += subdirectory.getDirectorySize()
	}
	d.size = sum
	return sum
}

func (d *directory) addSubDirectory(subDirectory *directory) {
	if d.subdirectories != nil {
		d.subdirectories[subDirectory.name] = subDirectory
	} else {
		d.subdirectories = make(map[string]*directory)
		d.subdirectories[subDirectory.name] = subDirectory
	}
}

func main() {
	input, err := os.ReadFile("./day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 01", "Total size of directories not bigger than 100.000", part01(string(input)))
	fmt.Println("Part 02", "Size of smallest directory to delete", part02(string(input)))

}

func part01(input string) int {
	lines := strings.Split(input, "\n")

	root := getFileSystem(lines)

	var dirSizes = make(map[string]int)
	for s, d := range root.subdirectories {
		dirSizes["/"+s] = d.size
		getSubdirecotrySize(d, dirSizes, s)
	}

	result := 0

	for s, i := range dirSizes {
		if i <= 100000 {
			result += dirSizes[s]
		}
		fmt.Println(s, ":", dirSizes[s])
	}

	return result
}

func part02(input string) int {
	lines := strings.Split(input, "\n")

	root := getFileSystem(lines)

	neededSpace := 30000000 - (70000000 - root.size)

	var dirSizes = make(map[string]int)
	for s, d := range root.subdirectories {
		dirSizes["/"+s] = d.size
		getSubdirecotrySize(d, dirSizes, s)
	}

	result := math.MaxInt
	for _, i := range dirSizes {
		if i >= neededSpace && i <= result {
			result = i
		}
	}

	return result
}

func getFileSystem(lines []string) *directory {
	root := &directory{name: "/"}

	currentDirectory := root

	for _, line := range lines {
		if isCommand(line) {
			commandSplit := strings.Split(line, " ")
			if commandSplit[1] == "cd" {
				if len(commandSplit) == 3 && commandSplit[2] != "/" && commandSplit[2] != ".." { // '$ cd foo'
					for _, subdirectory := range currentDirectory.subdirectories {
						if subdirectory.name == commandSplit[2] {
							currentDirectory = subdirectory
						}
					}
				} else if len(commandSplit) == 3 {
					if commandSplit[2] == "/" { // '$ cd /'
						currentDirectory = root
					} else if commandSplit[2] == ".." { // '$ cd ..'
						currentDirectory = currentDirectory.parent
					}
				}
			}
		}
		if isDirectory(line) {
			sub := &directory{name: strings.Split(line, " ")[1], parent: currentDirectory}
			currentDirectory.addSubDirectory(sub)
		}

		if isFile(line) {
			f := strings.Split(line, " ")
			file := &file{
				name: f[1],
			}
			if size, _ := strconv.Atoi(f[0]); size != 0 {
				file.size = size
			}
			currentDirectory.addFile(file)
		}
	}

	root.getDirectorySize()

	return root

}

func getSubdirecotrySize(d *directory, dirSizes map[string]int, parent string) {
	for s, d := range d.subdirectories {
		dirSizes[parent+"/"+s] = d.size
		getSubdirecotrySize(d, dirSizes, s)
	}
}

func isFile(input string) bool {
	regex := regexp.MustCompile(`\d+`)
	return len(regex.FindStringIndex(input)) > 0
}

func isDirectory(input string) bool {
	return strings.HasPrefix(input, "dir")
}

func isCommand(input string) bool {
	return strings.HasPrefix(input, "$")
}
