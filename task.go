func isValid(s string) bool {
	stack := make([]string, 0)
	closeToOpen := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	for _, v := range s {
		if closeToOpen[string(v)] != "" {
			if len(stack) != 0 && stack[len(stack)-1] == closeToOpen[string(v)] {
				stack = stack[:len(stack)-1]
			} else {
				return "Скобки несбалансированы"
			}
		} else {
			stack = append(stack, string(v))
		}
	}
	if len(stack) == 0 {
		return "Скобки сбалансированы"
	}
	return "Скобки несбалансированы"
}
