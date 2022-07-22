package handlers

import (
	"GatewayService/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Login(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	req, _ := http.NewRequest(http.MethodPost,
		utils.UsersServiceRoot.Next().Host+UsersServiceApi+"/login", r.Body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func Register(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	req, _ := http.NewRequest(http.MethodPost,
		utils.UsersServiceRoot.Next().Host+UsersServiceApi+"/register", r.Body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func FindAllUsers(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")

	response, err := http.Get(
		utils.UsersServiceRoot.Next().Host + UsersServiceApi + "/getUsers?page=" + page + "&size=" + size)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func UpdateUser(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	req, _ := http.NewRequest(http.MethodPut,
		utils.UsersServiceRoot.Next().Host+UsersServiceApi+"/updateUser", r.Body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func DeleteUser(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	params := mux.Vars(r)
	userId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodDelete,
		utils.UsersServiceRoot.Next().Host+UsersServiceApi+"/deleteUser/"+strconv.FormatUint(uint64(userId), 10),
		r.Body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func BanUser(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	params := mux.Vars(r)
	userId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodPatch,
		utils.UsersServiceRoot.Next().Host+UsersServiceApi+"/banUser/"+strconv.FormatUint(uint64(userId), 10),
		r.Body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func UnbanUser(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	params := mux.Vars(r)
	userId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodPatch,
		utils.UsersServiceRoot.Next().Host+UsersServiceApi+"/unbanUser/"+strconv.FormatUint(uint64(userId), 10),
		r.Body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}
