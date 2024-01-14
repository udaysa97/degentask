package fetchuserbalance

import (
	"fmt"
	"net/http"
	"text/template"
	"txcrawlerdegen/api/common"
	"txcrawlerdegen/api/types"
	"txcrawlerdegen/constants"
	userbalance "txcrawlerdegen/internal/service/fetchuserbalanceservice"

	"github.com/gin-gonic/gin"
)

func FetchUserBalance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestBody := &types.FetchBalanceRequest{}
		response := types.ResponseDTO[types.FetchBalanceResponse]{}
		requestBody.WalletAddress = "0xDa9CE944a37d218c3302F6B82a094844C6ECEb17"
		// Error if validation fails
		if err := common.ReadAndValidateRequestBody(ctx.Request, requestBody); err != nil {
			response.Status = types.StatusError
			errorResp := types.ErrorResponse{}
			errorResp.Code = http.StatusBadRequest
			errorResp.ErrorCode = constants.ERROR_TYPES[constants.BAD_REUQEST_ERROR].ErrorCode
			errorResp.Message = fmt.Sprintf("Validation: %s", err)
			response.Error = &errorResp
			response.Success = false
			ctx.JSON(http.StatusBadRequest, response)
			return

		}
		// Check if balance is returned or error
		balance, err := userbalance.InitUserBalanceSvc().FetchUserBalance(ctx, requestBody.WalletAddress, requestBody.TokenSymbol)
		if err != nil {
			response.Status = types.StatusError
			errorResp := types.ErrorResponse{}
			errorResp.Code = http.StatusInternalServerError
			errorResp.ErrorCode = constants.ERROR_TYPES[constants.BAD_REUQEST_ERROR].ErrorCode
			errorResp.Message = err.Error()
			response.Error = &errorResp
			response.Success = false
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		response.Status = types.StatusSuccess
		successRespose := types.FetchBalanceResponse{}
		successRespose.Balance = balance
		response.Result = &successRespose
		response.Success = true
		ctx.JSON(http.StatusOK, response)
		return
	}
}

func FetchEthTxs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestBody := &types.FetchBalanceRequest{}
		response := types.ResponseDTO[types.FetchBalanceResponse]{}
		requestBody.WalletAddress = "0xDa9CE944a37d218c3302F6B82a094844C6ECEb17"
		// Error if validation fails
		if err := common.ReadAndValidateRequestBody(ctx.Request, requestBody); err != nil {
			response.Status = types.StatusError
			errorResp := types.ErrorResponse{}
			errorResp.Code = http.StatusBadRequest
			errorResp.ErrorCode = constants.ERROR_TYPES[constants.BAD_REUQEST_ERROR].ErrorCode
			errorResp.Message = fmt.Sprintf("Validation: %s", err)
			response.Error = &errorResp
			response.Success = false
			ctx.JSON(http.StatusBadRequest, response)
			return

		}
		// Check if balance is returned or error
		userbalance.InitUserBalanceSvc().FetchWalletTxsEth(ctx, requestBody.WalletAddress)

		return
	}
}

func renderHTMLTemplate(w http.ResponseWriter, tmpl string, data types.BalanceData) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
