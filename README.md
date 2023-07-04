# Go API client for hwmux

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.19.0
- Package version: 2.19.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import hwmux "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), hwmux.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), hwmux.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), hwmux.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), hwmux.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://hwmux.silabs.net*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CallbackApi* | [**CallbackRetrieve**](docs/CallbackApi.md#callbackretrieve) | **Get** /callback/ | 
*DevicesApi* | [**DevicesCreate**](docs/DevicesApi.md#devicescreate) | **Post** /api/devices/ | 
*DevicesApi* | [**DevicesDestroy**](docs/DevicesApi.md#devicesdestroy) | **Delete** /api/devices/{id}/ | 
*DevicesApi* | [**DevicesList**](docs/DevicesApi.md#deviceslist) | **Get** /api/devices/ | 
*DevicesApi* | [**DevicesListMyList**](docs/DevicesApi.md#deviceslistmylist) | **Get** /api/devices/list_my/ | 
*DevicesApi* | [**DevicesLocationRetrieve**](docs/DevicesApi.md#deviceslocationretrieve) | **Get** /api/devices/{device_pk}/location/ | 
*DevicesApi* | [**DevicesPartialUpdate**](docs/DevicesApi.md#devicespartialupdate) | **Patch** /api/devices/{id}/ | 
*DevicesApi* | [**DevicesPermissionsPartialUpdate**](docs/DevicesApi.md#devicespermissionspartialupdate) | **Patch** /api/devices/{id}/permissions/ | 
*DevicesApi* | [**DevicesPermissionsRetrieve**](docs/DevicesApi.md#devicespermissionsretrieve) | **Get** /api/devices/{id}/permissions/ | 
*DevicesApi* | [**DevicesPermissionsUpdate**](docs/DevicesApi.md#devicespermissionsupdate) | **Put** /api/devices/{id}/permissions/ | 
*DevicesApi* | [**DevicesReleaseUpdate**](docs/DevicesApi.md#devicesreleaseupdate) | **Put** /api/devices/{id}/release/ | 
*DevicesApi* | [**DevicesReserveUpdate**](docs/DevicesApi.md#devicesreserveupdate) | **Put** /api/devices/{id}/reserve/ | 
*DevicesApi* | [**DevicesRetrieve**](docs/DevicesApi.md#devicesretrieve) | **Get** /api/devices/{id}/ | 
*DevicesApi* | [**DevicesSearchList**](docs/DevicesApi.md#devicessearchlist) | **Get** /api/devices/search/ | 
*DevicesApi* | [**DevicesSetOfflineCreate**](docs/DevicesApi.md#devicessetofflinecreate) | **Post** /api/devices/set_offline/ | 
*DevicesApi* | [**DevicesStatusCreate**](docs/DevicesApi.md#devicesstatuscreate) | **Post** /api/devices/{id}/status/ | 
*DevicesApi* | [**DevicesUpdate**](docs/DevicesApi.md#devicesupdate) | **Put** /api/devices/{id}/ | 
*GroupsApi* | [**GroupsAvailableList**](docs/GroupsApi.md#groupsavailablelist) | **Get** /api/groups/available/ | 
*GroupsApi* | [**GroupsCreate**](docs/GroupsApi.md#groupscreate) | **Post** /api/groups/ | 
*GroupsApi* | [**GroupsCreateWithDevicesCreate**](docs/GroupsApi.md#groupscreatewithdevicescreate) | **Post** /api/groups/create-with-devices | 
*GroupsApi* | [**GroupsDestroy**](docs/GroupsApi.md#groupsdestroy) | **Delete** /api/groups/{id}/ | 
*GroupsApi* | [**GroupsList**](docs/GroupsApi.md#groupslist) | **Get** /api/groups/ | 
*GroupsApi* | [**GroupsMyList**](docs/GroupsApi.md#groupsmylist) | **Get** /api/groups/my/ | 
*GroupsApi* | [**GroupsPartialUpdate**](docs/GroupsApi.md#groupspartialupdate) | **Patch** /api/groups/{id}/ | 
*GroupsApi* | [**GroupsPermissionsPartialUpdate**](docs/GroupsApi.md#groupspermissionspartialupdate) | **Patch** /api/groups/{id}/permissions/ | 
*GroupsApi* | [**GroupsPermissionsRetrieve**](docs/GroupsApi.md#groupspermissionsretrieve) | **Get** /api/groups/{id}/permissions/ | 
*GroupsApi* | [**GroupsPermissionsUpdate**](docs/GroupsApi.md#groupspermissionsupdate) | **Put** /api/groups/{id}/permissions/ | 
*GroupsApi* | [**GroupsReleaseByNameUpdate**](docs/GroupsApi.md#groupsreleasebynameupdate) | **Put** /api/groups/{group_name}/release_by_name/ | 
*GroupsApi* | [**GroupsReleaseUpdate**](docs/GroupsApi.md#groupsreleaseupdate) | **Put** /api/groups/{id}/release/ | 
*GroupsApi* | [**GroupsReserveByNameUpdate**](docs/GroupsApi.md#groupsreservebynameupdate) | **Put** /api/groups/{group_name}/reserve_by_name/ | 
*GroupsApi* | [**GroupsReserveUpdate**](docs/GroupsApi.md#groupsreserveupdate) | **Put** /api/groups/{id}/reserve/ | 
*GroupsApi* | [**GroupsRetrieve**](docs/GroupsApi.md#groupsretrieve) | **Get** /api/groups/{id}/ | 
*GroupsApi* | [**GroupsStatusCreate**](docs/GroupsApi.md#groupsstatuscreate) | **Post** /api/groups/{id}/status/ | 
*GroupsApi* | [**GroupsUpdate**](docs/GroupsApi.md#groupsupdate) | **Put** /api/groups/{id}/ | 
*LabelsApi* | [**LabelsCreate**](docs/LabelsApi.md#labelscreate) | **Post** /api/labels/ | 
*LabelsApi* | [**LabelsDestroy**](docs/LabelsApi.md#labelsdestroy) | **Delete** /api/labels/{id}/ | 
*LabelsApi* | [**LabelsList**](docs/LabelsApi.md#labelslist) | **Get** /api/labels/ | 
*LabelsApi* | [**LabelsPartialUpdate**](docs/LabelsApi.md#labelspartialupdate) | **Patch** /api/labels/{id}/ | 
*LabelsApi* | [**LabelsPermissionsPartialUpdate**](docs/LabelsApi.md#labelspermissionspartialupdate) | **Patch** /api/labels/{id}/permissions/ | 
*LabelsApi* | [**LabelsPermissionsRetrieve**](docs/LabelsApi.md#labelspermissionsretrieve) | **Get** /api/labels/{id}/permissions/ | 
*LabelsApi* | [**LabelsPermissionsUpdate**](docs/LabelsApi.md#labelspermissionsupdate) | **Put** /api/labels/{id}/permissions/ | 
*LabelsApi* | [**LabelsReserveAnyUpdate**](docs/LabelsApi.md#labelsreserveanyupdate) | **Put** /api/labels/{id}/reserve_any/ | 
*LabelsApi* | [**LabelsRetrieve**](docs/LabelsApi.md#labelsretrieve) | **Get** /api/labels/{id}/ | 
*LabelsApi* | [**LabelsUpdate**](docs/LabelsApi.md#labelsupdate) | **Put** /api/labels/{id}/ | 
*LogsApi* | [**LogsDestroy**](docs/LogsApi.md#logsdestroy) | **Delete** /api/logs/{id} | 
*LogsApi* | [**LogsList**](docs/LogsApi.md#logslist) | **Get** /api/logs/ | 
*LogsApi* | [**LogsRetrieve**](docs/LogsApi.md#logsretrieve) | **Get** /api/logs/{id}/ | 
*PartFamiliesApi* | [**PartFamiliesCreate**](docs/PartFamiliesApi.md#partfamiliescreate) | **Post** /api/part-families/ | 
*PartFamiliesApi* | [**PartFamiliesDestroy**](docs/PartFamiliesApi.md#partfamiliesdestroy) | **Delete** /api/part-families/{name}/ | 
*PartFamiliesApi* | [**PartFamiliesList**](docs/PartFamiliesApi.md#partfamilieslist) | **Get** /api/part-families/ | 
*PartFamiliesApi* | [**PartFamiliesPartialUpdate**](docs/PartFamiliesApi.md#partfamiliespartialupdate) | **Patch** /api/part-families/{name}/ | 
*PartFamiliesApi* | [**PartFamiliesPartsCreate**](docs/PartFamiliesApi.md#partfamiliespartscreate) | **Post** /api/part-families/{part_family_pk}/parts/ | 
*PartFamiliesApi* | [**PartFamiliesPartsDestroy**](docs/PartFamiliesApi.md#partfamiliespartsdestroy) | **Delete** /api/part-families/{part_family_pk}/parts/{id}/ | 
*PartFamiliesApi* | [**PartFamiliesPartsList**](docs/PartFamiliesApi.md#partfamiliespartslist) | **Get** /api/part-families/{part_family_pk}/parts/ | 
*PartFamiliesApi* | [**PartFamiliesPartsPartialUpdate**](docs/PartFamiliesApi.md#partfamiliespartspartialupdate) | **Patch** /api/part-families/{part_family_pk}/parts/{id}/ | 
*PartFamiliesApi* | [**PartFamiliesPartsRetrieve**](docs/PartFamiliesApi.md#partfamiliespartsretrieve) | **Get** /api/part-families/{part_family_pk}/parts/{id}/ | 
*PartFamiliesApi* | [**PartFamiliesPartsUpdate**](docs/PartFamiliesApi.md#partfamiliespartsupdate) | **Put** /api/part-families/{part_family_pk}/parts/{id}/ | 
*PartFamiliesApi* | [**PartFamiliesRetrieve**](docs/PartFamiliesApi.md#partfamiliesretrieve) | **Get** /api/part-families/{name}/ | 
*PartFamiliesApi* | [**PartFamiliesUpdate**](docs/PartFamiliesApi.md#partfamiliesupdate) | **Put** /api/part-families/{name}/ | 
*PartsApi* | [**PartsList**](docs/PartsApi.md#partslist) | **Get** /api/parts/ | 
*PartsApi* | [**PartsRetrieve**](docs/PartsApi.md#partsretrieve) | **Get** /api/parts/{part_no}/ | 
*PermissionsApi* | [**PermissionsGroupsCreate**](docs/PermissionsApi.md#permissionsgroupscreate) | **Post** /api/permissions/groups/ | 
*PermissionsApi* | [**PermissionsGroupsDestroy**](docs/PermissionsApi.md#permissionsgroupsdestroy) | **Delete** /api/permissions/groups/{name_or_id}/ | 
*PermissionsApi* | [**PermissionsGroupsDeviceGroupsCreate**](docs/PermissionsApi.md#permissionsgroupsdevicegroupscreate) | **Post** /api/permissions/groups/{name_or_id}/device_groups/ | 
*PermissionsApi* | [**PermissionsGroupsDeviceGroupsDestroy**](docs/PermissionsApi.md#permissionsgroupsdevicegroupsdestroy) | **Delete** /api/permissions/groups/{name_or_id}/device_groups/{id}/ | 
*PermissionsApi* | [**PermissionsGroupsDeviceGroupsList**](docs/PermissionsApi.md#permissionsgroupsdevicegroupslist) | **Get** /api/permissions/groups/{name_or_id}/device_groups/ | 
*PermissionsApi* | [**PermissionsGroupsDeviceGroupsPartialUpdate**](docs/PermissionsApi.md#permissionsgroupsdevicegroupspartialupdate) | **Patch** /api/permissions/groups/{name_or_id}/device_groups/{id}/ | 
*PermissionsApi* | [**PermissionsGroupsDeviceGroupsUpdate**](docs/PermissionsApi.md#permissionsgroupsdevicegroupsupdate) | **Put** /api/permissions/groups/{name_or_id}/device_groups/{id}/ | 
*PermissionsApi* | [**PermissionsGroupsDevicesCreate**](docs/PermissionsApi.md#permissionsgroupsdevicescreate) | **Post** /api/permissions/groups/{name_or_id}/devices/ | 
*PermissionsApi* | [**PermissionsGroupsDevicesDestroy**](docs/PermissionsApi.md#permissionsgroupsdevicesdestroy) | **Delete** /api/permissions/groups/{name_or_id}/devices/{id}/ | 
*PermissionsApi* | [**PermissionsGroupsDevicesList**](docs/PermissionsApi.md#permissionsgroupsdeviceslist) | **Get** /api/permissions/groups/{name_or_id}/devices/ | 
*PermissionsApi* | [**PermissionsGroupsDevicesPartialUpdate**](docs/PermissionsApi.md#permissionsgroupsdevicespartialupdate) | **Patch** /api/permissions/groups/{name_or_id}/devices/{id}/ | 
*PermissionsApi* | [**PermissionsGroupsDevicesUpdate**](docs/PermissionsApi.md#permissionsgroupsdevicesupdate) | **Put** /api/permissions/groups/{name_or_id}/devices/{id}/ | 
*PermissionsApi* | [**PermissionsGroupsLabelsCreate**](docs/PermissionsApi.md#permissionsgroupslabelscreate) | **Post** /api/permissions/groups/{name_or_id}/labels/ | 
*PermissionsApi* | [**PermissionsGroupsLabelsDestroy**](docs/PermissionsApi.md#permissionsgroupslabelsdestroy) | **Delete** /api/permissions/groups/{name_or_id}/labels/{id}/ | 
*PermissionsApi* | [**PermissionsGroupsLabelsList**](docs/PermissionsApi.md#permissionsgroupslabelslist) | **Get** /api/permissions/groups/{name_or_id}/labels/ | 
*PermissionsApi* | [**PermissionsGroupsLabelsPartialUpdate**](docs/PermissionsApi.md#permissionsgroupslabelspartialupdate) | **Patch** /api/permissions/groups/{name_or_id}/labels/{id}/ | 
*PermissionsApi* | [**PermissionsGroupsLabelsUpdate**](docs/PermissionsApi.md#permissionsgroupslabelsupdate) | **Put** /api/permissions/groups/{name_or_id}/labels/{id}/ | 
*PermissionsApi* | [**PermissionsGroupsList**](docs/PermissionsApi.md#permissionsgroupslist) | **Get** /api/permissions/groups/ | 
*PermissionsApi* | [**PermissionsGroupsPartialUpdate**](docs/PermissionsApi.md#permissionsgroupspartialupdate) | **Patch** /api/permissions/groups/{name_or_id}/ | 
*PermissionsApi* | [**PermissionsGroupsRetrieve**](docs/PermissionsApi.md#permissionsgroupsretrieve) | **Get** /api/permissions/groups/{name_or_id}/ | 
*PermissionsApi* | [**PermissionsGroupsUpdate**](docs/PermissionsApi.md#permissionsgroupsupdate) | **Put** /api/permissions/groups/{name_or_id}/ | 
*PermissionsApi* | [**PermissionsGroupsUsersCreate**](docs/PermissionsApi.md#permissionsgroupsuserscreate) | **Post** /api/permissions/groups/{name_or_id}/users/ | 
*PermissionsApi* | [**PermissionsGroupsUsersDestroy**](docs/PermissionsApi.md#permissionsgroupsusersdestroy) | **Delete** /api/permissions/groups/{name_or_id}/users/{username_or_id}/ | 
*ReservationsApi* | [**ReservationsActionableList**](docs/ReservationsApi.md#reservationsactionablelist) | **Get** /api/reservations/actionable/ | 
*ReservationsApi* | [**ReservationsActiveList**](docs/ReservationsApi.md#reservationsactivelist) | **Get** /api/reservations/active/ | 
*ReservationsApi* | [**ReservationsCancelUpdate**](docs/ReservationsApi.md#reservationscancelupdate) | **Put** /api/reservations/{id}/cancel/ | 
*ReservationsApi* | [**ReservationsCreate**](docs/ReservationsApi.md#reservationscreate) | **Post** /api/reservations/ | 
*ReservationsApi* | [**ReservationsList**](docs/ReservationsApi.md#reservationslist) | **Get** /api/reservations/ | 
*ReservationsApi* | [**ReservationsMetadataPartialUpdate**](docs/ReservationsApi.md#reservationsmetadatapartialupdate) | **Patch** /api/reservations/{id}/metadata/ | 
*ReservationsApi* | [**ReservationsMetadataUpdate**](docs/ReservationsApi.md#reservationsmetadataupdate) | **Put** /api/reservations/{id}/metadata/ | 
*ReservationsApi* | [**ReservationsReleaseUpdate**](docs/ReservationsApi.md#reservationsreleaseupdate) | **Put** /api/reservations/{id}/release/ | 
*ReservationsApi* | [**ReservationsRetrieve**](docs/ReservationsApi.md#reservationsretrieve) | **Get** /api/reservations/{id}/ | 
*ReservationsApi* | [**ReservationsUpdate**](docs/ReservationsApi.md#reservationsupdate) | **Put** /api/reservations/{id}/ | 
*RoomsApi* | [**RoomsList**](docs/RoomsApi.md#roomslist) | **Get** /api/rooms/ | 
*RoomsApi* | [**RoomsRetrieve**](docs/RoomsApi.md#roomsretrieve) | **Get** /api/rooms/{name}/ | 
*SchemaApi* | [**SchemaDownloadRetrieve**](docs/SchemaApi.md#schemadownloadretrieve) | **Get** /schema/download | 
*SigninApi* | [**SigninRetrieve**](docs/SigninApi.md#signinretrieve) | **Get** /signin/ | 
*SignoutApi* | [**SignoutRetrieve**](docs/SignoutApi.md#signoutretrieve) | **Get** /signout/ | 
*SitesApi* | [**SitesCreate**](docs/SitesApi.md#sitescreate) | **Post** /api/sites/ | 
*SitesApi* | [**SitesDestroy**](docs/SitesApi.md#sitesdestroy) | **Delete** /api/sites/{name}/ | 
*SitesApi* | [**SitesList**](docs/SitesApi.md#siteslist) | **Get** /api/sites/ | 
*SitesApi* | [**SitesPartialUpdate**](docs/SitesApi.md#sitespartialupdate) | **Patch** /api/sites/{name}/ | 
*SitesApi* | [**SitesRetrieve**](docs/SitesApi.md#sitesretrieve) | **Get** /api/sites/{name}/ | 
*SitesApi* | [**SitesRoomsCreate**](docs/SitesApi.md#sitesroomscreate) | **Post** /api/sites/{site_pk}/rooms/ | 
*SitesApi* | [**SitesRoomsDestroy**](docs/SitesApi.md#sitesroomsdestroy) | **Delete** /api/sites/{site_pk}/rooms/{id}/ | 
*SitesApi* | [**SitesRoomsList**](docs/SitesApi.md#sitesroomslist) | **Get** /api/sites/{site_pk}/rooms/ | 
*SitesApi* | [**SitesRoomsPartialUpdate**](docs/SitesApi.md#sitesroomspartialupdate) | **Patch** /api/sites/{site_pk}/rooms/{id}/ | 
*SitesApi* | [**SitesRoomsRetrieve**](docs/SitesApi.md#sitesroomsretrieve) | **Get** /api/sites/{site_pk}/rooms/{id}/ | 
*SitesApi* | [**SitesRoomsUpdate**](docs/SitesApi.md#sitesroomsupdate) | **Put** /api/sites/{site_pk}/rooms/{id}/ | 
*SitesApi* | [**SitesUpdate**](docs/SitesApi.md#sitesupdate) | **Put** /api/sites/{name}/ | 
*TokenAuthApi* | [**TokenAuthCreate**](docs/TokenAuthApi.md#tokenauthcreate) | **Post** /api-token-auth/ | 
*UserApi* | [**UserCreate**](docs/UserApi.md#usercreate) | **Post** /api/user/ | 
*UserApi* | [**UserCurrentRetrieve**](docs/UserApi.md#usercurrentretrieve) | **Get** /api/user/current/ | 
*UserApi* | [**UserCurrentTokenCreate**](docs/UserApi.md#usercurrenttokencreate) | **Post** /api/user/current/token/ | 
*UserApi* | [**UserCurrentTokenRetrieve**](docs/UserApi.md#usercurrenttokenretrieve) | **Get** /api/user/current/token/ | 
*UserApi* | [**UserDestroy**](docs/UserApi.md#userdestroy) | **Delete** /api/user/{username_or_id}/ | 
*UserApi* | [**UserList**](docs/UserApi.md#userlist) | **Get** /api/user/ | 
*UserApi* | [**UserPartialUpdate**](docs/UserApi.md#userpartialupdate) | **Patch** /api/user/{username_or_id}/ | 
*UserApi* | [**UserRetrieve**](docs/UserApi.md#userretrieve) | **Get** /api/user/{username_or_id}/ | 
*UserApi* | [**UserTokenCreate**](docs/UserApi.md#usertokencreate) | **Post** /api/user/{username_or_id}/token/ | 
*UserApi* | [**UserTokenRetrieve**](docs/UserApi.md#usertokenretrieve) | **Get** /api/user/{username_or_id}/token/ | 
*UserApi* | [**UserUpdate**](docs/UserApi.md#userupdate) | **Put** /api/user/{username_or_id}/ | 


## Documentation For Models

 - [AuthToken](docs/AuthToken.md)
 - [DeviceGroup](docs/DeviceGroup.md)
 - [DeviceGroupSerializerWithDevicePk](docs/DeviceGroupSerializerWithDevicePk.md)
 - [DeviceSerializerPublic](docs/DeviceSerializerPublic.md)
 - [EventEnum](docs/EventEnum.md)
 - [Label](docs/Label.md)
 - [LabelSerializerWithPermissions](docs/LabelSerializerWithPermissions.md)
 - [LightDevice](docs/LightDevice.md)
 - [Location](docs/Location.md)
 - [LocationSerializerWriteOnly](docs/LocationSerializerWriteOnly.md)
 - [Log](docs/Log.md)
 - [LoggedInUser](docs/LoggedInUser.md)
 - [NestedDeviceGroup](docs/NestedDeviceGroup.md)
 - [ObjectPermissions](docs/ObjectPermissions.md)
 - [PaginatedDeviceGroupList](docs/PaginatedDeviceGroupList.md)
 - [PaginatedDeviceSerializerPublicList](docs/PaginatedDeviceSerializerPublicList.md)
 - [PaginatedLabelList](docs/PaginatedLabelList.md)
 - [PaginatedLogList](docs/PaginatedLogList.md)
 - [PaginatedLoggedInUserList](docs/PaginatedLoggedInUserList.md)
 - [PaginatedPartFamilyList](docs/PaginatedPartFamilyList.md)
 - [PaginatedPartList](docs/PaginatedPartList.md)
 - [PaginatedPermissionGroupList](docs/PaginatedPermissionGroupList.md)
 - [PaginatedReservationSessionSerializerReadOnlyList](docs/PaginatedReservationSessionSerializerReadOnlyList.md)
 - [PaginatedRoomList](docs/PaginatedRoomList.md)
 - [PaginatedSiteList](docs/PaginatedSiteList.md)
 - [Part](docs/Part.md)
 - [PartFamily](docs/PartFamily.md)
 - [PatchedDeviceGroupSerializerWithDevicePk](docs/PatchedDeviceGroupSerializerWithDevicePk.md)
 - [PatchedLabelSerializerWithPermissions](docs/PatchedLabelSerializerWithPermissions.md)
 - [PatchedLoggedInUser](docs/PatchedLoggedInUser.md)
 - [PatchedObjectPermissions](docs/PatchedObjectPermissions.md)
 - [PatchedPart](docs/PatchedPart.md)
 - [PatchedPartFamily](docs/PatchedPartFamily.md)
 - [PatchedPermissionGroup](docs/PatchedPermissionGroup.md)
 - [PatchedResourcePermissions](docs/PatchedResourcePermissions.md)
 - [PatchedRoom](docs/PatchedRoom.md)
 - [PatchedSite](docs/PatchedSite.md)
 - [PatchedWriteOnlyDevice](docs/PatchedWriteOnlyDevice.md)
 - [PermissionGroup](docs/PermissionGroup.md)
 - [PermissionsEnum](docs/PermissionsEnum.md)
 - [ReservationRequest](docs/ReservationRequest.md)
 - [ReservationSessionSerializerReadOnly](docs/ReservationSessionSerializerReadOnly.md)
 - [ReservationSessionSerializerReadOnlyOwner](docs/ReservationSessionSerializerReadOnlyOwner.md)
 - [ResourcePermissions](docs/ResourcePermissions.md)
 - [ResourceStatusRequest](docs/ResourceStatusRequest.md)
 - [Room](docs/Room.md)
 - [Site](docs/Site.md)
 - [StatusEnum](docs/StatusEnum.md)
 - [Token](docs/Token.md)
 - [User](docs/User.md)
 - [WriteOnlyDevice](docs/WriteOnlyDevice.md)


## Documentation For Authorization



### cookieAuth

- **Type**: API key
- **API key parameter name**: sessionid
- **Location**: 

Note, each API key must be added to a map of `map[string]APIKey` where the key is: sessionid and passed in as the auth context for each request.


### tokenAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



