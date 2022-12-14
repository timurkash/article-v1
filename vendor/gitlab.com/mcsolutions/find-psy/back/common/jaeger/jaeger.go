package jaeger

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"
	semConv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"os"
)

// SetTracerProvider Set global trace provider
func SetTracerProvider(url, serviceName string) error {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return err
	}
	tp := traceSdk.NewTracerProvider(
		// Set the sampling rate based on the parent span to 100%
		traceSdk.WithSampler(traceSdk.ParentBased(traceSdk.TraceIDRatioBased(1.0))),
		// Always be sure to batch in production.
		traceSdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		traceSdk.WithResource(resource.NewSchemaless(
			semConv.ServiceNameKey.String(serviceName),
			attribute.String("env", os.Getenv("KRATOS_ENVIRONMENT")),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}
