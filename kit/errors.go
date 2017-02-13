package kit

type NotFoundError struct {
	Name string
	Version string
	IsDev bool
}

func (e NotFoundError) Error() string {
	if e.IsDev {
		return "Development kit (in dev/) not found"
	}
	return "Kit " + e.Name + "/" + e.Version + " not found"
}

func IsNotFound(e error) bool {
	_, ok := e.(NotFoundError);
	return ok;
}
