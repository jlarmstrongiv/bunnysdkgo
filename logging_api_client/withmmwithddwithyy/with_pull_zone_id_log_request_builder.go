package withmmwithddwithyy

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i19c2d0514172c4d19b08b93a7c8434f533d211eeb5aa52a5110e93db0232ec7f "github.com/jlarmstrongiv/bunnysdkgo/logging_api_client/withmmwithddwithyy/withpullzoneidlog"
)

// WithPullZoneIdLogRequestBuilder builds and executes requests for operations under \{mm}-{dd}-{yy}\{pullZoneId}.log
type WithPullZoneIdLogRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithPullZoneIdLogRequestBuilderGetQueryParameters [LoggingApi API Docs](https://docs.bunny.net/docs/cdn-logging-api)
type WithPullZoneIdLogRequestBuilderGetQueryParameters struct {
    End *int64 `uriparametername:"end"`
    Order *i19c2d0514172c4d19b08b93a7c8434f533d211eeb5aa52a5110e93db0232ec7f.GetOrderQueryParameterType `uriparametername:"order"`
    Search *string `uriparametername:"search"`
    Start *int64 `uriparametername:"start"`
    Status *string `uriparametername:"status"`
}
// NewWithPullZoneIdLogRequestBuilderInternal instantiates a new WithPullZoneIdLogRequestBuilder and sets the default values.
func NewWithPullZoneIdLogRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, pullZoneId *int64)(*WithPullZoneIdLogRequestBuilder) {
    m := &WithPullZoneIdLogRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/{mm}-{dd}-{yy}/{pullZoneId}.log?end={end}&order={order}&search={search}&start={start}&status={status}", pathParameters),
    }
    if pullZoneId != nil {
        m.BaseRequestBuilder.PathParameters["pullZoneId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(*pullZoneId, 10)
    }
    return m
}
// NewWithPullZoneIdLogRequestBuilder instantiates a new WithPullZoneIdLogRequestBuilder and sets the default values.
func NewWithPullZoneIdLogRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithPullZoneIdLogRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithPullZoneIdLogRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get [LoggingApi API Docs](https://docs.bunny.net/docs/cdn-logging-api)
// returns a []byte when successful
func (m *WithPullZoneIdLogRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[WithPullZoneIdLogRequestBuilderGetQueryParameters])([]byte, error) {
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
func (m *WithPullZoneIdLogRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[WithPullZoneIdLogRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/gzip")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WithPullZoneIdLogRequestBuilder when successful
func (m *WithPullZoneIdLogRequestBuilder) WithUrl(rawUrl string)(*WithPullZoneIdLogRequestBuilder) {
    return NewWithPullZoneIdLogRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
