package database

import "github.com/Masterminds/squirrel"

var SqlBuilder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
