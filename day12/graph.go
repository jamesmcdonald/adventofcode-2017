package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	value int
	edges []int
}

type graphType map[int]node

func (graph *graphType) parse(input string) {
	parts := strings.Split(input, " <-> ")
	if len(parts) != 2 {
		panic(input)
	}
	val, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		panic(input)
	}
	var edges []int
	for _, e := range strings.Split(strings.TrimSpace(parts[1]), ", ") {
		edge, err := strconv.Atoi(e)
		if err != nil {
			panic(input)
		}
		edges = append(edges, edge)
	}
	graph.add(val, edges)
}

func (graph *graphType) add(value int, edges []int) {
	if _, ok := (*graph)[value]; !ok {
		(*graph)[value] = node{value: value}
	}
	n := (*graph)[value]
	for _, edge := range edges {
		n.edges = append(n.edges, edge)
	}
	(*graph)[value] = n
}

func (graph graphType) dowalk(val int, result *map[int]bool) {
	if _, ok := (*result)[val]; ok {
		return
	}
	(*result)[val] = true
	for _, e := range graph[val].edges {
		graph.dowalk(e, result)
	}
}

func (graph graphType) walk(start int) []int {
	result := make(map[int]bool)
	graph.dowalk(start, &result)
	var nodes []int
	for k := range result {
		nodes = append(nodes, k)
	}
	sort.Ints(nodes)
	return nodes
}

func (graph graphType) countgroups() int {
	var groups [][]int
	for n := range graph {
		found := false
		group := graph.walk(n)
		for _, g := range groups {
			if len(g) != len(group) {
				continue
			}
			next := false
			for i := range g {
				if g[i] != group[i] {
					next = true
					break
				}
			}
			if next {
				continue
			}
			found = true
			break
		}
		if !found {
			groups = append(groups, group)
		}
	}
	return len(groups)
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	graph := make(graphType)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		graph.parse(scanner.Text())
	}
	fmt.Println(len(graph.walk(0)))
	fmt.Println(graph.countgroups())
}
