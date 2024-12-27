# mimir-scrapper

https://www.youtube.com/watch?v=wUSgA8WEy4Q
https://github.com/seversky/gachifinder/blob/master/scrape/scrape.go

mimir-scrapper/
├── cmd/                     # Hlavní vstupní body aplikace
│   └── scrapper/
│       └── main.go          # Entry point aplikace
├── internal/                # Kód, který není určený k opakovanému použití externě
│   ├── fetcher/             # Balíček pro získávání dat (HTTP požadavky, apod.)
│   │   ├── fetcher.go
│   ├── parser/              # Balíček pro zpracování/parsing HTML, JSON, apod.
│   │   ├── parser.go
│   ├── storage/             # Balíček pro ukládání dat (DB, soubory)
│   │   ├── storage.go
│   └── config/              # Balíček pro konfiguraci aplikace
│       ├── config.go
├── pkg/                     # Knihovny a kód, který lze opakovaně použít i v jiných projektech
│   ├── logger/              # Balíček pro logování
│   │   ├── logger.go
│   └── utils/               # Obecné utility
│       ├── utils.go
├── configs/                 # Šablony nebo výchozí konfigurace
│   └── config.yaml
├── test/                    # Externí testovací data a aplikace
│   ├── data/
│   └── mocks/
├── scripts/                 # Skripty pro build, deploy, analýzu, apod.
│   └── run_scraper.sh
├── go.mod                   # Go module soubor
└── README.md                # Dokumentace projektu

