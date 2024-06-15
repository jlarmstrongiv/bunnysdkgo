package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i23546f19c5d6462aadb661ab85e60299bfae70903a2193c41d88ff0400676c0c "github.com/jlarmstrongiv/bunnysdkgo/generated/EdgeStorageApiClient/models"
)

// WithPathItemRequestBuilder builds and executes requests for operations under \{storageZoneName}\{path}
type WithPathItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByFileName gets an item from the github.com/jlarmstrongiv/bunnysdkgo/generated/EdgeStorageApiClient.item.item.item collection
// returns a *ItemWithFileNameItemRequestBuilder when successful
func (m *WithPathItemRequestBuilder) ByFileName(fileName string)(*ItemWithFileNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if fileName != "" {
        urlTplParams["fileName"] = fileName
    }
    return NewItemWithFileNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithPathItemRequestBuilderInternal instantiates a new WithPathItemRequestBuilder and sets the default values.
func NewWithPathItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithPathItemRequestBuilder) {
    m := &WithPathItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/{storageZoneName}/{path}", pathParameters),
    }
    return m
}
// NewWithPathItemRequestBuilder instantiates a new WithPathItemRequestBuilder and sets the default values.
func NewWithPathItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithPathItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithPathItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [ListFiles API Docs](https://docs.bunny.net/reference/get_-storagezonename-path-)
// returns a []Fileable when successful
func (m *WithPathItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i23546f19c5d6462aadb661ab85e60299bfae70903a2193c41d88ff0400676c0c.Fileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i23546f19c5d6462aadb661ab85e60299bfae70903a2193c41d88ff0400676c0c.CreateFileFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i23546f19c5d6462aadb661ab85e60299bfae70903a2193c41d88ff0400676c0c.Fileable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i23546f19c5d6462aadb661ab85e60299bfae70903a2193c41d88ff0400676c0c.Fileable)
        }
    }
    return val, nil
}
// ToGetRequestInformation [ListFiles API Docs](https://docs.bunny.net/reference/get_-storagezonename-path-)
// returns a *RequestInformation when successful
func (m *WithPathItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WithPathItemRequestBuilder when successful
func (m *WithPathItemRequestBuilder) WithUrl(rawUrl string)(*WithPathItemRequestBuilder) {
    return NewWithPathItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
