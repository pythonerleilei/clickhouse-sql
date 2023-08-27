grammar ClickHouseParser;

import ClickhouseLexer;

options {
    tokenVocab = ClickhouseLexer;
}

statement
    : createStatement
    | alterStatement
    ;


createStatement
    : createDatabaseStatement
    | createTableStatement
    ;

createDatabaseStatement
    : CREATE DATABASE
        (IF NOT EXISTS)? 
        db_name=ID 
        (ON CLUSTER cluster_name=ID)? 
        (ENGINE EQUAL engine_name=expression)?
        (COMMENT comment=STRING)?
    ;

createTableStatement
    : CREATE TABLE
        (IF NOT EXISTS)?
        table_name=talbeIdentifier
    ;

alterStatement
    :   alterTableColumn
    ;

alterTableColumn
    : ALTER TABLE 
        table_name=talbeIdentifier 
        (ON CLUSTER cluster_name=ID)? 
        ADD COLUMN 
        (IF NOT EXISTS)? 
        column_type=columnType

    ;

talbeIdentifier: (ID DOT)? ID;

columnType: expression;

expression: ID (LEFT_P parameters? RIGHT_P)? ;

parameters: parameter (COMMA parameter)* ;

parameter
    : NUMBER
    | STRING
    | expression
    ;