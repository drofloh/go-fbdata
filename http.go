package fbdata

const (
	// APIBaseURL ...
	APIBaseURL = "http://api.football-data.org/v2"
)

// Response ...
type Response struct {
	// The response body
	Body []byte

	// Some useful information is returned via response headers, eg:
	// www.football-data.org/documentation/api#response-headers

	// Not authenticated
	//  X-Requests-Available: 41
	//  X-RequestCounter-Reset: 52231
	//  X-API-Version: v2
	//  X-Authenticated-Client: anonymous

	// Authenticated
	// 	X-Requests-Available-Minute: 9
	//  X-RequestCounter-Reset: 60
	//  X-API-Version: v2
	//  X-Authenticated-Client: Andy

	// Shows the detected API-client or 'anonymous'
	AuthenticatedClient string
	// Defines the seconds left to reset your request counter.
	RequestCounterReset int
	// Shows the remaining requests before being blocked.
	RequestsAvailable int
}

func (c *CLI) request() (Response, err) {

}
