package err

const (
	Auth    ErrDomain = "auth"
	Assets  ErrDomain = "assets"
	Profile ErrDomain = "profile"
	Article ErrDomain = "article"
)

func (e ErrDomain) Error() string {
	return string(e)
}

type ErrApp struct {
	MapErr Mapper
}

func NewErrApp() *ErrApp {
	return &ErrApp{
		MapErr: make(Mapper),
	}
}

func (e *ErrApp) Register(domain ErrDomain, errs []ErrEntry) {
	if e.MapErr[domain] == nil {
		e.MapErr[domain] = make(map[error]ErrEntry)
	}
	for _, err := range errs {
		e.MapErr[domain][err.Err] = err
	}
}

func (e *ErrApp) Translate(err error, domain ErrDomain) *ErrEntry {
	if err == nil {
		return nil
	}
	var Err *ErrEntry
	for _, errs := range e.MapErr[domain] {
		if errs.Err == err {
			Err = &ErrEntry{
				Err:        err,
				Code:       errs.Code,
				StatusCode: errs.StatusCode,
				Message:    errs.Message,
			}
			return Err
		}
	}

	return &ErrEntry{
		Err:        InternalErr,
		Code:       InternalErrCode,
		StatusCode: 500,
		Message:    "Internal Server Error",
	}
}
