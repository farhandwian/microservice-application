package app

import (
	"context"
	"strings"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type middleware struct {
}

// Define middleware untuk melakukan verifikasi token JWT.
func jwtMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Mendapatkan token dari metadata permintaan gRPC.
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "No metadata found in request")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "Authorization token not found")
	}

	// Memisahkan token dari header "Bearer".
	tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")

	// Validasi dan parse token JWT.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ganti "your-secret-key" dengan secret key yang digunakan untuk menandatangani token JWT.
		// Secret key ini harus sama dengan secret key yang digunakan untuk meng-generate token.
		// Anda dapat menaruh secret key ini di environment variable atau tempat yang aman.
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid token: %v", err)
	}

	if !token.Valid {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
	}

	// Lanjutkan dengan mengeksekusi gRPC handler jika token valid.
	return handler(ctx, req)
}
