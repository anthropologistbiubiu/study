package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/prometheus/client_golang/prometheus"
	"os"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"payhub/internal/conf"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	//Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

var (
	// Name is the name of the compiled software.
	Name = "metrics"
	// Version is the version of the compiled software.
	// Version = "v1.0.0"

	_metricSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "server",
		Subsystem: "requests",
		Name:      "duration_sec",
		Help:      "server requests duration(sec).",
		Buckets:   []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.250, 0.5, 1},
	}, []string{"kind", "operation"})

	_metricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "client",
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "The total number of processed requests",
	}, []string{"kind", "operation", "code", "reason"})
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {

	/*
		consulConfig := api.DefaultConfig()
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			log.Fatalf("failed to create consul client: %v", err)
		}

		// 创建 Kratos 的 Consul 注册器
		//r := consulRegistry.Registrar(context.Context(), consulClient)

	*/
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		//kratos.Registrar(r),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {

	flag.Parse()
	/*
		logger := log.With(log.NewStdLogger(os.Stdout),
			"ts", log.DefaultTimestamp,
			"caller", log.DefaultCaller,
			"service.id", id,
			"service.name", Name,
			"service.version", Version,
			"trace.id", tracing.TraceID(),
			"span.id", tracing.SpanID(),
		)
	*/
	infoFile, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	errorFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer infoFile.Close()
	defer errorFile.Close()

	// 创建不同级别的日志处理器
	infoLogger := log.With(log.NewStdLogger(infoFile),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"level", log.LevelInfo)
	/*
		errorLogger := log.With(log.NewStdLogger(errorFile),
			"ts", log.DefaultTimestamp,
			"caller", log.DefaultCaller,
			"level", log.LevelError)
	*/
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, infoLogger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
