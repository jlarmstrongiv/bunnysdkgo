package shield

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/shield"
)

// ShieldZoneGetByPullzoneWithPullZoneItemRequestBuilder builds and executes requests for operations under \shield\shield-zone\get-by-pullzone\{pullZoneId}
type ShieldZoneGetByPullzoneWithPullZoneItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewShieldZoneGetByPullzoneWithPullZoneItemRequestBuilderInternal instantiates a new ShieldZoneGetByPullzoneWithPullZoneItemRequestBuilder and sets the default values.
func NewShieldZoneGetByPullzoneWithPullZoneItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ShieldZoneGetByPullzoneWithPullZoneItemRequestBuilder) {
    m := &ShieldZoneGetByPullzoneWithPullZoneItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shield/shield-zone/get-by-pullzone/{pullZoneId}", pathParameters),
    }
    return m
}
// NewShieldZoneGetByPullzoneWithPullZoneItemRequestBuilder instantiates a new ShieldZoneGetByPullzoneWithPullZoneItemRequestBuilder and sets the default values.
func NewShieldZoneGetByPullzoneWithPullZoneItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ShieldZoneGetByPullzoneWithPullZoneItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewShieldZoneGetByPullzoneWithPullZoneItemRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a GetShieldZoneResponseable when successful
// returns a ProblemDetails error when the service returns a 401 status code
func (m *ShieldZoneGetByPullzoneWithPullZoneItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.GetShieldZoneResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.CreateProblemDetailsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.CreateGetShieldZoneResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.GetShieldZoneResponseable), nil
}
// returns a *RequestInformation when successful
func (m *ShieldZoneGetByPullzoneWithPullZoneItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json, text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ShieldZoneGetByPullzoneWithPullZoneItemRequestBuilder when successful
func (m *ShieldZoneGetByPullzoneWithPullZoneItemRequestBuilder) WithUrl(rawUrl string)(*ShieldZoneGetByPullzoneWithPullZoneItemRequestBuilder) {
    return NewShieldZoneGetByPullzoneWithPullZoneItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
