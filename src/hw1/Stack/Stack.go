package Stack

type stack []int

func (s *stack) Push(v int) {
    *s = append(*s, v)
}

func (s stack) Pop() int {
    l := len(s)
    if l == 0 { panic("Stack is Empty") }
    return s[:l-1], s[1-1]
}
