package category

type Meta_t struct {
    TotalCount          int             `json:"total_count"`        // 검색된 문서 수
    PageableCount       int             `json:"pageable_count"`     // total_count 중 노출 가능 문서 수 (최대값: 45)
    IsEnd               bool            `json:"is_end"`             // 현재 페이지가 마지막 페이지인지 여부
    SameName            RegionInfo_t    `json:"same_name"`
}

type RegionInfo_t struct {
    Region              []string        `json:"region"`             // 질의어에서 인식된 지역의 리스트
    Keyword             string          `json:"keyword"`            // 질의어에서 지역 정보를 제외한 키워드
    SelectedRegion      string          `json:"selected_region"`    // 인식된 지역 리스트 중, 현재 검색에 사용된 지역 정보
}

type Document_t struct {
    ID                  string          `json:"id"`                     // 장소 ID
    PlaceName           string          `json:"place_name"`             // 장소명, 업체명
    CategoryName        string          `json:"category_name"`          // 카테고리 이름
    CategoryGroupCode   string          `json:"category_group_code"`    // 중요 카테고리만 그룹핑한 카테고리 그룹 코드
    CategoryGroupName   string          `json:"category_group_name"`    // 중요 카테고리만 그룹핑한 카테고리 그룹명
    Phone               string          `json:"phone"`                  // 전화번호
    AddressName         string          `json:"address_name"`           // 전체 지번 주소
    RoadAddressName     string          `json:"road_address_name"`      // 전체 도로명 주소
    X                   string          `json:"x"`                      // X 좌표값, 경위도인 경우 longitude (경도)
    Y                   string          `json:"y"`                      // Y 좌표값, 경위도인 경우 latitude(위도)
    PlaceUrl            string          `json:"place_url"`              // 장소 상세페이지 URL
    Distance            string          `json:"distance"`               // 중심좌표까지의 거리 (단, x,y 파라미터를 준 경우에만 존재) 단위 meter
}

type Response_t struct {
    Meta                Meta_t          `json:"meta"`
    Documents           []Document_t    `json:"documents"`
}
