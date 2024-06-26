package dnszone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DnsRecord_GeolocationInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The City property
    city *string
    // The Country property
    country *string
    // The Latitude property
    latitude *float64
    // The Longitude property
    longitude *float64
}
// NewDnsRecord_GeolocationInfo instantiates a new DnsRecord_GeolocationInfo and sets the default values.
func NewDnsRecord_GeolocationInfo()(*DnsRecord_GeolocationInfo) {
    m := &DnsRecord_GeolocationInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDnsRecord_GeolocationInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDnsRecord_GeolocationInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDnsRecord_GeolocationInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DnsRecord_GeolocationInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCity gets the City property value. The City property
// returns a *string when successful
func (m *DnsRecord_GeolocationInfo) GetCity()(*string) {
    return m.city
}
// GetCountry gets the Country property value. The Country property
// returns a *string when successful
func (m *DnsRecord_GeolocationInfo) GetCountry()(*string) {
    return m.country
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DnsRecord_GeolocationInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["Latitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatitude(val)
        }
        return nil
    }
    res["Longitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLongitude(val)
        }
        return nil
    }
    return res
}
// GetLatitude gets the Latitude property value. The Latitude property
// returns a *float64 when successful
func (m *DnsRecord_GeolocationInfo) GetLatitude()(*float64) {
    return m.latitude
}
// GetLongitude gets the Longitude property value. The Longitude property
// returns a *float64 when successful
func (m *DnsRecord_GeolocationInfo) GetLongitude()(*float64) {
    return m.longitude
}
// Serialize serializes information the current object
func (m *DnsRecord_GeolocationInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteFloat64Value("Latitude", m.GetLatitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("Longitude", m.GetLongitude())
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
func (m *DnsRecord_GeolocationInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCity sets the City property value. The City property
func (m *DnsRecord_GeolocationInfo) SetCity(value *string)() {
    m.city = value
}
// SetCountry sets the Country property value. The Country property
func (m *DnsRecord_GeolocationInfo) SetCountry(value *string)() {
    m.country = value
}
// SetLatitude sets the Latitude property value. The Latitude property
func (m *DnsRecord_GeolocationInfo) SetLatitude(value *float64)() {
    m.latitude = value
}
// SetLongitude sets the Longitude property value. The Longitude property
func (m *DnsRecord_GeolocationInfo) SetLongitude(value *float64)() {
    m.longitude = value
}
type DnsRecord_GeolocationInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCity()(*string)
    GetCountry()(*string)
    GetLatitude()(*float64)
    GetLongitude()(*float64)
    SetCity(value *string)()
    SetCountry(value *string)()
    SetLatitude(value *float64)()
    SetLongitude(value *float64)()
}
