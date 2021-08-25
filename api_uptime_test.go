package statuscake_test

import (
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/StatusCakeDev/statuscake-go"
)

func TestCreateUptimeTest(t *testing.T) {
	t.Run("returns a created status on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expectEqual(t, mustParse(t, r), url.Values{
				"basic_pass":         []string{"password"},
				"basic_user":         []string{"username"},
				"check_rate":         []string{"3600"},
				"confirmation":       []string{"6"},
				"contact_groups_csv": []string{"123"},
				"custom_header":      []string{"Authorization: Bearer abc123"},
				"dns_ip_csv":         []string{"1.1.1.1"},
				"dns_server":         []string{"dns.statuscake.com"},
				"do_not_find":        []string{"true"},
				"enable_ssl_alert":   []string{"true"},
				"final_endpoint":     []string{"https://www.statuscake.com/redirected"},
				"find_string":        []string{"Hello, world"},
				"follow_redirects":   []string{"true"},
				"host":               []string{"AWS"},
				"include_header":     []string{"true"},
				"name":               []string{"statuscake.com"},
				"paused":             []string{"true"},
				"port":               []string{"123"},
				"post_body":          []string{`{"key":"value"}`},
				"post_raw":           []string{"key=value"},
				"regions[]":          []string{"london"},
				"status_codes_csv":   []string{"200,201"},
				"tags_csv":           []string{"testing"},
				"test_type":          []string{"HTTP"},
				"timeout":            []string{"10"},
				"trigger_rate":       []string{"2"},
				"use_jar":            []string{"true"},
				"user_agent":         []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36 OPR/73.0.3856.344"},
				"website_url":        []string{"https://www.statuscake.com"},
			})

			w.WriteHeader(http.StatusCreated)
			w.Write(mustRead(t, "testdata/create-resource-success.json"))
		}))
		defer s.Close()

		res, err := c.CreateUptimeTest(context.Background()).
			Name("statuscake.com").
			Paused(true).
			TestType(statuscake.UptimeTestTypeHTTP).
			WebsiteURL("https://www.statuscake.com").
			CheckRate(statuscake.UptimeTestCheckRateOneHour).
			ContactGroups([]string{
				"123",
			}).
			BasicPass("password").
			BasicUser("username").
			Confirmation(6).
			CustomHeader("Authorization: Bearer abc123").
			DNSIPs([]string{"1.1.1.1"}).
			DNSServer("dns.statuscake.com").
			DoNotFind(true).
			EnableSSLAlert(true).
			FinalEndpoint("https://www.statuscake.com/redirected").
			FindString("Hello, world").
			FollowRedirects(true).
			Host("AWS").
			IncludeHeader(true).
			Port(123).
			PostBody(`{"key":"value"}`).
			PostRaw("key=value").
			Regions([]string{"london"}).
			StatusCodes([]string{
				"200",
				"201",
			}).
			Tags([]string{
				"testing",
			}).
			Timeout(10).
			TriggerRate(2).
			UseJAR(true).
			UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36 OPR/73.0.3856.344").
			Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, res.Data.NewID, "2")
	})

	t.Run("returns an error if the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(mustRead(t, "testdata/invalid-website-url-error.json"))
		}))
		defer s.Close()

		_, err := c.CreateUptimeTest(context.Background()).
			Name("statuscake.com").
			TestType(statuscake.UptimeTestTypeHTTP).
			WebsiteURL("this,is,not,valid").
			CheckRate(statuscake.UptimeTestCheckRateOneHour).
			ContactGroups([]string{
				"123",
			}).
			Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusBadRequest,
			Message: "The provided parameters are invalid. Check the errors output for details information.",
			Errors: map[string][]string{
				"website_url": []string{"Website Url is not a valid URL"},
			},
		})
	})
}

func TestDeleteUptimeTest(t *testing.T) {
	t.Run("returns a no content status on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		}))
		defer s.Close()

		err := c.DeleteUptimeTest(context.Background(), "2").Execute()
		if err != nil {
			t.Fatal(err.Error())
		}
	})

	t.Run("returns an error when the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(mustRead(t, "testdata/fetch-resource-error.json"))
		}))
		defer s.Close()

		err := c.DeleteUptimeTest(context.Background(), "3").Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusNotFound,
			Message: "No results found",
			Errors:  map[string][]string{},
		})
	})
}

func TestGetUptimeTest(t *testing.T) {
	t.Run("returns an uptime test on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(mustRead(t, "testdata/get-uptime-test-success.json"))
		}))
		defer s.Close()

		test, err := c.GetUptimeTest(context.Background(), "2").Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, test.Data, statuscake.UptimeTest{
			ID:         "2",
			Name:       "statuscake.com",
			Paused:     true,
			TestType:   statuscake.UptimeTestTypeHTTP,
			WebsiteURL: "https://www.statuscake.com",
			CheckRate:  statuscake.UptimeTestCheckRateOneHour,
			ContactGroups: []string{
				"123",
			},
			Confirmation:    6,
			CustomHeader:    statuscake.PtrString("Authorization: Bearer abc123"),
			DNSIPs:          []string{"1.1.1.1"},
			DNSServer:       statuscake.PtrString("dns.statuscake.com"),
			DoNotFind:       true,
			EnableSSLAlert:  true,
			FinalEndpoint:   statuscake.PtrString("https://www.statuscake.com/redirected"),
			FindString:      statuscake.PtrString("Hello, world"),
			FollowRedirects: true,
			Host:            statuscake.PtrString("AWS"),
			LastTested:      statuscake.PtrTime(time.Date(2020, 10, 26, 13, 0, 0, 0, time.UTC)),
			Port:            statuscake.PtrInt32(123),
			PostBody:        statuscake.PtrString(`{"key":"value"}`),
			PostRaw:         statuscake.PtrString("key=value"),
			Servers: []statuscake.MonitoringLocation{
				statuscake.MonitoringLocation{
					Description: "Singapore",
					Region:      "Singapore",
					IPv4:        statuscake.PtrString("128.199.222.65"),
					RegionCode:  "singapore",
					Status:      statuscake.MonitoringLocationStatusUp,
				},
				statuscake.MonitoringLocation{
					Description: "United Kingdom, London - 5",
					Region:      "United Kingdom / London",
					IPv4:        statuscake.PtrString("178.62.78.199"),
					IPv6:        statuscake.PtrString("2a03:b0c0:1:d0::5e:f001"),
					RegionCode:  "london",
					Status:      statuscake.MonitoringLocationStatusDown,
				},
				statuscake.MonitoringLocation{
					Description: "Germany, Frankfurt - 10",
					Region:      "Germany / Frankfurt",
					IPv4:        statuscake.PtrString("139.59.152.248"),
					RegionCode:  "frankfurt",
					Status:      statuscake.MonitoringLocationStatusUp,
				},
			},
			Status: statuscake.UptimeTestStatusUp,
			StatusCodes: []string{
				"200",
				"201",
			},
			Tags: []string{
				"testing",
			},
			Timeout:     10,
			TriggerRate: 2,
			Uptime:      100,
			UseJAR:      true,
			UserAgent:   statuscake.PtrString("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.102 Safari/537.3"),
		})
	})

	t.Run("returns an error when the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(mustRead(t, "testdata/fetch-resource-error.json"))
		}))
		defer s.Close()

		_, err := c.GetUptimeTest(context.Background(), "3").Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusNotFound,
			Message: "No results found",
			Errors:  map[string][]string{},
		})
	})
}

func TestListUptimeTests(t *testing.T) {
	t.Run("returns a list of uptime tests on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(mustRead(t, "testdata/list-uptime-tests-success.json"))
		}))
		defer s.Close()

		tests, err := c.ListUptimeTests(context.Background()).Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, tests.Data, []statuscake.UptimeTestOverview{
			statuscake.UptimeTestOverview{
				ID:            "1",
				Name:          "google.com",
				Paused:        false,
				TestType:      statuscake.UptimeTestTypeHTTP,
				WebsiteURL:    "https://www.google.com",
				CheckRate:     statuscake.UptimeTestCheckRateOneMinute,
				ContactGroups: []string{},
				Status:        statuscake.UptimeTestStatusUp,
				Tags:          []string{},
				Uptime:        statuscake.PtrFloat32(0),
			},
			statuscake.UptimeTestOverview{
				ID:         "2",
				Name:       "statuscake.com",
				Paused:     true,
				TestType:   statuscake.UptimeTestTypeHTTP,
				WebsiteURL: "https://www.statuscake.com",
				CheckRate:  statuscake.UptimeTestCheckRateOneHour,
				ContactGroups: []string{
					"123",
				},
				Status: statuscake.UptimeTestStatusUp,
				Tags: []string{
					"testing",
				},
				Uptime: statuscake.PtrFloat32(100),
			},
		})
	})

	t.Run("returns an error when the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(mustRead(t, "testdata/fetch-resource-error.json"))
		}))
		defer s.Close()

		_, err := c.ListUptimeTests(context.Background()).Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusNotFound,
			Message: "No results found",
			Errors:  map[string][]string{},
		})
	})
}

func TestListUptimeTestHistory(t *testing.T) {
	t.Run("returns a list of uptime test history results on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(mustRead(t, "testdata/list-uptime-test-history-success.json"))
		}))
		defer s.Close()

		histroy, err := c.ListUptimeTestHistory(context.Background(), "2").Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, histroy.Data, map[string]statuscake.UptimeTestHistoryResult{
			"1611839465": statuscake.UptimeTestHistoryResult{
				StatusCode:  statuscake.PtrInt32(200),
				Location:    statuscake.PtrString("BR1"),
				Performance: statuscake.PtrInt64(259),
				Created:     time.Date(2021, 1, 28, 13, 11, 5, 0, time.UTC),
			},
			"1611837664": statuscake.UptimeTestHistoryResult{
				StatusCode:  statuscake.PtrInt32(200),
				Location:    statuscake.PtrString("BR1"),
				Performance: statuscake.PtrInt64(359),
				Created:     time.Date(2021, 1, 28, 12, 41, 4, 0, time.UTC),
			},
			"1611835857": statuscake.UptimeTestHistoryResult{
				StatusCode:  statuscake.PtrInt32(200),
				Location:    statuscake.PtrString("BR1"),
				Performance: statuscake.PtrInt64(269),
				Created:     time.Date(2021, 1, 28, 12, 10, 57, 0, time.UTC),
			},
		})
	})

	t.Run("returns an error when the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(mustRead(t, "testdata/fetch-resource-error.json"))
		}))
		defer s.Close()

		_, err := c.ListUptimeTestHistory(context.Background(), "3").Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusNotFound,
			Message: "No results found",
			Errors:  map[string][]string{},
		})
	})
}

func TestListSentAlerts(t *testing.T) {
	t.Run("returns a list of uptime test alerts on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(mustRead(t, "testdata/list-uptime-sent-alerts-success.json"))
		}))
		defer s.Close()

		alerts, err := c.ListSentAlerts(context.Background(), "2").Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, alerts.Data, []statuscake.UptimeTestAlert{
			statuscake.UptimeTestAlert{
				ID:         "2",
				Status:     statuscake.UptimeTestStatusDown,
				StatusCode: 404,
				Triggered:  statuscake.PtrTime(time.Date(2021, 1, 28, 13, 36, 0, 0, time.UTC)),
			},
		})
	})

	t.Run("returns an error when the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(mustRead(t, "testdata/fetch-resource-error.json"))
		}))
		defer s.Close()

		_, err := c.ListSentAlerts(context.Background(), "3").Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusNotFound,
			Message: "No results found",
			Errors:  map[string][]string{},
		})
	})
}

func TestUpdateUptimeTest(t *testing.T) {
	t.Run("returns a no content status on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expectEqual(t, mustParse(t, r), url.Values{
				"check_rate":         []string{"1800"},
				"contact_groups_csv": []string{""},
				"do_not_find":        []string{"false"},
				"enable_ssl_alert":   []string{"false"},
				"final_endpoint":     []string{""},
				"follow_redirects":   []string{"false"},
				"name":               []string{"example.com"},
				"paused":             []string{"false"},
				"regions[]":          []string{"london", "paris"},
				"status_codes_csv":   []string{"100,200,400"},
				"tags_csv":           []string{"example"},
				"use_jar":            []string{"false"},
			})

			w.WriteHeader(http.StatusNoContent)
		}))
		defer s.Close()

		err := c.UpdateUptimeTest(context.Background(), "2").
			Name("example.com").
			Paused(false).
			CheckRate(statuscake.UptimeTestCheckRateThirtyMinutes).
			ContactGroups([]string{}).
			DoNotFind(false).
			EnableSSLAlert(false).
			FinalEndpoint("").
			FollowRedirects(false).
			Regions([]string{
				"london",
				"paris",
			}).
			StatusCodes([]string{
				"100",
				"200",
				"400",
			}).
			Tags([]string{
				"example",
			}).
			UseJAR(false).
			Execute()
		if err != nil {
			t.Fatal(err.Error())
		}
	})

	t.Run("returns an error if the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(mustRead(t, "testdata/fetch-resource-error.json"))
		}))
		defer s.Close()

		err := c.UpdateUptimeTest(context.Background(), "3").
			Name("example.com").
			Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusNotFound,
			Message: "No results found",
			Errors:  map[string][]string{},
		})
	})
}

func TestCreateMaintenanceWindow(t *testing.T) {
	t.Run("returns a created status on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expectEqual(t, mustParse(t, r), url.Values{
				"end_at":          []string{"2020-07-11T03:30:00Z"},
				"name":            []string{"Monthly Maintenance"},
				"repeat_interval": []string{"1m"},
				"start_at":        []string{"2020-07-11T03:00:00Z"},
				"tags_csv":        []string{"testing"},
				"tests_csv":       []string{"5719107"},
				"timezone":        []string{"UTC"},
			})

			w.WriteHeader(http.StatusCreated)
			w.Write(mustRead(t, "testdata/create-resource-success.json"))
		}))
		defer s.Close()

		res, err := c.CreateMaintenanceWindow(context.Background()).
			Name("Monthly Maintenance").
			Start(time.Date(2020, 7, 11, 3, 0, 0, 0, time.UTC)).
			End(time.Date(2020, 7, 11, 3, 30, 0, 0, time.UTC)).
			RepeatInterval(statuscake.MaintenanceWindowRepeatIntervalMonthly).
			Tests([]string{
				"5719107",
			}).
			Tags([]string{
				"testing",
			}).
			Timezone("UTC").
			Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, res.Data.NewID, "2")
	})

	t.Run("returns an error if the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(mustRead(t, "testdata/invalid-dates-error.json"))
		}))
		defer s.Close()

		_, err := c.CreateMaintenanceWindow(context.Background()).
			Name("Monthly Maintenance").
			Start(time.Date(2020, 7, 11, 3, 0, 0, 0, time.UTC)).
			End(time.Date(2020, 6, 11, 3, 30, 0, 0, time.UTC)).
			RepeatInterval(statuscake.MaintenanceWindowRepeatIntervalMonthly).
			Tests([]string{
				"5719107",
			}).
			Tags([]string{
				"testing",
			}).
			Timezone("UTC").
			Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusBadRequest,
			Message: "Maintenance windows must start before they end.",
			Errors:  map[string][]string{},
		})
	})
}

func TestDeleteMaintenanceWindow(t *testing.T) {
	t.Run("returns a no content status on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		}))
		defer s.Close()

		err := c.DeleteMaintenanceWindow(context.Background(), "2").Execute()
		if err != nil {
			t.Fatal(err.Error())
		}
	})

	t.Run("returns an error when the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(mustRead(t, "testdata/fetch-resource-error.json"))
		}))
		defer s.Close()

		err := c.DeleteMaintenanceWindow(context.Background(), "3").Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusNotFound,
			Message: "No results found",
			Errors:  map[string][]string{},
		})
	})
}

func TestGetMaintenanceWindow(t *testing.T) {
	t.Run("returns a maintenance window on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(mustRead(t, "testdata/get-maintenance-window-success.json"))
		}))
		defer s.Close()

		window, err := c.GetMaintenanceWindow(context.Background(), "2").Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, window.Data, statuscake.MaintenanceWindow{
			ID:             "2",
			Name:           "Monthly Maintenance",
			Start:          time.Date(2020, 7, 11, 3, 0, 0, 0, time.UTC),
			End:            time.Date(2020, 7, 11, 3, 30, 0, 0, time.UTC),
			RepeatInterval: statuscake.MaintenanceWindowRepeatIntervalMonthly,
			Tests: []string{
				"5719107",
			},
			Tags: []string{
				"testing",
			},
			State:    statuscake.MaintenanceWindowStatePending,
			Timezone: "UTC",
		})
	})

	t.Run("returns an error when the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(mustRead(t, "testdata/fetch-resource-error.json"))
		}))
		defer s.Close()

		_, err := c.GetMaintenanceWindow(context.Background(), "3").Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusNotFound,
			Message: "No results found",
			Errors:  map[string][]string{},
		})
	})
}

func TestListMaintenanceWindows(t *testing.T) {
	t.Run("returns a list of maintenance windows on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(mustRead(t, "testdata/list-maintenance-windows-success.json"))
		}))
		defer s.Close()

		windows, err := c.ListMaintenanceWindows(context.Background()).Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, windows.Data, []statuscake.MaintenanceWindow{
			statuscake.MaintenanceWindow{
				ID:             "1",
				Name:           "Scheduled Maintenance",
				Start:          time.Date(2020, 11, 2, 14, 0, 0, 0, time.UTC),
				End:            time.Date(2020, 11, 3, 16, 0, 0, 0, time.UTC),
				RepeatInterval: statuscake.MaintenanceWindowRepeatIntervalNever,
				Tests: []string{
					"5522841",
				},
				Tags: []string{
					"scheduled",
				},
				State:    statuscake.MaintenanceWindowStatePending,
				Timezone: "UTC",
			},
			statuscake.MaintenanceWindow{
				ID:             "2",
				Name:           "Monthly Maintenance",
				Start:          time.Date(2020, 7, 11, 3, 0, 0, 0, time.UTC),
				End:            time.Date(2020, 7, 11, 3, 30, 0, 0, time.UTC),
				RepeatInterval: statuscake.MaintenanceWindowRepeatIntervalMonthly,
				Tests: []string{
					"5719107",
				},
				Tags: []string{
					"testing",
				},
				State:    statuscake.MaintenanceWindowStatePending,
				Timezone: "UTC",
			},
		})
	})

	t.Run("returns an error when the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(mustRead(t, "testdata/fetch-resource-error.json"))
		}))
		defer s.Close()

		_, err := c.ListMaintenanceWindows(context.Background()).Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusNotFound,
			Message: "No results found",
			Errors:  map[string][]string{},
		})
	})
}

func TestUpdateMaintenanceWindow(t *testing.T) {
	t.Run("returns a no content status on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expectEqual(t, mustParse(t, r), url.Values{
				"name":            []string{"Weekly Maintenance"},
				"repeat_interval": []string{"1w"},
			})

			w.WriteHeader(http.StatusNoContent)
		}))
		defer s.Close()

		err := c.UpdateMaintenanceWindow(context.Background(), "2").
			Name("Weekly Maintenance").
			RepeatInterval(statuscake.MaintenanceWindowRepeatIntervalWeekly).
			Execute()
		if err != nil {
			t.Fatal(err.Error())
		}
	})

	t.Run("returns an error if the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(mustRead(t, "testdata/fetch-resource-error.json"))
		}))
		defer s.Close()

		err := c.UpdateMaintenanceWindow(context.Background(), "3").
			Name("Weekly Maintenance").
			RepeatInterval(statuscake.MaintenanceWindowRepeatIntervalWeekly).
			Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusNotFound,
			Message: "No results found",
			Errors:  map[string][]string{},
		})
	})
}
