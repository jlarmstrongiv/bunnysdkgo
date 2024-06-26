package library

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234 "github.com/jlarmstrongiv/bunnysdkgo/stream_api_client/models/managevideos"
)

// ItemVideosItemRepackageRequestBuilder builds and executes requests for operations under \library\{libraryId}\videos\{videoId}\repackage
type ItemVideosItemRepackageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemVideosItemRepackageRequestBuilderPostQueryParameters [RepackageVideo API Docs](https://docs.bunny.net/reference/video_repackage)
type ItemVideosItemRepackageRequestBuilderPostQueryParameters struct {
    // Marks whether previous file versions should be kept in storage, allows for faster repackage later on. Default is true.
    KeepOriginalFiles *bool `uriparametername:"keepOriginalFiles"`
}
// NewItemVideosItemRepackageRequestBuilderInternal instantiates a new ItemVideosItemRepackageRequestBuilder and sets the default values.
func NewItemVideosItemRepackageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemRepackageRequestBuilder) {
    m := &ItemVideosItemRepackageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library/{libraryId}/videos/{videoId}/repackage?keepOriginalFiles={keepOriginalFiles}", pathParameters),
    }
    return m
}
// NewItemVideosItemRepackageRequestBuilder instantiates a new ItemVideosItemRepackageRequestBuilder and sets the default values.
func NewItemVideosItemRepackageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemRepackageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVideosItemRepackageRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [RepackageVideo API Docs](https://docs.bunny.net/reference/video_repackage)
// returns a Videoable when successful
func (m *ItemVideosItemRepackageRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemVideosItemRepackageRequestBuilderPostQueryParameters])(i5844a9acb9b86d746987515c853816c5c1b43b73d4191fff3bb0559feea76234.Videoable, error) {
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
// ToPostRequestInformation [RepackageVideo API Docs](https://docs.bunny.net/reference/video_repackage)
// returns a *RequestInformation when successful
func (m *ItemVideosItemRepackageRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemVideosItemRepackageRequestBuilderPostQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemVideosItemRepackageRequestBuilder when successful
func (m *ItemVideosItemRepackageRequestBuilder) WithUrl(rawUrl string)(*ItemVideosItemRepackageRequestBuilder) {
    return NewItemVideosItemRepackageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
