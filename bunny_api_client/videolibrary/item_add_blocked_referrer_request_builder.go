package videolibrary

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAddBlockedReferrerRequestBuilder builds and executes requests for operations under \videolibrary\{id}\addBlockedReferrer
type ItemAddBlockedReferrerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemAddBlockedReferrerRequestBuilderInternal instantiates a new ItemAddBlockedReferrerRequestBuilder and sets the default values.
func NewItemAddBlockedReferrerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddBlockedReferrerRequestBuilder) {
    m := &ItemAddBlockedReferrerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/videolibrary/{id}/addBlockedReferrer", pathParameters),
    }
    return m
}
// NewItemAddBlockedReferrerRequestBuilder instantiates a new ItemAddBlockedReferrerRequestBuilder and sets the default values.
func NewItemAddBlockedReferrerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddBlockedReferrerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAddBlockedReferrerRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [AddBlockedReferer API Docs](https://docs.bunny.net/reference/videolibrarypublic_addblockedreferrer)
func (m *ItemAddBlockedReferrerRequestBuilder) Post(ctx context.Context, body ItemAddBlockedReferrerPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPostRequestInformation [AddBlockedReferer API Docs](https://docs.bunny.net/reference/videolibrarypublic_addblockedreferrer)
// returns a *RequestInformation when successful
func (m *ItemAddBlockedReferrerRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemAddBlockedReferrerPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAddBlockedReferrerRequestBuilder when successful
func (m *ItemAddBlockedReferrerRequestBuilder) WithUrl(rawUrl string)(*ItemAddBlockedReferrerRequestBuilder) {
    return NewItemAddBlockedReferrerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
