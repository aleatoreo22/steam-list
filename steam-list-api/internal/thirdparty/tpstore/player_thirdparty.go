package tpstore

type StoreThirdParty interface {
	StoreID() string
	GetPlayerGames(playerID string, page int, itemsPerPage int) (*[]string, error)
}
