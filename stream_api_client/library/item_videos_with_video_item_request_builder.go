package library

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ic6a2b8141d365338b44f1e735d82ac4081bc99131cbeb3992752f448db804d1d "github.com/jlarmstrongiv/bunnysdkgo/stream_api_client/models"
    i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234 "github.com/jlarmstrongiv/bunnysdkgo/stream_api_client/models/managevideos"
)

// ItemVideosWithVideoItemRequestBuilder builds and executes requests for operations under \library\{libraryId}\videos\{videoId}
type ItemVideosWithVideoItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemVideosWithVideoItemRequestBuilderPutQueryParameters [UploadVideo API Docs](https://docs.bunny.net/reference/video_uploadvideo)
type ItemVideosWithVideoItemRequestBuilderPutQueryParameters struct {
    EnabledResolutions *string `uriparametername:"enabledResolutions"`
}
// Captions the captions property
// returns a *ItemVideosItemCaptionsRequestBuilder when successful
func (m *ItemVideosWithVideoItemRequestBuilder) Captions()(*ItemVideosItemCaptionsRequestBuilder) {
    return NewItemVideosItemCaptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemVideosWithVideoItemRequestBuilderInternal instantiates a new ItemVideosWithVideoItemRequestBuilder and sets the default values.
func NewItemVideosWithVideoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosWithVideoItemRequestBuilder) {
    m := &ItemVideosWithVideoItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library/{libraryId}/videos/{videoId}{?enabledResolutions}", pathParameters),
    }
    return m
}
// NewItemVideosWithVideoItemRequestBuilder instantiates a new ItemVideosWithVideoItemRequestBuilder and sets the default values.
func NewItemVideosWithVideoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosWithVideoItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVideosWithVideoItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete [DeleteVideo API Docs](https://docs.bunny.net/reference/video_deletevideo)
// returns a StructuredSuccessResponseable when successful
func (m *ItemVideosWithVideoItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ic6a2b8141d365338b44f1e735d82ac4081bc99131cbeb3992752f448db804d1d.StructuredSuccessResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic6a2b8141d365338b44f1e735d82ac4081bc99131cbeb3992752f448db804d1d.CreateStructuredSuccessResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic6a2b8141d365338b44f1e735d82ac4081bc99131cbeb3992752f448db804d1d.StructuredSuccessResponseable), nil
}
// Get [GetVideo API Docs](https://docs.bunny.net/reference/video_getvideo)
// returns a Videoable when successful
func (m *ItemVideosWithVideoItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234.Videoable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234.CreateVideoFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234.Videoable), nil
}
// Heatmap the heatmap property
// returns a *ItemVideosItemHeatmapRequestBuilder when successful
func (m *ItemVideosWithVideoItemRequestBuilder) Heatmap()(*ItemVideosItemHeatmapRequestBuilder) {
    return NewItemVideosItemHeatmapRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Play the play property
// returns a *ItemVideosItemPlayRequestBuilder when successful
func (m *ItemVideosWithVideoItemRequestBuilder) Play()(*ItemVideosItemPlayRequestBuilder) {
    return NewItemVideosItemPlayRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post [UpdateVideo API Docs](https://docs.bunny.net/reference/video_updatevideo)
// returns a StructuredSuccessResponseable when successful
func (m *ItemVideosWithVideoItemRequestBuilder) Post(ctx context.Context, body i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234.VideoCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ic6a2b8141d365338b44f1e735d82ac4081bc99131cbeb3992752f448db804d1d.StructuredSuccessResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic6a2b8141d365338b44f1e735d82ac4081bc99131cbeb3992752f448db804d1d.CreateStructuredSuccessResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic6a2b8141d365338b44f1e735d82ac4081bc99131cbeb3992752f448db804d1d.StructuredSuccessResponseable), nil
}
// Put [UploadVideo API Docs](https://docs.bunny.net/reference/video_uploadvideo)
// returns a StructuredSuccessResponseable when successful
func (m *ItemVideosWithVideoItemRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemVideosWithVideoItemRequestBuilderPutQueryParameters])(ic6a2b8141d365338b44f1e735d82ac4081bc99131cbeb3992752f448db804d1d.StructuredSuccessResponseable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic6a2b8141d365338b44f1e735d82ac4081bc99131cbeb3992752f448db804d1d.CreateStructuredSuccessResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic6a2b8141d365338b44f1e735d82ac4081bc99131cbeb3992752f448db804d1d.StructuredSuccessResponseable), nil
}
// Reencode the reencode property
// returns a *ItemVideosItemReencodeRequestBuilder when successful
func (m *ItemVideosWithVideoItemRequestBuilder) Reencode()(*ItemVideosItemReencodeRequestBuilder) {
    return NewItemVideosItemReencodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repackage the repackage property
// returns a *ItemVideosItemRepackageRequestBuilder when successful
func (m *ItemVideosWithVideoItemRequestBuilder) Repackage()(*ItemVideosItemRepackageRequestBuilder) {
    return NewItemVideosItemRepackageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Thumbnail the thumbnail property
// returns a *ItemVideosItemThumbnailRequestBuilder when successful
func (m *ItemVideosWithVideoItemRequestBuilder) Thumbnail()(*ItemVideosItemThumbnailRequestBuilder) {
    return NewItemVideosItemThumbnailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation [DeleteVideo API Docs](https://docs.bunny.net/reference/video_deletevideo)
// returns a *RequestInformation when successful
func (m *ItemVideosWithVideoItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation [GetVideo API Docs](https://docs.bunny.net/reference/video_getvideo)
// returns a *RequestInformation when successful
func (m *ItemVideosWithVideoItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation [UpdateVideo API Docs](https://docs.bunny.net/reference/video_updatevideo)
// returns a *RequestInformation when successful
func (m *ItemVideosWithVideoItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234.VideoCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToPutRequestInformation [UploadVideo API Docs](https://docs.bunny.net/reference/video_uploadvideo)
// returns a *RequestInformation when successful
func (m *ItemVideosWithVideoItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemVideosWithVideoItemRequestBuilderPutQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetStreamContentAndContentType(body, "application/octet-stream")
    return requestInfo, nil
}
// Transcribe the transcribe property
// returns a *ItemVideosItemTranscribeRequestBuilder when successful
func (m *ItemVideosWithVideoItemRequestBuilder) Transcribe()(*ItemVideosItemTranscribeRequestBuilder) {
    return NewItemVideosItemTranscribeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemVideosWithVideoItemRequestBuilder when successful
func (m *ItemVideosWithVideoItemRequestBuilder) WithUrl(rawUrl string)(*ItemVideosWithVideoItemRequestBuilder) {
    return NewItemVideosWithVideoItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
