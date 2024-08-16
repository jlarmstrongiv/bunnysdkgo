package library

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ic6a2b8141d365338b44f1e735d82ac4081bc99131cbeb3992752f448db804d1d "github.com/jlarmstrongiv/bunnysdkgo/stream_api_client/models"
)

// ItemVideosFetchRequestBuilder builds and executes requests for operations under \library\{libraryId}\videos\fetch
type ItemVideosFetchRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemVideosFetchRequestBuilderPostQueryParameters [FetchVideo API Docs](https://docs.bunny.net/reference/video_fetchnewvideo)
type ItemVideosFetchRequestBuilderPostQueryParameters struct {
    CollectionId *string `uriparametername:"collectionId"`
    // Video time in ms to extract the main video thumbnail.
    ThumbnailTime *int32 `uriparametername:"thumbnailTime"`
}
// NewItemVideosFetchRequestBuilderInternal instantiates a new ItemVideosFetchRequestBuilder and sets the default values.
func NewItemVideosFetchRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosFetchRequestBuilder) {
    m := &ItemVideosFetchRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library/{libraryId}/videos/fetch{?collectionId,thumbnailTime}", pathParameters),
    }
    return m
}
// NewItemVideosFetchRequestBuilder instantiates a new ItemVideosFetchRequestBuilder and sets the default values.
func NewItemVideosFetchRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosFetchRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVideosFetchRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [FetchVideo API Docs](https://docs.bunny.net/reference/video_fetchnewvideo)
// returns a StructuredSuccessResponseable when successful
func (m *ItemVideosFetchRequestBuilder) Post(ctx context.Context, body ItemVideosFetchPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemVideosFetchRequestBuilderPostQueryParameters])(ic6a2b8141d365338b44f1e735d82ac4081bc99131cbeb3992752f448db804d1d.StructuredSuccessResponseable, error) {
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
// ToPostRequestInformation [FetchVideo API Docs](https://docs.bunny.net/reference/video_fetchnewvideo)
// returns a *RequestInformation when successful
func (m *ItemVideosFetchRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemVideosFetchPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemVideosFetchRequestBuilderPostQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemVideosFetchRequestBuilder when successful
func (m *ItemVideosFetchRequestBuilder) WithUrl(rawUrl string)(*ItemVideosFetchRequestBuilder) {
    return NewItemVideosFetchRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
