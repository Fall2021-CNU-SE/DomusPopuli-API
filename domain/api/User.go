package api

type Preference_t struct {
    Budget      uint        `json:"budged"`
    Workspace   string      `json:"workspace"`
    Facilities  []string    `json:"facilities"`
    Favors      []string    `json:"favors"`
}
