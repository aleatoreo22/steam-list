package steamworksModel

type Root[T any] struct {
	Response T `json:"response"`
}
