package videolibrary

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder builds and executes requests for operations under \videolibrary\{id}\removeAllowedReferrer
type ItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilderInternal instantiates a new ItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder and sets the default values.
func NewItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder) {
    m := &ItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/videolibrary/{id}/removeAllowedReferrer", pathParameters),
    }
    return m
}
// NewItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder instantiates a new ItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder and sets the default values.
func NewItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [RemoveAllowedReferer API Docs](https://docs.bunny.net/reference/videolibrarypublic_removeallowedreferrer)
func (m *ItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder) Post(ctx context.Context, body ItemRemoveallowedreferrerRemoveAllowedReferrerPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPostRequestInformation [RemoveAllowedReferer API Docs](https://docs.bunny.net/reference/videolibrarypublic_removeallowedreferrer)
// returns a *RequestInformation when successful
func (m *ItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemRemoveallowedreferrerRemoveAllowedReferrerPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder when successful
func (m *ItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder) WithUrl(rawUrl string)(*ItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder) {
    return NewItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
