package storagezone

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i6e286ba4efe217dd2b9036726264aa0551ed098ea36182223475a366531acd7b "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models/storagezone"
)

// StoragezoneRequestBuilder builds and executes requests for operations under \storagezone
type StoragezoneRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// StoragezoneRequestBuilderGetQueryParameters [ListStorageZones API Docs](https://docs.bunny.net/reference/storagezonepublic_index)
type StoragezoneRequestBuilderGetQueryParameters struct {
    IncludeDeleted *bool `uriparametername:"includeDeleted"`
    Page *int32 `uriparametername:"page"`
    PerPage *int32 `uriparametername:"perPage"`
    Search *string `uriparametername:"search"`
}
// ById gets an item from the github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient.storagezone.item collection
// returns a *StoragezoneItemRequestBuilder when successful
func (m *StoragezoneRequestBuilder) ById(id int64)(*StoragezoneItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(id, 10)
    return NewStoragezoneItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Checkavailability the checkavailability property
// returns a *CheckavailabilityRequestBuilder when successful
func (m *StoragezoneRequestBuilder) Checkavailability()(*CheckavailabilityRequestBuilder) {
    return NewCheckavailabilityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewStoragezoneRequestBuilderInternal instantiates a new StoragezoneRequestBuilder and sets the default values.
func NewStoragezoneRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StoragezoneRequestBuilder) {
    m := &StoragezoneRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storagezone?includeDeleted={includeDeleted}&page={page}&perPage={perPage}{&search*}", pathParameters),
    }
    return m
}
// NewStoragezoneRequestBuilder instantiates a new StoragezoneRequestBuilder and sets the default values.
func NewStoragezoneRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StoragezoneRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStoragezoneRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [ListStorageZones API Docs](https://docs.bunny.net/reference/storagezonepublic_index)
// returns a []StorageZoneable when successful
func (m *StoragezoneRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[StoragezoneRequestBuilderGetQueryParameters])([]i6e286ba4efe217dd2b9036726264aa0551ed098ea36182223475a366531acd7b.StorageZoneable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i6e286ba4efe217dd2b9036726264aa0551ed098ea36182223475a366531acd7b.CreateStorageZoneFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i6e286ba4efe217dd2b9036726264aa0551ed098ea36182223475a366531acd7b.StorageZoneable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i6e286ba4efe217dd2b9036726264aa0551ed098ea36182223475a366531acd7b.StorageZoneable)
        }
    }
    return val, nil
}
// Post [AddStorageZone API Docs](https://docs.bunny.net/reference/storagezonepublic_add)
// returns a StorageZoneable when successful
func (m *StoragezoneRequestBuilder) Post(ctx context.Context, body i6e286ba4efe217dd2b9036726264aa0551ed098ea36182223475a366531acd7b.StorageZoneCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i6e286ba4efe217dd2b9036726264aa0551ed098ea36182223475a366531acd7b.StorageZoneable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ResetReadOnlyPassword the resetReadOnlyPassword property
// returns a *ResetreadonlypasswordResetReadOnlyPasswordRequestBuilder when successful
func (m *StoragezoneRequestBuilder) ResetReadOnlyPassword()(*ResetreadonlypasswordResetReadOnlyPasswordRequestBuilder) {
    return NewResetreadonlypasswordResetReadOnlyPasswordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation [ListStorageZones API Docs](https://docs.bunny.net/reference/storagezonepublic_index)
// returns a *RequestInformation when successful
func (m *StoragezoneRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[StoragezoneRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation [AddStorageZone API Docs](https://docs.bunny.net/reference/storagezonepublic_add)
// returns a *RequestInformation when successful
func (m *StoragezoneRequestBuilder) ToPostRequestInformation(ctx context.Context, body i6e286ba4efe217dd2b9036726264aa0551ed098ea36182223475a366531acd7b.StorageZoneCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/storagezone", m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *StoragezoneRequestBuilder when successful
func (m *StoragezoneRequestBuilder) WithUrl(rawUrl string)(*StoragezoneRequestBuilder) {
    return NewStoragezoneRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
