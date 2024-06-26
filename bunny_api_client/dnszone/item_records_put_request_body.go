package dnszone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemRecordsPutRequestBody struct {
    // The Accelerated property
    accelerated *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Comment property
    comment *string
    // The Disabled property
    disabled *bool
    // The EnviromentalVariables property
    enviromentalVariables []ItemRecordsPutRequestBody_EnviromentalVariablesable
    // The Flags property
    flags *int32
    // The GeolocationLatitude property
    geolocationLatitude *float64
    // The GeolocationLongitude property
    geolocationLongitude *float64
    // The LatencyZone property
    latencyZone *string
    // The MonitorType property
    monitorType *float64
    // The Name property
    name *string
    // The Port property
    port *int32
    // The Priority property
    priority *int32
    // The PullZoneId property
    pullZoneId *int64
    // The ScriptId property
    scriptId *int64
    // The SmartRoutingType property
    smartRoutingType *float64
    // The Tag property
    tag *string
    // The Ttl property
    ttl *float64
    // The Type property
    typeEscaped *float64
    // The Value property
    value *string
    // The Weight property
    weight *int32
}
// NewItemRecordsPutRequestBody instantiates a new ItemRecordsPutRequestBody and sets the default values.
func NewItemRecordsPutRequestBody()(*ItemRecordsPutRequestBody) {
    m := &ItemRecordsPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemRecordsPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemRecordsPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemRecordsPutRequestBody(), nil
}
// GetAccelerated gets the Accelerated property value. The Accelerated property
// returns a *bool when successful
func (m *ItemRecordsPutRequestBody) GetAccelerated()(*bool) {
    return m.accelerated
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemRecordsPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the Comment property value. The Comment property
// returns a *string when successful
func (m *ItemRecordsPutRequestBody) GetComment()(*string) {
    return m.comment
}
// GetDisabled gets the Disabled property value. The Disabled property
// returns a *bool when successful
func (m *ItemRecordsPutRequestBody) GetDisabled()(*bool) {
    return m.disabled
}
// GetEnviromentalVariables gets the EnviromentalVariables property value. The EnviromentalVariables property
// returns a []ItemRecordsPutRequestBody_EnviromentalVariablesable when successful
func (m *ItemRecordsPutRequestBody) GetEnviromentalVariables()([]ItemRecordsPutRequestBody_EnviromentalVariablesable) {
    return m.enviromentalVariables
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemRecordsPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Accelerated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccelerated(val)
        }
        return nil
    }
    res["Comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["Disabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisabled(val)
        }
        return nil
    }
    res["EnviromentalVariables"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemRecordsPutRequestBody_EnviromentalVariablesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemRecordsPutRequestBody_EnviromentalVariablesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ItemRecordsPutRequestBody_EnviromentalVariablesable)
                }
            }
            m.SetEnviromentalVariables(res)
        }
        return nil
    }
    res["Flags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlags(val)
        }
        return nil
    }
    res["GeolocationLatitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeolocationLatitude(val)
        }
        return nil
    }
    res["GeolocationLongitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeolocationLongitude(val)
        }
        return nil
    }
    res["LatencyZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatencyZone(val)
        }
        return nil
    }
    res["MonitorType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonitorType(val)
        }
        return nil
    }
    res["Name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["Port"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    res["Priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["PullZoneId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullZoneId(val)
        }
        return nil
    }
    res["ScriptId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScriptId(val)
        }
        return nil
    }
    res["SmartRoutingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmartRoutingType(val)
        }
        return nil
    }
    res["Tag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTag(val)
        }
        return nil
    }
    res["Ttl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTtl(val)
        }
        return nil
    }
    res["Type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    res["Value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    res["Weight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeight(val)
        }
        return nil
    }
    return res
}
// GetFlags gets the Flags property value. The Flags property
// returns a *int32 when successful
func (m *ItemRecordsPutRequestBody) GetFlags()(*int32) {
    return m.flags
}
// GetGeolocationLatitude gets the GeolocationLatitude property value. The GeolocationLatitude property
// returns a *float64 when successful
func (m *ItemRecordsPutRequestBody) GetGeolocationLatitude()(*float64) {
    return m.geolocationLatitude
}
// GetGeolocationLongitude gets the GeolocationLongitude property value. The GeolocationLongitude property
// returns a *float64 when successful
func (m *ItemRecordsPutRequestBody) GetGeolocationLongitude()(*float64) {
    return m.geolocationLongitude
}
// GetLatencyZone gets the LatencyZone property value. The LatencyZone property
// returns a *string when successful
func (m *ItemRecordsPutRequestBody) GetLatencyZone()(*string) {
    return m.latencyZone
}
// GetMonitorType gets the MonitorType property value. The MonitorType property
// returns a *float64 when successful
func (m *ItemRecordsPutRequestBody) GetMonitorType()(*float64) {
    return m.monitorType
}
// GetName gets the Name property value. The Name property
// returns a *string when successful
func (m *ItemRecordsPutRequestBody) GetName()(*string) {
    return m.name
}
// GetPort gets the Port property value. The Port property
// returns a *int32 when successful
func (m *ItemRecordsPutRequestBody) GetPort()(*int32) {
    return m.port
}
// GetPriority gets the Priority property value. The Priority property
// returns a *int32 when successful
func (m *ItemRecordsPutRequestBody) GetPriority()(*int32) {
    return m.priority
}
// GetPullZoneId gets the PullZoneId property value. The PullZoneId property
// returns a *int64 when successful
func (m *ItemRecordsPutRequestBody) GetPullZoneId()(*int64) {
    return m.pullZoneId
}
// GetScriptId gets the ScriptId property value. The ScriptId property
// returns a *int64 when successful
func (m *ItemRecordsPutRequestBody) GetScriptId()(*int64) {
    return m.scriptId
}
// GetSmartRoutingType gets the SmartRoutingType property value. The SmartRoutingType property
// returns a *float64 when successful
func (m *ItemRecordsPutRequestBody) GetSmartRoutingType()(*float64) {
    return m.smartRoutingType
}
// GetTag gets the Tag property value. The Tag property
// returns a *string when successful
func (m *ItemRecordsPutRequestBody) GetTag()(*string) {
    return m.tag
}
// GetTtl gets the Ttl property value. The Ttl property
// returns a *float64 when successful
func (m *ItemRecordsPutRequestBody) GetTtl()(*float64) {
    return m.ttl
}
// GetTypeEscaped gets the Type property value. The Type property
// returns a *float64 when successful
func (m *ItemRecordsPutRequestBody) GetTypeEscaped()(*float64) {
    return m.typeEscaped
}
// GetValue gets the Value property value. The Value property
// returns a *string when successful
func (m *ItemRecordsPutRequestBody) GetValue()(*string) {
    return m.value
}
// GetWeight gets the Weight property value. The Weight property
// returns a *int32 when successful
func (m *ItemRecordsPutRequestBody) GetWeight()(*int32) {
    return m.weight
}
// Serialize serializes information the current object
func (m *ItemRecordsPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("Accelerated", m.GetAccelerated())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("Disabled", m.GetDisabled())
        if err != nil {
            return err
        }
    }
    if m.GetEnviromentalVariables() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEnviromentalVariables()))
        for i, v := range m.GetEnviromentalVariables() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("EnviromentalVariables", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("Flags", m.GetFlags())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("GeolocationLatitude", m.GetGeolocationLatitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("GeolocationLongitude", m.GetGeolocationLongitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("LatencyZone", m.GetLatencyZone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("MonitorType", m.GetMonitorType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("Port", m.GetPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("Priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("PullZoneId", m.GetPullZoneId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("ScriptId", m.GetScriptId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("SmartRoutingType", m.GetSmartRoutingType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Tag", m.GetTag())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("Ttl", m.GetTtl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("Type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Value", m.GetValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("Weight", m.GetWeight())
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
// SetAccelerated sets the Accelerated property value. The Accelerated property
func (m *ItemRecordsPutRequestBody) SetAccelerated(value *bool)() {
    m.accelerated = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemRecordsPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the Comment property value. The Comment property
func (m *ItemRecordsPutRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetDisabled sets the Disabled property value. The Disabled property
func (m *ItemRecordsPutRequestBody) SetDisabled(value *bool)() {
    m.disabled = value
}
// SetEnviromentalVariables sets the EnviromentalVariables property value. The EnviromentalVariables property
func (m *ItemRecordsPutRequestBody) SetEnviromentalVariables(value []ItemRecordsPutRequestBody_EnviromentalVariablesable)() {
    m.enviromentalVariables = value
}
// SetFlags sets the Flags property value. The Flags property
func (m *ItemRecordsPutRequestBody) SetFlags(value *int32)() {
    m.flags = value
}
// SetGeolocationLatitude sets the GeolocationLatitude property value. The GeolocationLatitude property
func (m *ItemRecordsPutRequestBody) SetGeolocationLatitude(value *float64)() {
    m.geolocationLatitude = value
}
// SetGeolocationLongitude sets the GeolocationLongitude property value. The GeolocationLongitude property
func (m *ItemRecordsPutRequestBody) SetGeolocationLongitude(value *float64)() {
    m.geolocationLongitude = value
}
// SetLatencyZone sets the LatencyZone property value. The LatencyZone property
func (m *ItemRecordsPutRequestBody) SetLatencyZone(value *string)() {
    m.latencyZone = value
}
// SetMonitorType sets the MonitorType property value. The MonitorType property
func (m *ItemRecordsPutRequestBody) SetMonitorType(value *float64)() {
    m.monitorType = value
}
// SetName sets the Name property value. The Name property
func (m *ItemRecordsPutRequestBody) SetName(value *string)() {
    m.name = value
}
// SetPort sets the Port property value. The Port property
func (m *ItemRecordsPutRequestBody) SetPort(value *int32)() {
    m.port = value
}
// SetPriority sets the Priority property value. The Priority property
func (m *ItemRecordsPutRequestBody) SetPriority(value *int32)() {
    m.priority = value
}
// SetPullZoneId sets the PullZoneId property value. The PullZoneId property
func (m *ItemRecordsPutRequestBody) SetPullZoneId(value *int64)() {
    m.pullZoneId = value
}
// SetScriptId sets the ScriptId property value. The ScriptId property
func (m *ItemRecordsPutRequestBody) SetScriptId(value *int64)() {
    m.scriptId = value
}
// SetSmartRoutingType sets the SmartRoutingType property value. The SmartRoutingType property
func (m *ItemRecordsPutRequestBody) SetSmartRoutingType(value *float64)() {
    m.smartRoutingType = value
}
// SetTag sets the Tag property value. The Tag property
func (m *ItemRecordsPutRequestBody) SetTag(value *string)() {
    m.tag = value
}
// SetTtl sets the Ttl property value. The Ttl property
func (m *ItemRecordsPutRequestBody) SetTtl(value *float64)() {
    m.ttl = value
}
// SetTypeEscaped sets the Type property value. The Type property
func (m *ItemRecordsPutRequestBody) SetTypeEscaped(value *float64)() {
    m.typeEscaped = value
}
// SetValue sets the Value property value. The Value property
func (m *ItemRecordsPutRequestBody) SetValue(value *string)() {
    m.value = value
}
// SetWeight sets the Weight property value. The Weight property
func (m *ItemRecordsPutRequestBody) SetWeight(value *int32)() {
    m.weight = value
}
type ItemRecordsPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccelerated()(*bool)
    GetComment()(*string)
    GetDisabled()(*bool)
    GetEnviromentalVariables()([]ItemRecordsPutRequestBody_EnviromentalVariablesable)
    GetFlags()(*int32)
    GetGeolocationLatitude()(*float64)
    GetGeolocationLongitude()(*float64)
    GetLatencyZone()(*string)
    GetMonitorType()(*float64)
    GetName()(*string)
    GetPort()(*int32)
    GetPriority()(*int32)
    GetPullZoneId()(*int64)
    GetScriptId()(*int64)
    GetSmartRoutingType()(*float64)
    GetTag()(*string)
    GetTtl()(*float64)
    GetTypeEscaped()(*float64)
    GetValue()(*string)
    GetWeight()(*int32)
    SetAccelerated(value *bool)()
    SetComment(value *string)()
    SetDisabled(value *bool)()
    SetEnviromentalVariables(value []ItemRecordsPutRequestBody_EnviromentalVariablesable)()
    SetFlags(value *int32)()
    SetGeolocationLatitude(value *float64)()
    SetGeolocationLongitude(value *float64)()
    SetLatencyZone(value *string)()
    SetMonitorType(value *float64)()
    SetName(value *string)()
    SetPort(value *int32)()
    SetPriority(value *int32)()
    SetPullZoneId(value *int64)()
    SetScriptId(value *int64)()
    SetSmartRoutingType(value *float64)()
    SetTag(value *string)()
    SetTtl(value *float64)()
    SetTypeEscaped(value *float64)()
    SetValue(value *string)()
    SetWeight(value *int32)()
}
