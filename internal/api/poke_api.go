package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) GetLocationAreas(pageURL *string) (RespShallowLocations, error) {
	requestURL, err := url.JoinPath(BaseURL, "location-area")
	if err != nil {
		return RespShallowLocations{}, err
	}

	if pageURL != nil {
		requestURL = *pageURL
	}

	locationsResp, err := GetResponseData[RespShallowLocations](c, requestURL)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}

func (c *Client) GetAreaData(areaName string) (AreaData, error) {
	requestURL, err := url.JoinPath(BaseURL, "location-area", areaName)
	if err != nil {
		return AreaData{}, err
	}

	areaData, err := GetResponseData[AreaData](c, requestURL)
	if err != nil {
		return AreaData{}, err
	}

	return areaData, nil
}

func (c *Client) GetPokemonData(pokemonName string) (PokemonData, error) {
	requestURL, err := url.JoinPath(BaseURL, "pokemon", pokemonName)
	if err != nil {
		return PokemonData{}, err
	}

	pokemonData, err := GetResponseData[PokemonData](c, requestURL)
	if err != nil {
		return PokemonData{}, err
	}

	return pokemonData, nil

}

func GetResponseData[T any](client *Client, reqURL string) (T, error) {
	var dataType T
	if val, ok := client.cache.Get(reqURL); ok {
		err := json.Unmarshal(val, &dataType)
		if err != nil {
			return dataType, err
		}

		return dataType, nil
	}

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return dataType, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return dataType, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return dataType, fmt.Errorf("bad status %d from %s", resp.StatusCode, reqURL)
	}

	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return dataType, err
	}

	err = json.Unmarshal(raw, &dataType)
	if err != nil {
		return dataType, err
	}

	client.cache.Add(reqURL, raw)

	return dataType, nil
}
