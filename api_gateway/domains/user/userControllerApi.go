package user

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"tublessin/api_gateway/utils"
	"tublessin/common/model"

	"github.com/gorilla/mux"
)

type UserControllerApi struct {
	UserUsecaseApi UserUsecaseApiInterface
}

func NewLoginControllerApi(userService model.UserClient) *UserControllerApi {
	return &UserControllerApi{UserUsecaseApi: NewUserUsecaseApi(userService)}
}

func (c UserControllerApi) HandleGetUserProfileByID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		userId := mux.Vars(r)["id"]
		result, err := c.UserUsecaseApi.HandleGetUserProfileByID(userId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.UserResponeMessage{Response: "User Id Not Found", Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c UserControllerApi) HandleUpdateUserProfilePicture() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		getId := mux.Vars(r)["id"]

		fileName, err := utils.SaveFileToStorage(r, getId, "user")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&model.UserResponeMessage{Response: "Uploading Image Failed", Code: "500"})
			return
		}

		result, err := c.UserUsecaseApi.HandleUpdateUserProfilePicture(getId, fileName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: "Uploading Image Failed", Code: "500"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

// Ini function untuk serve yang ada di harddisk
func (s UserControllerApi) HandleServeUserFile() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dir, err := os.Getwd()
		if err != nil {
			log.Println(err)
			return
		}

		fileId := mux.Vars(r)["namaFile"]
		fileLocation := filepath.Join(dir, "fileServer", "user", fileId)

		w.Header().Set("Content-Type", "image/jpeg")
		http.ServeFile(w, r, fileLocation)
	}
}

func (c UserControllerApi) HandleUpdateUserProfileByID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var userProfile model.UserProfile
		json.NewDecoder(r.Body).Decode(&userProfile)
		convertId, _ := strconv.Atoi(mux.Vars(r)["id"])
		userProfile.Id = int32(convertId)

		result, err := c.UserUsecaseApi.HandleUpdateUserProfileByID(&userProfile)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.UserResponeMessage{Response: "Updating User Profile Failed", Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c UserControllerApi) HandleUpdateUserLocation() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var userProfile model.UserProfile
		var userLocation *model.UserLocation
		json.NewDecoder(r.Body).Decode(&userLocation)

		convertId, _ := strconv.Atoi(mux.Vars(r)["id"])
		userProfile.Id = int32(convertId)
		userProfile.Location = userLocation

		result, err := c.UserUsecaseApi.HandleUpdateUserLocation(&userProfile)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.UserResponeMessage{Response: "Updating User Location Failed", Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}
