package service

type (
    DataSourceMgr interface {
        AuthCheck(key, secret string) bool
    }

    DataSourceService struct {
    }

)


func NewDataSourceMgr() (DataSourceMgr, error) {
    return &DataSourceService{}, nil
}


func (ds *DataSourceService) AuthCheck(key, secret string) bool {
    return true
}