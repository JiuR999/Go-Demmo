package main

import (
	"AndroidToolServer-Go/common/utils"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"strings"
)

func main() {

	//generate(utils.CreateMySQLDB(), "zb_dream")
	/*db, _ := gorm.Open(mysql.Open(MysqlConfig))
	ma, _ := query.Use(db).ApArticle.Take()
	fmt.Println(ma)*/

	pgrDB, err := utils.CreatePostgreSQLDB()
	if err == nil {
		generate(pgrDB, "dgs_general_failures")
	}
}

func generate(db *gorm.DB, tables ...string) {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./query",
		Mode:              gen.WithoutContext,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
	})
	g.UseDB(db)

	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		appendString := false
		if strings.Contains(columnName, "_id") {
			appendString = true
		}
		camelNameArray := strings.Split(columnName, "_")
		for i, _ := range camelNameArray {
			if i == 0 {
				continue
			}
			camelNameArray[i] = strings.Title(camelNameArray[i])
		}
		camelName := strings.Join(camelNameArray, "")

		if appendString {
			return camelName + ",string"
		}
		return camelName
	})
	//将非默认字段名定义为自动时间戳和软删除字段
	//自动时间戳默认字段名为“update_time","create_time",表字段类型为INT 或 DATETIME
	//软删除字段默认字段名为：”deleted_at,表字段数据类型为DATETIME
	autoUpdateTimeField := gen.FieldGORMTag("update_time", func(tag field.GormTag) field.GormTag {
		tag.Set("column", "update_time")
		tag.Set("type", "datetime(0)")
		tag.Set("autoUpdateTime")
		return tag
	})

	autoCreateTimeField := gen.FieldGORMTag("create_time", func(tag field.GormTag) field.GormTag {
		tag.Set("column", "create_time")
		tag.Set("type", "datetime(0)")
		tag.Set("autoUpdateTime")
		return tag
	})

	softDeleteField := gen.FieldType("delete_time", "soft_delete.DeletedAt")

	fieldOpts := []gen.ModelOpt{
		jsonField, autoCreateTimeField, autoUpdateTimeField, softDeleteField,
	}

	if tables != nil {
		//创建模型的结构体 生成的文件在query目录;先创建结果不会被覆盖
		for _, table := range tables {
			model := g.GenerateModel(table, fieldOpts...)
			g.ApplyBasic(model)
		}
	} else {
		//创建全部模型文件，并覆盖前面创建的同名模型
		allModel := g.GenerateAllTable(fieldOpts...)
		g.ApplyBasic(allModel...)
	}
	g.Execute()
}
