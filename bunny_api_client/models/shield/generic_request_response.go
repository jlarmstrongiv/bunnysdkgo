package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GenericRequestResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The errorKey property
    errorKey *string
    // The message property
    message *string
    // The statusCode property
    statusCode *float64
    // The success property
    success *bool
}
// NewGenericRequestResponse instantiates a new GenericRequestResponse and sets the default values.
func NewGenericRequestResponse()(*GenericRequestResponse) {
    m := &GenericRequestResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGenericRequestResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGenericRequestResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGenericRequestResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GenericRequestResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetErrorKey gets the errorKey property value. The errorKey property
// returns a *string when successful
func (m *GenericRequestResponse) GetErrorKey()(*string) {
    return m.errorKey
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GenericRequestResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["errorKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorKey(val)
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
        val, err := n.GetFloat64Value()
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
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *GenericRequestResponse) GetMessage()(*string) {
    return m.message
}
// GetStatusCode gets the statusCode property value. The statusCode property
// returns a *float64 when successful
func (m *GenericRequestResponse) GetStatusCode()(*float64) {
    return m.statusCode
}
// GetSuccess gets the success property value. The success property
// returns a *bool when successful
func (m *GenericRequestResponse) GetSuccess()(*bool) {
    return m.success
}
// Serialize serializes information the current object
func (m *GenericRequestResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("errorKey", m.GetErrorKey())
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
        err := writer.WriteFloat64Value("statusCode", m.GetStatusCode())
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
func (m *GenericRequestResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetErrorKey sets the errorKey property value. The errorKey property
func (m *GenericRequestResponse) SetErrorKey(value *string)() {
    m.errorKey = value
}
// SetMessage sets the message property value. The message property
func (m *GenericRequestResponse) SetMessage(value *string)() {
    m.message = value
}
// SetStatusCode sets the statusCode property value. The statusCode property
func (m *GenericRequestResponse) SetStatusCode(value *float64)() {
    m.statusCode = value
}
// SetSuccess sets the success property value. The success property
func (m *GenericRequestResponse) SetSuccess(value *bool)() {
    m.success = value
}
type GenericRequestResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetErrorKey()(*string)
    GetMessage()(*string)
    GetStatusCode()(*float64)
    GetSuccess()(*bool)
    SetErrorKey(value *string)()
    SetMessage(value *string)()
    SetStatusCode(value *float64)()
    SetSuccess(value *bool)()
}
