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

> 下滑 for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Taylor's Petstore example to demonstrate openapi 3.0

Base URLs:

* <a href="http://localhost:8080">http://localhost:8080</a>

* <a href="https://taylorpetstore.dev.tencent.com">https://taylorpetstore.dev.tencent.com</a>

* <a href="https://taylorpetstore.test.tencent.com">https://taylorpetstore.test.tencent.com</a>

Email: <a href="mailto:taylorzyx@hotmail.com">Taylor Zhang</a> Web: <a href="https://taylorzyx.hashnode.dev/">Taylor Zhang</a> 
License: <a href="https://www.apache.org/licenses/LICENSE-2.0.html">Apache 2.0</a>

<h1 id="taylor-s-petstore-default">Default</h1>

## 获取所有 pet 的信息

<a id="opIdfindPets"></a>

`GET /pets`

获取系统中所有用户能够访问的 pet 的信息。

<h3 id="获取所有-pet-的信息-parameters">输入参数</h3>

|参数名称|必选|类型|描述|
|---|---|---|---|
|tags|false|array[string]|用来过滤 pet 的标签|
|limit|false|integer(int32)|返回的最大 pet 数量|

> Example responses

> 成功的 pet 列表

```json
[
  {
    "id": 1000,
    "name": "Miao",
    "tag": "cat"
  }
]
```

> default Response

```json
{
  "code": 0,
  "message": "string"
}
```

<h3 id="获取所有-pet-的信息-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功的 pet 列表|Inline|
|default|Default|错误信息|[Error](#schemaerror)|

<h3 id="获取所有-pet-的信息-responseschema">输出参数</h3>

Status Code **200**

|参数名称|类型|描述|
|---|---|---|
|*anonymous*|[allOf]|none|

*allOf*

|参数名称|类型|描述|
|---|---|---|
|*anonymous*|object|none|
|name|string|pet 的名字|
|tag|string|pet 的标签|

*and*

|参数名称|类型|描述|
|---|---|---|
|*anonymous*|object|none|
|id|integer(int64)|pet 的唯一 ID|

<aside class="success">
This operation does not require authentication
</aside>

## 创建新的 pet

<a id="opIdaddPet"></a>

`POST /pets`

创建新的 pet，如果 pet 已经存在，则返回错误。

> Body parameter

```json
{
  "name": "Miao",
  "tag": "cat"
}
```

<h3 id="创建新的-pet-parameters">输入参数</h3>

|参数名称|必选|类型|描述|
|---|---|---|---|
|body|true|[NewPet](#schemanewpet)|要创建并保存的 pet 信息|

> Example responses

> 成功创建的 pet 信息

```json
{
  "id": 1000,
  "name": "Miao",
  "tag": "cat"
}
```

> default Response

```json
{
  "code": 0,
  "message": "string"
}
```

<h3 id="创建新的-pet-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功创建的 pet 信息|[Pet](#schemapet)|
|default|Default|unexpected error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## 根据 ID 获取 pet 信息

<a id="opIdfindPetByID"></a>

`GET /pets/{id}`

根据 ID 获取 pet 信息，如果 pet 不存在，则返回错误。

<h3 id="根据-id-获取-pet-信息-parameters">输入参数</h3>

|参数名称|必选|类型|描述|
|---|---|---|---|
|id|true|integer(int64)|要获取的 pet 的 ID|

> Example responses

> 200 Response

```json
null
```

<h3 id="根据-id-获取-pet-信息-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|获取到的 pet 信息|[Pet](#schemapet)|
|default|Default|错误信息|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## 删除 pet

<a id="opIddeletePet"></a>

`DELETE /pets/{id}`

根据 ID 删除 pet，如果 pet 不存在，则返回错误。

<h3 id="删除-pet-parameters">输入参数</h3>

|参数名称|必选|类型|描述|
|---|---|---|---|
|id|true|integer(int64)|要删除的 pet 的 ID|

> Example responses

> default Response

```json
{
  "code": 0,
  "message": "string"
}
```

<h3 id="删除-pet-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|204|[No Content](https://tools.ietf.org/html/rfc7231#section-6.3.5)|成功删除 pet|None|
|default|Default|错误信息|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

# 复杂类型

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

|名称|类型|描述|
|---|---|---|
|*anonymous*|[NewPet](#schemanewpet)|none|

and

|名称|类型|描述|
|---|---|---|
|*anonymous*|object|none|
|id|integer(int64)|pet 的唯一 ID|

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

|名称|类型|描述|
|---|---|---|
|name|string|pet 的名字|
|tag|string|pet 的标签|

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

|名称|类型|描述|
|---|---|---|
|code|integer(int32)|错误码|
|message|string|错误信息|

