package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CreateWafRateLimitRuleModel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The blockTime property
    blockTime *float64
    // The operatorType property
    operatorType *float64
    // The requestCount property
    requestCount *int32
    // The severityType property
    severityType *float64
    // The timeframe property
    timeframe *float64
    // The transformationTypes property
    transformationTypes []float64
    // The value property
    value *string
    // The variableTypes property
    variableTypes CreateWafRateLimitRuleModel_variableTypesable
}
// NewCreateWafRateLimitRuleModel instantiates a new CreateWafRateLimitRuleModel and sets the default values.
func NewCreateWafRateLimitRuleModel()(*CreateWafRateLimitRuleModel) {
    m := &CreateWafRateLimitRuleModel{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCreateWafRateLimitRuleModelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCreateWafRateLimitRuleModelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCreateWafRateLimitRuleModel(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CreateWafRateLimitRuleModel) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBlockTime gets the blockTime property value. The blockTime property
// returns a *float64 when successful
func (m *CreateWafRateLimitRuleModel) GetBlockTime()(*float64) {
    return m.blockTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CreateWafRateLimitRuleModel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["blockTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockTime(val)
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
    res["requestCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestCount(val)
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
    res["timeframe"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeframe(val)
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
        val, err := n.GetObjectValue(CreateCreateWafRateLimitRuleModel_variableTypesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVariableTypes(val.(CreateWafRateLimitRuleModel_variableTypesable))
        }
        return nil
    }
    return res
}
// GetOperatorType gets the operatorType property value. The operatorType property
// returns a *float64 when successful
func (m *CreateWafRateLimitRuleModel) GetOperatorType()(*float64) {
    return m.operatorType
}
// GetRequestCount gets the requestCount property value. The requestCount property
// returns a *int32 when successful
func (m *CreateWafRateLimitRuleModel) GetRequestCount()(*int32) {
    return m.requestCount
}
// GetSeverityType gets the severityType property value. The severityType property
// returns a *float64 when successful
func (m *CreateWafRateLimitRuleModel) GetSeverityType()(*float64) {
    return m.severityType
}
// GetTimeframe gets the timeframe property value. The timeframe property
// returns a *float64 when successful
func (m *CreateWafRateLimitRuleModel) GetTimeframe()(*float64) {
    return m.timeframe
}
// GetTransformationTypes gets the transformationTypes property value. The transformationTypes property
// returns a []float64 when successful
func (m *CreateWafRateLimitRuleModel) GetTransformationTypes()([]float64) {
    return m.transformationTypes
}
// GetValue gets the value property value. The value property
// returns a *string when successful
func (m *CreateWafRateLimitRuleModel) GetValue()(*string) {
    return m.value
}
// GetVariableTypes gets the variableTypes property value. The variableTypes property
// returns a CreateWafRateLimitRuleModel_variableTypesable when successful
func (m *CreateWafRateLimitRuleModel) GetVariableTypes()(CreateWafRateLimitRuleModel_variableTypesable) {
    return m.variableTypes
}
// Serialize serializes information the current object
func (m *CreateWafRateLimitRuleModel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("blockTime", m.GetBlockTime())
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
        err := writer.WriteInt32Value("requestCount", m.GetRequestCount())
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
    {
        err := writer.WriteFloat64Value("timeframe", m.GetTimeframe())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreateWafRateLimitRuleModel) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBlockTime sets the blockTime property value. The blockTime property
func (m *CreateWafRateLimitRuleModel) SetBlockTime(value *float64)() {
    m.blockTime = value
}
// SetOperatorType sets the operatorType property value. The operatorType property
func (m *CreateWafRateLimitRuleModel) SetOperatorType(value *float64)() {
    m.operatorType = value
}
// SetRequestCount sets the requestCount property value. The requestCount property
func (m *CreateWafRateLimitRuleModel) SetRequestCount(value *int32)() {
    m.requestCount = value
}
// SetSeverityType sets the severityType property value. The severityType property
func (m *CreateWafRateLimitRuleModel) SetSeverityType(value *float64)() {
    m.severityType = value
}
// SetTimeframe sets the timeframe property value. The timeframe property
func (m *CreateWafRateLimitRuleModel) SetTimeframe(value *float64)() {
    m.timeframe = value
}
// SetTransformationTypes sets the transformationTypes property value. The transformationTypes property
func (m *CreateWafRateLimitRuleModel) SetTransformationTypes(value []float64)() {
    m.transformationTypes = value
}
// SetValue sets the value property value. The value property
func (m *CreateWafRateLimitRuleModel) SetValue(value *string)() {
    m.value = value
}
// SetVariableTypes sets the variableTypes property value. The variableTypes property
func (m *CreateWafRateLimitRuleModel) SetVariableTypes(value CreateWafRateLimitRuleModel_variableTypesable)() {
    m.variableTypes = value
}
type CreateWafRateLimitRuleModelable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBlockTime()(*float64)
    GetOperatorType()(*float64)
    GetRequestCount()(*int32)
    GetSeverityType()(*float64)
    GetTimeframe()(*float64)
    GetTransformationTypes()([]float64)
    GetValue()(*string)
    GetVariableTypes()(CreateWafRateLimitRuleModel_variableTypesable)
    SetBlockTime(value *float64)()
    SetOperatorType(value *float64)()
    SetRequestCount(value *int32)()
    SetSeverityType(value *float64)()
    SetTimeframe(value *float64)()
    SetTransformationTypes(value []float64)()
    SetValue(value *string)()
    SetVariableTypes(value CreateWafRateLimitRuleModel_variableTypesable)()
}
