package instagram

import (
	"fmt"
	"net/http"
	"net/url"
)

type RealtimeResponse struct {
	Data Realtime
	Meta RealtimeMeta
}

type RealtimeService struct {
	client *Client
}

type RealtimeMeta struct {
	Code int `json:"code,omitempty"`
}

// Realtime represents a realtime subscription on Instagram's service.
type Realtime struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	Object      string `json:"object,omitempty"`
	ObjectID    string `json:"object_id,omitempty"`
	Aspect      string `json:"aspect,omitempty"`
	CallbackURL string `json:"callback_url,omitempty"`
}

// MediaComments gets a full list of comments on a media.
//
// Instagram API docs: http://instagram.com/developer/endpoints/comments/#get_media_comments
func (s *RealtimeService) SubscribeToTag(tag, callbackURL, verifyToken string) (*RealtimeResponse, error) {
	u := "subscriptions/"

	params := url.Values{
		"aspect":        {"media"},
		"object":        {"tag"},
		"object_id":     {tag},
		"callback_url":  {callbackURL},
		"client_id":     {s.client.ClientID},
		"client_secret": {s.client.ClientSecret},
		"verify_token":  {verifyToken},
	}

	req, err := s.client.NewRequest("POST", u, params.Encode())
	if err != nil {
		return nil, err
	}

	realtimeResponse := new(RealtimeResponse)

	_, err = s.client.Do(req, realtimeResponse)
	if err != nil {
		return nil, err
	}

	return realtimeResponse, err
}

//An example RealTimeSubscribe ResponseWriter. This can be plugged directly into
// any standard http server. Note, however, that this particular implementation does
// no checking that the verifyToken is correct.
func ServeInstagramRealtimeSubscribe(w http.ResponseWriter, r *http.Request) {
	//You would think this would be "PostFormValue", but Instagram is sending the
	// challenge and response data as query strings in a POST, not as POST elements in a POST
	verify := r.FormValue("hub.challenge")

	fmt.Fprintf(w, verify)
}
