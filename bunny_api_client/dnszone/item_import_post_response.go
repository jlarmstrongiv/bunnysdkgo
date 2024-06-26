package dnszone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemImportPostResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The RecordsFailed property
    recordsFailed *int32
    // The RecordsSkipped property
    recordsSkipped *int32
    // The RecordsSuccessful property
    recordsSuccessful *int32
}
// NewItemImportPostResponse instantiates a new ItemImportPostResponse and sets the default values.
func NewItemImportPostResponse()(*ItemImportPostResponse) {
    m := &ItemImportPostResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemImportPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemImportPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemImportPostResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemImportPostResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemImportPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["RecordsFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordsFailed(val)
        }
        return nil
    }
    res["RecordsSkipped"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordsSkipped(val)
        }
        return nil
    }
    res["RecordsSuccessful"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordsSuccessful(val)
        }
        return nil
    }
    return res
}
// GetRecordsFailed gets the RecordsFailed property value. The RecordsFailed property
// returns a *int32 when successful
func (m *ItemImportPostResponse) GetRecordsFailed()(*int32) {
    return m.recordsFailed
}
// GetRecordsSkipped gets the RecordsSkipped property value. The RecordsSkipped property
// returns a *int32 when successful
func (m *ItemImportPostResponse) GetRecordsSkipped()(*int32) {
    return m.recordsSkipped
}
// GetRecordsSuccessful gets the RecordsSuccessful property value. The RecordsSuccessful property
// returns a *int32 when successful
func (m *ItemImportPostResponse) GetRecordsSuccessful()(*int32) {
    return m.recordsSuccessful
}
// Serialize serializes information the current object
func (m *ItemImportPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("RecordsFailed", m.GetRecordsFailed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("RecordsSkipped", m.GetRecordsSkipped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("RecordsSuccessful", m.GetRecordsSuccessful())
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
func (m *ItemImportPostResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRecordsFailed sets the RecordsFailed property value. The RecordsFailed property
func (m *ItemImportPostResponse) SetRecordsFailed(value *int32)() {
    m.recordsFailed = value
}
// SetRecordsSkipped sets the RecordsSkipped property value. The RecordsSkipped property
func (m *ItemImportPostResponse) SetRecordsSkipped(value *int32)() {
    m.recordsSkipped = value
}
// SetRecordsSuccessful sets the RecordsSuccessful property value. The RecordsSuccessful property
func (m *ItemImportPostResponse) SetRecordsSuccessful(value *int32)() {
    m.recordsSuccessful = value
}
type ItemImportPostResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRecordsFailed()(*int32)
    GetRecordsSkipped()(*int32)
    GetRecordsSuccessful()(*int32)
    SetRecordsFailed(value *int32)()
    SetRecordsSkipped(value *int32)()
    SetRecordsSuccessful(value *int32)()
}
