package setting

import (
    "gorm.io/driver/mysql"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func getMySQLConnection() gorm.Dialector {
    return mysql.Open(Conf.Database.Username + ":" +
        Conf.Database.Password + "@tcp(" + Conf.Database.Host +
        ":" + Conf.Database.Port + ")/" +
        Conf.Database.Database + "?charset=utf8mb4&parseTime=True&loc=Local")
}

func getPostgresConnection() gorm.Dialector {
    return postgres.Open("host=" + Conf.Database.Host +
        " user=" + Conf.Database.Username +
        " password=" + Conf.Database.Password +
        " dbname=" + Conf.Database.Database +
        " port=" + Conf.Database.Port +
        " sslmode=disable TimeZone=Asia/Shanghai")
}

func GetDBDriver() gorm.Dialector {
    switch Conf.Database.Type {
    case DatabaseTypePostgres:
        return getPostgresConnection()

    default:
        return getMySQLConnection()
    }
}