package cleanupunconfiguredresolutions

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CleanupUnconfiguredResolutionsResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The resolutions were successfully deleted
    data CleanupUnconfiguredResolutionsResult_dataable
    // Response message description
    message *string
    // The response status code
    statusCode *int32
    // Determines if the request was successful
    success *bool
}
// NewCleanupUnconfiguredResolutionsResult instantiates a new CleanupUnconfiguredResolutionsResult and sets the default values.
func NewCleanupUnconfiguredResolutionsResult()(*CleanupUnconfiguredResolutionsResult) {
    m := &CleanupUnconfiguredResolutionsResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCleanupUnconfiguredResolutionsResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCleanupUnconfiguredResolutionsResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCleanupUnconfiguredResolutionsResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CleanupUnconfiguredResolutionsResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetData gets the data property value. The resolutions were successfully deleted
// returns a CleanupUnconfiguredResolutionsResult_dataable when successful
func (m *CleanupUnconfiguredResolutionsResult) GetData()(CleanupUnconfiguredResolutionsResult_dataable) {
    return m.data
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CleanupUnconfiguredResolutionsResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCleanupUnconfiguredResolutionsResult_dataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetData(val.(CleanupUnconfiguredResolutionsResult_dataable))
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["statusCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusCode(val)
        }
        return nil
    }
    res["success"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccess(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. Response message description
// returns a *string when successful
func (m *CleanupUnconfiguredResolutionsResult) GetMessage()(*string) {
    return m.message
}
// GetStatusCode gets the statusCode property value. The response status code
// returns a *int32 when successful
func (m *CleanupUnconfiguredResolutionsResult) GetStatusCode()(*int32) {
    return m.statusCode
}
// GetSuccess gets the success property value. Determines if the request was successful
// returns a *bool when successful
func (m *CleanupUnconfiguredResolutionsResult) GetSuccess()(*bool) {
    return m.success
}
// Serialize serializes information the current object
func (m *CleanupUnconfiguredResolutionsResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("data", m.GetData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("statusCode", m.GetStatusCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("success", m.GetSuccess())
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
func (m *CleanupUnconfiguredResolutionsResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetData sets the data property value. The resolutions were successfully deleted
func (m *CleanupUnconfiguredResolutionsResult) SetData(value CleanupUnconfiguredResolutionsResult_dataable)() {
    m.data = value
}
// SetMessage sets the message property value. Response message description
func (m *CleanupUnconfiguredResolutionsResult) SetMessage(value *string)() {
    m.message = value
}
// SetStatusCode sets the statusCode property value. The response status code
func (m *CleanupUnconfiguredResolutionsResult) SetStatusCode(value *int32)() {
    m.statusCode = value
}
// SetSuccess sets the success property value. Determines if the request was successful
func (m *CleanupUnconfiguredResolutionsResult) SetSuccess(value *bool)() {
    m.success = value
}
type CleanupUnconfiguredResolutionsResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetData()(CleanupUnconfiguredResolutionsResult_dataable)
    GetMessage()(*string)
    GetStatusCode()(*int32)
    GetSuccess()(*bool)
    SetData(value CleanupUnconfiguredResolutionsResult_dataable)()
    SetMessage(value *string)()
    SetStatusCode(value *int32)()
    SetSuccess(value *bool)()
}
