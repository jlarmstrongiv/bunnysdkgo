package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ibdd3cc897f8f3833299954aef4508f6bfff545c7346b36641af0214c751e5b85 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/pullzone/edgerule"
)

// ItemEdgerulesAddOrUpdateRequestBuilder builds and executes requests for operations under \pullzone\{-id}\edgerules\addOrUpdate
type ItemEdgerulesAddOrUpdateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemEdgerulesAddOrUpdateRequestBuilderInternal instantiates a new ItemEdgerulesAddOrUpdateRequestBuilder and sets the default values.
func NewItemEdgerulesAddOrUpdateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEdgerulesAddOrUpdateRequestBuilder) {
    m := &ItemEdgerulesAddOrUpdateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/edgerules/addOrUpdate", pathParameters),
    }
    return m
}
// NewItemEdgerulesAddOrUpdateRequestBuilder instantiates a new ItemEdgerulesAddOrUpdateRequestBuilder and sets the default values.
func NewItemEdgerulesAddOrUpdateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEdgerulesAddOrUpdateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEdgerulesAddOrUpdateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [AddUpdateEdgeRule API Docs](https://docs.bunny.net/reference/pullzonepublic_addedgerule)
// returns a EdgeRuleable when successful
func (m *ItemEdgerulesAddOrUpdateRequestBuilder) Post(ctx context.Context, body ibdd3cc897f8f3833299954aef4508f6bfff545c7346b36641af0214c751e5b85.EdgeRuleable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ibdd3cc897f8f3833299954aef4508f6bfff545c7346b36641af0214c751e5b85.EdgeRuleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibdd3cc897f8f3833299954aef4508f6bfff545c7346b36641af0214c751e5b85.CreateEdgeRuleFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibdd3cc897f8f3833299954aef4508f6bfff545c7346b36641af0214c751e5b85.EdgeRuleable), nil
}
// ToPostRequestInformation [AddUpdateEdgeRule API Docs](https://docs.bunny.net/reference/pullzonepublic_addedgerule)
// returns a *RequestInformation when successful
func (m *ItemEdgerulesAddOrUpdateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ibdd3cc897f8f3833299954aef4508f6bfff545c7346b36641af0214c751e5b85.EdgeRuleable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemEdgerulesAddOrUpdateRequestBuilder when successful
func (m *ItemEdgerulesAddOrUpdateRequestBuilder) WithUrl(rawUrl string)(*ItemEdgerulesAddOrUpdateRequestBuilder) {
    return NewItemEdgerulesAddOrUpdateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
