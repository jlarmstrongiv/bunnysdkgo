package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemRemovehostnameRemoveHostnameRequestBuilder builds and executes requests for operations under \pullzone\{-id}\removeHostname
type ItemRemovehostnameRemoveHostnameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemRemovehostnameRemoveHostnameRequestBuilderInternal instantiates a new ItemRemovehostnameRemoveHostnameRequestBuilder and sets the default values.
func NewItemRemovehostnameRemoveHostnameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemovehostnameRemoveHostnameRequestBuilder) {
    m := &ItemRemovehostnameRemoveHostnameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/removeHostname", pathParameters),
    }
    return m
}
// NewItemRemovehostnameRemoveHostnameRequestBuilder instantiates a new ItemRemovehostnameRemoveHostnameRequestBuilder and sets the default values.
func NewItemRemovehostnameRemoveHostnameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemovehostnameRemoveHostnameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRemovehostnameRemoveHostnameRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete [RemoveCustomHostname API Docs](https://docs.bunny.net/reference/pullzonepublic_removehostname)
func (m *ItemRemovehostnameRemoveHostnameRequestBuilder) Delete(ctx context.Context, body ItemRemovehostnameRemoveHostnameDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
func (m *ItemRemovehostnameRemoveHostnameRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body ItemRemovehostnameRemoveHostnameDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRemovehostnameRemoveHostnameRequestBuilder when successful
func (m *ItemRemovehostnameRemoveHostnameRequestBuilder) WithUrl(rawUrl string)(*ItemRemovehostnameRemoveHostnameRequestBuilder) {
    return NewItemRemovehostnameRemoveHostnameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
