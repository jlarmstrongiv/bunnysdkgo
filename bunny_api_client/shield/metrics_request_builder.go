package shield

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MetricsRequestBuilder builds and executes requests for operations under \shield\metrics
type MetricsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewMetricsRequestBuilderInternal instantiates a new MetricsRequestBuilder and sets the default values.
func NewMetricsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetricsRequestBuilder) {
    m := &MetricsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shield/metrics", pathParameters),
    }
    return m
}
// NewMetricsRequestBuilder instantiates a new MetricsRequestBuilder and sets the default values.
func NewMetricsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetricsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMetricsRequestBuilderInternal(urlParams, requestAdapter)
}
// Waf the waf property
// returns a *MetricsWafRequestBuilder when successful
func (m *MetricsRequestBuilder) Waf()(*MetricsWafRequestBuilder) {
    return NewMetricsWafRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
