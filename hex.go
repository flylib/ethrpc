package ethrpc

import (
	"bytes"
	"math/big"
)

type hexInt int

//hexInt解锁JSON
func (i *hexInt) UnmarshalJSON(data []byte) error {
	result, err := ParseInt(string(bytes.Trim(data, `"`)))
	*i = hexInt(result)

	return err
}

type hexBig big.Int

//hexBig解锁JSON
func (i *hexBig) UnmarshalJSON(data []byte) error {
	result, err := ParseBigInt(string(bytes.Trim(data, `"`)))
	*i = hexBig(result)

	return err
}

