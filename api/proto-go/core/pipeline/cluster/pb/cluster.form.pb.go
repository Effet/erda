// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: cluster.proto

package pb

import (
	json "encoding/json"
	url "net/url"
	strings "strings"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*ClusterHookRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ClusterHookResponse)(nil)

// ClusterHookRequest implement urlenc.URLValuesUnmarshaler.
func (m *ClusterHookRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "event":
				m.Event = vals[0]
			case "action":
				m.Action = vals[0]
			case "orgID":
				m.OrgID = vals[0]
			case "projectID":
				m.ProjectID = vals[0]
			case "applicationID":
				m.ApplicationID = vals[0]
			case "env":
				m.Env = vals[0]
			case "timeStamp":
				m.TimeStamp = vals[0]
			case "content":
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Content = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Content = val
					} else {
						m.Content = structpb.NewStringValue(vals[0])
					}
				}
			}
		}
	}
	return nil
}

// ClusterHookResponse implement urlenc.URLValuesUnmarshaler.
func (m *ClusterHookResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}