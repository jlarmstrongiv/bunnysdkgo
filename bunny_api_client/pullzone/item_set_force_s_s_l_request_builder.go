package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSetForceSSLRequestBuilder builds and executes requests for operations under \pullzone\{-id}\setForceSSL
type ItemSetForceSSLRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemSetForceSSLRequestBuilderInternal instantiates a new ItemSetForceSSLRequestBuilder and sets the default values.
func NewItemSetForceSSLRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSetForceSSLRequestBuilder) {
    m := &ItemSetForceSSLRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/setForceSSL", pathParameters),
    }
    return m
}
// NewItemSetForceSSLRequestBuilder instantiates a new ItemSetForceSSLRequestBuilder and sets the default values.
func NewItemSetForceSSLRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSetForceSSLRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSetForceSSLRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [SetForceSsl API Docs](https://docs.bunny.net/reference/pullzonepublic_setforcessl)
func (m *ItemSetForceSSLRequestBuilder) Post(ctx context.Context, body ItemSetForceSSLPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
func (m *ItemSetForceSSLRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSetForceSSLPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSetForceSSLRequestBuilder when successful
func (m *ItemSetForceSSLRequestBuilder) WithUrl(rawUrl string)(*ItemSetForceSSLRequestBuilder) {
    return NewItemSetForceSSLRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
