package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemEdgerulesItemSetEdgeRuleEnabledPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Id property
    id *int64
    // The Value property
    value *bool
}
// NewItemEdgerulesItemSetEdgeRuleEnabledPostRequestBody instantiates a new ItemEdgerulesItemSetEdgeRuleEnabledPostRequestBody and sets the default values.
func NewItemEdgerulesItemSetEdgeRuleEnabledPostRequestBody()(*ItemEdgerulesItemSetEdgeRuleEnabledPostRequestBody) {
    m := &ItemEdgerulesItemSetEdgeRuleEnabledPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemEdgerulesItemSetEdgeRuleEnabledPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemEdgerulesItemSetEdgeRuleEnabledPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemEdgerulesItemSetEdgeRuleEnabledPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemEdgerulesItemSetEdgeRuleEnabledPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemEdgerulesItemSetEdgeRuleEnabledPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["Value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetId gets the Id property value. The Id property
// returns a *int64 when successful
func (m *ItemEdgerulesItemSetEdgeRuleEnabledPostRequestBody) GetId()(*int64) {
    return m.id
}
// GetValue gets the Value property value. The Value property
// returns a *bool when successful
func (m *ItemEdgerulesItemSetEdgeRuleEnabledPostRequestBody) GetValue()(*bool) {
    return m.value
}
// Serialize serializes information the current object
func (m *ItemEdgerulesItemSetEdgeRuleEnabledPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("Id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("Value", m.GetValue())
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
func (m *ItemEdgerulesItemSetEdgeRuleEnabledPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the Id property value. The Id property
func (m *ItemEdgerulesItemSetEdgeRuleEnabledPostRequestBody) SetId(value *int64)() {
    m.id = value
}
// SetValue sets the Value property value. The Value property
func (m *ItemEdgerulesItemSetEdgeRuleEnabledPostRequestBody) SetValue(value *bool)() {
    m.value = value
}
type ItemEdgerulesItemSetEdgeRuleEnabledPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int64)
    GetValue()(*bool)
    SetId(value *int64)()
    SetValue(value *bool)()
}
