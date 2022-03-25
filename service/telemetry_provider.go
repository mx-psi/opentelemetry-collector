// Copyright  The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service // import "go.opentelemetry.io/collector/service"

import (
	"fmt"

	"go.opentelemetry.io/contrib/zpages"
	"go.opentelemetry.io/otel/metric/nonrecording"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/service/internal"
	"go.opentelemetry.io/collector/service/internal/telemetrylogs"
)

// TelemetryProvider is the provider of telemetry.
type TelemetryProvider interface {
	// SetupTelemetry creates component.TelemetrySettings and possibly handles global telemetry setup.
	SetupTelemetry(cfg config.ServiceTelemetry) (component.TelemetrySettings, error)
	// ZPages returns the zpages.SpanProcessor associated with telemetry.
	ZPages() *zpages.SpanProcessor
}

var _ TelemetryProvider = (*defaultTelemetryFactory)(nil)

type defaultTelemetryFactory struct {
	zPagesSpanProcessor *zpages.SpanProcessor
	loggingOptions      []zap.Option
}

// NewDefaultTelemetryProvider creates the default telemetry provider.
func NewDefaultTelemetryProvider(loggingOptions ...zap.Option) TelemetryProvider {
	return &defaultTelemetryFactory{
		zPagesSpanProcessor: zpages.NewSpanProcessor(),
		loggingOptions:      loggingOptions,
	}
}

// SetupTelemetry implements the TelemetryProvider interface.
func (t *defaultTelemetryFactory) SetupTelemetry(cfg config.ServiceTelemetry) (set component.TelemetrySettings, err error) {
	set.Logger, err = telemetrylogs.NewLogger(cfg.Logs, t.loggingOptions)
	if err != nil {
		return component.TelemetrySettings{}, fmt.Errorf("failed to get logger: %w", err)
	}

	telemetrylogs.SetColGRPCLogger(set.Logger, cfg.Logs.Level)

	set.TracerProvider = sdktrace.NewTracerProvider(
		sdktrace.WithSampler(internal.AlwaysRecord()),
		sdktrace.WithSpanProcessor(t.zPagesSpanProcessor))

	set.MeterProvider = nonrecording.NewNoopMeterProvider()
	set.MetricsLevel = cfg.Metrics.Level

	return set, nil
}

// ZPages implements the TelemetryProvider interface.
func (t *defaultTelemetryFactory) ZPages() *zpages.SpanProcessor {
	return t.zPagesSpanProcessor
}
