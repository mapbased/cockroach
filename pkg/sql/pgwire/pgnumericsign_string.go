// Code generated by "stringer -type=pgNumericSign"; DO NOT EDIT.

package pgwire

import "fmt"

const (
	_pgNumericSign_name_0 = "pgNumericPos"
	_pgNumericSign_name_1 = "pgNumericNeg"
)

var (
	_pgNumericSign_index_0 = [...]uint8{0, 12}
	_pgNumericSign_index_1 = [...]uint8{0, 12}
)

func (i pgNumericSign) String() string {
	switch {
	case i == 0:
		return _pgNumericSign_name_0
	case i == 16384:
		return _pgNumericSign_name_1
	default:
		return fmt.Sprintf("pgNumericSign(%d)", i)
	}
}
