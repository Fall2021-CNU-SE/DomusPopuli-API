package domain

type Coordinate_t struct {
    Long        string
    Lat         string
}

type CheckList_t struct {
    H0  uint    `json:"h0"`
    H1  uint    `json:"h1"`
    H2  uint    `json:"h2"`
    B0  uint    `json:"b0"`
    B1  uint    `json:"b1"`
    B2  uint    `json:"b2"`
    K0  uint    `json:"k0"`
    K1  uint    `json:"k1"`
    K2  uint    `json:"k2"`
    E0  uint    `json:"e0"`
    O0  uint    `json:"o0"`
    O1  uint    `json:"o1"`
    O2  uint    `json:"o2"`
}
