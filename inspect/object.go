package inspect

import "encoding/json"

//go:generate easyjson -all object.go

// NewObject returns a new *Object instance with it's attributes populated from
// the given input JSON data.
func NewObject(inputJSON []byte) (*Object, error) {
	var obj Object
	err := json.Unmarshal(inputJSON, &obj)

	return &obj, err
}

// Object is a representation of a Ruby heap object as exported from Ruby via
// `ObjectSpace.dump_all`.
type Object struct {
	Address    string           `json:"address"`
	ByteSize   int64            `json:"bytesize"`
	Capacity   int              `json:"capacity"`
	Class      string           `json:"class"`
	Default    string           `json:"default"`
	Embedded   bool             `json:"embedded"`
	Encoding   string           `json:"encoding"`
	Fd         int              `json:"fd"`
	File       string           `json:"file"`
	Flags      ObjectFlags      `json:"flags"`
	Frozen     bool             `json:"frozen"`
	Fstring    bool             `json:"fstring"`
	Generation int              `json:"generation"`
	ImemoType  string           `json:"imemo_type"`
	Ivars      int              `json:"ivars"`
	Length     int64            `json:"length"`
	Line       int              `json:"line"`
	MemSize    int64            `json:"memsize"`
	Method     string           `json:"method"`
	Name       string           `json:"name"`
	References ObjectReferences `json:"references"`
	Root       string           `json:"root"`
	Shared     bool             `json:"shared"`
	Size       int64            `json:"size"`
	Struct     string           `json:"struct"`
	Type       string           `json:"type"`
	Value      string           `json:"value"`
}

// Index returns a unique index for the given Object.
func (s *Object) Index() string {
	return s.Address + ":" + s.Type
}

// ObjectFlags represents the available flags on an Object.
type ObjectFlags struct {
	Marked        bool `json:"marked"`
	Old           bool `json:"old"`
	Uncollectible bool `json:"uncollectible"`
	WbProtected   bool `json:"wb_protected"`
}

// ObjectReferences represents the list of references in an Object.
type ObjectReferences []string
