package videolibrary

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie59f1fc195391a433ed92a6780c54480efff53cfc74cc731d5a61ccfbf9645d9 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/streamvideolibrary"
)

// VideolibraryItemRequestBuilder builds and executes requests for operations under \videolibrary\{id}
type VideolibraryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VideolibraryItemRequestBuilderGetQueryParameters [GetVideoLibrary API Docs](https://docs.bunny.net/reference/videolibrarypublic_index2)
type VideolibraryItemRequestBuilderGetQueryParameters struct {
    IncludeAccessKey *bool `uriparametername:"includeAccessKey"`
}
// AddAllowedReferrer the addAllowedReferrer property
// returns a *ItemAddallowedreferrerAddAllowedReferrerRequestBuilder when successful
func (m *VideolibraryItemRequestBuilder) AddAllowedReferrer()(*ItemAddallowedreferrerAddAllowedReferrerRequestBuilder) {
    return NewItemAddallowedreferrerAddAllowedReferrerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddBlockedReferrer the addBlockedReferrer property
// returns a *ItemAddblockedreferrerAddBlockedReferrerRequestBuilder when successful
func (m *VideolibraryItemRequestBuilder) AddBlockedReferrer()(*ItemAddblockedreferrerAddBlockedReferrerRequestBuilder) {
    return NewItemAddblockedreferrerAddBlockedReferrerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewVideolibraryItemRequestBuilderInternal instantiates a new VideolibraryItemRequestBuilder and sets the default values.
func NewVideolibraryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VideolibraryItemRequestBuilder) {
    m := &VideolibraryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/videolibrary/{id}?includeAccessKey={includeAccessKey}", pathParameters),
    }
    return m
}
// NewVideolibraryItemRequestBuilder instantiates a new VideolibraryItemRequestBuilder and sets the default values.
func NewVideolibraryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VideolibraryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVideolibraryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete [DeleteVideoLibrary API Docs](https://docs.bunny.net/reference/videolibrarypublic_delete)
func (m *VideolibraryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get [GetVideoLibrary API Docs](https://docs.bunny.net/reference/videolibrarypublic_index2)
// returns a VideoLibraryable when successful
func (m *VideolibraryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[VideolibraryItemRequestBuilderGetQueryParameters])(ie59f1fc195391a433ed92a6780c54480efff53cfc74cc731d5a61ccfbf9645d9.VideoLibraryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Post [UpdateVideoLibrary API Docs](https://docs.bunny.net/reference/videolibrarypublic_update)
// returns a VideoLibraryable when successful
func (m *VideolibraryItemRequestBuilder) Post(ctx context.Context, body ItemVideolibraryPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie59f1fc195391a433ed92a6780c54480efff53cfc74cc731d5a61ccfbf9645d9.VideoLibraryable, error) {
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
// RemoveAllowedReferrer the removeAllowedReferrer property
// returns a *ItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder when successful
func (m *VideolibraryItemRequestBuilder) RemoveAllowedReferrer()(*ItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder) {
    return NewItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveBlockedReferrer the removeBlockedReferrer property
// returns a *ItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder when successful
func (m *VideolibraryItemRequestBuilder) RemoveBlockedReferrer()(*ItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder) {
    return NewItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResetApiKey the resetApiKey property
// returns a *ItemResetapikeyResetApiKeyRequestBuilder when successful
func (m *VideolibraryItemRequestBuilder) ResetApiKey()(*ItemResetapikeyResetApiKeyRequestBuilder) {
    return NewItemResetapikeyResetApiKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation [DeleteVideoLibrary API Docs](https://docs.bunny.net/reference/videolibrarypublic_delete)
// returns a *RequestInformation when successful
func (m *VideolibraryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/videolibrary/{id}", m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// ToGetRequestInformation [GetVideoLibrary API Docs](https://docs.bunny.net/reference/videolibrarypublic_index2)
// returns a *RequestInformation when successful
func (m *VideolibraryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[VideolibraryItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation [UpdateVideoLibrary API Docs](https://docs.bunny.net/reference/videolibrarypublic_update)
// returns a *RequestInformation when successful
func (m *VideolibraryItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemVideolibraryPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/videolibrary/{id}", m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// Watermark the watermark property
// returns a *ItemWatermarkRequestBuilder when successful
func (m *VideolibraryItemRequestBuilder) Watermark()(*ItemWatermarkRequestBuilder) {
    return NewItemWatermarkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VideolibraryItemRequestBuilder when successful
func (m *VideolibraryItemRequestBuilder) WithUrl(rawUrl string)(*VideolibraryItemRequestBuilder) {
    return NewVideolibraryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
