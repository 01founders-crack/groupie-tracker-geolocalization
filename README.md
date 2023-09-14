# groupie-tracker

groupie-tracker

You should create a .env file
For ex:

```bash
ACCESS_TOKEN=YOUR_MAPBOX_TOKEN
GMAPS_TOKEN=YOUR_GMPAS_TOKEN
```

```bash
.
├── LICENSE
├── README.md
├── backend
│   ├── external
│   │   └── groupietracker_api.go
│   ├── handlers
│   │   ├── common.go
│   │   ├── error_handler.go
│   │   ├── group_handler.go
│   │   └── root_handler.go
│   ├── helpers
│   │   └── helpers.go
│   ├── main.go
│   ├── mapboxgeo
│   │   └── mapboxgeo.go
│   ├── models
│   │   ├── artist.go
│   │   ├── artist_data.go
│   │   ├── combined_data.go
│   │   └── relations.go
│   └── tests
├── certificates
│   ├── fullchain.crt
│   ├── server.crt
│   └── server.key
├── frontend
│   ├── 404.html
│   ├── 500.html
│   ├── footer.html
│   ├── google_maps.html
│   ├── group.html
│   ├── group_info.html
│   ├── header.html
│   ├── images
│   │   ├── 01FCLogo.webp
│   │   ├── 404error.png
│   │   ├── MusicBG.svg
│   │   ├── donate.svg
│   │   ├── github.svg
│   │   └── twitter.svg
│   ├── index.html
│   ├── layout.html
│   ├── music_card.html
│   ├── scripts
│   │   └── main.js
│   └── styles
│       └── style.css
├── go.mod
└── run_page.sh

```
