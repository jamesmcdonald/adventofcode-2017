package main

import (
	"bytes"
	"fmt"
)

func parse(input string) {
	_, items := lex(input)
	groupDepth := 0
	groupSum := 0
	gc := 0
	var stack bytes.Buffer
	for i := range items {
		switch i.typ {
		case itemGroupStart:
			groupDepth++
			groupSum += groupDepth
			stack.WriteRune('{')
			//fmt.Printf("+ %d %d %s\n", groupDepth, groupSum, stack.String())
		case itemGroupEnd:
			groupDepth--
			stack.WriteRune('}')
			//fmt.Printf("- %d %d %s\n", groupDepth, groupSum, stack.String())
		case itemGarbage:
			stack.WriteRune('G')
			gc += len(i.val)
		case itemComma:
			stack.WriteRune(',')
		case itemError:
			fmt.Println(i)
		}
	}
	fmt.Println(groupSum, gc)
}
