package models

type Usuarios struct {
	Id_Usuario            int    `json:"id_usuario"`
	Email                 string `json:"email"`
	Nombres               string `json:"nombres"`
	Apellidos             string `json:"apellidos"`
	Dni                   int    `json:"dni"`
	Universidad           string `json:"universidad"`
	Celular               int    `json:"celular"`
	Imagen                string `json:"imagen"`
	Is_Active             bool   `json:"is_active"`
	Usuario_Administrador bool   `json:"usuario_administrador"`
	Password              string `json:"password"`
}
