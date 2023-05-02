---
title: Taylor's Petstore v1.0.0
language_tabs:
  - go: Go
language_clients:
  - go: ""
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2

---

<h1 id="taylor-s-petstore">Taylor's Petstore v1.0.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Taylor's Petstore example to demonstrate openapi 3.0

Base URLs:

* <a href="http://localhost:8080">http://localhost:8080</a>

* <a href="https://taylorpetstore.dev.tencent.com">https://taylorpetstore.dev.tencent.com</a>

* <a href="https://taylorpetstore.test.tencent.com">https://taylorpetstore.test.tencent.com</a>

Email: <a href="mailto:taylorzyx@hotmail.com">Taylor Zhang</a> Web: <a href="https://taylorzyx.hashnode.dev/">Taylor Zhang</a> 
License: <a href="https://www.apache.org/licenses/LICENSE-2.0.html">Apache 2.0</a>

<h1 id="taylor-s-petstore-default">Default</h1>

## Returns all pets

<a id="opIdfindPets"></a>

`GET /pets`

Returns all pets from the system that the user has access to

<h3 id="returns-all-pets-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|tags|query|array[string]|false|tags to filter by|
|limit|query|integer(int32)|false|maximum number of results to return|

> Example responses

> 200 Response

```json
[
  null
]
```

<h3 id="returns-all-pets-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|pet response|Inline|
|default|Default|unexpected error|[Error](#schemaerror)|

<h3 id="returns-all-pets-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[allOf]|false|none|none|

*allOf*

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» *anonymous*|object|false|none|none|
|»» name|string|true|none|Name of the pet|
|»» tag|string|false|none|Type of the pet|

*and*

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» *anonymous*|object|false|none|none|
|»» id|integer(int64)|true|none|Unique id of the pet|

<aside class="success">
This operation does not require authentication
</aside>

## Creates a new pet

<a id="opIdaddPet"></a>

`POST /pets`

Creates a new pet in the store. Duplicates are allowed

> Body parameter

```json
{
  "name": "string",
  "tag": "string"
}
```

<h3 id="creates-a-new-pet-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[NewPet](#schemanewpet)|true|Pet to add to the store|

> Example responses

> 200 Response

```json
null
```

<h3 id="creates-a-new-pet-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|pet response|[Pet](#schemapet)|
|default|Default|unexpected error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## Returns a pet by ID

<a id="opIdfindPetByID"></a>

`GET /pets/{id}`

Returns a pet based on a single ID

<h3 id="returns-a-pet-by-id-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|integer(int64)|true|ID of pet to fetch|

> Example responses

> 200 Response

```json
null
```

<h3 id="returns-a-pet-by-id-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|pet response|[Pet](#schemapet)|
|default|Default|unexpected error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## Deletes a pet by ID

<a id="opIddeletePet"></a>

`DELETE /pets/{id}`

deletes a single pet based on the ID supplied

<h3 id="deletes-a-pet-by-id-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|integer(int64)|true|ID of pet to delete|

> Example responses

> default Response

```json
{
  "code": 0,
  "message": "string"
}
```

<h3 id="deletes-a-pet-by-id-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|204|[No Content](https://tools.ietf.org/html/rfc7231#section-6.3.5)|pet deleted|None|
|default|Default|unexpected error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_Pet">Pet</h2>

<a id="schemapet"></a>
<a id="schema_Pet"></a>
<a id="tocSpet"></a>
<a id="tocspet"></a>

```json
null

```

### Properties

allOf

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[NewPet](#schemanewpet)|false|none|none|

and

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|object|false|none|none|
|» id|integer(int64)|true|none|Unique id of the pet|

<h2 id="tocS_NewPet">NewPet</h2>

<a id="schemanewpet"></a>
<a id="schema_NewPet"></a>
<a id="tocSnewpet"></a>
<a id="tocsnewpet"></a>

```json
{
  "name": "string",
  "tag": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|name|string|true|none|Name of the pet|
|tag|string|false|none|Type of the pet|

<h2 id="tocS_Error">Error</h2>

<a id="schemaerror"></a>
<a id="schema_Error"></a>
<a id="tocSerror"></a>
<a id="tocserror"></a>

```json
{
  "code": 0,
  "message": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|code|integer(int32)|true|none|Error code|
|message|string|true|none|Error message|

