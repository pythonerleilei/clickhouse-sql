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
        (ENGINE EQUAL ID (LEFT_P engine_params=parameters RIGHT_P)?)?
        COMMENT comment=STRING
    ;

parameters: parameter (COMMA parameter)*;
parameter
    : NUMBER
    | STRING
    ;