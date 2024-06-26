package dnszone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib14791b6c35b6ec4957489e332a1285cf193eb3d7db2b1fd1f20fa697ab6fcae "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models"
    i8affc51df26196299d4adfa5b2f725f33e0381e875a98cdcece21a52073b3ffd "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/dnszone/updatednsrecord"
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
        "400": ib14791b6c35b6ec4957489e332a1285cf193eb3d7db2b1fd1f20fa697ab6fcae.CreateStructuredBadRequestResponseFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Post [UpdateDnsRecord API Docs](https://docs.bunny.net/reference/dnszonepublic_updaterecord)
// returns a StructuredBadRequestResponse error when the service returns a 400 status code
func (m *ItemRecordsRecordsItemRequestBuilder) Post(ctx context.Context, body i8affc51df26196299d4adfa5b2f725f33e0381e875a98cdcece21a52073b3ffd.OptionalParametersable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": ib14791b6c35b6ec4957489e332a1285cf193eb3d7db2b1fd1f20fa697ab6fcae.CreateStructuredBadRequestResponseFromDiscriminatorValue,
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
func (m *ItemRecordsRecordsItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body i8affc51df26196299d4adfa5b2f725f33e0381e875a98cdcece21a52073b3ffd.OptionalParametersable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
