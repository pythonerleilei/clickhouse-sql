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
        (ON CLUSTER cluster_name=ID)? 
        LEFT_P
        column_defs
        (COMMA index_defs) ?
        RIGHT_P
        ENGINE EQUAL engine=ID ( LEFT_P engine_paras=parameters RIGHT_P)?
        ORDER BY oder_by=expression
        (PARTITION BY partition_by=expression)?
        (PRIMARY KEY primary_key=expression)?
        (SAMPLE BY sample_by=expression)?
        (ttl_def)?
        (SETTINGS kvs=key_values)?
        (COMMENT STRING)?
    ;

column_defs: column_def (COMMA column_def)* ;

column_def
        : name=ID data_type=expression 
            (DEFAULT default_value=expression) ?
            (CODEC LEFT_P codec=ID RIGHT_P) ?
            (COMMENT STRING)?
        ;

index_defs: index_def (COMMA index_def)* ;
index_def
        : INDEX ID eval=expression TYPE type_exp=expression (GRANULARITY NUMBER)? ;

ttl_def
    : TTL column=ID PLUS INTERVAL NUMBER unit=(SECOND | MINUTE | HOUR | DAY | WEEK | MONTH | QUARTER | YEAR) 
    ;

key_values: key_value (COMMA key_value)* ;
key_value: ID EQUAL parameter ;

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

expression
    : ID (LEFT_P parameters? RIGHT_P)? 
    | LEFT_P parameters? RIGHT_P
    ;

parameters: parameter (COMMA parameter)* ;

parameter
    : NUMBER
    | STRING
    | expression
    ;