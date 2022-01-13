// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: source.proto

package pb

import (
	context "context"
	http1 "net/http"
	strings "strings"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// SourceServiceHandler is the server API for SourceService service.
type SourceServiceHandler interface {
	// POST /api/pipeline-sources
	Create(context.Context, *PipelineSourceCreateRequest) (*PipelineSourceCreateResponse, error)
	// PUT /api/pipeline-sources/{pipelineSourceID}
	Update(context.Context, *PipelineSourceUpdateRequest) (*PipelineSourceUpdateResponse, error)
	// DELETE /api/pipeline-sources/{pipelineSourceID}
	Delete(context.Context, *PipelineSourceDeleteRequest) (*PipelineSourceDeleteResponse, error)
	// GET /api/pipeline-sources/{pipelineSourceID}
	Get(context.Context, *PipelineSourceGetRequest) (*PipelineSourceGetResponse, error)
	// GET /api/pipeline-sources
	List(context.Context, *PipelineSourceListRequest) (*PipelineSourceListResponse, error)
}

// RegisterSourceServiceHandler register SourceServiceHandler to http.Router.
func RegisterSourceServiceHandler(r http.Router, srv SourceServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		handler := func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
		if h.HTTPInterceptor != nil {
			handler = h.HTTPInterceptor(handler)
		}
		return handler
	}

	add_Create := func(method, path string, fn func(context.Context, *PipelineSourceCreateRequest) (*PipelineSourceCreateResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineSourceCreateRequest))
		}
		var Create_info transport.ServiceInfo
		if h.Interceptor != nil {
			Create_info = transport.NewServiceInfo("erda.core.pipeline.source.SourceService", "Create", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Create_info)
				}
				r = r.WithContext(ctx)
				var in PipelineSourceCreateRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_Update := func(method, path string, fn func(context.Context, *PipelineSourceUpdateRequest) (*PipelineSourceUpdateResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineSourceUpdateRequest))
		}
		var Update_info transport.ServiceInfo
		if h.Interceptor != nil {
			Update_info = transport.NewServiceInfo("erda.core.pipeline.source.SourceService", "Update", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Update_info)
				}
				r = r.WithContext(ctx)
				var in PipelineSourceUpdateRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "pipelineSourceID":
							in.PipelineSourceID = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_Delete := func(method, path string, fn func(context.Context, *PipelineSourceDeleteRequest) (*PipelineSourceDeleteResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineSourceDeleteRequest))
		}
		var Delete_info transport.ServiceInfo
		if h.Interceptor != nil {
			Delete_info = transport.NewServiceInfo("erda.core.pipeline.source.SourceService", "Delete", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Delete_info)
				}
				r = r.WithContext(ctx)
				var in PipelineSourceDeleteRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "pipelineSourceID":
							in.PipelineSourceID = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_Get := func(method, path string, fn func(context.Context, *PipelineSourceGetRequest) (*PipelineSourceGetResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineSourceGetRequest))
		}
		var Get_info transport.ServiceInfo
		if h.Interceptor != nil {
			Get_info = transport.NewServiceInfo("erda.core.pipeline.source.SourceService", "Get", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Get_info)
				}
				r = r.WithContext(ctx)
				var in PipelineSourceGetRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "pipelineSourceID":
							in.PipelineSourceID = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_List := func(method, path string, fn func(context.Context, *PipelineSourceListRequest) (*PipelineSourceListResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineSourceListRequest))
		}
		var List_info transport.ServiceInfo
		if h.Interceptor != nil {
			List_info = transport.NewServiceInfo("erda.core.pipeline.source.SourceService", "List", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, List_info)
				}
				r = r.WithContext(ctx)
				var in PipelineSourceListRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_Create("POST", "/api/pipeline-sources", srv.Create)
	add_Update("PUT", "/api/pipeline-sources/{pipelineSourceID}", srv.Update)
	add_Delete("DELETE", "/api/pipeline-sources/{pipelineSourceID}", srv.Delete)
	add_Get("GET", "/api/pipeline-sources/{pipelineSourceID}", srv.Get)
	add_List("GET", "/api/pipeline-sources", srv.List)
}