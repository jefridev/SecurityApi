package roles

import (
	"encoding/json"
	"net/http"
)

// CreateHandler add new role to database.
func CreateHandler(w http.ResponseWriter, r *http.Request) {

	var u RoleRegistration
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	role := &RoleModel{
		Name:        u.Name,
		Description: u.Description,
		Status:      u.Status,
	}

	err = SaveOne(role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m := map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "Role created",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(m)
}

// SetRoleToUserHandler add role to user to database.
func SetRoleToUserHandler(w http.ResponseWriter, r *http.Request) {

	var u UserRoleRegistration
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userRole := &UserRoleModel{
		UserID: u.UserID,
		RoleID: u.RoleID,
		Status: u.Status,
	}

	err = SaveOne(userRole)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m := map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "User set with role.",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(m)
}
