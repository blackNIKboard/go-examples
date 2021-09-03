package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"testing"
)

func TestNewRoot(t *testing.T) {
	type args struct {
		num *big.Int
		deg int
	}
	var tests []struct {
		name    string
		args    args
		want    big.Int
		wantErr bool
	}

	for i := 7; i <= 10; i++ {
		midDeg := int(math.Mod(float64(i), 7) + 2)
		midArgs := args{
			num: new(big.Int).Exp(new(big.Int).SetInt64(int64(i)), new(big.Int).SetInt64(int64(midDeg)), nil),
			deg: midDeg,
		}
		midWant := new(big.Int).SetInt64(int64(i))
		name := "test" + strconv.Itoa(i)

		spew.Dump(name, midArgs, midWant)

		tests = append(tests, struct {
			name    string
			args    args
			want    big.Int
			wantErr bool
		}{
			name:    name,
			args:    midArgs,
			want:    *midWant,
			wantErr: false,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRoot(tt.args.num, tt.args.deg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRoot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("NewRoot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SumOfDigits(t *testing.T) {
	type args struct {
		num *big.Int
	}
	var tests []struct {
		name    string
		args    args
		want    big.Int
		wantErr bool
	}

	for i := 1; i <= 1521; i++ {
		midSum := 0
		midStr := ""
		name := "test" + strconv.Itoa(i)

		for j := 1; j <= i; j++ {
			tmp := math.Mod(float64(j), 10)
			midSum += int(tmp)
			midStr += strconv.Itoa(int(tmp))
		}

		midStrI, ok := new(big.Int).SetString(midStr, 10)
		if !ok {
			panic(fmt.Errorf("error converting string"))
		}

		tests = append(tests, struct {
			name    string
			args    args
			want    big.Int
			wantErr bool
		}{
			name:    name,
			args:    args{num: midStrI},
			want:    *new(big.Int).SetInt64(int64(midSum)),
			wantErr: false})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumOfDigits(tt.args.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("sumOfDigits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("sumOfDigits() got = %v, want %v", got, tt.want)
			}
		})
	}
}
