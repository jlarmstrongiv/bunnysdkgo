package pullzone

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemWafStatisticsRequestBuilder builds and executes requests for operations under \pullzone\{-id}\waf\statistics
type ItemWafStatisticsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
type ItemWafStatisticsRequestBuilderGetQueryParameters struct {
    // The start date of the statistics. If no value is passed, the last 30 days will be returned.
    DateFrom *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"dateFrom"`
    // The end date of the statistics. If no value is passed, the last 30 days will be returned.
    DateTo *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"dateTo"`
    // If true, the statistics data will be returned in hourly grouping.
    Hourly *bool `uriparametername:"hourly"`
}
// NewItemWafStatisticsRequestBuilderInternal instantiates a new ItemWafStatisticsRequestBuilder and sets the default values.
func NewItemWafStatisticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWafStatisticsRequestBuilder) {
    m := &ItemWafStatisticsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/waf/statistics{?dateFrom*,dateTo*,hourly*}", pathParameters),
    }
    return m
}
// NewItemWafStatisticsRequestBuilder instantiates a new ItemWafStatisticsRequestBuilder and sets the default values.
func NewItemWafStatisticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWafStatisticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemWafStatisticsRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a ItemWafStatisticsGetResponseable when successful
func (m *ItemWafStatisticsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemWafStatisticsRequestBuilderGetQueryParameters])(ItemWafStatisticsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemWafStatisticsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemWafStatisticsGetResponseable), nil
}
// returns a *RequestInformation when successful
func (m *ItemWafStatisticsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemWafStatisticsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemWafStatisticsRequestBuilder when successful
func (m *ItemWafStatisticsRequestBuilder) WithUrl(rawUrl string)(*ItemWafStatisticsRequestBuilder) {
    return NewItemWafStatisticsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
