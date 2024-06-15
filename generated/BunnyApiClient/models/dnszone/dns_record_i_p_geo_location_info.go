package dnszone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DnsRecord_IPGeoLocationInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ASN property
    aSN *int64
    // The City property
    city *string
    // The Country property
    country *string
    // The CountryCode property
    countryCode *string
    // The OrganizationName property
    organizationName *string
}
// NewDnsRecord_IPGeoLocationInfo instantiates a new DnsRecord_IPGeoLocationInfo and sets the default values.
func NewDnsRecord_IPGeoLocationInfo()(*DnsRecord_IPGeoLocationInfo) {
    m := &DnsRecord_IPGeoLocationInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDnsRecord_IPGeoLocationInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDnsRecord_IPGeoLocationInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDnsRecord_IPGeoLocationInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DnsRecord_IPGeoLocationInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetASN gets the ASN property value. The ASN property
// returns a *int64 when successful
func (m *DnsRecord_IPGeoLocationInfo) GetASN()(*int64) {
    return m.aSN
}
// GetCity gets the City property value. The City property
// returns a *string when successful
func (m *DnsRecord_IPGeoLocationInfo) GetCity()(*string) {
    return m.city
}
// GetCountry gets the Country property value. The Country property
// returns a *string when successful
func (m *DnsRecord_IPGeoLocationInfo) GetCountry()(*string) {
    return m.country
}
// GetCountryCode gets the CountryCode property value. The CountryCode property
// returns a *string when successful
func (m *DnsRecord_IPGeoLocationInfo) GetCountryCode()(*string) {
    return m.countryCode
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DnsRecord_IPGeoLocationInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ASN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetASN(val)
        }
        return nil
    }
    res["City"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["Country"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountry(val)
        }
        return nil
    }
    res["CountryCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryCode(val)
        }
        return nil
    }
    res["OrganizationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationName(val)
        }
        return nil
    }
    return res
}
// GetOrganizationName gets the OrganizationName property value. The OrganizationName property
// returns a *string when successful
func (m *DnsRecord_IPGeoLocationInfo) GetOrganizationName()(*string) {
    return m.organizationName
}
// Serialize serializes information the current object
func (m *DnsRecord_IPGeoLocationInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("ASN", m.GetASN())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("City", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Country", m.GetCountry())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("CountryCode", m.GetCountryCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("OrganizationName", m.GetOrganizationName())
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
func (m *DnsRecord_IPGeoLocationInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetASN sets the ASN property value. The ASN property
func (m *DnsRecord_IPGeoLocationInfo) SetASN(value *int64)() {
    m.aSN = value
}
// SetCity sets the City property value. The City property
func (m *DnsRecord_IPGeoLocationInfo) SetCity(value *string)() {
    m.city = value
}
// SetCountry sets the Country property value. The Country property
func (m *DnsRecord_IPGeoLocationInfo) SetCountry(value *string)() {
    m.country = value
}
// SetCountryCode sets the CountryCode property value. The CountryCode property
func (m *DnsRecord_IPGeoLocationInfo) SetCountryCode(value *string)() {
    m.countryCode = value
}
// SetOrganizationName sets the OrganizationName property value. The OrganizationName property
func (m *DnsRecord_IPGeoLocationInfo) SetOrganizationName(value *string)() {
    m.organizationName = value
}
type DnsRecord_IPGeoLocationInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetASN()(*int64)
    GetCity()(*string)
    GetCountry()(*string)
    GetCountryCode()(*string)
    GetOrganizationName()(*string)
    SetASN(value *int64)()
    SetCity(value *string)()
    SetCountry(value *string)()
    SetCountryCode(value *string)()
    SetOrganizationName(value *string)()
}
