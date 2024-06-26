package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAddblockedipAddBlockedIpRequestBuilder builds and executes requests for operations under \pullzone\{-id}\addBlockedIp
type ItemAddblockedipAddBlockedIpRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemAddblockedipAddBlockedIpRequestBuilderInternal instantiates a new ItemAddblockedipAddBlockedIpRequestBuilder and sets the default values.
func NewItemAddblockedipAddBlockedIpRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddblockedipAddBlockedIpRequestBuilder) {
    m := &ItemAddblockedipAddBlockedIpRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/addBlockedIp", pathParameters),
    }
    return m
}
// NewItemAddblockedipAddBlockedIpRequestBuilder instantiates a new ItemAddblockedipAddBlockedIpRequestBuilder and sets the default values.
func NewItemAddblockedipAddBlockedIpRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddblockedipAddBlockedIpRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAddblockedipAddBlockedIpRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [AddBlockedIp API Docs](https://docs.bunny.net/reference/pullzonepublic_addblockedip)
func (m *ItemAddblockedipAddBlockedIpRequestBuilder) Post(ctx context.Context, body ItemAddblockedipAddBlockedIpPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPostRequestInformation [AddBlockedIp API Docs](https://docs.bunny.net/reference/pullzonepublic_addblockedip)
// returns a *RequestInformation when successful
func (m *ItemAddblockedipAddBlockedIpRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemAddblockedipAddBlockedIpPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAddblockedipAddBlockedIpRequestBuilder when successful
func (m *ItemAddblockedipAddBlockedIpRequestBuilder) WithUrl(rawUrl string)(*ItemAddblockedipAddBlockedIpRequestBuilder) {
    return NewItemAddblockedipAddBlockedIpRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
