package entities

type User struct {
	ID             int
	Login          string
	HashedPassword string
	FIO            string
	APIToken       string
	Actions        []Actions
}
