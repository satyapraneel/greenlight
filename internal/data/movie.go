package data

import (
	"time"
)

// {
// 	"id": 123,
// 	"title": "Casablanca",
// 	"runtime": 102,
// 	"genres": [
// 	"drama",
// 	"romance",
// 	"war"
// 	],
// 	"version": 1
// 	}

type Movie struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Title     string    `json:"title,omitempty"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty,string"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version,omitempty"`
}

// type MovieAlias Movie

// incase of json Alternative way of chaning Runtime (or any parameter detail)
// func (m Movie) MarshalJSON() ([]byte, error) {
// 	aux := struct {
// 		MovieAlias
// 		Runtime string `json:"runtime"`
// 	}{
// 		MovieAlias: MovieAlias(m),
// 		Runtime:    fmt.Sprintf("%d mins", m.Runtime),
// 	}

// 	return json.Marshal(aux)
// }
