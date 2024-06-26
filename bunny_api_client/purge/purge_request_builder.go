package purge

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PurgeRequestBuilder builds and executes requests for operations under \purge
type PurgeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PurgeRequestBuilderGetQueryParameters [PurgeUrlGet API Docs](https://docs.bunny.net/reference/purgepublic_index)
type PurgeRequestBuilderGetQueryParameters struct {
    Async *bool `uriparametername:"async"`
    HeaderName *string `uriparametername:"headerName"`
    HeaderValue *string `uriparametername:"headerValue"`
    Url *string `uriparametername:"url"`
}
// PurgeRequestBuilderPostQueryParameters [PurgeUrlPost API Docs](https://docs.bunny.net/reference/purgepublic_indexpost)
type PurgeRequestBuilderPostQueryParameters struct {
    Async *bool `uriparametername:"async"`
    Url *string `uriparametername:"url"`
}
// NewPurgeRequestBuilderInternal instantiates a new PurgeRequestBuilder and sets the default values.
func NewPurgeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PurgeRequestBuilder) {
    m := &PurgeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/purge?async={async}&headerName={headerName}&headerValue={headerValue}&url={url}", pathParameters),
    }
    return m
}
// NewPurgeRequestBuilder instantiates a new PurgeRequestBuilder and sets the default values.
func NewPurgeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PurgeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPurgeRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [PurgeUrlGet API Docs](https://docs.bunny.net/reference/purgepublic_index)
// returns a []byte when successful
func (m *PurgeRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[PurgeRequestBuilderGetQueryParameters])([]byte, error) {
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
// Post [PurgeUrlPost API Docs](https://docs.bunny.net/reference/purgepublic_indexpost)
// returns a []byte when successful
func (m *PurgeRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[PurgeRequestBuilderPostQueryParameters])([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToGetRequestInformation [PurgeUrlGet API Docs](https://docs.bunny.net/reference/purgepublic_index)
// returns a *RequestInformation when successful
func (m *PurgeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[PurgeRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// ToPostRequestInformation [PurgeUrlPost API Docs](https://docs.bunny.net/reference/purgepublic_indexpost)
// returns a *RequestInformation when successful
func (m *PurgeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[PurgeRequestBuilderPostQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/purge?async={async}&url={url}", m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PurgeRequestBuilder when successful
func (m *PurgeRequestBuilder) WithUrl(rawUrl string)(*PurgeRequestBuilder) {
    return NewPurgeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
