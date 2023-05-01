package main

import (
	"github.com/xuri/excelize/v2"
	"strings"
)

type ExcelReader interface {
	SheetName() string
	Each(callback func(dataMap map[int]string) bool) error
}

type excelReaderImpl struct {
	sheetName       string
	dataStartedFrom int
	fieldMap        map[int]string
	excelFile       *excelize.File
}

func New(fileName string) (ExcelReader, error) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, err
	}

	reader := &excelReaderImpl{
		excelFile:       f,
		sheetName:       f.GetSheetName(0),
		fieldMap:        make(map[int]string),
		dataStartedFrom: 0,
	}

	return reader, nil
}

func (reader *excelReaderImpl) SheetName() string {
	return reader.sheetName
}

func (reader *excelReaderImpl) Each(callback func(dataMap map[int]string) bool) error {
	if r, err := reader.excelFile.Rows(reader.sheetName); err != nil {
		return err
	} else {
		for rowCount := 0; r.Next(); rowCount++ {

			if columns, err := r.Columns(); err == nil {
				dataMap := make(map[int]string)

				for i, content := range columns {
					data := strings.TrimSpace(content)
					dataMap[i] = data
				}
				if len(dataMap) > 0 {
					if !callback(dataMap) {
						return nil
					}
				}
			}
		}
		return nil
	}
}
