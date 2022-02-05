// SPDX-License-Identifier: Apache-2.0

package api

type ResponseErrorData struct {
	code string
	message string
}

type ResponseError struct {
	error ResponseErrorData
}

type ResponseFileData struct {
	key string
	lastModified string
}

type ResponseFile struct {
	data ResponseFileData
}
