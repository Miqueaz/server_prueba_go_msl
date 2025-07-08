package user_sanitizer

type UserSanitizer struct {
	PrimerNombre    string  `db:"primernombre"`
	SegundoNombre   *string `db:"segundonombre"`
	PrimerApellido  string  `db:"primerapellido"`
	SegundoApellido *string `db:"segundoapellido"`
	Matricula       string  `db:"matricula"`
	Correo          string  `db:"correo"`
}