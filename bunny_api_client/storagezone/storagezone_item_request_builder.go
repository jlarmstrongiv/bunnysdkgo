package storagezone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i35697601df5820fe555c46235c796ce9ca5e9d758cc28eaf0fc2ed84d4a3aaa1 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/storagezone"
)

// StoragezoneItemRequestBuilder builds and executes requests for operations under \storagezone\{id}
type StoragezoneItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Connections the connections property
// returns a *ItemConnectionsRequestBuilder when successful
func (m *StoragezoneItemRequestBuilder) Connections()(*ItemConnectionsRequestBuilder) {
    return NewItemConnectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewStoragezoneItemRequestBuilderInternal instantiates a new StoragezoneItemRequestBuilder and sets the default values.
func NewStoragezoneItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StoragezoneItemRequestBuilder) {
    m := &StoragezoneItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storagezone/{id}", pathParameters),
    }
    return m
}
// NewStoragezoneItemRequestBuilder instantiates a new StoragezoneItemRequestBuilder and sets the default values.
func NewStoragezoneItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StoragezoneItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStoragezoneItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *StoragezoneItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get [GetStorageZone API Docs](https://docs.bunny.net/reference/storagezonepublic_index2)
// returns a StorageZoneable when successful
func (m *StoragezoneItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i35697601df5820fe555c46235c796ce9ca5e9d758cc28eaf0fc2ed84d4a3aaa1.StorageZoneable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i35697601df5820fe555c46235c796ce9ca5e9d758cc28eaf0fc2ed84d4a3aaa1.CreateStorageZoneFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i35697601df5820fe555c46235c796ce9ca5e9d758cc28eaf0fc2ed84d4a3aaa1.StorageZoneable), nil
}
// Post [UpdateStorageZone API Docs](https://docs.bunny.net/reference/storagezonepublic_update)
func (m *StoragezoneItemRequestBuilder) Post(ctx context.Context, body i35697601df5820fe555c46235c796ce9ca5e9d758cc28eaf0fc2ed84d4a3aaa1.StorageZoneCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ResetPassword the resetPassword property
// returns a *ItemResetpasswordResetPasswordRequestBuilder when successful
func (m *StoragezoneItemRequestBuilder) ResetPassword()(*ItemResetpasswordResetPasswordRequestBuilder) {
    return NewItemResetpasswordResetPasswordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Statistics the statistics property
// returns a *ItemStatisticsRequestBuilder when successful
func (m *StoragezoneItemRequestBuilder) Statistics()(*ItemStatisticsRequestBuilder) {
    return NewItemStatisticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// returns a *RequestInformation when successful
func (m *StoragezoneItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// ToGetRequestInformation [GetStorageZone API Docs](https://docs.bunny.net/reference/storagezonepublic_index2)
// returns a *RequestInformation when successful
func (m *StoragezoneItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation [UpdateStorageZone API Docs](https://docs.bunny.net/reference/storagezonepublic_update)
// returns a *RequestInformation when successful
func (m *StoragezoneItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body i35697601df5820fe555c46235c796ce9ca5e9d758cc28eaf0fc2ed84d4a3aaa1.StorageZoneCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *StoragezoneItemRequestBuilder when successful
func (m *StoragezoneItemRequestBuilder) WithUrl(rawUrl string)(*StoragezoneItemRequestBuilder) {
    return NewStoragezoneItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
