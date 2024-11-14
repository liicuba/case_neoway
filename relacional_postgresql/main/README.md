## Este arquivo contém:

1. go.mod: gerenciamento de dependências em projetos Go;
2. go.sum: para garantir a integridade das dependências do projeto em Go;
3. main.go: contém a estrutura da API, a chamada das funções de validação e deixa rotas para possiveis get para o front, e chama a função salvaBanco para gravar somente os itens válidos no banco, separando-os em Usuario (para os CPFs) e Empresa (para os CNPJs);
4. testa_separador.go: contém a função para detectar os separadores do arquivo e organizá-lo;
5. testa_nulos.go: contém a função para testar se o arquivo contém nulos na primary key e descartá-los;
6. testa_duplicados.go: contém a função para testar se o arquivo duplicados na primary key e descartá-los;
7. tira_acentos.go: contém a função para tirar os acentos dos títulos das colunas;
8. padr_maiusc.go: contém a função para padronizar as letras maiúsculas;
9. valida.go: contém a função para:
  - Valida o campo CPF, define qual tipos eles são: CPF ou CNPJs, se caso não for nenhum dos dois, ele tira.
  - Verifica se CPFs e CNPJs são válidos de acordo com as leis brasileiras, descartando os não validos.
10. banco.go: gravas os itens válidos, de acordo com sua classe, no banco postgres.

## O fluxo não está funcional conforme explicado anteriormente.


