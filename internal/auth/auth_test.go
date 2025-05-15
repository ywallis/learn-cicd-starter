package auth_test

import (
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		headers http.Header
		want    string
		wantErr bool
	}{{

		name: "pass",
		// headers: http.Header{"Authorization": []string{"ApiKey abc"]}},
		headers: func() http.Header {
			h := http.Header{}
			h.Set("Authorization", "ApiKey abc")
			return h
		}(),
		want:    "abc",
		wantErr: false,
	}, {

		name: "wrong type",
		// headers: http.Header{"Authorization": []string{"ApiKey abc"]}},
		headers: func() http.Header {
			h := http.Header{}
			h.Set("Authorization", "Bearer abc")
			return h
		}(),
		want:    "",
		wantErr: true,
	},{

		name: "empty value",
		// headers: http.Header{"Authorization": []string{"ApiKey abc"]}},
		headers: func() http.Header {
			h := http.Header{}
			h.Set("Authorization", "")
			return h
		}(),
		want:    "",
		wantErr: true,
	},{

		name: "no field",
		// headers: http.Header{"Authorization": []string{"ApiKey abc"]}},
		headers: func() http.Header {
			h := http.Header{}
			return h
		}(),
		want:    "",
		wantErr: true,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := auth.GetAPIKey(tt.headers)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetAPIKey() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetAPIKey() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
