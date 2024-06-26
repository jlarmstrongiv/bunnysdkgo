package dnszone

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DnsZone struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The CustomNameserversEnabled property
    customNameserversEnabled *bool
    // The DateCreated property
    dateCreated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The DateModified property
    dateModified *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Domain property
    domain *string
    // The Id property
    id *int64
    // The LogAnonymizationType property
    logAnonymizationType *float64
    // The LoggingEnabled property
    loggingEnabled *bool
    // Determines if the log anonymization should be enabled
    loggingIPAnonymizationEnabled *bool
    // The Nameserver1 property
    nameserver1 *string
    // The Nameserver2 property
    nameserver2 *string
    // The NameserversDetected property
    nameserversDetected *bool
    // The NameserversNextCheck property
    nameserversNextCheck *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Records property
    records []DnsRecordable
    // The SoaEmail property
    soaEmail *string
}
// NewDnsZone instantiates a new DnsZone and sets the default values.
func NewDnsZone()(*DnsZone) {
    m := &DnsZone{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDnsZoneFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDnsZoneFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDnsZone(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DnsZone) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustomNameserversEnabled gets the CustomNameserversEnabled property value. The CustomNameserversEnabled property
// returns a *bool when successful
func (m *DnsZone) GetCustomNameserversEnabled()(*bool) {
    return m.customNameserversEnabled
}
// GetDateCreated gets the DateCreated property value. The DateCreated property
// returns a *Time when successful
func (m *DnsZone) GetDateCreated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dateCreated
}
// GetDateModified gets the DateModified property value. The DateModified property
// returns a *Time when successful
func (m *DnsZone) GetDateModified()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dateModified
}
// GetDomain gets the Domain property value. The Domain property
// returns a *string when successful
func (m *DnsZone) GetDomain()(*string) {
    return m.domain
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DnsZone) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["CustomNameserversEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomNameserversEnabled(val)
        }
        return nil
    }
    res["DateCreated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateCreated(val)
        }
        return nil
    }
    res["DateModified"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateModified(val)
        }
        return nil
    }
    res["Domain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomain(val)
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
    res["LogAnonymizationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogAnonymizationType(val)
        }
        return nil
    }
    res["LoggingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoggingEnabled(val)
        }
        return nil
    }
    res["LoggingIPAnonymizationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoggingIPAnonymizationEnabled(val)
        }
        return nil
    }
    res["Nameserver1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNameserver1(val)
        }
        return nil
    }
    res["Nameserver2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNameserver2(val)
        }
        return nil
    }
    res["NameserversDetected"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNameserversDetected(val)
        }
        return nil
    }
    res["NameserversNextCheck"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNameserversNextCheck(val)
        }
        return nil
    }
    res["Records"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDnsRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DnsRecordable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DnsRecordable)
                }
            }
            m.SetRecords(res)
        }
        return nil
    }
    res["SoaEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoaEmail(val)
        }
        return nil
    }
    return res
}
// GetId gets the Id property value. The Id property
// returns a *int64 when successful
func (m *DnsZone) GetId()(*int64) {
    return m.id
}
// GetLogAnonymizationType gets the LogAnonymizationType property value. The LogAnonymizationType property
// returns a *float64 when successful
func (m *DnsZone) GetLogAnonymizationType()(*float64) {
    return m.logAnonymizationType
}
// GetLoggingEnabled gets the LoggingEnabled property value. The LoggingEnabled property
// returns a *bool when successful
func (m *DnsZone) GetLoggingEnabled()(*bool) {
    return m.loggingEnabled
}
// GetLoggingIPAnonymizationEnabled gets the LoggingIPAnonymizationEnabled property value. Determines if the log anonymization should be enabled
// returns a *bool when successful
func (m *DnsZone) GetLoggingIPAnonymizationEnabled()(*bool) {
    return m.loggingIPAnonymizationEnabled
}
// GetNameserver1 gets the Nameserver1 property value. The Nameserver1 property
// returns a *string when successful
func (m *DnsZone) GetNameserver1()(*string) {
    return m.nameserver1
}
// GetNameserver2 gets the Nameserver2 property value. The Nameserver2 property
// returns a *string when successful
func (m *DnsZone) GetNameserver2()(*string) {
    return m.nameserver2
}
// GetNameserversDetected gets the NameserversDetected property value. The NameserversDetected property
// returns a *bool when successful
func (m *DnsZone) GetNameserversDetected()(*bool) {
    return m.nameserversDetected
}
// GetNameserversNextCheck gets the NameserversNextCheck property value. The NameserversNextCheck property
// returns a *Time when successful
func (m *DnsZone) GetNameserversNextCheck()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.nameserversNextCheck
}
// GetRecords gets the Records property value. The Records property
// returns a []DnsRecordable when successful
func (m *DnsZone) GetRecords()([]DnsRecordable) {
    return m.records
}
// GetSoaEmail gets the SoaEmail property value. The SoaEmail property
// returns a *string when successful
func (m *DnsZone) GetSoaEmail()(*string) {
    return m.soaEmail
}
// Serialize serializes information the current object
func (m *DnsZone) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("CustomNameserversEnabled", m.GetCustomNameserversEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("LogAnonymizationType", m.GetLogAnonymizationType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("LoggingEnabled", m.GetLoggingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("LoggingIPAnonymizationEnabled", m.GetLoggingIPAnonymizationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Nameserver1", m.GetNameserver1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Nameserver2", m.GetNameserver2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("SoaEmail", m.GetSoaEmail())
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
func (m *DnsZone) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustomNameserversEnabled sets the CustomNameserversEnabled property value. The CustomNameserversEnabled property
func (m *DnsZone) SetCustomNameserversEnabled(value *bool)() {
    m.customNameserversEnabled = value
}
// SetDateCreated sets the DateCreated property value. The DateCreated property
func (m *DnsZone) SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateCreated = value
}
// SetDateModified sets the DateModified property value. The DateModified property
func (m *DnsZone) SetDateModified(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateModified = value
}
// SetDomain sets the Domain property value. The Domain property
func (m *DnsZone) SetDomain(value *string)() {
    m.domain = value
}
// SetId sets the Id property value. The Id property
func (m *DnsZone) SetId(value *int64)() {
    m.id = value
}
// SetLogAnonymizationType sets the LogAnonymizationType property value. The LogAnonymizationType property
func (m *DnsZone) SetLogAnonymizationType(value *float64)() {
    m.logAnonymizationType = value
}
// SetLoggingEnabled sets the LoggingEnabled property value. The LoggingEnabled property
func (m *DnsZone) SetLoggingEnabled(value *bool)() {
    m.loggingEnabled = value
}
// SetLoggingIPAnonymizationEnabled sets the LoggingIPAnonymizationEnabled property value. Determines if the log anonymization should be enabled
func (m *DnsZone) SetLoggingIPAnonymizationEnabled(value *bool)() {
    m.loggingIPAnonymizationEnabled = value
}
// SetNameserver1 sets the Nameserver1 property value. The Nameserver1 property
func (m *DnsZone) SetNameserver1(value *string)() {
    m.nameserver1 = value
}
// SetNameserver2 sets the Nameserver2 property value. The Nameserver2 property
func (m *DnsZone) SetNameserver2(value *string)() {
    m.nameserver2 = value
}
// SetNameserversDetected sets the NameserversDetected property value. The NameserversDetected property
func (m *DnsZone) SetNameserversDetected(value *bool)() {
    m.nameserversDetected = value
}
// SetNameserversNextCheck sets the NameserversNextCheck property value. The NameserversNextCheck property
func (m *DnsZone) SetNameserversNextCheck(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.nameserversNextCheck = value
}
// SetRecords sets the Records property value. The Records property
func (m *DnsZone) SetRecords(value []DnsRecordable)() {
    m.records = value
}
// SetSoaEmail sets the SoaEmail property value. The SoaEmail property
func (m *DnsZone) SetSoaEmail(value *string)() {
    m.soaEmail = value
}
type DnsZoneable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomNameserversEnabled()(*bool)
    GetDateCreated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDateModified()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDomain()(*string)
    GetId()(*int64)
    GetLogAnonymizationType()(*float64)
    GetLoggingEnabled()(*bool)
    GetLoggingIPAnonymizationEnabled()(*bool)
    GetNameserver1()(*string)
    GetNameserver2()(*string)
    GetNameserversDetected()(*bool)
    GetNameserversNextCheck()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRecords()([]DnsRecordable)
    GetSoaEmail()(*string)
    SetCustomNameserversEnabled(value *bool)()
    SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDateModified(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDomain(value *string)()
    SetId(value *int64)()
    SetLogAnonymizationType(value *float64)()
    SetLoggingEnabled(value *bool)()
    SetLoggingIPAnonymizationEnabled(value *bool)()
    SetNameserver1(value *string)()
    SetNameserver2(value *string)()
    SetNameserversDetected(value *bool)()
    SetNameserversNextCheck(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRecords(value []DnsRecordable)()
    SetSoaEmail(value *string)()
}
