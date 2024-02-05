// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/common/dao/internal"
)

// casbinRuleDao is the data access object for table casbin_rule.
// You can define custom methods on it to extend its functionality as you wish.
type casbinRuleDao struct {
	*internal.CasbinRuleDao
}

var (
	// CasbinRule is globally public accessible object for table casbin_rule operations.
	CasbinRule = casbinRuleDao{
		internal.NewCasbinRuleDao(),
	}
)

// Fill with you ideas below.
