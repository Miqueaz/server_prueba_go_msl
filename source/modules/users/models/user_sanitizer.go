package user_model

type UserSanitizer struct {
	PrimerNombre    string
	SegundoNombre   *string
	PrimerApellido  string
	SegundoApellido *string
	Matricula       string
	Correo          string
	Role            *string
}
