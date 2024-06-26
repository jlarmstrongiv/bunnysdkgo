package library

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i828e151d2f6e3625c9020e4e883f6a5feb6568e66aa48c9ecf7c18565695df71 "github.com/jlarmstrongiv/bunnysdkgo/stream_api_client/models/managevideos/getvideoplaydata"
)

// ItemVideosItemPlayRequestBuilder builds and executes requests for operations under \library\{libraryId}\videos\{videoId}\play
type ItemVideosItemPlayRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemVideosItemPlayRequestBuilderGetQueryParameters [GetVideoPlayData API Docs](https://docs.bunny.net/reference/video_getvideoplaydata)
type ItemVideosItemPlayRequestBuilderGetQueryParameters struct {
    Expires *int64 `uriparametername:"expires"`
    Token *string `uriparametername:"token"`
}
// NewItemVideosItemPlayRequestBuilderInternal instantiates a new ItemVideosItemPlayRequestBuilder and sets the default values.
func NewItemVideosItemPlayRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemPlayRequestBuilder) {
    m := &ItemVideosItemPlayRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library/{libraryId}/videos/{videoId}/play{?expires*,token*}", pathParameters),
    }
    return m
}
// NewItemVideosItemPlayRequestBuilder instantiates a new ItemVideosItemPlayRequestBuilder and sets the default values.
func NewItemVideosItemPlayRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemPlayRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVideosItemPlayRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [GetVideoPlayData API Docs](https://docs.bunny.net/reference/video_getvideoplaydata)
// returns a VideoPlayDataable when successful
func (m *ItemVideosItemPlayRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemVideosItemPlayRequestBuilderGetQueryParameters])(i828e151d2f6e3625c9020e4e883f6a5feb6568e66aa48c9ecf7c18565695df71.VideoPlayDataable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i828e151d2f6e3625c9020e4e883f6a5feb6568e66aa48c9ecf7c18565695df71.CreateVideoPlayDataFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i828e151d2f6e3625c9020e4e883f6a5feb6568e66aa48c9ecf7c18565695df71.VideoPlayDataable), nil
}
// ToGetRequestInformation [GetVideoPlayData API Docs](https://docs.bunny.net/reference/video_getvideoplaydata)
// returns a *RequestInformation when successful
func (m *ItemVideosItemPlayRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemVideosItemPlayRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemVideosItemPlayRequestBuilder when successful
func (m *ItemVideosItemPlayRequestBuilder) WithUrl(rawUrl string)(*ItemVideosItemPlayRequestBuilder) {
    return NewItemVideosItemPlayRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
