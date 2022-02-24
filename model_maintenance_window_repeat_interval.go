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

// MaintenanceWindowRepeatInterval How often the maintenance window should occur
type MaintenanceWindowRepeatInterval string

const (
	// MaintenanceWindowRepeatIntervalNever a maintenance window that never reoccurs.
	MaintenanceWindowRepeatIntervalNever MaintenanceWindowRepeatInterval = "never"
	// MaintenanceWindowRepeatIntervalDaily a maintenance window occuring daily.
	MaintenanceWindowRepeatIntervalDaily MaintenanceWindowRepeatInterval = "1d"
	// MaintenanceWindowRepeatIntervalWeekly a maintenance window occuring weekly.
	MaintenanceWindowRepeatIntervalWeekly MaintenanceWindowRepeatInterval = "1w"
	// MaintenanceWindowRepeatIntervalBiweekly a maintenance window occuring biweekly.
	MaintenanceWindowRepeatIntervalBiweekly MaintenanceWindowRepeatInterval = "2w"
	// MaintenanceWindowRepeatIntervalMonthly a maintenance window occuring monthly.
	MaintenanceWindowRepeatIntervalMonthly MaintenanceWindowRepeatInterval = "1m"
)

// Unmarshal JSON data into any of the pointers in the type.
func (v *MaintenanceWindowRepeatInterval) UnmarshalJSON(src []byte) error {
	var value string
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}

	ev := MaintenanceWindowRepeatInterval(value)
	if !ev.Valid() {
		return fmt.Errorf("%+v is not a valid MaintenanceWindowRepeatInterval", value)
	}

	*v = ev
	return nil
}

// Valid determines if the value is valid.
func (v MaintenanceWindowRepeatInterval) Valid() bool {
	return v == MaintenanceWindowRepeatIntervalNever || v == MaintenanceWindowRepeatIntervalDaily || v == MaintenanceWindowRepeatIntervalWeekly || v == MaintenanceWindowRepeatIntervalBiweekly || v == MaintenanceWindowRepeatIntervalMonthly
}

// MaintenanceWindowRepeatIntervalValues returns the values of MaintenanceWindowRepeatInterval.
func MaintenanceWindowRepeatIntervalValues() []string {
	return []string{
		"never",
		"1d",
		"1w",
		"2w",
		"1m",
	}
}