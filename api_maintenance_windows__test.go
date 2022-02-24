// +build consumer

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

package statuscake_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/pact-foundation/pact-go/v2/matchers"
	. "github.com/pact-foundation/pact-go/v2/sugar"

	"github.com/StatusCakeDev/statuscake-go"
)

func TestCreateMaintenanceWindow(t *testing.T) {
	t.Run("returns a created status on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing uptime test",
				Parameters: map[string]interface{}{
					"test_id": 1,
				},
			}).
			UponReceiving("A request to create a valid maintenance window").
			WithRequest(http.MethodPost, S("/v1/maintenance-windows")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
				"Content-Type":  []Matcher{S("application/x-www-form-urlencoded")},
			}).
			WithBody("application/x-www-form-urlencoded", []byte(
				"end_at=2020-07-11T03%3A30%3A00Z&"+
					"name=Monthly+Maintenance&"+
					"repeat_interval=1m&"+
					"start_at=2020-07-11T03%3A00%3A00Z&"+
					"tags%5B%5D=testing&"+
					"tests%5B%5D=1&"+
					"timezone=UTC",
			)).
			WillRespondWith(http.StatusCreated).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": matchers.StructMatcher{
					"new_id": Like("1"),
				},
			})

		executeTest(t, func(c *statuscake.Client) error {
			res, _ := c.CreateMaintenanceWindow(context.Background()).
				Name("Monthly Maintenance").
				End(time.Date(2020, 7, 11, 3, 30, 0, 0, time.UTC)).
				RepeatInterval(statuscake.MaintenanceWindowRepeatIntervalMonthly).
				Start(time.Date(2020, 7, 11, 3, 0, 0, 0, time.UTC)).
				Tags([]string{
					"testing",
				}).
				Tests([]string{
					"1",
				}).
				Timezone("UTC").
				Execute()

			return equal(res.Data.NewID, "1")
		})
	})

	t.Run("returns an error if the request fails", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to create an invalid maintenance window").
			WithRequest(http.MethodPost, S("/v1/maintenance-windows")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
				"Content-Type":  []Matcher{S("application/x-www-form-urlencoded")},
			}).
			WithBody("application/x-www-form-urlencoded", []byte(
				"end_at=2020-07-11T03%3A30%3A00Z&"+
					"name=Monthly+Maintenance&"+
					"start_at=2020-07-11T03%3A00%3A00Z&"+
					"timezone=UTC",
			)).
			WillRespondWith(http.StatusBadRequest).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"message": Like("Both test and tags cannot be empty."),
				"errors":  matchers.StructMatcher{},
			})

		executeTest(t, func(c *statuscake.Client) error {
			_, err := c.CreateMaintenanceWindow(context.Background()).
				Name("Monthly Maintenance").
				End(time.Date(2020, 7, 11, 3, 30, 0, 0, time.UTC)).
				Start(time.Date(2020, 7, 11, 3, 0, 0, 0, time.UTC)).
				Timezone("UTC").
				Execute()

			return equal(err, statuscake.APIError{
				Status:  http.StatusBadRequest,
				Message: "Both test and tags cannot be empty.",
				Errors:  map[string][]string{},
			})
		})
	})
}

func TestDeleteMaintenanceWindow(t *testing.T) {
	t.Run("returns a no content status on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing maintenance window",
			}).
			UponReceiving("A request to delete a maintenance window").
			WithRequest(http.MethodDelete, FromProviderState("/v1/maintenance-windows/${id}", "/v1/maintenance-windows/1")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusNoContent)

		executeTest(t, func(c *statuscake.Client) error {
			return c.DeleteMaintenanceWindow(context.Background(), "1").Execute()
		})
	})

	t.Run("returns an error when the maintenance window does not exist", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to delete a maintenance window").
			WithRequest(http.MethodDelete, S("/v1/maintenance-windows/2")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusNotFound).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"message": Like("No results found"),
				"errors":  matchers.StructMatcher{},
			})

		executeTest(t, func(c *statuscake.Client) error {
			err := c.DeleteMaintenanceWindow(context.Background(), "2").Execute()
			return equal(err, statuscake.APIError{
				Status:  http.StatusNotFound,
				Message: "No results found",
				Errors:  map[string][]string{},
			})
		})
	})
}

func TestGetMaintenanceWindow(t *testing.T) {
	t.Run("returns a maintenance window on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing maintenance window and uptime test",
			}).
			UponReceiving("A request to get a maintenance window").
			WithRequest(http.MethodGet, FromProviderState("/v1/maintenance-windows/${id}", "/v1/maintenance-windows/1")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusOK).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": matchers.StructMatcher{
					"id":              FromProviderState("${id}", "1"),
					"name":            Like("Monthly Maintenance"),
					"end_at":          Timestamp(),
					"repeat_interval": Like("1m"),
					"start_at":        Timestamp(),
					"state":           Like("pending"),
					"tags":            EachLike("testing", 1),
					"tests":           EachLike("1", 1),
					"timezone":        "UTC",
				},
			})

		executeTest(t, func(c *statuscake.Client) error {
			window, _ := c.GetMaintenanceWindow(context.Background(), "1").Execute()
			return equal(window.Data, statuscake.MaintenanceWindow{
				ID:             "1",
				Name:           "Monthly Maintenance",
				End:            time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
				RepeatInterval: statuscake.MaintenanceWindowRepeatIntervalMonthly,
				Start:          time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
				State:          statuscake.MaintenanceWindowStatePending,
				Tags: []string{
					"testing",
				},
				Tests: []string{
					"1",
				},
				Timezone: "UTC",
			})
		})
	})

	t.Run("returns an error when the maintenance window does not exist", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to get a maintenance windoe").
			WithRequest(http.MethodGet, S("/v1/maintenance-windows/2")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusNotFound).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"message": Like("No results found"),
				"errors":  matchers.StructMatcher{},
			})

		executeTest(t, func(c *statuscake.Client) error {
			_, err := c.GetMaintenanceWindow(context.Background(), "2").Execute()
			return equal(err, statuscake.APIError{
				Status:  http.StatusNotFound,
				Message: "No results found",
				Errors:  map[string][]string{},
			})
		})
	})
}

func TestListMaintenanceWindows(t *testing.T) {
	t.Run("returns a list of maintenance windows on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "Existing maintenance windows and uptime test",
			}).
			UponReceiving("A request to get a list of maintenance windows").
			WithRequest(http.MethodGet, S("/v1/maintenance-windows")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusOK).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": EachLike(
					matchers.StructMatcher{
						"id":              FromProviderState("${id}", "1"),
						"name":            Like("Monthly Maintenance"),
						"end_at":          Timestamp(),
						"repeat_interval": Like("1m"),
						"start_at":        Timestamp(),
						"state":           Like("pending"),
						"tags":            EachLike("testing", 1),
						"tests":           EachLike("1", 1),
						"timezone":        "UTC",
					}, 1,
				),
				"metadata": matchers.StructMatcher{
					"page":        Like(1),
					"per_page":    Like(25),
					"page_count":  Like(1),
					"total_count": Like(5),
				},
			})

		executeTest(t, func(c *statuscake.Client) error {
			windows, _ := c.ListMaintenanceWindows(context.Background()).Execute()
			return equal(windows.Data, []statuscake.MaintenanceWindow{
				statuscake.MaintenanceWindow{
					ID:             "1",
					Name:           "Monthly Maintenance",
					End:            time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
					RepeatInterval: statuscake.MaintenanceWindowRepeatIntervalMonthly,
					Start:          time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
					State:          statuscake.MaintenanceWindowStatePending,
					Tags: []string{
						"testing",
					},
					Tests: []string{
						"1",
					},
					Timezone: "UTC",
				},
			})
		})
	})

	t.Run("returns an empty list when there are no maintenance windows", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to get a list of maintenance windows").
			WithRequest(http.MethodGet, S("/v1/maintenance-windows")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusOK).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": Like([]interface{}{}),
				"metadata": matchers.StructMatcher{
					"page":        Like(1),
					"per_page":    Like(25),
					"page_count":  Like(1),
					"total_count": 0,
				},
			})

		executeTest(t, func(c *statuscake.Client) error {
			windows, _ := c.ListMaintenanceWindows(context.Background()).Execute()
			return equal(windows.Data, []statuscake.MaintenanceWindow{})
		})
	})
}

func TestUpdateMaintenanceWindow(t *testing.T) {
	t.Run("returns a no content status on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing maintenance window",
			}).
			UponReceiving("A request to update a maintenance window").
			WithRequest(http.MethodPut, FromProviderState("/v1/maintenance-windows/${id}", "/v1/maintenance-windows/1")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
				"Content-Type":  []Matcher{S("application/x-www-form-urlencoded")},
			}).
			WithBody("application/x-www-form-urlencoded", []byte(
				"name=Weekly+Maintenance&"+
					"repeat_interval=1w",
			)).
			WillRespondWith(http.StatusNoContent)

		executeTest(t, func(c *statuscake.Client) error {
			return c.UpdateMaintenanceWindow(context.Background(), "1").
				Name("Weekly Maintenance").
				RepeatInterval(statuscake.MaintenanceWindowRepeatIntervalWeekly).
				Execute()
		})
	})

	t.Run("returns an error if the request fails", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing maintenance window",
			}).
			UponReceiving("A request to update an invalid maintenance window").
			WithRequest(http.MethodPut, FromProviderState("/v1/maintenance-windows/${id}", "/v1/maintenance-windows/1")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
				"Content-Type":  []Matcher{S("application/x-www-form-urlencoded")},
			}).
			WithBody("application/x-www-form-urlencoded", []byte(
				"tags%5B%5D=&"+
					"tests%5B%5D=",
			)).
			WillRespondWith(http.StatusBadRequest).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"message": Like("Both test and tags cannot be empty."),
				"errors":  matchers.StructMatcher{},
			})

		executeTest(t, func(c *statuscake.Client) error {
			err := c.UpdateMaintenanceWindow(context.Background(), "1").
				Tags([]string{}).
				Tests([]string{}).
				Execute()

			return equal(err, statuscake.APIError{
				Status:  http.StatusBadRequest,
				Message: "Both test and tags cannot be empty.",
				Errors:  map[string][]string{},
			})
		})
	})

	t.Run("returns an error when the maintenance window does not exist", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to update a maintenance window").
			WithRequest(http.MethodPut, S("/v1/maintenance-windows/2")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
				"Content-Type":  []Matcher{S("application/x-www-form-urlencoded")},
			}).
			WithBody("application/x-www-form-urlencoded", []byte(
				"name=Weekly+Maintenance&"+
					"repeat_interval=1w",
			)).
			WillRespondWith(http.StatusNotFound).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"message": Like("No results found"),
				"errors":  matchers.StructMatcher{},
			})

		executeTest(t, func(c *statuscake.Client) error {
			err := c.UpdateMaintenanceWindow(context.Background(), "2").
				Name("Weekly Maintenance").
				RepeatInterval(statuscake.MaintenanceWindowRepeatIntervalWeekly).
				Execute()

			return equal(err, statuscake.APIError{
				Status:  http.StatusNotFound,
				Message: "No results found",
				Errors:  map[string][]string{},
			})
		})
	})
}