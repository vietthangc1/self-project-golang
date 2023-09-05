# Backend Recommendation Project
##### Inspired by Tiki recommendation system

This is the project adapted from Tiki recommendation, with multiple blocks and models config to deliver products to customers

## Table of Contents
- [Concepts & Definition](#concepts-and-definition)
- [Before running the project](#before-running-the-project)
- [User Admin/Customer Identities](#user-admin-customer-identities)
- [Main server APIS](#main-server-apis)
- [Product server APIS](#product-server-apis)
- [Requirements for google sheets used for model configuration](#requirements-for-google-sheets-used-for-model-configuration)
- [Next steps](#next-steps)

## Concepts and Definition
- There are serverl blocks (determined by `block_code`) on the website to recommend suitable products for each customer (determined by `customer_id`)
- Each block is configured by one or more models (determined by `model_code` or `model_id`)
- Each model has a configuration file to guide the products (determined by `product_id`) ranked for each customer
- The api would call to get suggested products for customer of a specific block

## Before running the project
#### 1. Applications
Run Redis and Mysql through Docker
```
docker compose up
```
#### 2. Environment Variables
```
export MYSQL_ADDR // the DNS connection uri to mysql
export REDIS_ADDRS // redis host
export TOKEN_HOUR_LIFESPAN // redis TTL
export TOKEN_SECRET_KEY // secret key for JW Token
export MONGO_ADDR= // the connection uri to mongodb
export GOOGLE_APPLICATION_CREDENTIALS // gcloud credentials, use for read gg sheets file
```
#### 3. Servers
Run concurrently in 4 terminals:
- Product server internal (default host: localhost:8081): This server shall return product information (without athorization to be called by the main server)
```
make run-product-internal
```
- Product server external (default host: localhost:8082): This server shall return product information by APIS (with authorization to get info)
```
make run-product
```
- Worker server: The cron job service update product info to cache in a schedule interval
```
make run-worker
```
- Main server (default host: localhost:8080): Main apis to return products for blocks and some for debugging
```
make run-server
```

## User Admin Customer Identities
Use middlewares for identify user admin/customer
- User Admin: input token through headers `X-User-Token`
- Customer: input token through headers `X-Customer-Token` or input the id of customer through `customer_id` query params

## Main server APIS
#### Main API
| Route | Input | Functions |
| ------ | ------ | ------ |
| GET `/data` | query: `block_code`, `page_size`: number of products returned, `begin_cursor`: the position of first product (start and default with 0) | Return list of products (with full information) suggested for `customer_id` in the block `block_code` |
#### Model APIS
| Route | Input | Functions |
| ------ | ------ | ------ |
| POST `/model/create` | body: `code`, `source`: {`sheet_id`: id of google sheets file used for configuration, `sheet_name`: worksheet name}| Create a new model |
| GET `/model/id/:id` | params: `id`: id of model to get info | Get Model info |
| GET `/model/code/:code` | params: `code`: model_code | Get Model info |
| GET `/model/score` | query: `model`: model_code | Return array of `{product_id, score}` for customer |
#### Block APIS
| Route | Input | Functions |
| ------ | ------ | ------ |
| POST `/block/create` | body: `code`, `description`, `model_ids`: array of model ids used in this block| Create a new block |
| GET `/block/id/:id` | params: `id`: id of block to get info | Get Block info |
| GET `/block/code/:code` | params: `code`: block_code of block to get info | Get Block info |
#### User APIS
| Route | Input | Functions |
| ------ | ------ | ------ |
| POST `/user/create` | body: `name`, `role`: 1 for read-only, 2 for write, `email`, `password`| Create a new user |
| GET `/user/get/:id` | params: `id`| Get User info |
| POST `/login` | body: `email`, `password`| Return user token |

## Product server APIS
| Route | Input | Functions |
| ------ | ------ | ------ |
| POST `/create` | body: `name`, `price`, `category`, `sub_category`, `sku`, `image_uri`| Create a new product (`user_role` = 2)|
| GET `/get/:id` | params: `id`| Get Product info (`user_role` = 1 or 2) |

## Requirements for google sheets used for model configuration
- Column A: `key`, this column specifies customer_id, use key "-" for other customers not custom in this model.
- Column B: `product id`: specifies product ids
- Column C: `score`: specifies score of `product_id` in col B and `custoemr_id` in col A. The higher score, the higher position of products when called by the `/data` api

Reference: [Model example configuration file](https://docs.google.com/spreadsheets/d/1efmk3-azGTV_h0oLyUjXkYamZDzF3dxZvCstiI8c47s/edit#gid=0)

## Next steps
Include Kafka (or other message broker service):
- To capture change in product info
- Add Order API and use Kafka to save the purchased products to exclude it from product lists