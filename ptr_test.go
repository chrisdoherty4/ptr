package ptr_test

import "github.com/chrisdoherty4/ptr"

// Examples

var _ *int = ptr.Int(1)

var _ *int = ptr.To(1)

type Custom struct {
	I int
}

var _ *Custom = ptr.To(Custom{})
