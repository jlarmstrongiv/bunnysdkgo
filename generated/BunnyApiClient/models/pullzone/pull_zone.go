package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PullZone struct {
    // The list of extensions that will return the CORS headers
    accessControlOriginHeaderExtensions []string
    // Determines if the Add Canonical Header is enabled for this Pull Zone
    addCanonicalHeader *bool
    // Determines if the Pull Zone should forward the current hostname to the origin
    addHostHeader *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Pull Zone specific pricing discount for Africa region.
    africaDiscount *int32
    // The list of referrer hostnames that are allowed to access the pull zone.Requests containing the header Referer: hostname that is not on the list will be rejected.If empty, all the referrers are allowed
    allowedReferrers []string
    // The Pull Zone specific pricing discount for Asia & Oceania region.
    asiaOceaniaDiscount *int32
    // Determines if the AWS Signing is enabled
    aWSSigningEnabled *bool
    // The AWS Signing region key
    aWSSigningKey *string
    // The AWS Signing region name
    aWSSigningRegionName *string
    // The AWS Signing region secret
    aWSSigningSecret *string
    // The list of blocked countries with the two-letter Alpha2 ISO codes
    blockedCountries []string
    // The list of IPs that are blocked from accessing the pull zone. Requests coming from the following IPs will be rejected. If empty, all the IPs will be allowed
    blockedIps []string
    // The list of referrer hostnames that are not allowed to access the pull zone. Requests containing the header Referer: hostname that is on the list will be rejected. If empty, all the referrers are allowed
    blockedReferrers []string
    // The BlockNoneReferrer property
    blockNoneReferrer *bool
    // If true, POST requests to the zone will be blocked
    blockPostRequests *bool
    // If true, access to root path will return a 403 error
    blockRootPathAccess *bool
    // The list of budget redirected countries with the two-letter Alpha2 ISO codes
    budgetRedirectedCountries []string
    // The BunnyAiImageBlueprints property
    bunnyAiImageBlueprints []BunnyAiImageBlueprintable
    // Excessive requests are delayed until their number exceeds the maximum burst size.
    burstSize *int32
    // The override cache time for the pull zone
    cacheControlMaxAgeOverride *int64
    // The override cache time for the pull zone for the end client
    cacheControlPublicMaxAgeOverride *int64
    // Determines if bunny.net should be caching error responses
    cacheErrorResponses *bool
    // The CacheVersion property
    cacheVersion *float64
    // The CNAME domain of the pull zone for setting up custom hostnames
    cnameDomain *string
    // The number of connections limited per IP for this zone
    connectionLimitPerIPCount *int32
    // Contains the list of vary parameters that will be used for vary cache by cookie string. If empty, cookie vary will not be used.
    cookieVaryParameters []string
    // Determines if the cookies are disabled for the pull zone
    disableCookies *bool
    // If true, the built-in let's encrypt is disabled and requests are passed to the origin.
    disableLetsEncrypt *bool
    // The ID of the DNS record tied to this pull zone
    dnsRecordId *int64
    // The cached version of the DNS record value
    dnsRecordValue *string
    // The ID of the DNS zone tied to this pull zone
    dnsZoneId *int64
    // The EdgeScriptExecutionPhase property
    edgeScriptExecutionPhase *float64
    // The ID of the edge script that the pull zone is linked to
    edgeScriptId *int64
    // Determines if the CORS headers should be enabled
    enableAccessControlOriginHeader *bool
    // If set to true, any hostnames added to this Pull Zone will automatically enable SSL.
    enableAutoSSL *bool
    // Determines if the AVIF Vary feature is enabled.
    enableAvifVary *bool
    // The EnableBunnyImageAi property
    enableBunnyImageAi *bool
    // Determines if the cache slice (Optimize for video) feature is enabled for the Pull Zone
    enableCacheSlice *bool
    // Determines if the Cookie Vary feature is enabled.
    enableCookieVary *bool
    // Determines if the Country Code Vary feature is enabled.
    enableCountryCodeVary *bool
    // Determines if the Pull Zone is currently enabled, active and running
    enabled *bool
    // Determines if the delivery from the Africa region is enabled for this pull zone
    enableGeoZoneAF *bool
    // Determines if the delivery from the Asian / Oceanian region is enabled for this pull zone
    enableGeoZoneASIA *bool
    // Determines if the delivery from the European region is enabled for this pull zone
    enableGeoZoneEU *bool
    // Determines if the delivery from the South American region is enabled for this pull zone
    enableGeoZoneSA *bool
    // Determines if the delivery from the North American region is enabled for this pull zone
    enableGeoZoneUS *bool
    // Determines if the Hostname Vary feature is enabled.
    enableHostnameVary *bool
    // Determines if the logging is enabled for this Pull Zone
    enableLogging *bool
    // Determines if the Mobile Vary feature is enabled.
    enableMobileVary *bool
    // If true the server will use the origin shield feature
    enableOriginShield *bool
    // If set to true the query string ordering property is enabled.
    enableQueryStringOrdering *bool
    // Determines if request coalescing is currently enabled.
    enableRequestCoalescing *bool
    // The EnableSafeHop property
    enableSafeHop *bool
    // Determines if smart caching is enabled for this zone
    enableSmartCache *bool
    // Determines if the TLS 1 is enabled on the Pull Zone
    enableTLS1 *bool
    // Determines if the TLS 1.1 is enabled on the Pull Zone
    enableTLS1_1 *bool
    // Determines if the WebP Vary feature is enabled.
    enableWebPVary *bool
    // Contains the custom error page code that will be returned
    errorPageCustomCode *string
    // Determines if custom error page code should be enabled.
    errorPageEnableCustomCode *bool
    // Determines if the statuspage widget should be displayed on the error pages
    errorPageEnableStatuspageWidget *bool
    // The statuspage code that will be used to build the status widget
    errorPageStatuspageCode *string
    // Determines if the error pages should be whitelabel or not
    errorPageWhitelabel *bool
    // The Pull Zone specific pricing discount for EU and US region.
    eUUSDiscount *int32
    // Determines if the zone will follow origin redirects
    followRedirects *bool
    // The list of hostnames linked to this Pull Zone
    hostnames []Hostnameable
    // The unique ID of the pull zone.
    id *int64
    // True if the Pull Zone is ignoring query strings when serving cached objects
    ignoreQueryStrings *bool
    // The amount of data after the rate limit will be activated
    limitRateAfter *float64
    // The maximum rate at which the zone will transfer data in kb/s. 0 for unlimited
    limitRatePerSecond *int32
    // The LogAnonymizationType property
    logAnonymizationType *float64
    // The LogFormat property
    logFormat *float64
    // Determines if the log forwarding is enabled
    logForwardingEnabled *bool
    // The LogForwardingFormat property
    logForwardingFormat *float64
    // The log forwarding hostname
    logForwardingHostname *string
    // The log forwarding port
    logForwardingPort *int32
    // The LogForwardingProtocol property
    logForwardingProtocol *float64
    // The log forwarding token value
    logForwardingToken *string
    // Determines if the log anonymization should be enabled
    loggingIPAnonymizationEnabled *bool
    // Determines if the permanent logging feature is enabled
    loggingSaveToStorage *bool
    // The ID of the logging storage zone that is configured for this Pull Zone
    loggingStorageZoneId *int64
    // The MagicContainersAppId property
    magicContainersAppId *string
    // The MagicContainersEndpointId property
    magicContainersEndpointId *int64
    // The MiddlewareScriptId property
    middlewareScriptId *int64
    // The monthly limit of bandwidth in bytes that the pullzone is allowed to use
    monthlyBandwidthLimit *int64
    // The amount of bandwidth in bytes that the pull zone used this month
    monthlyBandwidthUsed *int64
    // The total monthly charges for this so zone so far
    monthlyCharges *float64
    // The name of the pull zone.
    name *string
    // Determines if the automatic image optimization should be enabled
    optimizerAutomaticOptimizationEnabled *bool
    // Determines the maximum automatic image size for desktop clients
    optimizerDesktopMaxWidth *int32
    // Determines if the optimizer should be enabled for this zone
    optimizerEnabled *bool
    // Determines the image manipulation should be enabled
    optimizerEnableManipulationEngine *bool
    // Determines if the WebP optimization should be enabled
    optimizerEnableWebP *bool
    // Determines if the optimizer class list should be enforced
    optimizerForceClasses *bool
    // Determines the image quality for desktop clients
    optimizerImageQuality *int32
    // Determines if the CSS minification should be enabled
    optimizerMinifyCSS *bool
    // Determines if the JavaScript minification should be enabled
    optimizerMinifyJavaScript *bool
    // Determines the image quality for mobile clients
    optimizerMobileImageQuality *int32
    // Determines the maximum automatic image size for mobile clients
    optimizerMobileMaxWidth *int32
    // The OptimizerStaticHtmlEnabled property
    optimizerStaticHtmlEnabled *bool
    // The OptimizerStaticHtmlWordPressBypassCookie property
    optimizerStaticHtmlWordPressBypassCookie *string
    // The OptimizerStaticHtmlWordPressPath property
    optimizerStaticHtmlWordPressPath *string
    // Determines if image watermarking should be enabled
    optimizerWatermarkEnabled *bool
    // Sets the minimum image size to which the watermark will be added
    optimizerWatermarkMinImageSize *int32
    // Sets the offset of the watermark image
    optimizerWatermarkOffset *float64
    // The OptimizerWatermarkPosition property
    optimizerWatermarkPosition *float64
    // Sets the URL of the watermark image
    optimizerWatermarkUrl *string
    // The amount of seconds to wait when connecting to the origin. Otherwise the request will fail or retry.
    originConnectTimeout *int32
    // Determines the host header that will be sent to the origin
    originHostHeader *string
    // Returns the link short preview value for the pull zone origin connection.
    originLinkValue *string
    // The amount of seconds to wait when waiting for the origin reply. Otherwise the request will fail or retry.
    originResponseTimeout *int32
    // The number of retries to the origin server
    originRetries *int32
    // Determines if we should retry the request in case of a 5XX response.
    originRetry5XXResponses *bool
    // Determines if we should retry the request in case of a connection timeout.
    originRetryConnectionTimeout *bool
    // Determines the amount of time that the CDN should wait before retrying an origin request.
    originRetryDelay *int32
    // Determines if we should retry the request in case of a response timeout.
    originRetryResponseTimeout *bool
    // Determines if the origin shield concurrency limit is enabled.
    originShieldEnableConcurrencyLimit *bool
    // Determines the number of maximum concurrent requests allowed to the origin.
    originShieldMaxConcurrentRequests *int32
    // Determines the max number of origin requests that will remain in the queue
    originShieldMaxQueuedRequests *int32
    // Determines the max queue wait time
    originShieldQueueMaxWaitTime *int32
    // The zone code of the origin shield
    originShieldZoneCode *string
    // The OriginType property
    originType *float64
    // The origin URL of the pull zone where the files are fetched from.
    originUrl *string
    // The IP of the storage zone used for Perma-Cache
    permaCacheStorageZoneId *int64
    // The PermaCacheType property
    permaCacheType *int64
    // The custom preloading screen code
    preloadingScreenCode *string
    // Determines if the custom preloader screen is enabled
    preloadingScreenCodeEnabled *bool
    // The delay in milliseconds after which the preloading screen will be displayed
    preloadingScreenDelay *int32
    // Determines if the preloading screen is currently enabled
    preloadingScreenEnabled *bool
    // The preloading screen logo URL
    preloadingScreenLogoUrl *string
    // The PreloadingScreenShowOnFirstVisit property
    preloadingScreenShowOnFirstVisit *bool
    // The PreloadingScreenTheme property
    preloadingScreenTheme *float64
    // The custom price override for this zone
    priceOverride *float64
    // Contains the list of vary parameters that will be used for vary cache by query string. If empty, all parameters will be used to construct the key
    queryStringVaryParameters []string
    // Determines the lock time for coalesced requests.
    requestCoalescingTimeout *int32
    // Max number of requests per IP per second
    requestLimit *int32
    // The list of routing filters enabled for this zone
    routingFilters []PullZone_RoutingFilters
    // The ShieldDDosProtectionEnabled property
    shieldDDosProtectionEnabled *bool
    // The ShieldDDosProtectionType property
    shieldDDosProtectionType *float64
    // The Pull Zone specific pricing discount for South America region.
    southAmericaDiscount *int32
    // The ID of the storage zone that the pull zone is linked to
    storageZoneId *int64
    // The Suspended property
    suspended *bool
    // The Type property
    typeEscaped *float64
    // Determines if cache update is performed in the background.
    useBackgroundUpdate *bool
    // The UserId property
    userId *string
    // Determines if we should use stale cache while the origin is offline
    useStaleWhileOffline *bool
    // Determines if we should use stale cache while cache is updating
    useStaleWhileUpdating *bool
    // Determines if the Pull Zone should verify the origin SSL certificate
    verifyOriginSSL *bool
    // The ID of the video library that the zone is linked to
    videoLibraryId *int64
    // True if the URL secure token authentication security is enabled
    zoneSecurityEnabled *bool
    // True if the zone security hash should include the remote IP
    zoneSecurityIncludeHashRemoteIP *bool
    // The security key used for secure URL token authentication
    zoneSecurityKey *string
}
// NewPullZone instantiates a new PullZone and sets the default values.
func NewPullZone()(*PullZone) {
    m := &PullZone{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePullZoneFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePullZoneFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPullZone(), nil
}
// GetAccessControlOriginHeaderExtensions gets the AccessControlOriginHeaderExtensions property value. The list of extensions that will return the CORS headers
// returns a []string when successful
func (m *PullZone) GetAccessControlOriginHeaderExtensions()([]string) {
    return m.accessControlOriginHeaderExtensions
}
// GetAddCanonicalHeader gets the AddCanonicalHeader property value. Determines if the Add Canonical Header is enabled for this Pull Zone
// returns a *bool when successful
func (m *PullZone) GetAddCanonicalHeader()(*bool) {
    return m.addCanonicalHeader
}
// GetAddHostHeader gets the AddHostHeader property value. Determines if the Pull Zone should forward the current hostname to the origin
// returns a *bool when successful
func (m *PullZone) GetAddHostHeader()(*bool) {
    return m.addHostHeader
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PullZone) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAfricaDiscount gets the AfricaDiscount property value. The Pull Zone specific pricing discount for Africa region.
// returns a *int32 when successful
func (m *PullZone) GetAfricaDiscount()(*int32) {
    return m.africaDiscount
}
// GetAllowedReferrers gets the AllowedReferrers property value. The list of referrer hostnames that are allowed to access the pull zone.Requests containing the header Referer: hostname that is not on the list will be rejected.If empty, all the referrers are allowed
// returns a []string when successful
func (m *PullZone) GetAllowedReferrers()([]string) {
    return m.allowedReferrers
}
// GetAsiaOceaniaDiscount gets the AsiaOceaniaDiscount property value. The Pull Zone specific pricing discount for Asia & Oceania region.
// returns a *int32 when successful
func (m *PullZone) GetAsiaOceaniaDiscount()(*int32) {
    return m.asiaOceaniaDiscount
}
// GetAWSSigningEnabled gets the AWSSigningEnabled property value. Determines if the AWS Signing is enabled
// returns a *bool when successful
func (m *PullZone) GetAWSSigningEnabled()(*bool) {
    return m.aWSSigningEnabled
}
// GetAWSSigningKey gets the AWSSigningKey property value. The AWS Signing region key
// returns a *string when successful
func (m *PullZone) GetAWSSigningKey()(*string) {
    return m.aWSSigningKey
}
// GetAWSSigningRegionName gets the AWSSigningRegionName property value. The AWS Signing region name
// returns a *string when successful
func (m *PullZone) GetAWSSigningRegionName()(*string) {
    return m.aWSSigningRegionName
}
// GetAWSSigningSecret gets the AWSSigningSecret property value. The AWS Signing region secret
// returns a *string when successful
func (m *PullZone) GetAWSSigningSecret()(*string) {
    return m.aWSSigningSecret
}
// GetBlockedCountries gets the BlockedCountries property value. The list of blocked countries with the two-letter Alpha2 ISO codes
// returns a []string when successful
func (m *PullZone) GetBlockedCountries()([]string) {
    return m.blockedCountries
}
// GetBlockedIps gets the BlockedIps property value. The list of IPs that are blocked from accessing the pull zone. Requests coming from the following IPs will be rejected. If empty, all the IPs will be allowed
// returns a []string when successful
func (m *PullZone) GetBlockedIps()([]string) {
    return m.blockedIps
}
// GetBlockedReferrers gets the BlockedReferrers property value. The list of referrer hostnames that are not allowed to access the pull zone. Requests containing the header Referer: hostname that is on the list will be rejected. If empty, all the referrers are allowed
// returns a []string when successful
func (m *PullZone) GetBlockedReferrers()([]string) {
    return m.blockedReferrers
}
// GetBlockNoneReferrer gets the BlockNoneReferrer property value. The BlockNoneReferrer property
// returns a *bool when successful
func (m *PullZone) GetBlockNoneReferrer()(*bool) {
    return m.blockNoneReferrer
}
// GetBlockPostRequests gets the BlockPostRequests property value. If true, POST requests to the zone will be blocked
// returns a *bool when successful
func (m *PullZone) GetBlockPostRequests()(*bool) {
    return m.blockPostRequests
}
// GetBlockRootPathAccess gets the BlockRootPathAccess property value. If true, access to root path will return a 403 error
// returns a *bool when successful
func (m *PullZone) GetBlockRootPathAccess()(*bool) {
    return m.blockRootPathAccess
}
// GetBudgetRedirectedCountries gets the BudgetRedirectedCountries property value. The list of budget redirected countries with the two-letter Alpha2 ISO codes
// returns a []string when successful
func (m *PullZone) GetBudgetRedirectedCountries()([]string) {
    return m.budgetRedirectedCountries
}
// GetBunnyAiImageBlueprints gets the BunnyAiImageBlueprints property value. The BunnyAiImageBlueprints property
// returns a []BunnyAiImageBlueprintable when successful
func (m *PullZone) GetBunnyAiImageBlueprints()([]BunnyAiImageBlueprintable) {
    return m.bunnyAiImageBlueprints
}
// GetBurstSize gets the BurstSize property value. Excessive requests are delayed until their number exceeds the maximum burst size.
// returns a *int32 when successful
func (m *PullZone) GetBurstSize()(*int32) {
    return m.burstSize
}
// GetCacheControlMaxAgeOverride gets the CacheControlMaxAgeOverride property value. The override cache time for the pull zone
// returns a *int64 when successful
func (m *PullZone) GetCacheControlMaxAgeOverride()(*int64) {
    return m.cacheControlMaxAgeOverride
}
// GetCacheControlPublicMaxAgeOverride gets the CacheControlPublicMaxAgeOverride property value. The override cache time for the pull zone for the end client
// returns a *int64 when successful
func (m *PullZone) GetCacheControlPublicMaxAgeOverride()(*int64) {
    return m.cacheControlPublicMaxAgeOverride
}
// GetCacheErrorResponses gets the CacheErrorResponses property value. Determines if bunny.net should be caching error responses
// returns a *bool when successful
func (m *PullZone) GetCacheErrorResponses()(*bool) {
    return m.cacheErrorResponses
}
// GetCacheVersion gets the CacheVersion property value. The CacheVersion property
// returns a *float64 when successful
func (m *PullZone) GetCacheVersion()(*float64) {
    return m.cacheVersion
}
// GetCnameDomain gets the CnameDomain property value. The CNAME domain of the pull zone for setting up custom hostnames
// returns a *string when successful
func (m *PullZone) GetCnameDomain()(*string) {
    return m.cnameDomain
}
// GetConnectionLimitPerIPCount gets the ConnectionLimitPerIPCount property value. The number of connections limited per IP for this zone
// returns a *int32 when successful
func (m *PullZone) GetConnectionLimitPerIPCount()(*int32) {
    return m.connectionLimitPerIPCount
}
// GetCookieVaryParameters gets the CookieVaryParameters property value. Contains the list of vary parameters that will be used for vary cache by cookie string. If empty, cookie vary will not be used.
// returns a []string when successful
func (m *PullZone) GetCookieVaryParameters()([]string) {
    return m.cookieVaryParameters
}
// GetDisableCookies gets the DisableCookies property value. Determines if the cookies are disabled for the pull zone
// returns a *bool when successful
func (m *PullZone) GetDisableCookies()(*bool) {
    return m.disableCookies
}
// GetDisableLetsEncrypt gets the DisableLetsEncrypt property value. If true, the built-in let's encrypt is disabled and requests are passed to the origin.
// returns a *bool when successful
func (m *PullZone) GetDisableLetsEncrypt()(*bool) {
    return m.disableLetsEncrypt
}
// GetDnsRecordId gets the DnsRecordId property value. The ID of the DNS record tied to this pull zone
// returns a *int64 when successful
func (m *PullZone) GetDnsRecordId()(*int64) {
    return m.dnsRecordId
}
// GetDnsRecordValue gets the DnsRecordValue property value. The cached version of the DNS record value
// returns a *string when successful
func (m *PullZone) GetDnsRecordValue()(*string) {
    return m.dnsRecordValue
}
// GetDnsZoneId gets the DnsZoneId property value. The ID of the DNS zone tied to this pull zone
// returns a *int64 when successful
func (m *PullZone) GetDnsZoneId()(*int64) {
    return m.dnsZoneId
}
// GetEdgeScriptExecutionPhase gets the EdgeScriptExecutionPhase property value. The EdgeScriptExecutionPhase property
// returns a *float64 when successful
func (m *PullZone) GetEdgeScriptExecutionPhase()(*float64) {
    return m.edgeScriptExecutionPhase
}
// GetEdgeScriptId gets the EdgeScriptId property value. The ID of the edge script that the pull zone is linked to
// returns a *int64 when successful
func (m *PullZone) GetEdgeScriptId()(*int64) {
    return m.edgeScriptId
}
// GetEnableAccessControlOriginHeader gets the EnableAccessControlOriginHeader property value. Determines if the CORS headers should be enabled
// returns a *bool when successful
func (m *PullZone) GetEnableAccessControlOriginHeader()(*bool) {
    return m.enableAccessControlOriginHeader
}
// GetEnableAutoSSL gets the EnableAutoSSL property value. If set to true, any hostnames added to this Pull Zone will automatically enable SSL.
// returns a *bool when successful
func (m *PullZone) GetEnableAutoSSL()(*bool) {
    return m.enableAutoSSL
}
// GetEnableAvifVary gets the EnableAvifVary property value. Determines if the AVIF Vary feature is enabled.
// returns a *bool when successful
func (m *PullZone) GetEnableAvifVary()(*bool) {
    return m.enableAvifVary
}
// GetEnableBunnyImageAi gets the EnableBunnyImageAi property value. The EnableBunnyImageAi property
// returns a *bool when successful
func (m *PullZone) GetEnableBunnyImageAi()(*bool) {
    return m.enableBunnyImageAi
}
// GetEnableCacheSlice gets the EnableCacheSlice property value. Determines if the cache slice (Optimize for video) feature is enabled for the Pull Zone
// returns a *bool when successful
func (m *PullZone) GetEnableCacheSlice()(*bool) {
    return m.enableCacheSlice
}
// GetEnableCookieVary gets the EnableCookieVary property value. Determines if the Cookie Vary feature is enabled.
// returns a *bool when successful
func (m *PullZone) GetEnableCookieVary()(*bool) {
    return m.enableCookieVary
}
// GetEnableCountryCodeVary gets the EnableCountryCodeVary property value. Determines if the Country Code Vary feature is enabled.
// returns a *bool when successful
func (m *PullZone) GetEnableCountryCodeVary()(*bool) {
    return m.enableCountryCodeVary
}
// GetEnabled gets the Enabled property value. Determines if the Pull Zone is currently enabled, active and running
// returns a *bool when successful
func (m *PullZone) GetEnabled()(*bool) {
    return m.enabled
}
// GetEnableGeoZoneAF gets the EnableGeoZoneAF property value. Determines if the delivery from the Africa region is enabled for this pull zone
// returns a *bool when successful
func (m *PullZone) GetEnableGeoZoneAF()(*bool) {
    return m.enableGeoZoneAF
}
// GetEnableGeoZoneASIA gets the EnableGeoZoneASIA property value. Determines if the delivery from the Asian / Oceanian region is enabled for this pull zone
// returns a *bool when successful
func (m *PullZone) GetEnableGeoZoneASIA()(*bool) {
    return m.enableGeoZoneASIA
}
// GetEnableGeoZoneEU gets the EnableGeoZoneEU property value. Determines if the delivery from the European region is enabled for this pull zone
// returns a *bool when successful
func (m *PullZone) GetEnableGeoZoneEU()(*bool) {
    return m.enableGeoZoneEU
}
// GetEnableGeoZoneSA gets the EnableGeoZoneSA property value. Determines if the delivery from the South American region is enabled for this pull zone
// returns a *bool when successful
func (m *PullZone) GetEnableGeoZoneSA()(*bool) {
    return m.enableGeoZoneSA
}
// GetEnableGeoZoneUS gets the EnableGeoZoneUS property value. Determines if the delivery from the North American region is enabled for this pull zone
// returns a *bool when successful
func (m *PullZone) GetEnableGeoZoneUS()(*bool) {
    return m.enableGeoZoneUS
}
// GetEnableHostnameVary gets the EnableHostnameVary property value. Determines if the Hostname Vary feature is enabled.
// returns a *bool when successful
func (m *PullZone) GetEnableHostnameVary()(*bool) {
    return m.enableHostnameVary
}
// GetEnableLogging gets the EnableLogging property value. Determines if the logging is enabled for this Pull Zone
// returns a *bool when successful
func (m *PullZone) GetEnableLogging()(*bool) {
    return m.enableLogging
}
// GetEnableMobileVary gets the EnableMobileVary property value. Determines if the Mobile Vary feature is enabled.
// returns a *bool when successful
func (m *PullZone) GetEnableMobileVary()(*bool) {
    return m.enableMobileVary
}
// GetEnableOriginShield gets the EnableOriginShield property value. If true the server will use the origin shield feature
// returns a *bool when successful
func (m *PullZone) GetEnableOriginShield()(*bool) {
    return m.enableOriginShield
}
// GetEnableQueryStringOrdering gets the EnableQueryStringOrdering property value. If set to true the query string ordering property is enabled.
// returns a *bool when successful
func (m *PullZone) GetEnableQueryStringOrdering()(*bool) {
    return m.enableQueryStringOrdering
}
// GetEnableRequestCoalescing gets the EnableRequestCoalescing property value. Determines if request coalescing is currently enabled.
// returns a *bool when successful
func (m *PullZone) GetEnableRequestCoalescing()(*bool) {
    return m.enableRequestCoalescing
}
// GetEnableSafeHop gets the EnableSafeHop property value. The EnableSafeHop property
// returns a *bool when successful
func (m *PullZone) GetEnableSafeHop()(*bool) {
    return m.enableSafeHop
}
// GetEnableSmartCache gets the EnableSmartCache property value. Determines if smart caching is enabled for this zone
// returns a *bool when successful
func (m *PullZone) GetEnableSmartCache()(*bool) {
    return m.enableSmartCache
}
// GetEnableTLS1 gets the EnableTLS1 property value. Determines if the TLS 1 is enabled on the Pull Zone
// returns a *bool when successful
func (m *PullZone) GetEnableTLS1()(*bool) {
    return m.enableTLS1
}
// GetEnableTLS11 gets the EnableTLS1_1 property value. Determines if the TLS 1.1 is enabled on the Pull Zone
// returns a *bool when successful
func (m *PullZone) GetEnableTLS11()(*bool) {
    return m.enableTLS1_1
}
// GetEnableWebPVary gets the EnableWebPVary property value. Determines if the WebP Vary feature is enabled.
// returns a *bool when successful
func (m *PullZone) GetEnableWebPVary()(*bool) {
    return m.enableWebPVary
}
// GetErrorPageCustomCode gets the ErrorPageCustomCode property value. Contains the custom error page code that will be returned
// returns a *string when successful
func (m *PullZone) GetErrorPageCustomCode()(*string) {
    return m.errorPageCustomCode
}
// GetErrorPageEnableCustomCode gets the ErrorPageEnableCustomCode property value. Determines if custom error page code should be enabled.
// returns a *bool when successful
func (m *PullZone) GetErrorPageEnableCustomCode()(*bool) {
    return m.errorPageEnableCustomCode
}
// GetErrorPageEnableStatuspageWidget gets the ErrorPageEnableStatuspageWidget property value. Determines if the statuspage widget should be displayed on the error pages
// returns a *bool when successful
func (m *PullZone) GetErrorPageEnableStatuspageWidget()(*bool) {
    return m.errorPageEnableStatuspageWidget
}
// GetErrorPageStatuspageCode gets the ErrorPageStatuspageCode property value. The statuspage code that will be used to build the status widget
// returns a *string when successful
func (m *PullZone) GetErrorPageStatuspageCode()(*string) {
    return m.errorPageStatuspageCode
}
// GetErrorPageWhitelabel gets the ErrorPageWhitelabel property value. Determines if the error pages should be whitelabel or not
// returns a *bool when successful
func (m *PullZone) GetErrorPageWhitelabel()(*bool) {
    return m.errorPageWhitelabel
}
// GetEUUSDiscount gets the EUUSDiscount property value. The Pull Zone specific pricing discount for EU and US region.
// returns a *int32 when successful
func (m *PullZone) GetEUUSDiscount()(*int32) {
    return m.eUUSDiscount
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PullZone) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["AccessControlOriginHeaderExtensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAccessControlOriginHeaderExtensions(res)
        }
        return nil
    }
    res["AddCanonicalHeader"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddCanonicalHeader(val)
        }
        return nil
    }
    res["AddHostHeader"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddHostHeader(val)
        }
        return nil
    }
    res["AfricaDiscount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAfricaDiscount(val)
        }
        return nil
    }
    res["AllowedReferrers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAllowedReferrers(res)
        }
        return nil
    }
    res["AsiaOceaniaDiscount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAsiaOceaniaDiscount(val)
        }
        return nil
    }
    res["AWSSigningEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAWSSigningEnabled(val)
        }
        return nil
    }
    res["AWSSigningKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAWSSigningKey(val)
        }
        return nil
    }
    res["AWSSigningRegionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAWSSigningRegionName(val)
        }
        return nil
    }
    res["AWSSigningSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAWSSigningSecret(val)
        }
        return nil
    }
    res["BlockedCountries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetBlockedCountries(res)
        }
        return nil
    }
    res["BlockedIps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetBlockedIps(res)
        }
        return nil
    }
    res["BlockedReferrers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetBlockedReferrers(res)
        }
        return nil
    }
    res["BlockNoneReferrer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockNoneReferrer(val)
        }
        return nil
    }
    res["BlockPostRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockPostRequests(val)
        }
        return nil
    }
    res["BlockRootPathAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockRootPathAccess(val)
        }
        return nil
    }
    res["BudgetRedirectedCountries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetBudgetRedirectedCountries(res)
        }
        return nil
    }
    res["BunnyAiImageBlueprints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBunnyAiImageBlueprintFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BunnyAiImageBlueprintable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BunnyAiImageBlueprintable)
                }
            }
            m.SetBunnyAiImageBlueprints(res)
        }
        return nil
    }
    res["BurstSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBurstSize(val)
        }
        return nil
    }
    res["CacheControlMaxAgeOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCacheControlMaxAgeOverride(val)
        }
        return nil
    }
    res["CacheControlPublicMaxAgeOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCacheControlPublicMaxAgeOverride(val)
        }
        return nil
    }
    res["CacheErrorResponses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCacheErrorResponses(val)
        }
        return nil
    }
    res["CacheVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCacheVersion(val)
        }
        return nil
    }
    res["CnameDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCnameDomain(val)
        }
        return nil
    }
    res["ConnectionLimitPerIPCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionLimitPerIPCount(val)
        }
        return nil
    }
    res["CookieVaryParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetCookieVaryParameters(res)
        }
        return nil
    }
    res["DisableCookies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableCookies(val)
        }
        return nil
    }
    res["DisableLetsEncrypt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableLetsEncrypt(val)
        }
        return nil
    }
    res["DnsRecordId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDnsRecordId(val)
        }
        return nil
    }
    res["DnsRecordValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDnsRecordValue(val)
        }
        return nil
    }
    res["DnsZoneId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDnsZoneId(val)
        }
        return nil
    }
    res["EdgeScriptExecutionPhase"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeScriptExecutionPhase(val)
        }
        return nil
    }
    res["EdgeScriptId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeScriptId(val)
        }
        return nil
    }
    res["EnableAccessControlOriginHeader"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableAccessControlOriginHeader(val)
        }
        return nil
    }
    res["EnableAutoSSL"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableAutoSSL(val)
        }
        return nil
    }
    res["EnableAvifVary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableAvifVary(val)
        }
        return nil
    }
    res["EnableBunnyImageAi"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableBunnyImageAi(val)
        }
        return nil
    }
    res["EnableCacheSlice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableCacheSlice(val)
        }
        return nil
    }
    res["EnableCookieVary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableCookieVary(val)
        }
        return nil
    }
    res["EnableCountryCodeVary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableCountryCodeVary(val)
        }
        return nil
    }
    res["Enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["EnableGeoZoneAF"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableGeoZoneAF(val)
        }
        return nil
    }
    res["EnableGeoZoneASIA"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableGeoZoneASIA(val)
        }
        return nil
    }
    res["EnableGeoZoneEU"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableGeoZoneEU(val)
        }
        return nil
    }
    res["EnableGeoZoneSA"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableGeoZoneSA(val)
        }
        return nil
    }
    res["EnableGeoZoneUS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableGeoZoneUS(val)
        }
        return nil
    }
    res["EnableHostnameVary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableHostnameVary(val)
        }
        return nil
    }
    res["EnableLogging"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableLogging(val)
        }
        return nil
    }
    res["EnableMobileVary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableMobileVary(val)
        }
        return nil
    }
    res["EnableOriginShield"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableOriginShield(val)
        }
        return nil
    }
    res["EnableQueryStringOrdering"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableQueryStringOrdering(val)
        }
        return nil
    }
    res["EnableRequestCoalescing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableRequestCoalescing(val)
        }
        return nil
    }
    res["EnableSafeHop"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSafeHop(val)
        }
        return nil
    }
    res["EnableSmartCache"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSmartCache(val)
        }
        return nil
    }
    res["EnableTLS1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableTLS1(val)
        }
        return nil
    }
    res["EnableTLS1_1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableTLS11(val)
        }
        return nil
    }
    res["EnableWebPVary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableWebPVary(val)
        }
        return nil
    }
    res["ErrorPageCustomCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorPageCustomCode(val)
        }
        return nil
    }
    res["ErrorPageEnableCustomCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorPageEnableCustomCode(val)
        }
        return nil
    }
    res["ErrorPageEnableStatuspageWidget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorPageEnableStatuspageWidget(val)
        }
        return nil
    }
    res["ErrorPageStatuspageCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorPageStatuspageCode(val)
        }
        return nil
    }
    res["ErrorPageWhitelabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorPageWhitelabel(val)
        }
        return nil
    }
    res["EUUSDiscount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEUUSDiscount(val)
        }
        return nil
    }
    res["FollowRedirects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFollowRedirects(val)
        }
        return nil
    }
    res["Hostnames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostnameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Hostnameable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Hostnameable)
                }
            }
            m.SetHostnames(res)
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
    res["IgnoreQueryStrings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnoreQueryStrings(val)
        }
        return nil
    }
    res["LimitRateAfter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLimitRateAfter(val)
        }
        return nil
    }
    res["LimitRatePerSecond"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLimitRatePerSecond(val)
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
    res["LogFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogFormat(val)
        }
        return nil
    }
    res["LogForwardingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogForwardingEnabled(val)
        }
        return nil
    }
    res["LogForwardingFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogForwardingFormat(val)
        }
        return nil
    }
    res["LogForwardingHostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogForwardingHostname(val)
        }
        return nil
    }
    res["LogForwardingPort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogForwardingPort(val)
        }
        return nil
    }
    res["LogForwardingProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogForwardingProtocol(val)
        }
        return nil
    }
    res["LogForwardingToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogForwardingToken(val)
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
    res["LoggingSaveToStorage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoggingSaveToStorage(val)
        }
        return nil
    }
    res["LoggingStorageZoneId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoggingStorageZoneId(val)
        }
        return nil
    }
    res["MagicContainersAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMagicContainersAppId(val)
        }
        return nil
    }
    res["MagicContainersEndpointId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMagicContainersEndpointId(val)
        }
        return nil
    }
    res["MiddlewareScriptId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiddlewareScriptId(val)
        }
        return nil
    }
    res["MonthlyBandwidthLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonthlyBandwidthLimit(val)
        }
        return nil
    }
    res["MonthlyBandwidthUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonthlyBandwidthUsed(val)
        }
        return nil
    }
    res["MonthlyCharges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonthlyCharges(val)
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
    res["OptimizerAutomaticOptimizationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerAutomaticOptimizationEnabled(val)
        }
        return nil
    }
    res["OptimizerDesktopMaxWidth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerDesktopMaxWidth(val)
        }
        return nil
    }
    res["OptimizerEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerEnabled(val)
        }
        return nil
    }
    res["OptimizerEnableManipulationEngine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerEnableManipulationEngine(val)
        }
        return nil
    }
    res["OptimizerEnableWebP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerEnableWebP(val)
        }
        return nil
    }
    res["OptimizerForceClasses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerForceClasses(val)
        }
        return nil
    }
    res["OptimizerImageQuality"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerImageQuality(val)
        }
        return nil
    }
    res["OptimizerMinifyCSS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerMinifyCSS(val)
        }
        return nil
    }
    res["OptimizerMinifyJavaScript"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerMinifyJavaScript(val)
        }
        return nil
    }
    res["OptimizerMobileImageQuality"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerMobileImageQuality(val)
        }
        return nil
    }
    res["OptimizerMobileMaxWidth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerMobileMaxWidth(val)
        }
        return nil
    }
    res["OptimizerStaticHtmlEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerStaticHtmlEnabled(val)
        }
        return nil
    }
    res["OptimizerStaticHtmlWordPressBypassCookie"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerStaticHtmlWordPressBypassCookie(val)
        }
        return nil
    }
    res["OptimizerStaticHtmlWordPressPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerStaticHtmlWordPressPath(val)
        }
        return nil
    }
    res["OptimizerWatermarkEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerWatermarkEnabled(val)
        }
        return nil
    }
    res["OptimizerWatermarkMinImageSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerWatermarkMinImageSize(val)
        }
        return nil
    }
    res["OptimizerWatermarkOffset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerWatermarkOffset(val)
        }
        return nil
    }
    res["OptimizerWatermarkPosition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerWatermarkPosition(val)
        }
        return nil
    }
    res["OptimizerWatermarkUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizerWatermarkUrl(val)
        }
        return nil
    }
    res["OriginConnectTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginConnectTimeout(val)
        }
        return nil
    }
    res["OriginHostHeader"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginHostHeader(val)
        }
        return nil
    }
    res["OriginLinkValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginLinkValue(val)
        }
        return nil
    }
    res["OriginResponseTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginResponseTimeout(val)
        }
        return nil
    }
    res["OriginRetries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginRetries(val)
        }
        return nil
    }
    res["OriginRetry5XXResponses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginRetry5XXResponses(val)
        }
        return nil
    }
    res["OriginRetryConnectionTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginRetryConnectionTimeout(val)
        }
        return nil
    }
    res["OriginRetryDelay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginRetryDelay(val)
        }
        return nil
    }
    res["OriginRetryResponseTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginRetryResponseTimeout(val)
        }
        return nil
    }
    res["OriginShieldEnableConcurrencyLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginShieldEnableConcurrencyLimit(val)
        }
        return nil
    }
    res["OriginShieldMaxConcurrentRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginShieldMaxConcurrentRequests(val)
        }
        return nil
    }
    res["OriginShieldMaxQueuedRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginShieldMaxQueuedRequests(val)
        }
        return nil
    }
    res["OriginShieldQueueMaxWaitTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginShieldQueueMaxWaitTime(val)
        }
        return nil
    }
    res["OriginShieldZoneCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginShieldZoneCode(val)
        }
        return nil
    }
    res["OriginType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginType(val)
        }
        return nil
    }
    res["OriginUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginUrl(val)
        }
        return nil
    }
    res["PermaCacheStorageZoneId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermaCacheStorageZoneId(val)
        }
        return nil
    }
    res["PermaCacheType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermaCacheType(val)
        }
        return nil
    }
    res["PreloadingScreenCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreloadingScreenCode(val)
        }
        return nil
    }
    res["PreloadingScreenCodeEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreloadingScreenCodeEnabled(val)
        }
        return nil
    }
    res["PreloadingScreenDelay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreloadingScreenDelay(val)
        }
        return nil
    }
    res["PreloadingScreenEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreloadingScreenEnabled(val)
        }
        return nil
    }
    res["PreloadingScreenLogoUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreloadingScreenLogoUrl(val)
        }
        return nil
    }
    res["PreloadingScreenShowOnFirstVisit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreloadingScreenShowOnFirstVisit(val)
        }
        return nil
    }
    res["PreloadingScreenTheme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreloadingScreenTheme(val)
        }
        return nil
    }
    res["PriceOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriceOverride(val)
        }
        return nil
    }
    res["QueryStringVaryParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetQueryStringVaryParameters(res)
        }
        return nil
    }
    res["RequestCoalescingTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestCoalescingTimeout(val)
        }
        return nil
    }
    res["RequestLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestLimit(val)
        }
        return nil
    }
    res["RoutingFilters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePullZone_RoutingFilters)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PullZone_RoutingFilters, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*PullZone_RoutingFilters))
                }
            }
            m.SetRoutingFilters(res)
        }
        return nil
    }
    res["ShieldDDosProtectionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShieldDDosProtectionEnabled(val)
        }
        return nil
    }
    res["ShieldDDosProtectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShieldDDosProtectionType(val)
        }
        return nil
    }
    res["SouthAmericaDiscount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSouthAmericaDiscount(val)
        }
        return nil
    }
    res["StorageZoneId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageZoneId(val)
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
    res["UseBackgroundUpdate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseBackgroundUpdate(val)
        }
        return nil
    }
    res["UserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["UseStaleWhileOffline"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseStaleWhileOffline(val)
        }
        return nil
    }
    res["UseStaleWhileUpdating"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseStaleWhileUpdating(val)
        }
        return nil
    }
    res["VerifyOriginSSL"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifyOriginSSL(val)
        }
        return nil
    }
    res["VideoLibraryId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoLibraryId(val)
        }
        return nil
    }
    res["ZoneSecurityEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZoneSecurityEnabled(val)
        }
        return nil
    }
    res["ZoneSecurityIncludeHashRemoteIP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZoneSecurityIncludeHashRemoteIP(val)
        }
        return nil
    }
    res["ZoneSecurityKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZoneSecurityKey(val)
        }
        return nil
    }
    return res
}
// GetFollowRedirects gets the FollowRedirects property value. Determines if the zone will follow origin redirects
// returns a *bool when successful
func (m *PullZone) GetFollowRedirects()(*bool) {
    return m.followRedirects
}
// GetHostnames gets the Hostnames property value. The list of hostnames linked to this Pull Zone
// returns a []Hostnameable when successful
func (m *PullZone) GetHostnames()([]Hostnameable) {
    return m.hostnames
}
// GetId gets the Id property value. The unique ID of the pull zone.
// returns a *int64 when successful
func (m *PullZone) GetId()(*int64) {
    return m.id
}
// GetIgnoreQueryStrings gets the IgnoreQueryStrings property value. True if the Pull Zone is ignoring query strings when serving cached objects
// returns a *bool when successful
func (m *PullZone) GetIgnoreQueryStrings()(*bool) {
    return m.ignoreQueryStrings
}
// GetLimitRateAfter gets the LimitRateAfter property value. The amount of data after the rate limit will be activated
// returns a *float64 when successful
func (m *PullZone) GetLimitRateAfter()(*float64) {
    return m.limitRateAfter
}
// GetLimitRatePerSecond gets the LimitRatePerSecond property value. The maximum rate at which the zone will transfer data in kb/s. 0 for unlimited
// returns a *int32 when successful
func (m *PullZone) GetLimitRatePerSecond()(*int32) {
    return m.limitRatePerSecond
}
// GetLogAnonymizationType gets the LogAnonymizationType property value. The LogAnonymizationType property
// returns a *float64 when successful
func (m *PullZone) GetLogAnonymizationType()(*float64) {
    return m.logAnonymizationType
}
// GetLogFormat gets the LogFormat property value. The LogFormat property
// returns a *float64 when successful
func (m *PullZone) GetLogFormat()(*float64) {
    return m.logFormat
}
// GetLogForwardingEnabled gets the LogForwardingEnabled property value. Determines if the log forwarding is enabled
// returns a *bool when successful
func (m *PullZone) GetLogForwardingEnabled()(*bool) {
    return m.logForwardingEnabled
}
// GetLogForwardingFormat gets the LogForwardingFormat property value. The LogForwardingFormat property
// returns a *float64 when successful
func (m *PullZone) GetLogForwardingFormat()(*float64) {
    return m.logForwardingFormat
}
// GetLogForwardingHostname gets the LogForwardingHostname property value. The log forwarding hostname
// returns a *string when successful
func (m *PullZone) GetLogForwardingHostname()(*string) {
    return m.logForwardingHostname
}
// GetLogForwardingPort gets the LogForwardingPort property value. The log forwarding port
// returns a *int32 when successful
func (m *PullZone) GetLogForwardingPort()(*int32) {
    return m.logForwardingPort
}
// GetLogForwardingProtocol gets the LogForwardingProtocol property value. The LogForwardingProtocol property
// returns a *float64 when successful
func (m *PullZone) GetLogForwardingProtocol()(*float64) {
    return m.logForwardingProtocol
}
// GetLogForwardingToken gets the LogForwardingToken property value. The log forwarding token value
// returns a *string when successful
func (m *PullZone) GetLogForwardingToken()(*string) {
    return m.logForwardingToken
}
// GetLoggingIPAnonymizationEnabled gets the LoggingIPAnonymizationEnabled property value. Determines if the log anonymization should be enabled
// returns a *bool when successful
func (m *PullZone) GetLoggingIPAnonymizationEnabled()(*bool) {
    return m.loggingIPAnonymizationEnabled
}
// GetLoggingSaveToStorage gets the LoggingSaveToStorage property value. Determines if the permanent logging feature is enabled
// returns a *bool when successful
func (m *PullZone) GetLoggingSaveToStorage()(*bool) {
    return m.loggingSaveToStorage
}
// GetLoggingStorageZoneId gets the LoggingStorageZoneId property value. The ID of the logging storage zone that is configured for this Pull Zone
// returns a *int64 when successful
func (m *PullZone) GetLoggingStorageZoneId()(*int64) {
    return m.loggingStorageZoneId
}
// GetMagicContainersAppId gets the MagicContainersAppId property value. The MagicContainersAppId property
// returns a *string when successful
func (m *PullZone) GetMagicContainersAppId()(*string) {
    return m.magicContainersAppId
}
// GetMagicContainersEndpointId gets the MagicContainersEndpointId property value. The MagicContainersEndpointId property
// returns a *int64 when successful
func (m *PullZone) GetMagicContainersEndpointId()(*int64) {
    return m.magicContainersEndpointId
}
// GetMiddlewareScriptId gets the MiddlewareScriptId property value. The MiddlewareScriptId property
// returns a *int64 when successful
func (m *PullZone) GetMiddlewareScriptId()(*int64) {
    return m.middlewareScriptId
}
// GetMonthlyBandwidthLimit gets the MonthlyBandwidthLimit property value. The monthly limit of bandwidth in bytes that the pullzone is allowed to use
// returns a *int64 when successful
func (m *PullZone) GetMonthlyBandwidthLimit()(*int64) {
    return m.monthlyBandwidthLimit
}
// GetMonthlyBandwidthUsed gets the MonthlyBandwidthUsed property value. The amount of bandwidth in bytes that the pull zone used this month
// returns a *int64 when successful
func (m *PullZone) GetMonthlyBandwidthUsed()(*int64) {
    return m.monthlyBandwidthUsed
}
// GetMonthlyCharges gets the MonthlyCharges property value. The total monthly charges for this so zone so far
// returns a *float64 when successful
func (m *PullZone) GetMonthlyCharges()(*float64) {
    return m.monthlyCharges
}
// GetName gets the Name property value. The name of the pull zone.
// returns a *string when successful
func (m *PullZone) GetName()(*string) {
    return m.name
}
// GetOptimizerAutomaticOptimizationEnabled gets the OptimizerAutomaticOptimizationEnabled property value. Determines if the automatic image optimization should be enabled
// returns a *bool when successful
func (m *PullZone) GetOptimizerAutomaticOptimizationEnabled()(*bool) {
    return m.optimizerAutomaticOptimizationEnabled
}
// GetOptimizerDesktopMaxWidth gets the OptimizerDesktopMaxWidth property value. Determines the maximum automatic image size for desktop clients
// returns a *int32 when successful
func (m *PullZone) GetOptimizerDesktopMaxWidth()(*int32) {
    return m.optimizerDesktopMaxWidth
}
// GetOptimizerEnabled gets the OptimizerEnabled property value. Determines if the optimizer should be enabled for this zone
// returns a *bool when successful
func (m *PullZone) GetOptimizerEnabled()(*bool) {
    return m.optimizerEnabled
}
// GetOptimizerEnableManipulationEngine gets the OptimizerEnableManipulationEngine property value. Determines the image manipulation should be enabled
// returns a *bool when successful
func (m *PullZone) GetOptimizerEnableManipulationEngine()(*bool) {
    return m.optimizerEnableManipulationEngine
}
// GetOptimizerEnableWebP gets the OptimizerEnableWebP property value. Determines if the WebP optimization should be enabled
// returns a *bool when successful
func (m *PullZone) GetOptimizerEnableWebP()(*bool) {
    return m.optimizerEnableWebP
}
// GetOptimizerForceClasses gets the OptimizerForceClasses property value. Determines if the optimizer class list should be enforced
// returns a *bool when successful
func (m *PullZone) GetOptimizerForceClasses()(*bool) {
    return m.optimizerForceClasses
}
// GetOptimizerImageQuality gets the OptimizerImageQuality property value. Determines the image quality for desktop clients
// returns a *int32 when successful
func (m *PullZone) GetOptimizerImageQuality()(*int32) {
    return m.optimizerImageQuality
}
// GetOptimizerMinifyCSS gets the OptimizerMinifyCSS property value. Determines if the CSS minification should be enabled
// returns a *bool when successful
func (m *PullZone) GetOptimizerMinifyCSS()(*bool) {
    return m.optimizerMinifyCSS
}
// GetOptimizerMinifyJavaScript gets the OptimizerMinifyJavaScript property value. Determines if the JavaScript minification should be enabled
// returns a *bool when successful
func (m *PullZone) GetOptimizerMinifyJavaScript()(*bool) {
    return m.optimizerMinifyJavaScript
}
// GetOptimizerMobileImageQuality gets the OptimizerMobileImageQuality property value. Determines the image quality for mobile clients
// returns a *int32 when successful
func (m *PullZone) GetOptimizerMobileImageQuality()(*int32) {
    return m.optimizerMobileImageQuality
}
// GetOptimizerMobileMaxWidth gets the OptimizerMobileMaxWidth property value. Determines the maximum automatic image size for mobile clients
// returns a *int32 when successful
func (m *PullZone) GetOptimizerMobileMaxWidth()(*int32) {
    return m.optimizerMobileMaxWidth
}
// GetOptimizerStaticHtmlEnabled gets the OptimizerStaticHtmlEnabled property value. The OptimizerStaticHtmlEnabled property
// returns a *bool when successful
func (m *PullZone) GetOptimizerStaticHtmlEnabled()(*bool) {
    return m.optimizerStaticHtmlEnabled
}
// GetOptimizerStaticHtmlWordPressBypassCookie gets the OptimizerStaticHtmlWordPressBypassCookie property value. The OptimizerStaticHtmlWordPressBypassCookie property
// returns a *string when successful
func (m *PullZone) GetOptimizerStaticHtmlWordPressBypassCookie()(*string) {
    return m.optimizerStaticHtmlWordPressBypassCookie
}
// GetOptimizerStaticHtmlWordPressPath gets the OptimizerStaticHtmlWordPressPath property value. The OptimizerStaticHtmlWordPressPath property
// returns a *string when successful
func (m *PullZone) GetOptimizerStaticHtmlWordPressPath()(*string) {
    return m.optimizerStaticHtmlWordPressPath
}
// GetOptimizerWatermarkEnabled gets the OptimizerWatermarkEnabled property value. Determines if image watermarking should be enabled
// returns a *bool when successful
func (m *PullZone) GetOptimizerWatermarkEnabled()(*bool) {
    return m.optimizerWatermarkEnabled
}
// GetOptimizerWatermarkMinImageSize gets the OptimizerWatermarkMinImageSize property value. Sets the minimum image size to which the watermark will be added
// returns a *int32 when successful
func (m *PullZone) GetOptimizerWatermarkMinImageSize()(*int32) {
    return m.optimizerWatermarkMinImageSize
}
// GetOptimizerWatermarkOffset gets the OptimizerWatermarkOffset property value. Sets the offset of the watermark image
// returns a *float64 when successful
func (m *PullZone) GetOptimizerWatermarkOffset()(*float64) {
    return m.optimizerWatermarkOffset
}
// GetOptimizerWatermarkPosition gets the OptimizerWatermarkPosition property value. The OptimizerWatermarkPosition property
// returns a *float64 when successful
func (m *PullZone) GetOptimizerWatermarkPosition()(*float64) {
    return m.optimizerWatermarkPosition
}
// GetOptimizerWatermarkUrl gets the OptimizerWatermarkUrl property value. Sets the URL of the watermark image
// returns a *string when successful
func (m *PullZone) GetOptimizerWatermarkUrl()(*string) {
    return m.optimizerWatermarkUrl
}
// GetOriginConnectTimeout gets the OriginConnectTimeout property value. The amount of seconds to wait when connecting to the origin. Otherwise the request will fail or retry.
// returns a *int32 when successful
func (m *PullZone) GetOriginConnectTimeout()(*int32) {
    return m.originConnectTimeout
}
// GetOriginHostHeader gets the OriginHostHeader property value. Determines the host header that will be sent to the origin
// returns a *string when successful
func (m *PullZone) GetOriginHostHeader()(*string) {
    return m.originHostHeader
}
// GetOriginLinkValue gets the OriginLinkValue property value. Returns the link short preview value for the pull zone origin connection.
// returns a *string when successful
func (m *PullZone) GetOriginLinkValue()(*string) {
    return m.originLinkValue
}
// GetOriginResponseTimeout gets the OriginResponseTimeout property value. The amount of seconds to wait when waiting for the origin reply. Otherwise the request will fail or retry.
// returns a *int32 when successful
func (m *PullZone) GetOriginResponseTimeout()(*int32) {
    return m.originResponseTimeout
}
// GetOriginRetries gets the OriginRetries property value. The number of retries to the origin server
// returns a *int32 when successful
func (m *PullZone) GetOriginRetries()(*int32) {
    return m.originRetries
}
// GetOriginRetry5XXResponses gets the OriginRetry5XXResponses property value. Determines if we should retry the request in case of a 5XX response.
// returns a *bool when successful
func (m *PullZone) GetOriginRetry5XXResponses()(*bool) {
    return m.originRetry5XXResponses
}
// GetOriginRetryConnectionTimeout gets the OriginRetryConnectionTimeout property value. Determines if we should retry the request in case of a connection timeout.
// returns a *bool when successful
func (m *PullZone) GetOriginRetryConnectionTimeout()(*bool) {
    return m.originRetryConnectionTimeout
}
// GetOriginRetryDelay gets the OriginRetryDelay property value. Determines the amount of time that the CDN should wait before retrying an origin request.
// returns a *int32 when successful
func (m *PullZone) GetOriginRetryDelay()(*int32) {
    return m.originRetryDelay
}
// GetOriginRetryResponseTimeout gets the OriginRetryResponseTimeout property value. Determines if we should retry the request in case of a response timeout.
// returns a *bool when successful
func (m *PullZone) GetOriginRetryResponseTimeout()(*bool) {
    return m.originRetryResponseTimeout
}
// GetOriginShieldEnableConcurrencyLimit gets the OriginShieldEnableConcurrencyLimit property value. Determines if the origin shield concurrency limit is enabled.
// returns a *bool when successful
func (m *PullZone) GetOriginShieldEnableConcurrencyLimit()(*bool) {
    return m.originShieldEnableConcurrencyLimit
}
// GetOriginShieldMaxConcurrentRequests gets the OriginShieldMaxConcurrentRequests property value. Determines the number of maximum concurrent requests allowed to the origin.
// returns a *int32 when successful
func (m *PullZone) GetOriginShieldMaxConcurrentRequests()(*int32) {
    return m.originShieldMaxConcurrentRequests
}
// GetOriginShieldMaxQueuedRequests gets the OriginShieldMaxQueuedRequests property value. Determines the max number of origin requests that will remain in the queue
// returns a *int32 when successful
func (m *PullZone) GetOriginShieldMaxQueuedRequests()(*int32) {
    return m.originShieldMaxQueuedRequests
}
// GetOriginShieldQueueMaxWaitTime gets the OriginShieldQueueMaxWaitTime property value. Determines the max queue wait time
// returns a *int32 when successful
func (m *PullZone) GetOriginShieldQueueMaxWaitTime()(*int32) {
    return m.originShieldQueueMaxWaitTime
}
// GetOriginShieldZoneCode gets the OriginShieldZoneCode property value. The zone code of the origin shield
// returns a *string when successful
func (m *PullZone) GetOriginShieldZoneCode()(*string) {
    return m.originShieldZoneCode
}
// GetOriginType gets the OriginType property value. The OriginType property
// returns a *float64 when successful
func (m *PullZone) GetOriginType()(*float64) {
    return m.originType
}
// GetOriginUrl gets the OriginUrl property value. The origin URL of the pull zone where the files are fetched from.
// returns a *string when successful
func (m *PullZone) GetOriginUrl()(*string) {
    return m.originUrl
}
// GetPermaCacheStorageZoneId gets the PermaCacheStorageZoneId property value. The IP of the storage zone used for Perma-Cache
// returns a *int64 when successful
func (m *PullZone) GetPermaCacheStorageZoneId()(*int64) {
    return m.permaCacheStorageZoneId
}
// GetPermaCacheType gets the PermaCacheType property value. The PermaCacheType property
// returns a *int64 when successful
func (m *PullZone) GetPermaCacheType()(*int64) {
    return m.permaCacheType
}
// GetPreloadingScreenCode gets the PreloadingScreenCode property value. The custom preloading screen code
// returns a *string when successful
func (m *PullZone) GetPreloadingScreenCode()(*string) {
    return m.preloadingScreenCode
}
// GetPreloadingScreenCodeEnabled gets the PreloadingScreenCodeEnabled property value. Determines if the custom preloader screen is enabled
// returns a *bool when successful
func (m *PullZone) GetPreloadingScreenCodeEnabled()(*bool) {
    return m.preloadingScreenCodeEnabled
}
// GetPreloadingScreenDelay gets the PreloadingScreenDelay property value. The delay in milliseconds after which the preloading screen will be displayed
// returns a *int32 when successful
func (m *PullZone) GetPreloadingScreenDelay()(*int32) {
    return m.preloadingScreenDelay
}
// GetPreloadingScreenEnabled gets the PreloadingScreenEnabled property value. Determines if the preloading screen is currently enabled
// returns a *bool when successful
func (m *PullZone) GetPreloadingScreenEnabled()(*bool) {
    return m.preloadingScreenEnabled
}
// GetPreloadingScreenLogoUrl gets the PreloadingScreenLogoUrl property value. The preloading screen logo URL
// returns a *string when successful
func (m *PullZone) GetPreloadingScreenLogoUrl()(*string) {
    return m.preloadingScreenLogoUrl
}
// GetPreloadingScreenShowOnFirstVisit gets the PreloadingScreenShowOnFirstVisit property value. The PreloadingScreenShowOnFirstVisit property
// returns a *bool when successful
func (m *PullZone) GetPreloadingScreenShowOnFirstVisit()(*bool) {
    return m.preloadingScreenShowOnFirstVisit
}
// GetPreloadingScreenTheme gets the PreloadingScreenTheme property value. The PreloadingScreenTheme property
// returns a *float64 when successful
func (m *PullZone) GetPreloadingScreenTheme()(*float64) {
    return m.preloadingScreenTheme
}
// GetPriceOverride gets the PriceOverride property value. The custom price override for this zone
// returns a *float64 when successful
func (m *PullZone) GetPriceOverride()(*float64) {
    return m.priceOverride
}
// GetQueryStringVaryParameters gets the QueryStringVaryParameters property value. Contains the list of vary parameters that will be used for vary cache by query string. If empty, all parameters will be used to construct the key
// returns a []string when successful
func (m *PullZone) GetQueryStringVaryParameters()([]string) {
    return m.queryStringVaryParameters
}
// GetRequestCoalescingTimeout gets the RequestCoalescingTimeout property value. Determines the lock time for coalesced requests.
// returns a *int32 when successful
func (m *PullZone) GetRequestCoalescingTimeout()(*int32) {
    return m.requestCoalescingTimeout
}
// GetRequestLimit gets the RequestLimit property value. Max number of requests per IP per second
// returns a *int32 when successful
func (m *PullZone) GetRequestLimit()(*int32) {
    return m.requestLimit
}
// GetRoutingFilters gets the RoutingFilters property value. The list of routing filters enabled for this zone
// returns a []PullZone_RoutingFilters when successful
func (m *PullZone) GetRoutingFilters()([]PullZone_RoutingFilters) {
    return m.routingFilters
}
// GetShieldDDosProtectionEnabled gets the ShieldDDosProtectionEnabled property value. The ShieldDDosProtectionEnabled property
// returns a *bool when successful
func (m *PullZone) GetShieldDDosProtectionEnabled()(*bool) {
    return m.shieldDDosProtectionEnabled
}
// GetShieldDDosProtectionType gets the ShieldDDosProtectionType property value. The ShieldDDosProtectionType property
// returns a *float64 when successful
func (m *PullZone) GetShieldDDosProtectionType()(*float64) {
    return m.shieldDDosProtectionType
}
// GetSouthAmericaDiscount gets the SouthAmericaDiscount property value. The Pull Zone specific pricing discount for South America region.
// returns a *int32 when successful
func (m *PullZone) GetSouthAmericaDiscount()(*int32) {
    return m.southAmericaDiscount
}
// GetStorageZoneId gets the StorageZoneId property value. The ID of the storage zone that the pull zone is linked to
// returns a *int64 when successful
func (m *PullZone) GetStorageZoneId()(*int64) {
    return m.storageZoneId
}
// GetSuspended gets the Suspended property value. The Suspended property
// returns a *bool when successful
func (m *PullZone) GetSuspended()(*bool) {
    return m.suspended
}
// GetTypeEscaped gets the Type property value. The Type property
// returns a *float64 when successful
func (m *PullZone) GetTypeEscaped()(*float64) {
    return m.typeEscaped
}
// GetUseBackgroundUpdate gets the UseBackgroundUpdate property value. Determines if cache update is performed in the background.
// returns a *bool when successful
func (m *PullZone) GetUseBackgroundUpdate()(*bool) {
    return m.useBackgroundUpdate
}
// GetUserId gets the UserId property value. The UserId property
// returns a *string when successful
func (m *PullZone) GetUserId()(*string) {
    return m.userId
}
// GetUseStaleWhileOffline gets the UseStaleWhileOffline property value. Determines if we should use stale cache while the origin is offline
// returns a *bool when successful
func (m *PullZone) GetUseStaleWhileOffline()(*bool) {
    return m.useStaleWhileOffline
}
// GetUseStaleWhileUpdating gets the UseStaleWhileUpdating property value. Determines if we should use stale cache while cache is updating
// returns a *bool when successful
func (m *PullZone) GetUseStaleWhileUpdating()(*bool) {
    return m.useStaleWhileUpdating
}
// GetVerifyOriginSSL gets the VerifyOriginSSL property value. Determines if the Pull Zone should verify the origin SSL certificate
// returns a *bool when successful
func (m *PullZone) GetVerifyOriginSSL()(*bool) {
    return m.verifyOriginSSL
}
// GetVideoLibraryId gets the VideoLibraryId property value. The ID of the video library that the zone is linked to
// returns a *int64 when successful
func (m *PullZone) GetVideoLibraryId()(*int64) {
    return m.videoLibraryId
}
// GetZoneSecurityEnabled gets the ZoneSecurityEnabled property value. True if the URL secure token authentication security is enabled
// returns a *bool when successful
func (m *PullZone) GetZoneSecurityEnabled()(*bool) {
    return m.zoneSecurityEnabled
}
// GetZoneSecurityIncludeHashRemoteIP gets the ZoneSecurityIncludeHashRemoteIP property value. True if the zone security hash should include the remote IP
// returns a *bool when successful
func (m *PullZone) GetZoneSecurityIncludeHashRemoteIP()(*bool) {
    return m.zoneSecurityIncludeHashRemoteIP
}
// GetZoneSecurityKey gets the ZoneSecurityKey property value. The security key used for secure URL token authentication
// returns a *string when successful
func (m *PullZone) GetZoneSecurityKey()(*string) {
    return m.zoneSecurityKey
}
// Serialize serializes information the current object
func (m *PullZone) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccessControlOriginHeaderExtensions() != nil {
        err := writer.WriteCollectionOfStringValues("AccessControlOriginHeaderExtensions", m.GetAccessControlOriginHeaderExtensions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("AddCanonicalHeader", m.GetAddCanonicalHeader())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("AddHostHeader", m.GetAddHostHeader())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedReferrers() != nil {
        err := writer.WriteCollectionOfStringValues("AllowedReferrers", m.GetAllowedReferrers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("AWSSigningEnabled", m.GetAWSSigningEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("AWSSigningKey", m.GetAWSSigningKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("AWSSigningRegionName", m.GetAWSSigningRegionName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("AWSSigningSecret", m.GetAWSSigningSecret())
        if err != nil {
            return err
        }
    }
    if m.GetBlockedCountries() != nil {
        err := writer.WriteCollectionOfStringValues("BlockedCountries", m.GetBlockedCountries())
        if err != nil {
            return err
        }
    }
    if m.GetBlockedIps() != nil {
        err := writer.WriteCollectionOfStringValues("BlockedIps", m.GetBlockedIps())
        if err != nil {
            return err
        }
    }
    if m.GetBlockedReferrers() != nil {
        err := writer.WriteCollectionOfStringValues("BlockedReferrers", m.GetBlockedReferrers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("BlockNoneReferrer", m.GetBlockNoneReferrer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("BlockPostRequests", m.GetBlockPostRequests())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("BlockRootPathAccess", m.GetBlockRootPathAccess())
        if err != nil {
            return err
        }
    }
    if m.GetBudgetRedirectedCountries() != nil {
        err := writer.WriteCollectionOfStringValues("BudgetRedirectedCountries", m.GetBudgetRedirectedCountries())
        if err != nil {
            return err
        }
    }
    if m.GetBunnyAiImageBlueprints() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBunnyAiImageBlueprints()))
        for i, v := range m.GetBunnyAiImageBlueprints() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("BunnyAiImageBlueprints", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("BurstSize", m.GetBurstSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("CacheControlMaxAgeOverride", m.GetCacheControlMaxAgeOverride())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("CacheControlPublicMaxAgeOverride", m.GetCacheControlPublicMaxAgeOverride())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("CacheErrorResponses", m.GetCacheErrorResponses())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("ConnectionLimitPerIPCount", m.GetConnectionLimitPerIPCount())
        if err != nil {
            return err
        }
    }
    if m.GetCookieVaryParameters() != nil {
        err := writer.WriteCollectionOfStringValues("CookieVaryParameters", m.GetCookieVaryParameters())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("DisableCookies", m.GetDisableCookies())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("DisableLetsEncrypt", m.GetDisableLetsEncrypt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("EdgeScriptExecutionPhase", m.GetEdgeScriptExecutionPhase())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("EdgeScriptId", m.GetEdgeScriptId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableAccessControlOriginHeader", m.GetEnableAccessControlOriginHeader())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableAutoSSL", m.GetEnableAutoSSL())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableAvifVary", m.GetEnableAvifVary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableBunnyImageAi", m.GetEnableBunnyImageAi())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableCacheSlice", m.GetEnableCacheSlice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableCookieVary", m.GetEnableCookieVary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableCountryCodeVary", m.GetEnableCountryCodeVary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableGeoZoneAF", m.GetEnableGeoZoneAF())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableGeoZoneASIA", m.GetEnableGeoZoneASIA())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableGeoZoneEU", m.GetEnableGeoZoneEU())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableGeoZoneSA", m.GetEnableGeoZoneSA())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableGeoZoneUS", m.GetEnableGeoZoneUS())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableHostnameVary", m.GetEnableHostnameVary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableLogging", m.GetEnableLogging())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableMobileVary", m.GetEnableMobileVary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableOriginShield", m.GetEnableOriginShield())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableQueryStringOrdering", m.GetEnableQueryStringOrdering())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableRequestCoalescing", m.GetEnableRequestCoalescing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableSafeHop", m.GetEnableSafeHop())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableSmartCache", m.GetEnableSmartCache())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableTLS1", m.GetEnableTLS1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableTLS1_1", m.GetEnableTLS11())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("EnableWebPVary", m.GetEnableWebPVary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ErrorPageCustomCode", m.GetErrorPageCustomCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ErrorPageEnableCustomCode", m.GetErrorPageEnableCustomCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ErrorPageEnableStatuspageWidget", m.GetErrorPageEnableStatuspageWidget())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ErrorPageStatuspageCode", m.GetErrorPageStatuspageCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ErrorPageWhitelabel", m.GetErrorPageWhitelabel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("FollowRedirects", m.GetFollowRedirects())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("IgnoreQueryStrings", m.GetIgnoreQueryStrings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("LimitRateAfter", m.GetLimitRateAfter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("LimitRatePerSecond", m.GetLimitRatePerSecond())
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
        err := writer.WriteFloat64Value("LogFormat", m.GetLogFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("LogForwardingEnabled", m.GetLogForwardingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("LogForwardingFormat", m.GetLogForwardingFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("LogForwardingHostname", m.GetLogForwardingHostname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("LogForwardingPort", m.GetLogForwardingPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("LogForwardingProtocol", m.GetLogForwardingProtocol())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("LogForwardingToken", m.GetLogForwardingToken())
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
        err := writer.WriteBoolValue("LoggingSaveToStorage", m.GetLoggingSaveToStorage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("LoggingStorageZoneId", m.GetLoggingStorageZoneId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("MagicContainersAppId", m.GetMagicContainersAppId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("MagicContainersEndpointId", m.GetMagicContainersEndpointId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("MiddlewareScriptId", m.GetMiddlewareScriptId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("MonthlyBandwidthLimit", m.GetMonthlyBandwidthLimit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("MonthlyCharges", m.GetMonthlyCharges())
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
        err := writer.WriteBoolValue("OptimizerAutomaticOptimizationEnabled", m.GetOptimizerAutomaticOptimizationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("OptimizerDesktopMaxWidth", m.GetOptimizerDesktopMaxWidth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("OptimizerEnabled", m.GetOptimizerEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("OptimizerEnableManipulationEngine", m.GetOptimizerEnableManipulationEngine())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("OptimizerEnableWebP", m.GetOptimizerEnableWebP())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("OptimizerForceClasses", m.GetOptimizerForceClasses())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("OptimizerImageQuality", m.GetOptimizerImageQuality())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("OptimizerMinifyCSS", m.GetOptimizerMinifyCSS())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("OptimizerMinifyJavaScript", m.GetOptimizerMinifyJavaScript())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("OptimizerMobileImageQuality", m.GetOptimizerMobileImageQuality())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("OptimizerMobileMaxWidth", m.GetOptimizerMobileMaxWidth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("OptimizerStaticHtmlEnabled", m.GetOptimizerStaticHtmlEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("OptimizerStaticHtmlWordPressBypassCookie", m.GetOptimizerStaticHtmlWordPressBypassCookie())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("OptimizerStaticHtmlWordPressPath", m.GetOptimizerStaticHtmlWordPressPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("OptimizerWatermarkEnabled", m.GetOptimizerWatermarkEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("OptimizerWatermarkMinImageSize", m.GetOptimizerWatermarkMinImageSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("OptimizerWatermarkOffset", m.GetOptimizerWatermarkOffset())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("OptimizerWatermarkPosition", m.GetOptimizerWatermarkPosition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("OptimizerWatermarkUrl", m.GetOptimizerWatermarkUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("OriginConnectTimeout", m.GetOriginConnectTimeout())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("OriginHostHeader", m.GetOriginHostHeader())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("OriginResponseTimeout", m.GetOriginResponseTimeout())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("OriginRetries", m.GetOriginRetries())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("OriginRetry5XXResponses", m.GetOriginRetry5XXResponses())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("OriginRetryConnectionTimeout", m.GetOriginRetryConnectionTimeout())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("OriginRetryDelay", m.GetOriginRetryDelay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("OriginRetryResponseTimeout", m.GetOriginRetryResponseTimeout())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("OriginShieldEnableConcurrencyLimit", m.GetOriginShieldEnableConcurrencyLimit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("OriginShieldMaxConcurrentRequests", m.GetOriginShieldMaxConcurrentRequests())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("OriginShieldMaxQueuedRequests", m.GetOriginShieldMaxQueuedRequests())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("OriginShieldQueueMaxWaitTime", m.GetOriginShieldQueueMaxWaitTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("OriginShieldZoneCode", m.GetOriginShieldZoneCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("OriginType", m.GetOriginType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("OriginUrl", m.GetOriginUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("PermaCacheStorageZoneId", m.GetPermaCacheStorageZoneId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("PermaCacheType", m.GetPermaCacheType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("PreloadingScreenCode", m.GetPreloadingScreenCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("PreloadingScreenCodeEnabled", m.GetPreloadingScreenCodeEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("PreloadingScreenDelay", m.GetPreloadingScreenDelay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("PreloadingScreenEnabled", m.GetPreloadingScreenEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("PreloadingScreenLogoUrl", m.GetPreloadingScreenLogoUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("PreloadingScreenShowOnFirstVisit", m.GetPreloadingScreenShowOnFirstVisit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("PreloadingScreenTheme", m.GetPreloadingScreenTheme())
        if err != nil {
            return err
        }
    }
    if m.GetQueryStringVaryParameters() != nil {
        err := writer.WriteCollectionOfStringValues("QueryStringVaryParameters", m.GetQueryStringVaryParameters())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("RequestCoalescingTimeout", m.GetRequestCoalescingTimeout())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("RequestLimit", m.GetRequestLimit())
        if err != nil {
            return err
        }
    }
    if m.GetRoutingFilters() != nil {
        err := writer.WriteCollectionOfStringValues("RoutingFilters", SerializePullZone_RoutingFilters(m.GetRoutingFilters()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ShieldDDosProtectionEnabled", m.GetShieldDDosProtectionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("ShieldDDosProtectionType", m.GetShieldDDosProtectionType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("StorageZoneId", m.GetStorageZoneId())
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
        err := writer.WriteBoolValue("UseBackgroundUpdate", m.GetUseBackgroundUpdate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("UseStaleWhileOffline", m.GetUseStaleWhileOffline())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("UseStaleWhileUpdating", m.GetUseStaleWhileUpdating())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("VerifyOriginSSL", m.GetVerifyOriginSSL())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ZoneSecurityEnabled", m.GetZoneSecurityEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ZoneSecurityIncludeHashRemoteIP", m.GetZoneSecurityIncludeHashRemoteIP())
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
// SetAccessControlOriginHeaderExtensions sets the AccessControlOriginHeaderExtensions property value. The list of extensions that will return the CORS headers
func (m *PullZone) SetAccessControlOriginHeaderExtensions(value []string)() {
    m.accessControlOriginHeaderExtensions = value
}
// SetAddCanonicalHeader sets the AddCanonicalHeader property value. Determines if the Add Canonical Header is enabled for this Pull Zone
func (m *PullZone) SetAddCanonicalHeader(value *bool)() {
    m.addCanonicalHeader = value
}
// SetAddHostHeader sets the AddHostHeader property value. Determines if the Pull Zone should forward the current hostname to the origin
func (m *PullZone) SetAddHostHeader(value *bool)() {
    m.addHostHeader = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PullZone) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAfricaDiscount sets the AfricaDiscount property value. The Pull Zone specific pricing discount for Africa region.
func (m *PullZone) SetAfricaDiscount(value *int32)() {
    m.africaDiscount = value
}
// SetAllowedReferrers sets the AllowedReferrers property value. The list of referrer hostnames that are allowed to access the pull zone.Requests containing the header Referer: hostname that is not on the list will be rejected.If empty, all the referrers are allowed
func (m *PullZone) SetAllowedReferrers(value []string)() {
    m.allowedReferrers = value
}
// SetAsiaOceaniaDiscount sets the AsiaOceaniaDiscount property value. The Pull Zone specific pricing discount for Asia & Oceania region.
func (m *PullZone) SetAsiaOceaniaDiscount(value *int32)() {
    m.asiaOceaniaDiscount = value
}
// SetAWSSigningEnabled sets the AWSSigningEnabled property value. Determines if the AWS Signing is enabled
func (m *PullZone) SetAWSSigningEnabled(value *bool)() {
    m.aWSSigningEnabled = value
}
// SetAWSSigningKey sets the AWSSigningKey property value. The AWS Signing region key
func (m *PullZone) SetAWSSigningKey(value *string)() {
    m.aWSSigningKey = value
}
// SetAWSSigningRegionName sets the AWSSigningRegionName property value. The AWS Signing region name
func (m *PullZone) SetAWSSigningRegionName(value *string)() {
    m.aWSSigningRegionName = value
}
// SetAWSSigningSecret sets the AWSSigningSecret property value. The AWS Signing region secret
func (m *PullZone) SetAWSSigningSecret(value *string)() {
    m.aWSSigningSecret = value
}
// SetBlockedCountries sets the BlockedCountries property value. The list of blocked countries with the two-letter Alpha2 ISO codes
func (m *PullZone) SetBlockedCountries(value []string)() {
    m.blockedCountries = value
}
// SetBlockedIps sets the BlockedIps property value. The list of IPs that are blocked from accessing the pull zone. Requests coming from the following IPs will be rejected. If empty, all the IPs will be allowed
func (m *PullZone) SetBlockedIps(value []string)() {
    m.blockedIps = value
}
// SetBlockedReferrers sets the BlockedReferrers property value. The list of referrer hostnames that are not allowed to access the pull zone. Requests containing the header Referer: hostname that is on the list will be rejected. If empty, all the referrers are allowed
func (m *PullZone) SetBlockedReferrers(value []string)() {
    m.blockedReferrers = value
}
// SetBlockNoneReferrer sets the BlockNoneReferrer property value. The BlockNoneReferrer property
func (m *PullZone) SetBlockNoneReferrer(value *bool)() {
    m.blockNoneReferrer = value
}
// SetBlockPostRequests sets the BlockPostRequests property value. If true, POST requests to the zone will be blocked
func (m *PullZone) SetBlockPostRequests(value *bool)() {
    m.blockPostRequests = value
}
// SetBlockRootPathAccess sets the BlockRootPathAccess property value. If true, access to root path will return a 403 error
func (m *PullZone) SetBlockRootPathAccess(value *bool)() {
    m.blockRootPathAccess = value
}
// SetBudgetRedirectedCountries sets the BudgetRedirectedCountries property value. The list of budget redirected countries with the two-letter Alpha2 ISO codes
func (m *PullZone) SetBudgetRedirectedCountries(value []string)() {
    m.budgetRedirectedCountries = value
}
// SetBunnyAiImageBlueprints sets the BunnyAiImageBlueprints property value. The BunnyAiImageBlueprints property
func (m *PullZone) SetBunnyAiImageBlueprints(value []BunnyAiImageBlueprintable)() {
    m.bunnyAiImageBlueprints = value
}
// SetBurstSize sets the BurstSize property value. Excessive requests are delayed until their number exceeds the maximum burst size.
func (m *PullZone) SetBurstSize(value *int32)() {
    m.burstSize = value
}
// SetCacheControlMaxAgeOverride sets the CacheControlMaxAgeOverride property value. The override cache time for the pull zone
func (m *PullZone) SetCacheControlMaxAgeOverride(value *int64)() {
    m.cacheControlMaxAgeOverride = value
}
// SetCacheControlPublicMaxAgeOverride sets the CacheControlPublicMaxAgeOverride property value. The override cache time for the pull zone for the end client
func (m *PullZone) SetCacheControlPublicMaxAgeOverride(value *int64)() {
    m.cacheControlPublicMaxAgeOverride = value
}
// SetCacheErrorResponses sets the CacheErrorResponses property value. Determines if bunny.net should be caching error responses
func (m *PullZone) SetCacheErrorResponses(value *bool)() {
    m.cacheErrorResponses = value
}
// SetCacheVersion sets the CacheVersion property value. The CacheVersion property
func (m *PullZone) SetCacheVersion(value *float64)() {
    m.cacheVersion = value
}
// SetCnameDomain sets the CnameDomain property value. The CNAME domain of the pull zone for setting up custom hostnames
func (m *PullZone) SetCnameDomain(value *string)() {
    m.cnameDomain = value
}
// SetConnectionLimitPerIPCount sets the ConnectionLimitPerIPCount property value. The number of connections limited per IP for this zone
func (m *PullZone) SetConnectionLimitPerIPCount(value *int32)() {
    m.connectionLimitPerIPCount = value
}
// SetCookieVaryParameters sets the CookieVaryParameters property value. Contains the list of vary parameters that will be used for vary cache by cookie string. If empty, cookie vary will not be used.
func (m *PullZone) SetCookieVaryParameters(value []string)() {
    m.cookieVaryParameters = value
}
// SetDisableCookies sets the DisableCookies property value. Determines if the cookies are disabled for the pull zone
func (m *PullZone) SetDisableCookies(value *bool)() {
    m.disableCookies = value
}
// SetDisableLetsEncrypt sets the DisableLetsEncrypt property value. If true, the built-in let's encrypt is disabled and requests are passed to the origin.
func (m *PullZone) SetDisableLetsEncrypt(value *bool)() {
    m.disableLetsEncrypt = value
}
// SetDnsRecordId sets the DnsRecordId property value. The ID of the DNS record tied to this pull zone
func (m *PullZone) SetDnsRecordId(value *int64)() {
    m.dnsRecordId = value
}
// SetDnsRecordValue sets the DnsRecordValue property value. The cached version of the DNS record value
func (m *PullZone) SetDnsRecordValue(value *string)() {
    m.dnsRecordValue = value
}
// SetDnsZoneId sets the DnsZoneId property value. The ID of the DNS zone tied to this pull zone
func (m *PullZone) SetDnsZoneId(value *int64)() {
    m.dnsZoneId = value
}
// SetEdgeScriptExecutionPhase sets the EdgeScriptExecutionPhase property value. The EdgeScriptExecutionPhase property
func (m *PullZone) SetEdgeScriptExecutionPhase(value *float64)() {
    m.edgeScriptExecutionPhase = value
}
// SetEdgeScriptId sets the EdgeScriptId property value. The ID of the edge script that the pull zone is linked to
func (m *PullZone) SetEdgeScriptId(value *int64)() {
    m.edgeScriptId = value
}
// SetEnableAccessControlOriginHeader sets the EnableAccessControlOriginHeader property value. Determines if the CORS headers should be enabled
func (m *PullZone) SetEnableAccessControlOriginHeader(value *bool)() {
    m.enableAccessControlOriginHeader = value
}
// SetEnableAutoSSL sets the EnableAutoSSL property value. If set to true, any hostnames added to this Pull Zone will automatically enable SSL.
func (m *PullZone) SetEnableAutoSSL(value *bool)() {
    m.enableAutoSSL = value
}
// SetEnableAvifVary sets the EnableAvifVary property value. Determines if the AVIF Vary feature is enabled.
func (m *PullZone) SetEnableAvifVary(value *bool)() {
    m.enableAvifVary = value
}
// SetEnableBunnyImageAi sets the EnableBunnyImageAi property value. The EnableBunnyImageAi property
func (m *PullZone) SetEnableBunnyImageAi(value *bool)() {
    m.enableBunnyImageAi = value
}
// SetEnableCacheSlice sets the EnableCacheSlice property value. Determines if the cache slice (Optimize for video) feature is enabled for the Pull Zone
func (m *PullZone) SetEnableCacheSlice(value *bool)() {
    m.enableCacheSlice = value
}
// SetEnableCookieVary sets the EnableCookieVary property value. Determines if the Cookie Vary feature is enabled.
func (m *PullZone) SetEnableCookieVary(value *bool)() {
    m.enableCookieVary = value
}
// SetEnableCountryCodeVary sets the EnableCountryCodeVary property value. Determines if the Country Code Vary feature is enabled.
func (m *PullZone) SetEnableCountryCodeVary(value *bool)() {
    m.enableCountryCodeVary = value
}
// SetEnabled sets the Enabled property value. Determines if the Pull Zone is currently enabled, active and running
func (m *PullZone) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetEnableGeoZoneAF sets the EnableGeoZoneAF property value. Determines if the delivery from the Africa region is enabled for this pull zone
func (m *PullZone) SetEnableGeoZoneAF(value *bool)() {
    m.enableGeoZoneAF = value
}
// SetEnableGeoZoneASIA sets the EnableGeoZoneASIA property value. Determines if the delivery from the Asian / Oceanian region is enabled for this pull zone
func (m *PullZone) SetEnableGeoZoneASIA(value *bool)() {
    m.enableGeoZoneASIA = value
}
// SetEnableGeoZoneEU sets the EnableGeoZoneEU property value. Determines if the delivery from the European region is enabled for this pull zone
func (m *PullZone) SetEnableGeoZoneEU(value *bool)() {
    m.enableGeoZoneEU = value
}
// SetEnableGeoZoneSA sets the EnableGeoZoneSA property value. Determines if the delivery from the South American region is enabled for this pull zone
func (m *PullZone) SetEnableGeoZoneSA(value *bool)() {
    m.enableGeoZoneSA = value
}
// SetEnableGeoZoneUS sets the EnableGeoZoneUS property value. Determines if the delivery from the North American region is enabled for this pull zone
func (m *PullZone) SetEnableGeoZoneUS(value *bool)() {
    m.enableGeoZoneUS = value
}
// SetEnableHostnameVary sets the EnableHostnameVary property value. Determines if the Hostname Vary feature is enabled.
func (m *PullZone) SetEnableHostnameVary(value *bool)() {
    m.enableHostnameVary = value
}
// SetEnableLogging sets the EnableLogging property value. Determines if the logging is enabled for this Pull Zone
func (m *PullZone) SetEnableLogging(value *bool)() {
    m.enableLogging = value
}
// SetEnableMobileVary sets the EnableMobileVary property value. Determines if the Mobile Vary feature is enabled.
func (m *PullZone) SetEnableMobileVary(value *bool)() {
    m.enableMobileVary = value
}
// SetEnableOriginShield sets the EnableOriginShield property value. If true the server will use the origin shield feature
func (m *PullZone) SetEnableOriginShield(value *bool)() {
    m.enableOriginShield = value
}
// SetEnableQueryStringOrdering sets the EnableQueryStringOrdering property value. If set to true the query string ordering property is enabled.
func (m *PullZone) SetEnableQueryStringOrdering(value *bool)() {
    m.enableQueryStringOrdering = value
}
// SetEnableRequestCoalescing sets the EnableRequestCoalescing property value. Determines if request coalescing is currently enabled.
func (m *PullZone) SetEnableRequestCoalescing(value *bool)() {
    m.enableRequestCoalescing = value
}
// SetEnableSafeHop sets the EnableSafeHop property value. The EnableSafeHop property
func (m *PullZone) SetEnableSafeHop(value *bool)() {
    m.enableSafeHop = value
}
// SetEnableSmartCache sets the EnableSmartCache property value. Determines if smart caching is enabled for this zone
func (m *PullZone) SetEnableSmartCache(value *bool)() {
    m.enableSmartCache = value
}
// SetEnableTLS1 sets the EnableTLS1 property value. Determines if the TLS 1 is enabled on the Pull Zone
func (m *PullZone) SetEnableTLS1(value *bool)() {
    m.enableTLS1 = value
}
// SetEnableTLS11 sets the EnableTLS1_1 property value. Determines if the TLS 1.1 is enabled on the Pull Zone
func (m *PullZone) SetEnableTLS11(value *bool)() {
    m.enableTLS1_1 = value
}
// SetEnableWebPVary sets the EnableWebPVary property value. Determines if the WebP Vary feature is enabled.
func (m *PullZone) SetEnableWebPVary(value *bool)() {
    m.enableWebPVary = value
}
// SetErrorPageCustomCode sets the ErrorPageCustomCode property value. Contains the custom error page code that will be returned
func (m *PullZone) SetErrorPageCustomCode(value *string)() {
    m.errorPageCustomCode = value
}
// SetErrorPageEnableCustomCode sets the ErrorPageEnableCustomCode property value. Determines if custom error page code should be enabled.
func (m *PullZone) SetErrorPageEnableCustomCode(value *bool)() {
    m.errorPageEnableCustomCode = value
}
// SetErrorPageEnableStatuspageWidget sets the ErrorPageEnableStatuspageWidget property value. Determines if the statuspage widget should be displayed on the error pages
func (m *PullZone) SetErrorPageEnableStatuspageWidget(value *bool)() {
    m.errorPageEnableStatuspageWidget = value
}
// SetErrorPageStatuspageCode sets the ErrorPageStatuspageCode property value. The statuspage code that will be used to build the status widget
func (m *PullZone) SetErrorPageStatuspageCode(value *string)() {
    m.errorPageStatuspageCode = value
}
// SetErrorPageWhitelabel sets the ErrorPageWhitelabel property value. Determines if the error pages should be whitelabel or not
func (m *PullZone) SetErrorPageWhitelabel(value *bool)() {
    m.errorPageWhitelabel = value
}
// SetEUUSDiscount sets the EUUSDiscount property value. The Pull Zone specific pricing discount for EU and US region.
func (m *PullZone) SetEUUSDiscount(value *int32)() {
    m.eUUSDiscount = value
}
// SetFollowRedirects sets the FollowRedirects property value. Determines if the zone will follow origin redirects
func (m *PullZone) SetFollowRedirects(value *bool)() {
    m.followRedirects = value
}
// SetHostnames sets the Hostnames property value. The list of hostnames linked to this Pull Zone
func (m *PullZone) SetHostnames(value []Hostnameable)() {
    m.hostnames = value
}
// SetId sets the Id property value. The unique ID of the pull zone.
func (m *PullZone) SetId(value *int64)() {
    m.id = value
}
// SetIgnoreQueryStrings sets the IgnoreQueryStrings property value. True if the Pull Zone is ignoring query strings when serving cached objects
func (m *PullZone) SetIgnoreQueryStrings(value *bool)() {
    m.ignoreQueryStrings = value
}
// SetLimitRateAfter sets the LimitRateAfter property value. The amount of data after the rate limit will be activated
func (m *PullZone) SetLimitRateAfter(value *float64)() {
    m.limitRateAfter = value
}
// SetLimitRatePerSecond sets the LimitRatePerSecond property value. The maximum rate at which the zone will transfer data in kb/s. 0 for unlimited
func (m *PullZone) SetLimitRatePerSecond(value *int32)() {
    m.limitRatePerSecond = value
}
// SetLogAnonymizationType sets the LogAnonymizationType property value. The LogAnonymizationType property
func (m *PullZone) SetLogAnonymizationType(value *float64)() {
    m.logAnonymizationType = value
}
// SetLogFormat sets the LogFormat property value. The LogFormat property
func (m *PullZone) SetLogFormat(value *float64)() {
    m.logFormat = value
}
// SetLogForwardingEnabled sets the LogForwardingEnabled property value. Determines if the log forwarding is enabled
func (m *PullZone) SetLogForwardingEnabled(value *bool)() {
    m.logForwardingEnabled = value
}
// SetLogForwardingFormat sets the LogForwardingFormat property value. The LogForwardingFormat property
func (m *PullZone) SetLogForwardingFormat(value *float64)() {
    m.logForwardingFormat = value
}
// SetLogForwardingHostname sets the LogForwardingHostname property value. The log forwarding hostname
func (m *PullZone) SetLogForwardingHostname(value *string)() {
    m.logForwardingHostname = value
}
// SetLogForwardingPort sets the LogForwardingPort property value. The log forwarding port
func (m *PullZone) SetLogForwardingPort(value *int32)() {
    m.logForwardingPort = value
}
// SetLogForwardingProtocol sets the LogForwardingProtocol property value. The LogForwardingProtocol property
func (m *PullZone) SetLogForwardingProtocol(value *float64)() {
    m.logForwardingProtocol = value
}
// SetLogForwardingToken sets the LogForwardingToken property value. The log forwarding token value
func (m *PullZone) SetLogForwardingToken(value *string)() {
    m.logForwardingToken = value
}
// SetLoggingIPAnonymizationEnabled sets the LoggingIPAnonymizationEnabled property value. Determines if the log anonymization should be enabled
func (m *PullZone) SetLoggingIPAnonymizationEnabled(value *bool)() {
    m.loggingIPAnonymizationEnabled = value
}
// SetLoggingSaveToStorage sets the LoggingSaveToStorage property value. Determines if the permanent logging feature is enabled
func (m *PullZone) SetLoggingSaveToStorage(value *bool)() {
    m.loggingSaveToStorage = value
}
// SetLoggingStorageZoneId sets the LoggingStorageZoneId property value. The ID of the logging storage zone that is configured for this Pull Zone
func (m *PullZone) SetLoggingStorageZoneId(value *int64)() {
    m.loggingStorageZoneId = value
}
// SetMagicContainersAppId sets the MagicContainersAppId property value. The MagicContainersAppId property
func (m *PullZone) SetMagicContainersAppId(value *string)() {
    m.magicContainersAppId = value
}
// SetMagicContainersEndpointId sets the MagicContainersEndpointId property value. The MagicContainersEndpointId property
func (m *PullZone) SetMagicContainersEndpointId(value *int64)() {
    m.magicContainersEndpointId = value
}
// SetMiddlewareScriptId sets the MiddlewareScriptId property value. The MiddlewareScriptId property
func (m *PullZone) SetMiddlewareScriptId(value *int64)() {
    m.middlewareScriptId = value
}
// SetMonthlyBandwidthLimit sets the MonthlyBandwidthLimit property value. The monthly limit of bandwidth in bytes that the pullzone is allowed to use
func (m *PullZone) SetMonthlyBandwidthLimit(value *int64)() {
    m.monthlyBandwidthLimit = value
}
// SetMonthlyBandwidthUsed sets the MonthlyBandwidthUsed property value. The amount of bandwidth in bytes that the pull zone used this month
func (m *PullZone) SetMonthlyBandwidthUsed(value *int64)() {
    m.monthlyBandwidthUsed = value
}
// SetMonthlyCharges sets the MonthlyCharges property value. The total monthly charges for this so zone so far
func (m *PullZone) SetMonthlyCharges(value *float64)() {
    m.monthlyCharges = value
}
// SetName sets the Name property value. The name of the pull zone.
func (m *PullZone) SetName(value *string)() {
    m.name = value
}
// SetOptimizerAutomaticOptimizationEnabled sets the OptimizerAutomaticOptimizationEnabled property value. Determines if the automatic image optimization should be enabled
func (m *PullZone) SetOptimizerAutomaticOptimizationEnabled(value *bool)() {
    m.optimizerAutomaticOptimizationEnabled = value
}
// SetOptimizerDesktopMaxWidth sets the OptimizerDesktopMaxWidth property value. Determines the maximum automatic image size for desktop clients
func (m *PullZone) SetOptimizerDesktopMaxWidth(value *int32)() {
    m.optimizerDesktopMaxWidth = value
}
// SetOptimizerEnabled sets the OptimizerEnabled property value. Determines if the optimizer should be enabled for this zone
func (m *PullZone) SetOptimizerEnabled(value *bool)() {
    m.optimizerEnabled = value
}
// SetOptimizerEnableManipulationEngine sets the OptimizerEnableManipulationEngine property value. Determines the image manipulation should be enabled
func (m *PullZone) SetOptimizerEnableManipulationEngine(value *bool)() {
    m.optimizerEnableManipulationEngine = value
}
// SetOptimizerEnableWebP sets the OptimizerEnableWebP property value. Determines if the WebP optimization should be enabled
func (m *PullZone) SetOptimizerEnableWebP(value *bool)() {
    m.optimizerEnableWebP = value
}
// SetOptimizerForceClasses sets the OptimizerForceClasses property value. Determines if the optimizer class list should be enforced
func (m *PullZone) SetOptimizerForceClasses(value *bool)() {
    m.optimizerForceClasses = value
}
// SetOptimizerImageQuality sets the OptimizerImageQuality property value. Determines the image quality for desktop clients
func (m *PullZone) SetOptimizerImageQuality(value *int32)() {
    m.optimizerImageQuality = value
}
// SetOptimizerMinifyCSS sets the OptimizerMinifyCSS property value. Determines if the CSS minification should be enabled
func (m *PullZone) SetOptimizerMinifyCSS(value *bool)() {
    m.optimizerMinifyCSS = value
}
// SetOptimizerMinifyJavaScript sets the OptimizerMinifyJavaScript property value. Determines if the JavaScript minification should be enabled
func (m *PullZone) SetOptimizerMinifyJavaScript(value *bool)() {
    m.optimizerMinifyJavaScript = value
}
// SetOptimizerMobileImageQuality sets the OptimizerMobileImageQuality property value. Determines the image quality for mobile clients
func (m *PullZone) SetOptimizerMobileImageQuality(value *int32)() {
    m.optimizerMobileImageQuality = value
}
// SetOptimizerMobileMaxWidth sets the OptimizerMobileMaxWidth property value. Determines the maximum automatic image size for mobile clients
func (m *PullZone) SetOptimizerMobileMaxWidth(value *int32)() {
    m.optimizerMobileMaxWidth = value
}
// SetOptimizerStaticHtmlEnabled sets the OptimizerStaticHtmlEnabled property value. The OptimizerStaticHtmlEnabled property
func (m *PullZone) SetOptimizerStaticHtmlEnabled(value *bool)() {
    m.optimizerStaticHtmlEnabled = value
}
// SetOptimizerStaticHtmlWordPressBypassCookie sets the OptimizerStaticHtmlWordPressBypassCookie property value. The OptimizerStaticHtmlWordPressBypassCookie property
func (m *PullZone) SetOptimizerStaticHtmlWordPressBypassCookie(value *string)() {
    m.optimizerStaticHtmlWordPressBypassCookie = value
}
// SetOptimizerStaticHtmlWordPressPath sets the OptimizerStaticHtmlWordPressPath property value. The OptimizerStaticHtmlWordPressPath property
func (m *PullZone) SetOptimizerStaticHtmlWordPressPath(value *string)() {
    m.optimizerStaticHtmlWordPressPath = value
}
// SetOptimizerWatermarkEnabled sets the OptimizerWatermarkEnabled property value. Determines if image watermarking should be enabled
func (m *PullZone) SetOptimizerWatermarkEnabled(value *bool)() {
    m.optimizerWatermarkEnabled = value
}
// SetOptimizerWatermarkMinImageSize sets the OptimizerWatermarkMinImageSize property value. Sets the minimum image size to which the watermark will be added
func (m *PullZone) SetOptimizerWatermarkMinImageSize(value *int32)() {
    m.optimizerWatermarkMinImageSize = value
}
// SetOptimizerWatermarkOffset sets the OptimizerWatermarkOffset property value. Sets the offset of the watermark image
func (m *PullZone) SetOptimizerWatermarkOffset(value *float64)() {
    m.optimizerWatermarkOffset = value
}
// SetOptimizerWatermarkPosition sets the OptimizerWatermarkPosition property value. The OptimizerWatermarkPosition property
func (m *PullZone) SetOptimizerWatermarkPosition(value *float64)() {
    m.optimizerWatermarkPosition = value
}
// SetOptimizerWatermarkUrl sets the OptimizerWatermarkUrl property value. Sets the URL of the watermark image
func (m *PullZone) SetOptimizerWatermarkUrl(value *string)() {
    m.optimizerWatermarkUrl = value
}
// SetOriginConnectTimeout sets the OriginConnectTimeout property value. The amount of seconds to wait when connecting to the origin. Otherwise the request will fail or retry.
func (m *PullZone) SetOriginConnectTimeout(value *int32)() {
    m.originConnectTimeout = value
}
// SetOriginHostHeader sets the OriginHostHeader property value. Determines the host header that will be sent to the origin
func (m *PullZone) SetOriginHostHeader(value *string)() {
    m.originHostHeader = value
}
// SetOriginLinkValue sets the OriginLinkValue property value. Returns the link short preview value for the pull zone origin connection.
func (m *PullZone) SetOriginLinkValue(value *string)() {
    m.originLinkValue = value
}
// SetOriginResponseTimeout sets the OriginResponseTimeout property value. The amount of seconds to wait when waiting for the origin reply. Otherwise the request will fail or retry.
func (m *PullZone) SetOriginResponseTimeout(value *int32)() {
    m.originResponseTimeout = value
}
// SetOriginRetries sets the OriginRetries property value. The number of retries to the origin server
func (m *PullZone) SetOriginRetries(value *int32)() {
    m.originRetries = value
}
// SetOriginRetry5XXResponses sets the OriginRetry5XXResponses property value. Determines if we should retry the request in case of a 5XX response.
func (m *PullZone) SetOriginRetry5XXResponses(value *bool)() {
    m.originRetry5XXResponses = value
}
// SetOriginRetryConnectionTimeout sets the OriginRetryConnectionTimeout property value. Determines if we should retry the request in case of a connection timeout.
func (m *PullZone) SetOriginRetryConnectionTimeout(value *bool)() {
    m.originRetryConnectionTimeout = value
}
// SetOriginRetryDelay sets the OriginRetryDelay property value. Determines the amount of time that the CDN should wait before retrying an origin request.
func (m *PullZone) SetOriginRetryDelay(value *int32)() {
    m.originRetryDelay = value
}
// SetOriginRetryResponseTimeout sets the OriginRetryResponseTimeout property value. Determines if we should retry the request in case of a response timeout.
func (m *PullZone) SetOriginRetryResponseTimeout(value *bool)() {
    m.originRetryResponseTimeout = value
}
// SetOriginShieldEnableConcurrencyLimit sets the OriginShieldEnableConcurrencyLimit property value. Determines if the origin shield concurrency limit is enabled.
func (m *PullZone) SetOriginShieldEnableConcurrencyLimit(value *bool)() {
    m.originShieldEnableConcurrencyLimit = value
}
// SetOriginShieldMaxConcurrentRequests sets the OriginShieldMaxConcurrentRequests property value. Determines the number of maximum concurrent requests allowed to the origin.
func (m *PullZone) SetOriginShieldMaxConcurrentRequests(value *int32)() {
    m.originShieldMaxConcurrentRequests = value
}
// SetOriginShieldMaxQueuedRequests sets the OriginShieldMaxQueuedRequests property value. Determines the max number of origin requests that will remain in the queue
func (m *PullZone) SetOriginShieldMaxQueuedRequests(value *int32)() {
    m.originShieldMaxQueuedRequests = value
}
// SetOriginShieldQueueMaxWaitTime sets the OriginShieldQueueMaxWaitTime property value. Determines the max queue wait time
func (m *PullZone) SetOriginShieldQueueMaxWaitTime(value *int32)() {
    m.originShieldQueueMaxWaitTime = value
}
// SetOriginShieldZoneCode sets the OriginShieldZoneCode property value. The zone code of the origin shield
func (m *PullZone) SetOriginShieldZoneCode(value *string)() {
    m.originShieldZoneCode = value
}
// SetOriginType sets the OriginType property value. The OriginType property
func (m *PullZone) SetOriginType(value *float64)() {
    m.originType = value
}
// SetOriginUrl sets the OriginUrl property value. The origin URL of the pull zone where the files are fetched from.
func (m *PullZone) SetOriginUrl(value *string)() {
    m.originUrl = value
}
// SetPermaCacheStorageZoneId sets the PermaCacheStorageZoneId property value. The IP of the storage zone used for Perma-Cache
func (m *PullZone) SetPermaCacheStorageZoneId(value *int64)() {
    m.permaCacheStorageZoneId = value
}
// SetPermaCacheType sets the PermaCacheType property value. The PermaCacheType property
func (m *PullZone) SetPermaCacheType(value *int64)() {
    m.permaCacheType = value
}
// SetPreloadingScreenCode sets the PreloadingScreenCode property value. The custom preloading screen code
func (m *PullZone) SetPreloadingScreenCode(value *string)() {
    m.preloadingScreenCode = value
}
// SetPreloadingScreenCodeEnabled sets the PreloadingScreenCodeEnabled property value. Determines if the custom preloader screen is enabled
func (m *PullZone) SetPreloadingScreenCodeEnabled(value *bool)() {
    m.preloadingScreenCodeEnabled = value
}
// SetPreloadingScreenDelay sets the PreloadingScreenDelay property value. The delay in milliseconds after which the preloading screen will be displayed
func (m *PullZone) SetPreloadingScreenDelay(value *int32)() {
    m.preloadingScreenDelay = value
}
// SetPreloadingScreenEnabled sets the PreloadingScreenEnabled property value. Determines if the preloading screen is currently enabled
func (m *PullZone) SetPreloadingScreenEnabled(value *bool)() {
    m.preloadingScreenEnabled = value
}
// SetPreloadingScreenLogoUrl sets the PreloadingScreenLogoUrl property value. The preloading screen logo URL
func (m *PullZone) SetPreloadingScreenLogoUrl(value *string)() {
    m.preloadingScreenLogoUrl = value
}
// SetPreloadingScreenShowOnFirstVisit sets the PreloadingScreenShowOnFirstVisit property value. The PreloadingScreenShowOnFirstVisit property
func (m *PullZone) SetPreloadingScreenShowOnFirstVisit(value *bool)() {
    m.preloadingScreenShowOnFirstVisit = value
}
// SetPreloadingScreenTheme sets the PreloadingScreenTheme property value. The PreloadingScreenTheme property
func (m *PullZone) SetPreloadingScreenTheme(value *float64)() {
    m.preloadingScreenTheme = value
}
// SetPriceOverride sets the PriceOverride property value. The custom price override for this zone
func (m *PullZone) SetPriceOverride(value *float64)() {
    m.priceOverride = value
}
// SetQueryStringVaryParameters sets the QueryStringVaryParameters property value. Contains the list of vary parameters that will be used for vary cache by query string. If empty, all parameters will be used to construct the key
func (m *PullZone) SetQueryStringVaryParameters(value []string)() {
    m.queryStringVaryParameters = value
}
// SetRequestCoalescingTimeout sets the RequestCoalescingTimeout property value. Determines the lock time for coalesced requests.
func (m *PullZone) SetRequestCoalescingTimeout(value *int32)() {
    m.requestCoalescingTimeout = value
}
// SetRequestLimit sets the RequestLimit property value. Max number of requests per IP per second
func (m *PullZone) SetRequestLimit(value *int32)() {
    m.requestLimit = value
}
// SetRoutingFilters sets the RoutingFilters property value. The list of routing filters enabled for this zone
func (m *PullZone) SetRoutingFilters(value []PullZone_RoutingFilters)() {
    m.routingFilters = value
}
// SetShieldDDosProtectionEnabled sets the ShieldDDosProtectionEnabled property value. The ShieldDDosProtectionEnabled property
func (m *PullZone) SetShieldDDosProtectionEnabled(value *bool)() {
    m.shieldDDosProtectionEnabled = value
}
// SetShieldDDosProtectionType sets the ShieldDDosProtectionType property value. The ShieldDDosProtectionType property
func (m *PullZone) SetShieldDDosProtectionType(value *float64)() {
    m.shieldDDosProtectionType = value
}
// SetSouthAmericaDiscount sets the SouthAmericaDiscount property value. The Pull Zone specific pricing discount for South America region.
func (m *PullZone) SetSouthAmericaDiscount(value *int32)() {
    m.southAmericaDiscount = value
}
// SetStorageZoneId sets the StorageZoneId property value. The ID of the storage zone that the pull zone is linked to
func (m *PullZone) SetStorageZoneId(value *int64)() {
    m.storageZoneId = value
}
// SetSuspended sets the Suspended property value. The Suspended property
func (m *PullZone) SetSuspended(value *bool)() {
    m.suspended = value
}
// SetTypeEscaped sets the Type property value. The Type property
func (m *PullZone) SetTypeEscaped(value *float64)() {
    m.typeEscaped = value
}
// SetUseBackgroundUpdate sets the UseBackgroundUpdate property value. Determines if cache update is performed in the background.
func (m *PullZone) SetUseBackgroundUpdate(value *bool)() {
    m.useBackgroundUpdate = value
}
// SetUserId sets the UserId property value. The UserId property
func (m *PullZone) SetUserId(value *string)() {
    m.userId = value
}
// SetUseStaleWhileOffline sets the UseStaleWhileOffline property value. Determines if we should use stale cache while the origin is offline
func (m *PullZone) SetUseStaleWhileOffline(value *bool)() {
    m.useStaleWhileOffline = value
}
// SetUseStaleWhileUpdating sets the UseStaleWhileUpdating property value. Determines if we should use stale cache while cache is updating
func (m *PullZone) SetUseStaleWhileUpdating(value *bool)() {
    m.useStaleWhileUpdating = value
}
// SetVerifyOriginSSL sets the VerifyOriginSSL property value. Determines if the Pull Zone should verify the origin SSL certificate
func (m *PullZone) SetVerifyOriginSSL(value *bool)() {
    m.verifyOriginSSL = value
}
// SetVideoLibraryId sets the VideoLibraryId property value. The ID of the video library that the zone is linked to
func (m *PullZone) SetVideoLibraryId(value *int64)() {
    m.videoLibraryId = value
}
// SetZoneSecurityEnabled sets the ZoneSecurityEnabled property value. True if the URL secure token authentication security is enabled
func (m *PullZone) SetZoneSecurityEnabled(value *bool)() {
    m.zoneSecurityEnabled = value
}
// SetZoneSecurityIncludeHashRemoteIP sets the ZoneSecurityIncludeHashRemoteIP property value. True if the zone security hash should include the remote IP
func (m *PullZone) SetZoneSecurityIncludeHashRemoteIP(value *bool)() {
    m.zoneSecurityIncludeHashRemoteIP = value
}
// SetZoneSecurityKey sets the ZoneSecurityKey property value. The security key used for secure URL token authentication
func (m *PullZone) SetZoneSecurityKey(value *string)() {
    m.zoneSecurityKey = value
}
type PullZoneable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessControlOriginHeaderExtensions()([]string)
    GetAddCanonicalHeader()(*bool)
    GetAddHostHeader()(*bool)
    GetAfricaDiscount()(*int32)
    GetAllowedReferrers()([]string)
    GetAsiaOceaniaDiscount()(*int32)
    GetAWSSigningEnabled()(*bool)
    GetAWSSigningKey()(*string)
    GetAWSSigningRegionName()(*string)
    GetAWSSigningSecret()(*string)
    GetBlockedCountries()([]string)
    GetBlockedIps()([]string)
    GetBlockedReferrers()([]string)
    GetBlockNoneReferrer()(*bool)
    GetBlockPostRequests()(*bool)
    GetBlockRootPathAccess()(*bool)
    GetBudgetRedirectedCountries()([]string)
    GetBunnyAiImageBlueprints()([]BunnyAiImageBlueprintable)
    GetBurstSize()(*int32)
    GetCacheControlMaxAgeOverride()(*int64)
    GetCacheControlPublicMaxAgeOverride()(*int64)
    GetCacheErrorResponses()(*bool)
    GetCacheVersion()(*float64)
    GetCnameDomain()(*string)
    GetConnectionLimitPerIPCount()(*int32)
    GetCookieVaryParameters()([]string)
    GetDisableCookies()(*bool)
    GetDisableLetsEncrypt()(*bool)
    GetDnsRecordId()(*int64)
    GetDnsRecordValue()(*string)
    GetDnsZoneId()(*int64)
    GetEdgeScriptExecutionPhase()(*float64)
    GetEdgeScriptId()(*int64)
    GetEnableAccessControlOriginHeader()(*bool)
    GetEnableAutoSSL()(*bool)
    GetEnableAvifVary()(*bool)
    GetEnableBunnyImageAi()(*bool)
    GetEnableCacheSlice()(*bool)
    GetEnableCookieVary()(*bool)
    GetEnableCountryCodeVary()(*bool)
    GetEnabled()(*bool)
    GetEnableGeoZoneAF()(*bool)
    GetEnableGeoZoneASIA()(*bool)
    GetEnableGeoZoneEU()(*bool)
    GetEnableGeoZoneSA()(*bool)
    GetEnableGeoZoneUS()(*bool)
    GetEnableHostnameVary()(*bool)
    GetEnableLogging()(*bool)
    GetEnableMobileVary()(*bool)
    GetEnableOriginShield()(*bool)
    GetEnableQueryStringOrdering()(*bool)
    GetEnableRequestCoalescing()(*bool)
    GetEnableSafeHop()(*bool)
    GetEnableSmartCache()(*bool)
    GetEnableTLS1()(*bool)
    GetEnableTLS11()(*bool)
    GetEnableWebPVary()(*bool)
    GetErrorPageCustomCode()(*string)
    GetErrorPageEnableCustomCode()(*bool)
    GetErrorPageEnableStatuspageWidget()(*bool)
    GetErrorPageStatuspageCode()(*string)
    GetErrorPageWhitelabel()(*bool)
    GetEUUSDiscount()(*int32)
    GetFollowRedirects()(*bool)
    GetHostnames()([]Hostnameable)
    GetId()(*int64)
    GetIgnoreQueryStrings()(*bool)
    GetLimitRateAfter()(*float64)
    GetLimitRatePerSecond()(*int32)
    GetLogAnonymizationType()(*float64)
    GetLogFormat()(*float64)
    GetLogForwardingEnabled()(*bool)
    GetLogForwardingFormat()(*float64)
    GetLogForwardingHostname()(*string)
    GetLogForwardingPort()(*int32)
    GetLogForwardingProtocol()(*float64)
    GetLogForwardingToken()(*string)
    GetLoggingIPAnonymizationEnabled()(*bool)
    GetLoggingSaveToStorage()(*bool)
    GetLoggingStorageZoneId()(*int64)
    GetMagicContainersAppId()(*string)
    GetMagicContainersEndpointId()(*int64)
    GetMiddlewareScriptId()(*int64)
    GetMonthlyBandwidthLimit()(*int64)
    GetMonthlyBandwidthUsed()(*int64)
    GetMonthlyCharges()(*float64)
    GetName()(*string)
    GetOptimizerAutomaticOptimizationEnabled()(*bool)
    GetOptimizerDesktopMaxWidth()(*int32)
    GetOptimizerEnabled()(*bool)
    GetOptimizerEnableManipulationEngine()(*bool)
    GetOptimizerEnableWebP()(*bool)
    GetOptimizerForceClasses()(*bool)
    GetOptimizerImageQuality()(*int32)
    GetOptimizerMinifyCSS()(*bool)
    GetOptimizerMinifyJavaScript()(*bool)
    GetOptimizerMobileImageQuality()(*int32)
    GetOptimizerMobileMaxWidth()(*int32)
    GetOptimizerStaticHtmlEnabled()(*bool)
    GetOptimizerStaticHtmlWordPressBypassCookie()(*string)
    GetOptimizerStaticHtmlWordPressPath()(*string)
    GetOptimizerWatermarkEnabled()(*bool)
    GetOptimizerWatermarkMinImageSize()(*int32)
    GetOptimizerWatermarkOffset()(*float64)
    GetOptimizerWatermarkPosition()(*float64)
    GetOptimizerWatermarkUrl()(*string)
    GetOriginConnectTimeout()(*int32)
    GetOriginHostHeader()(*string)
    GetOriginLinkValue()(*string)
    GetOriginResponseTimeout()(*int32)
    GetOriginRetries()(*int32)
    GetOriginRetry5XXResponses()(*bool)
    GetOriginRetryConnectionTimeout()(*bool)
    GetOriginRetryDelay()(*int32)
    GetOriginRetryResponseTimeout()(*bool)
    GetOriginShieldEnableConcurrencyLimit()(*bool)
    GetOriginShieldMaxConcurrentRequests()(*int32)
    GetOriginShieldMaxQueuedRequests()(*int32)
    GetOriginShieldQueueMaxWaitTime()(*int32)
    GetOriginShieldZoneCode()(*string)
    GetOriginType()(*float64)
    GetOriginUrl()(*string)
    GetPermaCacheStorageZoneId()(*int64)
    GetPermaCacheType()(*int64)
    GetPreloadingScreenCode()(*string)
    GetPreloadingScreenCodeEnabled()(*bool)
    GetPreloadingScreenDelay()(*int32)
    GetPreloadingScreenEnabled()(*bool)
    GetPreloadingScreenLogoUrl()(*string)
    GetPreloadingScreenShowOnFirstVisit()(*bool)
    GetPreloadingScreenTheme()(*float64)
    GetPriceOverride()(*float64)
    GetQueryStringVaryParameters()([]string)
    GetRequestCoalescingTimeout()(*int32)
    GetRequestLimit()(*int32)
    GetRoutingFilters()([]PullZone_RoutingFilters)
    GetShieldDDosProtectionEnabled()(*bool)
    GetShieldDDosProtectionType()(*float64)
    GetSouthAmericaDiscount()(*int32)
    GetStorageZoneId()(*int64)
    GetSuspended()(*bool)
    GetTypeEscaped()(*float64)
    GetUseBackgroundUpdate()(*bool)
    GetUserId()(*string)
    GetUseStaleWhileOffline()(*bool)
    GetUseStaleWhileUpdating()(*bool)
    GetVerifyOriginSSL()(*bool)
    GetVideoLibraryId()(*int64)
    GetZoneSecurityEnabled()(*bool)
    GetZoneSecurityIncludeHashRemoteIP()(*bool)
    GetZoneSecurityKey()(*string)
    SetAccessControlOriginHeaderExtensions(value []string)()
    SetAddCanonicalHeader(value *bool)()
    SetAddHostHeader(value *bool)()
    SetAfricaDiscount(value *int32)()
    SetAllowedReferrers(value []string)()
    SetAsiaOceaniaDiscount(value *int32)()
    SetAWSSigningEnabled(value *bool)()
    SetAWSSigningKey(value *string)()
    SetAWSSigningRegionName(value *string)()
    SetAWSSigningSecret(value *string)()
    SetBlockedCountries(value []string)()
    SetBlockedIps(value []string)()
    SetBlockedReferrers(value []string)()
    SetBlockNoneReferrer(value *bool)()
    SetBlockPostRequests(value *bool)()
    SetBlockRootPathAccess(value *bool)()
    SetBudgetRedirectedCountries(value []string)()
    SetBunnyAiImageBlueprints(value []BunnyAiImageBlueprintable)()
    SetBurstSize(value *int32)()
    SetCacheControlMaxAgeOverride(value *int64)()
    SetCacheControlPublicMaxAgeOverride(value *int64)()
    SetCacheErrorResponses(value *bool)()
    SetCacheVersion(value *float64)()
    SetCnameDomain(value *string)()
    SetConnectionLimitPerIPCount(value *int32)()
    SetCookieVaryParameters(value []string)()
    SetDisableCookies(value *bool)()
    SetDisableLetsEncrypt(value *bool)()
    SetDnsRecordId(value *int64)()
    SetDnsRecordValue(value *string)()
    SetDnsZoneId(value *int64)()
    SetEdgeScriptExecutionPhase(value *float64)()
    SetEdgeScriptId(value *int64)()
    SetEnableAccessControlOriginHeader(value *bool)()
    SetEnableAutoSSL(value *bool)()
    SetEnableAvifVary(value *bool)()
    SetEnableBunnyImageAi(value *bool)()
    SetEnableCacheSlice(value *bool)()
    SetEnableCookieVary(value *bool)()
    SetEnableCountryCodeVary(value *bool)()
    SetEnabled(value *bool)()
    SetEnableGeoZoneAF(value *bool)()
    SetEnableGeoZoneASIA(value *bool)()
    SetEnableGeoZoneEU(value *bool)()
    SetEnableGeoZoneSA(value *bool)()
    SetEnableGeoZoneUS(value *bool)()
    SetEnableHostnameVary(value *bool)()
    SetEnableLogging(value *bool)()
    SetEnableMobileVary(value *bool)()
    SetEnableOriginShield(value *bool)()
    SetEnableQueryStringOrdering(value *bool)()
    SetEnableRequestCoalescing(value *bool)()
    SetEnableSafeHop(value *bool)()
    SetEnableSmartCache(value *bool)()
    SetEnableTLS1(value *bool)()
    SetEnableTLS11(value *bool)()
    SetEnableWebPVary(value *bool)()
    SetErrorPageCustomCode(value *string)()
    SetErrorPageEnableCustomCode(value *bool)()
    SetErrorPageEnableStatuspageWidget(value *bool)()
    SetErrorPageStatuspageCode(value *string)()
    SetErrorPageWhitelabel(value *bool)()
    SetEUUSDiscount(value *int32)()
    SetFollowRedirects(value *bool)()
    SetHostnames(value []Hostnameable)()
    SetId(value *int64)()
    SetIgnoreQueryStrings(value *bool)()
    SetLimitRateAfter(value *float64)()
    SetLimitRatePerSecond(value *int32)()
    SetLogAnonymizationType(value *float64)()
    SetLogFormat(value *float64)()
    SetLogForwardingEnabled(value *bool)()
    SetLogForwardingFormat(value *float64)()
    SetLogForwardingHostname(value *string)()
    SetLogForwardingPort(value *int32)()
    SetLogForwardingProtocol(value *float64)()
    SetLogForwardingToken(value *string)()
    SetLoggingIPAnonymizationEnabled(value *bool)()
    SetLoggingSaveToStorage(value *bool)()
    SetLoggingStorageZoneId(value *int64)()
    SetMagicContainersAppId(value *string)()
    SetMagicContainersEndpointId(value *int64)()
    SetMiddlewareScriptId(value *int64)()
    SetMonthlyBandwidthLimit(value *int64)()
    SetMonthlyBandwidthUsed(value *int64)()
    SetMonthlyCharges(value *float64)()
    SetName(value *string)()
    SetOptimizerAutomaticOptimizationEnabled(value *bool)()
    SetOptimizerDesktopMaxWidth(value *int32)()
    SetOptimizerEnabled(value *bool)()
    SetOptimizerEnableManipulationEngine(value *bool)()
    SetOptimizerEnableWebP(value *bool)()
    SetOptimizerForceClasses(value *bool)()
    SetOptimizerImageQuality(value *int32)()
    SetOptimizerMinifyCSS(value *bool)()
    SetOptimizerMinifyJavaScript(value *bool)()
    SetOptimizerMobileImageQuality(value *int32)()
    SetOptimizerMobileMaxWidth(value *int32)()
    SetOptimizerStaticHtmlEnabled(value *bool)()
    SetOptimizerStaticHtmlWordPressBypassCookie(value *string)()
    SetOptimizerStaticHtmlWordPressPath(value *string)()
    SetOptimizerWatermarkEnabled(value *bool)()
    SetOptimizerWatermarkMinImageSize(value *int32)()
    SetOptimizerWatermarkOffset(value *float64)()
    SetOptimizerWatermarkPosition(value *float64)()
    SetOptimizerWatermarkUrl(value *string)()
    SetOriginConnectTimeout(value *int32)()
    SetOriginHostHeader(value *string)()
    SetOriginLinkValue(value *string)()
    SetOriginResponseTimeout(value *int32)()
    SetOriginRetries(value *int32)()
    SetOriginRetry5XXResponses(value *bool)()
    SetOriginRetryConnectionTimeout(value *bool)()
    SetOriginRetryDelay(value *int32)()
    SetOriginRetryResponseTimeout(value *bool)()
    SetOriginShieldEnableConcurrencyLimit(value *bool)()
    SetOriginShieldMaxConcurrentRequests(value *int32)()
    SetOriginShieldMaxQueuedRequests(value *int32)()
    SetOriginShieldQueueMaxWaitTime(value *int32)()
    SetOriginShieldZoneCode(value *string)()
    SetOriginType(value *float64)()
    SetOriginUrl(value *string)()
    SetPermaCacheStorageZoneId(value *int64)()
    SetPermaCacheType(value *int64)()
    SetPreloadingScreenCode(value *string)()
    SetPreloadingScreenCodeEnabled(value *bool)()
    SetPreloadingScreenDelay(value *int32)()
    SetPreloadingScreenEnabled(value *bool)()
    SetPreloadingScreenLogoUrl(value *string)()
    SetPreloadingScreenShowOnFirstVisit(value *bool)()
    SetPreloadingScreenTheme(value *float64)()
    SetPriceOverride(value *float64)()
    SetQueryStringVaryParameters(value []string)()
    SetRequestCoalescingTimeout(value *int32)()
    SetRequestLimit(value *int32)()
    SetRoutingFilters(value []PullZone_RoutingFilters)()
    SetShieldDDosProtectionEnabled(value *bool)()
    SetShieldDDosProtectionType(value *float64)()
    SetSouthAmericaDiscount(value *int32)()
    SetStorageZoneId(value *int64)()
    SetSuspended(value *bool)()
    SetTypeEscaped(value *float64)()
    SetUseBackgroundUpdate(value *bool)()
    SetUserId(value *string)()
    SetUseStaleWhileOffline(value *bool)()
    SetUseStaleWhileUpdating(value *bool)()
    SetVerifyOriginSSL(value *bool)()
    SetVideoLibraryId(value *int64)()
    SetZoneSecurityEnabled(value *bool)()
    SetZoneSecurityIncludeHashRemoteIP(value *bool)()
    SetZoneSecurityKey(value *string)()
}
