package shield

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MetricsWafProcessedRequestBuilder builds and executes requests for operations under \shield\metrics\waf\processed
type MetricsWafProcessedRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByShieldZoneId gets an item from the github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client.shield.metrics.waf.processed.item collection
// returns a *MetricsWafProcessedWithShieldZoneItemRequestBuilder when successful
func (m *MetricsWafProcessedRequestBuilder) ByShieldZoneId(shieldZoneId int32)(*MetricsWafProcessedWithShieldZoneItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["shieldZoneId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(shieldZoneId), 10)
    return NewMetricsWafProcessedWithShieldZoneItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMetricsWafProcessedRequestBuilderInternal instantiates a new MetricsWafProcessedRequestBuilder and sets the default values.
func NewMetricsWafProcessedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetricsWafProcessedRequestBuilder) {
    m := &MetricsWafProcessedRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shield/metrics/waf/processed", pathParameters),
    }
    return m
}
// NewMetricsWafProcessedRequestBuilder instantiates a new MetricsWafProcessedRequestBuilder and sets the default values.
func NewMetricsWafProcessedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetricsWafProcessedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMetricsWafProcessedRequestBuilderInternal(urlParams, requestAdapter)
}
