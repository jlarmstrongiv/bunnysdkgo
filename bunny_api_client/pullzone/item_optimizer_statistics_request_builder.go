package pullzone

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOptimizerStatisticsRequestBuilder builds and executes requests for operations under \pullzone\{-id}\optimizer\statistics
type ItemOptimizerStatisticsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOptimizerStatisticsRequestBuilderGetQueryParameters [GetOptimizerStatistics API Docs](https://docs.bunny.net/reference/pullzonepublic_optimizerstatistics)
type ItemOptimizerStatisticsRequestBuilderGetQueryParameters struct {
    // The start date of the statistics. If no value is passed, the last 30 days will be returned.
    DateFrom *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"dateFrom"`
    // The end date of the statistics. If no value is passed, the last 30 days will be returned.
    DateTo *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"dateTo"`
    // If true, the statistics data will be returned in hourly grouping.
    Hourly *bool `uriparametername:"hourly"`
}
// NewItemOptimizerStatisticsRequestBuilderInternal instantiates a new ItemOptimizerStatisticsRequestBuilder and sets the default values.
func NewItemOptimizerStatisticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOptimizerStatisticsRequestBuilder) {
    m := &ItemOptimizerStatisticsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/optimizer/statistics{?dateFrom*,dateTo*,hourly*}", pathParameters),
    }
    return m
}
// NewItemOptimizerStatisticsRequestBuilder instantiates a new ItemOptimizerStatisticsRequestBuilder and sets the default values.
func NewItemOptimizerStatisticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOptimizerStatisticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOptimizerStatisticsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [GetOptimizerStatistics API Docs](https://docs.bunny.net/reference/pullzonepublic_optimizerstatistics)
// returns a ItemOptimizerStatisticsGetResponseable when successful
func (m *ItemOptimizerStatisticsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemOptimizerStatisticsRequestBuilderGetQueryParameters])(ItemOptimizerStatisticsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOptimizerStatisticsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOptimizerStatisticsGetResponseable), nil
}
// ToGetRequestInformation [GetOptimizerStatistics API Docs](https://docs.bunny.net/reference/pullzonepublic_optimizerstatistics)
// returns a *RequestInformation when successful
func (m *ItemOptimizerStatisticsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemOptimizerStatisticsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemOptimizerStatisticsRequestBuilder when successful
func (m *ItemOptimizerStatisticsRequestBuilder) WithUrl(rawUrl string)(*ItemOptimizerStatisticsRequestBuilder) {
    return NewItemOptimizerStatisticsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
