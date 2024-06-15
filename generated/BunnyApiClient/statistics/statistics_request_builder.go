package statistics

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ibc615e5a22b00b81a37f4c68e6b77616779f7dfea48f81c56396a82d7119732e "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models/statistics/getstatistics"
)

// StatisticsRequestBuilder builds and executes requests for operations under \statistics
type StatisticsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// StatisticsRequestBuilderGetQueryParameters [GetStatistics API Docs](https://docs.bunny.net/reference/statisticspublic_index)
type StatisticsRequestBuilderGetQueryParameters struct {
    DateFrom *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"dateFrom"`
    DateTo *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"dateTo"`
    Hourly *bool `uriparametername:"hourly"`
    LoadErrors *bool `uriparametername:"loadErrors"`
    PullZone *int64 `uriparametername:"pullZone"`
    ServerZoneId *int64 `uriparametername:"serverZoneId"`
}
// NewStatisticsRequestBuilderInternal instantiates a new StatisticsRequestBuilder and sets the default values.
func NewStatisticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StatisticsRequestBuilder) {
    m := &StatisticsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/statistics{?dateFrom*,dateTo*,hourly*,loadErrors*,pullZone*,serverZoneId*}", pathParameters),
    }
    return m
}
// NewStatisticsRequestBuilder instantiates a new StatisticsRequestBuilder and sets the default values.
func NewStatisticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StatisticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStatisticsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [GetStatistics API Docs](https://docs.bunny.net/reference/statisticspublic_index)
// returns a Statisticsable when successful
func (m *StatisticsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[StatisticsRequestBuilderGetQueryParameters])(ibc615e5a22b00b81a37f4c68e6b77616779f7dfea48f81c56396a82d7119732e.Statisticsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibc615e5a22b00b81a37f4c68e6b77616779f7dfea48f81c56396a82d7119732e.CreateStatisticsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibc615e5a22b00b81a37f4c68e6b77616779f7dfea48f81c56396a82d7119732e.Statisticsable), nil
}
// ToGetRequestInformation [GetStatistics API Docs](https://docs.bunny.net/reference/statisticspublic_index)
// returns a *RequestInformation when successful
func (m *StatisticsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[StatisticsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *StatisticsRequestBuilder when successful
func (m *StatisticsRequestBuilder) WithUrl(rawUrl string)(*StatisticsRequestBuilder) {
    return NewStatisticsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
