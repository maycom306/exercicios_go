package main

import (
	"sync"
	"strconv"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
    "github.com/gofiber/fiber/v3/middleware/recover"
)
// ============================================================
//  EXERCÍCIO 6 — API REST com Fiber
// ============================================================
//
// OBJETIVO:
//   Criar uma API REST usando o framework Fiber.
//
// SETUP (rode antes de começar):
//   go mod init ex06_fiber
//   go get github.com/gofiber/fiber/v3

//
// INSTRUÇÕES:
//
//  1. Crie a struct `Produto` com campos:
//       ID    int    `json:"id"`
//       Nome  string `json:"nome"`
//       Preco float64 `json:"preco"`
//
//  2. Crie um "banco de dados" em memória:
//       var produtos = map[int]Produto{}
//       var proximoID = 1
//       var mu sync.RWMutex   // para proteger acesso concorrente
//
//  3. Implemente os seguintes handlers HTTP:
//
//     GET  /tprodutos
//       → Retorna todos os produos em JSON
//       → Use mu.RLock() / mu.RUnlock() para leitura segura
//
//     GET  /produtos/:id
//       → Lê o parâmetro com: c.Params("id")
//       → Converta para int com strconv.Atoi
//       → Retorna 404 se não encontrado
//
//     POST /produtos
//       → Lê o body JSON usando c.BodyParser(&p)
//       → Atribui um ID automático
//       → Retorna 201 Created com o produto criado em JSON
//
//     DELETE /produtos/:id
//       → Remove o produto pelo ID
//       → Retorna 204 No Content se removido, 404 se não encontrado
//
//  4. No main(), configure o app Fiber e suba o servidor:
//       app := fiber.New()
//       app.Use(logger.New())    // log de requisições
//       app.Use(recover.New())   // recupera panics
//       // registre as rotas aqui
//       app.Listen(":8080")
//
// DICA — Retornar JSON:
//   Use c.Status(...).JSON(dados) para retornar dados em JSON.
//
// TESTE com curl ou Postman:
//   curl http://localhost:8080/produtos
//   curl -X POST http://localhost:8080/produtos \
//        -H "Content-Type: application/json" \
//        -d '{"nome":"Notebook","preco":3500.00}'
//   curl http://localhost:8080/produtos/1
//   curl -X DELETE http://localhost:8080/produtos/1
//
// ============================================================
type Produto struct{
	ID int `json:"id"`
	Nome string `json:"nome"`
	Preco float64 `json:"preco"`
}

type Db struct{
	produtos  map[int]Produto
	proximoID int
	mu sync.RWMutex
}


func main(){

	db := Db{
		produtos: make(map[int]Produto),
		proximoID: 1,
	}

	app := fiber.New(fiber.Config{
		AppName: "Appv1.0",
		ServerHeader: "Exercicio 6",
	})
	
	
	app.Post("/produtos", func(c fiber.Ctx) error {
	p := new(Produto)

	if err := c.Bind().Body(p); err != nil {
		return err
	}

	db.mu.Lock()
	p.ID = db.proximoID
	db.produtos[p.ID] = *p
	db.proximoID++
	db.mu.Unlock()

	return c.Status(fiber.StatusCreated).JSON(p)
})
	

	app.Get("/produtos/:id", func(c fiber.Ctx) error{
		idstring := c.Params("id")

		id, err := strconv.Atoi(idstring)

		if err != nil{
			return c.SendStatus(fiber.ErrBadRequest.Code)
		}

		produto, existe := db.produtos[id]
		if !existe{
			return c.SendStatus(fiber.ErrNotFound.Code)
		}
		
		return c.JSON(produto)
	})

	app.Get("/tprodutos", func(c fiber.Ctx) error{
		
		db.mu.RLock()
		defer db.mu.RUnlock()

		produtos := []Produto{}

		for _, Produto := range db.produtos{
			produtos = append(produtos, Produto)
			
		}
		return c.JSON(produtos)
	})

	app.Delete("/produtos/:id", func(c fiber.Ctx)error{
		db.mu.Lock()
		defer db.mu.Unlock()

		idstring := c.Params("id")

		id, err := strconv.Atoi(idstring)
		if err != nil{
			return c.SendStatus(fiber.StatusBadRequest)
		}
		
		_, existe := db.produtos[id]
		if !existe{
			return c.SendStatus(fiber.StatusNotFound)
		}
		delete(db.produtos, id)
		return c.SendStatus(fiber.StatusNoContent)
	})

	app.Listen(":8080")
	app.Use(logger.New())
	app.Use(recover.New())
}