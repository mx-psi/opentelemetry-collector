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

package service // import "go.opentelemetry.io/collector/service"

import (
	"go.opentelemetry.io/contrib/zpages"
	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configmapprovider"
	"go.opentelemetry.io/collector/config/configunmarshaler"
)

// svcSettings holds configuration for building a new service.
type svcSettings struct {
	Telemetry           component.TelemetrySettings
	Factories           component.Factories
	Config              *config.Config
	ZPagesSpanProcessor *zpages.SpanProcessor
	AsyncErrorChannel   chan error
	BuildInfo           component.BuildInfo
}

// CollectorSettings holds configuration for creating a new Collector.
type CollectorSettings struct {
	Factories               component.Factories
	ConfigMapProvider       configmapprovider.Provider
	ConfigUnmarshaler       configunmarshaler.ConfigUnmarshaler
	BuildInfo               component.BuildInfo
	LoggingOptions          []zap.Option
	DisableGracefulShutdown bool
}
