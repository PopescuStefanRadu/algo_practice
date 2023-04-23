package regexp

//
//type ConditionalReader interface {
//	ReadConditionally(rs []rune, pos int) (lastRead int, matches bool, hasMore bool)
//}
//
//type ExactReader struct {
//	expectation rune
//}
//
//func (er ExactReader) ReadConditionally(rs []rune, pos int) (lastRead int, matches bool, hasMore bool) {
//	if pos >= len(rs) {
//		return pos, false, false
//	}
//	return pos + 1, rs[pos] == er.expectation, false
//}
//
//type ZeroOrMoreReader struct {
//	reader  ConditionalReader
//	lastPos *int
//}
//
//func (zr ZeroOrMoreReader) ReadConditionally(rs []rune, pos int) (lastRead int, matches bool, hasMore bool) {
//	if zr.lastPos == nil {
//		var read int
//		var matched bool
//		for r, m, _ := zr.reader.ReadConditionally(rs, pos); m; {
//			matched = true
//			read = r
//			pos = read
//		}
//		zr.lastPos = &read
//		return read, matched, true
//	}
//	if pos == *zr.lastPos {
//		*zr.lastPos = pos
//		return pos, true, false
//	}
//	*zr.lastPos = *zr.lastPos - 1
//	return *zr.lastPos, true, false
//}
//
//type AnyReader struct{}
//
//func (_ AnyReader) ReadConditionally(rs []rune, pos int) (lastRead int, matches bool, hasMore bool) {
//	return pos + 1, pos < len(rs), false
//}
//
//func isMatch(s string, p string) bool {
//	return IsMatch(s, p)
//}
//
//func IsMatch(s string, p string) bool {
//	if p == "" {
//		return s == p
//	}
//
//	pat := []rune(p)
//	var conditions []ConditionalReader
//
//	for i := 0; i < len(pat); {
//		nextPos := i + 1
//		if nextPos > len(pat)-1 {
//			cond := newFiniteMatchConditionalReader(pat[i])
//			conditions = append(conditions, cond)
//			i++
//			continue
//		}
//		if pat[nextPos] == '*' {
//			conditions = append(conditions, newZeroOrMoreConditionalReader(pat[i]))
//			i = i + 2
//			continue
//		}
//		conditions = append(conditions, newFiniteMatchConditionalReader(pat[i]))
//		i++
//	}
//
//	rs := []rune(s)
//
//	for i, condPos := 0, 0; i < len(rs); {
//		cond := conditions[condPos]
//		read, matches, more := cond.ReadConditionally()
//		if !matches {
//			// TODO reengineer
//		}
//
//	}
//
//	return false
//}
//
//func testConditions(rs []rune, conds []ConditionalReader, currentPos int) bool {
//
//}
//
//func newFiniteMatchConditionalReader(r rune) ConditionalReader {
//	if r == '.' {
//		return AnyReader{}
//	}
//	return ExactReader{expectation: r}
//}
//
//func newZeroOrMoreConditionalReader(r rune) ConditionalReader {
//	return ZeroOrMoreReader{
//		reader: newFiniteMatchConditionalReader(r),
//	}
//}
