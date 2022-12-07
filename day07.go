package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node interface {
}

type file struct {
	name   string
	size   int
	parent *dir
}

func (f *file) print(prefix string) {
	fmt.Printf("%s - %s (file, %d)\n", prefix, f.name, f.size)
}

type dir struct {
	name      string
	childDirs map[string]*dir
	files     map[string]*file
	size      int
	parent    *dir
}

func (d *dir) print(prefix string) {
	fmt.Printf("%s - %s (dir)\n", prefix, d.name)
	p := prefix + " "
	for _, v := range d.childDirs {
		v.print(p)
	}
	for _, f := range d.files {
		f.print(p)
	}
}

func (d *dir) calSize() int {
	for _, ch := range d.childDirs {
		d.size += ch.calSize()
	}
	for _, f := range d.files {
		d.size += f.size
	}
	return d.size
}

func (d *dir) getAtMost(n int) []*dir {
	var res []*dir
	for _, ch := range d.childDirs {
		res = append(res, ch.getAtMost(n)...)
	}

	if d.size <= n {
		res = append(res, d)
	}
	return res
}

func (d *dir) getAtLeast(n int) []*dir {
	var res []*dir
	for _, ch := range d.childDirs {
		res = append(res, ch.getAtLeast(n)...)
	}

	if d.size >= n {
		res = append(res, d)
	}
	return res
}

func Day07_FirstAnswer(inputs []string, n int) int {
	rootDir := day07_buildFilesystem(inputs)
	// rootDir.print("")
	rootDir.calSize()
	dirs := rootDir.getAtMost(n)
	var totalSize int
	for _, d := range dirs {
		totalSize += d.size
	}
	return totalSize
}

func Day07_SecondAnswer(inputs []string) int {
	totalSpace := 70000000
	rootDir := day07_buildFilesystem(inputs)
	// rootDir.print("")
	rootDir.calSize()
	unusedSpace := totalSpace - rootDir.size
	needSpace := 30000000 - unusedSpace
	dirs := rootDir.getAtLeast(needSpace)
	smallest := totalSpace
	for _, d := range dirs {
		if d.size < smallest {
			smallest = d.size
		}
	}
	return smallest
}

func day07_buildFilesystem(inputs []string) *dir {
	// command
	// cd name | .. | /
	// ls print dir and files in the current dir
	// dir name  (directory)
	// <size> name (file)
	var rootDir *dir
	var currDir *dir
	for _, input := range inputs {
		if strings.HasPrefix(input, "$ cd ..") {
			currDir = currDir.parent
		} else if strings.HasPrefix(input, "$ cd /") {
			rootDir = &dir{name: "/", childDirs: make(map[string]*dir), files: make(map[string]*file)}
			currDir = rootDir
		} else if strings.HasPrefix(input, "$ cd") {
			dirName := strings.Split(input, " ")[2]
			child, ok := currDir.childDirs[dirName]
			if !ok {
				child = &dir{name: dirName, parent: currDir, childDirs: make(map[string]*dir), files: make(map[string]*file)}
				currDir.childDirs[dirName] = child
			}
			currDir = child
		} else if strings.HasPrefix(input, "$ ls") {

		} else if strings.HasPrefix(input, "dir") {
			dirName := strings.Split(input, " ")[1]
			child, ok := currDir.childDirs[dirName]
			if !ok {
				child = &dir{name: dirName, parent: currDir, childDirs: make(map[string]*dir), files: make(map[string]*file)}
				currDir.childDirs[dirName] = child
			}
		} else {
			size, _ := strconv.Atoi(strings.Split(input, " ")[0])
			fileName := strings.Split(input, " ")[1]
			newFile := &file{name: fileName, size: size, parent: currDir}
			currDir.files[fileName] = newFile
		}

	}

	return rootDir
}
