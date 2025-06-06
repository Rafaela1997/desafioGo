# desafioGo


API REST em Go usando [Gin](https://github.com/gin-gonic/gin), [GORM](https://gorm.io/) e banco de dados SQLite (com driver puro em Go).

## Funcionalidades

- Cria um banco de frutas automaticamente (`fruits.db`)
- Executa um crawler que consome a API pública de frutas:  
  [https://www.fruityvice.com/api/fruit/all](https://www.fruityvice.com/api/fruit/all)
- Salva os dados no banco SQLite
- Endpoint `/api/fruits/report-sugar` para agrupar frutas em:
  - **high_sugar** (açúcar ≥ 10)
  - **low_sugar** (açúcar < 10)


## Como rodar o projeto
- ```bash
go run .
### Pré-requisitos
- Go instalado → https://go.dev/dl/
- Não precisa de `gcc` ou `CGO` (usa driver puro em Go)

### Instalar dependências
```bash
go mod tidy

### A API ficará disponível em:
- http://localhost:8080

