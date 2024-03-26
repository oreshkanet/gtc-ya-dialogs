package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/go-session/session"
	"github.com/oreshkanet/gtc-ya-dialogs/ya-cloud/log"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"time"
)

var oauth2 *server.Server

func initServer(ctx context.Context) {
	if oauth2 != nil {
		return
	}

	log.Info(ctx, "initializing alice app")

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	// manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
	manager.MapAccessGenerate(generates.NewAccessGenerate())

	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "https://social.yandex.ru",
	})
	manager.MapClientStorage(clientStore)

	srvConfig := server.NewConfig()
	oauth2 = server.NewServer(srvConfig, manager)

	/*
		oauth2.SetPasswordAuthorizationHandler(func(ctx context.Context, clientID, username, password string) (userID string, err error) {
			if username == "test" && password == "test" {
				userID = "test"
			}
			return
		})
		oauth2.SetUserAuthorizationHandler(userAuthorizeHandler)

	*/
	oauth2.SetAllowGetAccessRequest(true)
	oauth2.SetClientInfoHandler(server.ClientFormHandler)

	oauth2.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Error(ctx, fmt.Sprintf("Internal Error: %v", err.Error()))
		return
	})

	oauth2.SetResponseErrorHandler(func(re *errors.Response) {
		log.Error(ctx, fmt.Sprintf("Response Error: %v", re.Error.Error()))
	})

	return
}

func Handler(ctx context.Context, req *Request) (*Response, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	ctx, err := initLogging(ctx)
	if err != nil {
		//return nil, err
	}

	defer log.Sync(ctx)

	if req.Request != nil {
		log.Info(ctx, "Request.Command")
		log.Info(ctx, req.Request.Command)
	}

	v, err := json.Marshal(req)
	if err != nil {
		log.Error(ctx, err.Error())
	} else {
		log.Info(ctx, string(v))
	}

	return &Response{}, nil
}

func Handler1(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	ctx, err := initLogging(ctx)
	if err != nil {
		//return nil, err
	}

	defer log.Sync(ctx)

	log.Info(ctx, "URL")
	if r.URL != nil {
		log.Info(ctx, r.URL.String())
	}

	log.Info(ctx, "HEADER")
	if len(r.Header) > 0 {
		v, err := json.Marshal(r.Header)
		if err == nil {
			log.Info(ctx, string(v))
		}
	}

	log.Info(ctx, "BODDY")
	b, err := r.GetBody()
	if err == nil {
		var b1 []byte
		b1, err = io.ReadAll(b)
		if err == nil {
			log.Info(ctx, string(b1))
		}
	}

	log.Info(ctx, "BODDY")
	v, err := json.Marshal(r.URL.Query())
	log.Info(ctx, string(v))

	//initServer(ctx)

	//_, _ = fmt.Fprint(os.Stderr, "fmt.Fprint\n")

	/*
		switch req.URL.Path {
		case "/auth":
			return Auth(ctx, event)
		default:
			return Execute(ctx, event)
		}
	*/

	/*
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error(ctx, err.Error())
			return
		}

	*/

	//log.Debug(ctx, r.URL.String())
	//log.Error(ctx, string(v))
	//log.Debug(ctx, string(bytes))

	/*

	 */

	if r.URL.Query().Get("response_type") != "" {
		Auth(ctx, w, r)
		return
	} else if r.URL.Query().Get("code") != "" {
		Token(ctx, w, r)
		return
	}

	var reqBodyBytes []byte
	if r.Body != nil {
		reqBodyBytes, err = io.ReadAll(r.Body)
		if err != nil {
			log.Error(ctx, err.Error())
			return
		}
	}
	log.Info(ctx, string(reqBodyBytes))

	data := &Request{}
	err = json.Unmarshal(reqBodyBytes, data)
	if err != nil {
		log.Error(ctx, err.Error())
	}

	w.Header().Set("X-Custom-Header", "Test")
	w.WriteHeader(200)
	//name := req.URL.Query().Get("name")
	io.WriteString(w, fmt.Sprintf("%s: %s", r.URL.String(), r.Body))
}

// ?client_id=test&redirect_uri=https%3A%2F%2Fsocial.yandex.net%2Fbroker%2Fredirect&response_type=code&state=https%3A%2F%2Fsocial.yandex.ru%2Fbroker2%2Fauthz_in_web%2F1a02abf8b7ed449c9a4a69cd07000560%2Fcallback: {}

func Auth(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	r.Form = form

	store.Delete("ReturnUri")
	store.Save()

	err = oauth2.HandleAuthorizeRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func Token(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	err := oauth2.HandleTokenRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	//if dumpvar {
	//	_ = dumpRequest(os.Stdout, "userAuthorizeHandler", r) // Ignore the error
	//}
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		store.Set("ReturnUri", r.Form)
		store.Save()

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	userID = uid.(string)
	store.Delete("LoggedInUserID")
	store.Save()
	return
}

func initLogging(ctx context.Context) (context.Context, error) {
	instanceID := "GTC"

	zapConf := zap.NewProductionConfig()
	zapConf.Level.Enabled(zap.DebugLevel)
	zapConf.OutputPaths = []string{"stderr"}
	logger, err := zapConf.Build(zap.AddCallerSkip(3))
	if err != nil {
		return nil, err
	}
	logger = logger.With(zap.String("instanceID", instanceID))
	return log.CtxWithLogger(ctx, logger), nil
}

// Структура запроса API Gateway v1
type APIGatewayRequest struct {
	OperationID string `json:"operationId"`
	Resource    string `json:"resource"`

	HTTPMethod string `json:"httpMethod"`

	Path           string            `json:"path"`
	PathParameters map[string]string `json:"pathParameters"`

	Headers           map[string]string   `json:"headers"`
	MultiValueHeaders map[string][]string `json:"multiValueHeaders"`

	QueryStringParameters           map[string]string   `json:"queryStringParameters"`
	MultiValueQueryStringParameters map[string][]string `json:"multiValueQueryStringParameters"`

	Parameters           map[string]string   `json:"parameters"`
	MultiValueParameters map[string][]string `json:"multiValueParameters"`

	Body            []byte `json:"body"`
	IsBase64Encoded bool   `json:"isBase64Encoded,omitempty"`

	RequestContext interface{} `json:"requestContext"`
}

// Структура ответа API Gateway v1
type APIGatewayResponse struct {
	StatusCode        int                 `json:"statusCode"`
	Headers           map[string]string   `json:"headers"`
	MultiValueHeaders map[string][]string `json:"multiValueHeaders"`
	Body              string              `json:"body"`
	IsBase64Encoded   bool                `json:"isBase64Encoded,omitempty"`
}

type Request1 struct {
	Name string `json:"name"`
}

func Greet(ctx context.Context, event *APIGatewayRequest) (*APIGatewayResponse, error) {
	ctx, err := initLogging(ctx)
	if err != nil {
		//return nil, err
	}

	defer log.Sync(ctx)

	log.Info(ctx, "Handler")

	// В журнале будет напечатано название HTTP-метода, с помощью которого осуществлен запрос, а также путь
	v, err := json.Marshal(event.PathParameters)
	log.Info(ctx, string(v))
	v, err = json.Marshal(event.QueryStringParameters)
	log.Info(ctx, string(v))
	log.Info(ctx, event.Resource)

	req := &Request{}

	// Поле event.Body запроса преобразуется в объект типа Request для получения переданного имени
	if err := json.Unmarshal(event.Body, &req); err != nil {
		return nil, fmt.Errorf("an error has occurred when parsing body: %v", err)
	}

	// Тело ответа.
	return &APIGatewayResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("Hello, %s", req.Meta.ClientId),
	}, nil
}
