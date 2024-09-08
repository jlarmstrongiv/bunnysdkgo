package optimizer
type Properties_crop_gravity int

const (
    CENTER_PROPERTIES_CROP_GRAVITY Properties_crop_gravity = iota
    FORGET_PROPERTIES_CROP_GRAVITY
    EAST_PROPERTIES_CROP_GRAVITY
    NORTH_PROPERTIES_CROP_GRAVITY
    SOUTH_PROPERTIES_CROP_GRAVITY
    WEST_PROPERTIES_CROP_GRAVITY
    NORTHEAST_PROPERTIES_CROP_GRAVITY
    NORTHWEST_PROPERTIES_CROP_GRAVITY
    SOUTHEAST_PROPERTIES_CROP_GRAVITY
    SOUTHWEST_PROPERTIES_CROP_GRAVITY
)

func (i Properties_crop_gravity) String() string {
    return []string{"center", "forget", "east", "north", "south", "west", "northeast", "northwest", "southeast", "southwest"}[i]
}
func ParseProperties_crop_gravity(v string) (any, error) {
    result := CENTER_PROPERTIES_CROP_GRAVITY
    switch v {
        case "center":
            result = CENTER_PROPERTIES_CROP_GRAVITY
        case "forget":
            result = FORGET_PROPERTIES_CROP_GRAVITY
        case "east":
            result = EAST_PROPERTIES_CROP_GRAVITY
        case "north":
            result = NORTH_PROPERTIES_CROP_GRAVITY
        case "south":
            result = SOUTH_PROPERTIES_CROP_GRAVITY
        case "west":
            result = WEST_PROPERTIES_CROP_GRAVITY
        case "northeast":
            result = NORTHEAST_PROPERTIES_CROP_GRAVITY
        case "northwest":
            result = NORTHWEST_PROPERTIES_CROP_GRAVITY
        case "southeast":
            result = SOUTHEAST_PROPERTIES_CROP_GRAVITY
        case "southwest":
            result = SOUTHWEST_PROPERTIES_CROP_GRAVITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeProperties_crop_gravity(values []Properties_crop_gravity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Properties_crop_gravity) isMultiValue() bool {
    return false
}
