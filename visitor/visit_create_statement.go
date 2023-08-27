package visitor

import (
	parser "clickhouse-sql/parser"
	"strconv"
)

func (v *BaseAstClickhouseParserVisitor) VisitCreateStatement(ctx *parser.CreateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

type CreateDatabase struct {
	FlagNotExists	bool
	DatabaseName	string
	ClusterName		string
	EngineName		string
	EngineParams	[]*Expression
	Comment			string
}

// create database
func (v *BaseAstClickhouseParserVisitor) VisitCreateDatabaseStatement(ctx *parser.CreateDatabaseStatementContext) interface{} {
	createDatabase := &CreateDatabase{}
	if ctx.EXISTS() != nil {
		createDatabase.FlagNotExists = true
	}

	createDatabase.DatabaseName = ctx.GetDb_name().GetText();
	if (ctx.GetCluster_name() != nil) {
		createDatabase.ClusterName = ctx.GetCluster_name().GetText()
	}
	
	if ctx.GetEngine_name() != nil {
		e := v.Visit(ctx.GetEngine_name()).(*Expression)
		createDatabase.EngineParams = e.children
		createDatabase.EngineName = e.Name()
	}
	
	if (ctx.COMMENT() != nil ){
		createDatabase.Comment = ctx.GetComment().GetText();
	}
	return createDatabase
}

type CreateTable struct {
	FlagNotExists			bool
	TableName				string
	ClusterName				string
	ColumnInfos				[]*ColumnInfo
	IndexInfos				[]*IndexInfo
	EngineName				string
	EngineParams			[]*Expression
	OrderByExps				[]*Expression
	PartitionByExps			[]*Expression
	PrimaryKeys				[]*Expression
	SampleByExps			[]*Expression
	TtlInfo					*TtlInfo
	SettingInfos			[]*SettingInfo
	Comment					string
}

type ColumnInfo struct {
	Name		string
	Type		*Expression
	Default		*Expression
	Codec		string
	Comment		string
}

type IndexInfo struct {
	Name		string
	EvalExp		*Expression
	TypeExp		*Expression
	Granularity	int
}

type TtlInfo struct {
	ColumnName	string
	Interval	int
	Unit		string
}

type SettingInfo struct {
	Name	string
	Value	string
}

// create table
func (v *BaseAstClickhouseParserVisitor) VisitCreateTableStatement(ctx *parser.CreateTableStatementContext) interface{} {
	createTable := &CreateTable{}
	if ctx.EXISTS() != nil {
		createTable.FlagNotExists = true
	}
	createTable.TableName = ctx.GetTable_name().GetText()
	if ctx.CLUSTER() != nil {
		createTable.ClusterName = ctx.GetCluster_name().GetText()
	}
	createTable.ColumnInfos = v.Visit(ctx.Column_defs()).([]*ColumnInfo)
	if ctx.Index_defs() != nil{
		createTable.IndexInfos = v.Visit(ctx.Index_defs()).([]*IndexInfo)
	}
	createTable.EngineName = ctx.GetEngine().GetText()
	if ctx.GetEngine_paras() != nil {
		createTable.EngineParams = v.Visit(ctx.GetEngine_paras()).([]*Expression)
	}
	createTable.OrderByExps = v.Visit(ctx.GetOder_by()).([]*Expression)
	if ctx.GetPartition_by() != nil {
		createTable.PartitionByExps = v.Visit(ctx.GetPartition_by()).([]*Expression)
	}
	if ctx.GetPrimary_key() != nil {
		createTable.PrimaryKeys = v.Visit(ctx.GetPrimary_key()).([]*Expression)
	}
	if ctx.GetSample_by() != nil {
		createTable.SampleByExps = v.Visit(ctx.GetSample_by()).([]*Expression)
	}
	if ctx.Ttl_def() != nil {
		createTable.TtlInfo = v.Visit(ctx.Ttl_def()).(*TtlInfo)
	}
	if ctx.GetKvs() != nil {
		createTable.SettingInfos = v.Visit(ctx.GetKvs()).([]*SettingInfo)
	}
	if ctx.COMMENT() != nil {
		createTable.Comment = ctx.COMMENT().GetText()
	}
	return createTable
}

func (v *BaseAstClickhouseParserVisitor) VisitColumn_defs(ctx *parser.Column_defsContext) interface{} {
	var columnInfos []*ColumnInfo
	for _, columnDef := range ctx.AllColumn_def() {
		columnInfo := v.Visit(columnDef).(*ColumnInfo)
		columnInfos = append(columnInfos, columnInfo)
	}
	return columnInfos
}

func (v *BaseAstClickhouseParserVisitor) VisitColumn_def(ctx *parser.Column_defContext) interface{} {
	columnInfo := &ColumnInfo{}
	columnInfo.Name = ctx.GetName().GetText()
	columnInfo.Type = v.Visit(ctx.GetData_type()).(*Expression)
	if ctx.GetDefault_value() != nil {
		columnInfo.Default = v.Visit(ctx.GetDefault_value()).(*Expression)
	}
	if ctx.GetCodec() != nil {
		columnInfo.Codec = ctx.GetCodec().GetText()
	}
	if ctx.COMMENT() != nil {
		columnInfo.Comment = ctx.COMMENT().GetText()
	}
	return columnInfo
}

func (v *BaseAstClickhouseParserVisitor) VisitIndex_defs(ctx *parser.Index_defsContext) interface{} {
	var indexInfos []*IndexInfo
	for _, index_def := range ctx.AllIndex_def() {
		indexInfo := v.Visit(index_def).(*IndexInfo)
		indexInfos = append(indexInfos, indexInfo)
	}
	return indexInfos
}

func (v *BaseAstClickhouseParserVisitor) VisitIndex_def(ctx *parser.Index_defContext) interface{} {
	indexInfo := &IndexInfo{}
	indexInfo.Name = ctx.ID().GetText()
	indexInfo.EvalExp = v.Visit(ctx.GetEval()).(*Expression)
	indexInfo.TypeExp = v.Visit(ctx.GetType_exp()).(*Expression)
	if ctx.GRANULARITY() != nil {
		g, err := strconv.Atoi(ctx.NUMBER().GetText())
		if err != nil {
			panic(err.Error())
		}
		indexInfo.Granularity = g
	}
	return indexInfo
}

func (v *BaseAstClickhouseParserVisitor) VisitTtl_def(ctx *parser.Ttl_defContext) interface{} {
	ttlInfo := &TtlInfo{}
	ttlInfo.ColumnName = ctx.ID().GetText()
	ttlInfo.Interval, _ = strconv.Atoi(ctx.NUMBER().GetText())
	ttlInfo.Unit = ctx.GetUnit().GetText()
	return ttlInfo
}

func (v *BaseAstClickhouseParserVisitor) VisitKey_values(ctx *parser.Key_valuesContext) interface{} {
	var settingInfos []*SettingInfo
	for _, kv := range ctx.AllKey_value() {
		settingInfo := v.Visit(kv).(*SettingInfo)
		settingInfos = append(settingInfos, settingInfo)
	}
	return settingInfos
}

func (v *BaseAstClickhouseParserVisitor) VisitKey_value(ctx *parser.Key_valueContext) interface{} {
	settingInfo := &SettingInfo{}
	settingInfo.Name = ctx.ID().GetText()
	settingInfo.Value = ctx.Parameter().GetText()
	return settingInfo
}