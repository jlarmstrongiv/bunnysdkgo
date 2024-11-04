package shield

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ShieldRequestBuilder builds and executes requests for operations under \shield
type ShieldRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewShieldRequestBuilderInternal instantiates a new ShieldRequestBuilder and sets the default values.
func NewShieldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ShieldRequestBuilder) {
    m := &ShieldRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shield", pathParameters),
    }
    return m
}
// NewShieldRequestBuilder instantiates a new ShieldRequestBuilder and sets the default values.
func NewShieldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ShieldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewShieldRequestBuilderInternal(urlParams, requestAdapter)
}
// Metrics the metrics property
// returns a *MetricsRequestBuilder when successful
func (m *ShieldRequestBuilder) Metrics()(*MetricsRequestBuilder) {
    return NewMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RateLimit the rateLimit property
// returns a *RateLimitRequestBuilder when successful
func (m *ShieldRequestBuilder) RateLimit()(*RateLimitRequestBuilder) {
    return NewRateLimitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RateLimits the rateLimits property
// returns a *RateLimitsRequestBuilder when successful
func (m *ShieldRequestBuilder) RateLimits()(*RateLimitsRequestBuilder) {
    return NewRateLimitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShieldZone the shieldZone property
// returns a *ShieldZoneRequestBuilder when successful
func (m *ShieldRequestBuilder) ShieldZone()(*ShieldZoneRequestBuilder) {
    return NewShieldZoneRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShieldZones the shieldZones property
// returns a *ShieldZonesRequestBuilder when successful
func (m *ShieldRequestBuilder) ShieldZones()(*ShieldZonesRequestBuilder) {
    return NewShieldZonesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Waf the waf property
// returns a *WafRequestBuilder when successful
func (m *ShieldRequestBuilder) Waf()(*WafRequestBuilder) {
    return NewWafRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
