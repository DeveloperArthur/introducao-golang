/*
Em termos de desempenho, o segundo trecho de código é geralmente
mais eficiente em relação ao primeiro. Isso ocorre porque,
ao passar um ponteiro para o objeto aluno na função MudaNome,
evita-se a cópia do objeto inteiro, o que pode ser
especialmente vantajoso quando o objeto é grande.
*/
func Main() {
  var aluno models.Aluno
	aluno.Nome = "Arthur"
	aluno = MudaNome(aluno)
}

func MudaNome(aluno models.Aluno) models.Aluno {
	aluno.Nome = "nome alterado"
	return aluno
}

///////////////
func Main2() {
	var aluno models.Aluno
	aluno.Nome = "Arthur"
	MudaNome2(&aluno)
}

func MudaNome2(aluno *models.Aluno) {
	aluno.Nome = "nome alterado"
}
