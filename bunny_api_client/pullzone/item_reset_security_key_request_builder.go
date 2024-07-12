package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemResetSecurityKeyRequestBuilder builds and executes requests for operations under \pullzone\{-id}\resetSecurityKey
type ItemResetSecurityKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemResetSecurityKeyRequestBuilderInternal instantiates a new ItemResetSecurityKeyRequestBuilder and sets the default values.
func NewItemResetSecurityKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResetSecurityKeyRequestBuilder) {
    m := &ItemResetSecurityKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/resetSecurityKey", pathParameters),
    }
    return m
}
// NewItemResetSecurityKeyRequestBuilder instantiates a new ItemResetSecurityKeyRequestBuilder and sets the default values.
func NewItemResetSecurityKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResetSecurityKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemResetSecurityKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [ResetTokenKey API Docs](https://docs.bunny.net/reference/pullzonepublic_resetsecuritykey)
func (m *ItemResetSecurityKeyRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation [ResetTokenKey API Docs](https://docs.bunny.net/reference/pullzonepublic_resetsecuritykey)
// returns a *RequestInformation when successful
func (m *ItemResetSecurityKeyRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemResetSecurityKeyRequestBuilder when successful
func (m *ItemResetSecurityKeyRequestBuilder) WithUrl(rawUrl string)(*ItemResetSecurityKeyRequestBuilder) {
    return NewItemResetSecurityKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
