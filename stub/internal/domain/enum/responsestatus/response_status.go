package responsestatus

type ResponseStatus int

const (
	Unknown ResponseStatus = iota
	OK
	BusinessError
	SystemError
)

func (r ResponseStatus) String() string {
	switch r {
	case OK:
		return "OK"
	case BusinessError:
		return "FAILED"
	case SystemError:
		return "ERR"
	default:
		return "Unknown"
	}
}
