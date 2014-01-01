package foursquare

// https://developer.foursquare.com/docs/responses/venue
type Venue struct {
	Id         string
	Name       string
	Location   Location
	Categories []Category
}

type Location struct {
	Address     string
	CrossStreet string
	City        string
	State       string
	PostalCode  string
	Country     string
	Lat         float64
	Lng         float64
	Distance    float64
}

type Category struct {
	Id         string
	Name       string
	PluralName string
	ShortName  string
	Primary    bool
}
