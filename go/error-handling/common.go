package erratum

import "io"

const testVersion = 2

type TransientError struct {
	err error
}

func (e TransientError) Error() string {
	return e.err.Error()
}

type FrobError struct {
	defrobTag string
	inner     error
}

func (e FrobError) Error() string {
	return e.inner.Error()
}

type Resource interface {
	io.Closer

	// Frob does something with the input string.
	// Because this is an incredibly badly designed system if there is an error
	// it panics.
	//
	// The paniced error may be a FrobError in which case Defrob should be called
	// with the defrobTag string.
	Frob(string)

	Defrob(string)
}

// ResourceOpener is a function that creates a resource.
//
// It may return a wrapped error of type TransientError. In this case the resource
// is temporarily unavailable and the caller should retry soon.
type ResourceOpener func() (Resource, error)

func Use(ro ResourceOpener, input string) (ret error) {
	resource, err := ro()
	for err != nil {
		if _, ok := err.(TransientError); ok {
			resource, err = ro()

		} else {
			return err
		}
	}
	defer resource.Close()

	defer func() {
		if err := recover(); err != nil {
			ret = err.(error)

			if frobError, ok := err.(FrobError); ok {
				resource.Defrob(frobError.defrobTag)
			}
		}
	}()

	resource.Frob(input)

	return
}
