package library

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234 "github.com/jlarmstrongiv/bunnysdkgo/stream_api_client/models/managevideos"
)

// ItemVideosItemReencodeRequestBuilder builds and executes requests for operations under \library\{libraryId}\videos\{videoId}\reencode
type ItemVideosItemReencodeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemVideosItemReencodeRequestBuilderInternal instantiates a new ItemVideosItemReencodeRequestBuilder and sets the default values.
func NewItemVideosItemReencodeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemReencodeRequestBuilder) {
    m := &ItemVideosItemReencodeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library/{libraryId}/videos/{videoId}/reencode", pathParameters),
    }
    return m
}
// NewItemVideosItemReencodeRequestBuilder instantiates a new ItemVideosItemReencodeRequestBuilder and sets the default values.
func NewItemVideosItemReencodeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemReencodeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVideosItemReencodeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [ReencodeVideo API Docs](https://docs.bunny.net/reference/video_reencodevideo)
// returns a Videoable when successful
func (m *ItemVideosItemReencodeRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234.Videoable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation [ReencodeVideo API Docs](https://docs.bunny.net/reference/video_reencodevideo)
// returns a *RequestInformation when successful
func (m *ItemVideosItemReencodeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemVideosItemReencodeRequestBuilder when successful
func (m *ItemVideosItemReencodeRequestBuilder) WithUrl(rawUrl string)(*ItemVideosItemReencodeRequestBuilder) {
    return NewItemVideosItemReencodeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
