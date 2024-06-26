package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemOriginshieldQueuestatisticsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The constructed chart of origin shield concurrent requests
    concurrentRequestsChart ItemOriginshieldQueuestatisticsGetResponse_ConcurrentRequestsChartable
    // The constructed chart of origin shield requests chart
    queuedRequestsChart ItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChartable
}
// NewItemOriginshieldQueuestatisticsGetResponse instantiates a new ItemOriginshieldQueuestatisticsGetResponse and sets the default values.
func NewItemOriginshieldQueuestatisticsGetResponse()(*ItemOriginshieldQueuestatisticsGetResponse) {
    m := &ItemOriginshieldQueuestatisticsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemOriginshieldQueuestatisticsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemOriginshieldQueuestatisticsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOriginshieldQueuestatisticsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemOriginshieldQueuestatisticsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConcurrentRequestsChart gets the ConcurrentRequestsChart property value. The constructed chart of origin shield concurrent requests
// returns a ItemOriginshieldQueuestatisticsGetResponse_ConcurrentRequestsChartable when successful
func (m *ItemOriginshieldQueuestatisticsGetResponse) GetConcurrentRequestsChart()(ItemOriginshieldQueuestatisticsGetResponse_ConcurrentRequestsChartable) {
    return m.concurrentRequestsChart
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemOriginshieldQueuestatisticsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ConcurrentRequestsChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemOriginshieldQueuestatisticsGetResponse_ConcurrentRequestsChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConcurrentRequestsChart(val.(ItemOriginshieldQueuestatisticsGetResponse_ConcurrentRequestsChartable))
        }
        return nil
    }
    res["QueuedRequestsChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueuedRequestsChart(val.(ItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChartable))
        }
        return nil
    }
    return res
}
// GetQueuedRequestsChart gets the QueuedRequestsChart property value. The constructed chart of origin shield requests chart
// returns a ItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChartable when successful
func (m *ItemOriginshieldQueuestatisticsGetResponse) GetQueuedRequestsChart()(ItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChartable) {
    return m.queuedRequestsChart
}
// Serialize serializes information the current object
func (m *ItemOriginshieldQueuestatisticsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("ConcurrentRequestsChart", m.GetConcurrentRequestsChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("QueuedRequestsChart", m.GetQueuedRequestsChart())
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
func (m *ItemOriginshieldQueuestatisticsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConcurrentRequestsChart sets the ConcurrentRequestsChart property value. The constructed chart of origin shield concurrent requests
func (m *ItemOriginshieldQueuestatisticsGetResponse) SetConcurrentRequestsChart(value ItemOriginshieldQueuestatisticsGetResponse_ConcurrentRequestsChartable)() {
    m.concurrentRequestsChart = value
}
// SetQueuedRequestsChart sets the QueuedRequestsChart property value. The constructed chart of origin shield requests chart
func (m *ItemOriginshieldQueuestatisticsGetResponse) SetQueuedRequestsChart(value ItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChartable)() {
    m.queuedRequestsChart = value
}
type ItemOriginshieldQueuestatisticsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConcurrentRequestsChart()(ItemOriginshieldQueuestatisticsGetResponse_ConcurrentRequestsChartable)
    GetQueuedRequestsChart()(ItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChartable)
    SetConcurrentRequestsChart(value ItemOriginshieldQueuestatisticsGetResponse_ConcurrentRequestsChartable)()
    SetQueuedRequestsChart(value ItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChartable)()
}
