package api

type MakeHouse_t struct {
    Token           string  `json:"token" binding:"required"`
    Name            string  `json:"name"`
    Address1        string  `json:"address1"`
    Address2        string  `json:"address2"`
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
