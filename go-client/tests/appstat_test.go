package tests

import (
	"clientAcronis/client"
	"testing"
)

func TestAppStatMetrics(t *testing.T) {
	client.DeleteAll()
	client.PostMetrics("../client/metrics files/metric.txt")
	r := client.GetAppStatMetrics("Speed")
	if r != "{\"Sample_mean\":5}" {
		t.Errorf("Response was incorrect, got: " + r + " , want: {\"Sample_mean\":5}.")
	}
}
