/*
 * Nsmf_EventExposure
 *
 * Session Management Event Exposure Service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type NsmfEventExposure struct {
	Supi	string	`json:"supi,omitempty" yaml:"supi" bson:"supi" mapstructure:"Supi"`
	Gpsi	string	`json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi" mapstructure:"Gpsi"`
	// Any UE indication. This IE shall be present if the event subscription is applicable to any UE. Default value \"FALSE\" is used, if not present.
	AnyUeInd	bool	`json:"anyUeInd,omitempty" yaml:"anyUeInd" bson:"anyUeInd" mapstructure:"AnyUeInd"`
	GroupId		string	`json:"groupId,omitempty" yaml:"groupId" bson:"groupId" mapstructure:"GroupId"`
	PduSeId		int32	`json:"pduSeId,omitempty" yaml:"pduSeId" bson:"pduSeId" mapstructure:"PduSeId"`
	// Identifies an Individual SMF Notification Subscription. To enable that the value is used as part of a URI, the string shall only contain characters allowed according to the \"lower-with-hyphen\" naming convention defined in 3GPP TS 29.501 [2]. In an OpenAPI [10] schema, the format shall be designated as \"SubId\".
	SubId	string	`json:"subId,omitempty" yaml:"subId" bson:"subId" mapstructure:"SubId"`
	// Notification Correlation ID assigned by the NF service consumer.
	NotifId		string	`json:"notifId" yaml:"notifId" bson:"notifId" mapstructure:"NotifId"`
	NotifUri	string	`json:"notifUri" yaml:"notifUri" bson:"notifUri" mapstructure:"NotifUri"`
	// Alternate or backup IPv4 Addess(es) where to send Notifications.
	AltNotifIpv4Addrs	[]string	`json:"altNotifIpv4Addrs,omitempty" yaml:"altNotifIpv4Addrs" bson:"altNotifIpv4Addrs" mapstructure:"AltNotifIpv4Addrs"`
	// Alternate or backup IPv6 Addess(es) where to send Notifications.
	AltNotifIpv6Addrs	[]string	`json:"altNotifIpv6Addrs,omitempty" yaml:"altNotifIpv6Addrs" bson:"altNotifIpv6Addrs" mapstructure:"AltNotifIpv6Addrs"`
	// Subscribed events
	EventSubs	[]EventSubscription	`json:"eventSubs" yaml:"eventSubs" bson:"eventSubs" mapstructure:"EventSubs"`
	ImmeRep		bool			`json:"ImmeRep,omitempty" yaml:"ImmeRep" bson:"ImmeRep" mapstructure:"ImmeRep"`
	NotifMethod	NotificationMethod	`json:"notifMethod,omitempty" yaml:"notifMethod" bson:"notifMethod" mapstructure:"NotifMethod"`
	MaxReportNbr	int32			`json:"maxReportNbr,omitempty" yaml:"maxReportNbr" bson:"maxReportNbr" mapstructure:"MaxReportNbr"`
	Expiry		*time.Time		`json:"expiry,omitempty" yaml:"expiry" bson:"expiry" mapstructure:"Expiry"`
	RepPeriod	int32			`json:"repPeriod,omitempty" yaml:"repPeriod" bson:"repPeriod" mapstructure:"RepPeriod"`
	Guami		*Guami			`json:"guami,omitempty" yaml:"guami" bson:"guami" mapstructure:"Guami"`
	// If the NF service consumer is an AMF, it should provide the name of a service produced by the AMF that makes use of notifications about subscribed events.
	ServiveName		string	`json:"serviveName,omitempty" yaml:"serviveName" bson:"serviveName" mapstructure:"ServiveName"`
	SupportedFeatures	string	`json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures" mapstructure:"SupportedFeatures"`
}
