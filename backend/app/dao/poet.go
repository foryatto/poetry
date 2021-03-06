// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"backend/app/dao/internal"
)

// poetDao is the manager for logic model data accessing and custom defined data operations functions management. 
// You can define custom methods on it to extend its functionality as you wish.
type poetDao struct {
	*internal.PoetDao
}

var (
	// Poet is globally public accessible object for table poet operations.
	Poet = poetDao{
		internal.NewPoetDao(),
	}
)

// Fill with you ideas below.