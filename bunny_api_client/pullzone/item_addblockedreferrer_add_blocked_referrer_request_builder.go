package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAddblockedreferrerAddBlockedReferrerRequestBuilder builds and executes requests for operations under \pullzone\{-id}\addBlockedReferrer
type ItemAddblockedreferrerAddBlockedReferrerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemAddblockedreferrerAddBlockedReferrerRequestBuilderInternal instantiates a new ItemAddblockedreferrerAddBlockedReferrerRequestBuilder and sets the default values.
func NewItemAddblockedreferrerAddBlockedReferrerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddblockedreferrerAddBlockedReferrerRequestBuilder) {
    m := &ItemAddblockedreferrerAddBlockedReferrerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/addBlockedReferrer", pathParameters),
    }
    return m
}
// NewItemAddblockedreferrerAddBlockedReferrerRequestBuilder instantiates a new ItemAddblockedreferrerAddBlockedReferrerRequestBuilder and sets the default values.
func NewItemAddblockedreferrerAddBlockedReferrerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddblockedreferrerAddBlockedReferrerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAddblockedreferrerAddBlockedReferrerRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [AddBlockedReferer API Docs](https://docs.bunny.net/reference/pullzonepublic_addblockedreferrer)
func (m *ItemAddblockedreferrerAddBlockedReferrerRequestBuilder) Post(ctx context.Context, body ItemAddblockedreferrerAddBlockedReferrerPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPostRequestInformation [AddBlockedReferer API Docs](https://docs.bunny.net/reference/pullzonepublic_addblockedreferrer)
// returns a *RequestInformation when successful
func (m *ItemAddblockedreferrerAddBlockedReferrerRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemAddblockedreferrerAddBlockedReferrerPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAddblockedreferrerAddBlockedReferrerRequestBuilder when successful
func (m *ItemAddblockedreferrerAddBlockedReferrerRequestBuilder) WithUrl(rawUrl string)(*ItemAddblockedreferrerAddBlockedReferrerRequestBuilder) {
    return NewItemAddblockedreferrerAddBlockedReferrerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
