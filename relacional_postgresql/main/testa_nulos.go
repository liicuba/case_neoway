package main

// Filtra usuários com CPF preenchido (descarta CPFs que são "" ou "null")
func FiltraUsuariosComCpfPreenchido(listaUsuarios []Usuario) []Usuario {
    var usuariosValidos []Usuario
    for _, usuario := range listaUsuarios {
        if usuario.CPF != "" || usuario.CPF != "null" { // Mantém CPF preenchido
            usuariosValidos = append(usuariosValidos, usuario)
