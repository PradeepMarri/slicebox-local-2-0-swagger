package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Query represents the Query schema from the OpenAPI specification
type Query struct {
	Queryproperties []QueryProperty `json:"queryProperties"`
	Startindex int64 `json:"startIndex"`
	Count int64 `json:"count"`
	Filters QueryFilters `json:"filters,omitempty"`
	Order QueryOrder `json:"order,omitempty"`
}

// UserPass represents the UserPass schema from the OpenAPI specification
type UserPass struct {
	Pass string `json:"pass,omitempty"`
	User string `json:"user,omitempty"`
}

// Seriesidseriestype represents the Seriesidseriestype schema from the OpenAPI specification
type Seriesidseriestype struct {
	Seriesid int64 `json:"seriesid,omitempty"`
	Seriestype Seriestype `json:"seriestype,omitempty"`
}

// TagValue represents the TagValue schema from the OpenAPI specification
type TagValue struct {
	Tagpath TagPathTag `json:"tagPath,omitempty"`
	Value string `json:"value,omitempty"`
}

// UserInfo represents the UserInfo schema from the OpenAPI specification
type UserInfo struct {
	User string `json:"user,omitempty"`
	Id int64 `json:"id,omitempty"`
	Role string `json:"role,omitempty"`
}

// OutgoingImage represents the OutgoingImage schema from the OpenAPI specification
type OutgoingImage struct {
	Sent bool `json:"sent,omitempty"`
	Sequencenumber int64 `json:"sequenceNumber,omitempty"`
	Id int64 `json:"id,omitempty"`
	Imageid int64 `json:"imageId,omitempty"`
	Outgoingtransactionid int64 `json:"outgoingTransactionId,omitempty"`
}

// WatchedDirectory represents the WatchedDirectory schema from the OpenAPI specification
type WatchedDirectory struct {
	Path string `json:"path,omitempty"`
	Id int64 `json:"id,omitempty"`
}

// QueryFilters represents the QueryFilters schema from the OpenAPI specification
type QueryFilters struct {
	Seriestagids []int64 `json:"seriesTagIds,omitempty"`
	Seriestypeids []int64 `json:"seriesTypeIds,omitempty"`
	Sourcerefs []SourceRef `json:"sourceRefs,omitempty"`
}

// Seriestype represents the Seriestype schema from the OpenAPI specification
type Seriestype struct {
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Idsquery represents the Idsquery schema from the OpenAPI specification
type Idsquery struct {
	Ids []int64 `json:"ids"`
}

// OutgoingTransaction represents the OutgoingTransaction schema from the OpenAPI specification
type OutgoingTransaction struct {
	Status string `json:"status,omitempty"`
	Totalimagecount int64 `json:"totalImageCount,omitempty"`
	Updated int64 `json:"updated,omitempty"`
	Boxid int64 `json:"boxId,omitempty"`
	Boxname string `json:"boxName,omitempty"`
	Id int64 `json:"id,omitempty"`
	Profile AnonymizationProfile `json:"profile,omitempty"`
	Sentimagecount int64 `json:"sentImageCount,omitempty"`
}

// Seriestypeupdatestatus represents the Seriestypeupdatestatus schema from the OpenAPI specification
type Seriestypeupdatestatus struct {
	Running bool `json:"running"`
}

// BulkAnonymizationData represents the BulkAnonymizationData schema from the OpenAPI specification
type BulkAnonymizationData struct {
	Profile AnonymizationProfile `json:"profile,omitempty"`
	Imagetagvaluesset []ImageTagValues `json:"imageTagValuesSet,omitempty"`
}

// ImageAttribute represents the ImageAttribute schema from the OpenAPI specification
type ImageAttribute struct {
	Vr string `json:"vr,omitempty"`
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Element string `json:"element,omitempty"`
	Group string `json:"group,omitempty"`
	Length int `json:"length,omitempty"`
	Depth int `json:"depth,omitempty"`
	Multiplicity int `json:"multiplicity,omitempty"`
	Path string `json:"path,omitempty"`
}

// ImportSession represents the ImportSession schema from the OpenAPI specification
type ImportSession struct {
	Filesrejected int `json:"filesRejected,omitempty"`
	Lastupdated int64 `json:"lastUpdated,omitempty"`
	Name string `json:"name,omitempty"`
	Filesadded int `json:"filesAdded,omitempty"`
	Filesimported int `json:"filesImported,omitempty"`
	User string `json:"user,omitempty"`
	Userid int64 `json:"userId,omitempty"`
	Created int64 `json:"created,omitempty"`
	Id int64 `json:"id,omitempty"`
}

// Seriestyperule represents the Seriestyperule schema from the OpenAPI specification
type Seriestyperule struct {
	Id int64 `json:"id,omitempty"`
	Seriestypeid int64 `json:"seriesTypeId,omitempty"`
}

// AnonymizationKeyQuery represents the AnonymizationKeyQuery schema from the OpenAPI specification
type AnonymizationKeyQuery struct {
	Order QueryOrder `json:"order,omitempty"`
	Queryproperties []QueryProperty `json:"queryProperties"`
	Startindex int64 `json:"startIndex"`
	Count int64 `json:"count"`
}

// Study represents the Study schema from the OpenAPI specification
type Study struct {
	Patientage DicomPropertyValue `json:"patientAge,omitempty"`
	Patientid int64 `json:"patientId,omitempty"`
	Studydate DicomPropertyValue `json:"studyDate,omitempty"`
	Studydescription DicomPropertyValue `json:"studyDescription,omitempty"`
	Studyid DicomPropertyValue `json:"studyID,omitempty"`
	Studyinstanceuid DicomPropertyValue `json:"studyInstanceUID,omitempty"`
	Accessionnumber DicomPropertyValue `json:"accessionNumber,omitempty"`
	Id int64 `json:"id,omitempty"`
}

// ExportSetId represents the ExportSetId schema from the OpenAPI specification
type ExportSetId struct {
	Value int64 `json:"value,omitempty"`
}

// Source represents the Source schema from the OpenAPI specification
type Source struct {
	Sourcename string `json:"sourceName,omitempty"`
	Sourcetype string `json:"sourceType,omitempty"`
	Sourceid int64 `json:"sourceId,omitempty"`
}

// Image represents the Image schema from the OpenAPI specification
type Image struct {
	Sopinstanceuid DicomPropertyValue `json:"sopInstanceUID,omitempty"`
	Id int64 `json:"id,omitempty"`
	Imagetype DicomPropertyValue `json:"imageType,omitempty"`
	Instancenumber DicomPropertyValue `json:"instanceNumber,omitempty"`
	Seriesid int64 `json:"seriesId,omitempty"`
}

// NewUser represents the NewUser schema from the OpenAPI specification
type NewUser struct {
	Password string `json:"password,omitempty"`
	Role string `json:"role,omitempty"`
	User string `json:"user,omitempty"`
}

// ImageTagValues represents the ImageTagValues schema from the OpenAPI specification
type ImageTagValues struct {
	Tagvalues []TagValue `json:"tagValues,omitempty"`
	Imageid int64 `json:"imageId,omitempty"`
}

// TagMapping represents the TagMapping schema from the OpenAPI specification
type TagMapping struct {
	Tagpath TagPathTag `json:"tagPath,omitempty"`
	Value string `json:"value,omitempty"`
}

// QueryOrder represents the QueryOrder schema from the OpenAPI specification
type QueryOrder struct {
	Orderascending bool `json:"orderAscending,omitempty"`
	Orderby string `json:"orderBy,omitempty"`
}

// AnonymizationProfile represents the AnonymizationProfile schema from the OpenAPI specification
type AnonymizationProfile struct {
	Options []ConfidentialityOption `json:"options,omitempty"`
}

// Forwardingrule represents the Forwardingrule schema from the OpenAPI specification
type Forwardingrule struct {
	Source Source `json:"source,omitempty"`
	Destination Destination `json:"destination,omitempty"`
	Id int64 `json:"id,omitempty"`
	Keepimages bool `json:"keepImages,omitempty"`
}

// Box represents the Box schema from the OpenAPI specification
type Box struct {
	Online bool `json:"online,omitempty"`
	Profile AnonymizationProfile `json:"profile,omitempty"`
	Sendmethod string `json:"sendMethod,omitempty"`
	Token string `json:"token,omitempty"`
	Baseurl string `json:"baseUrl,omitempty"`
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// TagPathTrunk represents the TagPathTrunk schema from the OpenAPI specification
type TagPathTrunk struct {
	Previous TagPathTrunk `json:"previous,omitempty"`
	Tag int `json:"tag,omitempty"`
	Item string `json:"item,omitempty"`
}

// Seriestyperuleattribute represents the Seriestyperuleattribute schema from the OpenAPI specification
type Seriestyperuleattribute struct {
	Path string `json:"path,omitempty"`
	Seriestyperuleid int64 `json:"seriesTypeRuleId"`
	Value string `json:"value"`
	Element int `json:"element"`
	Group int `json:"group"`
	Id int64 `json:"id"`
}

// Patient represents the Patient schema from the OpenAPI specification
type Patient struct {
	Id int64 `json:"id,omitempty"`
	Patientbirthdate DicomPropertyValue `json:"patientBirthDate,omitempty"`
	Patientid DicomPropertyValue `json:"patientID,omitempty"`
	Patientname DicomPropertyValue `json:"patientName,omitempty"`
	Patientsex DicomPropertyValue `json:"patientSex,omitempty"`
}

// Seriestag represents the Seriestag schema from the OpenAPI specification
type Seriestag struct {
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// SourceTagFilter represents the SourceTagFilter schema from the OpenAPI specification
type SourceTagFilter struct {
	Id int64 `json:"id,omitempty"`
	Sourceid int64 `json:"sourceId,omitempty"`
	Sourcetype string `json:"sourceType,omitempty"`
	Tagfilterid int64 `json:"tagFilterId,omitempty"`
}

// FailedOutgoingTransactionImage represents the FailedOutgoingTransactionImage schema from the OpenAPI specification
type FailedOutgoingTransactionImage struct {
	Message string `json:"message,omitempty"`
	Transactionimage OutgoingTransactionImage `json:"transactionImage,omitempty"`
}

// ConfidentialityOption represents the ConfidentialityOption schema from the OpenAPI specification
type ConfidentialityOption struct {
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	Rank int `json:"rank,omitempty"`
}

// Scp represents the Scp schema from the OpenAPI specification
type Scp struct {
	Port int `json:"port,omitempty"`
	Aetitle string `json:"aeTitle,omitempty"`
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Seriesidseriestypesresult represents the Seriesidseriestypesresult schema from the OpenAPI specification
type Seriesidseriestypesresult struct {
	Seriesidseriestypes []Seriesidseriestype `json:"seriesidseriestypes,omitempty"`
}

// Filter represents the Filter schema from the OpenAPI specification
type Filter struct {
	Tags []TagPathTag `json:"tags,omitempty"`
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tagfiltertype string `json:"tagFilterType,omitempty"`
}

// OutgoingTransactionImage represents the OutgoingTransactionImage schema from the OpenAPI specification
type OutgoingTransactionImage struct {
	Transaction OutgoingTransaction `json:"transaction,omitempty"`
	Image OutgoingImage `json:"image,omitempty"`
}

// DicomPropertyValue represents the DicomPropertyValue schema from the OpenAPI specification
type DicomPropertyValue struct {
	Value string `json:"value,omitempty"`
}

// RemoteBox represents the RemoteBox schema from the OpenAPI specification
type RemoteBox struct {
	Baseurl string `json:"baseUrl,omitempty"`
	Defaultprofile AnonymizationProfile `json:"defaultProfile,omitempty"`
	Name string `json:"name,omitempty"`
}

// AnonymizationKey represents the AnonymizationKey schema from the OpenAPI specification
type AnonymizationKey struct {
	Imageid int64 `json:"imageId,omitempty"`
	Anonsopinstanceuid string `json:"anonSOPInstanceUID,omitempty"`
	Seriesinstanceuid string `json:"seriesInstanceUID,omitempty"`
	Studyinstanceuid string `json:"studyInstanceUID,omitempty"`
	Anonpatientid string `json:"anonPatientID,omitempty"`
	Anonpatientname string `json:"anonPatientName,omitempty"`
	Anonseriesinstanceuid string `json:"anonSeriesInstanceUID,omitempty"`
	Anonstudyinstanceuid string `json:"anonStudyInstanceUID,omitempty"`
	Patientid string `json:"patientID,omitempty"`
	Patientname string `json:"patientName,omitempty"`
	Sopinstanceuid string `json:"sopInstanceUID,omitempty"`
	Created int64 `json:"created,omitempty"`
	Id int64 `json:"id,omitempty"`
}

// ImageInformation represents the ImageInformation schema from the OpenAPI specification
type ImageInformation struct {
	Numberofframes int `json:"numberOfFrames,omitempty"`
	Frameindex int `json:"frameIndex,omitempty"`
	Maximumpixelvalue int `json:"maximumPixelValue,omitempty"`
	Minimumpixelvalue int `json:"minimumPixelValue,omitempty"`
}

// LogEntry represents the LogEntry schema from the OpenAPI specification
type LogEntry struct {
	Created int64 `json:"created,omitempty"`
	Entrytype string `json:"entryType,omitempty"`
	Id int64 `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	Subject string `json:"subject,omitempty"`
}

// User represents the User schema from the OpenAPI specification
type User struct {
	Hashedpassword string `json:"hashedPassword,omitempty"`
	Id int64 `json:"id"`
	Role string `json:"role"`
	User string `json:"user"`
}

// IncomingTransaction represents the IncomingTransaction schema from the OpenAPI specification
type IncomingTransaction struct {
	Boxname string `json:"boxName,omitempty"`
	Id int64 `json:"id,omitempty"`
	Outgoingtransactionid int64 `json:"outgoingTransactionId,omitempty"`
	Receivedimagecount int64 `json:"receivedImageCount,omitempty"`
	Status string `json:"status,omitempty"`
	Totalimagecount int64 `json:"totalImageCount,omitempty"`
	Updated int64 `json:"updated,omitempty"`
	Boxid int64 `json:"boxId,omitempty"`
}

// RemoteBoxConnectionData represents the RemoteBoxConnectionData schema from the OpenAPI specification
type RemoteBoxConnectionData struct {
	Defaultprofile AnonymizationProfile `json:"defaultProfile,omitempty"`
	Name string `json:"name,omitempty"`
}

// TagPathTag represents the TagPathTag schema from the OpenAPI specification
type TagPathTag struct {
	Tag int `json:"tag,omitempty"`
	Previous TagPathTrunk `json:"previous,omitempty"`
}

// AnonymizationData represents the AnonymizationData schema from the OpenAPI specification
type AnonymizationData struct {
	Profile AnonymizationProfile `json:"profile,omitempty"`
	Tagvalues []TagValue `json:"tagValues,omitempty"`
}

// Scu represents the Scu schema from the OpenAPI specification
type Scu struct {
	Name string `json:"name,omitempty"`
	Port int `json:"port,omitempty"`
	Aetitle string `json:"aeTitle,omitempty"`
	Host string `json:"host,omitempty"`
	Id int64 `json:"id,omitempty"`
}

// Series represents the Series schema from the OpenAPI specification
type Series struct {
	Id int64 `json:"id,omitempty"`
	Manufacturer DicomPropertyValue `json:"manufacturer,omitempty"`
	Seriesdate DicomPropertyValue `json:"seriesDate,omitempty"`
	Stationname DicomPropertyValue `json:"stationName,omitempty"`
	Protocolname DicomPropertyValue `json:"protocolName,omitempty"`
	Seriesdescription DicomPropertyValue `json:"seriesDescription,omitempty"`
	Seriesinstanceuid DicomPropertyValue `json:"seriesInstanceUID,omitempty"`
	Studyid int64 `json:"studyId,omitempty"`
	Frameofreferenceuid DicomPropertyValue `json:"frameOfReferenceUID,omitempty"`
	Modality DicomPropertyValue `json:"modality,omitempty"`
	Bodypartexamined DicomPropertyValue `json:"bodyPartExamined,omitempty"`
}

// FlatSeries represents the FlatSeries schema from the OpenAPI specification
type FlatSeries struct {
	Study Study `json:"study,omitempty"`
	Id int64 `json:"id,omitempty"`
	Patient Patient `json:"patient,omitempty"`
	Series Series `json:"series,omitempty"`
}

// SourceRef represents the SourceRef schema from the OpenAPI specification
type SourceRef struct {
	Sourcetype string `json:"sourceType,omitempty"`
	Sourceid int64 `json:"sourceId,omitempty"`
}

// Destination represents the Destination schema from the OpenAPI specification
type Destination struct {
	Destinationid int64 `json:"destinationId,omitempty"`
	Destinationname string `json:"destinationName,omitempty"`
	Destinationtype string `json:"destinationType,omitempty"`
}

// QueryProperty represents the QueryProperty schema from the OpenAPI specification
type QueryProperty struct {
	Operator string `json:"operator,omitempty"`
	Propertyname string `json:"propertyName,omitempty"`
	Propertyvalue string `json:"propertyValue,omitempty"`
}

// AnonymizationKeyValue represents the AnonymizationKeyValue schema from the OpenAPI specification
type AnonymizationKeyValue struct {
	Value string `json:"value,omitempty"`
	Anonymizationkeyid int64 `json:"anonymizationKeyId,omitempty"`
	Anonymizedvalue string `json:"anonymizedValue,omitempty"`
	Id int64 `json:"id,omitempty"`
	Tagpath TagPathTag `json:"tagPath,omitempty"`
}
