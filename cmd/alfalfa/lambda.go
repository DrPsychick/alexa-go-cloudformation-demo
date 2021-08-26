package main

import (
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo"
	"github.com/drpsychick/alexa-go-cloudformation-demo/lambda"
	"github.com/drpsychick/alexa-go-cloudformation-demo/lambda/middleware"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
	"github.com/hamba/cmd"
	"github.com/hamba/logger"
	"github.com/hamba/pkg/stats"
	"os"
	//"github.com/hamba/logger"
	"github.com/hamba/pkg/log"
	//"github.com/hamba/pkg/stats"
	//"github.com/hamba/statter/l2met"
	"gopkg.in/urfave/cli.v1"
	//"os"
	"time"
)

func runLambda(c *cli.Context) error {
	start := time.Now()

	ctx, err := cmd.NewContext(c)
	if err != nil {
		return err
	}

	// attach a unbuffered logger:
	lg := logger.New(logger.StreamHandler(os.Stdout, logger.LogfmtFormat()))
	ctx.AttachLogger(func(l log.Logger) log.Logger {
		return lg
	})

	st, err := cmd.NewStats(c, lg)
	if err != nil {
		return err
	}
	if st == stats.Null {
		log.Info(ctx, fmt.Sprintf("Flag '%s' is empty!", cmd.FlagStatsDSN))
	}
	ctx.AttachStatter(func(s stats.Statter) stats.Statter {
		return st
	})

	app, err := newApplication(ctx)
	if err != nil {
		log.Fatal(ctx, err.Error())
	}
	stats.Timing(ctx, "Boot", time.Since(start), 1.0)
	sb := newSkill()
	l := newLambda(app, sb)

	ms, err := sb.BuildModels()
	if err != nil {
		log.Fatal(ctx, err)
	}
	for l, m := range ms {
		log.Info(ctx, fmt.Sprintf("accepting locale '%s' invocation '%s'", l, m.Model.Language.Invocation))
	}
	defer ctx.Close()

	stats.Timing(ctx, "Ready", time.Since(start), 1.0)
	if err := alexa.Serve(l); err != nil {
		log.Fatal(ctx, err)
	}

	log.Fatal(ctx, "Serve() should not have returned")
	return nil
}

func newLambda(app *alfalfa.Application, sb *gen.SkillBuilder) alexa.Handler {
	h := lambda.NewMux(app, sb)

	h = middleware.WithRequestStats(h, app)
	return middleware.WithRecovery(h, app)
}
