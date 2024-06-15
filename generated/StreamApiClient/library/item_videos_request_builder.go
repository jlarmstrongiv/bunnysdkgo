package library

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i804737ef689cf33cc1996aaa50bf60df828fa9ff20878b7ccb1696459a4f4e28 "github.com/jlarmstrongiv/bunnysdkgo/generated/StreamApiClient/models/managevideos"
    i7b0ee2beff0fa8a7199afb13265886158631f5bb8c4eaac1b0549722fb71719c "github.com/jlarmstrongiv/bunnysdkgo/generated/StreamApiClient/library/item/videos"
)

// ItemVideosRequestBuilder builds and executes requests for operations under \library\{libraryId}\videos
type ItemVideosRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemVideosRequestBuilderGetQueryParameters [ListVideos API Docs](https://docs.bunny.net/reference/video_list)
type ItemVideosRequestBuilderGetQueryParameters struct {
    Collection *string `uriparametername:"collection"`
    ItemsPerPage *int32 `uriparametername:"itemsPerPage"`
    OrderBy *i7b0ee2beff0fa8a7199afb13265886158631f5bb8c4eaac1b0549722fb71719c.GetOrderByQueryParameterType `uriparametername:"orderBy"`
    Page *int32 `uriparametername:"page"`
    Search *string `uriparametername:"search"`
}
// ByVideoId gets an item from the github.com/jlarmstrongiv/bunnysdkgo/generated/StreamApiClient.library.item.videos.item collection
// returns a *ItemVideosWithVideoItemRequestBuilder when successful
func (m *ItemVideosRequestBuilder) ByVideoId(videoId string)(*ItemVideosWithVideoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if videoId != "" {
        urlTplParams["videoId"] = videoId
    }
    return NewItemVideosWithVideoItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemVideosRequestBuilderInternal instantiates a new ItemVideosRequestBuilder and sets the default values.
func NewItemVideosRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosRequestBuilder) {
    m := &ItemVideosRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library/{libraryId}/videos?itemsPerPage={itemsPerPage}&orderBy={orderBy}&page={page}{&collection*,search*}", pathParameters),
    }
    return m
}
// NewItemVideosRequestBuilder instantiates a new ItemVideosRequestBuilder and sets the default values.
func NewItemVideosRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVideosRequestBuilderInternal(urlParams, requestAdapter)
}
// Fetch the fetch property
// returns a *ItemVideosFetchRequestBuilder when successful
func (m *ItemVideosRequestBuilder) Fetch()(*ItemVideosFetchRequestBuilder) {
    return NewItemVideosFetchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get [ListVideos API Docs](https://docs.bunny.net/reference/video_list)
// returns a ItemVideosGetResponseable when successful
func (m *ItemVideosRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemVideosRequestBuilderGetQueryParameters])(ItemVideosGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemVideosGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemVideosGetResponseable), nil
}
// Post [CreateVideo API Docs](https://docs.bunny.net/reference/video_createvideo)
// returns a Videoable when successful
func (m *ItemVideosRequestBuilder) Post(ctx context.Context, body i804737ef689cf33cc1996aaa50bf60df828fa9ff20878b7ccb1696459a4f4e28.VideoCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i804737ef689cf33cc1996aaa50bf60df828fa9ff20878b7ccb1696459a4f4e28.Videoable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i804737ef689cf33cc1996aaa50bf60df828fa9ff20878b7ccb1696459a4f4e28.CreateVideoFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i804737ef689cf33cc1996aaa50bf60df828fa9ff20878b7ccb1696459a4f4e28.Videoable), nil
}
// ToGetRequestInformation [ListVideos API Docs](https://docs.bunny.net/reference/video_list)
// returns a *RequestInformation when successful
func (m *ItemVideosRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemVideosRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation [CreateVideo API Docs](https://docs.bunny.net/reference/video_createvideo)
// returns a *RequestInformation when successful
func (m *ItemVideosRequestBuilder) ToPostRequestInformation(ctx context.Context, body i804737ef689cf33cc1996aaa50bf60df828fa9ff20878b7ccb1696459a4f4e28.VideoCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/library/{libraryId}/videos", m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemVideosRequestBuilder when successful
func (m *ItemVideosRequestBuilder) WithUrl(rawUrl string)(*ItemVideosRequestBuilder) {
    return NewItemVideosRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
