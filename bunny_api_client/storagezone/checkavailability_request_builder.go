package storagezone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CheckavailabilityRequestBuilder builds and executes requests for operations under \storagezone\checkavailability
type CheckavailabilityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewCheckavailabilityRequestBuilderInternal instantiates a new CheckavailabilityRequestBuilder and sets the default values.
func NewCheckavailabilityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CheckavailabilityRequestBuilder) {
    m := &CheckavailabilityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storagezone/checkavailability", pathParameters),
    }
    return m
}
// NewCheckavailabilityRequestBuilder instantiates a new CheckavailabilityRequestBuilder and sets the default values.
func NewCheckavailabilityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CheckavailabilityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCheckavailabilityRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [CheckTheStorageZoneAvailability API Docs](https://docs.bunny.net/reference/storagezonepublic_checkavailability)
// returns a CheckavailabilityPostResponseable when successful
func (m *CheckavailabilityRequestBuilder) Post(ctx context.Context, body CheckavailabilityPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(CheckavailabilityPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCheckavailabilityPostResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CheckavailabilityPostResponseable), nil
}
// ToPostRequestInformation [CheckTheStorageZoneAvailability API Docs](https://docs.bunny.net/reference/storagezonepublic_checkavailability)
// returns a *RequestInformation when successful
func (m *CheckavailabilityRequestBuilder) ToPostRequestInformation(ctx context.Context, body CheckavailabilityPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CheckavailabilityRequestBuilder when successful
func (m *CheckavailabilityRequestBuilder) WithUrl(rawUrl string)(*CheckavailabilityRequestBuilder) {
    return NewCheckavailabilityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
