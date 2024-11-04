package shield

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MetricsWafTriggeredRequestBuilder builds and executes requests for operations under \shield\metrics\waf\triggered
type MetricsWafTriggeredRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByShieldZoneId gets an item from the github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client.shield.metrics.waf.triggered.item collection
// returns a *MetricsWafTriggeredWithShieldZoneItemRequestBuilder when successful
func (m *MetricsWafTriggeredRequestBuilder) ByShieldZoneId(shieldZoneId int32)(*MetricsWafTriggeredWithShieldZoneItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["shieldZoneId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(shieldZoneId), 10)
    return NewMetricsWafTriggeredWithShieldZoneItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMetricsWafTriggeredRequestBuilderInternal instantiates a new MetricsWafTriggeredRequestBuilder and sets the default values.
func NewMetricsWafTriggeredRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetricsWafTriggeredRequestBuilder) {
    m := &MetricsWafTriggeredRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shield/metrics/waf/triggered", pathParameters),
    }
    return m
}
// NewMetricsWafTriggeredRequestBuilder instantiates a new MetricsWafTriggeredRequestBuilder and sets the default values.
func NewMetricsWafTriggeredRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetricsWafTriggeredRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMetricsWafTriggeredRequestBuilderInternal(urlParams, requestAdapter)
}
