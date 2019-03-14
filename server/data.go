package server

import (
	"fmt"

	"github.com/daxsorbito/golang-reactjs/webpack"
)

// User represents current user session
type User struct {
	Email     string
	FirstName string
	LastName  string
}

// ViewData contains data for the view
type ViewData struct {
	CurrentUser  User
	assetsMapper webpack.AssetsMapper
}

// NewViewData creates a new data for the view
func NewViewData(buildPath string) (ViewData, error) {
	assetsMapper, err := webpack.NewAssetsMapper(buildPath)
	fmt.Println(&assetsMapper)
	if err != nil {
		return ViewData{}, err
	}

	return ViewData{
		CurrentUser: User{
			Email:     "bill@example.com",
			FirstName: "Bill",
			LastName:  "Black",
		},
		assetsMapper: assetsMapper,
	}, nil
}

// Webpack maps file name to path
func (d ViewData) Webpack(file string) string {
	return d.assetsMapper(file)
}
