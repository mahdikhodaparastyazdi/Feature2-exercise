package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Team struct {
	Name string
}

type Fplayer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
	Team string `json:"team`
}

var Persons = []Team{
	{Name: "Germany"},
	{Name: "England"},
	{Name: "France"},
	{Name: "Spain"},
	{Name: "Manchester United"},
	{Name: "Arsenal"},
	{Name: "Chelsea"},
	{Name: "Barcelona"},
	{Name: "Real Madrid"},
	{Name: "Bayern Munich"},
}

//Check team's api exist in our slice or not
func IsElementInSlice(searchstring string) bool {
	for _, v := range Persons {
		if v.Name == searchstring {
			return true
		}
	}
	return false
}

//Convert io.Reader to slice of bytes
func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

type Response struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   struct {
		Team struct {
			ID          int    `json:"id"`
			OptaID      int    `json:"optaId"`
			Country     string `json:"country"`
			CountryName string `json:"countryName"`
			Name        string `json:"name"`
			LogoUrls    []struct {
				Size string `json:"size"`
				URL  string `json:"url"`
			} `json:"logoUrls"`
			IsNational      bool `json:"isNational"`
			HasOfficialPage bool `json:"hasOfficialPage"`
			Competitions    []struct {
				CompetitionID   int    `json:"competitionId"`
				CompetitionName string `json:"competitionName"`
			} `json:"competitions"`
			Players []struct {
				ID           string `json:"id"`
				Country      string `json:"country"`
				FirstName    string `json:"firstName"`
				LastName     string `json:"lastName"`
				Name         string `json:"name"`
				Position     string `json:"position"`
				Number       int    `json:"number"`
				BirthDate    string `json:"birthDate"`
				Age          string `json:"age"`
				Height       int    `json:"height"`
				Weight       int    `json:"weight"`
				ThumbnailSrc string `json:"thumbnailSrc"`
				Affiliation  struct {
					Name         string `json:"name"`
					ThumbnailSrc string `json:"thumbnailSrc"`
				} `json:"affiliation"`
			} `json:"players"`
			Officials []struct {
				CountryName  string `json:"countryName"`
				ID           string `json:"id"`
				FirstName    string `json:"firstName"`
				LastName     string `json:"lastName"`
				Country      string `json:"country"`
				Position     string `json:"position"`
				ThumbnailSrc string `json:"thumbnailSrc"`
				Affiliation  struct {
					Name         string `json:"name"`
					ThumbnailSrc string `json:"thumbnailSrc"`
				} `json:"affiliation"`
			} `json:"officials"`
			Colors struct {
				ShirtColorHome string `json:"shirtColorHome"`
				ShirtColorAway string `json:"shirtColorAway"`
				CrestMainColor string `json:"crestMainColor"`
				MainColor      string `json:"mainColor"`
			} `json:"colors"`
		} `json:"team"`
	} `json:"data"`
	Message string `json:"message"`
}

var players []Fplayer

func main() {
	j := 10
	for i := 1; i < 23000 && j > 0; i++ {
		// var s string = fmt.Sprintf("%v", i)
		// fmt.Sprintf()

		//URL For Fetching Data
		url := "https://api-origin.onefootball.com/score-one-proxy/api/teams/en/" + fmt.Sprintf("%d", i) + ".json"
		method := "GET"

		payload := strings.NewReader(``)

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			continue
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer res.Body.Close()

		var rest Response

		//Convert io.Reader to slice of bytes
		body := StreamToByte(res.Body)
		//body, err := ioutil.ReadAll(res.Body)

		//Convert Slice of bytes to object of struct
		json.Unmarshal(body, &rest)

		//fmt.Println(rest.Data.Team.Name)

		//Let's check that this team is the Selected Team or Not
		if IsElementInSlice(rest.Data.Team.Name) {
			var temp Fplayer
			for _, player := range rest.Data.Team.Players {
				temp.ID = player.ID
				temp.Name = player.Name
				temp.Age = player.Age
				temp.Team = rest.Data.Team.Name
				players = append(players, temp)
			}
			j--

		}
		/*
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(body))
		*/
	}

	//Finding player who has two team and merge teams together
	for i, player := range players {
		for j := i + 1; j < len(players); j++ {
			if player.Name == players[j].Name {
				//fmt.Print(player.Name, player.Team, players[j].Team+"\n")
				//deleting reapited player
				players = append(players[:i], players[i+1:]...)
				// add second team
				players[j].Team = players[j].Team + ", " + players[i].Team
			}
		}
	}
	//Print players just Like output demand
	for i, player := range players {
		fmt.Print(". "+player.ID+"; "+player.Name+"; "+player.Age+"; "+player.Team+"\n", i)
	}

}
