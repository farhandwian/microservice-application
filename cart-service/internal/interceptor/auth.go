package interceptorr

import (
	"context"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Define middleware untuk melakukan verifikasi token JWT.
// func jwtMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
// 	// Mendapatkan token dari metadata permintaan gRPC.
// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if !ok {
// 		return nil, status.Errorf(codes.Unauthenticated, "No metadata found in request")
// 	}

// 	authHeader, ok := md["authorization"]
// 	if !ok || len(authHeader) == 0 {
// 		return nil, status.Errorf(codes.Unauthenticated, "Authorization token not found")
// 	}

// 	// Memisahkan token dari header "Bearer".
// 	tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")

// 	// Validasi dan parse token JWT.
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Ganti "your-secret-key" dengan secret key yang digunakan untuk menandatangani token JWT.
// 		// Secret key ini harus sama dengan secret key yang digunakan untuk meng-generate token.
// 		// Anda dapat menaruh secret key ini di environment variable atau tempat yang aman.
// 		return []byte("your-secret-key"), nil
// 	})

// 	if err != nil {
// 		return nil, status.Errorf(codes.Unauthenticated, "Invalid token: %v", err)
// 	}

// 	if !token.Valid {
// 		return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
// 	}

// 	// Lanjutkan dengan mengeksekusi gRPC handler jika token valid.
// 	return handler(ctx, req)
// }

// AuthInterceptor represents the struct for auth interceptor.
type AuthInterceptor struct {
	secretKey string
}

func NewAuthInterceptor(secretKey string) *AuthInterceptor {
	return &AuthInterceptor{secretKey: secretKey}
}

// Unary returns a server interceptor function to authenticate and authorize unary RPC.
func (a *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := a.authorize(ctx); err != nil {
			return nil, err
		}

		// Pass the context and request to the next handler in the chain.
		return handler(ctx, req)
	}
}

// authorize checks the context for the correct auth token.
func (a *AuthInterceptor) authorize(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	tokenStr := values[0]
	_, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Errorf(codes.Unauthenticated, "unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(a.secretKey), nil
	})

	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return nil
}
