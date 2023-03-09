package dto

type FilterPagination struct {
	Search         string `json:"search" form:"search"`
	SortBy         string `json:"sortBy" form:"sortBy"`
	OrderBy        string `json:"orderBy" form:"orderBy"`
	Page           uint32 `json:"page" form:"page"`
	PerPage        uint32 `json:"perPage" form:"perPage"`
	UserID         string `json:"user_id" form:"user_id"`
	HotelID        string `json:"hotel_id" form:"hotel_id"`
	ProvinceID     string `json:"provinceId" form:"provinceId"`
	CityID         string `json:"cityId" form:"cityId"`
	DistrictID     string `json:"districtId" form:"districtId"`
	ClassID        string `json:"classId" form:"classId"`
	CategoryID     string `json:"categoryId" form:"categoryId"`
	SpeciesID      string `json:"speciesId" form:"speciesId"`
	GroupID        string `json:"groupId" form:"groupId"`
	CageTypeID     string `json:"cageTypeId" form:"cageTypeId"`
	CageCategoryID string `json:"cageCategoryId" form:"cageCategoryId"`
	CageDetailID   string `json:"cageDetailId" form:"cageDetailId"`
	ServiceID      string `json:"serviceId" form:"serviceId"`
}

type Pagination struct {
	Page      uint `json:"page"`
	PerPage   uint `json:"perPage"`
	TotalData uint `json:"totalData"`
	TotalPage uint `json:"totalPage"`
}
