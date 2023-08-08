package service

import (
	"context"
	"encoding/json"
	"fmt"
	"sign/proto"
	"time"
)

type sign interface {
	GetServiceSign([]byte) string
}

var (
	mp        = make(map[string]sign)
	hashValue string
	hashCode  int32
)

func init() {
	mp["md5"] = &Md5Sign{}
	mp["sha256"] = &ShaSign{}
	mp["rsa"] = &RsaSign{}
}

type SignServer struct {
}

func (s *SignServer) mustEmbedUnimplementedSignServiceRequestServer() {}
func (s *SignServer) GetSign(ctx context.Context, req *proto.SignRequest) (*proto.SignReponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	if signTool, ok := mp[req.Type]; ok {
		signTool.GetServiceSign(data)
	}
	var flag bool
	timer := time.After(10 * time.Second)
	for !flag {
		select {
		case <-timer:
			fmt.Println("10s is coming")
			flag = true
		default:
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now().Second())
		}
	}
	response := &proto.SignReponse{
		Sign: hashValue,
		Code: hashCode,
	}
	return response, nil
}
