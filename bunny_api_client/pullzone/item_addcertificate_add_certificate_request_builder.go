package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAddcertificateAddCertificateRequestBuilder builds and executes requests for operations under \pullzone\{-id}\addCertificate
type ItemAddcertificateAddCertificateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemAddcertificateAddCertificateRequestBuilderInternal instantiates a new ItemAddcertificateAddCertificateRequestBuilder and sets the default values.
func NewItemAddcertificateAddCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddcertificateAddCertificateRequestBuilder) {
    m := &ItemAddcertificateAddCertificateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/addCertificate", pathParameters),
    }
    return m
}
// NewItemAddcertificateAddCertificateRequestBuilder instantiates a new ItemAddcertificateAddCertificateRequestBuilder and sets the default values.
func NewItemAddcertificateAddCertificateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddcertificateAddCertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAddcertificateAddCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [AddCustomCertificate API Docs](https://docs.bunny.net/reference/pullzonepublic_addcertificate)
func (m *ItemAddcertificateAddCertificateRequestBuilder) Post(ctx context.Context, body ItemAddcertificateAddCertificatePostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPostRequestInformation [AddCustomCertificate API Docs](https://docs.bunny.net/reference/pullzonepublic_addcertificate)
// returns a *RequestInformation when successful
func (m *ItemAddcertificateAddCertificateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemAddcertificateAddCertificatePostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAddcertificateAddCertificateRequestBuilder when successful
func (m *ItemAddcertificateAddCertificateRequestBuilder) WithUrl(rawUrl string)(*ItemAddcertificateAddCertificateRequestBuilder) {
    return NewItemAddcertificateAddCertificateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
