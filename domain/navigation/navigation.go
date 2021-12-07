/************************************************************************
 * Kakao map API - navigation                                           *
 * docs: https://developers.kakao.com/docs/latest/en/kakaonavi/rest-api *
 ************************************************************************/

package navigation

type Response_t struct {
    // Unique identifier of a direction request.
    TransId     string      `json:"trans_id"`

    // Route data. If alternatives is set to true, one or more routes are returned.
    Routes      []Route_t   `json:"routes"`
}

type Route_t struct {
    // Result code of directions search.
    ResultCode  int         `json:"result_code"`

    // Result message of directions search.
    ResultMsg   string      `json:"result_msg"`

    // Summary information of directions.
    Summary     Summary_t   `json:"summary"`

    // Route data by section.
    //  If waypoints are specified, sections are created as many as {number of waypoints + 1}.
    //      (Example: If the number of waypoints is two, three sections are created.
    //      section1: origin → waypoint1
    //      section2: waypoint1 → waypoint2
    //      section3: waypoint2 → destination)
    Sections    []Section_t `json:"sections"`
}

type Summary_t struct {
    // Origin information.
    Origin      Origin_t        `json:"origin"`

    // Destination information.
    Destination Destination_t   `json:"destination"`

    // Waypoint information.
    Waypoints   []Waypoint_t    `json:"waypoints"`

    // Set priority options for navigation.
    Priority    string          `json:"priority"`

    // Geo-bounding box that includes all routes.
    Bound       Bound_t         `json:"bound"`

    // Total fare.
    Fare        Fare_t          `json:"fare"`

    // Total distance in meters.
    Distance    int	            `json:"distance"`

    // Travel time to the destination in seconds.
    Duration    int	            `json:"duration"`
}

type Section_t struct {
    // Distance between sections in meters.
    Distance    int         `json:"distance"`

    // Total required time in seconds.
    Duration    int         `json:"duration"`

    // Geo-bounding box that includes all routes.
    //  Only returned if summary is set to false.
    Bound       Bound_t     `json:"bound"`

    // Road information.
    //  Only returned if summary is set to false.
    Roads       []Road_t    `json:"roads"`

    // Guidance information.
    //  Only returned if summary is set to false.
    Guides      []Guide_t   `json:"guides"`
}

type Origin_t struct {
    // Origin name.
    Name    string  `json:"name"`

    // X coordinate (longitude).
    X       float64  `json:"x"`

    // Y coordinate (latitude).
    Y       float64 `json:"y"`
}

type Destination_t struct {
    // Destination name.
    Name    string  `json:"name"`

    // X coordinate (longitude).
    X       float64 `json:"x"`

    // Y coordinate (latitude).
    Y       float64 `json:"y"`
}

type Waypoint_t struct {
    // Waypoint name.
    Name    string  `json:"name"`

    // X coordinate (longitude).
    X       float64 `json:"x"`

    // Y coordinate (latitude).
    Y       float64 `json:"y"`
}

type Bound_t struct {
    // X coordinate of the bottom left corner of the geo-bounding box.
    MinX    float64 `json:"min_x"`

    // Y coordinate of the bottom left corner of the geo-bounding box.
    MinY    float64 `json:"min_y"`

    // X coordinate of the top right corner of the geo-bounding box.
    MaxX    float64 `json:"max_x"`

    // Y coordinate of the top right corner of the geo-bounding box.
    MaxY    float64 `json:"max_y"`
}

type Fare_t struct {
    // Taxi fare in South Korean won (KRW).
    Taxi    int `json:"taxi"`

    // Toll fare in South Korean won (KRW).
    Toll    int `json:"toll"`
}

type Road_t struct {
    // Road name.
    Name            string      `json:"name"`

    // Road distance in meters.
    Distance        int         `json:"distance"`

    // Estimated travel time in seconds, which is considered as the actual required time.
    Duration        int         `json:"duration"`

    // Current traffic speed in kilometers per hour.
    TrafficSpeed    float64     `json:"traffic_speed"`

    // Current traffic information.
    TrafficState    int         `json:"traffic_state"`

    // One-dimensional array that includes X and Y coordinates.
    //  (Example: [127.10966790676201, 37.394469584427156, 127.10967141980313, 37.39512739646385])
    Vertexes        []float64   `json:"vertexes"`
}

type Guide_t struct {
    // Name of guidance.
    Name        string      `json:"name"`

    // X coordinate (longitude).
    X           float64     `json:"x"`

    // Y coordinate (latitude).
    Y           float64     `json:"y"`

    // Distance from the previous guidance point to the current guidance point in meters.
    Distance    int         `json:"distance"`

    // Required time from the previous guidance point to the current guidance point in seconds.
    Duration    int         `json:"duration"`

    // Guidance type.
    Type        int         `json:"type"`

    // Guidance message.
    Guidance    string      `json:"guidance"`

    // Index of the corresponding guidance.
    RoadIndex   int         `json:"road_index"`
}
