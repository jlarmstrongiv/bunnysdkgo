package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemRemovecertificateRemoveCertificateRequestBuilder builds and executes requests for operations under \pullzone\{-id}\removeCertificate
type ItemRemovecertificateRemoveCertificateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemRemovecertificateRemoveCertificateRequestBuilderInternal instantiates a new ItemRemovecertificateRemoveCertificateRequestBuilder and sets the default values.
func NewItemRemovecertificateRemoveCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemovecertificateRemoveCertificateRequestBuilder) {
    m := &ItemRemovecertificateRemoveCertificateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/removeCertificate", pathParameters),
    }
    return m
}
// NewItemRemovecertificateRemoveCertificateRequestBuilder instantiates a new ItemRemovecertificateRemoveCertificateRequestBuilder and sets the default values.
func NewItemRemovecertificateRemoveCertificateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemovecertificateRemoveCertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRemovecertificateRemoveCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete [RemoveCertificate API Docs](https://docs.bunny.net/reference/pullzonepublic_removecertificate)
func (m *ItemRemovecertificateRemoveCertificateRequestBuilder) Delete(ctx context.Context, body ItemRemovecertificateRemoveCertificateDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToDeleteRequestInformation [RemoveCertificate API Docs](https://docs.bunny.net/reference/pullzonepublic_removecertificate)
// returns a *RequestInformation when successful
func (m *ItemRemovecertificateRemoveCertificateRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body ItemRemovecertificateRemoveCertificateDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRemovecertificateRemoveCertificateRequestBuilder when successful
func (m *ItemRemovecertificateRemoveCertificateRequestBuilder) WithUrl(rawUrl string)(*ItemRemovecertificateRemoveCertificateRequestBuilder) {
    return NewItemRemovecertificateRemoveCertificateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
