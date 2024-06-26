package storagezone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemResetpasswordResetPasswordRequestBuilder builds and executes requests for operations under \storagezone\{id}\resetPassword
type ItemResetpasswordResetPasswordRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemResetpasswordResetPasswordRequestBuilderInternal instantiates a new ItemResetpasswordResetPasswordRequestBuilder and sets the default values.
func NewItemResetpasswordResetPasswordRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResetpasswordResetPasswordRequestBuilder) {
    m := &ItemResetpasswordResetPasswordRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storagezone/{id}/resetPassword", pathParameters),
    }
    return m
}
// NewItemResetpasswordResetPasswordRequestBuilder instantiates a new ItemResetpasswordResetPasswordRequestBuilder and sets the default values.
func NewItemResetpasswordResetPasswordRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResetpasswordResetPasswordRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemResetpasswordResetPasswordRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [ResetPassword API Docs](https://docs.bunny.net/reference/storagezonepublic_resetpassword)
func (m *ItemResetpasswordResetPasswordRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPostRequestInformation [ResetPassword API Docs](https://docs.bunny.net/reference/storagezonepublic_resetpassword)
// returns a *RequestInformation when successful
func (m *ItemResetpasswordResetPasswordRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemResetpasswordResetPasswordRequestBuilder when successful
func (m *ItemResetpasswordResetPasswordRequestBuilder) WithUrl(rawUrl string)(*ItemResetpasswordResetPasswordRequestBuilder) {
    return NewItemResetpasswordResetPasswordRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
