package shield

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/shield"
)

// RateLimitRequestBuilder builds and executes requests for operations under \shield\rate-limit
type RateLimitRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ById gets an item from the github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client.shield.rateLimit.item collection
// returns a *RateLimitRateLimitItemRequestBuilder when successful
func (m *RateLimitRequestBuilder) ById(id int32)(*RateLimitRateLimitItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(id), 10)
    return NewRateLimitRateLimitItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewRateLimitRequestBuilderInternal instantiates a new RateLimitRequestBuilder and sets the default values.
func NewRateLimitRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RateLimitRequestBuilder) {
    m := &RateLimitRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shield/rate-limit", pathParameters),
    }
    return m
}
// NewRateLimitRequestBuilder instantiates a new RateLimitRequestBuilder and sets the default values.
func NewRateLimitRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RateLimitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRateLimitRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a CustomWafRuleable when successful
// returns a UnauthorizedResult error when the service returns a 401 status code
func (m *RateLimitRequestBuilder) Post(ctx context.Context, body i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.CreateWafRateLimitRequestable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.CustomWafRuleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.CreateUnauthorizedResultFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.CreateCustomWafRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.CustomWafRuleable), nil
}
// returns a *RequestInformation when successful
func (m *RateLimitRequestBuilder) ToPostRequestInformation(ctx context.Context, body i04f2701327b610e1ec4da2c827c86b62c9eb445337fbacb67fed6b4d94fc7419.CreateWafRateLimitRequestable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json, text/plain;q=0.9")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RateLimitRequestBuilder when successful
func (m *RateLimitRequestBuilder) WithUrl(rawUrl string)(*RateLimitRequestBuilder) {
    return NewRateLimitRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
