package backpressure

import "github.com/LYH2263/go-backpressure/internal/audit"

func (e *Engine) Audit() *audit.Logger {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.audit
}
