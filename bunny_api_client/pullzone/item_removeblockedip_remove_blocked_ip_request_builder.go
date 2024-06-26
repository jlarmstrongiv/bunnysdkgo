package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemRemoveblockedipRemoveBlockedIpRequestBuilder builds and executes requests for operations under \pullzone\{-id}\removeBlockedIp
type ItemRemoveblockedipRemoveBlockedIpRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemRemoveblockedipRemoveBlockedIpRequestBuilderInternal instantiates a new ItemRemoveblockedipRemoveBlockedIpRequestBuilder and sets the default values.
func NewItemRemoveblockedipRemoveBlockedIpRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoveblockedipRemoveBlockedIpRequestBuilder) {
    m := &ItemRemoveblockedipRemoveBlockedIpRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/removeBlockedIp", pathParameters),
    }
    return m
}
// NewItemRemoveblockedipRemoveBlockedIpRequestBuilder instantiates a new ItemRemoveblockedipRemoveBlockedIpRequestBuilder and sets the default values.
func NewItemRemoveblockedipRemoveBlockedIpRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoveblockedipRemoveBlockedIpRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRemoveblockedipRemoveBlockedIpRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [RemoveBlockedIp API Docs](https://docs.bunny.net/reference/pullzonepublic_removeblockedip)
func (m *ItemRemoveblockedipRemoveBlockedIpRequestBuilder) Post(ctx context.Context, body ItemRemoveblockedipRemoveBlockedIpPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
func (m *ItemRemoveblockedipRemoveBlockedIpRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemRemoveblockedipRemoveBlockedIpPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRemoveblockedipRemoveBlockedIpRequestBuilder when successful
func (m *ItemRemoveblockedipRemoveBlockedIpRequestBuilder) WithUrl(rawUrl string)(*ItemRemoveblockedipRemoveBlockedIpRequestBuilder) {
    return NewItemRemoveblockedipRemoveBlockedIpRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
