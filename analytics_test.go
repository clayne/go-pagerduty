package pagerduty

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestAnalytics_GetAggregatedIncidentData(t *testing.T) {
	setup()
	defer teardown()

	analyticsRequest := AnalyticsRequest{
		Filters: &AnalyticsFilter{
			CreatedAtStart: "2021-01-01T15:00:32Z",
			CreatedAtEnd:   "2021-01-08T15:00:32Z",
			TeamIDs:        []string{"PCDYDX0"},
		},
		AggregateUnit: "day",
		TimeZone:      "Etc/UTC",
	}
	analyticsDataWanted := AnalyticsData{MeanSecondsToResolve: 34550, MeanSecondsToFirstAck: 70, MeanEngagedSeconds: 502, MeanAssignmentCount: 1, TotalBusinessHourInterruptions: 1, TotalEngagedSeconds: 2514, TotalIncidentCount: 5, RangeStart: "2021-01-06T00:00:00.000000"}
	analyticsFilterWanted := AnalyticsFilter{CreatedAtStart: "2021-01-06T09:21:41Z", CreatedAtEnd: "2021-01-13T09:21:41Z", TeamIDs: []string{"PCDYDX0"}}
	analyticsResponse := AnalyticsResponse{
		Data:          []AnalyticsData{analyticsDataWanted},
		Filters:       &analyticsFilterWanted,
		AggregateUnit: "day",
		TimeZone:      "Etc/UTC",
	}

	bytesAnalyticsResponse, err := json.Marshal(analyticsResponse)
	testErrCheck(t, "json.Marshal()", "", err)

	mux.HandleFunc("/analytics/metrics/incidents/all", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, _ = w.Write(bytesAnalyticsResponse)
	})

	client := defaultTestClient(server.URL, "foo")

	res, err := client.GetAggregatedIncidentData(context.Background(), analyticsRequest)
	want := AnalyticsResponse{
		Data:          []AnalyticsData{analyticsDataWanted},
		Filters:       &analyticsFilterWanted,
		AggregateUnit: "day",
		TimeZone:      "Etc/UTC",
	}
	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, want, res)
}

func TestAnalytics_GetAggregatedServiceData(t *testing.T) {
	setup()
	defer teardown()

	analyticsRequest := AnalyticsRequest{
		Filters: &AnalyticsFilter{
			CreatedAtStart: "2021-01-01T15:00:32Z",
			CreatedAtEnd:   "2021-01-08T15:00:32Z",
			TeamIDs:        []string{"PCDYDX0"},
		},
		AggregateUnit: "day",
		TimeZone:      "Etc/UTC",
	}
	analyticsDataWanted := AnalyticsData{MeanAssignmentCount: 1, MeanEngagedSeconds: 502, MeanEngagedUserCount: 0, MeanSecondsToResolve: 34550, MeanSecondsToFirstAck: 70, TotalBusinessHourInterruptions: 1, TotalEngagedSeconds: 2514, TotalIncidentCount: 5, RangeStart: "2021-01-06T00:00:00.000000", ServiceID: "PSEJLIN", ServiceName: "FooAlerts", TeamID: "PCDYDX0", TeamName: "FooTeam", UpTimePct: 89.86111111111111}
	analyticsFilterWanted := AnalyticsFilter{CreatedAtStart: "2021-01-06T09:21:41Z", CreatedAtEnd: "2021-01-13T09:21:41Z", TeamIDs: []string{"PCDYDX0"}}
	analyticsResponse := AnalyticsResponse{
		Data:          []AnalyticsData{analyticsDataWanted},
		Filters:       &analyticsFilterWanted,
		AggregateUnit: "day",
		TimeZone:      "Etc/UTC",
	}
	bytesAnalyticsResponse, err := json.Marshal(analyticsResponse)
	testErrCheck(t, "json.Marshal()", "", err)

	mux.HandleFunc("/analytics/metrics/incidents/services", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, _ = w.Write(bytesAnalyticsResponse)
	})

	client := defaultTestClient(server.URL, "foo")

	res, err := client.GetAggregatedServiceData(context.Background(), analyticsRequest)
	want := AnalyticsResponse{
		Data:          []AnalyticsData{analyticsDataWanted},
		Filters:       &analyticsFilterWanted,
		AggregateUnit: "day",
		TimeZone:      "Etc/UTC",
	}
	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, want, res)
}

func TestAnalytics_GetAggregatedTeamData(t *testing.T) {
	setup()
	defer teardown()

	analyticsRequest := AnalyticsRequest{
		Filters: &AnalyticsFilter{
			CreatedAtStart: "2021-01-01T15:00:32Z",
			CreatedAtEnd:   "2021-01-08T15:00:32Z",
			TeamIDs:        []string{"PCDYDX0"},
		},
		AggregateUnit: "day",
		TimeZone:      "Etc/UTC",
	}
	analyticsDataWanted := AnalyticsData{MeanAssignmentCount: 1, MeanEngagedSeconds: 502, MeanEngagedUserCount: 0, MeanSecondsToResolve: 34550, MeanSecondsToFirstAck: 70, TotalBusinessHourInterruptions: 1, TotalEngagedSeconds: 2514, TotalIncidentCount: 5, RangeStart: "2021-01-06T00:00:00.000000", TeamID: "PCDYDX0", TeamName: "FooTeam", UpTimePct: 89.86111111111111}
	analyticsFilterWanted := AnalyticsFilter{CreatedAtStart: "2021-01-06T09:21:41Z", CreatedAtEnd: "2021-01-13T09:21:41Z", TeamIDs: []string{"PCDYDX0"}}
	analyticsResponse := AnalyticsResponse{
		Data:          []AnalyticsData{analyticsDataWanted},
		Filters:       &analyticsFilterWanted,
		AggregateUnit: "day",
		TimeZone:      "Etc/UTC",
	}
	bytesAnalyticsResponse, err := json.Marshal(analyticsResponse)
	testErrCheck(t, "json.Marshal()", "", err)

	mux.HandleFunc("/analytics/metrics/incidents/teams", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, _ = w.Write(bytesAnalyticsResponse)
	})

	client := defaultTestClient(server.URL, "foo")

	res, err := client.GetAggregatedTeamData(context.Background(), analyticsRequest)
	want := AnalyticsResponse{
		Data:          []AnalyticsData{analyticsDataWanted},
		Filters:       &analyticsFilterWanted,
		AggregateUnit: "day",
		TimeZone:      "Etc/UTC",
	}
	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, want, res)
}

func TestAnalytics_GetAggregatedEscalationPolicyData(t *testing.T) {
	setup()
	defer teardown()

	analyticsRequest := AnalyticsRequest{
		Filters: &AnalyticsFilter{
			CreatedAtStart:      "2021-01-01T15:00:32Z",
			CreatedAtEnd:        "2021-01-08T15:00:32Z",
			EscalationPolicyIDs: []string{"PCDYDX0"},
		},
		AggregateUnit: "day",
		TimeZone:      "Etc/UTC",
	}
	analyticsDataWanted := AnalyticsData{MeanAssignmentCount: 1, MeanEngagedSeconds: 502, MeanEngagedUserCount: 0, MeanSecondsToResolve: 34550, MeanSecondsToFirstAck: 70, TotalBusinessHourInterruptions: 1, TotalEngagedSeconds: 2514, TotalIncidentCount: 5, RangeStart: "2021-01-06T00:00:00.000000", EscalationPolicyID: "PCDYDX0", EscalationPolicyName: "FooEscalationPolicy", UpTimePct: 89.86111111111111}
	analyticsFilterWanted := AnalyticsFilter{CreatedAtStart: "2021-01-06T09:21:41Z", CreatedAtEnd: "2021-01-13T09:21:41Z", EscalationPolicyIDs: []string{"PCDYDX0"}}
	analyticsResponse := AnalyticsResponse{
		Data:          []AnalyticsData{analyticsDataWanted},
		Filters:       &analyticsFilterWanted,
		AggregateUnit: "day",
		TimeZone:      "Etc/UTC",
	}
	bytesAnalyticsResponse, err := json.Marshal(analyticsResponse)
	testErrCheck(t, "json.Marshal()", "", err)

	mux.HandleFunc("/analytics/metrics/incidents/escalation_policies", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, _ = w.Write(bytesAnalyticsResponse)
	})

	client := defaultTestClient(server.URL, "foo")

	res, err := client.GetAggregatedEscalationPolicyData(context.Background(), analyticsRequest)
	want := AnalyticsResponse{
		Data:          []AnalyticsData{analyticsDataWanted},
		Filters:       &analyticsFilterWanted,
		AggregateUnit: "day",
		TimeZone:      "Etc/UTC",
	}
	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, want, res)
}

func TestAnalytics_GetAnalyticsIncidentsById(t *testing.T) {
	setup()
	defer teardown()

	id := "PFGEDX0"
	rawDataWanted := AnalyticsRawIncident{
		AssignmentCount:           10,
		BusinessHourInterruptions: 1,
		CreatedAt:                 "2021-01-06T00:00:00.000000",
		Description:               "Test Alert Description",
		EngagedSeconds:            404,
		EngagedUserCount:          2,
		EscalationCount:           1,
		ID:                        id,
		IncidentNumber:            324,
		IsMajor:                   true,
		OffHourInterruptions:      1,
		PriorityID:                "PSVSDF3",
		PriorityName:              "Urgent",
		ResolvedAt:                "2021-01-06T08:23:21.000000",
		SecondsToEngage:           3453,
		SecondsToFirstAck:         123,
		SecondsToMobilize:         435,
		SecondsToResolve:          324524,
		ServiceID:                 "PA4DFYS",
		ServiceName:               "My Test Service",
		SleepHourInterruptions:    0,
		SnoozedSeconds:            0,
		TeamID:                    "PCDYDX0",
		TeamName:                  "My Team",
		Urgency:                   "Urgent",
		UserDefinedEffortSeconds:  0,
	}

	bytesRawDataWanted, err := json.Marshal(rawDataWanted)
	testErrCheck(t, "json.Marshal()", "", err)

	url := fmt.Sprintf("%s/%s", "/analytics/raw/incidents", id)
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = w.Write(bytesRawDataWanted)
	})

	client := defaultTestClient(server.URL, "foo")

	res, err := client.GetAnalyticsIncidentsById(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, &rawDataWanted, res)
}

func TestAnalytics_GetAnalyticsIncidents(t *testing.T) {
	setup()
	defer teardown()

	id := "PFGEDX0"
	rawDataWanted := AnalyticsRawIncident{
		AssignmentCount:           10,
		BusinessHourInterruptions: 1,
		CreatedAt:                 "2021-01-06T00:00:00.000000",
		Description:               "Test Alert Description",
		EngagedSeconds:            404,
		EngagedUserCount:          2,
		EscalationCount:           1,
		ID:                        id,
		IncidentNumber:            324,
		IsMajor:                   true,
		OffHourInterruptions:      1,
		PriorityID:                "PSVSDF3",
		PriorityName:              "Urgent",
		ResolvedAt:                "2021-01-06T08:23:21.000000",
		SecondsToEngage:           3453,
		SecondsToFirstAck:         123,
		SecondsToMobilize:         435,
		SecondsToResolve:          324524,
		ServiceID:                 "PA4DFYS",
		ServiceName:               "My Test Service",
		SleepHourInterruptions:    0,
		SnoozedSeconds:            0,
		TeamID:                    "PCDYDX0",
		TeamName:                  "My Team",
		Urgency:                   "Urgent",
		UserDefinedEffortSeconds:  0,
	}

	rawDataRequest := AnalyticsRawIncidentsRequest{
		Filters: &AnalyticsFilter{
			CreatedAtStart: "2021-01-01T15:00:32Z",
			CreatedAtEnd:   "2021-01-08T15:00:32Z",
			TeamIDs:        []string{"PCDYDX0"},
		},
		StartingAfter: "eyJpZCI6IlEwTUlFTUtXTVNYOFZFIiwib3JkZXJfYnkiOiJjcmVhdGVkX2F0IiwidmFsdWUiOiIyMDIzLTA0LTMwVDA4OjU0OjAzIn0=",
		EndingBefore:  "eyJpZCI6IlEzQU1DSlhWU05HTjZEIiwib3JkZXJfYnkiOiJjcmVhdGVkX2F0IiwidmFsdWUiOiIyMDIzLTA0LTI2VDE4OjA3OjIyIn0=",
		Limit:         100,
		Order:         "asc",
		OrderBy:       "created_at",
		TimeZone:      "Etc/UTC",
	}
	rawFilterWanted := AnalyticsFilter{CreatedAtStart: "2021-01-06T09:21:41Z", CreatedAtEnd: "2021-01-13T09:21:41Z", TeamIDs: []string{"PCDYDX0"}}
	rawDataResponse := AnalyticsRawIncidentsResponse{
		Data:     []AnalyticsRawIncident{rawDataWanted},
		Filters:  &rawFilterWanted,
		First:    "eyJpZCI6IlEwTUlFTUtXTVNYOFZFIiwib3JkZXJfYnkiOiJjcmVhdGVkX2F0IiwidmFsdWUiOiIyMDIzLTA0LTMwVDA4OjU0OjAzIn0=",
		Last:     "eyJpZCI6IlEzQU1DSlhWU05HTjZEIiwib3JkZXJfYnkiOiJjcmVhdGVkX2F0IiwidmFsdWUiOiIyMDIzLTA0LTI2VDE4OjA3OjIyIn0=",
		Limit:    100,
		More:     true,
		Order:    "asc",
		OrderBy:  "created_at",
		TimeZone: "Etc/UTC",
	}

	bytesRawDataWanted, err := json.Marshal(rawDataResponse)
	testErrCheck(t, "json.Marshal()", "", err)

	mux.HandleFunc("/analytics/raw/incidents", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, _ = w.Write(bytesRawDataWanted)
	})

	client := defaultTestClient(server.URL, "foo")

	res, err := client.GetAnalyticsIncidents(context.Background(), rawDataRequest)
	want := &AnalyticsRawIncidentsResponse{
		Data:     []AnalyticsRawIncident{rawDataWanted},
		Filters:  &rawFilterWanted,
		First:    "eyJpZCI6IlEwTUlFTUtXTVNYOFZFIiwib3JkZXJfYnkiOiJjcmVhdGVkX2F0IiwidmFsdWUiOiIyMDIzLTA0LTMwVDA4OjU0OjAzIn0=",
		Last:     "eyJpZCI6IlEzQU1DSlhWU05HTjZEIiwib3JkZXJfYnkiOiJjcmVhdGVkX2F0IiwidmFsdWUiOiIyMDIzLTA0LTI2VDE4OjA3OjIyIn0=",
		Limit:    100,
		More:     true,
		Order:    "asc",
		OrderBy:  "created_at",
		TimeZone: "Etc/UTC",
	}
	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, want, res)
}
