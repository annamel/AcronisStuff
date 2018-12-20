package tests

import (
	"clientAcronis/client"
	"regexp"
	"testing"
)

func TestDeleteAll(t *testing.T) {
	client.DeleteAll()
	r := client.GetAll()
	if r != "{\"Logs\":\"\",\"Metrics\":\"\"}" {
		t.Errorf("Response was incorrect, got: " + r + " , want: {\"Logs\":\"\",\"Metrics\":\"\"}.")
	}
}

func TestPostLogs(t *testing.T) {
	before := client.GetAll()
	LogsPath := "../client/swagger.log"
	client.PostLogs(LogsPath)
	after := client.GetAll()
	if len(before) == len(after) || len(before) > len(after) {
		t.Errorf("Response was incorrect, data wasn't added")
	}
}

func TestPostMetrics(t *testing.T) {
	before := client.GetAll()
	MetricsPath := "../client/metric.txt"
	client.PostMetrics(MetricsPath)
	after := client.GetAll()
	if len(before) == len(after) || len(before) > len(after) {
		t.Errorf("Response was incorrect, data wasn't added")
	}
}

func TestDeleteMetricsById(t *testing.T) {
	before := client.GetAll()
	if before == "{\"Logs\":\"\",\"Metrics\":\"\"}" {
		t.Errorf("Response was incorrect, database is already empty.")
	}
	r, _ := regexp.Compile("Metrics\":\"[0-9a-z]+")
	id := r.FindString(before)
	id = id[10:]
	client.DeleteMetricsById(id)
	after := client.GetAll()
	if len(before) == len(after) || len(before) < len(after) {
		t.Errorf("Response was incorrect, data wasn't deleted")
	}
}

func TestDeleteLogsById(t *testing.T) {
	before := client.GetAll()
	if before == "{\"Logs\":\"\",\"Metrics\":\"\"}" {
		t.Errorf("Response was incorrect, database is already empty.")
	}
	r, _ := regexp.Compile("Logs\":\"[0-9a-z]+")
	id := r.FindString(before)
	id = id[7:]
	client.DeleteLogsById(id)
	after := client.GetAll()
	if len(before) == len(after) || len(before) < len(after) {
		t.Errorf("Response was incorrect, data wasn't deleted")
	}
}

func TestDeleteLogs(t *testing.T) {
	before := client.GetAll()
	if before == "{\"Logs\":\"\",\"Metrics\":\"\"}" {
		t.Errorf("Response was incorrect, database is already empty.")
	}
	client.DeleteAllLogs()
	after := client.GetAll()
	if len(before) == len(after) || len(before) < len(after) {
		t.Errorf("Response was incorrect, data wasn't deleted")
	}
}

func TestDeleteMetrics(t *testing.T) {
	before := client.GetAll()
	if before == "{\"Logs\":\"\",\"Metrics\":\"\"}" {
		t.Errorf("Response was incorrect, database is already empty.")
	}
	client.DeleteAllMetrics()
	after := client.GetAll()
	if len(before) == len(after) || len(before) < len(after) {
		t.Errorf("Response was incorrect, data wasn't deleted")
	}
}
