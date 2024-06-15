package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemPurgecachePurgeCacheRequestBuilder builds and executes requests for operations under \pullzone\{-id}\purgeCache
type ItemPurgecachePurgeCacheRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemPurgecachePurgeCacheRequestBuilderInternal instantiates a new ItemPurgecachePurgeCacheRequestBuilder and sets the default values.
func NewItemPurgecachePurgeCacheRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPurgecachePurgeCacheRequestBuilder) {
    m := &ItemPurgecachePurgeCacheRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/purgeCache", pathParameters),
    }
    return m
}
// NewItemPurgecachePurgeCacheRequestBuilder instantiates a new ItemPurgecachePurgeCacheRequestBuilder and sets the default values.
func NewItemPurgecachePurgeCacheRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPurgecachePurgeCacheRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPurgecachePurgeCacheRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [PurgeCache API Docs](https://docs.bunny.net/reference/pullzonepublic_purgecachepostbytag)
func (m *ItemPurgecachePurgeCacheRequestBuilder) Post(ctx context.Context, body ItemPurgecachePurgeCachePostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation [PurgeCache API Docs](https://docs.bunny.net/reference/pullzonepublic_purgecachepostbytag)
// returns a *RequestInformation when successful
func (m *ItemPurgecachePurgeCacheRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemPurgecachePurgeCachePostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPurgecachePurgeCacheRequestBuilder when successful
func (m *ItemPurgecachePurgeCacheRequestBuilder) WithUrl(rawUrl string)(*ItemPurgecachePurgeCacheRequestBuilder) {
    return NewItemPurgecachePurgeCacheRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
