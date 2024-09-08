package library

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ief80e4af15e1ae3976f892fc118dc6aaff4d47add746e056b912d7a9cfc17a6d "github.com/jlarmstrongiv/bunnysdkgo/stream_api_client/models/managevideos/cleanupunconfiguredresolutions"
)

// ItemVideosItemResolutionsCleanupRequestBuilder builds and executes requests for operations under \library\{libraryId}\videos\{videoId}\resolutions\cleanup
type ItemVideosItemResolutionsCleanupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemVideosItemResolutionsCleanupRequestBuilderPostQueryParameters [CleanupUnconfiguredResolutions API Docs](https://docs.bunny.net/reference/video_deleteresolutions)
type ItemVideosItemResolutionsCleanupRequestBuilderPostQueryParameters struct {
    DeleteMp4Files *bool `uriparametername:"deleteMp4Files"`
    DeleteNonConfiguredResolutions *bool `uriparametername:"deleteNonConfiguredResolutions"`
    DeleteOriginal *bool `uriparametername:"deleteOriginal"`
    ResolutionsToDelete *string `uriparametername:"resolutionsToDelete"`
}
// NewItemVideosItemResolutionsCleanupRequestBuilderInternal instantiates a new ItemVideosItemResolutionsCleanupRequestBuilder and sets the default values.
func NewItemVideosItemResolutionsCleanupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemResolutionsCleanupRequestBuilder) {
    m := &ItemVideosItemResolutionsCleanupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library/{libraryId}/videos/{videoId}/resolutions/cleanup{?deleteMp4Files,deleteNonConfiguredResolutions,deleteOriginal,resolutionsToDelete}", pathParameters),
    }
    return m
}
// NewItemVideosItemResolutionsCleanupRequestBuilder instantiates a new ItemVideosItemResolutionsCleanupRequestBuilder and sets the default values.
func NewItemVideosItemResolutionsCleanupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemResolutionsCleanupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVideosItemResolutionsCleanupRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [CleanupUnconfiguredResolutions API Docs](https://docs.bunny.net/reference/video_deleteresolutions)
// returns a CleanupUnconfiguredResolutionsResultable when successful
func (m *ItemVideosItemResolutionsCleanupRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemVideosItemResolutionsCleanupRequestBuilderPostQueryParameters])(ief80e4af15e1ae3976f892fc118dc6aaff4d47add746e056b912d7a9cfc17a6d.CleanupUnconfiguredResolutionsResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ief80e4af15e1ae3976f892fc118dc6aaff4d47add746e056b912d7a9cfc17a6d.CreateCleanupUnconfiguredResolutionsResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ief80e4af15e1ae3976f892fc118dc6aaff4d47add746e056b912d7a9cfc17a6d.CleanupUnconfiguredResolutionsResultable), nil
}
// ToPostRequestInformation [CleanupUnconfiguredResolutions API Docs](https://docs.bunny.net/reference/video_deleteresolutions)
// returns a *RequestInformation when successful
func (m *ItemVideosItemResolutionsCleanupRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemVideosItemResolutionsCleanupRequestBuilderPostQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemVideosItemResolutionsCleanupRequestBuilder when successful
func (m *ItemVideosItemResolutionsCleanupRequestBuilder) WithUrl(rawUrl string)(*ItemVideosItemResolutionsCleanupRequestBuilder) {
    return NewItemVideosItemResolutionsCleanupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
