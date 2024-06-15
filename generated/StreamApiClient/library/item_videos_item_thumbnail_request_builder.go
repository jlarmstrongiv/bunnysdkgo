package library

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0 "github.com/jlarmstrongiv/bunnysdkgo/generated/StreamApiClient/models"
)

// ItemVideosItemThumbnailRequestBuilder builds and executes requests for operations under \library\{libraryId}\videos\{videoId}\thumbnail
type ItemVideosItemThumbnailRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemVideosItemThumbnailRequestBuilderPostQueryParameters [SetThumbnail API Docs](https://docs.bunny.net/reference/video_setthumbnail)
type ItemVideosItemThumbnailRequestBuilderPostQueryParameters struct {
    ThumbnailUrl *string `uriparametername:"thumbnailUrl"`
}
// NewItemVideosItemThumbnailRequestBuilderInternal instantiates a new ItemVideosItemThumbnailRequestBuilder and sets the default values.
func NewItemVideosItemThumbnailRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemThumbnailRequestBuilder) {
    m := &ItemVideosItemThumbnailRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library/{libraryId}/videos/{videoId}/thumbnail?thumbnailUrl={thumbnailUrl}", pathParameters),
    }
    return m
}
// NewItemVideosItemThumbnailRequestBuilder instantiates a new ItemVideosItemThumbnailRequestBuilder and sets the default values.
func NewItemVideosItemThumbnailRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemThumbnailRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVideosItemThumbnailRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [SetThumbnail API Docs](https://docs.bunny.net/reference/video_setthumbnail)
// returns a StructuredSuccessResponseable when successful
func (m *ItemVideosItemThumbnailRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemVideosItemThumbnailRequestBuilderPostQueryParameters])(i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0.StructuredSuccessResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0.CreateStructuredSuccessResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0.StructuredSuccessResponseable), nil
}
// ToPostRequestInformation [SetThumbnail API Docs](https://docs.bunny.net/reference/video_setthumbnail)
// returns a *RequestInformation when successful
func (m *ItemVideosItemThumbnailRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemVideosItemThumbnailRequestBuilderPostQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemVideosItemThumbnailRequestBuilder when successful
func (m *ItemVideosItemThumbnailRequestBuilder) WithUrl(rawUrl string)(*ItemVideosItemThumbnailRequestBuilder) {
    return NewItemVideosItemThumbnailRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
