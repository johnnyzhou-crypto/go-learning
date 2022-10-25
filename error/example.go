//package main

//
//import (
//	"context"
//	"fmt"
//	"github.com/jackc/pgconn"
//	"github.com/pkg/errors"
//	"net/http"
//	"strconv"
//)
//
//type DBTX interface {
//	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
//	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
//	QueryRow(context.Context, string, ...interface{}) pgx.Row
//}
//
//type myId int
//
//type Result struct {
//	ID myId
//}
//
//// service 层包错误
//// 这个存储层使用外部依赖的 orm
//func getFromRepository(id int) (Result, error) {
//	result := Result{ID: myId(id)}
//	//err := DBTX(&result)
//	var err error
//	if err != nil {
//		msg := fmt.Sprintf("error getting the  result with id %d", id)
//		switch err {
//		case orm.NoResult:
//			err = errors.Wrapf(err, msg)
//		default:
//			err = errors.NotFound(err, msg)
//		}
//		return Result{}, err
//	}
//	return result, nil
//}
//
//// 封装错误后的结果将是
//// err.Error() -> error getting the result with id 10: whatever it comes from the orm
//
//// transfer 层
//func getInteractor(idString string) (Result, error) {
//	id, err := strconv.Atoi(idString)
//	if err != nil {
//		return Result{}, errors.Wrapf(err, "interactor converting id to int")
//	}
//	return repository.getFromRepository(id)
//}
//
////顶层
//r := mux.NewRouter()
//r.HandleFunc("/result/{id}", ResultHandler)
//func ResultHandler(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	result, err := interactor.getInteractor(vars["id"])
//	if err != nil {
//		handleError(w, err)
//	}
//	fmt.Fprintf(w, result)
//}
//func handleError(w http.ResponseWriter, err error) {
//	var status int
//	errorType := errors.GetType(err)
//	switch errorType {
//	case BadRequest:
//		status = http.StatusBadRequest
//	case NotFound:
//		status = http.StatusNotFound
//	default:
//		status = http.StatusInternalServerError
//	}
//	w.WriteHeader(status)
//
//	if errorType == errors.NoType {
//		log.Errorf(err)
//	}
//	fmt.Fprintf(w, "error %s", err.Error())
//
//	errorContext := errors.GetContext(err)
//	if errorContext != nil {
//		fmt.Printf(w, "context %v", errorContext)
//	}
//}
//
////————————————————
////原文作者：Summer
////转自链接：https://learnku.com/go/t/33210
////版权声明：著作权归作者所有。商业转载请联系作者获得授权，非商业转载请保留以上作者信息和原文链接。
