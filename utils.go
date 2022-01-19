package main

import (
	"fmt"
	"strconv"

	"github.com/sandertv/gophertunnel/query"
)

type (
	CardData struct {
		Percent      int
		PercentStyle string
		QueryResult
	}
	QueryResult struct {
		HostIP    string
		Player    string
		Hostname  string
		Version   string
		NumPlayer int
		MaxPlayer int
		HostPort  int
		Gametype  string
		GameID    string
		Plugins   string
		Map       string
	}
)

func getServerCardData(ip string) CardData {
	if ip == "" {
		return CardData{
			Percent: (12 * 100) / (20 + 1),
			QueryResult: QueryResult{
				HostIP:    ip,
				Player:    "Compico, TestPlayer, Herobrine",
				Hostname:  "test.examplename.ru",
				Version:   "1.18.1",
				NumPlayer: 12,
				MaxPlayer: 20,
				HostPort:  25565,
				Gametype:  "survival",
				GameID:    "minecraft",
				Plugins:   "none",
				Map:       "world",
			},
		}
	}
	res, err := query.Do(ip)
	fmt.Println(err.Error())
	var x = QueryResult{}
	x.NumPlayer, _ = strconv.Atoi(res["numplayers"])
	x.MaxPlayer, _ = strconv.Atoi(res["maxplayers"])
	x.HostPort, _ = strconv.Atoi(res["hostport"])
	x.Player = res["players"]
	x.HostIP = res["hostip"]
	x.Hostname = res["hostname"]
	x.Version = res["version"]
	x.Gametype = res["gametype"]
	x.GameID = res["game_id"]
	x.Plugins = res["plugins"]
	x.Map = res["map"]
	var dataservercard = CardData{
		Percent:     (x.NumPlayer * 100) / (x.MaxPlayer + 1),
		QueryResult: x,
	}
	if dataservercard.Percent < 8 {
		dataservercard.Percent = 8
	}
	dataservercard.PercentStyle = fmt.Sprintf("style=\"width: %d%%;\"", dataservercard.Percent)
	return dataservercard
}
