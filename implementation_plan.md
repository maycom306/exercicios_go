# Price Tracker — Go + Docker

Serviço que monitora preços de produtos em sites brasileiros e envia alertas quando o preço cai abaixo de um limite definido pelo usuário.

## Como vai funcionar

1. O usuário cadastra um produto via API REST (URL do produto + preço alvo)
2. Um worker em background verifica os preços periodicamente (ex: a cada 1h)
3. Quando o preço cai abaixo do alvo, envia notificação via **Telegram Bot**
4. Dados persistidos em **SQLite** (simples, sem dependências externas)

## Stack

- **Go** (Gin para API REST)
- **SQLite** via `modernc.org/sqlite` (pure Go, sem CGO)
- **Colly** para web scraping
- **Docker** com imagem final mínima (`scratch` ou `alpine`)

---

## Estrutura de Pastas

```
projetos/
└── price-tracker/
    ├── cmd/
    │   └── main.go           # entrypoint
    ├── internal/
    │   ├── api/
    │   │   └── handler.go    # rotas Gin (CRUD de produtos monitorados)
    │   ├── db/
    │   │   └── sqlite.go     # setup do banco e queries
    │   ├── scraper/
    │   │   └── scraper.go    # lógica de scraping com Colly
    │   ├── notifier/
    │   │   └── telegram.go   # envio de alerta no Telegram
    │   └── worker/
    │       └── worker.go     # goroutine que roda o check periódico
    ├── Dockerfile
    ├── docker-compose.yml
    ├── .env.example
    └── go.mod
```

## API Endpoints

| Método | Rota | Descrição |
|--------|------|-----------|
| `POST` | `/produtos` | Cadastra produto para monitorar |
| `GET`  | `/produtos` | Lista todos os produtos monitorados |
| `DELETE` | `/produtos/:id` | Remove produto |
| `GET`  | `/produtos/:id/historico` | Histórico de preços coletados |

## Variáveis de Ambiente (.env)

```
TELEGRAM_TOKEN=seu_token_aqui
TELEGRAM_CHAT_ID=seu_chat_id_aqui
CHECK_INTERVAL=60m
PORT=8080
```

## Open Questions

> [!IMPORTANT]
> O scraping funciona para sites específicos. Para a demo, vou usar o **Mercado Livre**
> (estrutura de HTML pública). Você quer incluir outros sites? (Amazon BR, Kabum, etc.)

> [!NOTE]
> O Telegram Bot é a forma mais simples de notificação. Você já tem um bot criado
> ou prefere que eu inclua as instruções de como criar um?

## Verification Plan

- `docker-compose up` deve subir o serviço sem erros
- `POST /produtos` cadastra e retorna o produto
- Worker roda o check e persiste o preço no histórico
- Alerta chega no Telegram quando preço cai
