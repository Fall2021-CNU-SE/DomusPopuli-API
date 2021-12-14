package api

type MakeHouse_t struct {
    Name            string  `json:"name"`
    Address1        string  `json:"address1"`
    Address2        string  `json:"address2"`
    Deposit         uint    `json:"deposit"`
    Maintenance     uint    `json:"maintenance"`
    RentalFee       uint    `json:"rentalfee"`
    RoomSize        uint    `json:"roomsize"`
    LeaseOrRent     bool    `json:"leaseorrent"`
}
