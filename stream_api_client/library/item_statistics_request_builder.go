package library

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemStatisticsRequestBuilder builds and executes requests for operations under \library\{libraryId}\statistics
type ItemStatisticsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemStatisticsRequestBuilderGetQueryParameters [GetVideoStatistics API Docs](https://docs.bunny.net/reference/video_getvideostatistics)
type ItemStatisticsRequestBuilderGetQueryParameters struct {
    // The start date of the statistics. If no value is passed, the last 30 days will be returned.
    DateFrom *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"dateFrom"`
    // The end date of the statistics. If no value is passed, the last 30 days will be returned.
    DateTo *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"dateTo"`
    // If true, the statistics data will be returned in hourly grouping.
    Hourly *bool `uriparametername:"hourly"`
    // The GUID of the video for which the statistics will be returned
    VideoGuid *string `uriparametername:"videoGuid"`
}
// NewItemStatisticsRequestBuilderInternal instantiates a new ItemStatisticsRequestBuilder and sets the default values.
func NewItemStatisticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemStatisticsRequestBuilder) {
    m := &ItemStatisticsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library/{libraryId}/statistics{?dateFrom*,dateTo*,hourly*,videoGuid*}", pathParameters),
    }
    return m
}
// NewItemStatisticsRequestBuilder instantiates a new ItemStatisticsRequestBuilder and sets the default values.
func NewItemStatisticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemStatisticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemStatisticsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [GetVideoStatistics API Docs](https://docs.bunny.net/reference/video_getvideostatistics)
// returns a ItemStatisticsGetResponseable when successful
func (m *ItemStatisticsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemStatisticsRequestBuilderGetQueryParameters])(ItemStatisticsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemStatisticsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemStatisticsGetResponseable), nil
}
// ToGetRequestInformation [GetVideoStatistics API Docs](https://docs.bunny.net/reference/video_getvideostatistics)
// returns a *RequestInformation when successful
func (m *ItemStatisticsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemStatisticsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemStatisticsRequestBuilder when successful
func (m *ItemStatisticsRequestBuilder) WithUrl(rawUrl string)(*ItemStatisticsRequestBuilder) {
    return NewItemStatisticsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
