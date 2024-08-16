package pullzone

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSafehopStatisticsRequestBuilder builds and executes requests for operations under \pullzone\{-id}\safehop\statistics
type ItemSafehopStatisticsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSafehopStatisticsRequestBuilderGetQueryParameters [GetSafeHopStatistics API Docs](https://docs.bunny.net/reference/pullzonepublic_safehopstatistics)
type ItemSafehopStatisticsRequestBuilderGetQueryParameters struct {
    // The start date of the statistics. If no value is passed, the last 30 days will be returned.
    DateFrom *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"dateFrom"`
    // The end date of the statistics. If no value is passed, the last 30 days will be returned.
    DateTo *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"dateTo"`
    // If true, the statistics data will be returned in hourly grouping.
    Hourly *bool `uriparametername:"hourly"`
}
// NewItemSafehopStatisticsRequestBuilderInternal instantiates a new ItemSafehopStatisticsRequestBuilder and sets the default values.
func NewItemSafehopStatisticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSafehopStatisticsRequestBuilder) {
    m := &ItemSafehopStatisticsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/safehop/statistics{?dateFrom,dateTo,hourly}", pathParameters),
    }
    return m
}
// NewItemSafehopStatisticsRequestBuilder instantiates a new ItemSafehopStatisticsRequestBuilder and sets the default values.
func NewItemSafehopStatisticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSafehopStatisticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSafehopStatisticsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [GetSafeHopStatistics API Docs](https://docs.bunny.net/reference/pullzonepublic_safehopstatistics)
// returns a ItemSafehopStatisticsGetResponseable when successful
func (m *ItemSafehopStatisticsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemSafehopStatisticsRequestBuilderGetQueryParameters])(ItemSafehopStatisticsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSafehopStatisticsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSafehopStatisticsGetResponseable), nil
}
// ToGetRequestInformation [GetSafeHopStatistics API Docs](https://docs.bunny.net/reference/pullzonepublic_safehopstatistics)
// returns a *RequestInformation when successful
func (m *ItemSafehopStatisticsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemSafehopStatisticsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSafehopStatisticsRequestBuilder when successful
func (m *ItemSafehopStatisticsRequestBuilder) WithUrl(rawUrl string)(*ItemSafehopStatisticsRequestBuilder) {
    return NewItemSafehopStatisticsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
