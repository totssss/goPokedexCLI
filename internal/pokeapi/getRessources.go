package pokeapi

func (c *Client) GetLocation(locationName string) (Location, error) {
	return getJSON[Location](c, BaseLocationURL+"/"+locationName)
}

func (c *Client) GetLocationsPage(url string) (LocationPage, error) {
	return getJSON[LocationPage](c, url)
}
