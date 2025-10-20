package gotest

type SampleStruct struct {
	Field1 string
	Field2 int
}

type ErrSample struct {
	Sample string
}

func (e *ErrSample) Error() string {
	return e.Sample
}

func (s *SampleStruct) AMethod() (bool, error) {
	// Simulate some logic
	if s.Field1 == "METHOD_ERROR" {
		return false, &ErrSample{Sample: "YOU GOT AN ERROR"}
	}
	return true, nil
}
