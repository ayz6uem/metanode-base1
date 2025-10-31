package task1

type Stack struct {
	elements []rune
}

func (s *Stack) Push(v rune) {
	s.elements = append(s.elements, v)
}

func (s *Stack) Pop() rune {
	if len(s.elements) == 0 {
		return 0
	}
	v := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return v
}

func (s *Stack) last() rune {
	if len(s.elements) == 0 {
		return 0
	}
	return s.elements[len(s.elements)-1]
}

func ValidBracket(s string) bool {
	var stack Stack
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			stack.Push(v)
		case ')':
			if stack.Pop() != '(' {
				return false
			}
		case ']':
			if stack.Pop() != '[' {
				return false
			}
		case '}':
			if stack.Pop() != '{' {
				return false
			}
		}
	}
	return true
}
