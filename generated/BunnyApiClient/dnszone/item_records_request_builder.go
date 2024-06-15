package dnszone

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iaddb6f6101023ea6e821fd917ca8105f1326bd51411c26d4851be34a87cbb944 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models"
    iad9fd93a01c50a131e0880d2d5343c7a8209f886cba9d78dcb28342fae5a5d5c "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models/dnszone"
)

// ItemRecordsRequestBuilder builds and executes requests for operations under \dnszone\{-id}\records
type ItemRecordsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ById gets an item from the github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient.dnszone.item.records.item collection
// returns a *ItemRecordsRecordsItemRequestBuilder when successful
func (m *ItemRecordsRequestBuilder) ById(id int64)(*ItemRecordsRecordsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(id, 10)
    return NewItemRecordsRecordsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemRecordsRequestBuilderInternal instantiates a new ItemRecordsRequestBuilder and sets the default values.
func NewItemRecordsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRecordsRequestBuilder) {
    m := &ItemRecordsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dnszone/{%2Did}/records", pathParameters),
    }
    return m
}
// NewItemRecordsRequestBuilder instantiates a new ItemRecordsRequestBuilder and sets the default values.
func NewItemRecordsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRecordsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRecordsRequestBuilderInternal(urlParams, requestAdapter)
}
// Put [AddDnsRecord API Docs](https://docs.bunny.net/reference/dnszonepublic_addrecord)
// returns a DnsRecordable when successful
// returns a StructuredBadRequestResponse error when the service returns a 400 status code
func (m *ItemRecordsRequestBuilder) Put(ctx context.Context, body ItemRecordsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(iad9fd93a01c50a131e0880d2d5343c7a8209f886cba9d78dcb28342fae5a5d5c.DnsRecordable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": iaddb6f6101023ea6e821fd917ca8105f1326bd51411c26d4851be34a87cbb944.CreateStructuredBadRequestResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iad9fd93a01c50a131e0880d2d5343c7a8209f886cba9d78dcb28342fae5a5d5c.CreateDnsRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iad9fd93a01c50a131e0880d2d5343c7a8209f886cba9d78dcb28342fae5a5d5c.DnsRecordable), nil
}
// ToPutRequestInformation [AddDnsRecord API Docs](https://docs.bunny.net/reference/dnszonepublic_addrecord)
// returns a *RequestInformation when successful
func (m *ItemRecordsRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemRecordsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRecordsRequestBuilder when successful
func (m *ItemRecordsRequestBuilder) WithUrl(rawUrl string)(*ItemRecordsRequestBuilder) {
    return NewItemRecordsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
