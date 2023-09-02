package domain

type Store struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewStore(storeID, name string) Store {
	return Store{
		ID:   storeID,
		Name: name,
	}
}
