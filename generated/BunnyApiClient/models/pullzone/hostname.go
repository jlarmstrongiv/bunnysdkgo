package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Hostname struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Contains the Base64Url encoded certificate for the hostname
    certificate *string
    // Contains the Base64Url encoded certificate key for the hostname
    certificateKey *string
    // Determines if the Force SSL feature is enabled
    forceSSL *bool
    // Determines if the hostname has an SSL certificate configured
    hasCertificate *bool
    // The unique ID of the hostname
    id *int64
    // Determines if this is a system hostname controlled by bunny.net
    isSystemHostname *bool
    // The hostname value for the domain name
    value *string
}
// NewHostname instantiates a new Hostname and sets the default values.
func NewHostname()(*Hostname) {
    m := &Hostname{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateHostnameFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHostnameFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHostname(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Hostname) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCertificate gets the Certificate property value. Contains the Base64Url encoded certificate for the hostname
// returns a *string when successful
func (m *Hostname) GetCertificate()(*string) {
    return m.certificate
}
// GetCertificateKey gets the CertificateKey property value. Contains the Base64Url encoded certificate key for the hostname
// returns a *string when successful
func (m *Hostname) GetCertificateKey()(*string) {
    return m.certificateKey
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Hostname) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Certificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificate(val)
        }
        return nil
    }
    res["CertificateKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateKey(val)
        }
        return nil
    }
    res["ForceSSL"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForceSSL(val)
        }
        return nil
    }
    res["HasCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasCertificate(val)
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
    res["IsSystemHostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSystemHostname(val)
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
    return res
}
// GetForceSSL gets the ForceSSL property value. Determines if the Force SSL feature is enabled
// returns a *bool when successful
func (m *Hostname) GetForceSSL()(*bool) {
    return m.forceSSL
}
// GetHasCertificate gets the HasCertificate property value. Determines if the hostname has an SSL certificate configured
// returns a *bool when successful
func (m *Hostname) GetHasCertificate()(*bool) {
    return m.hasCertificate
}
// GetId gets the Id property value. The unique ID of the hostname
// returns a *int64 when successful
func (m *Hostname) GetId()(*int64) {
    return m.id
}
// GetIsSystemHostname gets the IsSystemHostname property value. Determines if this is a system hostname controlled by bunny.net
// returns a *bool when successful
func (m *Hostname) GetIsSystemHostname()(*bool) {
    return m.isSystemHostname
}
// GetValue gets the Value property value. The hostname value for the domain name
// returns a *string when successful
func (m *Hostname) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *Hostname) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("Certificate", m.GetCertificate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("CertificateKey", m.GetCertificateKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ForceSSL", m.GetForceSSL())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("HasCertificate", m.GetHasCertificate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("Id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("IsSystemHostname", m.GetIsSystemHostname())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Hostname) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCertificate sets the Certificate property value. Contains the Base64Url encoded certificate for the hostname
func (m *Hostname) SetCertificate(value *string)() {
    m.certificate = value
}
// SetCertificateKey sets the CertificateKey property value. Contains the Base64Url encoded certificate key for the hostname
func (m *Hostname) SetCertificateKey(value *string)() {
    m.certificateKey = value
}
// SetForceSSL sets the ForceSSL property value. Determines if the Force SSL feature is enabled
func (m *Hostname) SetForceSSL(value *bool)() {
    m.forceSSL = value
}
// SetHasCertificate sets the HasCertificate property value. Determines if the hostname has an SSL certificate configured
func (m *Hostname) SetHasCertificate(value *bool)() {
    m.hasCertificate = value
}
// SetId sets the Id property value. The unique ID of the hostname
func (m *Hostname) SetId(value *int64)() {
    m.id = value
}
// SetIsSystemHostname sets the IsSystemHostname property value. Determines if this is a system hostname controlled by bunny.net
func (m *Hostname) SetIsSystemHostname(value *bool)() {
    m.isSystemHostname = value
}
// SetValue sets the Value property value. The hostname value for the domain name
func (m *Hostname) SetValue(value *string)() {
    m.value = value
}
type Hostnameable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificate()(*string)
    GetCertificateKey()(*string)
    GetForceSSL()(*bool)
    GetHasCertificate()(*bool)
    GetId()(*int64)
    GetIsSystemHostname()(*bool)
    GetValue()(*string)
    SetCertificate(value *string)()
    SetCertificateKey(value *string)()
    SetForceSSL(value *bool)()
    SetHasCertificate(value *bool)()
    SetId(value *int64)()
    SetIsSystemHostname(value *bool)()
    SetValue(value *string)()
}
