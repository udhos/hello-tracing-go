package traceutil

import (
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func SetPropagation() {
	// In order to propagate trace context over the wire, a propagator must be registered with the OpenTelemetry API.
	// https://opentelemetry.io/docs/instrumentation/go/manual/
	//otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader)),
		//propagation.Baggage{},
		//propagation.TraceContext{},
		//ot.OT{},
	))
}
