package library

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0 "github.com/jlarmstrongiv/bunnysdkgo/generated/StreamApiClient/models"
    i3dff0a221e8bacaf9c54f97693b09c700857cde1d99499662fda6479c5b6beaa "github.com/jlarmstrongiv/bunnysdkgo/generated/StreamApiClient/models/managecollections"
)

// ItemCollectionsWithCollectionItemRequestBuilder builds and executes requests for operations under \library\{libraryId}\collections\{collectionId}
type ItemCollectionsWithCollectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsWithCollectionItemRequestBuilderGetQueryParameters [GetCollection API Docs](https://docs.bunny.net/reference/collection_getcollection)
type ItemCollectionsWithCollectionItemRequestBuilderGetQueryParameters struct {
    IncludeThumbnails *bool `uriparametername:"includeThumbnails"`
}
// NewItemCollectionsWithCollectionItemRequestBuilderInternal instantiates a new ItemCollectionsWithCollectionItemRequestBuilder and sets the default values.
func NewItemCollectionsWithCollectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsWithCollectionItemRequestBuilder) {
    m := &ItemCollectionsWithCollectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library/{libraryId}/collections/{collectionId}?includeThumbnails={includeThumbnails}", pathParameters),
    }
    return m
}
// NewItemCollectionsWithCollectionItemRequestBuilder instantiates a new ItemCollectionsWithCollectionItemRequestBuilder and sets the default values.
func NewItemCollectionsWithCollectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsWithCollectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsWithCollectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete [DeleteCollection API Docs](https://docs.bunny.net/reference/collection_deletecollection)
// returns a StructuredSuccessResponseable when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0.StructuredSuccessResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0.CreateStructuredSuccessResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0.StructuredSuccessResponseable), nil
}
// Get [GetCollection API Docs](https://docs.bunny.net/reference/collection_getcollection)
// returns a Collectionable when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemCollectionsWithCollectionItemRequestBuilderGetQueryParameters])(i3dff0a221e8bacaf9c54f97693b09c700857cde1d99499662fda6479c5b6beaa.Collectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i3dff0a221e8bacaf9c54f97693b09c700857cde1d99499662fda6479c5b6beaa.CreateCollectionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i3dff0a221e8bacaf9c54f97693b09c700857cde1d99499662fda6479c5b6beaa.Collectionable), nil
}
// Post [UpdateCollection API Docs](https://docs.bunny.net/reference/collection_updatecollection)
// returns a StructuredSuccessResponseable when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) Post(ctx context.Context, body i3dff0a221e8bacaf9c54f97693b09c700857cde1d99499662fda6479c5b6beaa.Collectionable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0.StructuredSuccessResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0.CreateStructuredSuccessResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0.StructuredSuccessResponseable), nil
}
// ToDeleteRequestInformation [DeleteCollection API Docs](https://docs.bunny.net/reference/collection_deletecollection)
// returns a *RequestInformation when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/library/{libraryId}/collections/{collectionId}", m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation [GetCollection API Docs](https://docs.bunny.net/reference/collection_getcollection)
// returns a *RequestInformation when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemCollectionsWithCollectionItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation [UpdateCollection API Docs](https://docs.bunny.net/reference/collection_updatecollection)
// returns a *RequestInformation when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body i3dff0a221e8bacaf9c54f97693b09c700857cde1d99499662fda6479c5b6beaa.Collectionable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/library/{libraryId}/collections/{collectionId}", m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCollectionsWithCollectionItemRequestBuilder when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsWithCollectionItemRequestBuilder) {
    return NewItemCollectionsWithCollectionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
