package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CreateCustomWafRuleModel struct {
    // The actionType property
    actionType *float64
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The operatorType property
    operatorType *float64
    // The severityType property
    severityType *float64
    // The transformationTypes property
    transformationTypes []float64
    // The value property
    value *string
    // The variableTypes property
    variableTypes CreateCustomWafRuleModel_variableTypesable
}
// NewCreateCustomWafRuleModel instantiates a new CreateCustomWafRuleModel and sets the default values.
func NewCreateCustomWafRuleModel()(*CreateCustomWafRuleModel) {
    m := &CreateCustomWafRuleModel{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCreateCustomWafRuleModelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCreateCustomWafRuleModelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCreateCustomWafRuleModel(), nil
}
// GetActionType gets the actionType property value. The actionType property
// returns a *float64 when successful
func (m *CreateCustomWafRuleModel) GetActionType()(*float64) {
    return m.actionType
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CreateCustomWafRuleModel) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CreateCustomWafRuleModel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionType(val)
        }
        return nil
    }
    res["operatorType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatorType(val)
        }
        return nil
    }
    res["severityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverityType(val)
        }
        return nil
    }
    res["transformationTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("float64")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]float64, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*float64))
                }
            }
            m.SetTransformationTypes(res)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    res["variableTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCreateCustomWafRuleModel_variableTypesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVariableTypes(val.(CreateCustomWafRuleModel_variableTypesable))
        }
        return nil
    }
    return res
}
// GetOperatorType gets the operatorType property value. The operatorType property
// returns a *float64 when successful
func (m *CreateCustomWafRuleModel) GetOperatorType()(*float64) {
    return m.operatorType
}
// GetSeverityType gets the severityType property value. The severityType property
// returns a *float64 when successful
func (m *CreateCustomWafRuleModel) GetSeverityType()(*float64) {
    return m.severityType
}
// GetTransformationTypes gets the transformationTypes property value. The transformationTypes property
// returns a []float64 when successful
func (m *CreateCustomWafRuleModel) GetTransformationTypes()([]float64) {
    return m.transformationTypes
}
// GetValue gets the value property value. The value property
// returns a *string when successful
func (m *CreateCustomWafRuleModel) GetValue()(*string) {
    return m.value
}
// GetVariableTypes gets the variableTypes property value. The variableTypes property
// returns a CreateCustomWafRuleModel_variableTypesable when successful
func (m *CreateCustomWafRuleModel) GetVariableTypes()(CreateCustomWafRuleModel_variableTypesable) {
    return m.variableTypes
}
// Serialize serializes information the current object
func (m *CreateCustomWafRuleModel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("actionType", m.GetActionType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("operatorType", m.GetOperatorType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("severityType", m.GetSeverityType())
        if err != nil {
            return err
        }
    }
    if m.GetTransformationTypes() != nil {
        err := writer.WriteCollectionOfFloat64Values("transformationTypes", m.GetTransformationTypes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("variableTypes", m.GetVariableTypes())
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
// SetActionType sets the actionType property value. The actionType property
func (m *CreateCustomWafRuleModel) SetActionType(value *float64)() {
    m.actionType = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreateCustomWafRuleModel) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOperatorType sets the operatorType property value. The operatorType property
func (m *CreateCustomWafRuleModel) SetOperatorType(value *float64)() {
    m.operatorType = value
}
// SetSeverityType sets the severityType property value. The severityType property
func (m *CreateCustomWafRuleModel) SetSeverityType(value *float64)() {
    m.severityType = value
}
// SetTransformationTypes sets the transformationTypes property value. The transformationTypes property
func (m *CreateCustomWafRuleModel) SetTransformationTypes(value []float64)() {
    m.transformationTypes = value
}
// SetValue sets the value property value. The value property
func (m *CreateCustomWafRuleModel) SetValue(value *string)() {
    m.value = value
}
// SetVariableTypes sets the variableTypes property value. The variableTypes property
func (m *CreateCustomWafRuleModel) SetVariableTypes(value CreateCustomWafRuleModel_variableTypesable)() {
    m.variableTypes = value
}
type CreateCustomWafRuleModelable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionType()(*float64)
    GetOperatorType()(*float64)
    GetSeverityType()(*float64)
    GetTransformationTypes()([]float64)
    GetValue()(*string)
    GetVariableTypes()(CreateCustomWafRuleModel_variableTypesable)
    SetActionType(value *float64)()
    SetOperatorType(value *float64)()
    SetSeverityType(value *float64)()
    SetTransformationTypes(value []float64)()
    SetValue(value *string)()
    SetVariableTypes(value CreateCustomWafRuleModel_variableTypesable)()
}
