package videolibrary

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemRemoveAllowedReferrerRequestBuilder builds and executes requests for operations under \videolibrary\{id}\removeAllowedReferrer
type ItemRemoveAllowedReferrerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemRemoveAllowedReferrerRequestBuilderInternal instantiates a new ItemRemoveAllowedReferrerRequestBuilder and sets the default values.
func NewItemRemoveAllowedReferrerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoveAllowedReferrerRequestBuilder) {
    m := &ItemRemoveAllowedReferrerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/videolibrary/{id}/removeAllowedReferrer", pathParameters),
    }
    return m
}
// NewItemRemoveAllowedReferrerRequestBuilder instantiates a new ItemRemoveAllowedReferrerRequestBuilder and sets the default values.
func NewItemRemoveAllowedReferrerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoveAllowedReferrerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRemoveAllowedReferrerRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [RemoveAllowedReferer API Docs](https://docs.bunny.net/reference/videolibrarypublic_removeallowedreferrer)
func (m *ItemRemoveAllowedReferrerRequestBuilder) Post(ctx context.Context, body ItemRemoveAllowedReferrerPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
func (m *ItemRemoveAllowedReferrerRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemRemoveAllowedReferrerPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRemoveAllowedReferrerRequestBuilder when successful
func (m *ItemRemoveAllowedReferrerRequestBuilder) WithUrl(rawUrl string)(*ItemRemoveAllowedReferrerRequestBuilder) {
    return NewItemRemoveAllowedReferrerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
