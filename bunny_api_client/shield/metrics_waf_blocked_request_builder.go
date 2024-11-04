package shield

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MetricsWafBlockedRequestBuilder builds and executes requests for operations under \shield\metrics\waf\blocked
type MetricsWafBlockedRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByShieldZoneId gets an item from the github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client.shield.metrics.waf.blocked.item collection
// returns a *MetricsWafBlockedWithShieldZoneItemRequestBuilder when successful
func (m *MetricsWafBlockedRequestBuilder) ByShieldZoneId(shieldZoneId int32)(*MetricsWafBlockedWithShieldZoneItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["shieldZoneId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(shieldZoneId), 10)
    return NewMetricsWafBlockedWithShieldZoneItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMetricsWafBlockedRequestBuilderInternal instantiates a new MetricsWafBlockedRequestBuilder and sets the default values.
func NewMetricsWafBlockedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetricsWafBlockedRequestBuilder) {
    m := &MetricsWafBlockedRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shield/metrics/waf/blocked", pathParameters),
    }
    return m
}
// NewMetricsWafBlockedRequestBuilder instantiates a new MetricsWafBlockedRequestBuilder and sets the default values.
func NewMetricsWafBlockedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetricsWafBlockedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMetricsWafBlockedRequestBuilderInternal(urlParams, requestAdapter)
}
