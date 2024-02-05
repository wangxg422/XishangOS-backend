// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)

// sysAuthRuleDao is the data access object for table sys_auth_rule.
// You can define custom methods on it to extend its functionality as you wish.
type sysAuthRuleDao struct {
	*internal.SysAuthRuleDao
}

var (
	// SysAuthRule is globally public accessible object for table sys_auth_rule operations.
	SysAuthRule = sysAuthRuleDao{
		internal.NewSysAuthRuleDao(),
	}
)

// Fill with you ideas below.
