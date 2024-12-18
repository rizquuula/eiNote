package errorcode

type ErrorCode int

const (
	Null ErrorCode = iota
	RequestParsingError
	AuthorizationError
	SanitizeError
	DatabaseError
	SystemError
	NotFoundError
)

func (r ErrorCode) String() string {
	switch r {
	case RequestParsingError:
		return "R001"
	case AuthorizationError:
		return "A001"
	case SanitizeError:
		return "S001"
	case DatabaseError:
		return "D001"
	case SystemError:
		return "S002"
	case NotFoundError:
		return "N001"
	default:
		return "E000"
	}
}
