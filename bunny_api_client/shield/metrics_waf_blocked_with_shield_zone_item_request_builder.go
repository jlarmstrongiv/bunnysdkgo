package shield

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/shield"
)

// MetricsWafBlockedWithShieldZoneItemRequestBuilder builds and executes requests for operations under \shield\metrics\waf\blocked\{shieldZoneId}
type MetricsWafBlockedWithShieldZoneItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
type MetricsWafBlockedWithShieldZoneItemRequestBuilderGetQueryParameters struct {
    DataTimeframe *int32 `uriparametername:"dataTimeframe"`
    GroupTime *string `uriparametername:"groupTime"`
}
// NewMetricsWafBlockedWithShieldZoneItemRequestBuilderInternal instantiates a new MetricsWafBlockedWithShieldZoneItemRequestBuilder and sets the default values.
func NewMetricsWafBlockedWithShieldZoneItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetricsWafBlockedWithShieldZoneItemRequestBuilder) {
    m := &MetricsWafBlockedWithShieldZoneItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shield/metrics/waf/blocked/{shieldZoneId}{?dataTimeframe,groupTime}", pathParameters),
    }
    return m
}
// NewMetricsWafBlockedWithShieldZoneItemRequestBuilder instantiates a new MetricsWafBlockedWithShieldZoneItemRequestBuilder and sets the default values.
func NewMetricsWafBlockedWithShieldZoneItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetricsWafBlockedWithShieldZoneItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMetricsWafBlockedWithShieldZoneItemRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a WafMetricsResponseable when successful
// returns a UnauthorizedResult error when the service returns a 401 status code
func (m *MetricsWafBlockedWithShieldZoneItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[MetricsWafBlockedWithShieldZoneItemRequestBuilderGetQueryParameters])(i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.WafMetricsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.CreateUnauthorizedResultFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.CreateWafMetricsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.WafMetricsResponseable), nil
}
// returns a *RequestInformation when successful
func (m *MetricsWafBlockedWithShieldZoneItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[MetricsWafBlockedWithShieldZoneItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json, text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MetricsWafBlockedWithShieldZoneItemRequestBuilder when successful
func (m *MetricsWafBlockedWithShieldZoneItemRequestBuilder) WithUrl(rawUrl string)(*MetricsWafBlockedWithShieldZoneItemRequestBuilder) {
    return NewMetricsWafBlockedWithShieldZoneItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
