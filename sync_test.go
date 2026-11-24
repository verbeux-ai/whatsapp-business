package whatsapp_business

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestSyncUsesGraphEnum(t *testing.T) {
	tests := []struct {
		name string
		kind SyncType
		want string
	}{
		{name: "history constant", kind: SyncTypeHistory, want: "history"},
		{name: "contact state constant", kind: SyncTypeAppStateSync, want: "smb_app_state_sync"},
		{name: "canonical media reupload", kind: SyncType("media_reupload"), want: "media_reupload"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			provider := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var request struct {
					Product  string `json:"messaging_product"`
					SyncType string `json:"sync_type"`
				}
				if r.Method != http.MethodPost || r.URL.Path != "/fixture-phone/smb_app_data" ||
					r.Header.Get("Authorization") != "Bearer fixture-token" ||
					json.NewDecoder(r.Body).Decode(&request) != nil || request.Product != "whatsapp" || request.SyncType != test.want {
					w.WriteHeader(http.StatusBadRequest)
					_, _ = w.Write([]byte(`{"error":{"message":"invalid sync request"}}`))
					return
				}
				_, _ = w.Write([]byte(`{"messaging_product":"whatsapp","request_id":"fixture-request"}`))
			}))
			defer provider.Close()

			client := NewClient(WithBaseUrl(provider.URL), WithHttpClient(provider.Client()), WithToken("fixture-token"), WithPhoneNumberId("fixture-phone"))
			response, err := client.RequestSync(context.Background(), test.kind)
			if err != nil {
				t.Fatal(err)
			}
			if response.RequestID != "fixture-request" {
				t.Fatalf("Graph request ID = %q", response.RequestID)
			}
		})
	}
}
