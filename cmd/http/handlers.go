package http

import (
	"encoding/json"
	"net/http"
	"uala-followers-service/config"
	"uala-followers-service/internal/application"
)

func createFollow(deps *config.Dependencies) http.HandlerFunc {
	createFollow := application.NewCreateFollow(deps.FollowRepository, deps.EventPublisher)
	return func(w http.ResponseWriter, r *http.Request) {
		var cmd application.CreateFollowCommand
		err := json.NewDecoder(r.Body).Decode(&cmd)
		if err != nil {
			handleError(w, err)
			return
		}
		response, err := createFollow.Exec(r.Context(), &cmd)
		if err != nil {
			handleError(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			handleError(w, err)
			return
		}
	}
}

func getFollowers(deps *config.Dependencies) http.HandlerFunc {
	getFollowers := application.NewGetFollowers(deps.FollowRepository)
	return func(w http.ResponseWriter, r *http.Request) {
		var cmd application.GetFollowersCommand
		err := json.NewDecoder(r.Body).Decode(&cmd)
		if err != nil {
			handleError(w, err)
			return
		}
		response, err := getFollowers.Exec(r.Context(), &cmd)
		if err != nil {
			handleError(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			handleError(w, err)
			return
		}
	}
}

func getFollowings(deps *config.Dependencies) http.HandlerFunc {
	createFollow := application.NewGetFollowings(deps.FollowRepository)
	return func(w http.ResponseWriter, r *http.Request) {
		var cmd application.GetFollowingsCommand
		err := json.NewDecoder(r.Body).Decode(&cmd)
		if err != nil {
			handleError(w, err)
			return
		}
		response, err := createFollow.Exec(r.Context(), &cmd)
		if err != nil {
			handleError(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			handleError(w, err)
			return
		}
	}
}

func handleError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	var errorResp ErrorResponse

	switch {
	default:
		errorResp = ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Code:       "INTERNAL_ERROR",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorResp.StatusCode)

	jsonResp, jsonErr := json.Marshal(errorResp)
	if jsonErr != nil {
		http.Error(w, "Error processing response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
	return
}

type ErrorResponse struct {
	StatusCode int    `json:"status,omitempty"`
	Message    string `json:"message,omitempty"`
	Code       string `json:"code,omitempty"`
}
