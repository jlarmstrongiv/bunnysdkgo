package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSetforcesslSetForceSSLRequestBuilder builds and executes requests for operations under \pullzone\{-id}\setForceSSL
type ItemSetforcesslSetForceSSLRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemSetforcesslSetForceSSLRequestBuilderInternal instantiates a new ItemSetforcesslSetForceSSLRequestBuilder and sets the default values.
func NewItemSetforcesslSetForceSSLRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSetforcesslSetForceSSLRequestBuilder) {
    m := &ItemSetforcesslSetForceSSLRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/setForceSSL", pathParameters),
    }
    return m
}
// NewItemSetforcesslSetForceSSLRequestBuilder instantiates a new ItemSetforcesslSetForceSSLRequestBuilder and sets the default values.
func NewItemSetforcesslSetForceSSLRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSetforcesslSetForceSSLRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSetforcesslSetForceSSLRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [SetForceSsl API Docs](https://docs.bunny.net/reference/pullzonepublic_setforcessl)
func (m *ItemSetforcesslSetForceSSLRequestBuilder) Post(ctx context.Context, body ItemSetforcesslSetForceSSLPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPostRequestInformation [SetForceSsl API Docs](https://docs.bunny.net/reference/pullzonepublic_setforcessl)
// returns a *RequestInformation when successful
func (m *ItemSetforcesslSetForceSSLRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSetforcesslSetForceSSLPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSetforcesslSetForceSSLRequestBuilder when successful
func (m *ItemSetforcesslSetForceSSLRequestBuilder) WithUrl(rawUrl string)(*ItemSetforcesslSetForceSSLRequestBuilder) {
    return NewItemSetforcesslSetForceSSLRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
