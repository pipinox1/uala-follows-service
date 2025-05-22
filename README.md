
# Followers Service

Manage follows relation between users

## API Reference

#### Follow user

```http
  POST /api/v1/follow/user/${user_id}/follow
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `user_id` | `string` | **Required**. ID of the user who will follow |
| `Content-Type` | `string` | **Required**. application/json |

**Request Body:**

| Field | Type     | Description                       |
| :---- | :------- | :-------------------------------- |
| `followed_id` | `string` | **Required**. ID of the user to follow |

#### Get user followings

```http
  GET /api/v1/follow/user/${user_id}/followings
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `user_id` | `string` | **Required**. ID of the user to get followings for |

#### Get user followers

```http
  GET /api/v1/follow/user/${user_id}/followers
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `user_id` | `string` | **Required**. ID of the user to get followers for |

## How to Run?


In order to run you have to clone:
- `uala-posts-service`
- `uala-followers-service`
- `uala-timeline-service`

In the `uala-posts-service` repository, you'll find the `docker-compose` file used to start all services.

### Steps:


1. After cloning each repository, run the following command inside each one:

   ```bash
   make build
   ```

2. Once all images are built, navigate to the uala-posts-service repository and run:

   ```bash
   docker-compose up -d
   ```

### Service Ports:
- posts-service: 8080
- timeline-service: 8081
- followers-service: 8082