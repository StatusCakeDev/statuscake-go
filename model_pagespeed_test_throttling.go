/*
 * StatusCake API
 *
 * Copyright (c) 2022
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
 * API version: 1.0.0-beta.2
 * Contact: support@statuscake.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package statuscake

import (
	"encoding/json"
	"fmt"
)

// PagespeedTestThrottling Simulated throttling speed
type PagespeedTestThrottling string

const (
	// PagespeedTestThrottlingNone no throttling.
	PagespeedTestThrottlingNone PagespeedTestThrottling = "NONE"
	// PagespeedTestThrottlingFast3G fast 3G.
	PagespeedTestThrottlingFast3G PagespeedTestThrottling = "3G_FAST"
	// PagespeedTestThrottlingSlow3G slow 3G.
	PagespeedTestThrottlingSlow3G PagespeedTestThrottling = "3G_SLOW"
	// PagespeedTestThrottling4G 4G.
	PagespeedTestThrottling4G PagespeedTestThrottling = "4G"
	// PagespeedTestThrottlingEDGE EDGE.
	PagespeedTestThrottlingEDGE PagespeedTestThrottling = "EDGE"
	// PagespeedTestThrottlingGPRS GPRS.
	PagespeedTestThrottlingGPRS PagespeedTestThrottling = "GPRS"
)

// Unmarshal JSON data into any of the pointers in the type.
func (v *PagespeedTestThrottling) UnmarshalJSON(src []byte) error {
	var value string
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}

	ev := PagespeedTestThrottling(value)
	if !ev.Valid() {
		return fmt.Errorf("%+v is not a valid PagespeedTestThrottling", value)
	}

	*v = ev
	return nil
}

// Valid determines if the value is valid.
func (v PagespeedTestThrottling) Valid() bool {
	return v == PagespeedTestThrottlingNone || v == PagespeedTestThrottlingFast3G || v == PagespeedTestThrottlingSlow3G || v == PagespeedTestThrottling4G || v == PagespeedTestThrottlingEDGE || v == PagespeedTestThrottlingGPRS
}

// PagespeedTestThrottlingValues returns the values of PagespeedTestThrottling.
func PagespeedTestThrottlingValues() []string {
	return []string{
		"NONE",
		"3G_FAST",
		"3G_SLOW",
		"4G",
		"EDGE",
		"GPRS",
	}
}