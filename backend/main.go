package main

import (
	"context"
	"database/sql"
	"embed"
	"io/fs"
	"log"
	"net"
	"net/http"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	gameservice "github.com/PlatosCodes/desserted/backend/game_service"
	"github.com/PlatosCodes/desserted/backend/gapi"
	"github.com/PlatosCodes/desserted/backend/mailer"

	"github.com/PlatosCodes/desserted/backend/pb"
	"github.com/PlatosCodes/desserted/backend/util"
	"github.com/PlatosCodes/desserted/backend/ws"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	_ "embed"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	dbSource := "postgresql://root:bluecomet@localhost:5432/desserted?sslmode=disable"
	log.Println(dbSource)
	conn, err := sql.Open(config.DBDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	runDBMigration(config.MigrationURL, dbSource)

	store := db.NewStore(conn)
	gameService := gameservice.NewGameService(store)

	mailer := mailer.New(config.SmtpHost, config.SmtpPort, config.SmtpUsername,
		config.SmtpPassword, config.SmtpSender)
	if err != nil {
		log.Fatal("cannot create mailer:", err)
	}

	go runGatewayServer(config, store, gameService, mailer) //run in a separate routine
	runGrpcServer(config, store, mailer)
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}
	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}
	log.Println("db migrated successfully")
}

func runGrpcServer(config util.Config, store db.Store, mailer *mailer.Mailer) {
	server, _, err := gapi.NewServer(config, store, mailer)
	if err != nil {
		log.Fatal("cannot create server1", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDessertedServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}

// Embed Swagger docs into content
//
//go:embed doc/swagger/*
var swagger embed.FS

func runGatewayServer(config util.Config, store db.Store, gameService *gameservice.GameService, mailer *mailer.Mailer) {
	server, tokenMaker, err := gapi.NewServer(config, store, mailer)
	if err != nil {
		log.Fatal("cannot create server2", err)
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterDessertedHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatalf("cannot register handler server: %s", err)
	}

	limiterMiddleware := gapi.RateLimitMiddleware(5, 10)

	mux := http.NewServeMux()
	mux.Handle("/", limiterMiddleware(grpcMux))

	wsHub := ws.NewHub()
	go wsHub.Run()              // Run WebSocket Hub in its own goroutine
	go wsHub.HandleGameEvents() // Start handling game events

	messageQueue := ws.NewMessageQueue(100) // Adjust size as needed
	messageQueue.StartProcessing(wsHub)

	mux.Handle("/ws", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(wsHub, w, r, config, store, gameService, messageQueue, tokenMaker)
	}))

	// Create a sub filesystem that contains the swagger files
	swaggerFiles, err := fs.Sub(swagger, "doc/swagger")
	if err != nil {
		log.Fatalf("cannot create sub filesystem: %s", err)
	}

	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.FS(swaggerFiles))))

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{config.FrontendAddress}, // Adjust as needed
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "X-Requested-With"},
		AllowCredentials: true,
		Debug:            true, // Set to false in production
	})

	handler := corsHandler.Handler(mux)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("start HTTP gateway server at %s", listener.Addr().String())
	if err := http.Serve(listener, handler); err != nil {
		log.Fatal("cannot start HTTP gateway server:", err)
	}
}
