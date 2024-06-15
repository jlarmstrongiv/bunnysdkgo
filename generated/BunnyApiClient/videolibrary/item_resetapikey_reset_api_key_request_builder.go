package videolibrary

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemResetapikeyResetApiKeyRequestBuilder builds and executes requests for operations under \videolibrary\{id}\resetApiKey
type ItemResetapikeyResetApiKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemResetapikeyResetApiKeyRequestBuilderInternal instantiates a new ItemResetapikeyResetApiKeyRequestBuilder and sets the default values.
func NewItemResetapikeyResetApiKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResetapikeyResetApiKeyRequestBuilder) {
    m := &ItemResetapikeyResetApiKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/videolibrary/{id}/resetApiKey", pathParameters),
    }
    return m
}
// NewItemResetapikeyResetApiKeyRequestBuilder instantiates a new ItemResetapikeyResetApiKeyRequestBuilder and sets the default values.
func NewItemResetapikeyResetApiKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResetapikeyResetApiKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemResetapikeyResetApiKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [ResetPasswordPath API Docs](https://docs.bunny.net/reference/videolibrarypublic_resetpassword2)
func (m *ItemResetapikeyResetApiKeyRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPostRequestInformation [ResetPasswordPath API Docs](https://docs.bunny.net/reference/videolibrarypublic_resetpassword2)
// returns a *RequestInformation when successful
func (m *ItemResetapikeyResetApiKeyRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemResetapikeyResetApiKeyRequestBuilder when successful
func (m *ItemResetapikeyResetApiKeyRequestBuilder) WithUrl(rawUrl string)(*ItemResetapikeyResetApiKeyRequestBuilder) {
    return NewItemResetapikeyResetApiKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
