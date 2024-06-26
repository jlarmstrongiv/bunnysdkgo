package videolibrary

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie59f1fc195391a433ed92a6780c54480efff53cfc74cc731d5a61ccfbf9645d9 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/streamvideolibrary"
)

// VideolibraryRequestBuilder builds and executes requests for operations under \videolibrary
type VideolibraryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VideolibraryRequestBuilderGetQueryParameters [ListVideoLibraries API Docs](https://docs.bunny.net/reference/videolibrarypublic_index)
type VideolibraryRequestBuilderGetQueryParameters struct {
    IncludeAccessKey *bool `uriparametername:"includeAccessKey"`
    Page *int32 `uriparametername:"page"`
    PerPage *int32 `uriparametername:"perPage"`
    // The search term that will be used to filter the results
    Search *string `uriparametername:"search"`
}
// ById gets an item from the github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client.videolibrary.item collection
// returns a *VideolibraryItemRequestBuilder when successful
func (m *VideolibraryRequestBuilder) ById(id int64)(*VideolibraryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(id, 10)
    return NewVideolibraryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVideolibraryRequestBuilderInternal instantiates a new VideolibraryRequestBuilder and sets the default values.
func NewVideolibraryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VideolibraryRequestBuilder) {
    m := &VideolibraryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/videolibrary?includeAccessKey={includeAccessKey}&page={page}&perPage={perPage}&search={search}", pathParameters),
    }
    return m
}
// NewVideolibraryRequestBuilder instantiates a new VideolibraryRequestBuilder and sets the default values.
func NewVideolibraryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VideolibraryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVideolibraryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [ListVideoLibraries API Docs](https://docs.bunny.net/reference/videolibrarypublic_index)
// returns a VideolibraryGetResponseable when successful
func (m *VideolibraryRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[VideolibraryRequestBuilderGetQueryParameters])(VideolibraryGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVideolibraryGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VideolibraryGetResponseable), nil
}
// Languages the languages property
// returns a *LanguagesRequestBuilder when successful
func (m *VideolibraryRequestBuilder) Languages()(*LanguagesRequestBuilder) {
    return NewLanguagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post [AddVideoLibrary API Docs](https://docs.bunny.net/reference/videolibrarypublic_add)
// returns a VideoLibraryable when successful
func (m *VideolibraryRequestBuilder) Post(ctx context.Context, body VideolibraryPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie59f1fc195391a433ed92a6780c54480efff53cfc74cc731d5a61ccfbf9645d9.VideoLibraryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie59f1fc195391a433ed92a6780c54480efff53cfc74cc731d5a61ccfbf9645d9.CreateVideoLibraryFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie59f1fc195391a433ed92a6780c54480efff53cfc74cc731d5a61ccfbf9645d9.VideoLibraryable), nil
}
// ResetApiKey the resetApiKey property
// returns a *ResetapikeyResetApiKeyRequestBuilder when successful
func (m *VideolibraryRequestBuilder) ResetApiKey()(*ResetapikeyResetApiKeyRequestBuilder) {
    return NewResetapikeyResetApiKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation [ListVideoLibraries API Docs](https://docs.bunny.net/reference/videolibrarypublic_index)
// returns a *RequestInformation when successful
func (m *VideolibraryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[VideolibraryRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation [AddVideoLibrary API Docs](https://docs.bunny.net/reference/videolibrarypublic_add)
// returns a *RequestInformation when successful
func (m *VideolibraryRequestBuilder) ToPostRequestInformation(ctx context.Context, body VideolibraryPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/videolibrary", m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VideolibraryRequestBuilder when successful
func (m *VideolibraryRequestBuilder) WithUrl(rawUrl string)(*VideolibraryRequestBuilder) {
    return NewVideolibraryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
