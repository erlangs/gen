package model

import (
	"database/sql"
	//"time"

	//"github.com/satori/go.uuid"

	"gorm.io/gorm"
)

/*
DB Table Details
-------------------------------------


Table: realm
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] access_code_lifespan                           INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] user_action_lifespan                           INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 3] access_token_lifespan                          INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] account_theme                                  VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] admin_theme                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] email_theme                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] enabled                                        BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 8] events_enabled                                 BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 9] events_expiration                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[10] login_theme                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[11] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[12] not_before                                     INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[13] password_policy                                VARCHAR(2550)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 2550    default: []
[14] registration_allowed                           BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[15] remember_me                                    BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[16] reset_password_allowed                         BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[17] social                                         BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[18] ssl_required                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[19] sso_idle_timeout                               INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[20] sso_max_lifespan                               INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[21] update_profile_on_soc_login                    BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[22] verify_email                                   BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[23] master_admin_client                            VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[24] login_lifespan                                 INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[25] internationalization_enabled                   BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[26] default_locale                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[27] reg_email_as_username                          BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[28] admin_events_enabled                           BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[29] admin_events_details_enabled                   BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[30] edit_username_allowed                          BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[31] otp_policy_counter                             INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[32] otp_policy_window                              INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [1]
[33] otp_policy_period                              INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [30]
[34] otp_policy_digits                              INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [6]
[35] otp_policy_alg                                 VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: [HmacSHA1]
[36] otp_policy_type                                VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: [totp]
[37] browser_flow                                   VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[38] registration_flow                              VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[39] direct_grant_flow                              VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[40] reset_credentials_flow                         VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[41] client_auth_flow                               VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[42] offline_session_idle_timeout                   INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[43] revoke_refresh_token                           BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[44] access_token_life_implicit                     INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[45] login_with_email_allowed                       BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [true]
[46] duplicate_emails_allowed                       BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[47] docker_auth_flow                               VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[48] refresh_token_max_reuse                        INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[49] allow_user_managed_access                      BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[50] sso_max_lifespan_remember_me                   INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[51] sso_idle_timeout_remember_me                   INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[52] default_role                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "iwpedMQidDrpVhGclIdeIfisT",    "access_code_lifespan": 39,    "user_action_lifespan": 77,    "access_token_lifespan": 41,    "account_theme": "ShKkAqkWELrmcihJhWJNksVTl",    "admin_theme": "MUQXUMXdOXWlVrTSakHZxdcTo",    "email_theme": "RUonIshPcyaEduPMZjhRCFGYV",    "enabled": false,    "events_enabled": false,    "events_expiration": 24,    "login_theme": "QbwoJaLyFvLAIFWdblPpiwYnS",    "name": "jMyeuDMKMRJSNFNnpqktFwclI",    "not_before": 17,    "password_policy": "XPkWgNDGZljXhOkgIMcWBhMdh",    "registration_allowed": true,    "remember_me": true,    "reset_password_allowed": false,    "social": true,    "ssl_required": "ZQiFJNgVDxoetlMcKMqbEEpRj",    "sso_idle_timeout": 68,    "sso_max_lifespan": 74,    "update_profile_on_soc_login": false,    "verify_email": true,    "master_admin_client": "LQVmMHdOQbKBxIXRaMpbKOlJg",    "login_lifespan": 37,    "internationalization_enabled": true,    "default_locale": "vXtPRwiNeRIboXYRdsMElaqoL",    "reg_email_as_username": false,    "admin_events_enabled": false,    "admin_events_details_enabled": true,    "edit_username_allowed": true,    "otp_policy_counter": 37,    "otp_policy_window": 8,    "otp_policy_period": 28,    "otp_policy_digits": 29,    "otp_policy_alg": "NMTQEVYWucHqTaUWWnqEPhDAC",    "otp_policy_type": "BwStaXPiXfAeyDtMXIIeiXHsu",    "browser_flow": "bCwJfxUnZSMoEbbKXlerlSjaQ",    "registration_flow": "yuVTDmQpJqDFvGCnyQDTafDts",    "direct_grant_flow": "TDrDrhAyYJygPGIwnNQDwGHqE",    "reset_credentials_flow": "sNUcuxwUGnhaEgnOWGdjTIBTH",    "client_auth_flow": "yGNUyKfpiFXPwiolydEEYnawa",    "offline_session_idle_timeout": 83,    "revoke_refresh_token": true,    "access_token_life_implicit": 83,    "login_with_email_allowed": false,    "duplicate_emails_allowed": true,    "docker_auth_flow": "KRAXCEwDOwqwxGuLtCqvlRXKW",    "refresh_token_max_reuse": 30,    "allow_user_managed_access": true,    "sso_max_lifespan_remember_me": 36,    "sso_idle_timeout_remember_me": 25,    "default_role": "lhEjhlOFWUllcNExqreJJeDaS"}



*/

// Realm struct is a row record of the realm table in the keycloak database
type Realm struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR;size:36;" json:"id"`
	//[ 1] access_code_lifespan                           INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	AccessCodeLifespan sql.NullInt32 `gorm:"column:access_code_lifespan;type:INT4;" json:"access_code_lifespan"`
	//[ 2] user_action_lifespan                           INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	UserActionLifespan sql.NullInt32 `gorm:"column:user_action_lifespan;type:INT4;" json:"user_action_lifespan"`
	//[ 3] access_token_lifespan                          INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	AccessTokenLifespan sql.NullInt32 `gorm:"column:access_token_lifespan;type:INT4;" json:"access_token_lifespan"`
	//[ 4] account_theme                                  VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	AccountTheme sql.NullString `gorm:"column:account_theme;type:VARCHAR;size:255;" json:"account_theme"`
	//[ 5] admin_theme                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	AdminTheme sql.NullString `gorm:"column:admin_theme;type:VARCHAR;size:255;" json:"admin_theme"`
	//[ 6] email_theme                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	EmailTheme sql.NullString `gorm:"column:email_theme;type:VARCHAR;size:255;" json:"email_theme"`
	//[ 7] enabled                                        BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	Enabled bool `gorm:"column:enabled;type:BOOL;default:false;" json:"enabled"`
	//[ 8] events_enabled                                 BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	EventsEnabled bool `gorm:"column:events_enabled;type:BOOL;default:false;" json:"events_enabled"`
	//[ 9] events_expiration                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	EventsExpiration sql.NullInt64 `gorm:"column:events_expiration;type:INT8;" json:"events_expiration"`
	//[10] login_theme                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	LoginTheme sql.NullString `gorm:"column:login_theme;type:VARCHAR;size:255;" json:"login_theme"`
	//[11] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name sql.NullString `gorm:"column:name;type:VARCHAR;size:255;" json:"name"`
	//[12] not_before                                     INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	NotBefore sql.NullInt32 `gorm:"column:not_before;type:INT4;" json:"not_before"`
	//[13] password_policy                                VARCHAR(2550)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 2550    default: []
	PasswordPolicy sql.NullString `gorm:"column:password_policy;type:VARCHAR;size:2550;" json:"password_policy"`
	//[14] registration_allowed                           BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	RegistrationAllowed bool `gorm:"column:registration_allowed;type:BOOL;default:false;" json:"registration_allowed"`
	//[15] remember_me                                    BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	RememberMe bool `gorm:"column:remember_me;type:BOOL;default:false;" json:"remember_me"`
	//[16] reset_password_allowed                         BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	ResetPasswordAllowed bool `gorm:"column:reset_password_allowed;type:BOOL;default:false;" json:"reset_password_allowed"`
	//[17] social                                         BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	Social bool `gorm:"column:social;type:BOOL;default:false;" json:"social"`
	//[18] ssl_required                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	SslRequired sql.NullString `gorm:"column:ssl_required;type:VARCHAR;size:255;" json:"ssl_required"`
	//[19] sso_idle_timeout                               INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	SsoIdleTimeout sql.NullInt32 `gorm:"column:sso_idle_timeout;type:INT4;" json:"sso_idle_timeout"`
	//[20] sso_max_lifespan                               INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	SsoMaxLifespan sql.NullInt32 `gorm:"column:sso_max_lifespan;type:INT4;" json:"sso_max_lifespan"`
	//[21] update_profile_on_soc_login                    BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	UpdateProfileOnSocLogin bool `gorm:"column:update_profile_on_soc_login;type:BOOL;default:false;" json:"update_profile_on_soc_login"`
	//[22] verify_email                                   BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	VerifyEmail bool `gorm:"column:verify_email;type:BOOL;default:false;" json:"verify_email"`
	//[23] master_admin_client                            VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	MasterAdminClient sql.NullString `gorm:"column:master_admin_client;type:VARCHAR;size:36;" json:"master_admin_client"`
	//[24] login_lifespan                                 INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	LoginLifespan sql.NullInt32 `gorm:"column:login_lifespan;type:INT4;" json:"login_lifespan"`
	//[25] internationalization_enabled                   BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	InternationalizationEnabled bool `gorm:"column:internationalization_enabled;type:BOOL;default:false;" json:"internationalization_enabled"`
	//[26] default_locale                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	DefaultLocale sql.NullString `gorm:"column:default_locale;type:VARCHAR;size:255;" json:"default_locale"`
	//[27] reg_email_as_username                          BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	RegEmailAsUsername bool `gorm:"column:reg_email_as_username;type:BOOL;default:false;" json:"reg_email_as_username"`
	//[28] admin_events_enabled                           BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	AdminEventsEnabled bool `gorm:"column:admin_events_enabled;type:BOOL;default:false;" json:"admin_events_enabled"`
	//[29] admin_events_details_enabled                   BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	AdminEventsDetailsEnabled bool `gorm:"column:admin_events_details_enabled;type:BOOL;default:false;" json:"admin_events_details_enabled"`
	//[30] edit_username_allowed                          BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	EditUsernameAllowed bool `gorm:"column:edit_username_allowed;type:BOOL;default:false;" json:"edit_username_allowed"`
	//[31] otp_policy_counter                             INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
	OtpPolicyCounter sql.NullInt32 `gorm:"column:otp_policy_counter;type:INT4;default:0;" json:"otp_policy_counter"`
	//[32] otp_policy_window                              INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [1]
	OtpPolicyWindow sql.NullInt32 `gorm:"column:otp_policy_window;type:INT4;default:1;" json:"otp_policy_window"`
	//[33] otp_policy_period                              INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [30]
	OtpPolicyPeriod sql.NullInt32 `gorm:"column:otp_policy_period;type:INT4;default:30;" json:"otp_policy_period"`
	//[34] otp_policy_digits                              INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [6]
	OtpPolicyDigits sql.NullInt32 `gorm:"column:otp_policy_digits;type:INT4;default:6;" json:"otp_policy_digits"`
	//[35] otp_policy_alg                                 VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: [HmacSHA1]
	OtpPolicyAlg sql.NullString `gorm:"column:otp_policy_alg;type:VARCHAR;size:36;default:HmacSHA1;" json:"otp_policy_alg"`
	//[36] otp_policy_type                                VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: [totp]
	OtpPolicyType sql.NullString `gorm:"column:otp_policy_type;type:VARCHAR;size:36;default:totp;" json:"otp_policy_type"`
	//[37] browser_flow                                   VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	BrowserFlow sql.NullString `gorm:"column:browser_flow;type:VARCHAR;size:36;" json:"browser_flow"`
	//[38] registration_flow                              VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RegistrationFlow sql.NullString `gorm:"column:registration_flow;type:VARCHAR;size:36;" json:"registration_flow"`
	//[39] direct_grant_flow                              VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	DirectGrantFlow sql.NullString `gorm:"column:direct_grant_flow;type:VARCHAR;size:36;" json:"direct_grant_flow"`
	//[40] reset_credentials_flow                         VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ResetCredentialsFlow sql.NullString `gorm:"column:reset_credentials_flow;type:VARCHAR;size:36;" json:"reset_credentials_flow"`
	//[41] client_auth_flow                               VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ClientAuthFlow sql.NullString `gorm:"column:client_auth_flow;type:VARCHAR;size:36;" json:"client_auth_flow"`
	//[42] offline_session_idle_timeout                   INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
	OfflineSessionIdleTimeout sql.NullInt32 `gorm:"column:offline_session_idle_timeout;type:INT4;default:0;" json:"offline_session_idle_timeout"`
	//[43] revoke_refresh_token                           BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	RevokeRefreshToken bool `gorm:"column:revoke_refresh_token;type:BOOL;default:false;" json:"revoke_refresh_token"`
	//[44] access_token_life_implicit                     INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
	AccessTokenLifeImplicit sql.NullInt32 `gorm:"column:access_token_life_implicit;type:INT4;default:0;" json:"access_token_life_implicit"`
	//[45] login_with_email_allowed                       BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [true]
	LoginWithEmailAllowed bool `gorm:"column:login_with_email_allowed;type:BOOL;default:true;" json:"login_with_email_allowed"`
	//[46] duplicate_emails_allowed                       BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	DuplicateEmailsAllowed bool `gorm:"column:duplicate_emails_allowed;type:BOOL;default:false;" json:"duplicate_emails_allowed"`
	//[47] docker_auth_flow                               VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	DockerAuthFlow sql.NullString `gorm:"column:docker_auth_flow;type:VARCHAR;size:36;" json:"docker_auth_flow"`
	//[48] refresh_token_max_reuse                        INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
	RefreshTokenMaxReuse sql.NullInt32 `gorm:"column:refresh_token_max_reuse;type:INT4;default:0;" json:"refresh_token_max_reuse"`
	//[49] allow_user_managed_access                      BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	AllowUserManagedAccess bool `gorm:"column:allow_user_managed_access;type:BOOL;default:false;" json:"allow_user_managed_access"`
	//[50] sso_max_lifespan_remember_me                   INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
	SsoMaxLifespanRememberMe int32 `gorm:"column:sso_max_lifespan_remember_me;type:INT4;default:0;" json:"sso_max_lifespan_remember_me"`
	//[51] sso_idle_timeout_remember_me                   INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
	SsoIdleTimeoutRememberMe int32 `gorm:"column:sso_idle_timeout_remember_me;type:INT4;default:0;" json:"sso_idle_timeout_remember_me"`
	//[52] default_role                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	DefaultRole sql.NullString `gorm:"column:default_role;type:VARCHAR;size:255;" json:"default_role"`
}

var realmTableInfo = &TableInfo{
	Name: "realm",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "access_code_lifespan",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "AccessCodeLifespan",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "access_code_lifespan",
			ProtobufFieldName:  "access_code_lifespan",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "user_action_lifespan",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "UserActionLifespan",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "user_action_lifespan",
			ProtobufFieldName:  "user_action_lifespan",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "access_token_lifespan",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "AccessTokenLifespan",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "access_token_lifespan",
			ProtobufFieldName:  "access_token_lifespan",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "account_theme",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "AccountTheme",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "account_theme",
			ProtobufFieldName:  "account_theme",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "admin_theme",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "AdminTheme",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "admin_theme",
			ProtobufFieldName:  "admin_theme",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "email_theme",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "EmailTheme",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "email_theme",
			ProtobufFieldName:  "email_theme",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "enabled",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "Enabled",
			GoFieldType:        "bool",
			JSONFieldName:      "enabled",
			ProtobufFieldName:  "enabled",
			ProtobufType:       "bool",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "events_enabled",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "EventsEnabled",
			GoFieldType:        "bool",
			JSONFieldName:      "events_enabled",
			ProtobufFieldName:  "events_enabled",
			ProtobufType:       "bool",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "events_expiration",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "EventsExpiration",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "events_expiration",
			ProtobufFieldName:  "events_expiration",
			ProtobufType:       "int64",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "login_theme",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "LoginTheme",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "login_theme",
			ProtobufFieldName:  "login_theme",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "not_before",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "NotBefore",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "not_before",
			ProtobufFieldName:  "not_before",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "password_policy",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(2550)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       2550,
			GoFieldName:        "PasswordPolicy",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "password_policy",
			ProtobufFieldName:  "password_policy",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "registration_allowed",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "RegistrationAllowed",
			GoFieldType:        "bool",
			JSONFieldName:      "registration_allowed",
			ProtobufFieldName:  "registration_allowed",
			ProtobufType:       "bool",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "remember_me",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "RememberMe",
			GoFieldType:        "bool",
			JSONFieldName:      "remember_me",
			ProtobufFieldName:  "remember_me",
			ProtobufType:       "bool",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "reset_password_allowed",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "ResetPasswordAllowed",
			GoFieldType:        "bool",
			JSONFieldName:      "reset_password_allowed",
			ProtobufFieldName:  "reset_password_allowed",
			ProtobufType:       "bool",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "social",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "Social",
			GoFieldType:        "bool",
			JSONFieldName:      "social",
			ProtobufFieldName:  "social",
			ProtobufType:       "bool",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "ssl_required",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "SslRequired",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "ssl_required",
			ProtobufFieldName:  "ssl_required",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "sso_idle_timeout",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "SsoIdleTimeout",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "sso_idle_timeout",
			ProtobufFieldName:  "sso_idle_timeout",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "sso_max_lifespan",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "SsoMaxLifespan",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "sso_max_lifespan",
			ProtobufFieldName:  "sso_max_lifespan",
			ProtobufType:       "int32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "update_profile_on_soc_login",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "UpdateProfileOnSocLogin",
			GoFieldType:        "bool",
			JSONFieldName:      "update_profile_on_soc_login",
			ProtobufFieldName:  "update_profile_on_soc_login",
			ProtobufType:       "bool",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "verify_email",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "VerifyEmail",
			GoFieldType:        "bool",
			JSONFieldName:      "verify_email",
			ProtobufFieldName:  "verify_email",
			ProtobufType:       "bool",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "master_admin_client",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "MasterAdminClient",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "master_admin_client",
			ProtobufFieldName:  "master_admin_client",
			ProtobufType:       "string",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "login_lifespan",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "LoginLifespan",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "login_lifespan",
			ProtobufFieldName:  "login_lifespan",
			ProtobufType:       "int32",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "internationalization_enabled",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "InternationalizationEnabled",
			GoFieldType:        "bool",
			JSONFieldName:      "internationalization_enabled",
			ProtobufFieldName:  "internationalization_enabled",
			ProtobufType:       "bool",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "default_locale",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "DefaultLocale",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "default_locale",
			ProtobufFieldName:  "default_locale",
			ProtobufType:       "string",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "reg_email_as_username",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "RegEmailAsUsername",
			GoFieldType:        "bool",
			JSONFieldName:      "reg_email_as_username",
			ProtobufFieldName:  "reg_email_as_username",
			ProtobufType:       "bool",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "admin_events_enabled",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "AdminEventsEnabled",
			GoFieldType:        "bool",
			JSONFieldName:      "admin_events_enabled",
			ProtobufFieldName:  "admin_events_enabled",
			ProtobufType:       "bool",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "admin_events_details_enabled",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "AdminEventsDetailsEnabled",
			GoFieldType:        "bool",
			JSONFieldName:      "admin_events_details_enabled",
			ProtobufFieldName:  "admin_events_details_enabled",
			ProtobufType:       "bool",
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "edit_username_allowed",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "EditUsernameAllowed",
			GoFieldType:        "bool",
			JSONFieldName:      "edit_username_allowed",
			ProtobufFieldName:  "edit_username_allowed",
			ProtobufType:       "bool",
			ProtobufPos:        31,
		},

		&ColumnInfo{
			Index:              31,
			Name:               "otp_policy_counter",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "OtpPolicyCounter",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "otp_policy_counter",
			ProtobufFieldName:  "otp_policy_counter",
			ProtobufType:       "int32",
			ProtobufPos:        32,
		},

		&ColumnInfo{
			Index:              32,
			Name:               "otp_policy_window",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "OtpPolicyWindow",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "otp_policy_window",
			ProtobufFieldName:  "otp_policy_window",
			ProtobufType:       "int32",
			ProtobufPos:        33,
		},

		&ColumnInfo{
			Index:              33,
			Name:               "otp_policy_period",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "OtpPolicyPeriod",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "otp_policy_period",
			ProtobufFieldName:  "otp_policy_period",
			ProtobufType:       "int32",
			ProtobufPos:        34,
		},

		&ColumnInfo{
			Index:              34,
			Name:               "otp_policy_digits",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "OtpPolicyDigits",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "otp_policy_digits",
			ProtobufFieldName:  "otp_policy_digits",
			ProtobufType:       "int32",
			ProtobufPos:        35,
		},

		&ColumnInfo{
			Index:              35,
			Name:               "otp_policy_alg",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "OtpPolicyAlg",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "otp_policy_alg",
			ProtobufFieldName:  "otp_policy_alg",
			ProtobufType:       "string",
			ProtobufPos:        36,
		},

		&ColumnInfo{
			Index:              36,
			Name:               "otp_policy_type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "OtpPolicyType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "otp_policy_type",
			ProtobufFieldName:  "otp_policy_type",
			ProtobufType:       "string",
			ProtobufPos:        37,
		},

		&ColumnInfo{
			Index:              37,
			Name:               "browser_flow",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "BrowserFlow",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "browser_flow",
			ProtobufFieldName:  "browser_flow",
			ProtobufType:       "string",
			ProtobufPos:        38,
		},

		&ColumnInfo{
			Index:              38,
			Name:               "registration_flow",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "RegistrationFlow",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "registration_flow",
			ProtobufFieldName:  "registration_flow",
			ProtobufType:       "string",
			ProtobufPos:        39,
		},

		&ColumnInfo{
			Index:              39,
			Name:               "direct_grant_flow",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "DirectGrantFlow",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "direct_grant_flow",
			ProtobufFieldName:  "direct_grant_flow",
			ProtobufType:       "string",
			ProtobufPos:        40,
		},

		&ColumnInfo{
			Index:              40,
			Name:               "reset_credentials_flow",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "ResetCredentialsFlow",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "reset_credentials_flow",
			ProtobufFieldName:  "reset_credentials_flow",
			ProtobufType:       "string",
			ProtobufPos:        41,
		},

		&ColumnInfo{
			Index:              41,
			Name:               "client_auth_flow",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "ClientAuthFlow",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "client_auth_flow",
			ProtobufFieldName:  "client_auth_flow",
			ProtobufType:       "string",
			ProtobufPos:        42,
		},

		&ColumnInfo{
			Index:              42,
			Name:               "offline_session_idle_timeout",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "OfflineSessionIdleTimeout",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "offline_session_idle_timeout",
			ProtobufFieldName:  "offline_session_idle_timeout",
			ProtobufType:       "int32",
			ProtobufPos:        43,
		},

		&ColumnInfo{
			Index:              43,
			Name:               "revoke_refresh_token",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "RevokeRefreshToken",
			GoFieldType:        "bool",
			JSONFieldName:      "revoke_refresh_token",
			ProtobufFieldName:  "revoke_refresh_token",
			ProtobufType:       "bool",
			ProtobufPos:        44,
		},

		&ColumnInfo{
			Index:              44,
			Name:               "access_token_life_implicit",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "AccessTokenLifeImplicit",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "access_token_life_implicit",
			ProtobufFieldName:  "access_token_life_implicit",
			ProtobufType:       "int32",
			ProtobufPos:        45,
		},

		&ColumnInfo{
			Index:              45,
			Name:               "login_with_email_allowed",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "LoginWithEmailAllowed",
			GoFieldType:        "bool",
			JSONFieldName:      "login_with_email_allowed",
			ProtobufFieldName:  "login_with_email_allowed",
			ProtobufType:       "bool",
			ProtobufPos:        46,
		},

		&ColumnInfo{
			Index:              46,
			Name:               "duplicate_emails_allowed",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "DuplicateEmailsAllowed",
			GoFieldType:        "bool",
			JSONFieldName:      "duplicate_emails_allowed",
			ProtobufFieldName:  "duplicate_emails_allowed",
			ProtobufType:       "bool",
			ProtobufPos:        47,
		},

		&ColumnInfo{
			Index:              47,
			Name:               "docker_auth_flow",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "DockerAuthFlow",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "docker_auth_flow",
			ProtobufFieldName:  "docker_auth_flow",
			ProtobufType:       "string",
			ProtobufPos:        48,
		},

		&ColumnInfo{
			Index:              48,
			Name:               "refresh_token_max_reuse",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "RefreshTokenMaxReuse",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "refresh_token_max_reuse",
			ProtobufFieldName:  "refresh_token_max_reuse",
			ProtobufType:       "int32",
			ProtobufPos:        49,
		},

		&ColumnInfo{
			Index:              49,
			Name:               "allow_user_managed_access",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "AllowUserManagedAccess",
			GoFieldType:        "bool",
			JSONFieldName:      "allow_user_managed_access",
			ProtobufFieldName:  "allow_user_managed_access",
			ProtobufType:       "bool",
			ProtobufPos:        50,
		},

		&ColumnInfo{
			Index:              50,
			Name:               "sso_max_lifespan_remember_me",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "SsoMaxLifespanRememberMe",
			GoFieldType:        "int32",
			JSONFieldName:      "sso_max_lifespan_remember_me",
			ProtobufFieldName:  "sso_max_lifespan_remember_me",
			ProtobufType:       "int32",
			ProtobufPos:        51,
		},

		&ColumnInfo{
			Index:              51,
			Name:               "sso_idle_timeout_remember_me",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "SsoIdleTimeoutRememberMe",
			GoFieldType:        "int32",
			JSONFieldName:      "sso_idle_timeout_remember_me",
			ProtobufFieldName:  "sso_idle_timeout_remember_me",
			ProtobufType:       "int32",
			ProtobufPos:        52,
		},

		&ColumnInfo{
			Index:              52,
			Name:               "default_role",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "DefaultRole",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "default_role",
			ProtobufFieldName:  "default_role",
			ProtobufType:       "string",
			ProtobufPos:        53,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *Realm) TableName() string {
	return "realm"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *Realm) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *Realm) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *Realm) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *Realm) TableInfo() *TableInfo {
	return realmTableInfo
}
