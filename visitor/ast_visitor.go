package visitor

import (
	parser "clickhouse-sql/parser"
)

type BaseAstClickhouseParserVisitor struct {
	*parser.BaseClickHouseParserVisitor
}

