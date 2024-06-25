package user

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type User struct {
    // The Id of the user
    accountId *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The remaining balance on the user's account
    balance *float64
    // The email where the invoices and billing messages will be sent
    billingEmail *string
    // The end date of the account's free trial. If this is in the past, the free trial has expired.
    billingFreeUntilDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The BillingType property
    billingType *float64
    // The city of the user
    city *string
    // The country name that the user is from
    companyName *string
    // The Alpha2 country code that the user is from
    country *string
    // The date when the user joined bunny.net
    dateJoined *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Determines if the DPA was accepted by the user or not
    dpaAccepted *bool
    // Determines the date on which the DPA was accepted
    dpaDateAccepted *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Determines which version of the DPA was accepted
    dpaVersionAccepted *int32
    // The email of the user
    email *string
    // Determines if the account's email has been successfully verified
    emailVerified *bool
    // Contains the list of available payment types for this account
    enabledPaymentTypes []string
    // The list of features that the user has enabled
    featureFlags []string
    // The first name of the user
    firstName *string
    // The FreeTrialExtendedDate property
    freeTrialExtendedDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The HasCompleteBillingProfile property
    hasCompleteBillingProfile *bool
    // The Id of the user
    id *string
    // Determines whether the user used a Single Sign On account
    isSsoAccount *bool
    // The last name of the user
    lastName *string
    // Determines if the payments are disabled on this account
    paymentsDisabled *bool
    // Determines if the account should receive notification emails from bunny.net
    receiveNotificationEmails *bool
    // Determines if the account should receive promotional emails from bunny.net
    receivePromotionalEmails *bool
    // Determines the roles that the user belongs to
    roles []string
    // The street address of the user
    streetAddress *string
    // Determines if the user's account is suspended
    suspended *bool
    // The total bandwidth used by the account.
    totalBandwidthUsed *int64
    // The total free trial bandwidth limit for this account
    trialBandwidthLimit *int64
    // Determines if the account has 2FA enabled
    twoFactorAuthenticationEnabled *bool
    // Returns the number of unread tickets on the user's account
    unreadSupportTicketCount *int32
    // The billing VAT number of the account
    vATNumber *string
    // The address zip code of the user
    zipCode *string
}
// NewUser instantiates a new User and sets the default values.
func NewUser()(*User) {
    m := &User{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUser(), nil
}
// GetAccountId gets the AccountId property value. The Id of the user
// returns a *string when successful
func (m *User) GetAccountId()(*string) {
    return m.accountId
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *User) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBalance gets the Balance property value. The remaining balance on the user's account
// returns a *float64 when successful
func (m *User) GetBalance()(*float64) {
    return m.balance
}
// GetBillingEmail gets the BillingEmail property value. The email where the invoices and billing messages will be sent
// returns a *string when successful
func (m *User) GetBillingEmail()(*string) {
    return m.billingEmail
}
// GetBillingFreeUntilDate gets the BillingFreeUntilDate property value. The end date of the account's free trial. If this is in the past, the free trial has expired.
// returns a *Time when successful
func (m *User) GetBillingFreeUntilDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.billingFreeUntilDate
}
// GetBillingType gets the BillingType property value. The BillingType property
// returns a *float64 when successful
func (m *User) GetBillingType()(*float64) {
    return m.billingType
}
// GetCity gets the City property value. The city of the user
// returns a *string when successful
func (m *User) GetCity()(*string) {
    return m.city
}
// GetCompanyName gets the CompanyName property value. The country name that the user is from
// returns a *string when successful
func (m *User) GetCompanyName()(*string) {
    return m.companyName
}
// GetCountry gets the Country property value. The Alpha2 country code that the user is from
// returns a *string when successful
func (m *User) GetCountry()(*string) {
    return m.country
}
// GetDateJoined gets the DateJoined property value. The date when the user joined bunny.net
// returns a *Time when successful
func (m *User) GetDateJoined()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dateJoined
}
// GetDpaAccepted gets the DpaAccepted property value. Determines if the DPA was accepted by the user or not
// returns a *bool when successful
func (m *User) GetDpaAccepted()(*bool) {
    return m.dpaAccepted
}
// GetDpaDateAccepted gets the DpaDateAccepted property value. Determines the date on which the DPA was accepted
// returns a *Time when successful
func (m *User) GetDpaDateAccepted()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dpaDateAccepted
}
// GetDpaVersionAccepted gets the DpaVersionAccepted property value. Determines which version of the DPA was accepted
// returns a *int32 when successful
func (m *User) GetDpaVersionAccepted()(*int32) {
    return m.dpaVersionAccepted
}
// GetEmail gets the Email property value. The email of the user
// returns a *string when successful
func (m *User) GetEmail()(*string) {
    return m.email
}
// GetEmailVerified gets the EmailVerified property value. Determines if the account's email has been successfully verified
// returns a *bool when successful
func (m *User) GetEmailVerified()(*bool) {
    return m.emailVerified
}
// GetEnabledPaymentTypes gets the EnabledPaymentTypes property value. Contains the list of available payment types for this account
// returns a []string when successful
func (m *User) GetEnabledPaymentTypes()([]string) {
    return m.enabledPaymentTypes
}
// GetFeatureFlags gets the FeatureFlags property value. The list of features that the user has enabled
// returns a []string when successful
func (m *User) GetFeatureFlags()([]string) {
    return m.featureFlags
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *User) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["AccountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountId(val)
        }
        return nil
    }
    res["Balance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBalance(val)
        }
        return nil
    }
    res["BillingEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillingEmail(val)
        }
        return nil
    }
    res["BillingFreeUntilDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillingFreeUntilDate(val)
        }
        return nil
    }
    res["BillingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillingType(val)
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
    res["CompanyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyName(val)
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
    res["DateJoined"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateJoined(val)
        }
        return nil
    }
    res["DpaAccepted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDpaAccepted(val)
        }
        return nil
    }
    res["DpaDateAccepted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDpaDateAccepted(val)
        }
        return nil
    }
    res["DpaVersionAccepted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDpaVersionAccepted(val)
        }
        return nil
    }
    res["Email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["EmailVerified"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailVerified(val)
        }
        return nil
    }
    res["EnabledPaymentTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetEnabledPaymentTypes(res)
        }
        return nil
    }
    res["FeatureFlags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetFeatureFlags(res)
        }
        return nil
    }
    res["FirstName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstName(val)
        }
        return nil
    }
    res["FreeTrialExtendedDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFreeTrialExtendedDate(val)
        }
        return nil
    }
    res["HasCompleteBillingProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasCompleteBillingProfile(val)
        }
        return nil
    }
    res["Id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["IsSsoAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSsoAccount(val)
        }
        return nil
    }
    res["LastName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastName(val)
        }
        return nil
    }
    res["PaymentsDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPaymentsDisabled(val)
        }
        return nil
    }
    res["ReceiveNotificationEmails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceiveNotificationEmails(val)
        }
        return nil
    }
    res["ReceivePromotionalEmails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivePromotionalEmails(val)
        }
        return nil
    }
    res["Roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRoles(res)
        }
        return nil
    }
    res["StreetAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStreetAddress(val)
        }
        return nil
    }
    res["Suspended"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuspended(val)
        }
        return nil
    }
    res["TotalBandwidthUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalBandwidthUsed(val)
        }
        return nil
    }
    res["TrialBandwidthLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrialBandwidthLimit(val)
        }
        return nil
    }
    res["TwoFactorAuthenticationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTwoFactorAuthenticationEnabled(val)
        }
        return nil
    }
    res["UnreadSupportTicketCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnreadSupportTicketCount(val)
        }
        return nil
    }
    res["VATNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVATNumber(val)
        }
        return nil
    }
    res["ZipCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZipCode(val)
        }
        return nil
    }
    return res
}
// GetFirstName gets the FirstName property value. The first name of the user
// returns a *string when successful
func (m *User) GetFirstName()(*string) {
    return m.firstName
}
// GetFreeTrialExtendedDate gets the FreeTrialExtendedDate property value. The FreeTrialExtendedDate property
// returns a *Time when successful
func (m *User) GetFreeTrialExtendedDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.freeTrialExtendedDate
}
// GetHasCompleteBillingProfile gets the HasCompleteBillingProfile property value. The HasCompleteBillingProfile property
// returns a *bool when successful
func (m *User) GetHasCompleteBillingProfile()(*bool) {
    return m.hasCompleteBillingProfile
}
// GetId gets the Id property value. The Id of the user
// returns a *string when successful
func (m *User) GetId()(*string) {
    return m.id
}
// GetIsSsoAccount gets the IsSsoAccount property value. Determines whether the user used a Single Sign On account
// returns a *bool when successful
func (m *User) GetIsSsoAccount()(*bool) {
    return m.isSsoAccount
}
// GetLastName gets the LastName property value. The last name of the user
// returns a *string when successful
func (m *User) GetLastName()(*string) {
    return m.lastName
}
// GetPaymentsDisabled gets the PaymentsDisabled property value. Determines if the payments are disabled on this account
// returns a *bool when successful
func (m *User) GetPaymentsDisabled()(*bool) {
    return m.paymentsDisabled
}
// GetReceiveNotificationEmails gets the ReceiveNotificationEmails property value. Determines if the account should receive notification emails from bunny.net
// returns a *bool when successful
func (m *User) GetReceiveNotificationEmails()(*bool) {
    return m.receiveNotificationEmails
}
// GetReceivePromotionalEmails gets the ReceivePromotionalEmails property value. Determines if the account should receive promotional emails from bunny.net
// returns a *bool when successful
func (m *User) GetReceivePromotionalEmails()(*bool) {
    return m.receivePromotionalEmails
}
// GetRoles gets the Roles property value. Determines the roles that the user belongs to
// returns a []string when successful
func (m *User) GetRoles()([]string) {
    return m.roles
}
// GetStreetAddress gets the StreetAddress property value. The street address of the user
// returns a *string when successful
func (m *User) GetStreetAddress()(*string) {
    return m.streetAddress
}
// GetSuspended gets the Suspended property value. Determines if the user's account is suspended
// returns a *bool when successful
func (m *User) GetSuspended()(*bool) {
    return m.suspended
}
// GetTotalBandwidthUsed gets the TotalBandwidthUsed property value. The total bandwidth used by the account.
// returns a *int64 when successful
func (m *User) GetTotalBandwidthUsed()(*int64) {
    return m.totalBandwidthUsed
}
// GetTrialBandwidthLimit gets the TrialBandwidthLimit property value. The total free trial bandwidth limit for this account
// returns a *int64 when successful
func (m *User) GetTrialBandwidthLimit()(*int64) {
    return m.trialBandwidthLimit
}
// GetTwoFactorAuthenticationEnabled gets the TwoFactorAuthenticationEnabled property value. Determines if the account has 2FA enabled
// returns a *bool when successful
func (m *User) GetTwoFactorAuthenticationEnabled()(*bool) {
    return m.twoFactorAuthenticationEnabled
}
// GetUnreadSupportTicketCount gets the UnreadSupportTicketCount property value. Returns the number of unread tickets on the user's account
// returns a *int32 when successful
func (m *User) GetUnreadSupportTicketCount()(*int32) {
    return m.unreadSupportTicketCount
}
// GetVATNumber gets the VATNumber property value. The billing VAT number of the account
// returns a *string when successful
func (m *User) GetVATNumber()(*string) {
    return m.vATNumber
}
// GetZipCode gets the ZipCode property value. The address zip code of the user
// returns a *string when successful
func (m *User) GetZipCode()(*string) {
    return m.zipCode
}
// Serialize serializes information the current object
func (m *User) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("AccountId", m.GetAccountId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("Balance", m.GetBalance())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("BillingEmail", m.GetBillingEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("BillingFreeUntilDate", m.GetBillingFreeUntilDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("BillingType", m.GetBillingType())
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
        err := writer.WriteStringValue("CompanyName", m.GetCompanyName())
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
        err := writer.WriteTimeValue("DateJoined", m.GetDateJoined())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("DpaAccepted", m.GetDpaAccepted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("DpaDateAccepted", m.GetDpaDateAccepted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("DpaVersionAccepted", m.GetDpaVersionAccepted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EmailVerified", m.GetEmailVerified())
        if err != nil {
            return err
        }
    }
    if m.GetEnabledPaymentTypes() != nil {
        err := writer.WriteCollectionOfStringValues("EnabledPaymentTypes", m.GetEnabledPaymentTypes())
        if err != nil {
            return err
        }
    }
    if m.GetFeatureFlags() != nil {
        err := writer.WriteCollectionOfStringValues("FeatureFlags", m.GetFeatureFlags())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("FirstName", m.GetFirstName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("FreeTrialExtendedDate", m.GetFreeTrialExtendedDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("HasCompleteBillingProfile", m.GetHasCompleteBillingProfile())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("IsSsoAccount", m.GetIsSsoAccount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("LastName", m.GetLastName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("PaymentsDisabled", m.GetPaymentsDisabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ReceiveNotificationEmails", m.GetReceiveNotificationEmails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ReceivePromotionalEmails", m.GetReceivePromotionalEmails())
        if err != nil {
            return err
        }
    }
    if m.GetRoles() != nil {
        err := writer.WriteCollectionOfStringValues("Roles", m.GetRoles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("StreetAddress", m.GetStreetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("Suspended", m.GetSuspended())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("TotalBandwidthUsed", m.GetTotalBandwidthUsed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("TrialBandwidthLimit", m.GetTrialBandwidthLimit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("TwoFactorAuthenticationEnabled", m.GetTwoFactorAuthenticationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("UnreadSupportTicketCount", m.GetUnreadSupportTicketCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("VATNumber", m.GetVATNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ZipCode", m.GetZipCode())
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
// SetAccountId sets the AccountId property value. The Id of the user
func (m *User) SetAccountId(value *string)() {
    m.accountId = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *User) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBalance sets the Balance property value. The remaining balance on the user's account
func (m *User) SetBalance(value *float64)() {
    m.balance = value
}
// SetBillingEmail sets the BillingEmail property value. The email where the invoices and billing messages will be sent
func (m *User) SetBillingEmail(value *string)() {
    m.billingEmail = value
}
// SetBillingFreeUntilDate sets the BillingFreeUntilDate property value. The end date of the account's free trial. If this is in the past, the free trial has expired.
func (m *User) SetBillingFreeUntilDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.billingFreeUntilDate = value
}
// SetBillingType sets the BillingType property value. The BillingType property
func (m *User) SetBillingType(value *float64)() {
    m.billingType = value
}
// SetCity sets the City property value. The city of the user
func (m *User) SetCity(value *string)() {
    m.city = value
}
// SetCompanyName sets the CompanyName property value. The country name that the user is from
func (m *User) SetCompanyName(value *string)() {
    m.companyName = value
}
// SetCountry sets the Country property value. The Alpha2 country code that the user is from
func (m *User) SetCountry(value *string)() {
    m.country = value
}
// SetDateJoined sets the DateJoined property value. The date when the user joined bunny.net
func (m *User) SetDateJoined(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateJoined = value
}
// SetDpaAccepted sets the DpaAccepted property value. Determines if the DPA was accepted by the user or not
func (m *User) SetDpaAccepted(value *bool)() {
    m.dpaAccepted = value
}
// SetDpaDateAccepted sets the DpaDateAccepted property value. Determines the date on which the DPA was accepted
func (m *User) SetDpaDateAccepted(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dpaDateAccepted = value
}
// SetDpaVersionAccepted sets the DpaVersionAccepted property value. Determines which version of the DPA was accepted
func (m *User) SetDpaVersionAccepted(value *int32)() {
    m.dpaVersionAccepted = value
}
// SetEmail sets the Email property value. The email of the user
func (m *User) SetEmail(value *string)() {
    m.email = value
}
// SetEmailVerified sets the EmailVerified property value. Determines if the account's email has been successfully verified
func (m *User) SetEmailVerified(value *bool)() {
    m.emailVerified = value
}
// SetEnabledPaymentTypes sets the EnabledPaymentTypes property value. Contains the list of available payment types for this account
func (m *User) SetEnabledPaymentTypes(value []string)() {
    m.enabledPaymentTypes = value
}
// SetFeatureFlags sets the FeatureFlags property value. The list of features that the user has enabled
func (m *User) SetFeatureFlags(value []string)() {
    m.featureFlags = value
}
// SetFirstName sets the FirstName property value. The first name of the user
func (m *User) SetFirstName(value *string)() {
    m.firstName = value
}
// SetFreeTrialExtendedDate sets the FreeTrialExtendedDate property value. The FreeTrialExtendedDate property
func (m *User) SetFreeTrialExtendedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.freeTrialExtendedDate = value
}
// SetHasCompleteBillingProfile sets the HasCompleteBillingProfile property value. The HasCompleteBillingProfile property
func (m *User) SetHasCompleteBillingProfile(value *bool)() {
    m.hasCompleteBillingProfile = value
}
// SetId sets the Id property value. The Id of the user
func (m *User) SetId(value *string)() {
    m.id = value
}
// SetIsSsoAccount sets the IsSsoAccount property value. Determines whether the user used a Single Sign On account
func (m *User) SetIsSsoAccount(value *bool)() {
    m.isSsoAccount = value
}
// SetLastName sets the LastName property value. The last name of the user
func (m *User) SetLastName(value *string)() {
    m.lastName = value
}
// SetPaymentsDisabled sets the PaymentsDisabled property value. Determines if the payments are disabled on this account
func (m *User) SetPaymentsDisabled(value *bool)() {
    m.paymentsDisabled = value
}
// SetReceiveNotificationEmails sets the ReceiveNotificationEmails property value. Determines if the account should receive notification emails from bunny.net
func (m *User) SetReceiveNotificationEmails(value *bool)() {
    m.receiveNotificationEmails = value
}
// SetReceivePromotionalEmails sets the ReceivePromotionalEmails property value. Determines if the account should receive promotional emails from bunny.net
func (m *User) SetReceivePromotionalEmails(value *bool)() {
    m.receivePromotionalEmails = value
}
// SetRoles sets the Roles property value. Determines the roles that the user belongs to
func (m *User) SetRoles(value []string)() {
    m.roles = value
}
// SetStreetAddress sets the StreetAddress property value. The street address of the user
func (m *User) SetStreetAddress(value *string)() {
    m.streetAddress = value
}
// SetSuspended sets the Suspended property value. Determines if the user's account is suspended
func (m *User) SetSuspended(value *bool)() {
    m.suspended = value
}
// SetTotalBandwidthUsed sets the TotalBandwidthUsed property value. The total bandwidth used by the account.
func (m *User) SetTotalBandwidthUsed(value *int64)() {
    m.totalBandwidthUsed = value
}
// SetTrialBandwidthLimit sets the TrialBandwidthLimit property value. The total free trial bandwidth limit for this account
func (m *User) SetTrialBandwidthLimit(value *int64)() {
    m.trialBandwidthLimit = value
}
// SetTwoFactorAuthenticationEnabled sets the TwoFactorAuthenticationEnabled property value. Determines if the account has 2FA enabled
func (m *User) SetTwoFactorAuthenticationEnabled(value *bool)() {
    m.twoFactorAuthenticationEnabled = value
}
// SetUnreadSupportTicketCount sets the UnreadSupportTicketCount property value. Returns the number of unread tickets on the user's account
func (m *User) SetUnreadSupportTicketCount(value *int32)() {
    m.unreadSupportTicketCount = value
}
// SetVATNumber sets the VATNumber property value. The billing VAT number of the account
func (m *User) SetVATNumber(value *string)() {
    m.vATNumber = value
}
// SetZipCode sets the ZipCode property value. The address zip code of the user
func (m *User) SetZipCode(value *string)() {
    m.zipCode = value
}
type Userable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountId()(*string)
    GetBalance()(*float64)
    GetBillingEmail()(*string)
    GetBillingFreeUntilDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetBillingType()(*float64)
    GetCity()(*string)
    GetCompanyName()(*string)
    GetCountry()(*string)
    GetDateJoined()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDpaAccepted()(*bool)
    GetDpaDateAccepted()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDpaVersionAccepted()(*int32)
    GetEmail()(*string)
    GetEmailVerified()(*bool)
    GetEnabledPaymentTypes()([]string)
    GetFeatureFlags()([]string)
    GetFirstName()(*string)
    GetFreeTrialExtendedDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHasCompleteBillingProfile()(*bool)
    GetId()(*string)
    GetIsSsoAccount()(*bool)
    GetLastName()(*string)
    GetPaymentsDisabled()(*bool)
    GetReceiveNotificationEmails()(*bool)
    GetReceivePromotionalEmails()(*bool)
    GetRoles()([]string)
    GetStreetAddress()(*string)
    GetSuspended()(*bool)
    GetTotalBandwidthUsed()(*int64)
    GetTrialBandwidthLimit()(*int64)
    GetTwoFactorAuthenticationEnabled()(*bool)
    GetUnreadSupportTicketCount()(*int32)
    GetVATNumber()(*string)
    GetZipCode()(*string)
    SetAccountId(value *string)()
    SetBalance(value *float64)()
    SetBillingEmail(value *string)()
    SetBillingFreeUntilDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetBillingType(value *float64)()
    SetCity(value *string)()
    SetCompanyName(value *string)()
    SetCountry(value *string)()
    SetDateJoined(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDpaAccepted(value *bool)()
    SetDpaDateAccepted(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDpaVersionAccepted(value *int32)()
    SetEmail(value *string)()
    SetEmailVerified(value *bool)()
    SetEnabledPaymentTypes(value []string)()
    SetFeatureFlags(value []string)()
    SetFirstName(value *string)()
    SetFreeTrialExtendedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHasCompleteBillingProfile(value *bool)()
    SetId(value *string)()
    SetIsSsoAccount(value *bool)()
    SetLastName(value *string)()
    SetPaymentsDisabled(value *bool)()
    SetReceiveNotificationEmails(value *bool)()
    SetReceivePromotionalEmails(value *bool)()
    SetRoles(value []string)()
    SetStreetAddress(value *string)()
    SetSuspended(value *bool)()
    SetTotalBandwidthUsed(value *int64)()
    SetTrialBandwidthLimit(value *int64)()
    SetTwoFactorAuthenticationEnabled(value *bool)()
    SetUnreadSupportTicketCount(value *int32)()
    SetVATNumber(value *string)()
    SetZipCode(value *string)()
}
