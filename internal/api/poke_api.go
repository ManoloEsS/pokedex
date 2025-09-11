package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type AreaData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationAreas(pageURL *string) (RespShallowLocations, error) {
	endpoint := "location-area"
	requestURL, err := url.JoinPath(BaseURL, endpoint)
	if err != nil {
		return RespShallowLocations{}, err
	}

	if pageURL != nil {
		requestURL = *pageURL
	}

	locationsResp, err := GetResponseData(c, requestURL, RespShallowLocations{})
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}

func (c *Client) GetAreaData(areaName string) (AreaData, error) {
	endpoint, err := url.JoinPath("location-area", areaName)
	if err != nil {
		return AreaData{}, err
	}

	requestURL, err := url.JoinPath(BaseURL, endpoint)
	if err != nil {
		return AreaData{}, err
	}

	areaData, err := GetResponseData(c, requestURL, AreaData{})
	if err != nil {
		return AreaData{}, err
	}

	return areaData, nil
}

func GetResponseData[T any](client *Client, reqURL string, dataType T) (T, error) {
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
