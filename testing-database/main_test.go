package main

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"

	"testing-database/pack"
)

func TestCreateTable(t *testing.T) {
	pack.ConnectDB(pack.ConnString)

	type args struct {
		ctx  context.Context
		conn *pgx.Conn
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{
			ctx:  context.Background(),
			conn: pack.Connection,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateTable(tt.args.ctx, tt.args.conn); (err != nil) != tt.wantErr {
				t.Errorf("CreateTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateTable1(t *testing.T) {
	pack.ConnectDB(pack.ConnString)

	type args struct {
		ctx  context.Context
		conn *pgx.Conn
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test111", args{
			ctx:  context.Background(),
			conn: pack.Connection,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateTable(tt.args.ctx, tt.args.conn); (err != nil) != tt.wantErr {
				t.Errorf("CreateTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
