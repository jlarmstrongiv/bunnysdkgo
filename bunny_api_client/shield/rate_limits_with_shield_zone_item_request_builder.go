package shield

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/shield"
)

// RateLimitsWithShieldZoneItemRequestBuilder builds and executes requests for operations under \shield\rate-limits\{shieldZoneId}
type RateLimitsWithShieldZoneItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
type RateLimitsWithShieldZoneItemRequestBuilderGetQueryParameters struct {
    Page *int32 `uriparametername:"page"`
    PerPage *int32 `uriparametername:"perPage"`
}
// NewRateLimitsWithShieldZoneItemRequestBuilderInternal instantiates a new RateLimitsWithShieldZoneItemRequestBuilder and sets the default values.
func NewRateLimitsWithShieldZoneItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RateLimitsWithShieldZoneItemRequestBuilder) {
    m := &RateLimitsWithShieldZoneItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shield/rate-limits/{shieldZoneId}{?page,perPage}", pathParameters),
    }
    return m
}
// NewRateLimitsWithShieldZoneItemRequestBuilder instantiates a new RateLimitsWithShieldZoneItemRequestBuilder and sets the default values.
func NewRateLimitsWithShieldZoneItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RateLimitsWithShieldZoneItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRateLimitsWithShieldZoneItemRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a GetCustomWafRulesResponseable when successful
// returns a UnauthorizedResult error when the service returns a 401 status code
func (m *RateLimitsWithShieldZoneItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[RateLimitsWithShieldZoneItemRequestBuilderGetQueryParameters])(i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.GetCustomWafRulesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.CreateUnauthorizedResultFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.CreateGetCustomWafRulesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.GetCustomWafRulesResponseable), nil
}
// returns a *RequestInformation when successful
func (m *RateLimitsWithShieldZoneItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[RateLimitsWithShieldZoneItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json, text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RateLimitsWithShieldZoneItemRequestBuilder when successful
func (m *RateLimitsWithShieldZoneItemRequestBuilder) WithUrl(rawUrl string)(*RateLimitsWithShieldZoneItemRequestBuilder) {
    return NewRateLimitsWithShieldZoneItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
