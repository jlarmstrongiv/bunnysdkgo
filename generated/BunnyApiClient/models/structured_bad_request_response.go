package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StructuredBadRequestResponse the server could not understand the request due to invalid syntax.
type StructuredBadRequestResponse struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ErrorKey property
    errorKey *string
    // The Field property
    field *string
    // The Message property
    message *string
}
// NewStructuredBadRequestResponse instantiates a new StructuredBadRequestResponse and sets the default values.
func NewStructuredBadRequestResponse()(*StructuredBadRequestResponse) {
    m := &StructuredBadRequestResponse{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStructuredBadRequestResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStructuredBadRequestResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStructuredBadRequestResponse(), nil
}
// Error the primary error message.
// returns a string when successful
func (m *StructuredBadRequestResponse) Error()(string) {
    return m.ApiError.Error()
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StructuredBadRequestResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetErrorKey gets the ErrorKey property value. The ErrorKey property
// returns a *string when successful
func (m *StructuredBadRequestResponse) GetErrorKey()(*string) {
    return m.errorKey
}
// GetField gets the Field property value. The Field property
// returns a *string when successful
func (m *StructuredBadRequestResponse) GetField()(*string) {
    return m.field
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StructuredBadRequestResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ErrorKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorKey(val)
        }
        return nil
    }
    res["Field"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetField(val)
        }
        return nil
    }
    res["Message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the Message property value. The Message property
// returns a *string when successful
func (m *StructuredBadRequestResponse) GetMessage()(*string) {
    return m.message
}
// Serialize serializes information the current object
func (m *StructuredBadRequestResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("ErrorKey", m.GetErrorKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Field", m.GetField())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Message", m.GetMessage())
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
func (m *StructuredBadRequestResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetErrorKey sets the ErrorKey property value. The ErrorKey property
func (m *StructuredBadRequestResponse) SetErrorKey(value *string)() {
    m.errorKey = value
}
// SetField sets the Field property value. The Field property
func (m *StructuredBadRequestResponse) SetField(value *string)() {
    m.field = value
}
// SetMessage sets the Message property value. The Message property
func (m *StructuredBadRequestResponse) SetMessage(value *string)() {
    m.message = value
}
type StructuredBadRequestResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetErrorKey()(*string)
    GetField()(*string)
    GetMessage()(*string)
    SetErrorKey(value *string)()
    SetField(value *string)()
    SetMessage(value *string)()
}
