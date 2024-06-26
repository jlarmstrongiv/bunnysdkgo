package storagezone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    id1540a1b5c14e26b8bd31bf71bcba2ceb0b9cb713071261357f35dc162cdfb13 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/storagezone/getstoragezoneconnections"
)

// ItemConnectionsRequestBuilder builds and executes requests for operations under \storagezone\{id}\connections
type ItemConnectionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemConnectionsRequestBuilderInternal instantiates a new ItemConnectionsRequestBuilder and sets the default values.
func NewItemConnectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectionsRequestBuilder) {
    m := &ItemConnectionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storagezone/{id}/connections", pathParameters),
    }
    return m
}
// NewItemConnectionsRequestBuilder instantiates a new ItemConnectionsRequestBuilder and sets the default values.
func NewItemConnectionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConnectionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [GetStorageZoneConnections API Docs](https://docs.bunny.net/reference/storagezonepublic_connections)
// returns a []Connectionable when successful
func (m *ItemConnectionsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]id1540a1b5c14e26b8bd31bf71bcba2ceb0b9cb713071261357f35dc162cdfb13.Connectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, id1540a1b5c14e26b8bd31bf71bcba2ceb0b9cb713071261357f35dc162cdfb13.CreateConnectionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]id1540a1b5c14e26b8bd31bf71bcba2ceb0b9cb713071261357f35dc162cdfb13.Connectionable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(id1540a1b5c14e26b8bd31bf71bcba2ceb0b9cb713071261357f35dc162cdfb13.Connectionable)
        }
    }
    return val, nil
}
// ToGetRequestInformation [GetStorageZoneConnections API Docs](https://docs.bunny.net/reference/storagezonepublic_connections)
// returns a *RequestInformation when successful
func (m *ItemConnectionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemConnectionsRequestBuilder when successful
func (m *ItemConnectionsRequestBuilder) WithUrl(rawUrl string)(*ItemConnectionsRequestBuilder) {
    return NewItemConnectionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
