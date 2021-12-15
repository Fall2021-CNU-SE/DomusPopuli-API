package api

type Preference_t struct {
    Budget      uint        `json:"budget"`
    Workspace   string      `json:"workspace"`
    Facilities  []string    `json:"facilities"`
    Favors      []string    `json:"favors"`
}
