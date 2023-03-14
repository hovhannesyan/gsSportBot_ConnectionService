package utils

import "github.com/hovhannesyan/gsSportBot_ConnectionService/pkg/pb"

func SetKeyToString(info *pb.SetInfo) string {
	return info.SetFor + ":" + info.Id + ":" + info.SetOf
}
