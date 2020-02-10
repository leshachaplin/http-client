package handlers

type UserModel struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TokenModel struct {
	Token string `json:"token"`
}
type ClaimModel struct {
	Key   string `json:"keyForClaim"`
	Value string `json:"valueForClaim"`
}
