package shield

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MetricsWafRequestBuilder builds and executes requests for operations under \shield\metrics\waf
type MetricsWafRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Blocked the blocked property
// returns a *MetricsWafBlockedRequestBuilder when successful
func (m *MetricsWafRequestBuilder) Blocked()(*MetricsWafBlockedRequestBuilder) {
    return NewMetricsWafBlockedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMetricsWafRequestBuilderInternal instantiates a new MetricsWafRequestBuilder and sets the default values.
func NewMetricsWafRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetricsWafRequestBuilder) {
    m := &MetricsWafRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shield/metrics/waf", pathParameters),
    }
    return m
}
// NewMetricsWafRequestBuilder instantiates a new MetricsWafRequestBuilder and sets the default values.
func NewMetricsWafRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetricsWafRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMetricsWafRequestBuilderInternal(urlParams, requestAdapter)
}
// Logged the logged property
// returns a *MetricsWafLoggedRequestBuilder when successful
func (m *MetricsWafRequestBuilder) Logged()(*MetricsWafLoggedRequestBuilder) {
    return NewMetricsWafLoggedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Processed the processed property
// returns a *MetricsWafProcessedRequestBuilder when successful
func (m *MetricsWafRequestBuilder) Processed()(*MetricsWafProcessedRequestBuilder) {
    return NewMetricsWafProcessedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Triggered the triggered property
// returns a *MetricsWafTriggeredRequestBuilder when successful
func (m *MetricsWafRequestBuilder) Triggered()(*MetricsWafTriggeredRequestBuilder) {
    return NewMetricsWafTriggeredRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
