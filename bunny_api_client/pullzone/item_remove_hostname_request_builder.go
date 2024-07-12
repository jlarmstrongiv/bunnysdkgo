package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemRemoveHostnameRequestBuilder builds and executes requests for operations under \pullzone\{-id}\removeHostname
type ItemRemoveHostnameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemRemoveHostnameRequestBuilderInternal instantiates a new ItemRemoveHostnameRequestBuilder and sets the default values.
func NewItemRemoveHostnameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoveHostnameRequestBuilder) {
    m := &ItemRemoveHostnameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/removeHostname", pathParameters),
    }
    return m
}
// NewItemRemoveHostnameRequestBuilder instantiates a new ItemRemoveHostnameRequestBuilder and sets the default values.
func NewItemRemoveHostnameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoveHostnameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRemoveHostnameRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete [RemoveCustomHostname API Docs](https://docs.bunny.net/reference/pullzonepublic_removehostname)
func (m *ItemRemoveHostnameRequestBuilder) Delete(ctx context.Context, body ItemRemoveHostnameDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation [RemoveCustomHostname API Docs](https://docs.bunny.net/reference/pullzonepublic_removehostname)
// returns a *RequestInformation when successful
func (m *ItemRemoveHostnameRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body ItemRemoveHostnameDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRemoveHostnameRequestBuilder when successful
func (m *ItemRemoveHostnameRequestBuilder) WithUrl(rawUrl string)(*ItemRemoveHostnameRequestBuilder) {
    return NewItemRemoveHostnameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
