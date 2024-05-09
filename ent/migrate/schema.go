// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ActiveCodeInfoColumns holds the columns for the "active_code_info" table.
	ActiveCodeInfoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "Status 1: normal 2: ban | 状态 1 正常 2 禁用", Default: 1},
		{Name: "active_key", Type: field.TypeString, Unique: true, Comment: "激活码", SchemaType: map[string]string{"mysql": "char(16)"}},
		{Name: "user_id", Type: field.TypeString, Comment: "用户id", SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "app_id", Type: field.TypeString, Comment: "appid", SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "active_ip", Type: field.TypeString, Comment: "激活的ip", SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "device_sn", Type: field.TypeString, Comment: "设备的sn码", SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "device_mac", Type: field.TypeString, Comment: "设备的mac", SchemaType: map[string]string{"mysql": "varchar(256)"}},
		{Name: "device_identity", Type: field.TypeString, Comment: "设备的身份", SchemaType: map[string]string{"mysql": "varchar(256)"}},
		{Name: "active_date", Type: field.TypeTime, Comment: "激活时间"},
		{Name: "active_type", Type: field.TypeInt64, Comment: "激活类型", Default: 0},
		{Name: "active_file", Type: field.TypeString, Comment: "激活的文件", SchemaType: map[string]string{"mysql": "varchar(256)"}},
		{Name: "version", Type: field.TypeString, Comment: "版本", SchemaType: map[string]string{"mysql": "varchar(256)"}},
		{Name: "start_date", Type: field.TypeTime, Comment: "开始时间"},
		{Name: "expire_date", Type: field.TypeTime, Comment: "结束时间"},
		{Name: "app_sdk_id", Type: field.TypeUUID, Nullable: true, Comment: "关联app_key"},
	}
	// ActiveCodeInfoTable holds the schema information for the "active_code_info" table.
	ActiveCodeInfoTable = &schema.Table{
		Name:       "active_code_info",
		Columns:    ActiveCodeInfoColumns,
		PrimaryKey: []*schema.Column{ActiveCodeInfoColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "active_code_info_app_sdk_active_code",
				Columns:    []*schema.Column{ActiveCodeInfoColumns[17]},
				RefColumns: []*schema.Column{AppSdkColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "activecodeinfo_active_key",
				Unique:  true,
				Columns: []*schema.Column{ActiveCodeInfoColumns[4]},
			},
		},
	}
	// AppInfoColumns holds the columns for the "app_info" table.
	AppInfoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "user_id", Type: field.TypeString, Comment: "用户id", SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "app_name", Type: field.TypeString, Comment: "app名字", SchemaType: map[string]string{"mysql": "varchar(256)"}},
		{Name: "summary", Type: field.TypeString, Comment: "摘要", SchemaType: map[string]string{"mysql": "varchar(1024)"}},
		{Name: "app_category", Type: field.TypeInt64, Comment: "种类"},
		{Name: "use_industry", Type: field.TypeInt64, Comment: "使用领域"},
		{Name: "app_category_name", Type: field.TypeString, Comment: "种类名字", SchemaType: map[string]string{"mysql": "varchar(256)"}},
		{Name: "use_industry_name", Type: field.TypeString, Comment: "使用领域名字", SchemaType: map[string]string{"mysql": "varchar(256)"}},
		{Name: "app_key", Type: field.TypeString, Comment: "钥匙", SchemaType: map[string]string{"mysql": "varchar(256)"}},
		{Name: "app_secret", Type: field.TypeString, Comment: "秘密", SchemaType: map[string]string{"mysql": "varchar(256)"}},
	}
	// AppInfoTable holds the schema information for the "app_info" table.
	AppInfoTable = &schema.Table{
		Name:       "app_info",
		Columns:    AppInfoColumns,
		PrimaryKey: []*schema.Column{AppInfoColumns[0]},
	}
	// AppSdkColumns holds the columns for the "app_sdk" table.
	AppSdkColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "sdk_key", Type: field.TypeString, Comment: "分配给app的sdk", SchemaType: map[string]string{"mysql": "char(32)"}},
		{Name: "app", Type: field.TypeUUID, Nullable: true, Comment: "app的id"},
		{Name: "sdk", Type: field.TypeUUID, Nullable: true, Comment: "sdk的id"},
	}
	// AppSdkTable holds the schema information for the "app_sdk" table.
	AppSdkTable = &schema.Table{
		Name:       "app_sdk",
		Columns:    AppSdkColumns,
		PrimaryKey: []*schema.Column{AppSdkColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_sdk_app_info_app_sdk",
				Columns:    []*schema.Column{AppSdkColumns[4]},
				RefColumns: []*schema.Column{AppInfoColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "app_sdk_sdk_info_app_sdk",
				Columns:    []*schema.Column{AppSdkColumns[5]},
				RefColumns: []*schema.Column{SdkInfoColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CategoryInfoColumns holds the columns for the "category_info" table.
	CategoryInfoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "name", Type: field.TypeString, Comment: "名字", SchemaType: map[string]string{"mysql": "varchar(64)"}},
	}
	// CategoryInfoTable holds the schema information for the "category_info" table.
	CategoryInfoTable = &schema.Table{
		Name:       "category_info",
		Columns:    CategoryInfoColumns,
		PrimaryKey: []*schema.Column{CategoryInfoColumns[0]},
	}
	// IndustryInfoColumns holds the columns for the "industry_info" table.
	IndustryInfoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "name", Type: field.TypeString, Comment: "名字", SchemaType: map[string]string{"mysql": "varchar(64)"}},
	}
	// IndustryInfoTable holds the schema information for the "industry_info" table.
	IndustryInfoTable = &schema.Table{
		Name:       "industry_info",
		Columns:    IndustryInfoColumns,
		PrimaryKey: []*schema.Column{IndustryInfoColumns[0]},
	}
	// SdkInfoColumns holds the columns for the "sdk_info" table.
	SdkInfoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "name", Type: field.TypeString, Unique: true, Comment: "名字", SchemaType: map[string]string{"mysql": "varchar(256)"}},
		{Name: "avatar", Type: field.TypeString, Comment: "头像", SchemaType: map[string]string{"mysql": "varchar(256)"}},
		{Name: "desc", Type: field.TypeInt64, Comment: "排序"},
		{Name: "download_url", Type: field.TypeString, Comment: "下载地址", SchemaType: map[string]string{"mysql": "varchar(512)"}},
	}
	// SdkInfoTable holds the schema information for the "sdk_info" table.
	SdkInfoTable = &schema.Table{
		Name:       "sdk_info",
		Columns:    SdkInfoColumns,
		PrimaryKey: []*schema.Column{SdkInfoColumns[0]},
	}
	// SdkUsageColumns holds the columns for the "sdk_usage" table.
	SdkUsageColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "user_id", Type: field.TypeString, Unique: true, Comment: "用户id", SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "all", Type: field.TypeInt64, Comment: "最大使用量", Default: 0},
		{Name: "used", Type: field.TypeInt64, Comment: "已使用量", Default: 0},
	}
	// SdkUsageTable holds the schema information for the "sdk_usage" table.
	SdkUsageTable = &schema.Table{
		Name:       "sdk_usage",
		Columns:    SdkUsageColumns,
		PrimaryKey: []*schema.Column{SdkUsageColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ActiveCodeInfoTable,
		AppInfoTable,
		AppSdkTable,
		CategoryInfoTable,
		IndustryInfoTable,
		SdkInfoTable,
		SdkUsageTable,
	}
)

func init() {
	ActiveCodeInfoTable.ForeignKeys[0].RefTable = AppSdkTable
	ActiveCodeInfoTable.Annotation = &entsql.Annotation{
		Table: "active_code_info",
	}
	AppInfoTable.Annotation = &entsql.Annotation{
		Table: "app_info",
	}
	AppSdkTable.ForeignKeys[0].RefTable = AppInfoTable
	AppSdkTable.ForeignKeys[1].RefTable = SdkInfoTable
	AppSdkTable.Annotation = &entsql.Annotation{
		Table: "app_sdk",
	}
	CategoryInfoTable.Annotation = &entsql.Annotation{
		Table: "category_info",
	}
	IndustryInfoTable.Annotation = &entsql.Annotation{
		Table: "industry_info",
	}
	SdkInfoTable.Annotation = &entsql.Annotation{
		Table: "sdk_info",
	}
	SdkUsageTable.Annotation = &entsql.Annotation{
		Table: "sdk_usage",
	}
}
