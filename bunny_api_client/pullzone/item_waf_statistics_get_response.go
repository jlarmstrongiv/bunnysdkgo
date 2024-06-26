package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemWafStatisticsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ThreatsBlockedChart property
    threatsBlockedChart ItemWafStatisticsGetResponse_ThreatsBlockedChartable
    // The ThreatsCheckedChart property
    threatsCheckedChart ItemWafStatisticsGetResponse_ThreatsCheckedChartable
    // The TotalBlockedRequests property
    totalBlockedRequests *int64
    // The TotalCheckedRequests property
    totalCheckedRequests *int64
}
// NewItemWafStatisticsGetResponse instantiates a new ItemWafStatisticsGetResponse and sets the default values.
func NewItemWafStatisticsGetResponse()(*ItemWafStatisticsGetResponse) {
    m := &ItemWafStatisticsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemWafStatisticsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemWafStatisticsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemWafStatisticsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemWafStatisticsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemWafStatisticsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ThreatsBlockedChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemWafStatisticsGetResponse_ThreatsBlockedChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreatsBlockedChart(val.(ItemWafStatisticsGetResponse_ThreatsBlockedChartable))
        }
        return nil
    }
    res["ThreatsCheckedChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemWafStatisticsGetResponse_ThreatsCheckedChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreatsCheckedChart(val.(ItemWafStatisticsGetResponse_ThreatsCheckedChartable))
        }
        return nil
    }
    res["TotalBlockedRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalBlockedRequests(val)
        }
        return nil
    }
    res["TotalCheckedRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCheckedRequests(val)
        }
        return nil
    }
    return res
}
// GetThreatsBlockedChart gets the ThreatsBlockedChart property value. The ThreatsBlockedChart property
// returns a ItemWafStatisticsGetResponse_ThreatsBlockedChartable when successful
func (m *ItemWafStatisticsGetResponse) GetThreatsBlockedChart()(ItemWafStatisticsGetResponse_ThreatsBlockedChartable) {
    return m.threatsBlockedChart
}
// GetThreatsCheckedChart gets the ThreatsCheckedChart property value. The ThreatsCheckedChart property
// returns a ItemWafStatisticsGetResponse_ThreatsCheckedChartable when successful
func (m *ItemWafStatisticsGetResponse) GetThreatsCheckedChart()(ItemWafStatisticsGetResponse_ThreatsCheckedChartable) {
    return m.threatsCheckedChart
}
// GetTotalBlockedRequests gets the TotalBlockedRequests property value. The TotalBlockedRequests property
// returns a *int64 when successful
func (m *ItemWafStatisticsGetResponse) GetTotalBlockedRequests()(*int64) {
    return m.totalBlockedRequests
}
// GetTotalCheckedRequests gets the TotalCheckedRequests property value. The TotalCheckedRequests property
// returns a *int64 when successful
func (m *ItemWafStatisticsGetResponse) GetTotalCheckedRequests()(*int64) {
    return m.totalCheckedRequests
}
// Serialize serializes information the current object
func (m *ItemWafStatisticsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("ThreatsBlockedChart", m.GetThreatsBlockedChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("ThreatsCheckedChart", m.GetThreatsCheckedChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("TotalBlockedRequests", m.GetTotalBlockedRequests())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("TotalCheckedRequests", m.GetTotalCheckedRequests())
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
func (m *ItemWafStatisticsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetThreatsBlockedChart sets the ThreatsBlockedChart property value. The ThreatsBlockedChart property
func (m *ItemWafStatisticsGetResponse) SetThreatsBlockedChart(value ItemWafStatisticsGetResponse_ThreatsBlockedChartable)() {
    m.threatsBlockedChart = value
}
// SetThreatsCheckedChart sets the ThreatsCheckedChart property value. The ThreatsCheckedChart property
func (m *ItemWafStatisticsGetResponse) SetThreatsCheckedChart(value ItemWafStatisticsGetResponse_ThreatsCheckedChartable)() {
    m.threatsCheckedChart = value
}
// SetTotalBlockedRequests sets the TotalBlockedRequests property value. The TotalBlockedRequests property
func (m *ItemWafStatisticsGetResponse) SetTotalBlockedRequests(value *int64)() {
    m.totalBlockedRequests = value
}
// SetTotalCheckedRequests sets the TotalCheckedRequests property value. The TotalCheckedRequests property
func (m *ItemWafStatisticsGetResponse) SetTotalCheckedRequests(value *int64)() {
    m.totalCheckedRequests = value
}
type ItemWafStatisticsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetThreatsBlockedChart()(ItemWafStatisticsGetResponse_ThreatsBlockedChartable)
    GetThreatsCheckedChart()(ItemWafStatisticsGetResponse_ThreatsCheckedChartable)
    GetTotalBlockedRequests()(*int64)
    GetTotalCheckedRequests()(*int64)
    SetThreatsBlockedChart(value ItemWafStatisticsGetResponse_ThreatsBlockedChartable)()
    SetThreatsCheckedChart(value ItemWafStatisticsGetResponse_ThreatsCheckedChartable)()
    SetTotalBlockedRequests(value *int64)()
    SetTotalCheckedRequests(value *int64)()
}
