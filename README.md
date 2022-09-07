


# Future Vuls Public API
Future Vuls Public API
  

## Informations

### Version

v1

### Contact

FutureVuls  https://futurevuls.tayori.com/form/1dd23480d6f0f7a0dd066f03be8ed7d1b41d10e7

## Content negotiation

### URI Schemes
  * http
  * https

### Consumes
  * application/gob
  * application/json
  * application/xml

### Produces
  * application/gob
  * application/json
  * application/xml

## Access control

### Security Schemes

#### api_key_header_Authorization (header: Authorization)

Secures endpoint by requiring an API key.

> **Type**: apikey

## All endpoints

###  cve

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /v1/cve/{cveID} | [cve get cve detail](#cve-get-cve-detail) | getCveDetail cve |
| GET | /v1/cves | [cve get cve list](#cve-get-cve-list) | getCveList cve |
  


###  health

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /healths | [health health](#health-health) | health health |
  


###  lockfile

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v1/lockfile | [lockfile add lockfile](#lockfile-add-lockfile) | addLockfile lockfile |
| DELETE | /v1/lockfile/{lockfileID} | [lockfile delete lockfile](#lockfile-delete-lockfile) | deleteLockfile lockfile |
| GET | /v1/lockfile/{lockfileID} | [lockfile get lockfile detail](#lockfile-get-lockfile-detail) | getLockfileDetail lockfile |
| GET | /v1/lockfiles | [lockfile get lockfile list](#lockfile-get-lockfile-list) | getLockfileList lockfile |
| PUT | /v1/lockfile/{lockfileID} | [lockfile update lockfile](#lockfile-update-lockfile) | updateLockfile lockfile |
  


###  pkg_cpe

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v1/pkgCpe/cpe | [pkg cpe add cpe](#pkg-cpe-add-cpe) | addCpe pkgCpe |
| DELETE | /v1/pkgCpe/cpe/{cpeID} | [pkg cpe delete cpe](#pkg-cpe-delete-cpe) | deleteCpe pkgCpe |
| DELETE | /v1/pkgCpe/cpe | [pkg cpe delete cpe deprecated](#pkg-cpe-delete-cpe-deprecated) | deleteCpe_deprecated pkgCpe |
| GET | /v1/pkgCpe/cpe/{cpeID} | [pkg cpe get cpe detail](#pkg-cpe-get-cpe-detail) | getCpeDetail pkgCpe |
| GET | /v1/pkgCpes | [pkg cpe get pkg cpe list](#pkg-cpe-get-pkg-cpe-list) | getPkgCpeList pkgCpe |
| GET | /v1/pkgCpe/pkg/{pkgID} | [pkg cpe get pkg detail](#pkg-cpe-get-pkg-detail) | getPkgDetail pkgCpe |
  


###  role

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| DELETE | /v1/role/{roleID} | [role delete role](#role-delete-role) | deleteRole role |
| GET | /v1/role/{roleID} | [role get role detail](#role-get-role-detail) | getRoleDetail role |
| GET | /v1/roles | [role get role list](#role-get-role-list) | getRoleList role |
| PUT | /v1/role/{roleID} | [role update role](#role-update-role) | updateRole role |
  


###  serverops

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v1/server/paste | [server create pkg paste server](#server-create-pkg-paste-server) | createPkgPasteServer server |
| DELETE | /v1/server/{serverID} | [server delete server](#server-delete-server) | deleteServer server |
| GET | /v1/server/{serverID} | [server get server detail](#server-get-server-detail) | getServerDetail server |
| GET | /v1/server/uuid/{serverUuid} | [server get server detail by UUID](#server-get-server-detail-by-uuid) | getServerDetailByUUID server |
| GET | /v1/servers | [server get server list](#server-get-server-list) | getServerList server |
| PUT | /v1/server/paste/{serverID} | [server update pkg paste server](#server-update-pkg-paste-server) | updatePkgPasteServer server |
| PUT | /v1/server/{serverID} | [server update server](#server-update-server) | updateServer server |
  


###  task

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v1/task/{taskID}/comment | [task add task comment](#task-add-task-comment) | addTaskComment task |
| GET | /v1/task/{taskID} | [task get task detail](#task-get-task-detail) | getTaskDetail task |
| GET | /v1/tasks | [task get task list](#task-get-task-list) | getTaskList task |
| PUT | /v1/task/{taskID} | [task update task](#task-update-task) | updateTask task |
| PUT | /v1/task/{taskID}/ignore | [task update task ignore](#task-update-task-ignore) | updateTaskIgnore task |
  


## Paths

### <span id="cve-get-cve-detail"></span> getCveDetail cve (*cve#getCveDetail*)

```
GET /v1/cve/{cveID}
```

cve detail

> [Read more](https://vuls.biz/ "Get Cve Detail")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cveID | `path` | string | `string` |  | ✓ |  | Cve ID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#cve-get-cve-detail-200) | OK | OK response. |  | [schema](#cve-get-cve-detail-200-schema) |

#### Responses


##### <span id="cve-get-cve-detail-200"></span> 200 - OK response.
Status: OK

###### <span id="cve-get-cve-detail-200-schema"></span> Schema
   
  

[CveGetCveDetailResponseBody](#cve-get-cve-detail-response-body)

### <span id="cve-get-cve-list"></span> getCveList cve (*cve#getCveList*)

```
GET /v1/cves
```

cve list

> [Read more](https://vuls.biz/ "Get Cve List")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Authorization | `header` | string | `string` |  |  |  | api key auth |
| filterCpeID | `query` | integer | `int64` |  |  |  | CpeID filter |
| filterPkgID | `query` | integer | `int64` |  |  |  | PackageID filter |
| filterRoleID | `query` | integer | `int64` |  |  |  | ServerRoleID filter |
| filterServerID | `query` | integer | `int64` |  |  |  | ServerID filter |
| limit | `query` | integer | `int64` |  |  | `20` | Limit |
| offset | `query` | integer | `int64` |  |  |  | Offset |
| page | `query` | integer | `int64` |  |  | `1` | Page Number |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#cve-get-cve-list-200) | OK | OK response. |  | [schema](#cve-get-cve-list-200-schema) |

#### Responses


##### <span id="cve-get-cve-list-200"></span> 200 - OK response.
Status: OK

###### <span id="cve-get-cve-list-200-schema"></span> Schema
   
  

[CveGetCveListResponseBody](#cve-get-cve-list-response-body)

### <span id="health-health"></span> health health (*health#health*)

```
GET /healths
```

health

> [Read more](https://vuls.biz/ "Health")

#### URI Schemes
  * http

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#health-health-200) | OK | OK response. |  | [schema](#health-health-200-schema) |

#### Responses


##### <span id="health-health-200"></span> 200 - OK response.
Status: OK

###### <span id="health-health-200-schema"></span> Schema
   
  



### <span id="lockfile-add-lockfile"></span> addLockfile lockfile (*lockfile#addLockfile*)

```
POST /v1/lockfile
```

add lockfile

> [Read more](https://vuls.biz/ "Add Lockfile")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Authorization | `header` | string | `string` |  |  |  | api key auth |
| AddLockfileRequestBody | `body` | [LockfileAddLockfileRequestBody](#lockfile-add-lockfile-request-body) | `models.LockfileAddLockfileRequestBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#lockfile-add-lockfile-200) | OK | OK response. |  | [schema](#lockfile-add-lockfile-200-schema) |

#### Responses


##### <span id="lockfile-add-lockfile-200"></span> 200 - OK response.
Status: OK

###### <span id="lockfile-add-lockfile-200-schema"></span> Schema
   
  

[LockfileAddLockfileResponseBody](#lockfile-add-lockfile-response-body)

### <span id="lockfile-delete-lockfile"></span> deleteLockfile lockfile (*lockfile#deleteLockfile*)

```
DELETE /v1/lockfile/{lockfileID}
```

lockfile delete

> [Read more](https://vuls.biz/ "Delete Lockfile")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| lockfileID | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | Lockfile ID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#lockfile-delete-lockfile-200) | OK | OK response. |  | [schema](#lockfile-delete-lockfile-200-schema) |

#### Responses


##### <span id="lockfile-delete-lockfile-200"></span> 200 - OK response.
Status: OK

###### <span id="lockfile-delete-lockfile-200-schema"></span> Schema

### <span id="lockfile-get-lockfile-detail"></span> getLockfileDetail lockfile (*lockfile#getLockfileDetail*)

```
GET /v1/lockfile/{lockfileID}
```

lockfile detail

> [Read more](https://vuls.biz/ "Get Lockfile Detail")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| lockfileID | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | Lockfile ID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#lockfile-get-lockfile-detail-200) | OK | OK response. |  | [schema](#lockfile-get-lockfile-detail-200-schema) |

#### Responses


##### <span id="lockfile-get-lockfile-detail-200"></span> 200 - OK response.
Status: OK

###### <span id="lockfile-get-lockfile-detail-200-schema"></span> Schema
   
  

[LockfileGetLockfileDetailResponseBody](#lockfile-get-lockfile-detail-response-body)

### <span id="lockfile-get-lockfile-list"></span> getLockfileList lockfile (*lockfile#getLockfileList*)

```
GET /v1/lockfiles
```

lockfile list

> [Read more](https://vuls.biz/ "Get Lockfile List")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Authorization | `header` | string | `string` |  |  |  | api key auth |
| filterPath | `query` | string | `string` |  |  |  | Path filter |
| filterServerID | `query` | int64 (formatted integer) | `int64` |  |  |  | ServerID filter |
| limit | `query` | integer | `int64` |  |  | `20` | Limit |
| offset | `query` | integer | `int64` |  |  |  | Offset |
| page | `query` | integer | `int64` |  |  | `1` | Page Number |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#lockfile-get-lockfile-list-200) | OK | OK response. |  | [schema](#lockfile-get-lockfile-list-200-schema) |

#### Responses


##### <span id="lockfile-get-lockfile-list-200"></span> 200 - OK response.
Status: OK

###### <span id="lockfile-get-lockfile-list-200-schema"></span> Schema
   
  

[LockfileGetLockfileListResponseBody](#lockfile-get-lockfile-list-response-body)

### <span id="lockfile-update-lockfile"></span> updateLockfile lockfile (*lockfile#updateLockfile*)

```
PUT /v1/lockfile/{lockfileID}
```

update lockfile

> [Read more](https://vuls.biz/ "Update Lockfile")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| lockfileID | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | Lockfile ID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |
| UpdateLockfileRequestBody | `body` | [LockfileUpdateLockfileRequestBody](#lockfile-update-lockfile-request-body) | `models.LockfileUpdateLockfileRequestBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#lockfile-update-lockfile-200) | OK | OK response. |  | [schema](#lockfile-update-lockfile-200-schema) |

#### Responses


##### <span id="lockfile-update-lockfile-200"></span> 200 - OK response.
Status: OK

###### <span id="lockfile-update-lockfile-200-schema"></span> Schema
   
  

[LockfileUpdateLockfileResponseBody](#lockfile-update-lockfile-response-body)

### <span id="pkg-cpe-add-cpe"></span> addCpe pkgCpe (*pkgCpe#addCpe*)

```
POST /v1/pkgCpe/cpe
```

add cpe

> [Read more](https://vuls.biz/ "Add cpe")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Authorization | `header` | string | `string` |  |  |  | api key auth |
| AddCpeRequestBody | `body` | [PkgCpeAddCpeRequestBody](#pkg-cpe-add-cpe-request-body) | `models.PkgCpeAddCpeRequestBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#pkg-cpe-add-cpe-200) | OK | OK response. |  | [schema](#pkg-cpe-add-cpe-200-schema) |

#### Responses


##### <span id="pkg-cpe-add-cpe-200"></span> 200 - OK response.
Status: OK

###### <span id="pkg-cpe-add-cpe-200-schema"></span> Schema
   
  

[PkgCpeAddCpeResponseBody](#pkg-cpe-add-cpe-response-body)

### <span id="pkg-cpe-delete-cpe"></span> deleteCpe pkgCpe (*pkgCpe#deleteCpe*)

```
DELETE /v1/pkgCpe/cpe/{cpeID}
```

delete cpe (urlにcpeIDを指定してください。cpeIDの指定のないアクセス方法は今後削除されます。)

> [Read more](https://vuls.biz/ "Delete cpe Detail")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cpeID | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | cpe ID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#pkg-cpe-delete-cpe-200) | OK | OK response. |  | [schema](#pkg-cpe-delete-cpe-200-schema) |

#### Responses


##### <span id="pkg-cpe-delete-cpe-200"></span> 200 - OK response.
Status: OK

###### <span id="pkg-cpe-delete-cpe-200-schema"></span> Schema

### <span id="pkg-cpe-delete-cpe-deprecated"></span> deleteCpe_deprecated pkgCpe (*pkgCpe#deleteCpe_deprecated*)

```
DELETE /v1/pkgCpe/cpe
```

[deprecated] urlにcpeIDを指定して利用してください。cpeIDの指定のないこちらのアクセス方法は今後削除されます。

> [Read more](https://vuls.biz/ "Delete cpe Detail")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Authorization | `header` | string | `string` |  |  |  | api key auth |
| deleteCpe_deprecated_request_body | `body` | [PkgCpeDeleteCpeDeprecatedRequestBody](#pkg-cpe-delete-cpe-deprecated-request-body) | `models.PkgCpeDeleteCpeDeprecatedRequestBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#pkg-cpe-delete-cpe-deprecated-200) | OK | OK response. |  | [schema](#pkg-cpe-delete-cpe-deprecated-200-schema) |

#### Responses


##### <span id="pkg-cpe-delete-cpe-deprecated-200"></span> 200 - OK response.
Status: OK

###### <span id="pkg-cpe-delete-cpe-deprecated-200-schema"></span> Schema

### <span id="pkg-cpe-get-cpe-detail"></span> getCpeDetail pkgCpe (*pkgCpe#getCpeDetail*)

```
GET /v1/pkgCpe/cpe/{cpeID}
```

cpe detail

> [Read more](https://vuls.biz/ "Get cpe Detail")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cpeID | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | cpe ID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#pkg-cpe-get-cpe-detail-200) | OK | OK response. |  | [schema](#pkg-cpe-get-cpe-detail-200-schema) |

#### Responses


##### <span id="pkg-cpe-get-cpe-detail-200"></span> 200 - OK response.
Status: OK

###### <span id="pkg-cpe-get-cpe-detail-200-schema"></span> Schema
   
  

[PkgCpeGetCpeDetailResponseBody](#pkg-cpe-get-cpe-detail-response-body)

### <span id="pkg-cpe-get-pkg-cpe-list"></span> getPkgCpeList pkgCpe (*pkgCpe#getPkgCpeList*)

```
GET /v1/pkgCpes
```

pkgCpe list

> [Read more](https://vuls.biz/ "Get PkgCpe List. Calling pkgCpes without filters can be very time-consuming. It is recommended that you specify filter as much as possible.")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Authorization | `header` | string | `string` |  |  |  | api key auth |
| filterCveID | `query` | string | `string` |  |  |  | CveID filter |
| filterRoleID | `query` | integer | `int64` |  |  |  | ServerRoleID filter |
| filterServerID | `query` | integer | `int64` |  |  |  | ServerID filter |
| filterTaskID | `query` | integer | `int64` |  |  |  | TaskID filter |
| limit | `query` | integer | `int64` |  |  | `20` | Limit |
| offset | `query` | integer | `int64` |  |  |  | Offset |
| page | `query` | integer | `int64` |  |  | `1` | Page Number |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#pkg-cpe-get-pkg-cpe-list-200) | OK | OK response. |  | [schema](#pkg-cpe-get-pkg-cpe-list-200-schema) |

#### Responses


##### <span id="pkg-cpe-get-pkg-cpe-list-200"></span> 200 - OK response.
Status: OK

###### <span id="pkg-cpe-get-pkg-cpe-list-200-schema"></span> Schema
   
  

[PkgCpeGetPkgCpeListResponseBody](#pkg-cpe-get-pkg-cpe-list-response-body)

### <span id="pkg-cpe-get-pkg-detail"></span> getPkgDetail pkgCpe (*pkgCpe#getPkgDetail*)

```
GET /v1/pkgCpe/pkg/{pkgID}
```

pkg detail

> [Read more](https://vuls.biz/ "Get Pkg Detail")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| pkgID | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | PackageID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#pkg-cpe-get-pkg-detail-200) | OK | OK response. |  | [schema](#pkg-cpe-get-pkg-detail-200-schema) |

#### Responses


##### <span id="pkg-cpe-get-pkg-detail-200"></span> 200 - OK response.
Status: OK

###### <span id="pkg-cpe-get-pkg-detail-200-schema"></span> Schema
   
  

[PkgCpeGetPkgDetailResponseBody](#pkg-cpe-get-pkg-detail-response-body)

### <span id="role-delete-role"></span> deleteRole role (*role#deleteRole*)

```
DELETE /v1/role/{roleID}
```

role delete

> [Read more](https://vuls.biz/ "Delete Role")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| roleID | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | Role ID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#role-delete-role-200) | OK | OK response. |  | [schema](#role-delete-role-200-schema) |

#### Responses


##### <span id="role-delete-role-200"></span> 200 - OK response.
Status: OK

###### <span id="role-delete-role-200-schema"></span> Schema

### <span id="role-get-role-detail"></span> getRoleDetail role (*role#getRoleDetail*)

```
GET /v1/role/{roleID}
```

role detail

> [Read more](https://vuls.biz/ "Get Role Detail")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| roleID | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | Role ID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#role-get-role-detail-200) | OK | OK response. |  | [schema](#role-get-role-detail-200-schema) |

#### Responses


##### <span id="role-get-role-detail-200"></span> 200 - OK response.
Status: OK

###### <span id="role-get-role-detail-200-schema"></span> Schema
   
  

[RoleGetRoleDetailResponseBody](#role-get-role-detail-response-body)

### <span id="role-get-role-list"></span> getRoleList role (*role#getRoleList*)

```
GET /v1/roles
```

role list

> [Read more](https://vuls.biz/ "Get Role List")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Authorization | `header` | string | `string` |  |  |  | api key auth |
| filterCveID | `query` | string | `string` |  |  |  | CveID filter |
| limit | `query` | integer | `int64` |  |  | `20` | Limit (default: 20, max: 100) |
| offset | `query` | integer | `int64` |  |  |  | Offset |
| page | `query` | integer | `int64` |  |  | `1` | Page Number (default: 1) |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#role-get-role-list-200) | OK | OK response. |  | [schema](#role-get-role-list-200-schema) |

#### Responses


##### <span id="role-get-role-list-200"></span> 200 - OK response.
Status: OK

###### <span id="role-get-role-list-200-schema"></span> Schema
   
  

[RoleGetRoleListResponseBody](#role-get-role-list-response-body)

### <span id="role-update-role"></span> updateRole role (*role#updateRole*)

```
PUT /v1/role/{roleID}
```

update role

> [Read more](https://vuls.biz/ "Update Role")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| roleID | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | Role ID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |
| UpdateRoleRequestBody | `body` | [RoleUpdateRoleRequestBody](#role-update-role-request-body) | `models.RoleUpdateRoleRequestBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#role-update-role-200) | OK | OK response. |  | [schema](#role-update-role-200-schema) |

#### Responses


##### <span id="role-update-role-200"></span> 200 - OK response.
Status: OK

###### <span id="role-update-role-200-schema"></span> Schema
   
  

[RoleUpdateRoleResponseBody](#role-update-role-response-body)

### <span id="server-create-pkg-paste-server"></span> createPkgPasteServer server (*server#createPkgPasteServer*)

```
POST /v1/server/paste
```

create paste server

> [Read more](https://vuls.biz/ "Create Paste Server")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Authorization | `header` | string | `string` |  |  |  | api key auth |
| CreatePkgPasteServerRequestBody | `body` | [ServerCreatePkgPasteServerRequestBody](#server-create-pkg-paste-server-request-body) | `models.ServerCreatePkgPasteServerRequestBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#server-create-pkg-paste-server-200) | OK | OK response. |  | [schema](#server-create-pkg-paste-server-200-schema) |

#### Responses


##### <span id="server-create-pkg-paste-server-200"></span> 200 - OK response.
Status: OK

###### <span id="server-create-pkg-paste-server-200-schema"></span> Schema
   
  

[ServerCreatePkgPasteServerResponseBody](#server-create-pkg-paste-server-response-body)

### <span id="server-delete-server"></span> deleteServer server (*server#deleteServer*)

```
DELETE /v1/server/{serverID}
```

server delete

> [Read more](https://vuls.biz/ "Delete Server")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| serverID | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | Server ID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#server-delete-server-200) | OK | OK response. |  | [schema](#server-delete-server-200-schema) |

#### Responses


##### <span id="server-delete-server-200"></span> 200 - OK response.
Status: OK

###### <span id="server-delete-server-200-schema"></span> Schema

### <span id="server-get-server-detail"></span> getServerDetail server (*server#getServerDetail*)

```
GET /v1/server/{serverID}
```

server detail

> [Read more](https://vuls.biz/ "Get Server Detail")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| serverID | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | Server ID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#server-get-server-detail-200) | OK | OK response. |  | [schema](#server-get-server-detail-200-schema) |

#### Responses


##### <span id="server-get-server-detail-200"></span> 200 - OK response.
Status: OK

###### <span id="server-get-server-detail-200-schema"></span> Schema
   
  

[ServerGetServerDetailResponseBody](#server-get-server-detail-response-body)

### <span id="server-get-server-detail-by-uuid"></span> getServerDetailByUUID server (*server#getServerDetailByUUID*)

```
GET /v1/server/uuid/{serverUuid}
```

server detail by UUID

> [Read more](https://vuls.biz/ "Get Server Detail")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| serverUuid | `path` | string | `string` |  | ✓ |  | Server UUID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#server-get-server-detail-by-uuid-200) | OK | OK response. |  | [schema](#server-get-server-detail-by-uuid-200-schema) |

#### Responses


##### <span id="server-get-server-detail-by-uuid-200"></span> 200 - OK response.
Status: OK

###### <span id="server-get-server-detail-by-uuid-200-schema"></span> Schema
   
  

[ServerGetServerDetailByUUIDResponseBody](#server-get-server-detail-by-uuid-response-body)

### <span id="server-get-server-list"></span> getServerList server (*server#getServerList*)

```
GET /v1/servers
```

server list

> [Read more](https://vuls.biz/ "Get Server List")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Authorization | `header` | string | `string` |  |  |  | api key auth |
| filterCveID | `query` | string | `string` |  |  |  | CveID filter |
| filterRoleID | `query` | integer | `int64` |  |  |  | ServerRoleID filter |
| filterTagName | `query` | string | `string` |  |  |  | ServerTagName filter |
| limit | `query` | integer | `int64` |  |  | `20` | Limit |
| offset | `query` | integer | `int64` |  |  |  | Offset |
| page | `query` | integer | `int64` |  |  | `1` | Page Number |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#server-get-server-list-200) | OK | OK response. |  | [schema](#server-get-server-list-200-schema) |

#### Responses


##### <span id="server-get-server-list-200"></span> 200 - OK response.
Status: OK

###### <span id="server-get-server-list-200-schema"></span> Schema
   
  

[ServerGetServerListResponseBody](#server-get-server-list-response-body)

### <span id="server-update-pkg-paste-server"></span> updatePkgPasteServer server (*server#updatePkgPasteServer*)

```
PUT /v1/server/paste/{serverID}
```

update paste server

> [Read more](https://vuls.biz/ "Update Paste Server")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| serverID | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | Server ID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |
| UpdatePkgPasteServerRequestBody | `body` | [ServerUpdatePkgPasteServerRequestBody](#server-update-pkg-paste-server-request-body) | `models.ServerUpdatePkgPasteServerRequestBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#server-update-pkg-paste-server-200) | OK | OK response. |  | [schema](#server-update-pkg-paste-server-200-schema) |

#### Responses


##### <span id="server-update-pkg-paste-server-200"></span> 200 - OK response.
Status: OK

###### <span id="server-update-pkg-paste-server-200-schema"></span> Schema

### <span id="server-update-server"></span> updateServer server (*server#updateServer*)

```
PUT /v1/server/{serverID}
```

update server

> [Read more](https://vuls.biz/ "Update Server")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| serverID | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | Server ID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |
| UpdateServerRequestBody | `body` | [ServerUpdateServerRequestBody](#server-update-server-request-body) | `models.ServerUpdateServerRequestBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#server-update-server-200) | OK | OK response. |  | [schema](#server-update-server-200-schema) |

#### Responses


##### <span id="server-update-server-200"></span> 200 - OK response.
Status: OK

###### <span id="server-update-server-200-schema"></span> Schema
   
  

[ServerUpdateServerResponseBody](#server-update-server-response-body)

### <span id="task-add-task-comment"></span> addTaskComment task (*task#addTaskComment*)

```
POST /v1/task/{taskID}/comment
```

add task comment

> [Read more](https://vuls.biz/ "Add Task Comment")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskID | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | Task ID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |
| AddTaskCommentRequestBody | `body` | [TaskAddTaskCommentRequestBody](#task-add-task-comment-request-body) | `models.TaskAddTaskCommentRequestBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#task-add-task-comment-200) | OK | OK response. |  | [schema](#task-add-task-comment-200-schema) |

#### Responses


##### <span id="task-add-task-comment-200"></span> 200 - OK response.
Status: OK

###### <span id="task-add-task-comment-200-schema"></span> Schema
   
  

[TaskAddTaskCommentResponseBody](#task-add-task-comment-response-body)

### <span id="task-get-task-detail"></span> getTaskDetail task (*task#getTaskDetail*)

```
GET /v1/task/{taskID}
```

task detail

> [Read more](https://vuls.biz/ "Get Task Detail")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskID | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | Task ID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#task-get-task-detail-200) | OK | OK response. |  | [schema](#task-get-task-detail-200-schema) |

#### Responses


##### <span id="task-get-task-detail-200"></span> 200 - OK response.
Status: OK

###### <span id="task-get-task-detail-200-schema"></span> Schema
   
  

[TaskGetTaskDetailResponseBody](#task-get-task-detail-response-body)

### <span id="task-get-task-list"></span> getTaskList task (*task#getTaskList*)

```
GET /v1/tasks
```

task list

> [Read more](https://vuls.biz/ "Get Task List")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Authorization | `header` | string | `string` |  |  |  | api key auth |
| filterCpeID | `query` | integer | `int64` |  |  |  | CpeID filter |
| filterCveID | `query` | string | `string` |  |  |  | CveID filter |
| filterIgnore | `query` | boolean | `bool` |  |  |  | Ignore filter(trueの場合は、非表示のものを取得しない。falseの場合は全件取得) |
| filterMainUserIDs | `query` | []integer | `[]int64` | `multi` |  |  | MainUserIDs filter |
| filterPkgID | `query` | integer | `int64` |  |  |  | PackageID filter |
| filterPriority | `query` | []string | `[]string` | `multi` |  |  | Priority filter |
| filterRoleID | `query` | integer | `int64` |  |  |  | ServerRoleID filter |
| filterServerID | `query` | integer | `int64` |  |  |  | ServerID filter |
| filterStatus | `query` | []string | `[]string` | `multi` |  | `["new","investigating","ongoing"]` | Status filter |
| filterSubUserIDs | `query` | []integer | `[]int64` | `multi` |  |  | SubUserIDs filter |
| limit | `query` | integer | `int64` |  |  | `20` | Limit |
| offset | `query` | integer | `int64` |  |  |  | Offset |
| page | `query` | integer | `int64` |  |  | `1` | Page Number |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#task-get-task-list-200) | OK | OK response. |  | [schema](#task-get-task-list-200-schema) |

#### Responses


##### <span id="task-get-task-list-200"></span> 200 - OK response.
Status: OK

###### <span id="task-get-task-list-200-schema"></span> Schema
   
  

[TaskGetTaskListResponseBody](#task-get-task-list-response-body)

### <span id="task-update-task"></span> updateTask task (*task#updateTask*)

```
PUT /v1/task/{taskID}
```

update task

> [Read more](https://vuls.biz/ "Update Task")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskID | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | Task ID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |
| UpdateTaskRequestBody | `body` | [TaskUpdateTaskRequestBody](#task-update-task-request-body) | `models.TaskUpdateTaskRequestBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#task-update-task-200) | OK | OK response. |  | [schema](#task-update-task-200-schema) |

#### Responses


##### <span id="task-update-task-200"></span> 200 - OK response.
Status: OK

###### <span id="task-update-task-200-schema"></span> Schema
   
  

[TaskUpdateTaskResponseBody](#task-update-task-response-body)

### <span id="task-update-task-ignore"></span> updateTaskIgnore task (*task#updateTaskIgnore*)

```
PUT /v1/task/{taskID}/ignore
```

update task ignore

> [Read more](https://vuls.biz/ "Update Task Ignore")

#### URI Schemes
  * http

#### Security Requirements
  * api_key_header_Authorization

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskID | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | Task ID |
| Authorization | `header` | string | `string` |  |  |  | api key auth |
| UpdateTaskIgnoreRequestBody | `body` | [TaskUpdateTaskIgnoreRequestBody](#task-update-task-ignore-request-body) | `models.TaskUpdateTaskIgnoreRequestBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#task-update-task-ignore-200) | OK | OK response. |  | [schema](#task-update-task-ignore-200-schema) |

#### Responses


##### <span id="task-update-task-ignore-200"></span> 200 - OK response.
Status: OK

###### <span id="task-update-task-ignore-200-schema"></span> Schema
   
  

[TaskUpdateTaskIgnoreResponseBody](#task-update-task-ignore-response-body)

## Models

### <span id="advisory-response-body"></span> AdvisoryResponseBody


> Advisory describes a advisory
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| advisoryID | string| `string` | ✓ | | AdvisoryID of advisory | `advisoryID` |
| osFamily | string| `string` | ✓ | | osFamily of advisory | `redhat` |



### <span id="affected-proc-response-body"></span> AffectedProcResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` | ✓ | | AffectedProc Name | `apache` |
| pid | string| `string` | ✓ | | PID | `12` |



### <span id="child-task-response-body"></span> ChildTaskResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| applyingPatchOn | date (formatted string)| `strfmt.Date` |  | | ApplyingPatchOn of task | `2018-07-13` |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | created time of task | `2018-07-14T08:13:28Z` |
| cveID | string| `string` | ✓ | | CVE ID of task | `CVE-2017-6799` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of task | `1` |
| ignore | boolean| `bool` | ✓ | | Ignore of task | `true` |
| ignoreUntil | string| `string` |  | | Ignore until of task | `vector` |
| mainUserID | int64 (formatted integer)| `int64` |  | | MainUserID of task | `1` |
| mainUserName | string| `string` |  | | MainUserName of task | `main-user-name` |
| priority | string| `string` | ✓ | | Priority of task | `high` |
| serverID | int64 (formatted integer)| `int64` | ✓ | | ServerID of task | `1` |
| status | string| `string` | ✓ | | Status of task | `new` |
| subUserID | int64 (formatted integer)| `int64` |  | | SubUserID of task | `1` |
| subUserName | string| `string` |  | | SubUserName of task | `sub-user-name` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time of task | `2018-07-14T08:13:28Z` |



### <span id="cve-get-cve-detail-response-body"></span> CveGetCveDetailResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| advisories | [][AdvisoryResponseBody](#advisory-response-body)| `[]*AdvisoryResponseBody` |  | | advisory of cve | `[{"advisoryID":"advisoryID","osFamily":"redhat"},{"advisoryID":"advisoryID","osFamily":"redhat"},{"advisoryID":"advisoryID","osFamily":"redhat"}]` |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | created time | `2018-07-14T08:13:28Z` |
| cveID | string| `string` | ✓ | | Cve ID string of cve | `CVE-2018-1234` |
| cvss | [ReadCloser](#read-closer)| `io.ReadCloser` | ✓ | | cvss of cve | `Unde dolorum distinctio voluptatem nemo ut.` |
| cwes | [][CweResponseBody](#cwe-response-body)| `[]*CweResponseBody` | ✓ | | cwes of cve | `[{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"}]` |
| envMetricV2s | [][EnvMetricV2ResponseBody](#env-metric-v2-response-body)| `[]*EnvMetricV2ResponseBody` | ✓ | | envMetricV2 of cve | `[{"cdp":"","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","roleID":1,"roleName":"roleName","td":"","updatedAt":"2018-07-14T08:13:28Z"},{"cdp":"","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","roleID":1,"roleName":"roleName","td":"","updatedAt":"2018-07-14T08:13:28Z"}]` |
| envMetricV3s | [][EnvMetricV3ResponseBody](#env-metric-v3-response-body)| `[]*EnvMetricV3ResponseBody` | ✓ | | envMetricV3 of cve | `[{"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","ma":"","mac":"","mav":"","mc":"","mpr":"","ms":"","mui":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","ma":"","mac":"","mav":"","mc":"","mpr":"","ms":"","mui":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"}]` |
| references | map of string| `map[string]string` | ✓ | | references of cve | `{"nvd":"https://xxxxxx"}` |
| secMetrics | [][SecMetricResponseBody](#sec-metric-response-body)| `[]*SecMetricResponseBody` | ✓ | | secMetric of cve | `[{"ar":"","cr":"","createdAt":"2018-07-14T08:13:28Z","ir":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},{"ar":"","cr":"","createdAt":"2018-07-14T08:13:28Z","ir":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},{"ar":"","cr":"","createdAt":"2018-07-14T08:13:28Z","ir":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"}]` |
| serverOsFamilies | []string| `[]string` | ✓ | | serverOsFamilies of cve | `["Odit suscipit suscipit.","Dolor reiciendis.","Est accusamus et repudiandae.","Illum atque dicta sapiente optio commodi."]` |
| tmpMetricV2 | [TmpMetricResponseBody](#tmp-metric-response-body)| `TmpMetricResponseBody` | ✓ | |  |  |
| tmpMetricV3 | [TmpMetricResponseBody](#tmp-metric-response-body)| `TmpMetricResponseBody` | ✓ | |  |  |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time | `2018-07-14T08:13:28Z` |



### <span id="cve-get-cve-list-response-body"></span> CveGetCveListResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cves | [][CveListResponseBody](#cve-list-response-body)| `[]*CveListResponseBody` |  | | Cves list | `[{"advisoryIDs":["advisoryID"],"allTaskCount":100,"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","cwes":[{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"}],"hasExploit":true,"hasMitigation":true,"hasWorkaround":true,"isNotActive":true,"isOwaspTopTen2017":true,"maxV2":9,"maxV3":9,"newTaskCount":10,"scoreV2s":{"nvd":9,"redhat":7},"scoreV3s":{"nvd":8,"redhat":9},"title":"title","topicCount":10,"topicLastUpdatedAt":"2018-07-14T08:13:28Z","updatedAt":"2018-07-14T08:13:28Z","vectorV2s":{"jvn":"AV:L/AC:M/Au:N/C:C/I:N/A:N","nvd":"AV:L/AC:M/Au:N/C:C/I:N/A:N"},"vectorV3s":{"jvn":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N","nvd":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N"}},{"advisoryIDs":["advisoryID"],"allTaskCount":100,"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","cwes":[{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"}],"hasExploit":true,"hasMitigation":true,"hasWorkaround":true,"isNotActive":true,"isOwaspTopTen2017":true,"maxV2":9,"maxV3":9,"newTaskCount":10,"scoreV2s":{"nvd":9,"redhat":7},"scoreV3s":{"nvd":8,"redhat":9},"title":"title","topicCount":10,"topicLastUpdatedAt":"2018-07-14T08:13:28Z","updatedAt":"2018-07-14T08:13:28Z","vectorV2s":{"jvn":"AV:L/AC:M/Au:N/C:C/I:N/A:N","nvd":"AV:L/AC:M/Au:N/C:C/I:N/A:N"},"vectorV3s":{"jvn":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N","nvd":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N"}},{"advisoryIDs":["advisoryID"],"allTaskCount":100,"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","cwes":[{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"}],"hasExploit":true,"hasMitigation":true,"hasWorkaround":true,"isNotActive":true,"isOwaspTopTen2017":true,"maxV2":9,"maxV3":9,"newTaskCount":10,"scoreV2s":{"nvd":9,"redhat":7},"scoreV3s":{"nvd":8,"redhat":9},"title":"title","topicCount":10,"topicLastUpdatedAt":"2018-07-14T08:13:28Z","updatedAt":"2018-07-14T08:13:28Z","vectorV2s":{"jvn":"AV:L/AC:M/Au:N/C:C/I:N/A:N","nvd":"AV:L/AC:M/Au:N/C:C/I:N/A:N"},"vectorV3s":{"jvn":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N","nvd":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N"}},{"advisoryIDs":["advisoryID"],"allTaskCount":100,"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","cwes":[{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"}],"hasExploit":true,"hasMitigation":true,"hasWorkaround":true,"isNotActive":true,"isOwaspTopTen2017":true,"maxV2":9,"maxV3":9,"newTaskCount":10,"scoreV2s":{"nvd":9,"redhat":7},"scoreV3s":{"nvd":8,"redhat":9},"title":"title","topicCount":10,"topicLastUpdatedAt":"2018-07-14T08:13:28Z","updatedAt":"2018-07-14T08:13:28Z","vectorV2s":{"jvn":"AV:L/AC:M/Au:N/C:C/I:N/A:N","nvd":"AV:L/AC:M/Au:N/C:C/I:N/A:N"},"vectorV3s":{"jvn":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N","nvd":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N"}}]` |
| paging | [PagingResponseBody](#paging-response-body)| `PagingResponseBody` |  | |  |  |



### <span id="cve-list-response-body"></span> CveListResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| advisoryIDs | []string| `[]string` |  | | advisoryIDs of cve | `["advisoryID"]` |
| allTaskCount | int64 (formatted integer)| `int64` | ✓ | | AllTaskCount of cve | `100` |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | created time | `2018-07-14T08:13:28Z` |
| cveID | string| `string` | ✓ | | Cve ID string of cve | `CVE-2018-1234` |
| cwes | [][CweResponseBody](#cwe-response-body)| `[]*CweResponseBody` | ✓ | | cwes of cve | `[{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"}]` |
| hasExploit | boolean| `bool` |  | | hasExploit of cve | `true` |
| hasMitigation | boolean| `bool` |  | | hasMitigation of cve | `true` |
| hasWorkaround | boolean| `bool` |  | | hasWorkaroundof cve | `true` |
| isNotActive | boolean| `bool` | ✓ | | Flag of active cve | `true` |
| isOwaspTopTen2017 | boolean| `bool` | ✓ | | isOwaspTopTen2017 of cve | `true` |
| maxV2 | double (formatted number)| `float64` | ✓ | | maxV2 of cve | `9` |
| maxV3 | double (formatted number)| `float64` | ✓ | | maxV3 of cve | `9` |
| newTaskCount | int64 (formatted integer)| `int64` | ✓ | | NewTaskCount of cve | `10` |
| scoreV2s | map of double (formatted number)| `map[string]float64` | ✓ | | cvss v2 scores of cve | `{"nvd":9,"redhat":7}` |
| scoreV3s | map of double (formatted number)| `map[string]float64` | ✓ | | cvss v3 scores of cve | `{"nvd":8,"redhat":9}` |
| title | string| `string` | ✓ | | Title of cve | `title` |
| topicCount | int64 (formatted integer)| `int64` | ✓ | | topicCount of cve | `10` |
| topicLastUpdatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | topicLastUpdatedAt of cve | `2018-07-14T08:13:28Z` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time | `2018-07-14T08:13:28Z` |
| vectorV2s | map of string| `map[string]string` | ✓ | | cvss v2 vectors of cve | `{"jvn":"AV:L/AC:M/Au:N/C:C/I:N/A:N","nvd":"AV:L/AC:M/Au:N/C:C/I:N/A:N"}` |
| vectorV3s | map of string| `map[string]string` | ✓ | | cvss v3 vectors of cve | `{"jvn":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N","nvd":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N"}` |



### <span id="cwe-response-body"></span> CweResponseBody


> Cwe describes a cwe
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cweID | string| `string` | ✓ | | cweID of cwe | `CWE-416` |
| english | string| `string` | ✓ | | english summary of cwe | `english summary` |
| japanese | string| `string` | ✓ | | japanese summary of cwe | `japanese summary` |
| owaspTopTen2017 | string| `string` |  | | owaspTopTen2017 of cwe | `10` |
| sourceType | string| `string` | ✓ | | sourceType of cwe | `nvd` |



### <span id="detection-method-response-body"></span> DetectionMethod ResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` | ✓ | | Detection Method Name | `vuls` |
| reliabilityScore | int64 (formatted integer)| `int64` | ✓ | | ReliabilityScore | `100` |



### <span id="detection-tool-response-body"></span> DetectionToolResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` | ✓ | | Detection Tool Name | `changelog` |
| patchAppliedAt | date-time (formatted string)| `strfmt.DateTime` |  | | PatchAppliedAt | `2018-07-14T08:13:28Z` |



### <span id="env-metric-v2-response-body"></span> EnvMetricV2ResponseBody


> EnvMetricV2 describes a envMetricV2
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cdp | string| `string` | ✓ | | CDP of envMetricV2 |  |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | created time | `2018-07-14T08:13:28Z` |
| cveID | string| `string` | ✓ | | CveID of envMetricV2 | `CVE-2018-1234` |
| roleID | int64 (formatted integer)| `int64` | ✓ | | ServerRoleID of envMetricV2 | `1` |
| roleName | string| `string` | ✓ | | ServerRoleName of envMetricV2 | `roleName` |
| td | string| `string` | ✓ | | TD of envMetricV2 |  |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time | `2018-07-14T08:13:28Z` |



### <span id="env-metric-v3-response-body"></span> EnvMetricV3ResponseBody


> EnvMetricV3 describes a envMetricV3
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | created time | `2018-07-14T08:13:28Z` |
| cveID | string| `string` | ✓ | | CveID of envMetricV3 | `CVE-2018-1234` |
| ma | string| `string` | ✓ | | MA of envMetricV3 |  |
| mac | string| `string` | ✓ | | MAC of envMetricV3 |  |
| mav | string| `string` | ✓ | | MAV of envMetricV3 |  |
| mc | string| `string` | ✓ | | MC of envMetricV3 |  |
| mpr | string| `string` | ✓ | | MPR of envMetricV3 |  |
| ms | string| `string` | ✓ | | MS of envMetricV3 |  |
| mui | string| `string` | ✓ | | MUI of envMetricV3 |  |
| roleID | int64 (formatted integer)| `int64` | ✓ | | ServerRoleID of envMetricV3 | `1` |
| roleName | string| `string` | ✓ | | ServerRoleName of envMetricV3 | `roleName` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time | `2018-07-14T08:13:28Z` |



### <span id="library-pkg-child-response-body"></span> LibraryPkgChildResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | crated time of package or cpe | `2018-07-14T08:13:28Z` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of library package | `1` |
| name | string| `string` | ✓ | | Name of library package | `package01` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time of package or cpe | `2018-07-14T08:13:28Z` |
| version | string| `string` | ✓ | | Version of library package | `1.0` |



### <span id="lockfile-add-lockfile-request-body"></span> LockfileAddLockfileRequestBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| fileContent | string| `string` | ✓ | | fileContent of the lockfile |  |
| path | string| `string` | ✓ | | Path of lockfile | `/FutureVuls/package.json` |
| serverID | int64 (formatted integer)| `int64` | ✓ | | ServerID | `1` |



### <span id="lockfile-add-lockfile-response-body"></span> LockfileAddLockfileResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| createdAt | string| `string` | ✓ | | created time of lockfile | `2018-07-14T08:13:28Z` |
| fileContent | string| `string` | ✓ | | FileContent of lockfile |  |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of lockfile | `1` |
| libraryPkgs | [][LibraryPkgChildResponseBody](#library-pkg-child-response-body)| `[]*LibraryPkgChildResponseBody` |  | | LibraryPkgs of lockfile | `[{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"}]` |
| path | string| `string` | ✓ | | Path of lockfile | `/FutureVuls/package.json` |
| server | [ServerChildResponseBody](#server-child-response-body)| `ServerChildResponseBody` |  | |  |  |
| updatedAt | string| `string` | ✓ | | updated time of lockfile | `2018-07-14T08:13:28Z` |



### <span id="lockfile-get-lockfile-detail-response-body"></span> LockfileGetLockfileDetailResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| createdAt | string| `string` | ✓ | | created time of lockfile | `2018-07-14T08:13:28Z` |
| fileContent | string| `string` | ✓ | | FileContent of lockfile |  |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of lockfile | `1` |
| libraryPkgs | [][LibraryPkgChildResponseBody](#library-pkg-child-response-body)| `[]*LibraryPkgChildResponseBody` |  | | LibraryPkgs of lockfile | `[{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"}]` |
| path | string| `string` | ✓ | | Path of lockfile | `/FutureVuls/package.json` |
| server | [ServerChildResponseBody](#server-child-response-body)| `ServerChildResponseBody` |  | |  |  |
| updatedAt | string| `string` | ✓ | | updated time of lockfile | `2018-07-14T08:13:28Z` |



### <span id="lockfile-get-lockfile-list-response-body"></span> LockfileGetLockfileListResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| lockfiles | [][LockfileListResponseBody](#lockfile-list-response-body)| `[]*LockfileListResponseBody` |  | | Lockfiles list | `[{"createdAt":"2018-07-14T08:13:28Z","fileContent":"","id":1,"libraryPkgs":[{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"}],"path":"/FutureVuls/package.json","server":{"createdAt":"2018-07-14T08:13:28Z","defaultUserId":1,"defaultUserName":"vuls user","hostUuid":"141df30a-ecd0-39f4-a8f4-1ef216a4b5f2","id":1,"lastScannedAt":"2018-07-14T08:13:28Z","lastUploadedAt":"2018-07-14T08:13:28Z","needKernelRestart":true,"osFamily":"centos","osVersion":"6","serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","serverroleId":1,"serverroleName":"server_role01","tags":["Fugit repellendus illo.","Aperiam ipsa voluptate autem unde.","Fuga accusamus aut."],"updatedAt":"2018-07-14T08:13:28Z"},"updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","fileContent":"","id":1,"libraryPkgs":[{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"}],"path":"/FutureVuls/package.json","server":{"createdAt":"2018-07-14T08:13:28Z","defaultUserId":1,"defaultUserName":"vuls user","hostUuid":"141df30a-ecd0-39f4-a8f4-1ef216a4b5f2","id":1,"lastScannedAt":"2018-07-14T08:13:28Z","lastUploadedAt":"2018-07-14T08:13:28Z","needKernelRestart":true,"osFamily":"centos","osVersion":"6","serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","serverroleId":1,"serverroleName":"server_role01","tags":["Fugit repellendus illo.","Aperiam ipsa voluptate autem unde.","Fuga accusamus aut."],"updatedAt":"2018-07-14T08:13:28Z"},"updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","fileContent":"","id":1,"libraryPkgs":[{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"}],"path":"/FutureVuls/package.json","server":{"createdAt":"2018-07-14T08:13:28Z","defaultUserId":1,"defaultUserName":"vuls user","hostUuid":"141df30a-ecd0-39f4-a8f4-1ef216a4b5f2","id":1,"lastScannedAt":"2018-07-14T08:13:28Z","lastUploadedAt":"2018-07-14T08:13:28Z","needKernelRestart":true,"osFamily":"centos","osVersion":"6","serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","serverroleId":1,"serverroleName":"server_role01","tags":["Fugit repellendus illo.","Aperiam ipsa voluptate autem unde.","Fuga accusamus aut."],"updatedAt":"2018-07-14T08:13:28Z"},"updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","fileContent":"","id":1,"libraryPkgs":[{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"}],"path":"/FutureVuls/package.json","server":{"createdAt":"2018-07-14T08:13:28Z","defaultUserId":1,"defaultUserName":"vuls user","hostUuid":"141df30a-ecd0-39f4-a8f4-1ef216a4b5f2","id":1,"lastScannedAt":"2018-07-14T08:13:28Z","lastUploadedAt":"2018-07-14T08:13:28Z","needKernelRestart":true,"osFamily":"centos","osVersion":"6","serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","serverroleId":1,"serverroleName":"server_role01","tags":["Fugit repellendus illo.","Aperiam ipsa voluptate autem unde.","Fuga accusamus aut."],"updatedAt":"2018-07-14T08:13:28Z"},"updatedAt":"2018-07-14T08:13:28Z"}]` |
| paging | [PagingResponseBody](#paging-response-body)| `PagingResponseBody` |  | |  |  |



### <span id="lockfile-list-response-body"></span> LockfileListResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| createdAt | string| `string` | ✓ | | created time of lockfile | `2018-07-14T08:13:28Z` |
| fileContent | string| `string` | ✓ | | FileContent of lockfile |  |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of lockfile | `1` |
| libraryPkgs | [][LibraryPkgChildResponseBody](#library-pkg-child-response-body)| `[]*LibraryPkgChildResponseBody` |  | | LibraryPkgs of lockfile | `[{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"}]` |
| path | string| `string` | ✓ | | Path of lockfile | `/FutureVuls/package.json` |
| server | [ServerChildResponseBody](#server-child-response-body)| `ServerChildResponseBody` |  | |  |  |
| updatedAt | string| `string` | ✓ | | updated time of lockfile | `2018-07-14T08:13:28Z` |



### <span id="lockfile-update-lockfile-request-body"></span> LockfileUpdateLockfileRequestBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| fileContent | string| `string` |  | | fileContent of the lockfile |  |
| path | string| `string` |  | | Path of lockfile | `/FutureVuls/package.json` |



### <span id="lockfile-update-lockfile-response-body"></span> LockfileUpdateLockfileResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| createdAt | string| `string` | ✓ | | created time of lockfile | `2018-07-14T08:13:28Z` |
| fileContent | string| `string` | ✓ | | FileContent of lockfile |  |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of lockfile | `1` |
| libraryPkgs | [][LibraryPkgChildResponseBody](#library-pkg-child-response-body)| `[]*LibraryPkgChildResponseBody` |  | | LibraryPkgs of lockfile | `[{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"}]` |
| path | string| `string` | ✓ | | Path of lockfile | `/FutureVuls/package.json` |
| server | [ServerChildResponseBody](#server-child-response-body)| `ServerChildResponseBody` |  | |  |  |
| updatedAt | string| `string` | ✓ | | updated time of lockfile | `2018-07-14T08:13:28Z` |



### <span id="need-restart-proc-response-body"></span> NeedRestartProcResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| initSystem | string| `string` | ✓ | | InitSystem of NeedRestartProc | `initSystem` |
| path | string| `string` | ✓ | | Path of NeedRestartProc | `path` |
| pid | string| `string` | ✓ | | PID | `12` |
| serviceName | string| `string` | ✓ | | ServiceName of NeedRestartProc | `serviceName` |



### <span id="paging-response-body"></span> PagingResponseBody


> Paging describes a paging object
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| limit | int64 (formatted integer)| `int64` | ✓ | | Limit | `20` |
| offset | int64 (formatted integer)| `int64` | ✓ | | Offset | `10` |
| page | int64 (formatted integer)| `int64` | ✓ | | Page | `1` |
| totalCount | int64 (formatted integer)| `int64` | ✓ | | TotalCount | `200` |
| totalPage | int64 (formatted integer)| `int64` | ✓ | | Total Page Size | `10` |



### <span id="pkg-cpe-add-cpe-request-body"></span> PkgCpeAddCpeRequestBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cpeName | string| `string` | ✓ | | Cpe Name(cpe uri or cpe format string) | `cpe:/a:berlios:discussion_forum_2k:3.3` |
| isURI | boolean| `bool` |  | `true`| isURI specifies whether cpeName is in URI format or FormatString format. | `true` |
| serverID | int64 (formatted integer)| `int64` | ✓ | | ServerID | `1` |



### <span id="pkg-cpe-add-cpe-response-body"></span> PkgCpeAddCpeResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cpeURI | string| `string` |  | | Cpe URI of cpe | `cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*` |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | crated time of package or cpe | `2018-07-14T08:13:28Z` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of package or cpe | `1` |
| name | string| `string` | ✓ | | Name of package or cpe | `package01` |
| release | string| `string` |  | | Release of package | `release` |
| server | [ServerChildResponseBody](#server-child-response-body)| `ServerChildResponseBody` | ✓ | |  |  |
| tasks | [][ChildTaskResponseBody](#child-task-response-body)| `[]*ChildTaskResponseBody` |  | | updated time of server | `[{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"}]` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time of package or cpe | `2018-07-14T08:13:28Z` |
| version | string| `string` | ✓ | | Version of package or cpe | `1.0` |



### <span id="pkg-cpe-child-response-body"></span> PkgCpeChildResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| affectedProcs | [][AffectedProcResponseBody](#affected-proc-response-body)| `[]*AffectedProcResponseBody` |  | | AffectedProcess list of package | `[{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"}]` |
| cpeID | int64 (formatted integer)| `int64` |  | | CpeID of cpe | `1` |
| cpeURI | string| `string` |  | | Cpe URI of cpe | `cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*` |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | crated time of package or cpe | `2018-07-14T08:13:28Z` |
| name | string| `string` | ✓ | | Name of package or cpe | `package01` |
| newRelease | string| `string` |  | | New Release of package | `new release` |
| newVersion | string| `string` |  | | New Version of package | `2.0` |
| pkgID | int64 (formatted integer)| `int64` |  | | Package ID of package | `1` |
| release | string| `string` |  | | Release of package | `release` |
| repository | string| `string` |  | | Repository of package | `repository` |
| serverID | int64 (formatted integer)| `int64` | ✓ | | ServerID of package or cpe | `1` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time of package or cpe | `2018-07-14T08:13:28Z` |
| version | string| `string` | ✓ | | Version of package or cpe | `1.0` |



### <span id="pkg-cpe-delete-cpe-deprecated-request-body"></span> PkgCpeDeleteCpeDeprecatedRequestBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cpeID | int64 (formatted integer)| `int64` | ✓ | | cpe ID | `4046142736569201700` |



### <span id="pkg-cpe-get-cpe-detail-response-body"></span> PkgCpeGetCpeDetailResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cpeURI | string| `string` |  | | Cpe URI of cpe | `cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*` |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | crated time of package or cpe | `2018-07-14T08:13:28Z` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of package or cpe | `1` |
| name | string| `string` | ✓ | | Name of package or cpe | `package01` |
| release | string| `string` |  | | Release of package | `release` |
| server | [ServerChildResponseBody](#server-child-response-body)| `ServerChildResponseBody` | ✓ | |  |  |
| tasks | [][ChildTaskResponseBody](#child-task-response-body)| `[]*ChildTaskResponseBody` |  | | updated time of server | `[{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"}]` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time of package or cpe | `2018-07-14T08:13:28Z` |
| version | string| `string` | ✓ | | Version of package or cpe | `1.0` |



### <span id="pkg-cpe-get-pkg-cpe-list-response-body"></span> PkgCpeGetPkgCpeListResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| paging | [PagingResponseBody](#paging-response-body)| `PagingResponseBody` |  | |  |  |
| pkgCpes | [][PkgCpeListResponseBody](#pkg-cpe-list-response-body)| `[]*PkgCpeListResponseBody` |  | | PkgCpes list | `[{"applicationName":"Application","cpeFS":"cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*","cpeID":1,"cpeURI":"cpe:/a:cisco:ios:15.2","createdAt":"2018-07-14T08:13:28Z","githubPkgID":1,"libraryPath":"/FutureVuls/go.sum","libraryPkgID":1,"name":"package01","needRestartProcs":[{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"}],"newRelease":"new release","newVersion":"2.0","notFixedYet":true,"ossLicense":"ossLicense","pkgID":1,"release":"release","repository":"repository","serverID":1,"serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","updatedAt":"2018-07-14T08:13:28Z","version":"1.0","wordpressPackageStatus":"wordpressPackageStatus","wordpressPkgID":1},{"applicationName":"Application","cpeFS":"cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*","cpeID":1,"cpeURI":"cpe:/a:cisco:ios:15.2","createdAt":"2018-07-14T08:13:28Z","githubPkgID":1,"libraryPath":"/FutureVuls/go.sum","libraryPkgID":1,"name":"package01","needRestartProcs":[{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"}],"newRelease":"new release","newVersion":"2.0","notFixedYet":true,"ossLicense":"ossLicense","pkgID":1,"release":"release","repository":"repository","serverID":1,"serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","updatedAt":"2018-07-14T08:13:28Z","version":"1.0","wordpressPackageStatus":"wordpressPackageStatus","wordpressPkgID":1},{"applicationName":"Application","cpeFS":"cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*","cpeID":1,"cpeURI":"cpe:/a:cisco:ios:15.2","createdAt":"2018-07-14T08:13:28Z","githubPkgID":1,"libraryPath":"/FutureVuls/go.sum","libraryPkgID":1,"name":"package01","needRestartProcs":[{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"}],"newRelease":"new release","newVersion":"2.0","notFixedYet":true,"ossLicense":"ossLicense","pkgID":1,"release":"release","repository":"repository","serverID":1,"serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","updatedAt":"2018-07-14T08:13:28Z","version":"1.0","wordpressPackageStatus":"wordpressPackageStatus","wordpressPkgID":1}]` |



### <span id="pkg-cpe-get-pkg-detail-response-body"></span> PkgCpeGetPkgDetailResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| affectedProcs | [][AffectedProcResponseBody](#affected-proc-response-body)| `[]*AffectedProcResponseBody` |  | | AffectedProcess list of package | `[{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"}]` |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | crated time of package or cpe | `2018-07-14T08:13:28Z` |
| id | int64 (formatted integer)| `int64` |  | | ID of package | `1` |
| name | string| `string` | ✓ | | Name of package or cpe | `package01` |
| needRestartProcs | [][NeedRestartProcResponseBody](#need-restart-proc-response-body)| `[]*NeedRestartProcResponseBody` |  | | NeedRestartProcess list of package | `[{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"}]` |
| newRelease | string| `string` |  | | New Release of package | `new release` |
| newVersion | string| `string` |  | | New Version of package | `2.0` |
| packageStatuses | map of string| `map[string]string` |  | | package status of package | `{"bash":"resolved"}` |
| release | string| `string` | ✓ | | Release of package | `release` |
| repository | string| `string` |  | | Repository of package | `repository` |
| server | [ServerChildResponseBody](#server-child-response-body)| `ServerChildResponseBody` | ✓ | |  |  |
| tasks | [][ChildTaskResponseBody](#child-task-response-body)| `[]*ChildTaskResponseBody` |  | | updated time of server | `[{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"}]` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time of package or cpe | `2018-07-14T08:13:28Z` |
| version | string| `string` | ✓ | | Version of package or cpe | `1.0` |



### <span id="pkg-cpe-list-response-body"></span> PkgCpeListResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| applicationName | string| `string` |  | | ApplicationName of library package | `Application` |
| cpeFS | string| `string` |  | | Cpe FS of cpe | `cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*` |
| cpeID | int64 (formatted integer)| `int64` |  | | CpeID of cpe | `1` |
| cpeURI | string| `string` |  | | Cpe URI of cpe | `cpe:/a:cisco:ios:15.2` |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | crated time of package or cpe | `2018-07-14T08:13:28Z` |
| githubPkgID | int64 (formatted integer)| `int64` |  | | githubPKGID of github pkg | `1` |
| libraryPath | string| `string` |  | | LibraryPath of library package | `/FutureVuls/go.sum` |
| libraryPkgID | int64 (formatted integer)| `int64` |  | | libraryPKGID of library pkg | `1` |
| name | string| `string` | ✓ | | Name of package or cpe | `package01` |
| needRestartProcs | [][NeedRestartProcResponseBody](#need-restart-proc-response-body)| `[]*NeedRestartProcResponseBody` |  | | NeedRestartProcess list of package | `[{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"}]` |
| newRelease | string| `string` |  | | New Release of package | `new release` |
| newVersion | string| `string` |  | | New Version of package | `2.0` |
| notFixedYet | boolean| `bool` |  | | Flag of Not fixed yet of package | `true` |
| ossLicense | string| `string` |  | | ossLicense of library package | `ossLicense` |
| pkgID | int64 (formatted integer)| `int64` |  | | Package ID of package | `1` |
| release | string| `string` |  | | Release of package | `release` |
| repository | string| `string` |  | | Repository of package | `repository` |
| serverID | int64 (formatted integer)| `int64` | ✓ | | ServerID of package or cpe | `1` |
| serverName | string| `string` | ✓ | | ServerName of package or cpe | `server01` |
| serverUuid | uuid (formatted string)| `strfmt.UUID` | ✓ | | ServerUUID of package or cpe | `abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time of package or cpe | `2018-07-14T08:13:28Z` |
| version | string| `string` | ✓ | | Version of package or cpe | `1.0` |
| wordpressPackageStatus | string| `string` |  | | WordpressPackageStatus of wordpress package | `wordpressPackageStatus` |
| wordpressPkgID | int64 (formatted integer)| `int64` |  | | wordpressPKGID of wordpress pkg | `1` |



### <span id="role-get-role-detail-response-body"></span> RoleGetRoleDetailResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| allTaskCount | int64 (formatted integer)| `int64` |  | | AllTaskCount of server role | `10` |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | created time of server role | `2018-07-14T08:13:28Z` |
| envMetricV2s | [][EnvMetricV2ResponseBody](#env-metric-v2-response-body)| `[]*EnvMetricV2ResponseBody` |  | | envMetricV2s of server role | `[{"cdp":"","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","roleID":1,"roleName":"roleName","td":"","updatedAt":"2018-07-14T08:13:28Z"},{"cdp":"","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","roleID":1,"roleName":"roleName","td":"","updatedAt":"2018-07-14T08:13:28Z"},{"cdp":"","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","roleID":1,"roleName":"roleName","td":"","updatedAt":"2018-07-14T08:13:28Z"}]` |
| envMetricV3s | [][EnvMetricV3ResponseBody](#env-metric-v3-response-body)| `[]*EnvMetricV3ResponseBody` |  | | envMetricV3s of server role | `[{"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","ma":"","mac":"","mav":"","mc":"","mpr":"","ms":"","mui":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","ma":"","mac":"","mav":"","mc":"","mpr":"","ms":"","mui":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","ma":"","mac":"","mav":"","mc":"","mpr":"","ms":"","mui":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"}]` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of server role | `1` |
| isDefault | boolean| `bool` |  | | isDefault of server role | `true` |
| name | string| `string` | ✓ | | Name of server role | `server-role-name` |
| newTaskCount | int64 (formatted integer)| `int64` |  | | NewTaskCount of server role | `10` |
| secMetric | [SecMetricResponseBody](#sec-metric-response-body)| `SecMetricResponseBody` |  | |  |  |
| servers | [][ServerChildResponseBody](#server-child-response-body)| `[]*ServerChildResponseBody` |  | | Servers of server role | `[{"createdAt":"2018-07-14T08:13:28Z","defaultUserId":1,"defaultUserName":"vuls user","hostUuid":"141df30a-ecd0-39f4-a8f4-1ef216a4b5f2","id":1,"lastScannedAt":"2018-07-14T08:13:28Z","lastUploadedAt":"2018-07-14T08:13:28Z","needKernelRestart":true,"osFamily":"centos","osVersion":"6","serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","serverroleId":1,"serverroleName":"server_role01","tags":["Fugit repellendus illo.","Aperiam ipsa voluptate autem unde.","Fuga accusamus aut."],"updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","defaultUserId":1,"defaultUserName":"vuls user","hostUuid":"141df30a-ecd0-39f4-a8f4-1ef216a4b5f2","id":1,"lastScannedAt":"2018-07-14T08:13:28Z","lastUploadedAt":"2018-07-14T08:13:28Z","needKernelRestart":true,"osFamily":"centos","osVersion":"6","serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","serverroleId":1,"serverroleName":"server_role01","tags":["Fugit repellendus illo.","Aperiam ipsa voluptate autem unde.","Fuga accusamus aut."],"updatedAt":"2018-07-14T08:13:28Z"}]` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time of server role | `2018-07-14T08:13:28Z` |



### <span id="role-get-role-list-response-body"></span> RoleGetRoleListResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| paging | [PagingResponseBody](#paging-response-body)| `PagingResponseBody` |  | |  |  |
| roles | [][RoleListResponseBody](#role-list-response-body)| `[]*RoleListResponseBody` |  | | ServerRole list | `[{"allTaskCount":10,"createdAt":"2018-07-14T08:13:28Z","id":1,"isDefault":true,"name":"server-role-name","newTaskCount":10,"secMetric":{"ar":"","cr":"","createdAt":"2018-07-14T08:13:28Z","ir":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},"serverCount":10,"updatedAt":"2018-07-14T08:13:28Z"},{"allTaskCount":10,"createdAt":"2018-07-14T08:13:28Z","id":1,"isDefault":true,"name":"server-role-name","newTaskCount":10,"secMetric":{"ar":"","cr":"","createdAt":"2018-07-14T08:13:28Z","ir":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},"serverCount":10,"updatedAt":"2018-07-14T08:13:28Z"},{"allTaskCount":10,"createdAt":"2018-07-14T08:13:28Z","id":1,"isDefault":true,"name":"server-role-name","newTaskCount":10,"secMetric":{"ar":"","cr":"","createdAt":"2018-07-14T08:13:28Z","ir":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},"serverCount":10,"updatedAt":"2018-07-14T08:13:28Z"}]` |



### <span id="role-list-response-body"></span> RoleListResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| allTaskCount | int64 (formatted integer)| `int64` |  | | AllTaskCount of server role | `10` |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | created time of server role | `2018-07-14T08:13:28Z` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of server role | `1` |
| isDefault | boolean| `bool` | ✓ | | isDefault of server role | `true` |
| name | string| `string` | ✓ | | Name of server role | `server-role-name` |
| newTaskCount | int64 (formatted integer)| `int64` |  | | NewTaskCount of server role | `10` |
| secMetric | [SecMetricResponseBody](#sec-metric-response-body)| `SecMetricResponseBody` |  | |  |  |
| serverCount | int64 (formatted integer)| `int64` |  | | Server Count of server role | `10` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time of server role | `2018-07-14T08:13:28Z` |



### <span id="role-update-role-request-body"></span> RoleUpdateRoleRequestBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| roleName | string| `string` |  | | RoleName of role | `new-role-name` |



### <span id="role-update-role-response-body"></span> RoleUpdateRoleResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| allTaskCount | int64 (formatted integer)| `int64` |  | | AllTaskCount of server role | `10` |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | created time of server role | `2018-07-14T08:13:28Z` |
| envMetricV2s | [][EnvMetricV2ResponseBody](#env-metric-v2-response-body)| `[]*EnvMetricV2ResponseBody` |  | | envMetricV2s of server role | `[{"cdp":"","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","roleID":1,"roleName":"roleName","td":"","updatedAt":"2018-07-14T08:13:28Z"},{"cdp":"","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","roleID":1,"roleName":"roleName","td":"","updatedAt":"2018-07-14T08:13:28Z"},{"cdp":"","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","roleID":1,"roleName":"roleName","td":"","updatedAt":"2018-07-14T08:13:28Z"},{"cdp":"","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","roleID":1,"roleName":"roleName","td":"","updatedAt":"2018-07-14T08:13:28Z"}]` |
| envMetricV3s | [][EnvMetricV3ResponseBody](#env-metric-v3-response-body)| `[]*EnvMetricV3ResponseBody` |  | | envMetricV3s of server role | `[{"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","ma":"","mac":"","mav":"","mc":"","mpr":"","ms":"","mui":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","ma":"","mac":"","mav":"","mc":"","mpr":"","ms":"","mui":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","ma":"","mac":"","mav":"","mc":"","mpr":"","ms":"","mui":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"}]` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of server role | `1` |
| isDefault | boolean| `bool` |  | | isDefault of server role | `true` |
| name | string| `string` | ✓ | | Name of server role | `server-role-name` |
| newTaskCount | int64 (formatted integer)| `int64` |  | | NewTaskCount of server role | `10` |
| secMetric | [SecMetricResponseBody](#sec-metric-response-body)| `SecMetricResponseBody` |  | |  |  |
| servers | [][ServerChildResponseBody](#server-child-response-body)| `[]*ServerChildResponseBody` |  | | Servers of server role | `[{"createdAt":"2018-07-14T08:13:28Z","defaultUserId":1,"defaultUserName":"vuls user","hostUuid":"141df30a-ecd0-39f4-a8f4-1ef216a4b5f2","id":1,"lastScannedAt":"2018-07-14T08:13:28Z","lastUploadedAt":"2018-07-14T08:13:28Z","needKernelRestart":true,"osFamily":"centos","osVersion":"6","serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","serverroleId":1,"serverroleName":"server_role01","tags":["Fugit repellendus illo.","Aperiam ipsa voluptate autem unde.","Fuga accusamus aut."],"updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","defaultUserId":1,"defaultUserName":"vuls user","hostUuid":"141df30a-ecd0-39f4-a8f4-1ef216a4b5f2","id":1,"lastScannedAt":"2018-07-14T08:13:28Z","lastUploadedAt":"2018-07-14T08:13:28Z","needKernelRestart":true,"osFamily":"centos","osVersion":"6","serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","serverroleId":1,"serverroleName":"server_role01","tags":["Fugit repellendus illo.","Aperiam ipsa voluptate autem unde.","Fuga accusamus aut."],"updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","defaultUserId":1,"defaultUserName":"vuls user","hostUuid":"141df30a-ecd0-39f4-a8f4-1ef216a4b5f2","id":1,"lastScannedAt":"2018-07-14T08:13:28Z","lastUploadedAt":"2018-07-14T08:13:28Z","needKernelRestart":true,"osFamily":"centos","osVersion":"6","serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","serverroleId":1,"serverroleName":"server_role01","tags":["Fugit repellendus illo.","Aperiam ipsa voluptate autem unde.","Fuga accusamus aut."],"updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","defaultUserId":1,"defaultUserName":"vuls user","hostUuid":"141df30a-ecd0-39f4-a8f4-1ef216a4b5f2","id":1,"lastScannedAt":"2018-07-14T08:13:28Z","lastUploadedAt":"2018-07-14T08:13:28Z","needKernelRestart":true,"osFamily":"centos","osVersion":"6","serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","serverroleId":1,"serverroleName":"server_role01","tags":["Fugit repellendus illo.","Aperiam ipsa voluptate autem unde.","Fuga accusamus aut."],"updatedAt":"2018-07-14T08:13:28Z"}]` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time of server role | `2018-07-14T08:13:28Z` |



### <span id="sec-metric-response-body"></span> SecMetricResponseBody


> SecMetric describes a secMetric
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ar | string| `string` | ✓ | | AR of secMetric |  |
| cr | string| `string` | ✓ | | CR of secMetric |  |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | created time | `2018-07-14T08:13:28Z` |
| ir | string| `string` | ✓ | | IR of secMetric |  |
| roleID | int64 (formatted integer)| `int64` | ✓ | | ServerRoleID of secMetric | `1` |
| roleName | string| `string` | ✓ | | ServerRoleName of secMetric | `roleName` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time | `2018-07-14T08:13:28Z` |



### <span id="server-child-response-body"></span> ServerChildResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| createdAt | string| `string` | ✓ | | crated time of server | `2018-07-14T08:13:28Z` |
| defaultUserId | int64 (formatted integer)| `int64` |  | | default user ID of server | `1` |
| defaultUserName | string| `string` |  | | default user name of server | `vuls user` |
| hostUuid | uuid (formatted string)| `strfmt.UUID` | ✓ | | UUID of server | `141df30a-ecd0-39f4-a8f4-1ef216a4b5f2` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of server | `1` |
| lastScannedAt | string| `string` |  | | last scanned time of server | `2018-07-14T08:13:28Z` |
| lastUploadedAt | string| `string` |  | | last uploaded time of server | `2018-07-14T08:13:28Z` |
| needKernelRestart | boolean| `bool` | ✓ | | Whether server needs kernel restart | `true` |
| osFamily | string| `string` | ✓ | | OS Name of server | `centos` |
| osVersion | string| `string` | ✓ | | OS Version of server | `6` |
| serverName | string| `string` | ✓ | | Name of server | `server01` |
| serverUuid | uuid (formatted string)| `strfmt.UUID` | ✓ | | UUID of server | `abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2` |
| serverroleId | int64 (formatted integer)| `int64` | ✓ | | ID of server role | `1` |
| serverroleName | string| `string` | ✓ | | Name of server role | `server_role01` |
| tags | []string| `[]string` |  | | tags is list of server tag | `["Consequatur atque et animi doloribus.","Aperiam id molestias dolore.","Consectetur asperiores."]` |
| updatedAt | string| `string` | ✓ | | updated time of server | `2018-07-14T08:13:28Z` |



### <span id="server-create-pkg-paste-server-request-body"></span> ServerCreatePkgPasteServerRequestBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| kernelRelease | string| `string` | ✓ | | Kernel Release | `kernel release` |
| kernelVersion | string| `string` |  | | Kernel Version | `kernel version` |
| osFamily | string| `string` | ✓ | | Server OS Family | `20` |
| osVersion | string| `string` | ✓ | | Server OS Version | `20` |
| pkgPasteText | string| `string` | ✓ | | pkg paste text |  |
| serverName | string| `string` | ✓ | | Server Name | `Server Name` |



### <span id="server-create-pkg-paste-server-response-body"></span> ServerCreatePkgPasteServerResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| createdAt | string| `string` | ✓ | | crated time of server | `2018-07-14T08:13:28Z` |
| defaultUserId | int64 (formatted integer)| `int64` |  | | default user ID of server | `1` |
| defaultUserName | string| `string` |  | | default user name of server | `vuls user` |
| hostUuid | uuid (formatted string)| `strfmt.UUID` | ✓ | | UUID of server | `141df30a-ecd0-39f4-a8f4-1ef216a4b5f2` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of server | `1` |
| lastScannedAt | string| `string` |  | | last scanned time of server | `2018-07-14T08:13:28Z` |
| lastUploadedAt | string| `string` |  | | last uploaded time of server | `2018-07-14T08:13:28Z` |
| needKernelRestart | boolean| `bool` | ✓ | | Whether server needs kernel restart | `true` |
| osFamily | string| `string` | ✓ | | OS Name of server | `centos` |
| osVersion | string| `string` | ✓ | | OS Version of server | `6` |
| platformInstanceId | string| `string` | ✓ | | platformInstanceId of server | `i-xxxxxxx` |
| platformName | string| `string` | ✓ | | platformName of server | `aws` |
| serverIpv4 | string| `string` | ✓ | | IPv4 of server | `192.168.0.2` |
| serverName | string| `string` | ✓ | | Name of server | `server01` |
| serverUuid | uuid (formatted string)| `strfmt.UUID` | ✓ | | UUID of server | `abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2` |
| serverroleId | int64 (formatted integer)| `int64` | ✓ | | ID of server role | `1` |
| serverroleName | string| `string` | ✓ | | Name of server role | `server_role01` |
| tags | [][ServerTagResponseBody](#server-tag-response-body)| `[]*ServerTagResponseBody` |  | | tags is list of server tag | `[{"id":1,"name":"tag"},{"id":1,"name":"tag"},{"id":1,"name":"tag"},{"id":1,"name":"tag"}]` |
| tasks | [][ChildTaskResponseBody](#child-task-response-body)| `[]*ChildTaskResponseBody` |  | | tasks of server | `[{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"}]` |
| updatedAt | string| `string` | ✓ | | updated time of server | `2018-07-14T08:13:28Z` |



### <span id="server-get-server-detail-by-uuid-response-body"></span> ServerGetServerDetailByUUIDResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| createdAt | string| `string` | ✓ | | crated time of server | `2018-07-14T08:13:28Z` |
| defaultUserId | int64 (formatted integer)| `int64` |  | | default user ID of server | `1` |
| defaultUserName | string| `string` |  | | default user name of server | `vuls user` |
| hostUuid | uuid (formatted string)| `strfmt.UUID` | ✓ | | UUID of server | `141df30a-ecd0-39f4-a8f4-1ef216a4b5f2` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of server | `1` |
| lastScannedAt | string| `string` |  | | last scanned time of server | `2018-07-14T08:13:28Z` |
| lastUploadedAt | string| `string` |  | | last uploaded time of server | `2018-07-14T08:13:28Z` |
| needKernelRestart | boolean| `bool` | ✓ | | Whether server needs kernel restart | `true` |
| osFamily | string| `string` | ✓ | | OS Name of server | `centos` |
| osVersion | string| `string` | ✓ | | OS Version of server | `6` |
| platformInstanceId | string| `string` | ✓ | | platformInstanceId of server | `i-xxxxxxx` |
| platformName | string| `string` | ✓ | | platformName of server | `aws` |
| serverIpv4 | string| `string` | ✓ | | IPv4 of server | `192.168.0.2` |
| serverName | string| `string` | ✓ | | Name of server | `server01` |
| serverUuid | uuid (formatted string)| `strfmt.UUID` | ✓ | | UUID of server | `abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2` |
| serverroleId | int64 (formatted integer)| `int64` | ✓ | | ID of server role | `1` |
| serverroleName | string| `string` | ✓ | | Name of server role | `server_role01` |
| tags | [][ServerTagResponseBody](#server-tag-response-body)| `[]*ServerTagResponseBody` |  | | tags is list of server tag | `[{"id":1,"name":"tag"},{"id":1,"name":"tag"}]` |
| tasks | [][ChildTaskResponseBody](#child-task-response-body)| `[]*ChildTaskResponseBody` |  | | tasks of server | `[{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"}]` |
| updatedAt | string| `string` | ✓ | | updated time of server | `2018-07-14T08:13:28Z` |



### <span id="server-get-server-detail-response-body"></span> ServerGetServerDetailResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| createdAt | string| `string` | ✓ | | crated time of server | `2018-07-14T08:13:28Z` |
| defaultUserId | int64 (formatted integer)| `int64` |  | | default user ID of server | `1` |
| defaultUserName | string| `string` |  | | default user name of server | `vuls user` |
| hostUuid | uuid (formatted string)| `strfmt.UUID` | ✓ | | UUID of server | `141df30a-ecd0-39f4-a8f4-1ef216a4b5f2` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of server | `1` |
| lastScannedAt | string| `string` |  | | last scanned time of server | `2018-07-14T08:13:28Z` |
| lastUploadedAt | string| `string` |  | | last uploaded time of server | `2018-07-14T08:13:28Z` |
| needKernelRestart | boolean| `bool` | ✓ | | Whether server needs kernel restart | `true` |
| osFamily | string| `string` | ✓ | | OS Name of server | `centos` |
| osVersion | string| `string` | ✓ | | OS Version of server | `6` |
| platformInstanceId | string| `string` | ✓ | | platformInstanceId of server | `i-xxxxxxx` |
| platformName | string| `string` | ✓ | | platformName of server | `aws` |
| serverIpv4 | string| `string` | ✓ | | IPv4 of server | `192.168.0.2` |
| serverName | string| `string` | ✓ | | Name of server | `server01` |
| serverUuid | uuid (formatted string)| `strfmt.UUID` | ✓ | | UUID of server | `abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2` |
| serverroleId | int64 (formatted integer)| `int64` | ✓ | | ID of server role | `1` |
| serverroleName | string| `string` | ✓ | | Name of server role | `server_role01` |
| tags | [][ServerTagResponseBody](#server-tag-response-body)| `[]*ServerTagResponseBody` |  | | tags is list of server tag | `[{"id":1,"name":"tag"},{"id":1,"name":"tag"},{"id":1,"name":"tag"},{"id":1,"name":"tag"}]` |
| tasks | [][ChildTaskResponseBody](#child-task-response-body)| `[]*ChildTaskResponseBody` |  | | tasks of server | `[{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"}]` |
| updatedAt | string| `string` | ✓ | | updated time of server | `2018-07-14T08:13:28Z` |



### <span id="server-get-server-list-response-body"></span> ServerGetServerListResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| paging | [PagingResponseBody](#paging-response-body)| `PagingResponseBody` |  | |  |  |
| servers | [][ServerListResponseBody](#server-list-response-body)| `[]*ServerListResponseBody` |  | | Servers list | `[{"createdAt":"2018-07-14T08:13:28Z","defaultUserId":1,"defaultUserName":"vuls user","hostUuid":"141df30a-ecd0-39f4-a8f4-1ef216a4b5f2","id":1,"lastScannedAt":"2018-07-14T08:13:28Z","lastUploadedAt":"2018-07-14T08:13:28Z","needKernelRestart":true,"osFamily":"centos","osVersion":"6","platformInstanceId":"i-xxxxxxx","platformName":"aws","serverIpv4":"192.168.0.2","serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","serverroleId":1,"serverroleName":"server_role01","successScanCount":7572510666884157000,"tags":[{"id":1,"name":"tag"},{"id":1,"name":"tag"},{"id":1,"name":"tag"}],"updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","defaultUserId":1,"defaultUserName":"vuls user","hostUuid":"141df30a-ecd0-39f4-a8f4-1ef216a4b5f2","id":1,"lastScannedAt":"2018-07-14T08:13:28Z","lastUploadedAt":"2018-07-14T08:13:28Z","needKernelRestart":true,"osFamily":"centos","osVersion":"6","platformInstanceId":"i-xxxxxxx","platformName":"aws","serverIpv4":"192.168.0.2","serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","serverroleId":1,"serverroleName":"server_role01","successScanCount":7572510666884157000,"tags":[{"id":1,"name":"tag"},{"id":1,"name":"tag"},{"id":1,"name":"tag"}],"updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","defaultUserId":1,"defaultUserName":"vuls user","hostUuid":"141df30a-ecd0-39f4-a8f4-1ef216a4b5f2","id":1,"lastScannedAt":"2018-07-14T08:13:28Z","lastUploadedAt":"2018-07-14T08:13:28Z","needKernelRestart":true,"osFamily":"centos","osVersion":"6","platformInstanceId":"i-xxxxxxx","platformName":"aws","serverIpv4":"192.168.0.2","serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","serverroleId":1,"serverroleName":"server_role01","successScanCount":7572510666884157000,"tags":[{"id":1,"name":"tag"},{"id":1,"name":"tag"},{"id":1,"name":"tag"}],"updatedAt":"2018-07-14T08:13:28Z"}]` |



### <span id="server-list-response-body"></span> ServerListResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| createdAt | string| `string` | ✓ | | crated time of server | `2018-07-14T08:13:28Z` |
| defaultUserId | int64 (formatted integer)| `int64` |  | | default user ID of server | `1` |
| defaultUserName | string| `string` |  | | default user name of server | `vuls user` |
| hostUuid | uuid (formatted string)| `strfmt.UUID` | ✓ | | UUID of server | `141df30a-ecd0-39f4-a8f4-1ef216a4b5f2` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of server | `1` |
| lastScannedAt | string| `string` |  | | last scanned time of server | `2018-07-14T08:13:28Z` |
| lastUploadedAt | string| `string` |  | | last uploaded time of server | `2018-07-14T08:13:28Z` |
| needKernelRestart | boolean| `bool` | ✓ | | Whether server needs kernel restart | `true` |
| osFamily | string| `string` | ✓ | | OS Name of server | `centos` |
| osVersion | string| `string` | ✓ | | OS Version of server | `6` |
| platformInstanceId | string| `string` | ✓ | | platformInstanceId of server | `i-xxxxxxx` |
| platformName | string| `string` | ✓ | | platformName of server | `aws` |
| serverIpv4 | string| `string` | ✓ | | IPv4 of server | `192.168.0.2` |
| serverName | string| `string` | ✓ | | Name of server | `server01` |
| serverUuid | uuid (formatted string)| `strfmt.UUID` | ✓ | | UUID of server | `abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2` |
| serverroleId | int64 (formatted integer)| `int64` | ✓ | | ID of server role | `1` |
| serverroleName | string| `string` | ✓ | | Name of server role | `server_role01` |
| successScanCount | int64 (formatted integer)| `int64` | ✓ | | successScanCount of server | `8268415631704044000` |
| tags | [][ServerTagResponseBody](#server-tag-response-body)| `[]*ServerTagResponseBody` |  | | tags is list of server tag | `[{"id":1,"name":"tag"},{"id":1,"name":"tag"}]` |
| updatedAt | string| `string` | ✓ | | updated time of server | `2018-07-14T08:13:28Z` |



### <span id="server-tag-response-body"></span> ServerTagResponseBody


> ServerTag describes a server tag
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| id | int64 (formatted integer)| `int64` | ✓ | | ID of server tag | `1` |
| name | string| `string` | ✓ | | Name of server tag | `tag` |



### <span id="server-update-pkg-paste-server-request-body"></span> ServerUpdatePkgPasteServerRequestBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| kernelRelease | string| `string` |  | | Kernel Release | `kernel release` |
| kernelVersion | string| `string` |  | | Kernel Version | `kernel version` |
| osVersion | string| `string` |  | | Server OS Version | `20` |
| pkgPasteText | string| `string` | ✓ | | pkg paste text |  |



### <span id="server-update-server-request-body"></span> ServerUpdateServerRequestBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| additionalInfo | string| `string` |  | | Additional information of the server | `This server is expensive` |
| defaultUserID | int64 (formatted integer)| `int64` |  | | DefaultUserID of server | `1` |
| roleID | int64 (formatted integer)| `int64` |  | | ServerRoleID of server | `1` |
| serverName | string| `string` |  | | ServerName of server | `new-server-name` |



### <span id="server-update-server-response-body"></span> ServerUpdateServerResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| createdAt | string| `string` | ✓ | | crated time of server | `2018-07-14T08:13:28Z` |
| defaultUserId | int64 (formatted integer)| `int64` |  | | default user ID of server | `1` |
| defaultUserName | string| `string` |  | | default user name of server | `vuls user` |
| hostUuid | uuid (formatted string)| `strfmt.UUID` | ✓ | | UUID of server | `141df30a-ecd0-39f4-a8f4-1ef216a4b5f2` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of server | `1` |
| lastScannedAt | string| `string` |  | | last scanned time of server | `2018-07-14T08:13:28Z` |
| lastUploadedAt | string| `string` |  | | last uploaded time of server | `2018-07-14T08:13:28Z` |
| needKernelRestart | boolean| `bool` | ✓ | | Whether server needs kernel restart | `true` |
| osFamily | string| `string` | ✓ | | OS Name of server | `centos` |
| osVersion | string| `string` | ✓ | | OS Version of server | `6` |
| platformInstanceId | string| `string` | ✓ | | platformInstanceId of server | `i-xxxxxxx` |
| platformName | string| `string` | ✓ | | platformName of server | `aws` |
| serverIpv4 | string| `string` | ✓ | | IPv4 of server | `192.168.0.2` |
| serverName | string| `string` | ✓ | | Name of server | `server01` |
| serverUuid | uuid (formatted string)| `strfmt.UUID` | ✓ | | UUID of server | `abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2` |
| serverroleId | int64 (formatted integer)| `int64` | ✓ | | ID of server role | `1` |
| serverroleName | string| `string` | ✓ | | Name of server role | `server_role01` |
| tags | [][ServerTagResponseBody](#server-tag-response-body)| `[]*ServerTagResponseBody` |  | | tags is list of server tag | `[{"id":1,"name":"tag"},{"id":1,"name":"tag"}]` |
| tasks | [][ChildTaskResponseBody](#child-task-response-body)| `[]*ChildTaskResponseBody` |  | | tasks of server | `[{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"}]` |
| updatedAt | string| `string` | ✓ | | updated time of server | `2018-07-14T08:13:28Z` |



### <span id="task-add-task-comment-request-body"></span> TaskAddTaskCommentRequestBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| commentContent | string| `string` | ✓ | | comment content | `comment` |



### <span id="task-add-task-comment-response-body"></span> TaskAddTaskCommentResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| advisoryIDs | []string| `[]string` |  | | advisoryIDs of cve | `["advisoryID"]` |
| applyingPatchOn | date (formatted string)| `strfmt.Date` |  | | ApplyingPatchOn of task | `2018-07-13` |
| comments | [][TaskCommentResponseBody](#task-comment-response-body)| `[]*TaskCommentResponseBody` |  | | Comment of task | `[{"comment":"comment","createdAt":"2018-07-14T08:13:28Z","id":1,"type":"system","updatedAt":"2018-07-14T08:13:28Z","userID":1,"userName":"user-name"},{"comment":"comment","createdAt":"2018-07-14T08:13:28Z","id":1,"type":"system","updatedAt":"2018-07-14T08:13:28Z","userID":1,"userName":"user-name"},{"comment":"comment","createdAt":"2018-07-14T08:13:28Z","id":1,"type":"system","updatedAt":"2018-07-14T08:13:28Z","userID":1,"userName":"user-name"},{"comment":"comment","createdAt":"2018-07-14T08:13:28Z","id":1,"type":"system","updatedAt":"2018-07-14T08:13:28Z","userID":1,"userName":"user-name"}]` |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | created time of task | `2018-07-14T08:13:28Z` |
| cveID | string| `string` | ✓ | | CVE ID of task | `CVE-2017-6799` |
| cvss | map of [ReadCloser](#read-closer)| `map[string]io.ReadCloser` |  | | Key Value of CveID and Cvss of task | `{"Quia dolor aut expedita velit est.":"Iste impedit rerum libero exercitationem deleniti et."}` |
| detectionMethods | [][DetectionMethodResponseBody](#detection-method-response-body)| `[]*DetectionMethodResponseBody` |  | | DetectionMethod of task | `[{"name":"vuls","reliabilityScore":100},{"name":"vuls","reliabilityScore":100},{"name":"vuls","reliabilityScore":100},{"name":"vuls","reliabilityScore":100}]` |
| detectionTools | [][DetectionToolResponseBody](#detection-tool-response-body)| `[]*DetectionToolResponseBody` |  | | DetectionTools of task | `[{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"},{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"}]` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of task | `1` |
| ignore | boolean| `bool` | ✓ | | Ignore of task | `true` |
| ignoreUntil | string| `string` |  | | Ignore until of task | `vector` |
| mainUserID | int64 (formatted integer)| `int64` |  | | MainUserID of task | `1` |
| mainUserName | string| `string` |  | | MainUserName of task | `main-user-name` |
| packageStatuses | map of string| `map[string]string` |  | | packageStatus of task | `{"bash":"resolved"}` |
| pkgCpes | [][PkgCpeChildResponseBody](#pkg-cpe-child-response-body)| `[]*PkgCpeChildResponseBody` |  | | Pcakge And Cpe list of task | `[{"affectedProcs":[{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"}],"cpeID":1,"cpeURI":"cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*","createdAt":"2018-07-14T08:13:28Z","name":"package01","newRelease":"new release","newVersion":"2.0","pkgID":1,"release":"release","repository":"repository","serverID":1,"updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"affectedProcs":[{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"}],"cpeID":1,"cpeURI":"cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*","createdAt":"2018-07-14T08:13:28Z","name":"package01","newRelease":"new release","newVersion":"2.0","pkgID":1,"release":"release","repository":"repository","serverID":1,"updatedAt":"2018-07-14T08:13:28Z","version":"1.0"}]` |
| priority | string| `string` | ✓ | | Priority of task | `high` |
| roleID | int64 (formatted integer)| `int64` | ✓ | | ServerRoleID of task | `1` |
| roleName | string| `string` | ✓ | | ServerRoleName of task | `server-role-name` |
| server | [ServerChildResponseBody](#server-child-response-body)| `ServerChildResponseBody` | ✓ | |  |  |
| serverID | int64 (formatted integer)| `int64` | ✓ | | ServerID of task | `1` |
| status | string| `string` | ✓ | | Status of task | `new` |
| subUserID | int64 (formatted integer)| `int64` |  | | SubUserID of task | `1` |
| subUserName | string| `string` |  | | SubUserName of task | `sub-user-name` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time of task | `2018-07-14T08:13:28Z` |



### <span id="task-comment-response-body"></span> TaskComment ResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| comment | string| `string` | ✓ | | Comment content of TaskComment | `comment` |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | created time of TaskComment | `2018-07-14T08:13:28Z` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of TaskComment | `1` |
| type | string| `string` | ✓ | | Type of TaskComment | `system` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time of TaskComment | `2018-07-14T08:13:28Z` |
| userID | int64 (formatted integer)| `int64` | ✓ | | UserID of TaskComment | `1` |
| userName | string| `string` | ✓ | | UserName of TaskComment | `user-name` |



### <span id="task-get-task-detail-response-body"></span> TaskGetTaskDetailResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| advisoryIDs | []string| `[]string` |  | | advisoryIDs of cve | `["advisoryID"]` |
| applyingPatchOn | date (formatted string)| `strfmt.Date` |  | | ApplyingPatchOn of task | `2018-07-13` |
| comments | [][TaskCommentResponseBody](#task-comment-response-body)| `[]*TaskCommentResponseBody` |  | | Comment of task | `[{"comment":"comment","createdAt":"2018-07-14T08:13:28Z","id":1,"type":"system","updatedAt":"2018-07-14T08:13:28Z","userID":1,"userName":"user-name"},{"comment":"comment","createdAt":"2018-07-14T08:13:28Z","id":1,"type":"system","updatedAt":"2018-07-14T08:13:28Z","userID":1,"userName":"user-name"},{"comment":"comment","createdAt":"2018-07-14T08:13:28Z","id":1,"type":"system","updatedAt":"2018-07-14T08:13:28Z","userID":1,"userName":"user-name"},{"comment":"comment","createdAt":"2018-07-14T08:13:28Z","id":1,"type":"system","updatedAt":"2018-07-14T08:13:28Z","userID":1,"userName":"user-name"}]` |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | created time of task | `2018-07-14T08:13:28Z` |
| cveID | string| `string` | ✓ | | CVE ID of task | `CVE-2017-6799` |
| cvss | map of [ReadCloser](#read-closer)| `map[string]io.ReadCloser` |  | | Key Value of CveID and Cvss of task | `{"Assumenda dolorem sed est.":"Tempore ea aperiam ipsum dolorem.","Aut molestias beatae et sed assumenda.":"Itaque iure illo maiores officia recusandae fugiat.","Nihil sint ea et.":"Et aut quasi consequatur sequi ex perspiciatis."}` |
| detectionMethods | [][DetectionMethodResponseBody](#detection-method-response-body)| `[]*DetectionMethodResponseBody` |  | | DetectionMethod of task | `[{"name":"vuls","reliabilityScore":100},{"name":"vuls","reliabilityScore":100},{"name":"vuls","reliabilityScore":100},{"name":"vuls","reliabilityScore":100}]` |
| detectionTools | [][DetectionToolResponseBody](#detection-tool-response-body)| `[]*DetectionToolResponseBody` |  | | DetectionTools of task | `[{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"},{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"},{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"}]` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of task | `1` |
| ignore | boolean| `bool` | ✓ | | Ignore of task | `true` |
| ignoreUntil | string| `string` |  | | Ignore until of task | `vector` |
| mainUserID | int64 (formatted integer)| `int64` |  | | MainUserID of task | `1` |
| mainUserName | string| `string` |  | | MainUserName of task | `main-user-name` |
| packageStatuses | map of string| `map[string]string` |  | | packageStatus of task | `{"bash":"resolved"}` |
| pkgCpes | [][PkgCpeChildResponseBody](#pkg-cpe-child-response-body)| `[]*PkgCpeChildResponseBody` |  | | Pcakge And Cpe list of task | `[{"affectedProcs":[{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"}],"cpeID":1,"cpeURI":"cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*","createdAt":"2018-07-14T08:13:28Z","name":"package01","newRelease":"new release","newVersion":"2.0","pkgID":1,"release":"release","repository":"repository","serverID":1,"updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"affectedProcs":[{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"}],"cpeID":1,"cpeURI":"cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*","createdAt":"2018-07-14T08:13:28Z","name":"package01","newRelease":"new release","newVersion":"2.0","pkgID":1,"release":"release","repository":"repository","serverID":1,"updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"affectedProcs":[{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"}],"cpeID":1,"cpeURI":"cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*","createdAt":"2018-07-14T08:13:28Z","name":"package01","newRelease":"new release","newVersion":"2.0","pkgID":1,"release":"release","repository":"repository","serverID":1,"updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"affectedProcs":[{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"}],"cpeID":1,"cpeURI":"cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*","createdAt":"2018-07-14T08:13:28Z","name":"package01","newRelease":"new release","newVersion":"2.0","pkgID":1,"release":"release","repository":"repository","serverID":1,"updatedAt":"2018-07-14T08:13:28Z","version":"1.0"}]` |
| priority | string| `string` | ✓ | | Priority of task | `high` |
| roleID | int64 (formatted integer)| `int64` | ✓ | | ServerRoleID of task | `1` |
| roleName | string| `string` | ✓ | | ServerRoleName of task | `server-role-name` |
| server | [ServerChildResponseBody](#server-child-response-body)| `ServerChildResponseBody` | ✓ | |  |  |
| serverID | int64 (formatted integer)| `int64` | ✓ | | ServerID of task | `1` |
| status | string| `string` | ✓ | | Status of task | `new` |
| subUserID | int64 (formatted integer)| `int64` |  | | SubUserID of task | `1` |
| subUserName | string| `string` |  | | SubUserName of task | `sub-user-name` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time of task | `2018-07-14T08:13:28Z` |



### <span id="task-get-task-list-response-body"></span> TaskGetTaskListResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| paging | [PagingResponseBody](#paging-response-body)| `PagingResponseBody` |  | |  |  |
| tasks | [][TaskListResponseBody](#task-list-response-body)| `[]*TaskListResponseBody` |  | | Task list | `[{"advisoryIDs":["advisoryID"],"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","detectionTools":[{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"},{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"},{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"}],"hasExploit":true,"hasMitigation":true,"hasWorkaround":true,"id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","osFamily":"centos","osVersion":"6","pkgCpeNames":["package1","package2"],"pkgNotFixedYet":true,"priority":"high","roleID":1,"roleName":"server-role-name","serverID":1,"serverName":"serverName","serverTags":["tag"],"serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"advisoryIDs":["advisoryID"],"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","detectionTools":[{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"},{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"},{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"}],"hasExploit":true,"hasMitigation":true,"hasWorkaround":true,"id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","osFamily":"centos","osVersion":"6","pkgCpeNames":["package1","package2"],"pkgNotFixedYet":true,"priority":"high","roleID":1,"roleName":"server-role-name","serverID":1,"serverName":"serverName","serverTags":["tag"],"serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"advisoryIDs":["advisoryID"],"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","detectionTools":[{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"},{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"},{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"}],"hasExploit":true,"hasMitigation":true,"hasWorkaround":true,"id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","osFamily":"centos","osVersion":"6","pkgCpeNames":["package1","package2"],"pkgNotFixedYet":true,"priority":"high","roleID":1,"roleName":"server-role-name","serverID":1,"serverName":"serverName","serverTags":["tag"],"serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"advisoryIDs":["advisoryID"],"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","detectionTools":[{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"},{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"},{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"}],"hasExploit":true,"hasMitigation":true,"hasWorkaround":true,"id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","osFamily":"centos","osVersion":"6","pkgCpeNames":["package1","package2"],"pkgNotFixedYet":true,"priority":"high","roleID":1,"roleName":"server-role-name","serverID":1,"serverName":"serverName","serverTags":["tag"],"serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"}]` |



### <span id="task-list-response-body"></span> TaskListResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| advisoryIDs | []string| `[]string` |  | | advisoryIDs of cve | `["advisoryID"]` |
| applyingPatchOn | date (formatted string)| `strfmt.Date` |  | | ApplyingPatchOn of task | `2018-07-13` |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | created time of task | `2018-07-14T08:13:28Z` |
| cveID | string| `string` | ✓ | | CVE ID of task | `CVE-2017-6799` |
| detectionTools | [][DetectionToolResponseBody](#detection-tool-response-body)| `[]*DetectionToolResponseBody` |  | | DetectionTools of task | `[{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"},{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"}]` |
| hasExploit | boolean| `bool` |  | | hasExploit of cve | `true` |
| hasMitigation | boolean| `bool` |  | | hasMitigation of cve | `true` |
| hasWorkaround | boolean| `bool` |  | | hasWorkaroundof cve | `true` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of task | `1` |
| ignore | boolean| `bool` | ✓ | | Ignore of task | `true` |
| ignoreUntil | string| `string` |  | | Ignore until of task | `vector` |
| mainUserID | int64 (formatted integer)| `int64` |  | | MainUserID of task | `1` |
| mainUserName | string| `string` |  | | MainUserName of task | `main-user-name` |
| osFamily | string| `string` | ✓ | | OS Name of server | `centos` |
| osVersion | string| `string` | ✓ | | OS Version of server | `6` |
| pkgCpeNames | []string| `[]string` |  | | Package And CPE Names of task | `["package1","package2"]` |
| pkgNotFixedYet | boolean| `bool` |  | | Flag of Not Fixed Yet of task | `true` |
| priority | string| `string` | ✓ | | Priority of task | `high` |
| roleID | int64 (formatted integer)| `int64` | ✓ | | ServerRoleID of task | `1` |
| roleName | string| `string` | ✓ | | ServerRoleName of task | `server-role-name` |
| serverID | int64 (formatted integer)| `int64` | ✓ | | ServerID of task | `1` |
| serverName | string| `string` | ✓ | | ServerName of task | `serverName` |
| serverTags | []string| `[]string` |  | | ServerTags of task | `["tag"]` |
| serverUuid | string| `string` | ✓ | | ServerUUID of task | `abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2` |
| status | string| `string` | ✓ | | Status of task | `new` |
| subUserID | int64 (formatted integer)| `int64` |  | | SubUserID of task | `1` |
| subUserName | string| `string` |  | | SubUserName of task | `sub-user-name` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time of task | `2018-07-14T08:13:28Z` |



### <span id="task-update-task-ignore-request-body"></span> TaskUpdateTaskIgnoreRequestBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ignoreUntil | string| `string` | ✓ | | ignore until. | `forever` |



### <span id="task-update-task-ignore-response-body"></span> TaskUpdateTaskIgnoreResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| advisoryIDs | []string| `[]string` |  | | advisoryIDs of cve | `["advisoryID"]` |
| applyingPatchOn | date (formatted string)| `strfmt.Date` |  | | ApplyingPatchOn of task | `2018-07-13` |
| comments | [][TaskCommentResponseBody](#task-comment-response-body)| `[]*TaskCommentResponseBody` |  | | Comment of task | `[{"comment":"comment","createdAt":"2018-07-14T08:13:28Z","id":1,"type":"system","updatedAt":"2018-07-14T08:13:28Z","userID":1,"userName":"user-name"},{"comment":"comment","createdAt":"2018-07-14T08:13:28Z","id":1,"type":"system","updatedAt":"2018-07-14T08:13:28Z","userID":1,"userName":"user-name"},{"comment":"comment","createdAt":"2018-07-14T08:13:28Z","id":1,"type":"system","updatedAt":"2018-07-14T08:13:28Z","userID":1,"userName":"user-name"}]` |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | created time of task | `2018-07-14T08:13:28Z` |
| cveID | string| `string` | ✓ | | CVE ID of task | `CVE-2017-6799` |
| cvss | map of [ReadCloser](#read-closer)| `map[string]io.ReadCloser` |  | | Key Value of CveID and Cvss of task | `{"Et quibusdam cumque tempore odio est quam.":"Dolores similique dolorem recusandae qui ea.","Odit et a rerum necessitatibus voluptatibus.":"Deleniti enim ut.","Voluptas culpa distinctio sed rem ab.":"A eos facilis nostrum."}` |
| detectionMethods | [][DetectionMethodResponseBody](#detection-method-response-body)| `[]*DetectionMethodResponseBody` |  | | DetectionMethod of task | `[{"name":"vuls","reliabilityScore":100},{"name":"vuls","reliabilityScore":100}]` |
| detectionTools | [][DetectionToolResponseBody](#detection-tool-response-body)| `[]*DetectionToolResponseBody` |  | | DetectionTools of task | `[{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"},{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"},{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"}]` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of task | `1` |
| ignore | boolean| `bool` | ✓ | | Ignore of task | `true` |
| ignoreUntil | string| `string` |  | | Ignore until of task | `vector` |
| mainUserID | int64 (formatted integer)| `int64` |  | | MainUserID of task | `1` |
| mainUserName | string| `string` |  | | MainUserName of task | `main-user-name` |
| packageStatuses | map of string| `map[string]string` |  | | packageStatus of task | `{"bash":"resolved"}` |
| pkgCpes | [][PkgCpeChildResponseBody](#pkg-cpe-child-response-body)| `[]*PkgCpeChildResponseBody` |  | | Pcakge And Cpe list of task | `[{"affectedProcs":[{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"}],"cpeID":1,"cpeURI":"cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*","createdAt":"2018-07-14T08:13:28Z","name":"package01","newRelease":"new release","newVersion":"2.0","pkgID":1,"release":"release","repository":"repository","serverID":1,"updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"affectedProcs":[{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"}],"cpeID":1,"cpeURI":"cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*","createdAt":"2018-07-14T08:13:28Z","name":"package01","newRelease":"new release","newVersion":"2.0","pkgID":1,"release":"release","repository":"repository","serverID":1,"updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"affectedProcs":[{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"}],"cpeID":1,"cpeURI":"cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*","createdAt":"2018-07-14T08:13:28Z","name":"package01","newRelease":"new release","newVersion":"2.0","pkgID":1,"release":"release","repository":"repository","serverID":1,"updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"affectedProcs":[{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"}],"cpeID":1,"cpeURI":"cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*","createdAt":"2018-07-14T08:13:28Z","name":"package01","newRelease":"new release","newVersion":"2.0","pkgID":1,"release":"release","repository":"repository","serverID":1,"updatedAt":"2018-07-14T08:13:28Z","version":"1.0"}]` |
| priority | string| `string` | ✓ | | Priority of task | `high` |
| roleID | int64 (formatted integer)| `int64` | ✓ | | ServerRoleID of task | `1` |
| roleName | string| `string` | ✓ | | ServerRoleName of task | `server-role-name` |
| server | [ServerChildResponseBody](#server-child-response-body)| `ServerChildResponseBody` | ✓ | |  |  |
| serverID | int64 (formatted integer)| `int64` | ✓ | | ServerID of task | `1` |
| status | string| `string` | ✓ | | Status of task | `new` |
| subUserID | int64 (formatted integer)| `int64` |  | | SubUserID of task | `1` |
| subUserName | string| `string` |  | | SubUserName of task | `sub-user-name` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time of task | `2018-07-14T08:13:28Z` |



### <span id="task-update-task-request-body"></span> TaskUpdateTaskRequestBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| applyingPatchOn | date (formatted string)| `strfmt.Date` |  | | applyingPatchOn (YYYY-MM-DD) UTC | `2018-07-14` |
| mainUserID | int64 (formatted integer)| `int64` |  | | mainUserID of task | `1` |
| priority | string| `string` |  | | Priority of task | `medium` |
| status | string| `string` |  | | Status of task | `ongoing` |
| subUserID | int64 (formatted integer)| `int64` |  | | subUserID of task | `2` |



### <span id="task-update-task-response-body"></span> TaskUpdateTaskResponseBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| advisoryIDs | []string| `[]string` |  | | advisoryIDs of cve | `["advisoryID"]` |
| applyingPatchOn | date (formatted string)| `strfmt.Date` |  | | ApplyingPatchOn of task | `2018-07-13` |
| comments | [][TaskCommentResponseBody](#task-comment-response-body)| `[]*TaskCommentResponseBody` |  | | Comment of task | `[{"comment":"comment","createdAt":"2018-07-14T08:13:28Z","id":1,"type":"system","updatedAt":"2018-07-14T08:13:28Z","userID":1,"userName":"user-name"},{"comment":"comment","createdAt":"2018-07-14T08:13:28Z","id":1,"type":"system","updatedAt":"2018-07-14T08:13:28Z","userID":1,"userName":"user-name"},{"comment":"comment","createdAt":"2018-07-14T08:13:28Z","id":1,"type":"system","updatedAt":"2018-07-14T08:13:28Z","userID":1,"userName":"user-name"}]` |
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | created time of task | `2018-07-14T08:13:28Z` |
| cveID | string| `string` | ✓ | | CVE ID of task | `CVE-2017-6799` |
| cvss | map of [ReadCloser](#read-closer)| `map[string]io.ReadCloser` |  | | Key Value of CveID and Cvss of task | `{"Et voluptatem et omnis quibusdam eum ex.":"Ut quas.","Saepe qui eligendi est accusamus dolores.":"Repellat doloribus neque quasi sequi impedit."}` |
| detectionMethods | [][DetectionMethodResponseBody](#detection-method-response-body)| `[]*DetectionMethodResponseBody` |  | | DetectionMethod of task | `[{"name":"vuls","reliabilityScore":100},{"name":"vuls","reliabilityScore":100},{"name":"vuls","reliabilityScore":100},{"name":"vuls","reliabilityScore":100}]` |
| detectionTools | [][DetectionToolResponseBody](#detection-tool-response-body)| `[]*DetectionToolResponseBody` |  | | DetectionTools of task | `[{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"},{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"},{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"},{"name":"changelog","patchAppliedAt":"2018-07-14T08:13:28Z"}]` |
| id | int64 (formatted integer)| `int64` | ✓ | | ID of task | `1` |
| ignore | boolean| `bool` | ✓ | | Ignore of task | `true` |
| ignoreUntil | string| `string` |  | | Ignore until of task | `vector` |
| mainUserID | int64 (formatted integer)| `int64` |  | | MainUserID of task | `1` |
| mainUserName | string| `string` |  | | MainUserName of task | `main-user-name` |
| packageStatuses | map of string| `map[string]string` |  | | packageStatus of task | `{"bash":"resolved"}` |
| pkgCpes | [][PkgCpeChildResponseBody](#pkg-cpe-child-response-body)| `[]*PkgCpeChildResponseBody` |  | | Pcakge And Cpe list of task | `[{"affectedProcs":[{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"}],"cpeID":1,"cpeURI":"cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*","createdAt":"2018-07-14T08:13:28Z","name":"package01","newRelease":"new release","newVersion":"2.0","pkgID":1,"release":"release","repository":"repository","serverID":1,"updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"affectedProcs":[{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"}],"cpeID":1,"cpeURI":"cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*","createdAt":"2018-07-14T08:13:28Z","name":"package01","newRelease":"new release","newVersion":"2.0","pkgID":1,"release":"release","repository":"repository","serverID":1,"updatedAt":"2018-07-14T08:13:28Z","version":"1.0"}]` |
| priority | string| `string` | ✓ | | Priority of task | `high` |
| roleID | int64 (formatted integer)| `int64` | ✓ | | ServerRoleID of task | `1` |
| roleName | string| `string` | ✓ | | ServerRoleName of task | `server-role-name` |
| server | [ServerChildResponseBody](#server-child-response-body)| `ServerChildResponseBody` | ✓ | |  |  |
| serverID | int64 (formatted integer)| `int64` | ✓ | | ServerID of task | `1` |
| status | string| `string` | ✓ | | Status of task | `new` |
| subUserID | int64 (formatted integer)| `int64` |  | | SubUserID of task | `1` |
| subUserName | string| `string` |  | | SubUserName of task | `sub-user-name` |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time of task | `2018-07-14T08:13:28Z` |



### <span id="tmp-metric-response-body"></span> TmpMetricResponseBody


> TmpMetric describes a tmpMetric
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| createdAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | created time | `2018-07-14T08:13:28Z` |
| e | string| `string` | ✓ | | E of tmpMetric |  |
| rc | string| `string` | ✓ | | RC of tmpMetric |  |
| rl | string| `string` | ✓ | | RL of tmpMetric |  |
| updatedAt | date-time (formatted string)| `strfmt.DateTime` | ✓ | | updated time | `2018-07-14T08:13:28Z` |


