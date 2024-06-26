package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemEdgerulesWithEdgeRuleItemRequestBuilder builds and executes requests for operations under \pullzone\{-id}\edgerules\{edgeRuleId}
type ItemEdgerulesWithEdgeRuleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemEdgerulesWithEdgeRuleItemRequestBuilderInternal instantiates a new ItemEdgerulesWithEdgeRuleItemRequestBuilder and sets the default values.
func NewItemEdgerulesWithEdgeRuleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEdgerulesWithEdgeRuleItemRequestBuilder) {
    m := &ItemEdgerulesWithEdgeRuleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}/edgerules/{edgeRuleId}", pathParameters),
    }
    return m
}
// NewItemEdgerulesWithEdgeRuleItemRequestBuilder instantiates a new ItemEdgerulesWithEdgeRuleItemRequestBuilder and sets the default values.
func NewItemEdgerulesWithEdgeRuleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEdgerulesWithEdgeRuleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEdgerulesWithEdgeRuleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete [DeleteEdgeRule API Docs](https://docs.bunny.net/reference/pullzonepublic_deleteedgerule)
func (m *ItemEdgerulesWithEdgeRuleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// SetEdgeRuleEnabled the setEdgeRuleEnabled property
// returns a *ItemEdgerulesItemSetedgeruleenabledSetEdgeRuleEnabledRequestBuilder when successful
func (m *ItemEdgerulesWithEdgeRuleItemRequestBuilder) SetEdgeRuleEnabled()(*ItemEdgerulesItemSetedgeruleenabledSetEdgeRuleEnabledRequestBuilder) {
    return NewItemEdgerulesItemSetedgeruleenabledSetEdgeRuleEnabledRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation [DeleteEdgeRule API Docs](https://docs.bunny.net/reference/pullzonepublic_deleteedgerule)
// returns a *RequestInformation when successful
func (m *ItemEdgerulesWithEdgeRuleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemEdgerulesWithEdgeRuleItemRequestBuilder when successful
func (m *ItemEdgerulesWithEdgeRuleItemRequestBuilder) WithUrl(rawUrl string)(*ItemEdgerulesWithEdgeRuleItemRequestBuilder) {
    return NewItemEdgerulesWithEdgeRuleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
