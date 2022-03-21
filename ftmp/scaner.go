package ftmp

import (
	"bufio"
	"io"
)

func NewScaner(r io.Reader, filePath string) *Scanner {
	return &Scanner {
		r:			bufio.NewReader(r),
		filePath:	filePath,
	}
}

func (s *Scanner) Next() bool {
	if s.rewind {
		s.rewind = false
		return true
	}

	for {
		if !s.scanTocken() {
			return false
		}
	}

	return true
}

func (s *Scanner) scanTocken() bool {
	switch s.nextTokenID {
	case Text:
		return s.readText()
	case TagName:
		return s.ReadTagName()
	case TagContents:
		return s.readTagContents()
	default:
		panic("unknown nextTockenID")
	}
}

func (s *Scanner) ReadTagName() bool

func (s *Scanner) readTagContents() bool

func (s *Scanner) readText() bool {
	s.t.init(Text, s.line, len(s.lineStr))
	ok := false
	for {
		if !s.nextByte() {
			ok = (len(s.t.Value) > 0)
			break
		}
		if s.c != '{' {
			s.appendByte()
			continue
		}
		if !s.nextByte() {
			s.appendByte()
			ok = true
			break
		}
		if s.c == '%' {
			s.nextTokenID = TagName
			ok = true
			if !s.nextByte() {
				s.appendByte()
				break
			}
			if s.c != '-' {
				s.unreadByte(s.c)
				break
			}
			s.t.Value = prevBlank.ReplaceAll(s.t.Value, nil)
			break
		}
		s.unreadByte('{')
		s.appendByte()
	}
	if s.stripSpaceDepth > 0 {
		s.t.Value = stripSpace(s.t.Value)
	} else if s.collapseSpaceDepth > 0 {
		s.t.Value = collapseSpace(s.t.Value)
	}
	return ok
}

func (s *Scanner) nextByte() bool {

	return true
}

func (s *Scanner) appendByte() {
	s.t.Value = append(s.t.Value, s.c)
}

func (t *Token) init(id, line, pos int) {
	t.ID = id
	t.Value = t.Value[:0]
	t.line = line
	t.pos = pos
}