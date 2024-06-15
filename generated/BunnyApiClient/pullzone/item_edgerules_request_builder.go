package pullzone

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemEdgerulesRequestBuilder builds and executes requests for operations under \pullzone\{-id}\edgerules
type ItemEdgerulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AddOrUpdate the addOrUpdate property
// returns a *ItemEdgerulesAddorupdateAddOrUpdateRequestBuilder when successful
func (m *ItemEdgerulesRequestBuilder) AddOrUpdate()(*ItemEdgerulesAddorupdateAddOrUpdateRequestBuilder) {
    return NewItemEdgerulesAddorupdateAddOrUpdateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByEdgeRuleId gets an item from the github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient.pullzone.item.edgerules.item collection
// returns a *ItemEdgerulesWithEdgeRuleItemRequestBuilder when successful
func (m *ItemEdgerulesRequestBuilder) ByEdgeRuleId(edgeRuleId string)(*ItemEdgerulesWithEdgeRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if edgeRuleId != "" {
        urlTplParams["edgeRuleId"] = edgeRuleId
    }
    return NewItemEdgerulesWithEdgeRuleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemEdgerulesRequestBuilderInternal instantiates a new ItemEdgerulesRequestBuilder and sets the default values.
func NewItemEdgerulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEdgerulesRequestBuilder) {
    m := &ItemEdgerulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/edgerules", pathParameters),
    }
    return m
}
// NewItemEdgerulesRequestBuilder instantiates a new ItemEdgerulesRequestBuilder and sets the default values.
func NewItemEdgerulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEdgerulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEdgerulesRequestBuilderInternal(urlParams, requestAdapter)
}
