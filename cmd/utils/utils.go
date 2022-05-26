package utils

import (
	"github.com/xpaymentsorg/go-xpayments/XPSx"
	"github.com/xpaymentsorg/go-xpayments/XPSxlending"
	"github.com/xpaymentsorg/go-xpayments/eth"
	"github.com/xpaymentsorg/go-xpayments/eth/downloader"
	"github.com/xpaymentsorg/go-xpayments/ethstats"
	"github.com/xpaymentsorg/go-xpayments/les"
	"github.com/xpaymentsorg/go-xpayments/node"
	whisper "github.com/xpaymentsorg/go-xpayments/whisper/whisperv6"
)

// RegisterEthService adds an Ethereum client to the stack.
func RegisterEthService(stack *node.Node, cfg *eth.Config) {
	var err error
	if cfg.SyncMode == downloader.LightSync {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			return les.New(ctx, cfg)
		})
	} else {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			var XPSXServ *XPSx.XPSX
			ctx.Service(&XPSXServ)
			var lendingServ *XPSxlending.Lending
			ctx.Service(&lendingServ)
			fullNode, err := eth.New(ctx, cfg, XPSXServ, lendingServ)
			if fullNode != nil && cfg.LightServ > 0 {
				ls, _ := les.NewLesServer(fullNode, cfg)
				fullNode.AddLesServer(ls)
			}
			return fullNode, err
		})
	}
	if err != nil {
		Fatalf("Failed to register the Ethereum service: %v", err)
	}
}

// RegisterShhService configures Whisper and adds it to the given node.
func RegisterShhService(stack *node.Node, cfg *whisper.Config) {
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return whisper.New(cfg), nil
	}); err != nil {
		Fatalf("Failed to register the Whisper service: %v", err)
	}
}

// RegisterEthStatsService configures the Ethereum Stats daemon and adds it to
// th egiven node.
func RegisterEthStatsService(stack *node.Node, url string) {
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		// Retrieve both eth and les services
		var ethServ *eth.Ethereum
		ctx.Service(&ethServ)

		var lesServ *les.LightEthereum
		ctx.Service(&lesServ)

		return ethstats.New(url, ethServ, lesServ)
	}); err != nil {
		Fatalf("Failed to register the Ethereum Stats service: %v", err)
	}
}

func RegisterXPSXService(stack *node.Node, cfg *XPSx.Config) {
	XPSX := XPSx.New(cfg)
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return XPSX, nil
	}); err != nil {
		Fatalf("Failed to register the XPSX service: %v", err)
	}

	// register XPSxlending service
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return XPSxlending.New(XPSX), nil
	}); err != nil {
		Fatalf("Failed to register the XPSXLending service: %v", err)
	}
}
