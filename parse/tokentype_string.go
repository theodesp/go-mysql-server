// Code generated by "stringer -type=TokenType"; DO NOT EDIT

package parse

import "fmt"

const _TokenType_name = "ErrorTokenEOFTokenLeftParenTokenRightParenTokenCommaTokenDotTokenKeywordTokenIdentifierTokenIntTokenFloatTokenStringTokenOpToken"

var _TokenType_index = [...]uint8{0, 10, 18, 32, 47, 57, 65, 77, 92, 100, 110, 121, 128}

func (i TokenType) String() string {
	if i >= TokenType(len(_TokenType_index)-1) {
		return fmt.Sprintf("TokenType(%d)", i)
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
