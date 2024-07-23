package dto

type SpotifyResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
	ExpriresIn  int64  `json:"expires_in,omitempty"`
	Success     bool   `json:"success"`
	Status      int64  `json:"status"`
}

/*

{
    "access_token": "BQBuhdCwsjgH6uEyfOuuHaVtsl0LqEShZQHLCIbOuFcVlDEjkCXMZQA79MHqEbDtq1yJz4666pEKQIspk9k0bhJgj1OQamq2tkIBi2vfcxS_J_r6POk",
    "token_type": "Bearer",
    "expires_in": 3600
}
*/
