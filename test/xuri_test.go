package test

//
//func Test_ExcelWriter(t *testing.T) {
//	f := excelize.NewFile()
//	defer func() {
//		if err := f.Close(); err != nil {
//			fmt.Println(err)
//		}
//	}()
//	// 创建一个工作表
//	index, err := f.NewSheet("Sheet2")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	// 设置单元格的值
//	f.SetCellValue("Sheet2", "A2", "Hello world.")
//	f.SetCellValue("Sheet1", "B2", 100)
//	// 设置工作簿的默认工作表
//	f.SetActiveSheet(index)
//
//	buf, err := f.WriteToBuffer()
//	if err != nil {
//		return
//	}
//
//	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
//
//	// 根据指定路径保存文件
//	fmt.Println("f", base64Str)
//}
//
//func Test_ExcelReader(t *testing.T) {
//	fileReader, err := os.OpenFile("users.xlsx", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//	if err != nil {
//		log.Fatalf("open excel file err, error:%+v\n", err)
//	}
//	file, err := excelize.OpenReader(fileReader)
//	if err != nil {
//		log.Fatalf("open excel file err, error:%+v\n", err)
//	}
//
//	for _, sheetName := range file.GetSheetList() {
//		rows, err := file.GetRows(sheetName)
//		if err != nil {
//			log.Printf("excel get rows error, err:%+v\n", err)
//		}
//		for index, row := range rows {
//			if index > 0 {
//				fmt.Println(index, row)
//				//strings.TrimSpace(cols[0])
//			}
//
//		}
//	}
//
//	//ParseEx(0, []*PeopleInfo{{}})
//
//	//ef := excelize.NewFile()
//	//WriteFirstRow(ef, 0, []*PeopleInfo{{PeopleNo: "test", PeopleName: "test", BirthDate: "2020-01-01", EmploymentStatus: "在职"}})
//	//ef.SaveAs("people_info.xlsx")
//
//}
//
//func WriteToExcel(stocks []PeopleInfo, fileName string) error {
//	file := excelize.NewFile()
//	rowsCount := len(stocks)
//	for i := 1; i <= rowsCount; i++ {
//		if err := file.SetSheetRow("stock",
//			fmt.Sprintf("A%d", i),
//			&[]interface{}{
//				stocks[i-1].PeopleNo,
//				stocks[i-1].PeopleName,
//				stocks[i-1].BirthDate,
//				stocks[i-1].EmploymentStatus,
//			}); err != nil {
//			return err
//		}
//	}
//	defer func(file *excelize.File) {
//		if err := file.Close(); err != nil {
//		}
//	}(file)
//	if err := file.SaveAs(fileName); err != nil {
//		return err
//	}
//	return nil
//}
//
//type PeopleInfo struct {
//	PeopleNo         string `ex:"head:工号;type:string;required;color:#0070C0"`
//	PeopleName       string `ex:"head:姓名;type:string;required"`
//	BirthDate        string `ex:"head:出生日期;type:date;omitempty"`
//	EmploymentStatus string `ex:"head:在职状态;type:string;select:在职,离职"`
//}
//
//type Setting struct {
//	Head      string
//	Type      string
//	Select    []string
//	Required  bool
//	OmitEmpty bool
//	Color     string
//}
//
//// 解析第idx个字段的ex
//func ParseEx(idx int, data interface{}) *Setting {
//	tp := reflect.ValueOf(data).Type().Elem().Elem() // 获得结构体的反射Type
//	field := tp.Field(idx)
//	exStr := field.Tag.Get("ex") // 获得tag中ex的值
//	setting := &Setting{}
//
//	pairs := strings.Split(exStr, ";")
//	for _, pair := range pairs {
//		kv := strings.Split(pair, ":")
//		key := strings.TrimSpace(kv[0])
//		value := "true"
//		if len(kv) == 2 {
//			value = strings.TrimSpace(kv[1])
//		}
//
//		println("pair kv", key, value)
//		switch key {
//		case "head":
//			setting.Head = value
//		case "type":
//			setting.Type = value
//		case "select":
//			setting.Select = strings.Split(value, ";")
//		case "required":
//			setting.Required = value == "true"
//		case "omitempty":
//			setting.OmitEmpty = value == "true"
//		case "color":
//			setting.Color = value
//		default:
//			println("unknown key:", key)
//		}
//	}
//	println("setting Head", setting.Head)
//	return setting
//}
//
//// 写入第1行数据的第idx个字段
//func WriteFirstRow(ef *excelize.File, idx int, data interface{}) error {
//	firstRow := reflect.ValueOf(data).Index(0).Elem() // 第1个数据的反射Value
//	v := firstRow.Field(idx)                          // 第idx个字段的反射Value
//	setting := ParseEx(idx, data)                     // 第idx个字段解析出来的ex信息
//
//	// 处理omitempty
//	if setting.OmitEmpty && v.IsZero() {
//		return nil
//	}
//
//	var val interface{}
//	// 处理type
//	switch setting.Type {
//	case "string":
//		val = v.String()
//	}
//
//	// Excel列号从1开始，所以列号是idx+1；行号从2开始，因为第1行要显示列名
//	axis, err := excelize.CoordinatesToCellName(idx+1, 2)
//	if err != nil {
//		return err
//	}
//
//	// 将数据写入默认工作表Sheet1中axis坐标处的单元格
//	return ef.SetCellValue("Sheet1", axis, val)
//}
