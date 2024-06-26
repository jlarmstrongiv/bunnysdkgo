package compute

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ScriptItemReleasesRequestBuilder builds and executes requests for operations under \compute\script\{id}\releases
type ScriptItemReleasesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ScriptItemReleasesRequestBuilderGetQueryParameters [ListComputeScriptReleases API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_getreleases)
type ScriptItemReleasesRequestBuilderGetQueryParameters struct {
    Page *int32 `uriparametername:"page"`
    PerPage *int32 `uriparametername:"perPage"`
}
// NewScriptItemReleasesRequestBuilderInternal instantiates a new ScriptItemReleasesRequestBuilder and sets the default values.
func NewScriptItemReleasesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScriptItemReleasesRequestBuilder) {
    m := &ScriptItemReleasesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compute/script/{id}/releases?page={page}&perPage={perPage}", pathParameters),
    }
    return m
}
// NewScriptItemReleasesRequestBuilder instantiates a new ScriptItemReleasesRequestBuilder and sets the default values.
func NewScriptItemReleasesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScriptItemReleasesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScriptItemReleasesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [ListComputeScriptReleases API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_getreleases)
// returns a ScriptItemReleasesGetResponseable when successful
func (m *ScriptItemReleasesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ScriptItemReleasesRequestBuilderGetQueryParameters])(ScriptItemReleasesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateScriptItemReleasesGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ScriptItemReleasesGetResponseable), nil
}
// ToGetRequestInformation [ListComputeScriptReleases API Docs](https://docs.bunny.net/reference/computeedgescriptpublic_getreleases)
// returns a *RequestInformation when successful
func (m *ScriptItemReleasesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ScriptItemReleasesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ScriptItemReleasesRequestBuilder when successful
func (m *ScriptItemReleasesRequestBuilder) WithUrl(rawUrl string)(*ScriptItemReleasesRequestBuilder) {
    return NewScriptItemReleasesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
