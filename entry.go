package main

func NewEntry(inputJSON []byte) (*Entry, error) {
	obj, err := NewObject(inputJSON)
	if err != nil {
		return nil, err
	}

	return &Entry{
		Object: obj,
		Index:  obj.Index(),
	}, err
}

// Entry is a parsed heap item object
type Entry struct {
	Object *Object
	Offset int64
	Index  string
}

func (s *Entry) Address() string {
	return s.Object.Address
}
