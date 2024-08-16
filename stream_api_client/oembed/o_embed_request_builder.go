package oembed

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OEmbedRequestBuilder builds and executes requests for operations under \OEmbed
type OEmbedRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OEmbedRequestBuilderGetQueryParameters [OEmbed API Docs](https://docs.bunny.net/reference/oembed_getoembed)
type OEmbedRequestBuilderGetQueryParameters struct {
    Expires *int64 `uriparametername:"expires"`
    MaxHeight *int32 `uriparametername:"maxHeight"`
    MaxWidth *int32 `uriparametername:"maxWidth"`
    Token *string `uriparametername:"token"`
    Url *string `uriparametername:"url"`
}
// NewOEmbedRequestBuilderInternal instantiates a new OEmbedRequestBuilder and sets the default values.
func NewOEmbedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OEmbedRequestBuilder) {
    m := &OEmbedRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/OEmbed?expires={expires}{&maxHeight,maxWidth,token,url}", pathParameters),
    }
    return m
}
// NewOEmbedRequestBuilder instantiates a new OEmbedRequestBuilder and sets the default values.
func NewOEmbedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OEmbedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOEmbedRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [OEmbed API Docs](https://docs.bunny.net/reference/oembed_getoembed)
// returns a OEmbedGetResponseable when successful
func (m *OEmbedRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[OEmbedRequestBuilderGetQueryParameters])(OEmbedGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOEmbedGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OEmbedGetResponseable), nil
}
// ToGetRequestInformation [OEmbed API Docs](https://docs.bunny.net/reference/oembed_getoembed)
// returns a *RequestInformation when successful
func (m *OEmbedRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[OEmbedRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OEmbedRequestBuilder when successful
func (m *OEmbedRequestBuilder) WithUrl(rawUrl string)(*OEmbedRequestBuilder) {
    return NewOEmbedRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
