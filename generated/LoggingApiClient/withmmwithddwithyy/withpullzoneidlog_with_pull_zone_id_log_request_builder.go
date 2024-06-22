package withmmwithddwithyy

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ifbc02c50c687852e2bbdb482dcb6c5411c62789b41331b62281d5577ff367f01 "github.com/jlarmstrongiv/bunnysdkgo/generated/LoggingApiClient/withmmwithddwithyy/withpullzoneidlog"
)

// WithpullzoneidlogWithPullZoneIdLogRequestBuilder builds and executes requests for operations under \{mm}-{dd}-{yy}\{pullZoneId}.log
type WithpullzoneidlogWithPullZoneIdLogRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithpullzoneidlogWithPullZoneIdLogRequestBuilderGetQueryParameters [LoggingApi API Docs](https://docs.bunny.net/docs/cdn-logging-api)
type WithpullzoneidlogWithPullZoneIdLogRequestBuilderGetQueryParameters struct {
    End *int64 `uriparametername:"end"`
    Order *ifbc02c50c687852e2bbdb482dcb6c5411c62789b41331b62281d5577ff367f01.GetOrderQueryParameterType `uriparametername:"order"`
    Search *string `uriparametername:"search"`
    Start *int64 `uriparametername:"start"`
    Status *string `uriparametername:"status"`
}
// NewWithpullzoneidlogWithPullZoneIdLogRequestBuilderInternal instantiates a new WithpullzoneidlogWithPullZoneIdLogRequestBuilder and sets the default values.
func NewWithpullzoneidlogWithPullZoneIdLogRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, pullZoneId *int64)(*WithpullzoneidlogWithPullZoneIdLogRequestBuilder) {
    m := &WithpullzoneidlogWithPullZoneIdLogRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/{mm}-{dd}-{yy}/{pullZoneId}.log?end={end}&order={order}&search={search}&start={start}&status={status}", pathParameters),
    }
    if pullZoneId != nil {
        m.BaseRequestBuilder.PathParameters["pullZoneId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(*pullZoneId, 10)
    }
    return m
}
// NewWithpullzoneidlogWithPullZoneIdLogRequestBuilder instantiates a new WithpullzoneidlogWithPullZoneIdLogRequestBuilder and sets the default values.
func NewWithpullzoneidlogWithPullZoneIdLogRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithpullzoneidlogWithPullZoneIdLogRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithpullzoneidlogWithPullZoneIdLogRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get [LoggingApi API Docs](https://docs.bunny.net/docs/cdn-logging-api)
// returns a []byte when successful
func (m *WithpullzoneidlogWithPullZoneIdLogRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[WithpullzoneidlogWithPullZoneIdLogRequestBuilderGetQueryParameters])([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation [LoggingApi API Docs](https://docs.bunny.net/docs/cdn-logging-api)
// returns a *RequestInformation when successful
func (m *WithpullzoneidlogWithPullZoneIdLogRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[WithpullzoneidlogWithPullZoneIdLogRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/gzip")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WithpullzoneidlogWithPullZoneIdLogRequestBuilder when successful
func (m *WithpullzoneidlogWithPullZoneIdLogRequestBuilder) WithUrl(rawUrl string)(*WithpullzoneidlogWithPullZoneIdLogRequestBuilder) {
    return NewWithpullzoneidlogWithPullZoneIdLogRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
