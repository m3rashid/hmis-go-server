package middlewares

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	redisStorage "github.com/gofiber/storage/redis"
	"github.com/m3rashid-org/hmis-go-server/internal/utils"
)

var Cors = cors.New(cors.Config{
	AllowOrigins:     utils.Ternary(os.Getenv("SERVER_ENV") == "development", "*", "*"),
	AllowHeaders:     "Origin, Content-Type, Accept",
	AllowCredentials: true,
})

var Logger = logger.New(logger.Config{
	Format: "[${time} :: ${pid}] ${locals:requestid} ${status}-${method} ${path} ${latency}\n",
})

var Compression = compress.New(compress.Config{
	Level: compress.LevelBestSpeed,
})

var Etag = etag.New(etag.Config{})

var csrfStorage = redisStorage.New()
var Csrf = csrf.New(csrf.Config{
	CookieHTTPOnly: true,
	Storage:        csrfStorage,
})

var Idempotency = idempotency.New(idempotency.Config{})

var limiterStorage = redisStorage.New()
var Limiter = limiter.New(limiter.Config{
	Storage:    limiterStorage,
	Max:        50,
	Expiration: 1 * time.Minute,
})

var Recover = recover.New(recover.Config{
	EnableStackTrace:  os.Getenv("SERVER_ENV") == "development",
	StackTraceHandler: func(c *fiber.Ctx, e interface{}) {},
})
