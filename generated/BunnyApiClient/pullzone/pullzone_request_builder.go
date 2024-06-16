package pullzone

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models/pullzone"
)

// PullzoneRequestBuilder builds and executes requests for operations under \pullzone
type PullzoneRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PullzoneRequestBuilderGetQueryParameters [ListPullZones API Docs](https://docs.bunny.net/reference/pullzonepublic_index)
type PullzoneRequestBuilderGetQueryParameters struct {
    // Determines if the result hostnames should contain the SSL certificate
    IncludeCertificate *bool `uriparametername:"includeCertificate"`
    Page *int32 `uriparametername:"page"`
    PerPage *int32 `uriparametername:"perPage"`
    // The search term that will be used to filter the results
    Search *string `uriparametername:"search"`
}
// ById gets an item from the github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient.pullzone.item collection
// returns a *ItemRequestBuilder when successful
func (m *PullzoneRequestBuilder) ById(id int64)(*ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["%2Did"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(id, 10)
    return NewItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Checkavailability the checkavailability property
// returns a *CheckavailabilityRequestBuilder when successful
func (m *PullzoneRequestBuilder) Checkavailability()(*CheckavailabilityRequestBuilder) {
    return NewCheckavailabilityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPullzoneRequestBuilderInternal instantiates a new PullzoneRequestBuilder and sets the default values.
func NewPullzoneRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PullzoneRequestBuilder) {
    m := &PullzoneRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone?includeCertificate={includeCertificate}&page={page}&perPage={perPage}{&search*}", pathParameters),
    }
    return m
}
// NewPullzoneRequestBuilder instantiates a new PullzoneRequestBuilder and sets the default values.
func NewPullzoneRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PullzoneRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPullzoneRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [ListPullZones API Docs](https://docs.bunny.net/reference/pullzonepublic_index)
// returns a PullzoneGetResponseable when successful
func (m *PullzoneRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[PullzoneRequestBuilderGetQueryParameters])(PullzoneGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreatePullzoneGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(PullzoneGetResponseable), nil
}
// LoadFreeCertificate the loadFreeCertificate property
// returns a *LoadfreecertificateLoadFreeCertificateRequestBuilder when successful
func (m *PullzoneRequestBuilder) LoadFreeCertificate()(*LoadfreecertificateLoadFreeCertificateRequestBuilder) {
    return NewLoadfreecertificateLoadFreeCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post [AddPullZone API Docs](https://docs.bunny.net/reference/pullzonepublic_add)
// returns a PullZoneable when successful
func (m *PullzoneRequestBuilder) Post(ctx context.Context, body id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.PullZoneCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.PullZoneable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.CreatePullZoneFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.PullZoneable), nil
}
// ToGetRequestInformation [ListPullZones API Docs](https://docs.bunny.net/reference/pullzonepublic_index)
// returns a *RequestInformation when successful
func (m *PullzoneRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[PullzoneRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation [AddPullZone API Docs](https://docs.bunny.net/reference/pullzonepublic_add)
// returns a *RequestInformation when successful
func (m *PullzoneRequestBuilder) ToPostRequestInformation(ctx context.Context, body id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.PullZoneCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/pullzone", m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PullzoneRequestBuilder when successful
func (m *PullzoneRequestBuilder) WithUrl(rawUrl string)(*PullzoneRequestBuilder) {
    return NewPullzoneRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
