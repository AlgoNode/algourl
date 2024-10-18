package encoder

import (
	"fmt"

	"github.com/algorand/go-algorand-sdk/v2/encoding/msgpack"
	"github.com/algorand/go-algorand-sdk/v2/types"
	qrcode "github.com/algorandfoundation/go-tinyqr"
	"github.com/google/go-querystring/query"
)

const ARC0026URLHANDLER = "algorand"

type AUrlTxn struct {
	AUrlTxnKeyCommon
	AUrlTxnKeyreg
}

type AUrlTxnKeyCommon struct {
	Sender string `url:"-"`
	Type   string `url:"type"`
	Fee    *uint64 `url:"fee,omitempty"`
}

type AUrlTxnKeyreg struct {
	VotePK          *string `url:"votekey,omitempty"`
	SelectionPK     *string `url:"selkey,omitempty"`
	StateProofPK    *string `url:"sprfkey,omitempty"`
	VoteFirst       *uint64 `url:"votefst,omitempty"`
	VoteLast        *uint64 `url:"votelst,omitempty"`
	VoteKeyDilution *uint64 `url:"votekd,omitempty"`
}

func (krg AUrlTxn) String() string {
	v, _ := query.Values(krg.AUrlTxnKeyCommon)
	url := fmt.Sprintf("%s://%s?%s",
		ARC0026URLHANDLER,
		krg.Sender,
		v.Encode(),
	)
	url += krg.AUrlTxnKeyreg.String()
	return url
}

func (krg AUrlTxnKeyreg) String() string {
	v, _ := query.Values(krg)
	if len(v) == 0 {
		return ""
	}
	return "&" + v.Encode()
}

type RawTxn struct {
	Txn types.Transaction `codec:"txn"`
}

func MakeQRKeyRegRequest(encodedTxn []byte) (*AUrlTxn, error) {
	var txn RawTxn

	msgpack.Decode(encodedTxn, &txn)
	if txn.Txn.Type != "keyreg" {
		return nil, fmt.Errorf("no support for transaction of type '%s'", txn.Txn.Type)
	}

	kr := &AUrlTxn{
		AUrlTxnKeyCommon: AUrlTxnKeyCommon{
			Sender: txn.Txn.Sender.String(),
			Type:   string(txn.Txn.Type),
		},
		AUrlTxnKeyreg: AUrlTxnKeyreg{
			VotePK:          urlEncodeBytesPtrOrNil(txn.Txn.VotePK[:]),
			SelectionPK:     urlEncodeBytesPtrOrNil(txn.Txn.SelectionPK[:]),
			StateProofPK:    urlEncodeBytesPtrOrNil(txn.Txn.StateProofPK[:]),
			VoteFirst:       toPtr(uint64(txn.Txn.VoteFirst)),
			VoteLast:        toPtr(uint64(txn.Txn.VoteLast)),
			VoteKeyDilution: toPtrOrNil(txn.Txn.VoteKeyDilution),
		},
	}
	if uint64(txn.Txn.Fee) != uint64(1000) {
		kr.AUrlTxnKeyCommon.Fee =	toPtr(uint64(txn.Txn.Fee))
	}

	return kr, nil
}

func (krg AUrlTxn) Print() {
	url := krg.String()
	fmt.Println("Paste below URL into your browser or scan QR code to online/offline the account")
	fmt.Println(url)
	fmt.Println()
	qrcode.Print(url)
}
