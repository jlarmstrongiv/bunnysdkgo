package dnszone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemStatisticsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The QueriesByTypeChart property
    queriesByTypeChart ItemStatisticsGetResponse_QueriesByTypeChartable
    // The QueriesServedChart property
    queriesServedChart ItemStatisticsGetResponse_QueriesServedChartable
    // The TotalQueriesServed property
    totalQueriesServed *int64
}
// NewItemStatisticsGetResponse instantiates a new ItemStatisticsGetResponse and sets the default values.
func NewItemStatisticsGetResponse()(*ItemStatisticsGetResponse) {
    m := &ItemStatisticsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemStatisticsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemStatisticsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemStatisticsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemStatisticsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemStatisticsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["QueriesByTypeChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemStatisticsGetResponse_QueriesByTypeChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueriesByTypeChart(val.(ItemStatisticsGetResponse_QueriesByTypeChartable))
        }
        return nil
    }
    res["QueriesServedChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemStatisticsGetResponse_QueriesServedChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueriesServedChart(val.(ItemStatisticsGetResponse_QueriesServedChartable))
        }
        return nil
    }
    res["TotalQueriesServed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalQueriesServed(val)
        }
        return nil
    }
    return res
}
// GetQueriesByTypeChart gets the QueriesByTypeChart property value. The QueriesByTypeChart property
// returns a ItemStatisticsGetResponse_QueriesByTypeChartable when successful
func (m *ItemStatisticsGetResponse) GetQueriesByTypeChart()(ItemStatisticsGetResponse_QueriesByTypeChartable) {
    return m.queriesByTypeChart
}
// GetQueriesServedChart gets the QueriesServedChart property value. The QueriesServedChart property
// returns a ItemStatisticsGetResponse_QueriesServedChartable when successful
func (m *ItemStatisticsGetResponse) GetQueriesServedChart()(ItemStatisticsGetResponse_QueriesServedChartable) {
    return m.queriesServedChart
}
// GetTotalQueriesServed gets the TotalQueriesServed property value. The TotalQueriesServed property
// returns a *int64 when successful
func (m *ItemStatisticsGetResponse) GetTotalQueriesServed()(*int64) {
    return m.totalQueriesServed
}
// Serialize serializes information the current object
func (m *ItemStatisticsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("QueriesByTypeChart", m.GetQueriesByTypeChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("QueriesServedChart", m.GetQueriesServedChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("TotalQueriesServed", m.GetTotalQueriesServed())
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
func (m *ItemStatisticsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetQueriesByTypeChart sets the QueriesByTypeChart property value. The QueriesByTypeChart property
func (m *ItemStatisticsGetResponse) SetQueriesByTypeChart(value ItemStatisticsGetResponse_QueriesByTypeChartable)() {
    m.queriesByTypeChart = value
}
// SetQueriesServedChart sets the QueriesServedChart property value. The QueriesServedChart property
func (m *ItemStatisticsGetResponse) SetQueriesServedChart(value ItemStatisticsGetResponse_QueriesServedChartable)() {
    m.queriesServedChart = value
}
// SetTotalQueriesServed sets the TotalQueriesServed property value. The TotalQueriesServed property
func (m *ItemStatisticsGetResponse) SetTotalQueriesServed(value *int64)() {
    m.totalQueriesServed = value
}
type ItemStatisticsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetQueriesByTypeChart()(ItemStatisticsGetResponse_QueriesByTypeChartable)
    GetQueriesServedChart()(ItemStatisticsGetResponse_QueriesServedChartable)
    GetTotalQueriesServed()(*int64)
    SetQueriesByTypeChart(value ItemStatisticsGetResponse_QueriesByTypeChartable)()
    SetQueriesServedChart(value ItemStatisticsGetResponse_QueriesServedChartable)()
    SetTotalQueriesServed(value *int64)()
}
