package main

import (
	"fmt"
)

func parse(input string) {
	_, items := lex(input)
	groupDepth := 0
	groupSum := 0
	gc := 0
	for i := range items {
		switch i.typ {
		case itemGroupStart:
			groupDepth++
			groupSum += groupDepth
		case itemGroupEnd:
			groupDepth--
		case itemGarbage:
			gc += len(i.val)
		case itemError:
			fmt.Println(i)
		}
	}
	fmt.Println(groupSum, gc)
}
