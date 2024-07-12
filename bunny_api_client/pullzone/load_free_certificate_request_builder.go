package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LoadFreeCertificateRequestBuilder builds and executes requests for operations under \pullzone\loadFreeCertificate
type LoadFreeCertificateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LoadFreeCertificateRequestBuilderGetQueryParameters [LoadFreeCertificate API Docs](https://docs.bunny.net/reference/pullzonepublic_loadfreecertificate)
type LoadFreeCertificateRequestBuilderGetQueryParameters struct {
    // The hostname that the certificate will be loaded for
    Hostname *string `uriparametername:"hostname"`
}
// NewLoadFreeCertificateRequestBuilderInternal instantiates a new LoadFreeCertificateRequestBuilder and sets the default values.
func NewLoadFreeCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LoadFreeCertificateRequestBuilder) {
    m := &LoadFreeCertificateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/loadFreeCertificate?hostname={hostname}", pathParameters),
    }
    return m
}
// NewLoadFreeCertificateRequestBuilder instantiates a new LoadFreeCertificateRequestBuilder and sets the default values.
func NewLoadFreeCertificateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LoadFreeCertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLoadFreeCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [LoadFreeCertificate API Docs](https://docs.bunny.net/reference/pullzonepublic_loadfreecertificate)
// returns a []byte when successful
func (m *LoadFreeCertificateRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[LoadFreeCertificateRequestBuilderGetQueryParameters])([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation [LoadFreeCertificate API Docs](https://docs.bunny.net/reference/pullzonepublic_loadfreecertificate)
// returns a *RequestInformation when successful
func (m *LoadFreeCertificateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[LoadFreeCertificateRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LoadFreeCertificateRequestBuilder when successful
func (m *LoadFreeCertificateRequestBuilder) WithUrl(rawUrl string)(*LoadFreeCertificateRequestBuilder) {
    return NewLoadFreeCertificateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
