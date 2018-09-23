package main

// Stack model
type Stack struct {
	container []string
}

func (s *Stack) push(value string) {
	s.container = append(s.container, value)
}

func (s *Stack) pop() {
	s.container = s.container[:len(s.container)-1]
}

func (s *Stack) front() string {
	if len(s.container) == 0 {
		return ""
	}

	return s.container[len(s.container)-1]
}

func isValid(s string) bool {

	parentheses := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	var stack Stack

	for _, symbol := range s {
		ch := string(symbol)

		if parentheses[ch] != "" && stack.front() == parentheses[ch] {
			stack.pop()
			continue
		}

		stack.push(ch)
	}

	return stack.front() == ""
}
