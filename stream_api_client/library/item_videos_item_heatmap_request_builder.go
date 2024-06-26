package library

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemVideosItemHeatmapRequestBuilder builds and executes requests for operations under \library\{libraryId}\videos\{videoId}\heatmap
type ItemVideosItemHeatmapRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemVideosItemHeatmapRequestBuilderInternal instantiates a new ItemVideosItemHeatmapRequestBuilder and sets the default values.
func NewItemVideosItemHeatmapRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemHeatmapRequestBuilder) {
    m := &ItemVideosItemHeatmapRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library/{libraryId}/videos/{videoId}/heatmap", pathParameters),
    }
    return m
}
// NewItemVideosItemHeatmapRequestBuilder instantiates a new ItemVideosItemHeatmapRequestBuilder and sets the default values.
func NewItemVideosItemHeatmapRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemHeatmapRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVideosItemHeatmapRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [GetVideoHeatmap API Docs](https://docs.bunny.net/reference/video_getvideoheatmap)
// returns a ItemVideosItemHeatmapGetResponseable when successful
func (m *ItemVideosItemHeatmapRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemVideosItemHeatmapGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemVideosItemHeatmapGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemVideosItemHeatmapGetResponseable), nil
}
// ToGetRequestInformation [GetVideoHeatmap API Docs](https://docs.bunny.net/reference/video_getvideoheatmap)
// returns a *RequestInformation when successful
func (m *ItemVideosItemHeatmapRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemVideosItemHeatmapRequestBuilder when successful
func (m *ItemVideosItemHeatmapRequestBuilder) WithUrl(rawUrl string)(*ItemVideosItemHeatmapRequestBuilder) {
    return NewItemVideosItemHeatmapRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
