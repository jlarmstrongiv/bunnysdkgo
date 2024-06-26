package storagezone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemStatisticsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The FileCountChart property
    fileCountChart ItemStatisticsGetResponse_FileCountChartable
    // The StorageUsedChart property
    storageUsedChart ItemStatisticsGetResponse_StorageUsedChartable
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
    res["FileCountChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemStatisticsGetResponse_FileCountChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileCountChart(val.(ItemStatisticsGetResponse_FileCountChartable))
        }
        return nil
    }
    res["StorageUsedChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemStatisticsGetResponse_StorageUsedChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageUsedChart(val.(ItemStatisticsGetResponse_StorageUsedChartable))
        }
        return nil
    }
    return res
}
// GetFileCountChart gets the FileCountChart property value. The FileCountChart property
// returns a ItemStatisticsGetResponse_FileCountChartable when successful
func (m *ItemStatisticsGetResponse) GetFileCountChart()(ItemStatisticsGetResponse_FileCountChartable) {
    return m.fileCountChart
}
// GetStorageUsedChart gets the StorageUsedChart property value. The StorageUsedChart property
// returns a ItemStatisticsGetResponse_StorageUsedChartable when successful
func (m *ItemStatisticsGetResponse) GetStorageUsedChart()(ItemStatisticsGetResponse_StorageUsedChartable) {
    return m.storageUsedChart
}
// Serialize serializes information the current object
func (m *ItemStatisticsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("FileCountChart", m.GetFileCountChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("StorageUsedChart", m.GetStorageUsedChart())
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
// SetFileCountChart sets the FileCountChart property value. The FileCountChart property
func (m *ItemStatisticsGetResponse) SetFileCountChart(value ItemStatisticsGetResponse_FileCountChartable)() {
    m.fileCountChart = value
}
// SetStorageUsedChart sets the StorageUsedChart property value. The StorageUsedChart property
func (m *ItemStatisticsGetResponse) SetStorageUsedChart(value ItemStatisticsGetResponse_StorageUsedChartable)() {
    m.storageUsedChart = value
}
type ItemStatisticsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFileCountChart()(ItemStatisticsGetResponse_FileCountChartable)
    GetStorageUsedChart()(ItemStatisticsGetResponse_StorageUsedChartable)
    SetFileCountChart(value ItemStatisticsGetResponse_FileCountChartable)()
    SetStorageUsedChart(value ItemStatisticsGetResponse_StorageUsedChartable)()
}
