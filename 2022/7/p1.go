package main

import (
	"os"
	"strconv"
	"strings"
	// "regexp"
)

type FSWalker struct {
	Fs         *Directory
	CurrentLoc *Directory
}

type File struct {
	parent *Directory
	Size   int
}

type Directory struct {
	parent      *Directory
	Directories map[string]*Directory
	Files       map[string]*File
	Size        int
	Root        bool
}

type FSCommand struct {
	op, arg, output string
}

func parse(in string) *[]FSCommand {
	split := strings.Split(in, "$")[1:]

	var instructionSet []FSCommand
	for _, v := range split {
		opString, output, _ := strings.Cut(v, "\n")
		opSet := strings.Split(opString, " ")

		// need to set parsing cds and outs etc.
		var this FSCommand
		this.op = opSet[1]
		if opSet[1] != "ls" {
			this.arg = opSet[2]
		}
		this.output = output
		instructionSet = append(instructionSet, this)
	}
	return &instructionSet
}

func (fs *FSWalker) executeCommand(cmd FSCommand) {
	if cmd.op == "cd" {
		if cmd.arg == ".." {
			fs.CurrentLoc = fs.CurrentLoc.parent
		} else if cmd.arg == "/" {
			fs.CurrentLoc = fs.Fs
		} else {
			fs.CurrentLoc = fs.CurrentLoc.Directories[cmd.arg]
		}
	} else if cmd.op == "ls" {
		out := strings.Split(cmd.output, "\n")
		out = out[:len(out)-1]
		for _, line := range out {

			pair := strings.Split(line, " ")
			if pair[0] == "dir" {
				if _, ok := fs.CurrentLoc.Directories[pair[1]]; ok {
					continue
				}
				fs.CurrentLoc.Directories[pair[1]] = &Directory{
					parent:      fs.CurrentLoc,
					Directories: make(map[string]*Directory),
					Files:       make(map[string]*File),
					Root:        false,
				}
			} else {
				size, err := strconv.Atoi(pair[0])
				if err != nil {
					panic(err)
				}
				if _, ok := fs.CurrentLoc.Files[pair[1]]; ok {
					continue
				}
				fs.CurrentLoc.Files[pair[1]] = &File{
					parent: fs.CurrentLoc,
					Size:   size,
				}
				addSizesRecursively(fs.CurrentLoc, size)
			}
		}
	}
}

func addSizesRecursively(d *Directory, size int) {
	d.Size += size
	if !d.Root {
		addSizesRecursively(d.parent, size)
	}
}

func main() {
	in, _ := os.ReadFile("input")
	commandSet := parse(string(in))

	fsPtr := &Directory{
		Directories: make(map[string]*Directory),
		Files:       make(map[string]*File),
		Root:        true,
	}
	fsWalker := FSWalker{
		Fs:         fsPtr,
		CurrentLoc: nil,
	}

	for _, command := range *commandSet {
		fsWalker.executeCommand(command)
	}

	var total int
	recurseAndPrintBig(fsWalker.Fs, "root", &total)
	// js, _ := json.MarshalIndent(fsWalker.Fs, ", ", "  ")
	// fmt.Println(string(js))

	// fmt.Println(total)
}

func recurseAndPrintBig(d *Directory, name string, total *int) {
	if d.Size > 100000 {
		*total += d.Size
	}
	for k, child := range d.Directories {
		recurseAndPrintBig(child, k, total)
	}
}
