package models

type Sellers struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Document  string `json:"document"`
	Address   string `json:"address"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Clients struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Document  string `json:"document"`
	Address   string `json:"address"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Products struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Value     int64  `json:"value"`
	IdSeller  int64  `json:"idSeller"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Deliveries struct {
	ID          int64  `json:"id"`
	ProductName string `json:"productName"`
	Origin      string `json:"origin"`
	Destiny     string `json:"destiny"`
	IdSeller    int64  `json:"idSeller"`
	IdClient    int64  `json:"idClient"`
	IdProduct   int64  `json:"idProduct"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	ClientName  string `json:"clientName"`
}
