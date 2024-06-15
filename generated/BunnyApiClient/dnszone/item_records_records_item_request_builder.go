package dnszone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iaddb6f6101023ea6e821fd917ca8105f1326bd51411c26d4851be34a87cbb944 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models"
    i6224d1b932828b2b917266486a1fe8bbcb1219f9049ed46c28da622d1dd60fce "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models/dnszone/updatednsrecord"
)

// ItemRecordsRecordsItemRequestBuilder builds and executes requests for operations under \dnszone\{-id}\records\{id}
type ItemRecordsRecordsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemRecordsRecordsItemRequestBuilderInternal instantiates a new ItemRecordsRecordsItemRequestBuilder and sets the default values.
func NewItemRecordsRecordsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRecordsRecordsItemRequestBuilder) {
    m := &ItemRecordsRecordsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dnszone/{%2Did}/records/{id}", pathParameters),
    }
    return m
}
// NewItemRecordsRecordsItemRequestBuilder instantiates a new ItemRecordsRecordsItemRequestBuilder and sets the default values.
func NewItemRecordsRecordsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRecordsRecordsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRecordsRecordsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete [DeleteDnsRecord API Docs](https://docs.bunny.net/reference/dnszonepublic_deleterecord)
// returns a StructuredBadRequestResponse error when the service returns a 400 status code
func (m *ItemRecordsRecordsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": iaddb6f6101023ea6e821fd917ca8105f1326bd51411c26d4851be34a87cbb944.CreateStructuredBadRequestResponseFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Post [UpdateDnsRecord API Docs](https://docs.bunny.net/reference/dnszonepublic_updaterecord)
// returns a StructuredBadRequestResponse error when the service returns a 400 status code
func (m *ItemRecordsRecordsItemRequestBuilder) Post(ctx context.Context, body i6224d1b932828b2b917266486a1fe8bbcb1219f9049ed46c28da622d1dd60fce.OptionalParametersable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": iaddb6f6101023ea6e821fd917ca8105f1326bd51411c26d4851be34a87cbb944.CreateStructuredBadRequestResponseFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation [DeleteDnsRecord API Docs](https://docs.bunny.net/reference/dnszonepublic_deleterecord)
// returns a *RequestInformation when successful
func (m *ItemRecordsRecordsItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation [UpdateDnsRecord API Docs](https://docs.bunny.net/reference/dnszonepublic_updaterecord)
// returns a *RequestInformation when successful
func (m *ItemRecordsRecordsItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body i6224d1b932828b2b917266486a1fe8bbcb1219f9049ed46c28da622d1dd60fce.OptionalParametersable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRecordsRecordsItemRequestBuilder when successful
func (m *ItemRecordsRecordsItemRequestBuilder) WithUrl(rawUrl string)(*ItemRecordsRecordsItemRequestBuilder) {
    return NewItemRecordsRecordsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
