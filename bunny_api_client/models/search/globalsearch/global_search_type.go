package globalsearch
type GlobalSearchType int

const (
    CDN_GLOBALSEARCHTYPE GlobalSearchType = iota
    STORAGE_GLOBALSEARCHTYPE
    DNS_GLOBALSEARCHTYPE
    SCRIPT_GLOBALSEARCHTYPE
    STREAM_GLOBALSEARCHTYPE
)

func (i GlobalSearchType) String() string {
    return []string{"cdn", "storage", "dns", "script", "stream"}[i]
}
func ParseGlobalSearchType(v string) (any, error) {
    result := CDN_GLOBALSEARCHTYPE
    switch v {
        case "cdn":
            result = CDN_GLOBALSEARCHTYPE
        case "storage":
            result = STORAGE_GLOBALSEARCHTYPE
        case "dns":
            result = DNS_GLOBALSEARCHTYPE
        case "script":
            result = SCRIPT_GLOBALSEARCHTYPE
        case "stream":
            result = STREAM_GLOBALSEARCHTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGlobalSearchType(values []GlobalSearchType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GlobalSearchType) isMultiValue() bool {
    return false
}
