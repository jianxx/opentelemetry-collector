// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package obsreporttest_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configtelemetry"
	"go.opentelemetry.io/collector/obsreport"
	"go.opentelemetry.io/collector/obsreport/obsreporttest"
)

const (
	transport = "fakeTransport"
	format    = "fakeFormat"
)

var (
	receiver = config.NewComponentID("fakeReicever")
	exporter = config.NewComponentID("fakeExporter")
)

func TestCheckReceiverTracesViews(t *testing.T) {
	set, err := obsreporttest.SetupRecordedMetricsTest()
	require.NoError(t, err)
	defer set.Shutdown(context.Background())

	rec := obsreport.NewReceiver(obsreport.ReceiverSettings{ReceiverID: receiver, Transport: transport})
	ctx := rec.StartTracesOp(context.Background())
	assert.NotNil(t, ctx)
	rec.EndTracesOp(
		ctx,
		format,
		7,
		nil)

	require.NoError(t, obsreporttest.CheckReceiverTraces(receiver, transport, 7, 0))
}

func TestCheckReceiverMetricsViews(t *testing.T) {
	set, err := obsreporttest.SetupRecordedMetricsTest()
	require.NoError(t, err)
	defer set.Shutdown(context.Background())

	rec := obsreport.NewReceiver(obsreport.ReceiverSettings{ReceiverID: receiver, Transport: transport})
	ctx := rec.StartMetricsOp(context.Background())
	assert.NotNil(t, ctx)
	rec.EndMetricsOp(ctx, format, 7, nil)

	require.NoError(t, obsreporttest.CheckReceiverMetrics(receiver, transport, 7, 0))
}

func TestCheckReceiverLogsViews(t *testing.T) {
	set, err := obsreporttest.SetupRecordedMetricsTest()
	require.NoError(t, err)
	defer set.Shutdown(context.Background())

	rec := obsreport.NewReceiver(obsreport.ReceiverSettings{ReceiverID: receiver, Transport: transport})
	ctx := rec.StartLogsOp(context.Background())
	assert.NotNil(t, ctx)
	rec.EndLogsOp(ctx, format, 7, nil)

	require.NoError(t, obsreporttest.CheckReceiverLogs(receiver, transport, 7, 0))
}

func TestCheckExporterTracesViews(t *testing.T) {
	set, err := obsreporttest.SetupRecordedMetricsTest()
	require.NoError(t, err)
	defer set.Shutdown(context.Background())

	obsrep := obsreport.NewExporter(obsreport.ExporterSettings{
		Level:                  configtelemetry.LevelNormal,
		ExporterID:             exporter,
		ExporterCreateSettings: set.ToExporterCreateSettings(),
	})
	ctx := obsrep.StartTracesOp(context.Background())
	assert.NotNil(t, ctx)

	obsrep.EndTracesOp(ctx, 7, nil)

	require.NoError(t, obsreporttest.CheckExporterTraces(exporter, 7, 0))
}

func TestCheckExporterMetricsViews(t *testing.T) {
	set, err := obsreporttest.SetupRecordedMetricsTest()
	require.NoError(t, err)
	defer set.Shutdown(context.Background())

	obsrep := obsreport.NewExporter(obsreport.ExporterSettings{
		Level:                  configtelemetry.LevelNormal,
		ExporterID:             exporter,
		ExporterCreateSettings: set.ToExporterCreateSettings(),
	})
	ctx := obsrep.StartMetricsOp(context.Background())
	assert.NotNil(t, ctx)

	obsrep.EndMetricsOp(ctx, 7, nil)

	require.NoError(t, obsreporttest.CheckExporterMetrics(exporter, 7, 0))
}

func TestCheckExporterLogsViews(t *testing.T) {
	set, err := obsreporttest.SetupRecordedMetricsTest()
	require.NoError(t, err)
	defer set.Shutdown(context.Background())

	obsrep := obsreport.NewExporter(obsreport.ExporterSettings{
		Level:                  configtelemetry.LevelNormal,
		ExporterID:             exporter,
		ExporterCreateSettings: set.ToExporterCreateSettings(),
	})
	ctx := obsrep.StartLogsOp(context.Background())
	assert.NotNil(t, ctx)
	obsrep.EndLogsOp(ctx, 7, nil)

	require.NoError(t, obsreporttest.CheckExporterLogs(exporter, 7, 0))
}
