/*
 * StatusCake API
 *
 * Copyright (c) 2023
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to
 * deal in the Software without restriction, including without limitation the
 * rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
 * sell copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
 * IN THE SOFTWARE.
 *
 * API version: 1.2.0
 * Contact: support@statuscake.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package statuscake

import (
	"encoding/json"
	"fmt"
)

// HeartbeatTestStatus The returned status of a heartbeat check
type HeartbeatTestStatus string

const (
	// HeartbeatTestStatusDown a heartbeat check with a down status.
	HeartbeatTestStatusDown HeartbeatTestStatus = "down"
	// HeartbeatTestStatusUp a heartbeat check with an up status.
	HeartbeatTestStatusUp HeartbeatTestStatus = "up"
)

// Unmarshal JSON data into any of the pointers in the type.
func (v *HeartbeatTestStatus) UnmarshalJSON(src []byte) error {
	var value string
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}

	ev := HeartbeatTestStatus(value)
	if !ev.Valid() {
		return fmt.Errorf("%+v is not a valid HeartbeatTestStatus", value)
	}

	*v = ev
	return nil
}

// Valid determines if the value is valid.
func (v HeartbeatTestStatus) Valid() bool {
	return v == HeartbeatTestStatusDown || v == HeartbeatTestStatusUp
}

// HeartbeatTestStatusValues returns the values of HeartbeatTestStatus.
func HeartbeatTestStatusValues() []string {
	return []string{
		"down",
		"up",
	}
}