package main

// TestaDuplicados verifica e filtra CPFs duplicados na lista de usuários válidos
func TestaDuplicados(usuariosValidos []Usuario) []Usuario {
	// Cria um mapa para verificar duplicidade
	cpfMap := make(map[string]bool)
	var usuariosUnicos []Usuario

	// Itera pela lista de usuários válidos
	for _, usuario := range usuariosValidos {
		// Se o CPF não estiver no mapa, adiciona ele
		if !cpfMap[usuario.CPF] && usuario.CPF != "" && usuario.CPF != "null" {
			usuariosUnicos = append(usuariosUnicos, usuario) // Adiciona usuário à lista
			cpfMap[usuario.CPF] = true // Marca o CPF como processado
		}
	}

	return usuariosUnicos
}
