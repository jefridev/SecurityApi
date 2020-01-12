package users

import "net/http"

import "encoding/json"

// CreateHandler add new user to database.
func CreateHandler(w http.ResponseWriter, r *http.Request) {

	var u UserRegistration
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &UserModel{
		Username: u.Username,
		Email:    u.Username,
		Bio:      u.Bio,
		Image:    &u.Image,
	}

	err = user.setPassword(u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = user.checkPassword(u.PasswordConfirmation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = SaveOne(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m := map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "User created",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(m)
}
