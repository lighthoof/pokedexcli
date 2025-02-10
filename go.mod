module github.com/lighthoof/pokedexcli

go 1.23.5

replace (
	github.com/lighthoof/pokedexcli/internal/pokeAPIHandler => /home/niriam/projects/github.com/lighthoof/pokedexcli/internal/pokeAPIHandler
	github.com/lighthoof/pokedexcli/internal/pokeCache => /home/niriam/projects/github.com/lighthoof/pokedexcli/internal/pokeCache
)

require (
	github.com/lighthoof/pokedexcli/internal/pokeAPIHandler v0.0.0-20250202132351-d65dde3b2174
	github.com/lighthoof/pokedexcli/internal/pokeCache v0.0.0-20250208201419-101ea4996a36
)
