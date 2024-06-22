package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemWafStatisticsGetResponse_ThreatsCheckedChart struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewItemWafStatisticsGetResponse_ThreatsCheckedChart instantiates a new ItemWafStatisticsGetResponse_ThreatsCheckedChart and sets the default values.
func NewItemWafStatisticsGetResponse_ThreatsCheckedChart()(*ItemWafStatisticsGetResponse_ThreatsCheckedChart) {
    m := &ItemWafStatisticsGetResponse_ThreatsCheckedChart{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemWafStatisticsGetResponse_ThreatsCheckedChartFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemWafStatisticsGetResponse_ThreatsCheckedChartFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemWafStatisticsGetResponse_ThreatsCheckedChart(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemWafStatisticsGetResponse_ThreatsCheckedChart) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemWafStatisticsGetResponse_ThreatsCheckedChart) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemWafStatisticsGetResponse_ThreatsCheckedChart) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemWafStatisticsGetResponse_ThreatsCheckedChart) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type ItemWafStatisticsGetResponse_ThreatsCheckedChartable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
