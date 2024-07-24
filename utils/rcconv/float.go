// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rcconv

import (
	"strconv"

	"github.com/focela/ratcatcher/encoding/rcbinary"
)

// Float32 converts `any` to float32.
func Float32(any interface{}) float32 {
	if any == nil {
		return 0
	}
	switch value := any.(type) {
	case float32:
		return value
	case float64:
		return float32(value)
	case []byte:
		return rcbinary.DecodeToFloat32(value)
	default:
		if f, ok := value.(iFloat32); ok {
			return f.Float32()
		}
		v, _ := strconv.ParseFloat(String(any), 64)
		return float32(v)
	}
}

// Float64 converts `any` to float64.
func Float64(any interface{}) float64 {
	if any == nil {
		return 0
	}
	switch value := any.(type) {
	case float32:
		return float64(value)
	case float64:
		return value
	case []byte:
		return rcbinary.DecodeToFloat64(value)
	default:
		if f, ok := value.(iFloat64); ok {
			return f.Float64()
		}
		v, _ := strconv.ParseFloat(String(any), 64)
		return v
	}
}
