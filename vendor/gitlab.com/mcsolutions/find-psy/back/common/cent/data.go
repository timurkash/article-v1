package cent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	"github.com/go-kratos/kratos/v2/log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func SetSchema(driver *sql.Driver, schema string) error {
	if schema != "" && schema != "public" {
		if _, err := driver.DB().Exec(
			fmt.Sprintf(
				"create schema if not exists %s;set search_path to %s;",
				schema,
				schema,
			),
		); err != nil {
			return err
		}
	}
	return nil
}

func DebugWithContext(driver *sql.Driver, logHelper *log.Helper) dialect.Driver {
	return dialect.DebugWithContext(driver, func(ctx context.Context, i ...interface{}) {
		logHelper.WithContext(ctx).Info(i...)
		_, span := otel.Tracer("entgo.io").Start(ctx,
			"Query",
			trace.WithAttributes(
				attribute.String("sql", fmt.Sprint(i...)),
			),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		span.End()
	})
}
