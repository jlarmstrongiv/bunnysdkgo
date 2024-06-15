package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LoadfreecertificateLoadFreeCertificateRequestBuilder builds and executes requests for operations under \pullzone\loadFreeCertificate
type LoadfreecertificateLoadFreeCertificateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LoadfreecertificateLoadFreeCertificateRequestBuilderGetQueryParameters [LoadFreeCertificate API Docs](https://docs.bunny.net/reference/pullzonepublic_loadfreecertificate)
type LoadfreecertificateLoadFreeCertificateRequestBuilderGetQueryParameters struct {
    // The hostname that the certificate will be loaded for
    Hostname *string `uriparametername:"hostname"`
}
// NewLoadfreecertificateLoadFreeCertificateRequestBuilderInternal instantiates a new LoadfreecertificateLoadFreeCertificateRequestBuilder and sets the default values.
func NewLoadfreecertificateLoadFreeCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LoadfreecertificateLoadFreeCertificateRequestBuilder) {
    m := &LoadfreecertificateLoadFreeCertificateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/loadFreeCertificate?hostname={hostname}", pathParameters),
    }
    return m
}
// NewLoadfreecertificateLoadFreeCertificateRequestBuilder instantiates a new LoadfreecertificateLoadFreeCertificateRequestBuilder and sets the default values.
func NewLoadfreecertificateLoadFreeCertificateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LoadfreecertificateLoadFreeCertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLoadfreecertificateLoadFreeCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [LoadFreeCertificate API Docs](https://docs.bunny.net/reference/pullzonepublic_loadfreecertificate)
// returns a []byte when successful
func (m *LoadfreecertificateLoadFreeCertificateRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[LoadfreecertificateLoadFreeCertificateRequestBuilderGetQueryParameters])([]byte, error) {
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
func (m *LoadfreecertificateLoadFreeCertificateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[LoadfreecertificateLoadFreeCertificateRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LoadfreecertificateLoadFreeCertificateRequestBuilder when successful
func (m *LoadfreecertificateLoadFreeCertificateRequestBuilder) WithUrl(rawUrl string)(*LoadfreecertificateLoadFreeCertificateRequestBuilder) {
    return NewLoadfreecertificateLoadFreeCertificateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
