// Code generated by "stringer -type=errorType"; DO NOT EDIT

package types

import "fmt"

const _errorType_name = "EmptyTensorIndexErrorShapeMismatchDtypeMismatchDimensionMismatchSizeMismatchAxisErrorIOErrorInvalidCmpOpOpErrorNotYetImplemented"

var _errorType_index = [...]uint8{0, 11, 21, 34, 47, 64, 76, 85, 92, 104, 111, 128}

func (i errorType) String() string {
	if i >= errorType(len(_errorType_index)-1) {
		return fmt.Sprintf("errorType(%d)", i)
	}
	return _errorType_name[_errorType_index[i]:_errorType_index[i+1]]
}