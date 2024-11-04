package shield

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MetricsWafLoggedRequestBuilder builds and executes requests for operations under \shield\metrics\waf\logged
type MetricsWafLoggedRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByShieldZoneId gets an item from the github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client.shield.metrics.waf.logged.item collection
// returns a *MetricsWafLoggedWithShieldZoneItemRequestBuilder when successful
func (m *MetricsWafLoggedRequestBuilder) ByShieldZoneId(shieldZoneId int32)(*MetricsWafLoggedWithShieldZoneItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["shieldZoneId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(shieldZoneId), 10)
    return NewMetricsWafLoggedWithShieldZoneItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMetricsWafLoggedRequestBuilderInternal instantiates a new MetricsWafLoggedRequestBuilder and sets the default values.
func NewMetricsWafLoggedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetricsWafLoggedRequestBuilder) {
    m := &MetricsWafLoggedRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shield/metrics/waf/logged", pathParameters),
    }
    return m
}
// NewMetricsWafLoggedRequestBuilder instantiates a new MetricsWafLoggedRequestBuilder and sets the default values.
func NewMetricsWafLoggedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetricsWafLoggedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMetricsWafLoggedRequestBuilderInternal(urlParams, requestAdapter)
}
