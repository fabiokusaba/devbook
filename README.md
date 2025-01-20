# DevBook

## Introdução
DevBook é uma rede social para criar publicações textuais onde teremos duas principais entidades:
* Usuários: além das operações básicas de CRUD será possível também o usuário seguir outro usuário, parar de seguir um usuário, buscar todos os usuários que segue, buscar todos os usuários que não segue e atualizar sua senha. Para lidarmos com tudo isso teremos duas tabelas no banco de dados sendo: usuários e seguidores.
* Publicações: além das operações básicas de CRUD será possível buscar publicações de acordo com os usuários que segue e também será possível curtir publicações. Para lidarmos com isso teremos uma tabela no banco de dados chamada: publicações.

## Estrutura da aplicação
Teremos uma aplicação web que vai chamar a nossa API que vai ser o meio de comunicação com o nosso banco de dados e para trabalharmos no seu desenvolvimento usaremos uma estruturação em pacotes que podem ser divididos em dois tipos: pacotes principais e pacotes auxiliares.

### Pacotes principais
Pacotes relacionados a estrutura da nossa API, vamos ter cinco pacotes principais:
* Main: arquivo executável
* Router: configurar o router e todas as rotas que estão embaixo dele
* Controllers: funções que vão lidar com as requisições http
* Modelos: guardar as entidades da aplicação
* Repositórios: interação com o banco de dados

### Pacotes auxiliares
Pacotes que vão lidar com as utilidades da API como um todo
* Config: configurações de variáveis de ambiente
* Banco: abrir a conexão com o banco de dados
* Autenticação: cuidar do login, criação de token, etc
* Middleware: camada que vai ficar entre a requisição e a resposta para verificar se o usuário está autenticado
* Segurança: criptografia de senhas, verificação de senhas
* Respostas: padronização das respostas da API