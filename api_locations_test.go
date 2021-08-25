package statuscake_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/StatusCakeDev/statuscake-go"
)

func TestListPagespeedMonitoringLocations(t *testing.T) {
	t.Run("returns a list of monitoring locations on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(mustRead(t, "testdata/list-pagespeed-locations-success.json"))
		}))
		defer s.Close()

		locations, err := c.ListPagespeedMonitoringLocations(context.Background()).Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, locations.Data, []statuscake.MonitoringLocation{
			statuscake.MonitoringLocation{
				Description: "Google Chrome 72.0.3626.121",
				Region:      "United Kingdom",
				IPv4:        statuscake.PtrString("178.62.47.83"),
				IPv6:        statuscake.PtrString("2a03:b0c0:1:d0::a4:6001"),
				RegionCode:  "United Kingdom",
				Status:      statuscake.MonitoringLocationStatusUp,
			},
			statuscake.MonitoringLocation{
				Description: "Google Chrome 72.0.3626.121",
				Region:      "United Kingdom",
				IPv4:        statuscake.PtrString("46.101.86.253"),
				IPv6:        statuscake.PtrString("2a03:b0c0:1:a1::3a5:3001"),
				RegionCode:  "United Kingdom",
				Status:      statuscake.MonitoringLocationStatusUp,
			},
			statuscake.MonitoringLocation{
				Description: "Google Chrome 72.0.3626.121",
				Region:      "United Kingdom",
				IPv4:        statuscake.PtrString("188.166.170.233"),
				RegionCode:  "United Kingdom",
				Status:      statuscake.MonitoringLocationStatusUp,
			},
		})
	})
}

func TestListUptimeMonitoringLocations(t *testing.T) {
	t.Run("returns a list of monitoring locations on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(mustRead(t, "testdata/list-uptime-locations-success.json"))
		}))
		defer s.Close()

		locations, err := c.ListUptimeMonitoringLocations(context.Background()).Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, locations.Data, []statuscake.MonitoringLocation{
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
		})
	})
}
