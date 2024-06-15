package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAddallowedreferrerAddAllowedReferrerRequestBuilder builds and executes requests for operations under \pullzone\{-id}\addAllowedReferrer
type ItemAddallowedreferrerAddAllowedReferrerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemAddallowedreferrerAddAllowedReferrerRequestBuilderInternal instantiates a new ItemAddallowedreferrerAddAllowedReferrerRequestBuilder and sets the default values.
func NewItemAddallowedreferrerAddAllowedReferrerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddallowedreferrerAddAllowedReferrerRequestBuilder) {
    m := &ItemAddallowedreferrerAddAllowedReferrerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/addAllowedReferrer", pathParameters),
    }
    return m
}
// NewItemAddallowedreferrerAddAllowedReferrerRequestBuilder instantiates a new ItemAddallowedreferrerAddAllowedReferrerRequestBuilder and sets the default values.
func NewItemAddallowedreferrerAddAllowedReferrerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddallowedreferrerAddAllowedReferrerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAddallowedreferrerAddAllowedReferrerRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [AddAllowedReferer API Docs](https://docs.bunny.net/reference/pullzonepublic_addallowedreferrer)
func (m *ItemAddallowedreferrerAddAllowedReferrerRequestBuilder) Post(ctx context.Context, body ItemAddallowedreferrerAddAllowedReferrerPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPostRequestInformation [AddAllowedReferer API Docs](https://docs.bunny.net/reference/pullzonepublic_addallowedreferrer)
// returns a *RequestInformation when successful
func (m *ItemAddallowedreferrerAddAllowedReferrerRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemAddallowedreferrerAddAllowedReferrerPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAddallowedreferrerAddAllowedReferrerRequestBuilder when successful
func (m *ItemAddallowedreferrerAddAllowedReferrerRequestBuilder) WithUrl(rawUrl string)(*ItemAddallowedreferrerAddAllowedReferrerRequestBuilder) {
    return NewItemAddallowedreferrerAddAllowedReferrerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
