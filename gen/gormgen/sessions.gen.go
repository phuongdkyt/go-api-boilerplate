// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gormgen

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/anhnmt/go-api-boilerplate/gen/pb"
	"github.com/anhnmt/go-api-boilerplate/internal/model"
)

func newSession(db *gorm.DB, opts ...gen.DOOption) session {
	_session := session{}

	_session.sessionDo.UseDB(db, opts...)
	_session.sessionDo.UseModel(&model.Session{})

	tableName := _session.sessionDo.TableName()
	_session.ALL = field.NewAsterisk(tableName)
	_session.ID = field.NewString(tableName, "id")
	_session.CreatedAt = field.NewTime(tableName, "created_at")
	_session.UpdatedAt = field.NewTime(tableName, "updated_at")
	_session.UserID = field.NewString(tableName, "user_id")
	_session.Fingerprint = field.NewString(tableName, "fingerprint")
	_session.UserAgent = field.NewString(tableName, "user_agent")
	_session.DeviceType = field.NewString(tableName, "device_type")
	_session.OS = field.NewString(tableName, "os")
	_session.Browser = field.NewString(tableName, "browser")
	_session.Device = field.NewString(tableName, "device")
	_session.IpAddress = field.NewString(tableName, "ip_address")
	_session.IsRevoked = field.NewBool(tableName, "is_revoked")
	_session.LastSeenAt = field.NewTime(tableName, "last_seen_at")
	_session.ExpiresAt = field.NewTime(tableName, "expires_at")

	_session.fillFieldMap()

	return _session
}

type session struct {
	sessionDo

	ALL         field.Asterisk
	ID          field.String
	CreatedAt   field.Time
	UpdatedAt   field.Time
	UserID      field.String
	Fingerprint field.String
	UserAgent   field.String
	DeviceType  field.String
	OS          field.String
	Browser     field.String
	Device      field.String
	IpAddress   field.String
	IsRevoked   field.Bool
	LastSeenAt  field.Time
	ExpiresAt   field.Time

	fieldMap map[string]field.Expr
}

func (s session) Table(newTableName string) *session {
	s.sessionDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s session) As(alias string) *session {
	s.sessionDo.DO = *(s.sessionDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *session) updateTableName(table string) *session {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewString(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.UserID = field.NewString(table, "user_id")
	s.Fingerprint = field.NewString(table, "fingerprint")
	s.UserAgent = field.NewString(table, "user_agent")
	s.DeviceType = field.NewString(table, "device_type")
	s.OS = field.NewString(table, "os")
	s.Browser = field.NewString(table, "browser")
	s.Device = field.NewString(table, "device")
	s.IpAddress = field.NewString(table, "ip_address")
	s.IsRevoked = field.NewBool(table, "is_revoked")
	s.LastSeenAt = field.NewTime(table, "last_seen_at")
	s.ExpiresAt = field.NewTime(table, "expires_at")

	s.fillFieldMap()

	return s
}

func (s *session) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *session) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 14)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["user_id"] = s.UserID
	s.fieldMap["fingerprint"] = s.Fingerprint
	s.fieldMap["user_agent"] = s.UserAgent
	s.fieldMap["device_type"] = s.DeviceType
	s.fieldMap["os"] = s.OS
	s.fieldMap["browser"] = s.Browser
	s.fieldMap["device"] = s.Device
	s.fieldMap["ip_address"] = s.IpAddress
	s.fieldMap["is_revoked"] = s.IsRevoked
	s.fieldMap["last_seen_at"] = s.LastSeenAt
	s.fieldMap["expires_at"] = s.ExpiresAt
}

func (s session) clone(db *gorm.DB) session {
	s.sessionDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s session) replaceDB(db *gorm.DB) session {
	s.sessionDo.ReplaceDB(db)
	return s
}

type sessionDo struct{ gen.DO }

type ISessionDo interface {
	gen.SubQuery
	Debug() ISessionDo
	WithContext(ctx context.Context) ISessionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISessionDo
	WriteDB() ISessionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISessionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISessionDo
	Not(conds ...gen.Condition) ISessionDo
	Or(conds ...gen.Condition) ISessionDo
	Select(conds ...field.Expr) ISessionDo
	Where(conds ...gen.Condition) ISessionDo
	Order(conds ...field.Expr) ISessionDo
	Distinct(cols ...field.Expr) ISessionDo
	Omit(cols ...field.Expr) ISessionDo
	Join(table schema.Tabler, on ...field.Expr) ISessionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISessionDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISessionDo
	Group(cols ...field.Expr) ISessionDo
	Having(conds ...gen.Condition) ISessionDo
	Limit(limit int) ISessionDo
	Offset(offset int) ISessionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISessionDo
	Unscoped() ISessionDo
	Create(values ...*model.Session) error
	CreateInBatches(values []*model.Session, batchSize int) error
	Save(values ...*model.Session) error
	First() (*model.Session, error)
	Take() (*model.Session, error)
	Last() (*model.Session, error)
	Find() ([]*model.Session, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Session, err error)
	FindInBatches(result *[]*model.Session, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Session) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISessionDo
	Assign(attrs ...field.AssignExpr) ISessionDo
	Joins(fields ...field.RelationField) ISessionDo
	Preload(fields ...field.RelationField) ISessionDo
	FirstOrInit() (*model.Session, error)
	FirstOrCreate() (*model.Session, error)
	FindByPage(offset int, limit int) (result []*model.Session, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISessionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FindByUserIDAndSessionID(userID string, sessionID string, limit int, offset int) (result []*pb.ActiveSessions, err error)
	CountByUserID(userID string) (result int, err error)
	UpdateRevokedByUserIDWithoutSessionID(userID string, sessionID string) (err error)
	FindByUserIDWithoutSessionID(userID string, sessionID string) (result []string, err error)
}

// select id,
// fingerprint,
// user_agent,
// os,
// device_type,
// device,
// ip_address,
// created_at as login_time,
// last_seen_at as last_seen
// {{if sessionID != ""}}
// , CASE
//
//	WHEN id = @sessionID THEN true
//	ELSE false
//
// END as is_current
// {{end}}
// from sessions
// where user_id = @userID
// and is_revoked = false
// and expires_at >= NOW() - INTERVAL '24 hours'
// order by
// {{if sessionID != ""}}
// is_current DESC,
// {{end}}
// last_seen_at DESC, updated_at DESC, expires_at DESC
// LIMIT {{if limit == 0}} 10 {{else}} @limit {{end}}
// OFFSET @offset;
func (s sessionDo) FindByUserIDAndSessionID(userID string, sessionID string, limit int, offset int) (result []*pb.ActiveSessions, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("select id, fingerprint, user_agent, os, device_type, device, ip_address, created_at as login_time, last_seen_at as last_seen ")
	if sessionID != "" {
		params = append(params, sessionID)
		generateSQL.WriteString(", CASE WHEN id = ? THEN true ELSE false END as is_current ")
	}
	params = append(params, userID)
	generateSQL.WriteString("from sessions where user_id = ? and is_revoked = false and expires_at >= NOW() - INTERVAL '24 hours' order by ")
	if sessionID != "" {
		generateSQL.WriteString("is_current DESC, ")
	}
	generateSQL.WriteString("last_seen_at DESC, updated_at DESC, expires_at DESC LIMIT ")
	if limit == 0 {
		generateSQL.WriteString("10 ")
	} else {
		params = append(params, limit)
		generateSQL.WriteString("? ")
	}
	params = append(params, offset)
	generateSQL.WriteString("OFFSET ?; ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// select count(1) AS total
// from sessions
// where user_id = @userID
// and is_revoked = false
// and expires_at >= NOW() - INTERVAL '24 hours'
func (s sessionDo) CountByUserID(userID string) (result int, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, userID)
	generateSQL.WriteString("select count(1) AS total from sessions where user_id = ? and is_revoked = false and expires_at >= NOW() - INTERVAL '24 hours' ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// update sessions
// set is_revoked = true
// where user_id = @userID
// {{if sessionID != ""}}
// and id <> @sessionID
// {{end}}
// and is_revoked = false
// and expires_at >= NOW() - INTERVAL '24 hours'
func (s sessionDo) UpdateRevokedByUserIDWithoutSessionID(userID string, sessionID string) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, userID)
	generateSQL.WriteString("update sessions set is_revoked = true where user_id = ? ")
	if sessionID != "" {
		params = append(params, sessionID)
		generateSQL.WriteString("and id <> ? ")
	}
	generateSQL.WriteString("and is_revoked = false and expires_at >= NOW() - INTERVAL '24 hours' ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// select id
// from sessions
// where user_id = @userID
// {{if sessionID != ""}}
// and id <> @sessionID
// {{end}}
// and is_revoked = false
// and expires_at >= NOW() - INTERVAL '24 hours'
func (s sessionDo) FindByUserIDWithoutSessionID(userID string, sessionID string) (result []string, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, userID)
	generateSQL.WriteString("select id from sessions where user_id = ? ")
	if sessionID != "" {
		params = append(params, sessionID)
		generateSQL.WriteString("and id <> ? ")
	}
	generateSQL.WriteString("and is_revoked = false and expires_at >= NOW() - INTERVAL '24 hours' ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (s sessionDo) Debug() ISessionDo {
	return s.withDO(s.DO.Debug())
}

func (s sessionDo) WithContext(ctx context.Context) ISessionDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sessionDo) ReadDB() ISessionDo {
	return s.Clauses(dbresolver.Read)
}

func (s sessionDo) WriteDB() ISessionDo {
	return s.Clauses(dbresolver.Write)
}

func (s sessionDo) Session(config *gorm.Session) ISessionDo {
	return s.withDO(s.DO.Session(config))
}

func (s sessionDo) Clauses(conds ...clause.Expression) ISessionDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sessionDo) Returning(value interface{}, columns ...string) ISessionDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sessionDo) Not(conds ...gen.Condition) ISessionDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sessionDo) Or(conds ...gen.Condition) ISessionDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sessionDo) Select(conds ...field.Expr) ISessionDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sessionDo) Where(conds ...gen.Condition) ISessionDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sessionDo) Order(conds ...field.Expr) ISessionDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sessionDo) Distinct(cols ...field.Expr) ISessionDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sessionDo) Omit(cols ...field.Expr) ISessionDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sessionDo) Join(table schema.Tabler, on ...field.Expr) ISessionDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sessionDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISessionDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sessionDo) RightJoin(table schema.Tabler, on ...field.Expr) ISessionDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sessionDo) Group(cols ...field.Expr) ISessionDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sessionDo) Having(conds ...gen.Condition) ISessionDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sessionDo) Limit(limit int) ISessionDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sessionDo) Offset(offset int) ISessionDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sessionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISessionDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sessionDo) Unscoped() ISessionDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sessionDo) Create(values ...*model.Session) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sessionDo) CreateInBatches(values []*model.Session, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sessionDo) Save(values ...*model.Session) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sessionDo) First() (*model.Session, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Session), nil
	}
}

func (s sessionDo) Take() (*model.Session, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Session), nil
	}
}

func (s sessionDo) Last() (*model.Session, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Session), nil
	}
}

func (s sessionDo) Find() ([]*model.Session, error) {
	result, err := s.DO.Find()
	return result.([]*model.Session), err
}

func (s sessionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Session, err error) {
	buf := make([]*model.Session, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sessionDo) FindInBatches(result *[]*model.Session, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sessionDo) Attrs(attrs ...field.AssignExpr) ISessionDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sessionDo) Assign(attrs ...field.AssignExpr) ISessionDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sessionDo) Joins(fields ...field.RelationField) ISessionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sessionDo) Preload(fields ...field.RelationField) ISessionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sessionDo) FirstOrInit() (*model.Session, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Session), nil
	}
}

func (s sessionDo) FirstOrCreate() (*model.Session, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Session), nil
	}
}

func (s sessionDo) FindByPage(offset int, limit int) (result []*model.Session, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sessionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sessionDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sessionDo) Delete(models ...*model.Session) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sessionDo) withDO(do gen.Dao) *sessionDo {
	s.DO = *do.(*gen.DO)
	return s
}
