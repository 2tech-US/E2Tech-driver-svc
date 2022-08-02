# E2TechEcommerce API document

This is a documentation guide for the Driver Service.

## Database

- note: `createdAt` `updatedAt` is auto generate in everys table

### driver

| Column         |  Type   |             |
| -------------- | :-----: | :---------: |
| phone          | String  | Primary key |
| name           | String  |             |
| dob            |  Date   |             |
| identification | String  |             |
| avatar         | String  |             |
| avaiable       | boolean |             |
| taxisId        |  Long   | foreign key |

### taxis

| Column   |  Type   |                            |
| -------- | :-----: | :------------------------: |
| taxisId  |  Long   | Primary key, auto generate |
| phone    | String  |        Primary key         |
| plate    | String  |                            |
| type     | String  |      [car, motorbike]      |
| isDelete | Boolean |     [ default = true]      |

## API (V1)

### driver

#### Get Drivers :

```
GET /api/v1/driver/admin/
```

Authorization token: `admin`

Query parameters available (Pagination):

- `limit`: `optional` `integer` `min=1`, limit the result
- `page`: `optional` `integer` `min=1`, page requested, need `limit` to work
- `sort`: `optional` `string` name of the field to sort results with `default is 'createdAt`
- `order_by`: `optional` `must be either asc or desc`, only work it `sort` is specified, default to `desc`
- `query`: `optional` `string`, search by phone's number

#### Get driver info:

```
GET /api/v1/driver/{phone}
```

Authorization token: `driver` , `admin`

#### Update driver info:

```
PUT /api/v1/driver/{phone}
```

Authorization token: `driver`

Body parameters available (`Json`):

- `name`: `optinal` `string`
- `dob`: `optinal` `date`
- `avatar`: `optinal` `string`

#### Update driver avaiable:

```
PUT /api/v1/driver/avaiable/{phone}
```

Authorization token: `driver`

Body parameters available (`Json`):

- `available`: `require` `boolean`

### Taxis

#### Get Taxis

```
GET /api/v1/driver/taxis/admin/
```

Authorization token: `admin`

Query parameters available (Pagination):

- `limit`: `optional` `integer` `min=1`, limit the result
- `page`: `optional` `integer` `min=1`, page requested, need `limit` to work
- `sort`: `optional` `string` name of the field to sort results with `default is 'createdAt`
- `order_by`: `optional` `must be either asc or desc`, only work it `sort` is specified, default to `desc`
- `query`: `optional` `string`, search by phone's number

#### Get Taxis's info own by driver

```
GET /api/v1/driver/taxis/{phone}
```

Authorization token: `driver` `admin`

#### Get Taxis's info by taxisId

```
GET /api/v1/driver/taxis/{taxisId}
```

Authorization token: `driver` `admin`

#### Create Taxis

```
POST /api/v1/driver/taxis/{phone}
```

Authorization token: `driver`

Body parameters available (`Json`):

- `plate`: `require` `string`
- `type`: `require` `string`

#### Delete Taxis

```
delete /api/v1/driver/taxis/{taxisId}
```

Authorization token: `driver`
