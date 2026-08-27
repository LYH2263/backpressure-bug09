package backpressure
type Stage struct {
        Name string
        Valve *Valve
        Next *Stage
}
func LinkStages(names []string, limit int) *Stage {
        if len(names) == 0 { return nil }
        head := &Stage{Name: names[0], Valve: NewValve(limit)}
        cur := head
        for _, n := range names[1:] {
                cur.Next = &Stage{Name: n, Valve: NewValve(limit)}
                cur = cur.Next
        }
        return head
}
func (s *Stage) Walk(fn func(*Stage)) {
        for cur := s; cur != nil; cur = cur.Next { fn(cur) }
}
