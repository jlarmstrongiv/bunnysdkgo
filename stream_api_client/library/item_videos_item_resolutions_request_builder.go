package library

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i47924754b506d970f97cf36a9ead38c4920c571fc0694d201c282131a43b83b8 "github.com/jlarmstrongiv/bunnysdkgo/stream_api_client/models/managevideos/videoresolutionsinfo"
)

// ItemVideosItemResolutionsRequestBuilder builds and executes requests for operations under \library\{libraryId}\videos\{videoId}\resolutions
type ItemVideosItemResolutionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Cleanup the cleanup property
// returns a *ItemVideosItemResolutionsCleanupRequestBuilder when successful
func (m *ItemVideosItemResolutionsRequestBuilder) Cleanup()(*ItemVideosItemResolutionsCleanupRequestBuilder) {
    return NewItemVideosItemResolutionsCleanupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemVideosItemResolutionsRequestBuilderInternal instantiates a new ItemVideosItemResolutionsRequestBuilder and sets the default values.
func NewItemVideosItemResolutionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemResolutionsRequestBuilder) {
    m := &ItemVideosItemResolutionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library/{libraryId}/videos/{videoId}/resolutions", pathParameters),
    }
    return m
}
// NewItemVideosItemResolutionsRequestBuilder instantiates a new ItemVideosItemResolutionsRequestBuilder and sets the default values.
func NewItemVideosItemResolutionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemResolutionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVideosItemResolutionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [VideoResolutionsInfo API Docs](https://docs.bunny.net/reference/video_getvideoresolutions)
// returns a VideoResolutionsInfoResultable when successful
func (m *ItemVideosItemResolutionsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i47924754b506d970f97cf36a9ead38c4920c571fc0694d201c282131a43b83b8.VideoResolutionsInfoResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i47924754b506d970f97cf36a9ead38c4920c571fc0694d201c282131a43b83b8.CreateVideoResolutionsInfoResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i47924754b506d970f97cf36a9ead38c4920c571fc0694d201c282131a43b83b8.VideoResolutionsInfoResultable), nil
}
// ToGetRequestInformation [VideoResolutionsInfo API Docs](https://docs.bunny.net/reference/video_getvideoresolutions)
// returns a *RequestInformation when successful
func (m *ItemVideosItemResolutionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemVideosItemResolutionsRequestBuilder when successful
func (m *ItemVideosItemResolutionsRequestBuilder) WithUrl(rawUrl string)(*ItemVideosItemResolutionsRequestBuilder) {
    return NewItemVideosItemResolutionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
