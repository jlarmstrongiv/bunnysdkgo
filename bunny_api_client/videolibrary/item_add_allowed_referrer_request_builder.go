package videolibrary

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAddAllowedReferrerRequestBuilder builds and executes requests for operations under \videolibrary\{id}\addAllowedReferrer
type ItemAddAllowedReferrerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemAddAllowedReferrerRequestBuilderInternal instantiates a new ItemAddAllowedReferrerRequestBuilder and sets the default values.
func NewItemAddAllowedReferrerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddAllowedReferrerRequestBuilder) {
    m := &ItemAddAllowedReferrerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/videolibrary/{id}/addAllowedReferrer", pathParameters),
    }
    return m
}
// NewItemAddAllowedReferrerRequestBuilder instantiates a new ItemAddAllowedReferrerRequestBuilder and sets the default values.
func NewItemAddAllowedReferrerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddAllowedReferrerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAddAllowedReferrerRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [AddAllowedReferer API Docs](https://docs.bunny.net/reference/videolibrarypublic_addallowedreferrer)
func (m *ItemAddAllowedReferrerRequestBuilder) Post(ctx context.Context, body ItemAddAllowedReferrerPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPostRequestInformation [AddAllowedReferer API Docs](https://docs.bunny.net/reference/videolibrarypublic_addallowedreferrer)
// returns a *RequestInformation when successful
func (m *ItemAddAllowedReferrerRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemAddAllowedReferrerPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAddAllowedReferrerRequestBuilder when successful
func (m *ItemAddAllowedReferrerRequestBuilder) WithUrl(rawUrl string)(*ItemAddAllowedReferrerRequestBuilder) {
    return NewItemAddAllowedReferrerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
