package metricsservice

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/proto/spire/common/hostservices"
	mock_hostservices "github.com/spiffe/spire/test/mock/proto/common/hostservices"
)

func setupPluginMetricsWrapper(ctx context.Context, m hostservices.MetricsService, log hclog.Logger) telemetry.Metrics {
	return WrapPluginMetricsForContext(ctx, m, log)
}

func TestWrapEmitKey(t *testing.T) {
	tests := []struct {
		desc  string
		inKey []string
		inVal float32
	}{
		{
			desc:  "base call",
			inKey: []string{"key1", "key2"},
			inVal: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			ctx := context.Background()

			mockMetricsService := mock_hostservices.NewMockMetricsService(mockCtrl)

			mockMetricsService.EXPECT().EmitKey(ctx, &hostservices.EmitKeyRequest{
				Key: tt.inKey,
				Val: tt.inVal,
			}).Return(&hostservices.EmitKeyResponse{}, nil)

			metricsWrapper := setupPluginMetricsWrapper(ctx, mockMetricsService, hclog.NewNullLogger())
			metricsWrapper.EmitKey(tt.inKey, tt.inVal)
		})
	}
}

func TestWrapSetGaugeWithLabels(t *testing.T) {
	tests := []struct {
		desc     string
		inKey    []string
		inVal    float32
		inLabels []telemetry.Label
	}{
		{
			desc:  "no labels",
			inKey: []string{"key1", "key2"},
			inVal: 3,
		},
		{
			desc:  "label",
			inKey: []string{"key1", "key2"},
			inVal: 3,
			inLabels: []telemetry.Label{
				{
					Name:  "label1",
					Value: "val1",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			ctx := context.Background()

			mockMetricsService := mock_hostservices.NewMockMetricsService(mockCtrl)

			mockMetricsService.EXPECT().SetGauge(ctx, &hostservices.SetGaugeRequest{
				Key:    tt.inKey,
				Val:    tt.inVal,
				Labels: convertToRPCLabels(tt.inLabels),
			}).Return(&hostservices.SetGaugeResponse{}, nil)

			metricsWrapper := setupPluginMetricsWrapper(ctx, mockMetricsService, hclog.NewNullLogger())
			metricsWrapper.SetGaugeWithLabels(tt.inKey, tt.inVal, tt.inLabels)
		})
	}
}

func TestWrapIncrCounterWithLabels(t *testing.T) {
	tests := []struct {
		desc     string
		inKey    []string
		inVal    float32
		inLabels []telemetry.Label
	}{
		{
			desc:  "no labels",
			inKey: []string{"key1", "key2"},
			inVal: 3,
		},
		{
			desc:  "label",
			inKey: []string{"key1", "key2"},
			inVal: 3,
			inLabels: []telemetry.Label{
				{
					Name:  "label1",
					Value: "val1",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			ctx := context.Background()

			mockMetricsService := mock_hostservices.NewMockMetricsService(mockCtrl)

			mockMetricsService.EXPECT().IncrCounter(ctx, &hostservices.IncrCounterRequest{
				Key:    tt.inKey,
				Val:    tt.inVal,
				Labels: convertToRPCLabels(tt.inLabels),
			}).Return(&hostservices.IncrCounterResponse{}, nil)

			metricsWrapper := setupPluginMetricsWrapper(ctx, mockMetricsService, hclog.NewNullLogger())
			metricsWrapper.IncrCounterWithLabels(tt.inKey, tt.inVal, tt.inLabels)
		})
	}
}

func TestWrapAddSampleWithLabels(t *testing.T) {
	tests := []struct {
		desc     string
		inKey    []string
		inVal    float32
		inLabels []telemetry.Label
	}{
		{
			desc:  "no labels",
			inKey: []string{"key1", "key2"},
			inVal: 3,
		},
		{
			desc:  "label",
			inKey: []string{"key1", "key2"},
			inVal: 3,
			inLabels: []telemetry.Label{
				{
					Name:  "label1",
					Value: "val1",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			ctx := context.Background()

			mockMetricsService := mock_hostservices.NewMockMetricsService(mockCtrl)

			mockMetricsService.EXPECT().AddSample(ctx, &hostservices.AddSampleRequest{
				Key:    tt.inKey,
				Val:    tt.inVal,
				Labels: convertToRPCLabels(tt.inLabels),
			}).Return(&hostservices.AddSampleResponse{}, nil)

			metricsWrapper := setupPluginMetricsWrapper(ctx, mockMetricsService, hclog.NewNullLogger())
			metricsWrapper.AddSampleWithLabels(tt.inKey, tt.inVal, tt.inLabels)
		})
	}
}

func TestWrapMeasureSinceWithLabels(t *testing.T) {
	tests := []struct {
		desc     string
		inKey    []string
		inTime   time.Time
		inLabels []telemetry.Label
	}{
		{
			desc:   "no labels",
			inKey:  []string{"key1", "key2"},
			inTime: time.Now(),
		},
		{
			desc:   "label",
			inKey:  []string{"key1", "key2"},
			inTime: time.Now(),
			inLabels: []telemetry.Label{
				{
					Name:  "label1",
					Value: "val1",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			ctx := context.Background()

			mockMetricsService := mock_hostservices.NewMockMetricsService(mockCtrl)

			mockMetricsService.EXPECT().MeasureSince(ctx, &hostservices.MeasureSinceRequest{
				Key:    tt.inKey,
				Time:   tt.inTime.UnixNano(),
				Labels: convertToRPCLabels(tt.inLabels),
			}).Return(&hostservices.MeasureSinceResponse{}, nil)

			metricsWrapper := setupPluginMetricsWrapper(ctx, mockMetricsService, hclog.NewNullLogger())
			metricsWrapper.MeasureSinceWithLabels(tt.inKey, tt.inTime, tt.inLabels)
		})
	}
}
