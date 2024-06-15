package storagezone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i6e286ba4efe217dd2b9036726264aa0551ed098ea36182223475a366531acd7b "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models/storagezone"
)

// StoragezoneItemRequestBuilder builds and executes requests for operations under \storagezone\{id}
type StoragezoneItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
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
func (m *StoragezoneItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i6e286ba4efe217dd2b9036726264aa0551ed098ea36182223475a366531acd7b.StorageZoneable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i6e286ba4efe217dd2b9036726264aa0551ed098ea36182223475a366531acd7b.CreateStorageZoneFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i6e286ba4efe217dd2b9036726264aa0551ed098ea36182223475a366531acd7b.StorageZoneable), nil
}
// Post [UpdateStorageZone API Docs](https://docs.bunny.net/reference/storagezonepublic_update)
func (m *StoragezoneItemRequestBuilder) Post(ctx context.Context, body i6e286ba4efe217dd2b9036726264aa0551ed098ea36182223475a366531acd7b.StorageZoneCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
func (m *StoragezoneItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body i6e286ba4efe217dd2b9036726264aa0551ed098ea36182223475a366531acd7b.StorageZoneCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
