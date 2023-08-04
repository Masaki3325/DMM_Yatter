package accounts

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
)

// Request body for `POST /v1/accounts`
type UpdateRequest struct {
	DisplayName *string
	Note        *string
	Avatar      *string
	Header      *string
}

// Handle request for `POST /v1/accounts/update_credentials`
func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	// var req UpdateRequest
	account := auth.AccountOf(r)
	// フォームからの値の取得
	displayName := r.FormValue("display_name")
	note := r.FormValue("note")

	// アバターの処理
	avatarURL, err := processFile(r, "avatar", account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ヘッダーの処理
	headerURL, err := processFile(r, "header", account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account.DisplayName = &displayName
	account.Note = &note
	account.Avatar = &avatarURL
	account.Header = &headerURL

	err = h.ar.UpdateAccount(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	addedAccount, err := h.ar.FindByUsername(r.Context(), account.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(addedAccount); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func processFile(r *http.Request, key string, account *object.Account) (string, error) {
	file, header, err := r.FormFile(key)
	if err == http.ErrMissingFile {
		return "", nil
	} else if err != nil {
		return "", err
	}
	defer file.Close()

	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		return "", err
	}

	// 元のファイル名から拡張子を取得
	ext := filepath.Ext(header.Filename)

	// ユーザー名とキー（avatarまたはheader）を使用して一意のファイル名を生成し、拡張子を追加
	tempFileName := filepath.Join("uploads", account.Username+"_"+key+ext)

	// ファイルを開く
	tempFile, err := os.OpenFile(tempFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	// ファイルをコピー
	if _, err := io.Copy(tempFile, file); err != nil {
		return "", err
	}

	// サーバーのURLを指定
	url := tempFileName

	return url, nil
}
