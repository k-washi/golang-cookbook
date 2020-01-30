# OAuth

1. Client
2. OAuth API
3. Other APIs

1 -> 2 : Request Access Token (AT)
2 -> 1 : Return valid access token

1 -> 3 : request with AT
3 -> 2 : Validate AT
2 -> 3 : AT validated
3 -> 1 : Return response