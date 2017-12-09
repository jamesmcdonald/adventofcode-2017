package main

import (
	"fmt"
	"unicode/utf8"
)

type lexer struct {
	name  string
	input string
	start int
	pos   int
	width int
	items chan item
}

type item struct {
	typ itemType
	val string
}

func (i item) String() string {
	switch i.typ {
	case itemEOF:
		return "EOF"
	case itemError:
		return i.val
	}
	return fmt.Sprintf("%q", i.val)
}

type itemType int

const (
	itemError itemType = iota
	itemEOF
	itemGroupStart
	itemGroupEnd
	itemGarbageStart
	itemGarbage
	itemGarbageEnd
	itemComma
)

const eof = -1

type stateFn func(*lexer) stateFn

func (l *lexer) run() {
	for state := lexLookingForGroup; state != nil; {
		state = state(l)
	}
	close(l.items)
}

func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{
		itemError,
		fmt.Sprintf(format, args...),
	}
	return nil
}

func lex(input string) (*lexer, chan item) {
	l := &lexer{
		input: input,
		items: make(chan item),
	}
	go l.run()
	return l, l.items
}

const groupStart = '{'
const groupEnd = '}'
const garbageStart = '<'
const garbageEnd = '>'
const escape = '!'
const comma = ','

func lexLookingForGroup(l *lexer) stateFn {
	if l.next() != groupStart {
		return l.errorf("No group start. Sad.")
	}
	l.emit(itemGroupStart)
	return lexInsideGroup
}

func lexInsideGroup(l *lexer) stateFn {
	val := l.next()
	switch val {
	case groupStart:
		l.emit(itemGroupStart)
		return lexInsideGroup

	case garbageStart:
		l.emit(itemGarbageStart)
		return lexInsideGarbage

	case groupEnd:
		l.emit(itemGroupEnd)
		return lexAfterGroup
	}
	return l.errorf("invalid input %q", val)
}

func lexAfterGroup(l *lexer) stateFn {
	val := l.next()
	for val == '\n' {
		val = l.next()
	}
	switch val {
	case comma:
		l.emit(itemComma)
		return lexInsideGroup
	case groupEnd:
		l.emit(itemGroupEnd)
		return lexAfterGroup
	case eof:
		l.emit(itemEOF)
		return nil
	}
	return l.errorf("invalid input %q", val)
}

func lexInsideGarbage(l *lexer) stateFn {
	for {
		if l.input[l.pos] == escape {
			if l.pos > l.start {
				l.emit(itemGarbage)
			}
			return lexEscape
		}
		if l.input[l.pos] == garbageEnd {
			if l.pos > l.start {
				l.emit(itemGarbage)
			}
			l.next()
			l.emit(itemGarbageEnd)
			return lexAfterGroup
		}
		l.next()
	}
}

func lexEscape(l *lexer) stateFn {
	l.next()
	l.next()
	l.ignore()
	return lexInsideGarbage
}
