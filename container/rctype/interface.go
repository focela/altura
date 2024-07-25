// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rctype

import (
	"sync/atomic"

	"github.com/focela/ratcatcher/internal/deepcopy"
	"github.com/focela/ratcatcher/internal/json"
	"github.com/focela/ratcatcher/utils/rcconv"
)

// Interface is a struct for concurrent-safe operation for type interface{}.
type Interface struct {
	value atomic.Value
}

// NewInterface creates and returns a concurrent-safe object for interface{} type,
// with given initial value `value`.
func NewInterface(value ...interface{}) *Interface {
	t := &Interface{}
	if len(value) > 0 && value[0] != nil {
		t.value.Store(value[0])
	}
	return t
}

// Clone clones and returns a new concurrent-safe object for interface{} type.
func (v *Interface) Clone() *Interface {
	return NewInterface(v.Val())
}

// Set atomically stores `value` into t.value and returns the previous value of t.value.
// Note: The parameter `value` cannot be nil.
func (v *Interface) Set(value interface{}) (old interface{}) {
	old = v.Val()
	v.value.Store(value)
	return
}

// Val atomically loads and returns t.value.
func (v *Interface) Val() interface{} {
	return v.value.Load()
}

// String implements String interface for string printing.
func (v *Interface) String() string {
	return rcconv.String(v.Val())
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (v Interface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Val())
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (v *Interface) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.UnmarshalUseNumber(b, &i); err != nil {
		return err
	}
	v.Set(i)
	return nil
}

// UnmarshalValue is an interface implement which sets any type of value for `v`.
func (v *Interface) UnmarshalValue(value interface{}) error {
	v.Set(value)
	return nil
}

// DeepCopy implements interface for deep copy of current type.
func (v *Interface) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewInterface(deepcopy.Copy(v.Val()))
}
