package student

type StudentDTO struct {
	ID 			uint   `json:"id,string,omitempty"`
	Identifier 	string `json:"identifier"`
	Name 		string `json:"name"`
	Email 		string `json:"email"`
	UKT 		uint   `json:"ukt,string"`
}

type CreateStudentDTO struct {
	ID 			uint   `json:"id,string,omitempty"`
	Name 		string `json:"name"`
	Email 		string `json:"email"`
	UKT 		uint   `json:"ukt,string"`
}

type UpdateStudentDTO struct {
	ID 			uint   `json:"id,string,omitempty"`
	Name 		string `json:"name"`
	Email 		string `json:"email"`
}