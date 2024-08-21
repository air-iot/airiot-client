package api_client_go

import (
	"context"
	"fmt"
	"io"
	netHttp "net/http"
	"net/url"

	"github.com/air-iot/api-client-go/v4/api"
	"github.com/air-iot/api-client-go/v4/apicontext"
	"github.com/air-iot/api-client-go/v4/apitransport"
	"github.com/air-iot/api-client-go/v4/config"
	"github.com/air-iot/api-client-go/v4/core"
	"github.com/air-iot/errors"
	"github.com/air-iot/json"
	"github.com/air-iot/logger"
)

func (c *Client) GetFileLicense(ctx context.Context, result interface{}) error {
	cli, err := c.CoreClient.GetLicenseServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.GetFileLicense(ctx, &api.QueryRequest{})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) UseLicense(ctx context.Context, projectId string, result interface{}) error {
	cli, err := c.CoreClient.GetLicenseServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.UseLicense(apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}), &api.QueryRequest{})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) UploadLicense(ctx context.Context, projectId, filename string, size int, r io.Reader) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}

	cli, err := c.CoreClient.GetLicenseServiceClient()
	if err != nil {
		return err
	}
	stream, err := cli.UploadLicense(apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId, "filename": filename}))
	if err != nil {
		return errors.NewResErrorMsg(err, "请求错误")
	}

	//defer stream.CloseAndRecv()

	buffer := make([]byte, 1024)

	bytesReadAll := 0

	for {
		bytesRead, err := r.Read(buffer)
		if err != nil {
			return errors.Wrap(err, "读取授权文件错误")
		}
		err = stream.Send(&core.UploadFileRequest{Data: buffer[:bytesRead]})
		if err != nil {
			return errors.Wrap(err, "grpc发送授权文件错误")
		}
		bytesReadAll += bytesRead

		if bytesReadAll == size {
			//err := stream.CloseSend()
			//if err != nil {
			//	return fmt.Errorf("CloseSend错误:%s", err.Error())
			//}

			err := stream.Send(&core.UploadFileRequest{Data: []byte("down")})
			if err != nil {
				return errors.Wrap(err, "grpc发送结束标志错误")
			}

			m := new(api.Response)
			err = stream.RecvMsg(m)
			if err != nil {
				return errors.Wrap(err, "读取server响应错误")
			}
			fmt.Printf("上传文件成功，服务器响应结果:%+v\n", m)

			return nil
		}

	}
}

func (c *Client) GetDriverLicense(ctx context.Context, projectId, driverId string, result interface{}) error {
	cli, err := c.CoreClient.GetLicenseServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.GetDriverLicense(apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}), &api.GetOrDeleteRequest{
		Id: driverId,
	})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) FindMachineCode(ctx context.Context, result interface{}) error {
	cli, err := c.CoreClient.GetLicenseServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.FindMachineCode(ctx, &api.QueryRequest{})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetCurrentUserInfo(ctx context.Context, projectId, token string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if token == "" {
		return errors.New("token is empty")
	}
	cli, err := c.CoreClient.GetUserServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.GetCurrentUserInfo(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId, config.XRequestHeaderAuthorization: token}),
		&core.LoginUserRequest{})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) UserPermissionUpdate(ctx context.Context, projectId, token string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if token == "" {
		return errors.New("token is empty")
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.CoreClient.GetUserServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.UserPermissionUpdate(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId, config.XRequestHeaderAuthorization: token}),
		&api.CreateRequest{Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) QueryUser(ctx context.Context, projectId string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetUserServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) QueryUserBackup(ctx context.Context, projectId string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetUserServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.QueryBackup(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteManyUserBackup(ctx context.Context, projectId string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetUserServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.DeleteManyBackup(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateManyUserBackup(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.CoreClient.GetUserServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.CreateManyBackup(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.CreateRequest{
			Data: bts,
		})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetUser(ctx context.Context, projectId, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return nil, errors.New("id为空")
	}
	cli, err := c.CoreClient.GetUserServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.Get(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	return parseRes(err, res, result)
}

func (c *Client) DeleteUser(ctx context.Context, projectId, id string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.CoreClient.GetUserServiceClient()
	if err != nil {
		return err
	}

	res, err := cli.Delete(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateUser(ctx context.Context, projectId, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}

	cli, err := c.CoreClient.GetUserServiceClient()
	if err != nil {
		return err
	}

	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据为空")
	}
	res, err := cli.Update(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.UpdateRequest{Id: id, Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) ReplaceUser(ctx context.Context, projectId, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	cli, err := c.CoreClient.GetUserServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据为空")
	}
	res, err := cli.Replace(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.UpdateRequest{Id: id, Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateLog(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}

	cli, err := c.CoreClient.GetLogServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.Create(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.CreateRequest{Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateUser(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.CoreClient.GetUserServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.Create(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.CreateRequest{Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) StatsQuery(ctx context.Context, projectId string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetTableSchemaServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.StatsQuery(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) QueryTableSchema(ctx context.Context, projectId string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetTableSchemaServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) RestQueryTableSchema(ctx context.Context, projectId string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	u := url.URL{Path: "/core/t/schema"}
	if query != nil {
		bts, err := json.Marshal(query)
		if err != nil {
			return errors.Wrap(err, "序列化查询参数为空")
		}
		params := url.Values{}
		params.Set("query", string(bts))
		u.RawQuery = params.Encode()
	}
	cli, err := c.CoreClient.GetRestClient()
	if err != nil {
		return err
	}
	if err := cli.Invoke(apitransport.NewClientContext(ctx, &apitransport.Transport{ReqHeader: map[string]string{config.XRequestProject: projectId}}), netHttp.MethodGet, u.RequestURI(), map[string]interface{}{}, result); err != nil {
		return errors.NewResError(err)
	}
	return nil
}

func (c *Client) QueryTableSchemaDeviceByDriverAndGroup(ctx context.Context, projectId, driverId, groupId string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if groupId == "" {
		return errors.New("实例组ID为空")
	}
	cli, err := c.CoreClient.GetTableSchemaServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.QueryDeviceByDriverAndGroup(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.GetDeviceRequest{Driver: driverId, Group: groupId})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) QueryOnlyTableSchemaDeviceByDriverAndGroup(ctx context.Context, projectId, driverId, groupId string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if groupId == "" {
		return errors.New("实例组ID为空")
	}
	cli, err := c.CoreClient.GetTableSchemaServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.QueryTableDeviceByDriverAndGroup(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.GetDeviceRequest{Driver: driverId, Group: groupId})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) FindDevice(ctx context.Context, projectId, driverId, groupId, tableId, deviceId string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if groupId == "" {
		return errors.New("实例组ID为空")
	}
	if deviceId == "" {
		return errors.New("设备ID为空")
	}
	cli, err := c.CoreClient.GetTableSchemaServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.FindDevice(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.GetDataDeviceRequest{Driver: driverId, Group: groupId, Table: tableId, Id: deviceId})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) QueryEmulator(ctx context.Context, projectId string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	cli, err := c.CoreClient.GetTableSchemaServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.QueryEmulator(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetTableSchema(ctx context.Context, projectId, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return nil, errors.New("id为空")
	}
	cli, err := c.CoreClient.GetTableSchemaServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.Get(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	return parseRes(err, res, result)
}

func (c *Client) DeleteTableSchema(ctx context.Context, projectId, id string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.CoreClient.GetTableSchemaServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.Delete(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateTableSchema(ctx context.Context, projectId, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	cli, err := c.CoreClient.GetTableSchemaServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据为空")
	}
	res, err := cli.Update(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.UpdateRequest{Id: id, Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) ReplaceTableSchema(ctx context.Context, projectId, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	cli, err := c.CoreClient.GetTableSchemaServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据为空")
	}
	res, err := cli.Replace(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.UpdateRequest{Id: id, Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateTableSchema(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.CoreClient.GetTableSchemaServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.Create(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.CreateRequest{Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) QueryTableRecord(ctx context.Context, projectId string, query, result interface{}) (int64, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return 0, errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetTableRecordServiceClient()
	if err != nil {
		return 0, err
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return 0, err
	}
	return res.GetCount(), nil
}

func (c *Client) GetTableRecord(ctx context.Context, projectId, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return nil, errors.New("id为空")
	}
	cli, err := c.CoreClient.GetTableRecordServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.Get(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	return parseRes(err, res, result)
}

func (c *Client) DeleteTableRecord(ctx context.Context, projectId, id string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.CoreClient.GetTableRecordServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.Delete(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateTableRecord(ctx context.Context, projectId, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	cli, err := c.CoreClient.GetTableRecordServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据为空")
	}
	res, err := cli.Update(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.UpdateRequest{Id: id, Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) ReplaceTableRecord(ctx context.Context, projectId, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	cli, err := c.CoreClient.GetTableRecordServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据为空")
	}
	res, err := cli.Replace(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.UpdateRequest{Id: id, Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateTableRecord(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.CoreClient.GetTableRecordServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.Create(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.CreateRequest{Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) QueryTableData(ctx context.Context, projectId, tableName string, query, result interface{}) (int64, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if tableName == "" {
		return 0, errors.New("表为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return 0, err
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return 0, errors.Wrap(err, "序列化查询参数为空")
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.QueryDataRequest{
			Table: tableName,
			Query: bts,
		})
	if _, err := parseRes(err, res, result); err != nil {
		return 0, err
	}
	return res.GetCount(), nil
}

func (c *Client) QueryTableDataByTableId(ctx context.Context, projectId, tableId string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if tableId == "" {
		return errors.New("记录id为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	res, err := cli.QueryByTableId(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.QueryDataRequest{
			Table: tableId,
			Query: bts,
		})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetTableData(ctx context.Context, projectId, tableName, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if tableName == "" {
		return nil, errors.New("表为空")
	}
	if id == "" {
		return nil, errors.New("id为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.Get(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.GetOrDeleteDataRequest{Table: tableName, Id: id})
	return parseRes(err, res, result)
}

func (c *Client) DeleteTableData(ctx context.Context, projectId, tableName, id string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if tableName == "" {
		return errors.New("表为空")
	}
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.Delete(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.GetOrDeleteDataRequest{Table: tableName, Id: id})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteManyTableData(ctx context.Context, projectId, tableName string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	if tableName == "" {
		return errors.New("表为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.DeleteMany(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.QueryDataRequest{Table: tableName, Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateTableData(ctx context.Context, projectId, tableName, id string, closeRequire bool, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if tableName == "" {
		return errors.New("表为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据为空")
	}
	res, err := cli.Update(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.UpdateDataRequest{Table: tableName, Id: id, Data: bts, CloseRequire: closeRequire})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) ReplaceTableData(ctx context.Context, projectId, tableName, id string, closeRequire bool, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	if tableName == "" {
		return errors.New("表为空")
	}
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据为空")
	}
	res, err := cli.Replace(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.UpdateDataRequest{Table: tableName, Id: id, Data: bts, CloseRequire: closeRequire})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateTableData(ctx context.Context, projectId, tableName string, closeRequire bool, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if tableName == "" {
		return errors.New("表为空")
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.Create(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.CreateDataRequest{
			Table:        tableName,
			Data:         bts,
			CloseRequire: closeRequire,
		})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateManyTableData(ctx context.Context, projectId, tableName string, closeRequire bool, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if tableName == "" {
		return errors.New("表为空")
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.CreateMany(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.CreateDataRequest{
			Table:        tableName,
			Data:         bts,
			CloseRequire: closeRequire,
		})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateMessage(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.CoreClient.GetMessageServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.Create(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.CreateRequest{Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) QueryMessage(ctx context.Context, projectId string, query, result interface{}) (int, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return 0, errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetMessageServiceClient()
	if err != nil {
		return 0, err
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return 0, err
	}
	return int(res.GetCount()), nil
}

func (c *Client) GetLog(ctx context.Context, projectId, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return nil, errors.New("id为空")
	}
	cli, err := c.CoreClient.GetLogServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.Get(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	return parseRes(err, res, result)
}

func (c *Client) QueryLog(ctx context.Context, projectId string, query, result interface{}) (int, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return 0, errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetLogServiceClient()
	if err != nil {
		return 0, err
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return 0, err
	}
	return int(res.GetCount()), nil
}

func (c *Client) QueryApiPermissionLog(ctx context.Context, projectId string, query, result interface{}) (int, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return 0, errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetLogServiceClient()
	if err != nil {
		return 0, err
	}
	res, err := cli.QueryApiPermission(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return 0, err
	}
	return int(res.GetCount()), nil
}

func (c *Client) PostLatest(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.CoreClient.GetDataQueryServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.PostLatest(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.CreateRequest{Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetQuery(ctx context.Context, projectId string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetDataQueryServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.GetQuery(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) PostQuery(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.CoreClient.GetDataQueryServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.PostQuery(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.CreateRequest{Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) QueryRole(ctx context.Context, projectId string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetRoleServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) AdminRoleCheck(ctx context.Context, projectId, token string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if token == "" {
		return errors.New("无Token认证信息")
	}
	cli, err := c.CoreClient.GetRoleServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.AdminRoleCheck(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId, config.XRequestHeaderAuthorization: token}),
		&api.EmptyRequest{})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetRole(ctx context.Context, projectId, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return nil, errors.New("id为空")
	}
	cli, err := c.CoreClient.GetRoleServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.Get(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	return parseRes(err, res, result)
}

func (c *Client) FindTableDataDeptByDeptIDs(ctx context.Context, projectId string, ids map[string]interface{}, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(ids)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	res, err := cli.FindTableDataDeptByDeptIDs(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.CreateRequest{
			Data: bts,
		})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateManyTableData(ctx context.Context, projectId, tableName string, closeRequire bool, query, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	if tableName == "" {
		return errors.New("表为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	btsUpdate, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据为空")
	}
	res, err := cli.UpdateMany(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.MultiUpdateDataRequest{Table: tableName, Query: bts, Data: btsUpdate, CloseRequire: closeRequire})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetWarningFilterIDs(ctx context.Context, projectId, token string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	if token == "" {
		return errors.New("无Token认证信息")
	}
	res, err := cli.GetWarningFilterIDs(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId, config.XRequestHeaderAuthorization: token}),
		&api.EmptyRequest{})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) QueryCatalog(ctx context.Context, projectId string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetCatalogServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetCatalog(ctx context.Context, projectId, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return nil, errors.New("id为空")
	}
	cli, err := c.CoreClient.GetCatalogServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.Get(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	return parseRes(err, res, result)
}

func (c *Client) QueryDept(ctx context.Context, projectId string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetDeptServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetDept(ctx context.Context, projectId, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return nil, errors.New("id为空")
	}
	cli, err := c.CoreClient.GetDeptServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.Get(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	return parseRes(err, res, result)
}

func (c *Client) QuerySetting(ctx context.Context, projectId string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetSettingServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) QueryApp(ctx context.Context, projectId string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetAppServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetApp(ctx context.Context, projectId, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return nil, errors.New("id为空")
	}
	cli, err := c.CoreClient.GetAppServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.Get(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	return parseRes(err, res, result)
}

func (c *Client) QuerySystemVariable(ctx context.Context, projectId string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetSystemVariableServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetSystemVariable(ctx context.Context, projectId, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return nil, errors.New("id为空")
	}
	cli, err := c.CoreClient.GetSystemVariableServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.Get(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	return parseRes(err, res, result)
}

func (c *Client) DeleteSystemVariable(ctx context.Context, projectId, id string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.CoreClient.GetSystemVariableServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.Delete(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateSystemVariable(ctx context.Context, projectId, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}

	cli, err := c.CoreClient.GetSystemVariableServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据为空")
	}
	res, err := cli.Update(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.UpdateRequest{Id: id, Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) ReplaceSystemVariable(ctx context.Context, projectId, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	cli, err := c.CoreClient.GetSystemVariableServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据为空")
	}
	res, err := cli.Replace(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.UpdateRequest{Id: id, Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateSystemVariable(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.CoreClient.GetSystemVariableServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.Create(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.CreateRequest{Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) QueryBackup(ctx context.Context, projectId string, query, result interface{}, count *int64) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetBackupServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	*count = res.GetCount()
	return nil
}

func (c *Client) GetBackup(ctx context.Context, projectId, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return nil, errors.New("id为空")
	}
	cli, err := c.CoreClient.GetBackupServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.Get(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	return parseRes(err, res, result)
}

func (c *Client) DeleteBackup(ctx context.Context, projectId, id string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.CoreClient.GetBackupServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.Delete(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateBackup(ctx context.Context, projectId, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}

	cli, err := c.CoreClient.GetBackupServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据为空")
	}
	res, err := cli.Update(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.UpdateRequest{Id: id, Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) ExportBackup(ctx context.Context, projectId string, query interface{}) (string, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return "", errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetBackupServiceClient()
	if err != nil {
		return "", err
	}
	res, err := cli.Export(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, nil); err != nil {
		return "", err
	}
	id := string(res.GetResult())
	return id, nil
}

func (c *Client) ImportBackup(ctx context.Context, projectId string, query interface{}) (string, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return "", errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetBackupServiceClient()
	if err != nil {
		return "", err
	}
	res, err := cli.Import(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, nil); err != nil {
		return "", err
	}
	id := string(res.GetResult())
	return id, nil
}

func (c *Client) UploadBackup(ctx context.Context, projectId, password string, size int, r io.Reader) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}

	cli, err := c.CoreClient.GetBackupServiceClient()
	if err != nil {
		return err
	}
	stream, err := cli.Upload(apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId, "password": password}))
	if err != nil {
		return errors.NewResErrorMsg(err, "请求错误")
	}

	//defer stream.CloseAndRecv()

	buffer := make([]byte, 1024)

	bytesReadAll := 0

	for {
		bytesRead, err := r.Read(buffer)
		if err != nil {
			return fmt.Errorf("读取备份文件错误:%s", err.Error())
		}
		err = stream.Send(&core.UploadFileRequest{Data: buffer[:bytesRead]})
		if err != nil {
			return fmt.Errorf("grpc发送备份文件错误:%s", err.Error())
		}
		bytesReadAll += bytesRead

		if bytesReadAll == size {
			err := stream.CloseSend()
			if err != nil {
				return fmt.Errorf("CloseSend错误:%s", err.Error())
			}

			//err := stream.Send(&core.UploadFileRequest{Data: []byte("down")})
			//if err != nil {
			//	return fmt.Errorf("grpc发送结束标志错误:%s", err.Error())
			//}
			//
			//m := new(api.Response)
			//err = stream.RecvMsg(m)
			//if err != nil {
			//	return fmt.Errorf("读取server响应错误:%s", err.Error())
			//}
			//fmt.Printf("上传文件成功，服务器响应结果:%+v\n", m)

			return nil
		}

	}
}

func (c *Client) DownloadBackup(ctx context.Context, projectId, id, password string, w io.Writer) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}

	cli, err := c.CoreClient.GetBackupServiceClient()
	if err != nil {
		return err
	}

	in := new(api.GetOrDeleteRequest)
	in.Id = id
	stream, err := cli.Download(apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId, "password": password, "id": id}), in)
	if err != nil {
		return errors.NewResErrorMsg(err, "请求错误")
	}

	defer func() {
		_ = stream.CloseSend()
	}()

	for {
		d, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return errors.Wrap(err, "stream.Recv错误")
		}

		logger.Infof("数据长度:%+v", len(d.GetData()))

		n, err := w.Write(d.GetData())
		if err != nil {
			return errors.Wrap(err, " w.Write错误")
		}
		logger.Infof("写入数据长度:%+v", n)
	}

	// 方法二
	//d, err := stream.Recv()
	//if err != nil {
	//	return errors.NewMsg("stream.Recv错误, %s", err.Error())
	//}
	//
	//logger.Infof("数据长度:%+v", len(d.GetData()))
	//
	//n, err := w.Write(d.GetData())
	//if err != nil {
	//	return errors.NewMsg(" w.Write错误, %s", err.Error())
	//}
	//logger.Infof("写入数据长度:%+v", n)
	//return nil

	//for {
	//
	//	err = stream.RecvMsg(&buffer)
	//	if err == io.EOF {
	//		return nil
	//	}
	//	if err != nil {
	//		return err
	//	}
	//
	//	w.Write(buffer)
	//}
}

func (c *Client) FindTagByID(ctx context.Context, projectId, tableId, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if tableId == "" {
		return nil, errors.New("表为空")
	}
	if id == "" {
		return nil, errors.New("设备id为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.FindTagByID(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.GetOrDeleteDataRequest{Table: tableId, Id: id})
	return parseRes(err, res, result)
}

func (c *Client) QueryTaskManager(ctx context.Context, projectId string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetTaskManagerServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetTaskManager(ctx context.Context, projectId, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return nil, errors.New("id为空")
	}
	cli, err := c.CoreClient.GetTaskManagerServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.Get(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	return parseRes(err, res, result)
}

func (c *Client) DeleteTaskManager(ctx context.Context, projectId, id string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.CoreClient.GetTaskManagerServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.Delete(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateTaskManager(ctx context.Context, projectId, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}

	cli, err := c.CoreClient.GetTaskManagerServiceClient()
	if err != nil {
		return err
	}

	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据为空")
	}
	res, err := cli.Update(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.UpdateRequest{Id: id, Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) ReplaceTaskManager(ctx context.Context, projectId, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	cli, err := c.CoreClient.GetTaskManagerServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据为空")
	}
	res, err := cli.Replace(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.UpdateRequest{Id: id, Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateTaskManager(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.CoreClient.GetTaskManagerServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.Create(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.CreateRequest{Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) FindTableCommandById(ctx context.Context, projectId, id string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.CoreClient.GetTableSchemaServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.FindCommandByID(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) FindTableDataCommandById(ctx context.Context, projectId, tableId, id string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if tableId == "" || id == "" {
		return errors.New("表或记录id为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.FindCommandByID(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.GetOrDeleteDataRequest{Table: tableId, Id: id})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) QueryTableDataByDB(ctx context.Context, projectId, tableName string, query, result interface{}) (int64, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if tableName == "" {
		return 0, errors.New("表为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return 0, err
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return 0, errors.Wrap(err, "序列化查询参数为空")
	}
	res, err := cli.QueryByDB(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.QueryDataRequest{
			Table: tableName,
			Query: bts,
		})
	if _, err := parseRes(err, res, result); err != nil {
		return 0, err
	}
	return res.GetCount(), nil
}

func (c *Client) GetTableDataByDB(ctx context.Context, projectId, tableName, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if tableName == "" {
		return nil, errors.New("表为空")
	}
	if id == "" {
		return nil, errors.New("id为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.GetByDB(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.GetOrDeleteDataRequest{Table: tableName, Id: id})
	return parseRes(err, res, result)
}

func (c *Client) DeleteTableDataByDB(ctx context.Context, projectId, tableName, id string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if tableName == "" {
		return errors.New("表为空")
	}
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.DeleteByDB(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.GetOrDeleteDataRequest{Table: tableName, Id: id})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteManyTableDataByDB(ctx context.Context, projectId, tableName string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	if tableName == "" {
		return errors.New("表为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.DeleteManyByDB(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.QueryDataRequest{Table: tableName, Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateTableDataByDB(ctx context.Context, projectId, tableName, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if tableName == "" {
		return errors.New("表为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据为空")
	}
	res, err := cli.UpdateByDB(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.UpdateDataRequest{Table: tableName, Id: id, Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) ReplaceTableDataByDB(ctx context.Context, projectId, tableName, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	if tableName == "" {
		return errors.New("表为空")
	}
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据为空")
	}
	res, err := cli.ReplaceByDB(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.UpdateDataRequest{Table: tableName, Id: id, Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateTableDataByDB(ctx context.Context, projectId, tableName string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if tableName == "" {
		return errors.New("表为空")
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.CreateByDB(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.CreateDataRequest{
			Table: tableName,
			Data:  bts,
		})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateManyTableDataByDB(ctx context.Context, projectId, tableName string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if tableName == "" {
		return errors.New("表为空")
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.CreateManyByDB(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.CreateDataRequest{
			Table: tableName,
			Data:  bts,
		})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateManyTableDataByDB(ctx context.Context, projectId, tableName string, updateDataList, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if tableName == "" {
		return errors.New("表为空")
	}
	cli, err := c.CoreClient.GetTableDataServiceClient()
	if err != nil {
		return err
	}
	btsUpdate, err := json.Marshal(updateDataList)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据为空")
	}
	res, err := cli.UpdateManyByDB(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&core.MultiUpdateDataRequest{Table: tableName, Data: btsUpdate})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateDashboard(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.CoreClient.GetDashboardServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.Create(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.CreateRequest{Data: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) QueryDashboard(ctx context.Context, projectId string, query, result interface{}) (int, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return 0, errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.CoreClient.GetDashboardServiceClient()
	if err != nil {
		return 0, err
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if _, err := parseRes(err, res, result); err != nil {
		return 0, err
	}
	return int(res.GetCount()), nil
}

// UploadFileFromUrl 将远程文件上传到媒体库
//
// sourceUrl 远程文件的下载 url
//
// catalog 上传到媒体库的目录
//
// filename 上传到媒体库后的文件名
//
// 上传成功后返回文件的访问地址
func (c *Client) UploadFileFromUrl(ctx context.Context, projectId string, sourceUrl string, catalog string, filename string) (string, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}

	body := map[string]string{
		"fileUrl":          sourceUrl,
		"mediaLibraryPath": catalog,
		"saveFileName":     filename,
	}

	cli, err := c.CoreClient.GetRestClient()
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := cli.Invoke(apitransport.NewClientContext(ctx, &apitransport.Transport{ReqHeader: map[string]string{config.XRequestProject: projectId}}),
		"POST", "/core/mediaLibrary/saveFileFromUrl",
		body, &result); err != nil {
		return "", errors.NewResErrorMsg(err, "请求错误")
	}

	fileUrl, ok := result["url"]
	if !ok {
		return "", errors.New("上传媒体库成功, 但未返回文件的 url")
	}

	fileUrlStr, ok := fileUrl.(string)
	if !ok {
		return "", fmt.Errorf("上传媒体库成功, 但返回文件的 url 不是字符串, %+v", fileUrl)
	}

	return fileUrlStr, nil
}
