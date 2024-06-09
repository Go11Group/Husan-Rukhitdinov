// PROBLEM
func (us *NewHttp) CreateProblem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "It's not POST request!", http.StatusMethodNotAllowed)
		return
	}
	var problem model.Problem
	err := json.NewDecoder(r.Body).Decode(&problem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = us.User.CreateProblem(problem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(problem)
}

func (us *NewHttp) UpdateProblem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "It's not PUT request!", http.StatusMethodNotAllowed)
		return
	}
	var problem model.Problem
	err := json.NewDecoder(r.Body).Decode(&problem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = us.User.UpdateProblem(problem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(problem)
}

func (us *NewHttp) DeleteProblem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "It's not DELETE request!", http.StatusMethodNotAllowed)
		return
	}
	var problem model.Problem
	err := json.NewDecoder(r.Body).Decode(&problem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = us.User.DeleteProblem(problem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(problem)
}