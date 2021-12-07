package coords

type Meta_t struct {
    TotalCount      int     `json:"total_count"`        // 검색된 문서 수
    PageableCount   int     `json:"pageable_count"`     // total_count 중 노출 가능 문서 수 (최대값: 45)
    IsEnd           bool    `json:"is_end"`             // 현재 페이지가 마지막 페이지인지 여부
        //값이 false면 다음 요청 시 page 값을 증가시켜 다음 페이지 요청 가능
}

type Address_t struct {
    AddressName         string  `json:"address_name"`           // 전체 지번 주소
    Region1DepthName    string  `json:"region_1depth_name"`     // 지역 1 Depth, 시도 단위
    Region2DepthName    string  `json:"region_2depth_name"`     // 지역 2 Depth, 구 단위
    Region3DepthName    string  `json:"region_3depth_name"`     // 지역 3 Depth, 동 단위
    Region3DepthHName   string  `json:"region_3depth_h_name"`   // 지역 3 Depth, 행정동 명칭
    HCode               string  `json:"h_code"`                 // 행정 코드
    BCode               string  `json:"b_code"`                 // 법정 코드
    MountainYN          string  `json:"mountain_yn"`            // 산 여부, Y 또는 N
    MainAddressNo       string  `json:"main_address_no"`        // 지번 주번지
    SubAddressNo        string  `json:"sub_address_no"`         // 지번 부번지, 없을 경우 빈 문자열("") 반환
    ZipCode             string  `json:"zip_code"`               // Deprecated 우편번호(6자리)
    X                   string  `json:"x"`                      // X 좌표값, 경위도인 경우 경도(longitude)
    Y                   string  `json:"y"`                      // Y 좌표값, 경위도인 경우 위도(latitude)
}

type RoadAddress_t struct {
    AddressName         string  `json:"address_name"`           // 전체 지번 주소
    Region1DepthName    string  `json:"region_1depth_name"`     // 지역 1 Depth, 시도 단위
    Region2DepthName    string  `json:"region_2depth_name"`     // 지역 2 Depth, 구 단위
    Region3DepthName    string  `json:"region_3depth_name"`     // 지역 3 Depth, 동 단위
    RoadName            string  `json:"road_name"`              // 도로명
    UndergroundYN       string  `json:"underground_yn"`         // 지하 여부, Y 또는 N
    MainBuildingNo      string  `json:"main_building_no"`       // 건물 본번
    SubBuildingNo       string  `json:"sub_building_no"`        // 건물 부번, 없을 경우 빈 문자열("") 반환
    BuildingName        string  `json:"building_name"`          // 건물 이름
    NoneNo              string  `json:"zone_no"`                // 우편번호(5자리)
    X                   string  `json:"x"`                      // X 좌표값, 경위도인 경우 경도(longitude)
    Y                   string  `json:"y"`                      // Y 좌표값, 경위도인 경우 위도(latitude)
}

type Document_t struct {
    AddressName     string          `json:"address_name"`       // 전체 지번 주소 또는 전체 도로명 주소, 입력에 따라 결정됨
    AddressType     string          `json:"address_type"`       // address_name의 값의 타입
        // "REGION"(지명) "ROAD"(도로명) "REGION_ADDR"(지번 주소) "ROAD_ADDR"(도로명 주소)
    X               string          `json:"x"`                  // X 좌표값, 경위도인 경우 경도(longitude)
    Y               string          `json:"y"`                  // Y 좌표값, 경위도인 경우 위도(latitude)
    Address         Address_t       `json:"addressAddress"`     // 지번 주소 상세 정보, 아래 Address 참고
    RoadAddress     RoadAddress_t   `json:"road_address"`       // 도로명 주소 상세 정보, 아래 RoadAaddress 참고
}

type Response_t struct {
    Meta        Meta_t              `json:"meta"`
    Documents   []Document_t        `json:"documents"`
}
