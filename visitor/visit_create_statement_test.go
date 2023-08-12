package visitor

import (
	"testing"
)

func TestVisitCreateDatabase(t *testing.T) {
	sql := `create database if not exists 
			test 
			on cluster test_cluster 
			ENGINE=Atomic 
			comment 'test database'`

	astVisitor := NewBaseAstClickhouseParserVisitor(sql)

	tree := astVisitor.Parser.Statement()

	result := astVisitor.Visit(tree)

	cd, ok := result.(*CreateDatabase)
	if !ok {
		t.Fatalf("parse create database failed, sql[%v]", sql)
	}

	if !cd.FlagNotExists {
		t.Fatalf("parse create database failed, sql[%v]\n, flag not exits expected[%v], get[%v]", sql, true, cd.FlagNotExists)
	}

	if cd.DatabaseName != "test" {
		t.Fatalf("parse create database failed, sql[%v]\n, database name expected[%v], get[%v]", sql, "test", cd.DatabaseName)
	}
	
	if cd.ClusterName != "test_cluster" {
		t.Fatalf("parse create database failed, sql[%v]\n, cluster name expected[%v], get[%v]", sql, "test_cluster", cd.ClusterName)
	}

	if cd.EngineName != "Atomic" {
		t.Fatalf("parse create database failed, sql[%v]\n, engine name expected[%v], get[%v]", sql, "Atomic", cd.EngineName)
	}

	if cd.Comment != "'test database'" {
		t.Fatalf("parse create database failed, sql[%v]\n, engine name expected[%v], get[%v]", sql, "'test database'", cd.Comment)
	}
}