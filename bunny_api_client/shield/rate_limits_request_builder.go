package shield

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RateLimitsRequestBuilder builds and executes requests for operations under \shield\rate-limits
type RateLimitsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByShieldZoneId gets an item from the github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client.shield.rateLimits.item collection
// returns a *RateLimitsWithShieldZoneItemRequestBuilder when successful
func (m *RateLimitsRequestBuilder) ByShieldZoneId(shieldZoneId int64)(*RateLimitsWithShieldZoneItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["shieldZoneId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(shieldZoneId, 10)
    return NewRateLimitsWithShieldZoneItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewRateLimitsRequestBuilderInternal instantiates a new RateLimitsRequestBuilder and sets the default values.
func NewRateLimitsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RateLimitsRequestBuilder) {
    m := &RateLimitsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shield/rate-limits", pathParameters),
    }
    return m
}
// NewRateLimitsRequestBuilder instantiates a new RateLimitsRequestBuilder and sets the default values.
func NewRateLimitsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RateLimitsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRateLimitsRequestBuilderInternal(urlParams, requestAdapter)
}
