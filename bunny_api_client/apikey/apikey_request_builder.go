package apikey

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ApikeyRequestBuilder builds and executes requests for operations under \apikey
type ApikeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ApikeyRequestBuilderGetQueryParameters [ListApiKeys API Docs](https://docs.bunny.net/reference/apikeypublic_listapikeys)
type ApikeyRequestBuilderGetQueryParameters struct {
    Page *int32 `uriparametername:"page"`
    PerPage *int32 `uriparametername:"perPage"`
}
// NewApikeyRequestBuilderInternal instantiates a new ApikeyRequestBuilder and sets the default values.
func NewApikeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApikeyRequestBuilder) {
    m := &ApikeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/apikey?page={page}&perPage={perPage}", pathParameters),
    }
    return m
}
// NewApikeyRequestBuilder instantiates a new ApikeyRequestBuilder and sets the default values.
func NewApikeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApikeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApikeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [ListApiKeys API Docs](https://docs.bunny.net/reference/apikeypublic_listapikeys)
// returns a ApikeyGetResponseable when successful
func (m *ApikeyRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ApikeyRequestBuilderGetQueryParameters])(ApikeyGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateApikeyGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ApikeyGetResponseable), nil
}
// ToGetRequestInformation [ListApiKeys API Docs](https://docs.bunny.net/reference/apikeypublic_listapikeys)
// returns a *RequestInformation when successful
func (m *ApikeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ApikeyRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ApikeyRequestBuilder when successful
func (m *ApikeyRequestBuilder) WithUrl(rawUrl string)(*ApikeyRequestBuilder) {
    return NewApikeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
