package videolibrary

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder builds and executes requests for operations under \videolibrary\{id}\removeBlockedReferrer
type ItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilderInternal instantiates a new ItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder and sets the default values.
func NewItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder) {
    m := &ItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/videolibrary/{id}/removeBlockedReferrer", pathParameters),
    }
    return m
}
// NewItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder instantiates a new ItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder and sets the default values.
func NewItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [RemoveBlockedReferer API Docs](https://docs.bunny.net/reference/videolibrarypublic_removeblockedreferrer)
func (m *ItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder) Post(ctx context.Context, body ItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
func (m *ItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder when successful
func (m *ItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder) WithUrl(rawUrl string)(*ItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder) {
    return NewItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
