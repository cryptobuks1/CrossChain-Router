package mongodb

const (
	tbRouterSwaps       string = "RouterSwaps"
	tbRouterSwapResults string = "RouterSwapResults"
)

// MgoSwap registered swap
type MgoSwap struct {
	Key         string `bson:"_id"` // fromChainID + txid + logindex
	SwapType    uint32 `bson:"swaptype"`
	TxID        string `bson:"txid"`
	TxTo        string `bson:"txto"`
	Bind        string `bson:"bind"`
	LogIndex    int    `bson:"logIndex"`
	FromChainID string `bson:"fromChainID"`
	ToChainID   string `bson:"toChainID"`
	SwapInfo    `bson:"swapinfo"`
	Status      SwapStatus `bson:"status"`
	Timestamp   int64      `bson:"timestamp"`
	Memo        string     `bson:"memo"`
}

// MgoSwapResult swap result (verified swap)
type MgoSwapResult struct {
	Key         string `bson:"_id"` // fromChainID + txid + logindex
	SwapType    uint32 `bson:"swaptype"`
	TxID        string `bson:"txid"`
	TxTo        string `bson:"txto"`
	TxHeight    uint64 `bson:"txheight"`
	TxTime      uint64 `bson:"txtime"`
	From        string `bson:"from"`
	To          string `bson:"to"`
	Bind        string `bson:"bind"`
	Value       string `bson:"value"`
	LogIndex    int    `bson:"logIndex"`
	FromChainID string `bson:"fromChainID"`
	ToChainID   string `bson:"toChainID"`
	SwapInfo    `bson:"swapinfo"`
	SwapTx      string     `bson:"swaptx"`
	OldSwapTxs  []string   `bson:"oldswaptxs"`
	SwapHeight  uint64     `bson:"swapheight"`
	SwapTime    uint64     `bson:"swaptime"`
	SwapValue   string     `bson:"swapvalue"`
	SwapNonce   uint64     `bson:"swapnonce"`
	Status      SwapStatus `bson:"status"`
	Timestamp   int64      `bson:"timestamp"`
	Memo        string     `bson:"memo"`
}

// SwapResultUpdateItems swap update items
type SwapResultUpdateItems struct {
	SwapTx     string
	OldSwapTxs []string
	SwapHeight uint64
	SwapTime   uint64
	SwapValue  string
	SwapNonce  uint64
	Status     SwapStatus
	Timestamp  int64
	Memo       string
}

// SwapInfo struct
type SwapInfo struct {
	*RouterSwapInfo  `bson:"routerSwapInfo,omitempty"  json:"routerSwapInfo,omitempty"`
	*AnyCallSwapInfo `bson:"anycallSwapInfo,omitempty" json:"anycallSwapInfo,omitempty"`
}

// RouterSwapInfo struct
type RouterSwapInfo struct {
	ForNative     bool     `bson:"forNative,omitempty"     json:"forNative,omitempty"`
	ForUnderlying bool     `bson:"forUnderlying,omitempty" json:"forUnderlying,omitempty"`
	Token         string   `bson:"token"                   json:"token"`
	TokenID       string   `bson:"tokenID"                 json:"tokenID"`
	Path          []string `bson:"path,omitempty"          json:"path,omitempty"`
	AmountOutMin  string   `bson:"amountOutMin,omitempty"  json:"amountOutMin,omitempty"`
}

// AnyCallSwapInfo struct
type AnyCallSwapInfo struct {
	CallFrom   string   `bson:"callFrom"   json:"callFrom"`
	CallTo     []string `bson:"callTo"     json:"callTo"`
	CallData   []string `bson:"callData"   json:"callData"`
	Callbacks  []string `bson:"callbacks"  json:"callbacks"`
	CallNonces []string `bson:"callNonces" json:"callNonces"`
}

// GetTokenID get tokenID
func (s *SwapInfo) GetTokenID() string {
	if s.RouterSwapInfo == nil {
		return ""
	}
	return s.TokenID
}
