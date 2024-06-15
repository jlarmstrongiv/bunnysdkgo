package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ic2c6a3f39062a40014f8015e866e73a07e96b7a67cc729ec34798c1fe3d87586 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models/pullzone/edgerule"
)

// ItemEdgerulesAddorupdateAddOrUpdateRequestBuilder builds and executes requests for operations under \pullzone\{-id}\edgerules\addOrUpdate
type ItemEdgerulesAddorupdateAddOrUpdateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemEdgerulesAddorupdateAddOrUpdateRequestBuilderInternal instantiates a new ItemEdgerulesAddorupdateAddOrUpdateRequestBuilder and sets the default values.
func NewItemEdgerulesAddorupdateAddOrUpdateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEdgerulesAddorupdateAddOrUpdateRequestBuilder) {
    m := &ItemEdgerulesAddorupdateAddOrUpdateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/edgerules/addOrUpdate", pathParameters),
    }
    return m
}
// NewItemEdgerulesAddorupdateAddOrUpdateRequestBuilder instantiates a new ItemEdgerulesAddorupdateAddOrUpdateRequestBuilder and sets the default values.
func NewItemEdgerulesAddorupdateAddOrUpdateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEdgerulesAddorupdateAddOrUpdateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEdgerulesAddorupdateAddOrUpdateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [AddUpdateEdgeRule API Docs](https://docs.bunny.net/reference/pullzonepublic_addedgerule)
// returns a EdgeRuleable when successful
func (m *ItemEdgerulesAddorupdateAddOrUpdateRequestBuilder) Post(ctx context.Context, body ic2c6a3f39062a40014f8015e866e73a07e96b7a67cc729ec34798c1fe3d87586.EdgeRuleable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ic2c6a3f39062a40014f8015e866e73a07e96b7a67cc729ec34798c1fe3d87586.EdgeRuleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic2c6a3f39062a40014f8015e866e73a07e96b7a67cc729ec34798c1fe3d87586.CreateEdgeRuleFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic2c6a3f39062a40014f8015e866e73a07e96b7a67cc729ec34798c1fe3d87586.EdgeRuleable), nil
}
// ToPostRequestInformation [AddUpdateEdgeRule API Docs](https://docs.bunny.net/reference/pullzonepublic_addedgerule)
// returns a *RequestInformation when successful
func (m *ItemEdgerulesAddorupdateAddOrUpdateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ic2c6a3f39062a40014f8015e866e73a07e96b7a67cc729ec34798c1fe3d87586.EdgeRuleable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemEdgerulesAddorupdateAddOrUpdateRequestBuilder when successful
func (m *ItemEdgerulesAddorupdateAddOrUpdateRequestBuilder) WithUrl(rawUrl string)(*ItemEdgerulesAddorupdateAddOrUpdateRequestBuilder) {
    return NewItemEdgerulesAddorupdateAddOrUpdateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
