package api

type MakeHouse_t struct {
    Token           string  `json:"token" binding:"required"`
    Name            string  `json:"name"`
    Address         string  `json:"address"`
    Deposit         uint    `json:"deposit"`
    Maintenance     uint    `json:"maintenance"`
    RentalFee       uint    `json:"rentalfee"`
    RoomSize        uint    `json:"roomsize"`
    IsRent          bool    `json:"isrent"`
}

type HouseSummary_t struct {
    Name            string  `json:"name"`
    TotalScore      float64 `json:"totalScore"`
}

type HouseDetail_t struct {
    Error           string                  `json:"error"`
    TotalScore      float64                 `json:"totalScore"`
    ChecklistScore  float64                 `json:"checklistScore"`
    FacilityScore   float64                 `json:"facilityScore"`
    Facilities      []FacilitySummary_t     `json:"facilities"`
}

type FacilitySummary_t struct {
    Name            string  `json:"name"`
    Address         string  `json:"address"`
}
