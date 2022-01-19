package util

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
)

type tool struct {
}

var Tool *tool

// IsValidAddress 判断是否为有效地址
func (*tool) IsValidAddress(address interface{}) bool {
	reg := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := address.(type) {
	case string:
		return reg.MatchString(v)
	case common.Address:
		return reg.MatchString(v.Hex())
	default:
		return false
	}
}

func (*tool) IsZeroAddress(address interface{}) bool {
	var addr common.Address
	switch v := address.(type) {
	case string:
		addr = common.HexToAddress(v)
	case common.Address:
		addr = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := addr.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

// ToDecimal 将wei转换为decimals
func (*tool) ToDecimal(value interface{}, decimals int) decimal.Decimal {
	val := new(big.Int)
	switch v := value.(type) {
	case string:
		val.SetString(v, 10)
	case *big.Int:
		val = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(val.String())
	result := num.Div(mul)

	return result
}

// ToWei 将decimals转换为wei
func (*tool) ToWei(amount interface{}, decimals int) *big.Int {
	a := decimal.NewFromFloat(0)
	switch v := amount.(type) {
	case string:
		a, _ = decimal.NewFromString(v)
	case float64:
		a = decimal.NewFromFloat(v)
	case int64:
		a = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		a = v
	case *decimal.Decimal:
		a = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := a.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

// CalcGasCost 通过给定的gas limit (uint) 和gas price (wei) 计算gas消耗
func (*tool) CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
	gasLimitBig := big.NewInt(int64(gasLimit))
	return gasLimitBig.Mul(gasLimitBig, gasPrice)
}

// SigRSV 解构R,S,V
func (*tool) SigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := isig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}

	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vS := sigstr[128:130]
	vI, _ := strconv.Atoi(vS)
	V := uint8(vI + 27)

	return R, S, V
}

