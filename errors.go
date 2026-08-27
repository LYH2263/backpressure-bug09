package backpressure
import "errors"
var (
        ErrClosed = errors.New("backpressure: closed")
        ErrInvalid = errors.New("backpressure: invalid")
        ErrNotFound = errors.New("backpressure: not found")
        ErrConflict = errors.New("backpressure: conflict")
)
