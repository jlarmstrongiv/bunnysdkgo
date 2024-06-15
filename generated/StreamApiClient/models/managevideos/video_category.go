package managevideos
// The automatically detected category of the video
type Video_category int

const (
    ADULT_VIDEO_CATEGORY Video_category = iota
    ANIMALSBIRDS_VIDEO_CATEGORY
    ANIMALSCATS_VIDEO_CATEGORY
    ANIMALSDOGS_VIDEO_CATEGORY
    ANIMATED_VIDEO_CATEGORY
    ANIME_VIDEO_CATEGORY
    GAMING_VIDEO_CATEGORY
    GRAPHICS_VIDEO_CATEGORY
    ILLUSTRATIONS_VIDEO_CATEGORY
    MOVIE_VIDEO_CATEGORY
    OTHER_VIDEO_CATEGORY
    OTHERPEOPLE_VIDEO_CATEGORY
    SPORTSBASKETBALL_VIDEO_CATEGORY
    SPORTSHOCKEY_VIDEO_CATEGORY
    SPORTSOTHER_VIDEO_CATEGORY
    SPORTSRACING_VIDEO_CATEGORY
    SPORTSSOCCER_VIDEO_CATEGORY
    SPORTSTENNIS_VIDEO_CATEGORY
    UNKNOWN_VIDEO_CATEGORY
)

func (i Video_category) String() string {
    return []string{"Adult", "Animals Birds", "Animals Cats", "Animals Dogs", "Animated", "Anime", "Gaming", "Graphics", "Illustrations", "Movie", "Other", "Other People", "Sports Basketball", "Sports Hockey", "Sports Other", "Sports Racing", "Sports Soccer", "Sports Tennis", "unknown"}[i]
}
func ParseVideo_category(v string) (any, error) {
    result := ADULT_VIDEO_CATEGORY
    switch v {
        case "Adult":
            result = ADULT_VIDEO_CATEGORY
        case "Animals Birds":
            result = ANIMALSBIRDS_VIDEO_CATEGORY
        case "Animals Cats":
            result = ANIMALSCATS_VIDEO_CATEGORY
        case "Animals Dogs":
            result = ANIMALSDOGS_VIDEO_CATEGORY
        case "Animated":
            result = ANIMATED_VIDEO_CATEGORY
        case "Anime":
            result = ANIME_VIDEO_CATEGORY
        case "Gaming":
            result = GAMING_VIDEO_CATEGORY
        case "Graphics":
            result = GRAPHICS_VIDEO_CATEGORY
        case "Illustrations":
            result = ILLUSTRATIONS_VIDEO_CATEGORY
        case "Movie":
            result = MOVIE_VIDEO_CATEGORY
        case "Other":
            result = OTHER_VIDEO_CATEGORY
        case "Other People":
            result = OTHERPEOPLE_VIDEO_CATEGORY
        case "Sports Basketball":
            result = SPORTSBASKETBALL_VIDEO_CATEGORY
        case "Sports Hockey":
            result = SPORTSHOCKEY_VIDEO_CATEGORY
        case "Sports Other":
            result = SPORTSOTHER_VIDEO_CATEGORY
        case "Sports Racing":
            result = SPORTSRACING_VIDEO_CATEGORY
        case "Sports Soccer":
            result = SPORTSSOCCER_VIDEO_CATEGORY
        case "Sports Tennis":
            result = SPORTSTENNIS_VIDEO_CATEGORY
        case "unknown":
            result = UNKNOWN_VIDEO_CATEGORY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVideo_category(values []Video_category) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Video_category) isMultiValue() bool {
    return false
}
