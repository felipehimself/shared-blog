## SharedBlog

Trata-se do backend de uma aplicação desenvolvido em Go. A aplicação é similar a um blog compartilhado, no qual usuários podem criar publicações que ficarão visíveis para outros usuários. 

Usuários que criarem uma conta e realizarem login podem comentar em postagens e também votar. Esta validação é feita por meio de middleware que verifica se o token enviado nos cookies da requisição é válido. 

### Endpoints

O projeto possui 15 endpoints:

##### Públicos (não requer autenticação)

- `/api/user/signup`: Cria um usuário;
- `/api/user/signin`: Faz login do usuário;
- `/api/user/signout`: Faz logout do usuário;
- `/api/user/is-authorized`: Verifica se o token do usuário é e está válido;
- `/api/post/get-posts`: Retorna todas as postagens;
- `/api/post/get-post/:postId`: Retorna uma única postagem;
- `/api/post/get-user-posts/:username`: Retorna uma única postagem de um usuário.

##### Protegidos (requer autenticação)

- ```/api/protected/post/create-post```: Cria uma postagem; 
- ```/api/protected/post/vote/:postId```: Vota em uma postagem;
- ```/api/protected/post/unvote/:postId```: Remove o voto de uma postagem;
- ```/api/protected/post/edit-post/:postId```: Edita uma postagem;
- ```/api/protected/post/delete-post/:postId```: Deleta uma postagem;
- ```/api/protected/comment/comment-post```: Comenta em uma postagem;
- ```/api/protected/comment/delete-comment/:commentId```: Deleta o comentário de uma postagem;
- ```/api/protected/topics/get-topics```: Retorna os tópicos cadastrados no banco de dados.

## Sumário

- [Tecnologias utilizadas](#tecnologias)
- [Instruções para rodar o projeto](#instrucoes)
- [Organização e estruturação do projeto](#organizacao)
- [Desenvolvimento](#desenvolvimento)

## Tecnologias Utilizadas <a name="tecnologias"></a>

#### Backend

- [**Fiber**](https://docs.gofiber.io/)
- [**MySQL**](https://www.mysql.com/)
- [**Crypto**](https://golang.org/x/crypto)
- [**JWT Go**](https://pkg.go.dev/github.com/dgrijalva/jwt-go@v3.2.0+incompatible)

## Instruções para rodar o projeto <a name="instrucoes"></a>

O modelo das tabelas utilizadas estão no diretório: ```sql/sql.sql ```

Você poderá criá-las localmente utilizando o  [MySQL Workbench](https://docs.gofiber.io/) ou um banco de dados remoto como o [DB4Free](db4free.net).

#### Será necessário ter instalado em sua máquina:

```
Git
Go
```

- Clone o repositório com o comando **git clone**:

```
git clone https://github.com/felipehimself/shared-blog
```

- Entre no diretório que acabou de ser criado:

```
cd sharedblog/backend
```

- Faça a instalação das dependências do projeto:

```
go get
```

- Inicialize o projeto:

```
go run main.go
```


## Organização e estruturação do projeto <a name="organizacao"></a>

O projeto está estruturado da seguinte forma:

```
   └───backend
    │   go.mod
    │   go.sum
    │   main.go
    │
    ├───sql
    │       sql.sql
    │
    └───src
        ├───config
        │       config.go
        │
        ├───controllers
        │       auth.go
        │       comment.go
        │       post.go
        │       topic.go
        │
        ├───database
        │       database.go
        │
        ├───middleware
        │       middleware.go
        │
        ├───models
        │       comment.go
        │       post.go
        │       topic.go
        │       user.go
        │
        ├───repositories
        │       comment.go
        │       post.go
        │       topic.go
        │       user.go
        │
        ├───responses
        │       responses.go
        │
        ├───routes
        │       routes.go
        │
        └───utils
                token.go
```

## Desenvolvimento <a name="desenvolvimento"></a>

#### Backend

#### [**Fiber**](https://docs.gofiber.io/)

Web framework inspirado pelo Express do NodeJS.

#### [**MySQL**](https://www.mysql.com/)

Banco de dados relacional.

#### [**Crypto**](https://golang.org/x/crypto)

Pacote utilizado para encriptar as senhas antes de salvá-las no banco de dados e para verificar se a mesma é válida quando o usuário realizar autenticação.

#### [**JWT Go**](https://pkg.go.dev/github.com/dgrijalva/jwt-go@v3.2.0+incompatible)

Pacote utilizado para gerar o token jwt quando o usuário efetua login e para verificar se o mesmo é válido nas requisições das rotas protegidas.