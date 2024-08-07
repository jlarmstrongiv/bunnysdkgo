package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemRemoveBlockedIpRequestBuilder builds and executes requests for operations under \pullzone\{-id}\removeBlockedIp
type ItemRemoveBlockedIpRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemRemoveBlockedIpRequestBuilderInternal instantiates a new ItemRemoveBlockedIpRequestBuilder and sets the default values.
func NewItemRemoveBlockedIpRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoveBlockedIpRequestBuilder) {
    m := &ItemRemoveBlockedIpRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/removeBlockedIp", pathParameters),
    }
    return m
}
// NewItemRemoveBlockedIpRequestBuilder instantiates a new ItemRemoveBlockedIpRequestBuilder and sets the default values.
func NewItemRemoveBlockedIpRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoveBlockedIpRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRemoveBlockedIpRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [RemoveBlockedIp API Docs](https://docs.bunny.net/reference/pullzonepublic_removeblockedip)
func (m *ItemRemoveBlockedIpRequestBuilder) Post(ctx context.Context, body ItemRemoveBlockedIpPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPostRequestInformation [RemoveBlockedIp API Docs](https://docs.bunny.net/reference/pullzonepublic_removeblockedip)
// returns a *RequestInformation when successful
func (m *ItemRemoveBlockedIpRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemRemoveBlockedIpPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRemoveBlockedIpRequestBuilder when successful
func (m *ItemRemoveBlockedIpRequestBuilder) WithUrl(rawUrl string)(*ItemRemoveBlockedIpRequestBuilder) {
    return NewItemRemoveBlockedIpRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
