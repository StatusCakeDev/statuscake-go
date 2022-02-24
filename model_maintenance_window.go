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
	"time"
)

// MaintenanceWindow struct for MaintenanceWindow
type MaintenanceWindow struct {
	// Maintenance window ID
	ID string `json:"id"`
	// Name of the maintenance window
	Name string `json:"name"`
	// End of the maintenance window (RFC3339 format)
	End            time.Time                       `json:"end_at"`
	RepeatInterval MaintenanceWindowRepeatInterval `json:"repeat_interval"`
	// Start of the maintenance window (RFC3339 format)
	Start time.Time              `json:"start_at"`
	State MaintenanceWindowState `json:"state"`
	// List of tags used to include matching uptime checks in this maintenance window
	Tags []string `json:"tags"`
	// List of uptime check IDs explicitly included in this maintenance window
	Tests []string `json:"tests"`
	// Standard [timezone](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List) associated with this maintenance window
	Timezone string `json:"timezone"`
}

// NewMaintenanceWindow instantiates a new MaintenanceWindow object.
// This constructor will assign default values to properties that have it
// defined, and makes sure properties required by API are set, but the set of
// arguments will change when the set of required properties is changed.
func NewMaintenanceWindow(id string, name string, endAt time.Time, repeatInterval MaintenanceWindowRepeatInterval, startAt time.Time, state MaintenanceWindowState, tags []string, tests []string, timezone string) *MaintenanceWindow {
	return &MaintenanceWindow{
		ID:             id,
		Name:           name,
		End:            endAt,
		RepeatInterval: repeatInterval,
		Start:          startAt,
		State:          state,
		Tags:           tags,
		Tests:          tests,
		Timezone:       timezone,
	}
}

// Marshal data from the in the struct to JSON.
func (o MaintenanceWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.ID
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["end_at"] = o.End
	}
	if true {
		toSerialize["repeat_interval"] = o.RepeatInterval
	}
	if true {
		toSerialize["start_at"] = o.Start
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["tests"] = o.Tests
	}
	if true {
		toSerialize["timezone"] = o.Timezone
	}
	return json.Marshal(toSerialize)
}