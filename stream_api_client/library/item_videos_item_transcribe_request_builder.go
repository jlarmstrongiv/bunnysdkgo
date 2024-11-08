package library

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ic6a2b8141d365338b44f1e735d82ac4081bc99131cbeb3992752f448db804d1d "github.com/jlarmstrongiv/bunnysdkgo/stream_api_client/models"
)

// ItemVideosItemTranscribeRequestBuilder builds and executes requests for operations under \library\{libraryId}\videos\{videoId}\transcribe
type ItemVideosItemTranscribeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemVideosItemTranscribeRequestBuilderPostQueryParameters [TranscribeVideo API Docs](https://docs.bunny.net/reference/video_transcribevideo)
type ItemVideosItemTranscribeRequestBuilderPostQueryParameters struct {
    Force *bool `uriparametername:"force"`
    // Video source language, use ISO 639-1 language code
    Language *string `uriparametername:"language"`
}
// NewItemVideosItemTranscribeRequestBuilderInternal instantiates a new ItemVideosItemTranscribeRequestBuilder and sets the default values.
func NewItemVideosItemTranscribeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemTranscribeRequestBuilder) {
    m := &ItemVideosItemTranscribeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library/{libraryId}/videos/{videoId}/transcribe?force={force}&language={language}", pathParameters),
    }
    return m
}
// NewItemVideosItemTranscribeRequestBuilder instantiates a new ItemVideosItemTranscribeRequestBuilder and sets the default values.
func NewItemVideosItemTranscribeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemTranscribeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVideosItemTranscribeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [TranscribeVideo API Docs](https://docs.bunny.net/reference/video_transcribevideo)
// returns a StructuredSuccessResponseable when successful
func (m *ItemVideosItemTranscribeRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemVideosItemTranscribeRequestBuilderPostQueryParameters])(ic6a2b8141d365338b44f1e735d82ac4081bc99131cbeb3992752f448db804d1d.StructuredSuccessResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation [TranscribeVideo API Docs](https://docs.bunny.net/reference/video_transcribevideo)
// returns a *RequestInformation when successful
func (m *ItemVideosItemTranscribeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemVideosItemTranscribeRequestBuilderPostQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemVideosItemTranscribeRequestBuilder when successful
func (m *ItemVideosItemTranscribeRequestBuilder) WithUrl(rawUrl string)(*ItemVideosItemTranscribeRequestBuilder) {
    return NewItemVideosItemTranscribeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
