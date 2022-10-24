package handlers

import (
	"context"
	"gitlab.com/tokene/keyserver-svc/internal/config"
	"gitlab.com/tokene/keyserver-svc/internal/data"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	kdfCtxKey
	walletsCtxKey
	walletsConfigCtxKey
	emailTokensCtxKey
	tfaConfigCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxKDFQ(entry data.KDFQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, kdfCtxKey, entry)
	}
}

func KDFQ(r *http.Request) data.KDFQ {
	return r.Context().Value(kdfCtxKey).(data.KDFQ).New()
}

func CtxWalletsQ(entry data.WalletsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, walletsCtxKey, entry)
	}
}

func WalletsQ(r *http.Request) data.WalletsQ {
	return r.Context().Value(walletsCtxKey).(data.WalletsQ).New()
}

func CtxWalletsConfig(wallets config.WalletsConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, walletsConfigCtxKey, wallets)
	}
}

func WalletsConfig(r *http.Request) config.WalletsConfig {
	return r.Context().Value(walletsConfigCtxKey).(config.WalletsConfig)
}

func CtxEmailTokensQ(entry data.EmailTokensQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, emailTokensCtxKey, entry)
	}
}

func EmailTokensQ(r *http.Request) data.EmailTokensQ {
	return r.Context().Value(emailTokensCtxKey).(data.EmailTokensQ).New()
}

func CtxTFAConfig(wallets config.TFAConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, tfaConfigCtxKey, wallets)
	}
}

func TFAConfig(r *http.Request) config.TFAConfig {
	return r.Context().Value(tfaConfigCtxKey).(config.TFAConfig)
}
