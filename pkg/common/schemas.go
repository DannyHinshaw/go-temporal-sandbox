package common

// BadJSON is a struct that is not JSON serializable, the error type will cause issues.
type BadJSON struct {
	SomeProp string // This is fine...
	Error    error  // This will cause the activity to fail without logging
}

