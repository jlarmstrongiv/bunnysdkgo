package library

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i850d43ce11cdefb3b02ff25c6bb9d6e464cec45fef37ec8775101de5fca8b880 "github.com/jlarmstrongiv/bunnysdkgo/stream_api_client/models/managecollections"
    ia7ab31254ef792422fbaae0bd58f41fbafcfebb55613f9a157dc838b5f75687b "github.com/jlarmstrongiv/bunnysdkgo/stream_api_client/library/item/collections"
)

// ItemCollectionsRequestBuilder builds and executes requests for operations under \library\{libraryId}\collections
type ItemCollectionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsRequestBuilderGetQueryParameters [GetCollectionList API Docs](https://docs.bunny.net/reference/collection_list)
type ItemCollectionsRequestBuilderGetQueryParameters struct {
    IncludeThumbnails *bool `uriparametername:"includeThumbnails"`
    ItemsPerPage *int32 `uriparametername:"itemsPerPage"`
    OrderBy *ia7ab31254ef792422fbaae0bd58f41fbafcfebb55613f9a157dc838b5f75687b.GetOrderByQueryParameterType `uriparametername:"orderBy"`
    Page *int32 `uriparametername:"page"`
    Search *string `uriparametername:"search"`
}
// ByCollectionId gets an item from the github.com/jlarmstrongiv/bunnysdkgo/stream_api_client.library.item.collections.item collection
// returns a *ItemCollectionsWithCollectionItemRequestBuilder when successful
func (m *ItemCollectionsRequestBuilder) ByCollectionId(collectionId string)(*ItemCollectionsWithCollectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if collectionId != "" {
        urlTplParams["collectionId"] = collectionId
    }
    return NewItemCollectionsWithCollectionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCollectionsRequestBuilderInternal instantiates a new ItemCollectionsRequestBuilder and sets the default values.
func NewItemCollectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsRequestBuilder) {
    m := &ItemCollectionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library/{libraryId}/collections?includeThumbnails={includeThumbnails}&itemsPerPage={itemsPerPage}&orderBy={orderBy}&page={page}{&search}", pathParameters),
    }
    return m
}
// NewItemCollectionsRequestBuilder instantiates a new ItemCollectionsRequestBuilder and sets the default values.
func NewItemCollectionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [GetCollectionList API Docs](https://docs.bunny.net/reference/collection_list)
// returns a ItemCollectionsGetResponseable when successful
func (m *ItemCollectionsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemCollectionsRequestBuilderGetQueryParameters])(ItemCollectionsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCollectionsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCollectionsGetResponseable), nil
}
// Post [CreateCollection API Docs](https://docs.bunny.net/reference/collection_createcollection)
// returns a Collectionable when successful
func (m *ItemCollectionsRequestBuilder) Post(ctx context.Context, body i850d43ce11cdefb3b02ff25c6bb9d6e464cec45fef37ec8775101de5fca8b880.Collectionable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i850d43ce11cdefb3b02ff25c6bb9d6e464cec45fef37ec8775101de5fca8b880.Collectionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i850d43ce11cdefb3b02ff25c6bb9d6e464cec45fef37ec8775101de5fca8b880.CreateCollectionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i850d43ce11cdefb3b02ff25c6bb9d6e464cec45fef37ec8775101de5fca8b880.Collectionable), nil
}
// ToGetRequestInformation [GetCollectionList API Docs](https://docs.bunny.net/reference/collection_list)
// returns a *RequestInformation when successful
func (m *ItemCollectionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemCollectionsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation [CreateCollection API Docs](https://docs.bunny.net/reference/collection_createcollection)
// returns a *RequestInformation when successful
func (m *ItemCollectionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i850d43ce11cdefb3b02ff25c6bb9d6e464cec45fef37ec8775101de5fca8b880.Collectionable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/library/{libraryId}/collections", m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCollectionsRequestBuilder when successful
func (m *ItemCollectionsRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsRequestBuilder) {
    return NewItemCollectionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
