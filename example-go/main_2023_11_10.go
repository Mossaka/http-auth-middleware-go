package main

import (handler "github.com/Mossaka/hello-wasi-http-go/target_world/2023_11_10")

func init() {
	a := HttpImpl{}
	handler.SetExportsWasiHttp0_2_0_rc_2023_11_10_IncomingHandler(a)
}

type HttpImpl struct {

}

func (h HttpImpl) Handle(request handler.ExportsWasiHttp0_2_0_rc_2023_11_10_IncomingHandlerIncomingRequest, response_out handler.ExportsWasiHttp0_2_0_rc_2023_11_10_IncomingHandlerResponseOutparam) {
	hdrs := handler.NewFields()
	response := handler.NewOutgoingResponse(hdrs)
	response.SetStatusCode(200)
	body := response.Body().Unwrap()
	res_result := handler.Ok[handler.WasiHttp0_2_0_rc_2023_11_10_TypesOutgoingResponse, handler.WasiHttp0_2_0_rc_2023_11_10_TypesErrorCode](response)
	handler.StaticResponseOutparamSet(response_out, res_result)

	out := body.Write().Unwrap()
	out.BlockingWriteAndFlush([]uint8("Hello world from Go!!!\n")).Unwrap()
	handler.StaticOutgoingBodyFinish(body, handler.None[handler.WasiHttp0_2_0_rc_2023_11_10_TypesTrailers]())
}

//go:generate wit-bindgen tiny-go wit/2023_11_10 --out-dir=target_world/2023_11_10 --gofmt
func main() {}