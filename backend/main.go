package main

// import (
// 	"context"
// 	"database/sql"
// 	"embed"
// 	"io/fs"
// 	"log"
// 	"net"
// 	"net/http"

// 	"github.com/golang-migrate/migrate"
// 	"github.com/grpc-ecosystem/grpc-gateway/runtime"
// 	_ "github.com/lib/pq"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/reflection"
// 	"google.golang.org/protobuf/encoding/protojson"
// )

// func main() {
// 	config, err := util.LoadConfig(".")
// 	if err != nil {
// 		log.Fatal("cannot load config", err)
// 	}

// 	conn, err := sql.Open(config.DBDriver, config.DBSource)
// 	if err != nil {
// 		log.Fatal("cannot connect to db:", err)
// 	}

// 	runDBMigration(config.MigrationURL, config.DBSource)

// 	store := db.NewStore(conn)
// 	go runGatewayServer(config, store) //run in a separate routine
// 	runGrpcServer(config, store)
// }

// func runDBMigration(migrationURL string, dbSource string) {
// 	migration, err := migrate.New(migrationURL, dbSource)
// 	if err != nil {
// 		log.Fatal("cannot create new migrate instance:", err)
// 	}
// 	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
// 		log.Fatal("failed to run migrate up:", err)
// 	}
// 	log.Println("db migrated successfully")
// }

// func runGrpcServer(config util.Config, store db.Store) {
// 	server, err := gapi.NewServer(config, store)
// 	if err != nil {
// 		log.Fatal("cannot create server")
// 	}

// 	grpcServer := grpc.NewServer()
// 	pb.RegisterMomsRecipesServer(grpcServer, server)
// 	reflection.Register(grpcServer)

// 	listener, err := net.Listen("tcp", config.GRPCServerAddress)
// 	if err != nil {
// 		log.Fatal("cannot create listener")
// 	}

// 	log.Printf("start gRPC server at %s", listener.Addr().String())
// 	err = grpcServer.Serve(listener)
// 	if err != nil {
// 		log.Fatal("cannot start gRPC server")
// 	}
// }

// // Embed Swagger docs into content
// //
// //go:embed doc/swagger/*
// var swagger embed.FS

// func runGatewayServer(config util.Config, store db.Store) {
// 	server, err := gapi.NewServer(config, store)
// 	if err != nil {
// 		log.Fatal("cannot create server")
// 	}

// 	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
// 		MarshalOptions: protojson.MarshalOptions{
// 			UseProtoNames: true,
// 		},
// 		UnmarshalOptions: protojson.UnmarshalOptions{
// 			DiscardUnknown: true,
// 		},
// 	})

// 	grpcMux := runtime.NewServeMux(jsonOption)

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	err = pb.RegisterMomsRecipesHandlerServer(ctx, grpcMux, server)
// 	if err != nil {
// 		log.Fatalf("cannot register handler server: %s", err)
// 	}

// 	mux := http.NewServeMux()
// 	mux.Handle("/", grpcMux)

// 	// Create a sub filesystem that contains the swagger files
// 	swaggerFiles, err := fs.Sub(swagger, "doc/swagger")
// 	if err != nil {
// 		log.Fatalf("cannot create sub filesystem: %s", err)
// 	}

// 	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.FS(swaggerFiles))))

// 	listener, err := net.Listen("tcp", config.HTTPServerAddress)
// 	if err != nil {
// 		log.Fatal("cannot create listener")
// 	}

// 	log.Printf("start HTTP gateway server at %s", listener.Addr().String())
// 	err = http.Serve(listener, mux)
// 	if err != nil {
// 		log.Fatal("cannot start HTTP gateway server")
// 	}
// }

// // func runGinServer(config util.Config, store db.Store) {
// // 	server, err := api.NewServer(config, store)
// // 	if err != nil {
// // 		log.Fatal("cannot create server")
// // 	}

// // 	err = server.Start(config.HTTPServerAddress)
// // 	if err != nil {
// // 		log.Fatal("cannot start server:", err)
// // 	}
// // }
