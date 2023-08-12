grammar ClickHouseParser;

import ClickhouseLexer;

options {
    tokenVocab = ClickhouseLexer;
}

statement
    : createStatement;


createStatement
    : createDatabaseStatement;

createDatabaseStatement
    : CREATE DATABASE
        (IF NOT EXISTS)? 
        db_name=ID 
        (ON CLUSTER cluster_name=ID)? 
        (ENGINE EQUAL engine_name=ID (LEFT_P engine_params+=parameter (COMMA parameter)* RIGHT_P)?)?
        (COMMENT comment=STRING)?
    ;

parameter
    : NUMBER
    | STRING
    ;