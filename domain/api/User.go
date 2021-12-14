package api

type Sign_t struct {
    Id  string  `json:"id"  binding:"required"`
    Pw  string  `json:"pw"  binding:"required"`
}

type Preference_t struct {
    Budget      uint        `json:"budged"`
    Workspace   string      `json:"workspace"`
    Facilities  []string    `json:"facilities"`
    Favors      []string    `json:"favors"`
}
