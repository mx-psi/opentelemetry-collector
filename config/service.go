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

package config // import "go.opentelemetry.io/collector/config"

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

// Service defines the configurable components of the service.
type Service struct {
	Pipelines  Pipelines        `mapstructure:"pipelines"`
	Telemetry  ServiceTelemetry `mapstructure:"telemetry"`
	Extensions []ComponentID    `mapstructure:"extensions"`
}

// ServiceTelemetry defines the configurable settings for service telemetry.
type ServiceTelemetry struct {
	Logs ServiceTelemetryLogs `mapstructure:"logs"`
}

func (srvT *ServiceTelemetry) validate() error {
	return srvT.Logs.validate()
}

// ServiceTelemetryLogs defines the configurable settings for service telemetry logs.
// This MUST be compatible with zap.Config. Cannot use directly zap.Config because
// the collector uses mapstructure and not yaml tags.
type ServiceTelemetryLogs struct {
	Encoding    string        `mapstructure:"encoding"`
	Level       zapcore.Level `mapstructure:"level"`
	Development bool          `mapstructure:"development"`
}

func (srvTL *ServiceTelemetryLogs) validate() error {
	if srvTL.Encoding != "json" && srvTL.Encoding != "console" {
		return fmt.Errorf(`service telemetry logs invalid encoding: %q, valid values are "json" and "console"`, srvTL.Encoding)
	}
	return nil
}

// DataType is a special Type that represents the data types supported by the collector. We currently support
// collecting metrics, traces and logs, this can expand in the future.
type DataType = Type

// Currently supported data types. Add new data types here when new types are supported in the future.
const (
	// TracesDataType is the data type tag for traces.
	TracesDataType DataType = "traces"

	// MetricsDataType is the data type tag for metrics.
	MetricsDataType DataType = "metrics"

	// LogsDataType is the data type tag for logs.
	LogsDataType DataType = "logs"
)

// Pipeline defines a single pipeline.
type Pipeline struct {
	Receivers  []ComponentID `mapstructure:"receivers"`
	Processors []ComponentID `mapstructure:"processors"`
	Exporters  []ComponentID `mapstructure:"exporters"`
}

// Pipelines is a map of names to Pipelines.
type Pipelines map[ComponentID]*Pipeline
