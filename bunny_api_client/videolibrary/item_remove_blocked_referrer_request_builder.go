package videolibrary

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemRemoveBlockedReferrerRequestBuilder builds and executes requests for operations under \videolibrary\{id}\removeBlockedReferrer
type ItemRemoveBlockedReferrerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemRemoveBlockedReferrerRequestBuilderInternal instantiates a new ItemRemoveBlockedReferrerRequestBuilder and sets the default values.
func NewItemRemoveBlockedReferrerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoveBlockedReferrerRequestBuilder) {
    m := &ItemRemoveBlockedReferrerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/videolibrary/{id}/removeBlockedReferrer", pathParameters),
    }
    return m
}
// NewItemRemoveBlockedReferrerRequestBuilder instantiates a new ItemRemoveBlockedReferrerRequestBuilder and sets the default values.
func NewItemRemoveBlockedReferrerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoveBlockedReferrerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRemoveBlockedReferrerRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [RemoveBlockedReferer API Docs](https://docs.bunny.net/reference/videolibrarypublic_removeblockedreferrer)
func (m *ItemRemoveBlockedReferrerRequestBuilder) Post(ctx context.Context, body ItemRemoveBlockedReferrerPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPostRequestInformation [RemoveBlockedReferer API Docs](https://docs.bunny.net/reference/videolibrarypublic_removeblockedreferrer)
// returns a *RequestInformation when successful
func (m *ItemRemoveBlockedReferrerRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemRemoveBlockedReferrerPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRemoveBlockedReferrerRequestBuilder when successful
func (m *ItemRemoveBlockedReferrerRequestBuilder) WithUrl(rawUrl string)(*ItemRemoveBlockedReferrerRequestBuilder) {
    return NewItemRemoveBlockedReferrerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
