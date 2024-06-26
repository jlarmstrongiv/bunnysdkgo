package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChart the constructed chart of origin shield requests chart
type ItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChart struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChart instantiates a new ItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChart and sets the default values.
func NewItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChart()(*ItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChart) {
    m := &ItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChart{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChartFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChartFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChart(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChart) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChart) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChart) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChart) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type ItemOriginshieldQueuestatisticsGetResponse_QueuedRequestsChartable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
