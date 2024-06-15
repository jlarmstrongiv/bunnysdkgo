package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemSafehopStatisticsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The constructed chart of requests retried
    requestsRetriedChart ItemSafehopStatisticsGetResponse_RequestsRetriedChartable
    // The constructed chart of requests saved
    requestsSavedChart ItemSafehopStatisticsGetResponse_RequestsSavedChartable
    // The total number of retried requests
    totalRequestsRetried *int64
    // The total number of saved requests
    totalRequestsSaved *int64
}
// NewItemSafehopStatisticsGetResponse instantiates a new ItemSafehopStatisticsGetResponse and sets the default values.
func NewItemSafehopStatisticsGetResponse()(*ItemSafehopStatisticsGetResponse) {
    m := &ItemSafehopStatisticsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSafehopStatisticsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSafehopStatisticsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSafehopStatisticsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemSafehopStatisticsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemSafehopStatisticsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["RequestsRetriedChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemSafehopStatisticsGetResponse_RequestsRetriedChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestsRetriedChart(val.(ItemSafehopStatisticsGetResponse_RequestsRetriedChartable))
        }
        return nil
    }
    res["RequestsSavedChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemSafehopStatisticsGetResponse_RequestsSavedChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestsSavedChart(val.(ItemSafehopStatisticsGetResponse_RequestsSavedChartable))
        }
        return nil
    }
    res["TotalRequestsRetried"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalRequestsRetried(val)
        }
        return nil
    }
    res["TotalRequestsSaved"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalRequestsSaved(val)
        }
        return nil
    }
    return res
}
// GetRequestsRetriedChart gets the RequestsRetriedChart property value. The constructed chart of requests retried
// returns a ItemSafehopStatisticsGetResponse_RequestsRetriedChartable when successful
func (m *ItemSafehopStatisticsGetResponse) GetRequestsRetriedChart()(ItemSafehopStatisticsGetResponse_RequestsRetriedChartable) {
    return m.requestsRetriedChart
}
// GetRequestsSavedChart gets the RequestsSavedChart property value. The constructed chart of requests saved
// returns a ItemSafehopStatisticsGetResponse_RequestsSavedChartable when successful
func (m *ItemSafehopStatisticsGetResponse) GetRequestsSavedChart()(ItemSafehopStatisticsGetResponse_RequestsSavedChartable) {
    return m.requestsSavedChart
}
// GetTotalRequestsRetried gets the TotalRequestsRetried property value. The total number of retried requests
// returns a *int64 when successful
func (m *ItemSafehopStatisticsGetResponse) GetTotalRequestsRetried()(*int64) {
    return m.totalRequestsRetried
}
// GetTotalRequestsSaved gets the TotalRequestsSaved property value. The total number of saved requests
// returns a *int64 when successful
func (m *ItemSafehopStatisticsGetResponse) GetTotalRequestsSaved()(*int64) {
    return m.totalRequestsSaved
}
// Serialize serializes information the current object
func (m *ItemSafehopStatisticsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("RequestsRetriedChart", m.GetRequestsRetriedChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("RequestsSavedChart", m.GetRequestsSavedChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("TotalRequestsRetried", m.GetTotalRequestsRetried())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("TotalRequestsSaved", m.GetTotalRequestsSaved())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSafehopStatisticsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRequestsRetriedChart sets the RequestsRetriedChart property value. The constructed chart of requests retried
func (m *ItemSafehopStatisticsGetResponse) SetRequestsRetriedChart(value ItemSafehopStatisticsGetResponse_RequestsRetriedChartable)() {
    m.requestsRetriedChart = value
}
// SetRequestsSavedChart sets the RequestsSavedChart property value. The constructed chart of requests saved
func (m *ItemSafehopStatisticsGetResponse) SetRequestsSavedChart(value ItemSafehopStatisticsGetResponse_RequestsSavedChartable)() {
    m.requestsSavedChart = value
}
// SetTotalRequestsRetried sets the TotalRequestsRetried property value. The total number of retried requests
func (m *ItemSafehopStatisticsGetResponse) SetTotalRequestsRetried(value *int64)() {
    m.totalRequestsRetried = value
}
// SetTotalRequestsSaved sets the TotalRequestsSaved property value. The total number of saved requests
func (m *ItemSafehopStatisticsGetResponse) SetTotalRequestsSaved(value *int64)() {
    m.totalRequestsSaved = value
}
type ItemSafehopStatisticsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRequestsRetriedChart()(ItemSafehopStatisticsGetResponse_RequestsRetriedChartable)
    GetRequestsSavedChart()(ItemSafehopStatisticsGetResponse_RequestsSavedChartable)
    GetTotalRequestsRetried()(*int64)
    GetTotalRequestsSaved()(*int64)
    SetRequestsRetriedChart(value ItemSafehopStatisticsGetResponse_RequestsRetriedChartable)()
    SetRequestsSavedChart(value ItemSafehopStatisticsGetResponse_RequestsSavedChartable)()
    SetTotalRequestsRetried(value *int64)()
    SetTotalRequestsSaved(value *int64)()
}
