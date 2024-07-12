package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAddHostnameRequestBuilder builds and executes requests for operations under \pullzone\{-id}\addHostname
type ItemAddHostnameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemAddHostnameRequestBuilderInternal instantiates a new ItemAddHostnameRequestBuilder and sets the default values.
func NewItemAddHostnameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddHostnameRequestBuilder) {
    m := &ItemAddHostnameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/addHostname", pathParameters),
    }
    return m
}
// NewItemAddHostnameRequestBuilder instantiates a new ItemAddHostnameRequestBuilder and sets the default values.
func NewItemAddHostnameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddHostnameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAddHostnameRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [AddCustomHostname API Docs](https://docs.bunny.net/reference/pullzonepublic_addhostname)
func (m *ItemAddHostnameRequestBuilder) Post(ctx context.Context, body ItemAddHostnamePostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPostRequestInformation [AddCustomHostname API Docs](https://docs.bunny.net/reference/pullzonepublic_addhostname)
// returns a *RequestInformation when successful
func (m *ItemAddHostnameRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemAddHostnamePostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAddHostnameRequestBuilder when successful
func (m *ItemAddHostnameRequestBuilder) WithUrl(rawUrl string)(*ItemAddHostnameRequestBuilder) {
    return NewItemAddHostnameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
