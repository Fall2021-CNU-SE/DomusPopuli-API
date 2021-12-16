package api

type Preference_t struct {
    Token       string      `json:"token" binding:"required"`
    Budget      uint        `json:"budget"`
    Workspace   string      `json:"workspace"`
    Facilities  []string    `json:"facilities"`
    Favors      []string    `json:"favors"`
}
