package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	restcontrollers "github.com/shubha-intelops/sphara/sphara/pkg/rest/server/controllers"
	"github.com/sinhashubham95/go-actuator"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/credentials"
	"os"
)

func main() {

	router := gin.Default()
	if len(serviceName) > 0 && len(collectorURL) > 0 {
		// add opentel
		cleanup := initTracer()
		defer func(func(context.Context) error) {
			_ = cleanup(context.Background())
		}(cleanup)
		router.Use(otelgin.Middleware(serviceName))
	}

	// add actuator
	addActuator(router)
	// add prometheus
	addPrometheus(router)

	signupController, err := restcontrollers.NewSignupController()
	if err != nil {
		log.Errorf("error occurred: %s", err)
		return
	}

	medicalconditionController, err := restcontrollers.NewMedicalConditionController()
	if err != nil {
		log.Errorf("error occurred: %s", err)
		return
	}

	medicalinsuranceController, err := restcontrollers.NewMedicalInsuranceController()
	if err != nil {
		log.Errorf("error occurred: %s", err)
		return
	}

	uploadidController, err := restcontrollers.NewUploadIdController()
	if err != nil {
		log.Errorf("error occurred: %s", err)
		return
	}

	emgcontactController, err := restcontrollers.NewEmgContactController()
	if err != nil {
		log.Errorf("error occurred: %s", err)
		return
	}

	ambulanceController, err := restcontrollers.NewAmbulanceController()
	if err != nil {
		log.Errorf("error occurred: %s", err)
		return
	}

	v1 := router.Group("/v1")
	{

		v1.GET("/signups/:id", signupController.FetchSignup)
		v1.POST("/signups", signupController.CreateSignup)
		v1.PUT("/signups/:id", signupController.UpdateSignup)
		v1.DELETE("/signups/:id", signupController.DeleteSignup)
		v1.GET("/signups", signupController.ListSignups)
		v1.PATCH("/signups/:id", signupController.PatchSignup)
		v1.HEAD("/signups", signupController.HeadSignup)
		v1.OPTIONS("/signups", signupController.OptionsSignup)

		v1.GET("/medicalconditions/:id", medicalconditionController.FetchMedicalCondition)
		v1.POST("/medicalconditions", medicalconditionController.CreateMedicalCondition)
		v1.PUT("/medicalconditions/:id", medicalconditionController.UpdateMedicalCondition)
		v1.DELETE("/medicalconditions/:id", medicalconditionController.DeleteMedicalCondition)
		v1.GET("/medicalconditions", medicalconditionController.ListMedicalConditions)
		v1.PATCH("/medicalconditions/:id", medicalconditionController.PatchMedicalCondition)
		v1.HEAD("/medicalconditions", medicalconditionController.HeadMedicalCondition)
		v1.OPTIONS("/medicalconditions", medicalconditionController.OptionsMedicalCondition)

		v1.GET("/medicalinsurances/:id", medicalinsuranceController.FetchMedicalInsurance)
		v1.POST("/medicalinsurances", medicalinsuranceController.CreateMedicalInsurance)
		v1.PUT("/medicalinsurances/:id", medicalinsuranceController.UpdateMedicalInsurance)
		v1.DELETE("/medicalinsurances/:id", medicalinsuranceController.DeleteMedicalInsurance)
		v1.GET("/medicalinsurances", medicalinsuranceController.ListMedicalInsurances)
		v1.PATCH("/medicalinsurances/:id", medicalinsuranceController.PatchMedicalInsurance)
		v1.HEAD("/medicalinsurances", medicalinsuranceController.HeadMedicalInsurance)
		v1.OPTIONS("/medicalinsurances", medicalinsuranceController.OptionsMedicalInsurance)

		v1.GET("/uploadids/:id", uploadidController.FetchUploadId)
		v1.POST("/uploadids", uploadidController.CreateUploadId)
		v1.PUT("/uploadids/:id", uploadidController.UpdateUploadId)
		v1.DELETE("/uploadids/:id", uploadidController.DeleteUploadId)
		v1.GET("/uploadids", uploadidController.ListUploadIds)
		v1.PATCH("/uploadids/:id", uploadidController.PatchUploadId)
		v1.HEAD("/uploadids", uploadidController.HeadUploadId)
		v1.OPTIONS("/uploadids", uploadidController.OptionsUploadId)

		v1.GET("/emgcontacts/:id", emgcontactController.FetchEmgContact)
		v1.POST("/emgcontacts", emgcontactController.CreateEmgContact)
		v1.PUT("/emgcontacts/:id", emgcontactController.UpdateEmgContact)
		v1.DELETE("/emgcontacts/:id", emgcontactController.DeleteEmgContact)
		v1.GET("/emgcontacts", emgcontactController.ListEmgContacts)
		v1.PATCH("/emgcontacts/:id", emgcontactController.PatchEmgContact)
		v1.HEAD("/emgcontacts", emgcontactController.HeadEmgContact)
		v1.OPTIONS("/emgcontacts", emgcontactController.OptionsEmgContact)

		v1.GET("/ambulances/:id", ambulanceController.FetchAmbulance)
		v1.POST("/ambulances", ambulanceController.CreateAmbulance)
		v1.PUT("/ambulances/:id", ambulanceController.UpdateAmbulance)
		v1.DELETE("/ambulances/:id", ambulanceController.DeleteAmbulance)
		v1.GET("/ambulances", ambulanceController.ListAmbulances)
		v1.PATCH("/ambulances/:id", ambulanceController.PatchAmbulance)
		v1.HEAD("/ambulances", ambulanceController.HeadAmbulance)
		v1.OPTIONS("/ambulances", ambulanceController.OptionsAmbulance)

	}

	Port := ":8080"
	log.Println("Server started")
	if err = router.Run(Port); err != nil {
		log.Errorf("error occurred: %s", err)
		return
	}

}

var (
	serviceName  = os.Getenv("SERVICE_NAME")
	collectorURL = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	insecure     = os.Getenv("INSECURE_MODE")
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func addPrometheus(router *gin.Engine) {
	router.GET("/metrics", prometheusHandler())
}

func addActuator(router *gin.Engine) {
	actuatorHandler := actuator.GetActuatorHandler(&actuator.Config{Endpoints: []int{
		actuator.Env,
		actuator.Info,
		actuator.Metrics,
		actuator.Ping,
		// actuator.Shutdown,
		actuator.ThreadDump,
	},
		Env:     "dev",
		Name:    "sphara",
		Port:    8080,
		Version: "0.0.1",
	})
	ginActuatorHandler := func(ctx *gin.Context) {
		actuatorHandler(ctx.Writer, ctx.Request)
	}
	router.GET("/actuator/*endpoint", ginActuatorHandler)
}

func initTracer() func(context.Context) error {
	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if len(insecure) > 0 {
		secureOption = otlptracegrpc.WithInsecure()
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(collectorURL),
		),
	)

	if err != nil {
		log.Errorf("error occurred: %s", err)
		return nil
	}
	restResources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("services.name", serviceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Printf("could not set restResources: %s", err)
	}

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(restResources),
		),
	)
	return exporter.Shutdown
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}
