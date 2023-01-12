package model

// 16line
type House struct {
	Id                    int    `json:"id"`
	LicenseNumber         string `json:"license_number" binding:"required"`
	UserId                int    `json:"user_id" binding:"required"`
	City                  string `json:"city" binding:"required"`
	Region                string `json:"region" binding:"required"`
	Title                 string `json:"title" binding:"required"`
	Imgs                  string `json:"imgs" binding:"required"`
	HouseType             string `json:"house_type" binding:"required"`
	Storey                int    `json:"storey" binding:"required"`
	TotalStorey           int    `json:"total_storey" binding:"required"`
	Decoration            string `json:"decoration" binding:"required"`
	IsElevator            bool   `json:"isElevator"`
	Address               string `json:"address" binding:"required"`
	Ownership_certificate string `json:"ownership_certificate" binding:"required"`
	State                 int    `json:"state"`
	IsAuth                bool   `json:"isAuth"`
	// Id                    int    `json:"id"`
	// LicenseNumber         string `json:"license_number" binding:"required"`
	// UserId                int    `json:"user_id"`
	// City                  string `json:"city"`
	// Region                string `json:"region"`
	// Title                 string `json:"title"`
	// Imgs                  string `json:"imgs"`
	// HouseType             string `json:"house_type"`
	// Storey                int    `json:"storey"`
	// TotalStorey           int    `json:"total_storey"`
	// Decoration            string `json:"decoration"`
	// IsElevator            bool   `json:"isElevator"`
	// Address               string `json:"address"`
	// Ownership_certificate string `json:"ownership_certificate"`
	// IsAuth                bool   `json:"isAuth"`
	// State                 int    `json:"state"`
}

// func NewHouse(house House) *House {
// 	return &House{
// 		LicenseNumber:         house.LicenseNumber,
// 		UserId:                house.UserId,
// 		City:                  house.City,
// 		Region:                house.Region,
// 		Title:                 house.Title,
// 		Imgs:                  house.Imgs,
// 		HouseType:             house.HouseType,
// 		Storey:                house.Storey,
// 		TotalStorey:           house.TotalStorey,
// 		Decoration:            house.Decoration,
// 		IsElevator:            house.IsElevator,
// 		Ownership_certificate: house.Ownership_certificate,
// 		Address:               house.Address,
// 		IsAuth:                true,
// 		State:                 1,
// 	}
// }
