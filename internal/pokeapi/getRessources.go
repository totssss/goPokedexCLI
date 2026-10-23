package pokeapi

func (c *Client) GetLocation(locationName string) (Location, error) {
	return getJSON[Location](c, BaseLocationURL+"/"+locationName)
}

func (c *Client) GetLocationsPage(url string) (LocationPage, error) {
	return getJSON[LocationPage](c, url)
}

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	return getJSON[Pokemon](c, BasePokemonURL+pokemonName)
}
