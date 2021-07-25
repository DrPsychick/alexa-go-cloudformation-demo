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
	"strings"

	//"github.com/hamba/logger"
	"github.com/hamba/pkg/log"
	//"github.com/hamba/pkg/stats"
	//"github.com/hamba/statter/l2met"
	"gopkg.in/urfave/cli.v1"
	//"os"
	"time"
)

func runLambda(c *cli.Context) error {
	ctx, err := cmd.NewContext(c)
	if err != nil {
		return err
	}

	// attach a unbuffered logger:
	lg := logger.New(logger.StreamHandler(os.Stdout, logger.LogfmtFormat()))
	ctx.AttachLogger(func(l log.Logger) log.Logger {
		return lg
	})

	// new statter, using unbuffered logger
	// TODO: this is broken: when running lambda locally in docker context flags are empty
	log.Info(ctx, "DSN: "+c.String(cmd.FlagStatsDSN))
	log.Info(ctx, "LogLevel: "+c.String(cmd.FlagLogLevel))
	log.Info(ctx, fmt.Sprintf("global: %s flags: %s", strings.Join(c.GlobalFlagNames(), ","), strings.Join(c.FlagNames(), ",")))
	st, err := cmd.NewStats(c, lg)
	if err != nil {
		return err
	}
	if st == stats.Null {
		log.Info(ctx, fmt.Sprintf("Flag '%s' is empty!", cmd.FlagStatsDSN))
	}
	ctx.AttachStatter(func(s stats.Statter) stats.Statter {
		return st
		//return l2met.New(lg, "")
	})
	//st.Gauge("bar", 123, 1.0)
	stats.Gauge(ctx, "foo", 234, 1.0)
	stats.Inc(ctx, "foo", 5, 1.0)

	app, err := newApplication(ctx)
	if err != nil {
		log.Fatal(ctx, err.Error())
	}
	sb := newSkill()
	l := newLambda(app, sb)
	// -> mux.HandleIntent("MyIntent", handleMyIntent(app), WithSlot("MySlot"))
	//sb.ModelFromIntentProvider(l)

	//l := newLambda(app, sb)

	//l.HandleIntent(loca.DemoIntent, lambda.HandleDemoIntent()).WithSlot(A)

	ms, err := sb.BuildModels()
	if err != nil {
		log.Fatal(ctx, err)
	}
	for l, m := range ms {
		log.Info(ctx, fmt.Sprintf("accepting locale '%s' invocation '%s'", l, m.Model.Language.Invocation))
	}
	// this is logged, but logger waits 1 sec before logging...
	defer log.Info(ctx, "just for fun")
	defer time.Sleep(1 * time.Second)
	log.Info(ctx, "will also appear immediately")
	defer ctx.Close()

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
