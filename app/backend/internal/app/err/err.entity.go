package err

type ErrDomain string
type ErrCode string
type ErrEntry struct {
	Err        error
	Code       ErrCode
	StatusCode int
	Message    string
}
type Mapper map[ErrDomain]map[error]ErrEntry
