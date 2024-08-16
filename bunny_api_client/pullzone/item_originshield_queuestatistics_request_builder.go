package pullzone

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOriginshieldQueuestatisticsRequestBuilder builds and executes requests for operations under \pullzone\{-id}\originshield\queuestatistics
type ItemOriginshieldQueuestatisticsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOriginshieldQueuestatisticsRequestBuilderGetQueryParameters [GetOriginShieldQueueStatistics API Docs](https://docs.bunny.net/reference/pullzonepublic_originshieldconcurrencystatistics)
type ItemOriginshieldQueuestatisticsRequestBuilderGetQueryParameters struct {
    // The start date of the statistics. If no value is passed, the last 30 days will be returned.
    DateFrom *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"dateFrom"`
    // The end date of the statistics. If no value is passed, the last 30 days will be returned.
    DateTo *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"dateTo"`
    // If true, the statistics data will be returned in hourly grouping.
    Hourly *bool `uriparametername:"hourly"`
}
// NewItemOriginshieldQueuestatisticsRequestBuilderInternal instantiates a new ItemOriginshieldQueuestatisticsRequestBuilder and sets the default values.
func NewItemOriginshieldQueuestatisticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOriginshieldQueuestatisticsRequestBuilder) {
    m := &ItemOriginshieldQueuestatisticsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/originshield/queuestatistics{?dateFrom,dateTo,hourly}", pathParameters),
    }
    return m
}
// NewItemOriginshieldQueuestatisticsRequestBuilder instantiates a new ItemOriginshieldQueuestatisticsRequestBuilder and sets the default values.
func NewItemOriginshieldQueuestatisticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOriginshieldQueuestatisticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOriginshieldQueuestatisticsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [GetOriginShieldQueueStatistics API Docs](https://docs.bunny.net/reference/pullzonepublic_originshieldconcurrencystatistics)
// returns a ItemOriginshieldQueuestatisticsGetResponseable when successful
func (m *ItemOriginshieldQueuestatisticsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemOriginshieldQueuestatisticsRequestBuilderGetQueryParameters])(ItemOriginshieldQueuestatisticsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOriginshieldQueuestatisticsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOriginshieldQueuestatisticsGetResponseable), nil
}
// ToGetRequestInformation [GetOriginShieldQueueStatistics API Docs](https://docs.bunny.net/reference/pullzonepublic_originshieldconcurrencystatistics)
// returns a *RequestInformation when successful
func (m *ItemOriginshieldQueuestatisticsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemOriginshieldQueuestatisticsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemOriginshieldQueuestatisticsRequestBuilder when successful
func (m *ItemOriginshieldQueuestatisticsRequestBuilder) WithUrl(rawUrl string)(*ItemOriginshieldQueuestatisticsRequestBuilder) {
    return NewItemOriginshieldQueuestatisticsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
