package dnszone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DnsRecord struct {
    // The Accelerated property
    accelerated *bool
    // The AcceleratedPullZoneId property
    acceleratedPullZoneId *int64
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Comment property
    comment *string
    // The Disabled property
    disabled *bool
    // The EnviromentalVariables property
    enviromentalVariables []DnsRecord_EnviromentalVariablesable
    // The Flags property
    flags *int32
    // The GeolocationInfo property
    geolocationInfo DnsRecord_GeolocationInfoable
    // The GeolocationLatitude property
    geolocationLatitude *float64
    // The GeolocationLongitude property
    geolocationLongitude *float64
    // The Id property
    id *int64
    // The IPGeoLocationInfo property
    iPGeoLocationInfo DnsRecord_IPGeoLocationInfoable
    // The LatencyZone property
    latencyZone *string
    // The LinkName property
    linkName *string
    // The MonitorStatus property
    monitorStatus *float64
    // The MonitorType property
    monitorType *float64
    // The Name property
    name *string
    // The Port property
    port *int32
    // The Priority property
    priority *int32
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
// NewDnsRecord instantiates a new DnsRecord and sets the default values.
func NewDnsRecord()(*DnsRecord) {
    m := &DnsRecord{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDnsRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDnsRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDnsRecord(), nil
}
// GetAccelerated gets the Accelerated property value. The Accelerated property
// returns a *bool when successful
func (m *DnsRecord) GetAccelerated()(*bool) {
    return m.accelerated
}
// GetAcceleratedPullZoneId gets the AcceleratedPullZoneId property value. The AcceleratedPullZoneId property
// returns a *int64 when successful
func (m *DnsRecord) GetAcceleratedPullZoneId()(*int64) {
    return m.acceleratedPullZoneId
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DnsRecord) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the Comment property value. The Comment property
// returns a *string when successful
func (m *DnsRecord) GetComment()(*string) {
    return m.comment
}
// GetDisabled gets the Disabled property value. The Disabled property
// returns a *bool when successful
func (m *DnsRecord) GetDisabled()(*bool) {
    return m.disabled
}
// GetEnviromentalVariables gets the EnviromentalVariables property value. The EnviromentalVariables property
// returns a []DnsRecord_EnviromentalVariablesable when successful
func (m *DnsRecord) GetEnviromentalVariables()([]DnsRecord_EnviromentalVariablesable) {
    return m.enviromentalVariables
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DnsRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["AcceleratedPullZoneId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceleratedPullZoneId(val)
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
        val, err := n.GetCollectionOfObjectValues(CreateDnsRecord_EnviromentalVariablesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DnsRecord_EnviromentalVariablesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DnsRecord_EnviromentalVariablesable)
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
    res["GeolocationInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDnsRecord_GeolocationInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeolocationInfo(val.(DnsRecord_GeolocationInfoable))
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
    res["IPGeoLocationInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDnsRecord_IPGeoLocationInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIPGeoLocationInfo(val.(DnsRecord_IPGeoLocationInfoable))
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
    res["LinkName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinkName(val)
        }
        return nil
    }
    res["MonitorStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonitorStatus(val)
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
func (m *DnsRecord) GetFlags()(*int32) {
    return m.flags
}
// GetGeolocationInfo gets the GeolocationInfo property value. The GeolocationInfo property
// returns a DnsRecord_GeolocationInfoable when successful
func (m *DnsRecord) GetGeolocationInfo()(DnsRecord_GeolocationInfoable) {
    return m.geolocationInfo
}
// GetGeolocationLatitude gets the GeolocationLatitude property value. The GeolocationLatitude property
// returns a *float64 when successful
func (m *DnsRecord) GetGeolocationLatitude()(*float64) {
    return m.geolocationLatitude
}
// GetGeolocationLongitude gets the GeolocationLongitude property value. The GeolocationLongitude property
// returns a *float64 when successful
func (m *DnsRecord) GetGeolocationLongitude()(*float64) {
    return m.geolocationLongitude
}
// GetId gets the Id property value. The Id property
// returns a *int64 when successful
func (m *DnsRecord) GetId()(*int64) {
    return m.id
}
// GetIPGeoLocationInfo gets the IPGeoLocationInfo property value. The IPGeoLocationInfo property
// returns a DnsRecord_IPGeoLocationInfoable when successful
func (m *DnsRecord) GetIPGeoLocationInfo()(DnsRecord_IPGeoLocationInfoable) {
    return m.iPGeoLocationInfo
}
// GetLatencyZone gets the LatencyZone property value. The LatencyZone property
// returns a *string when successful
func (m *DnsRecord) GetLatencyZone()(*string) {
    return m.latencyZone
}
// GetLinkName gets the LinkName property value. The LinkName property
// returns a *string when successful
func (m *DnsRecord) GetLinkName()(*string) {
    return m.linkName
}
// GetMonitorStatus gets the MonitorStatus property value. The MonitorStatus property
// returns a *float64 when successful
func (m *DnsRecord) GetMonitorStatus()(*float64) {
    return m.monitorStatus
}
// GetMonitorType gets the MonitorType property value. The MonitorType property
// returns a *float64 when successful
func (m *DnsRecord) GetMonitorType()(*float64) {
    return m.monitorType
}
// GetName gets the Name property value. The Name property
// returns a *string when successful
func (m *DnsRecord) GetName()(*string) {
    return m.name
}
// GetPort gets the Port property value. The Port property
// returns a *int32 when successful
func (m *DnsRecord) GetPort()(*int32) {
    return m.port
}
// GetPriority gets the Priority property value. The Priority property
// returns a *int32 when successful
func (m *DnsRecord) GetPriority()(*int32) {
    return m.priority
}
// GetSmartRoutingType gets the SmartRoutingType property value. The SmartRoutingType property
// returns a *float64 when successful
func (m *DnsRecord) GetSmartRoutingType()(*float64) {
    return m.smartRoutingType
}
// GetTag gets the Tag property value. The Tag property
// returns a *string when successful
func (m *DnsRecord) GetTag()(*string) {
    return m.tag
}
// GetTtl gets the Ttl property value. The Ttl property
// returns a *float64 when successful
func (m *DnsRecord) GetTtl()(*float64) {
    return m.ttl
}
// GetTypeEscaped gets the Type property value. The Type property
// returns a *float64 when successful
func (m *DnsRecord) GetTypeEscaped()(*float64) {
    return m.typeEscaped
}
// GetValue gets the Value property value. The Value property
// returns a *string when successful
func (m *DnsRecord) GetValue()(*string) {
    return m.value
}
// GetWeight gets the Weight property value. The Weight property
// returns a *int32 when successful
func (m *DnsRecord) GetWeight()(*int32) {
    return m.weight
}
// Serialize serializes information the current object
func (m *DnsRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DnsRecord) SetAccelerated(value *bool)() {
    m.accelerated = value
}
// SetAcceleratedPullZoneId sets the AcceleratedPullZoneId property value. The AcceleratedPullZoneId property
func (m *DnsRecord) SetAcceleratedPullZoneId(value *int64)() {
    m.acceleratedPullZoneId = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DnsRecord) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the Comment property value. The Comment property
func (m *DnsRecord) SetComment(value *string)() {
    m.comment = value
}
// SetDisabled sets the Disabled property value. The Disabled property
func (m *DnsRecord) SetDisabled(value *bool)() {
    m.disabled = value
}
// SetEnviromentalVariables sets the EnviromentalVariables property value. The EnviromentalVariables property
func (m *DnsRecord) SetEnviromentalVariables(value []DnsRecord_EnviromentalVariablesable)() {
    m.enviromentalVariables = value
}
// SetFlags sets the Flags property value. The Flags property
func (m *DnsRecord) SetFlags(value *int32)() {
    m.flags = value
}
// SetGeolocationInfo sets the GeolocationInfo property value. The GeolocationInfo property
func (m *DnsRecord) SetGeolocationInfo(value DnsRecord_GeolocationInfoable)() {
    m.geolocationInfo = value
}
// SetGeolocationLatitude sets the GeolocationLatitude property value. The GeolocationLatitude property
func (m *DnsRecord) SetGeolocationLatitude(value *float64)() {
    m.geolocationLatitude = value
}
// SetGeolocationLongitude sets the GeolocationLongitude property value. The GeolocationLongitude property
func (m *DnsRecord) SetGeolocationLongitude(value *float64)() {
    m.geolocationLongitude = value
}
// SetId sets the Id property value. The Id property
func (m *DnsRecord) SetId(value *int64)() {
    m.id = value
}
// SetIPGeoLocationInfo sets the IPGeoLocationInfo property value. The IPGeoLocationInfo property
func (m *DnsRecord) SetIPGeoLocationInfo(value DnsRecord_IPGeoLocationInfoable)() {
    m.iPGeoLocationInfo = value
}
// SetLatencyZone sets the LatencyZone property value. The LatencyZone property
func (m *DnsRecord) SetLatencyZone(value *string)() {
    m.latencyZone = value
}
// SetLinkName sets the LinkName property value. The LinkName property
func (m *DnsRecord) SetLinkName(value *string)() {
    m.linkName = value
}
// SetMonitorStatus sets the MonitorStatus property value. The MonitorStatus property
func (m *DnsRecord) SetMonitorStatus(value *float64)() {
    m.monitorStatus = value
}
// SetMonitorType sets the MonitorType property value. The MonitorType property
func (m *DnsRecord) SetMonitorType(value *float64)() {
    m.monitorType = value
}
// SetName sets the Name property value. The Name property
func (m *DnsRecord) SetName(value *string)() {
    m.name = value
}
// SetPort sets the Port property value. The Port property
func (m *DnsRecord) SetPort(value *int32)() {
    m.port = value
}
// SetPriority sets the Priority property value. The Priority property
func (m *DnsRecord) SetPriority(value *int32)() {
    m.priority = value
}
// SetSmartRoutingType sets the SmartRoutingType property value. The SmartRoutingType property
func (m *DnsRecord) SetSmartRoutingType(value *float64)() {
    m.smartRoutingType = value
}
// SetTag sets the Tag property value. The Tag property
func (m *DnsRecord) SetTag(value *string)() {
    m.tag = value
}
// SetTtl sets the Ttl property value. The Ttl property
func (m *DnsRecord) SetTtl(value *float64)() {
    m.ttl = value
}
// SetTypeEscaped sets the Type property value. The Type property
func (m *DnsRecord) SetTypeEscaped(value *float64)() {
    m.typeEscaped = value
}
// SetValue sets the Value property value. The Value property
func (m *DnsRecord) SetValue(value *string)() {
    m.value = value
}
// SetWeight sets the Weight property value. The Weight property
func (m *DnsRecord) SetWeight(value *int32)() {
    m.weight = value
}
type DnsRecordable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccelerated()(*bool)
    GetAcceleratedPullZoneId()(*int64)
    GetComment()(*string)
    GetDisabled()(*bool)
    GetEnviromentalVariables()([]DnsRecord_EnviromentalVariablesable)
    GetFlags()(*int32)
    GetGeolocationInfo()(DnsRecord_GeolocationInfoable)
    GetGeolocationLatitude()(*float64)
    GetGeolocationLongitude()(*float64)
    GetId()(*int64)
    GetIPGeoLocationInfo()(DnsRecord_IPGeoLocationInfoable)
    GetLatencyZone()(*string)
    GetLinkName()(*string)
    GetMonitorStatus()(*float64)
    GetMonitorType()(*float64)
    GetName()(*string)
    GetPort()(*int32)
    GetPriority()(*int32)
    GetSmartRoutingType()(*float64)
    GetTag()(*string)
    GetTtl()(*float64)
    GetTypeEscaped()(*float64)
    GetValue()(*string)
    GetWeight()(*int32)
    SetAccelerated(value *bool)()
    SetAcceleratedPullZoneId(value *int64)()
    SetComment(value *string)()
    SetDisabled(value *bool)()
    SetEnviromentalVariables(value []DnsRecord_EnviromentalVariablesable)()
    SetFlags(value *int32)()
    SetGeolocationInfo(value DnsRecord_GeolocationInfoable)()
    SetGeolocationLatitude(value *float64)()
    SetGeolocationLongitude(value *float64)()
    SetId(value *int64)()
    SetIPGeoLocationInfo(value DnsRecord_IPGeoLocationInfoable)()
    SetLatencyZone(value *string)()
    SetLinkName(value *string)()
    SetMonitorStatus(value *float64)()
    SetMonitorType(value *float64)()
    SetName(value *string)()
    SetPort(value *int32)()
    SetPriority(value *int32)()
    SetSmartRoutingType(value *float64)()
    SetTag(value *string)()
    SetTtl(value *float64)()
    SetTypeEscaped(value *float64)()
    SetValue(value *string)()
    SetWeight(value *int32)()
}
