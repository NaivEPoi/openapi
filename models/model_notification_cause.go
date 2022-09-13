/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type NotificationCause string

// List of NotificationCause
const (
	NotificationCause_QOS_FULFILLED		NotificationCause	= "QOS_FULFILLED"
	NotificationCause_QOS_NOT_FULFILLED	NotificationCause	= "QOS_NOT_FULFILLED"
	NotificationCause_UP_SEC_FULFILLED	NotificationCause	= "UP_SEC_FULFILLED"
	NotificationCause_UP_SEC_NOT_FULFILLED	NotificationCause	= "UP_SEC_NOT_FULFILLED"
)
