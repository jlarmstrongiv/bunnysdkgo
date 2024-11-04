package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i97f7eed2653035383217e2ff197d6949f30724c49cf23d1b582eff8dac419d81 "github.com/jlarmstrongiv/bunnysdkgo/edge_storage_api_client/models"
)

// ItemWithPathSlashRequestBuilder builds and executes requests for operations under \{storageZoneName}\{+path}\
type ItemWithPathSlashRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemWithPathSlashRequestBuilderInternal instantiates a new ItemWithPathSlashRequestBuilder and sets the default values.
func NewItemWithPathSlashRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, path *string)(*ItemWithPathSlashRequestBuilder) {
    m := &ItemWithPathSlashRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/{storageZoneName}/{+path}/", pathParameters),
    }
    if path != nil {
        m.BaseRequestBuilder.PathParameters["path"] = *path
    }
    return m
}
// NewItemWithPathSlashRequestBuilder instantiates a new ItemWithPathSlashRequestBuilder and sets the default values.
func NewItemWithPathSlashRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWithPathSlashRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemWithPathSlashRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get [ListFiles API Docs](https://docs.bunny.net/reference/get_-storagezonename-path-)
// returns a []Fileable when successful
func (m *ItemWithPathSlashRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i97f7eed2653035383217e2ff197d6949f30724c49cf23d1b582eff8dac419d81.Fileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i97f7eed2653035383217e2ff197d6949f30724c49cf23d1b582eff8dac419d81.CreateFileFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i97f7eed2653035383217e2ff197d6949f30724c49cf23d1b582eff8dac419d81.Fileable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i97f7eed2653035383217e2ff197d6949f30724c49cf23d1b582eff8dac419d81.Fileable)
        }
    }
    return val, nil
}
// ToGetRequestInformation [ListFiles API Docs](https://docs.bunny.net/reference/get_-storagezonename-path-)
// returns a *RequestInformation when successful
func (m *ItemWithPathSlashRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemWithPathSlashRequestBuilder when successful
func (m *ItemWithPathSlashRequestBuilder) WithUrl(rawUrl string)(*ItemWithPathSlashRequestBuilder) {
    return NewItemWithPathSlashRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
