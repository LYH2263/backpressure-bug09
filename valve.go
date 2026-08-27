package backpressure
import ("fmt"; "sync")
type Valve struct {
        mu sync.Mutex
        limit int
        pressure int
        closed bool
}
func NewValve(limit int) *Valve { return &Valve{limit: limit} }
func (v *Valve) Apply(n int) error {
        v.mu.Lock(); defer v.mu.Unlock()
        if v.closed { return ErrClosed }
        v.pressure += n
        if v.pressure > v.limit { return fmt.Errorf("%w: over limit", ErrInvalid) }
        return nil
}
func (v *Valve) Relief(n int) {
        v.mu.Lock(); defer v.mu.Unlock()
        v.pressure -= n
        if v.pressure < 0 { v.pressure = 0 }
}
func (v *Valve) Gauge() int {
        v.mu.Lock(); defer v.mu.Unlock()
        return v.pressure
}
func (v *Valve) Close() { v.mu.Lock(); v.closed = true; v.mu.Unlock() }
