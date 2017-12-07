package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// Process represents a process
type Process struct {
	Name        string
	Weight      int
	TotalWeight int
	Children    []*Process
	Parent      *Process
}

// ProcessMap is an entry in the process lookup table
type ProcessMap map[string]*Process

// LoadProcess creates a Process and loads it into the Table
func (pt *ProcessMap) LoadProcess(name string, weight int, children []string) {
	process := (*pt)[name]
	if process == nil {
		process = &Process{Name: name}
		(*pt)[name] = process
	}
	process.Weight = weight
	for _, childname := range children {
		child := (*pt)[childname]
		if child == nil {
			child = &Process{Name: childname}
			(*pt)[childname] = child
		}
		child.Parent = process
		process.Children = append(process.Children, child)
	}
}

func (pt ProcessMap) String() string {
	buffer := new(bytes.Buffer)
	for name, p := range pt {
		fmt.Fprintf(buffer, "%s: %v\n", name, *p)
	}
	return buffer.String()
}

// FindRoot finds the root of the ProcessMap
func (pt ProcessMap) FindRoot() Process {
	// Pick a random entry
	k := func() string {
		for k := range pt {
			return k
		}
		return ""
	}()
	for ; pt[k].Parent != nil; k = pt[k].Parent.Name {
	}
	return *pt[k]
}

func (p *Process) childweights() string {
	buf := make([]string, len(p.Children))
	for i, c := range p.Children {
		buf[i] = fmt.Sprintf("%s:%d:%d", c.Name, c.TotalWeight, c.Weight)
	}
	return strings.Join(buf, ", ")
}

func (p *Process) calcWeight() {
	p.TotalWeight = p.Weight
	for _, c := range p.Children {
		c.calcWeight()
		p.TotalWeight += c.TotalWeight
	}
	weight := 0
	for _, c := range p.Children {
		if weight == 0 {
			weight = c.TotalWeight
		} else if weight != c.TotalWeight {
			fmt.Printf("Unbalanced: %s [%s]\n", p.Name, p.childweights())
			break
		}
	}
}

// CalcWeights works out weights for the tree
func (pt ProcessMap) CalcWeights() {
	start := pt.FindRoot()
	(&start).calcWeight()
}

func dumper(p Process, indent string) {
	fmt.Printf("%s- %s (%d)\n", indent, p.Name, p.TotalWeight)
	for _, c := range p.Children {
		dumper(*c, indent+"  ")
	}
}

// Dump the contents of the ProcessMap
func (pt ProcessMap) Dump() {
	start := pt.FindRoot()
	dumper(start, "")
}

// processDetails struct to use for input channel
type processDetails struct {
	Name     string
	Weight   int
	Children []string
}

func readfile(r io.Reader) <-chan processDetails {
	output := make(chan processDetails)
	scanner := bufio.NewScanner(r)
	go func() {
		for scanner.Scan() {
			text := scanner.Text()

			pd := processDetails{}
			fmt.Sscanf(text, "%s (%d)", &pd.Name, &pd.Weight)

			rest := strings.Split(text, "->")
			if len(rest) > 1 {
				for _, c := range strings.Split(rest[1], ",") {
					pd.Children = append(pd.Children, strings.TrimSpace(c))
				}
			}
			output <- pd
		}
		close(output)
	}()
	return output
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	pt := make(ProcessMap)
	processstream := readfile(input)
	for p := range processstream {
		pt.LoadProcess(p.Name, p.Weight, p.Children)
	}
	fmt.Println(pt.FindRoot().Name)
	pt.CalcWeights()
	pt.Dump()
}
