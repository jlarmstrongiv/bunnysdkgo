package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemEdgerulesItemSetEdgeRuleEnabledRequestBuilder builds and executes requests for operations under \pullzone\{-id}\edgerules\{edgeRuleId}\setEdgeRuleEnabled
type ItemEdgerulesItemSetEdgeRuleEnabledRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemEdgerulesItemSetEdgeRuleEnabledRequestBuilderInternal instantiates a new ItemEdgerulesItemSetEdgeRuleEnabledRequestBuilder and sets the default values.
func NewItemEdgerulesItemSetEdgeRuleEnabledRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEdgerulesItemSetEdgeRuleEnabledRequestBuilder) {
    m := &ItemEdgerulesItemSetEdgeRuleEnabledRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/edgerules/{edgeRuleId}/setEdgeRuleEnabled", pathParameters),
    }
    return m
}
// NewItemEdgerulesItemSetEdgeRuleEnabledRequestBuilder instantiates a new ItemEdgerulesItemSetEdgeRuleEnabledRequestBuilder and sets the default values.
func NewItemEdgerulesItemSetEdgeRuleEnabledRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEdgerulesItemSetEdgeRuleEnabledRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEdgerulesItemSetEdgeRuleEnabledRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [SetEdgeRuleEnabled API Docs](https://docs.bunny.net/reference/pullzonepublic_setedgeruleenabled)
func (m *ItemEdgerulesItemSetEdgeRuleEnabledRequestBuilder) Post(ctx context.Context, body ItemEdgerulesItemSetEdgeRuleEnabledPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPostRequestInformation [SetEdgeRuleEnabled API Docs](https://docs.bunny.net/reference/pullzonepublic_setedgeruleenabled)
// returns a *RequestInformation when successful
func (m *ItemEdgerulesItemSetEdgeRuleEnabledRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemEdgerulesItemSetEdgeRuleEnabledPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemEdgerulesItemSetEdgeRuleEnabledRequestBuilder when successful
func (m *ItemEdgerulesItemSetEdgeRuleEnabledRequestBuilder) WithUrl(rawUrl string)(*ItemEdgerulesItemSetEdgeRuleEnabledRequestBuilder) {
    return NewItemEdgerulesItemSetEdgeRuleEnabledRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
