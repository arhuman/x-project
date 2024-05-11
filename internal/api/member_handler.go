package api

import (
	"biomassx/internal/member"
	"biomassx/models"
	"encoding/json"
	"net/http"
	"strconv"

	"git.doolta.com/doolta/go-kit/web"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"go.uber.org/zap"
)

// @task: write the functions implementation for the following handler, returning a api.Response encoded as json @run

// CreateMemberHandlerFunc godoc
// @Summary Create a new member
// @Tags member
// @Success 201 {object} models.Member{}
// @Router /v1/members/ [post]
func CreateMemberHandlerFunc(w http.ResponseWriter, r *http.Request) {
	log, services := web.GetServerContext(r)
	log.Debug("CreateMemberHandlerFunc called")
	member := services["member"].(interface{}).(*member.ServiceImpl)

	// Get member data from request body
	var newMember *models.Member

	err := json.NewDecoder(r.Body).Decode(&newMember)
	if err != nil {
		log.Error("Failed to decode request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the password before saving it
	bytes, err := bcrypt.GenerateFromPassword([]byte(newMember.Passwd), 13)
	if err != nil {
		log.Error("can't generate hash for password: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newMember.Passwd = string(bytes)

	// Create a new member
	err = member.Create(newMember)
	if err != nil {
		log.Error("Failed to create member", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    newMember,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// ReadMemberHandlerFunc godoc
// @Summary Get a member by ID
// @Tags member
// @Success 200 {object} models.Member{}
// @Router /v1/members/{id} [get]
func ReadMemberHandlerFunc(w http.ResponseWriter, r *http.Request) {
	log, services := web.GetServerContext(r)
	log.Debug("CreateMemberHandlerFunc called")
	member := services["member"].(interface{}).(*member.ServiceImpl)

	// Get member ID from path
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
	if err != nil {
		log.Error("Failed to parse member ID from URL", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Get member by ID
	m, err := member.Read(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    m,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateMemberHandlerFunc godoc
// @Summary Update a member by ID
// @Tags member
// @Success 200 {object} models.Member{}
// @Router /v1/members/{id} [put]
func UpdateMemberHandlerFunc(w http.ResponseWriter, r *http.Request) {
	log, services := web.GetServerContext(r)
	log.Debug("CreateMemberHandlerFunc called")
	member := services["member"].(interface{}).(*member.ServiceImpl)

	// Get member ID from path
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
	if err != nil {
		log.Error("Failed to parse member ID from URL", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Get member by ID
	m, err := member.Read(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get updated member data from request body
	err = json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update member
	err = member.Update(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    m,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteMemberHandlerFunc godoc
// @Summary Delete a member by ID
// @Tags member
// @Success 204 {object} models.Member{}
// @Router /v1/members/{id} [delete]
func DeleteMemberHandlerFunc(w http.ResponseWriter, r *http.Request) {
	log, services := web.GetServerContext(r)
	log.Debug("CreateMemberHandlerFunc called")
	service := services["member"].(interface{}).(*member.ServiceImpl)

	id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
	if err != nil {
		log.Error("Failed to parse member ID from URL", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Check member exists
	_, err = service.Read(uint(id))
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			http.Error(w, "Member does not exists", http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Delete member
	err = service.Delete(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Status:  http.StatusNoContent,
		Message: "Success",
		Data:    nil,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
