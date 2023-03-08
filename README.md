Go string helpers inspered by Laravel's :)
# Requirements
- Go version 1.19

# Setup
## Clone the project

- Clone this project by executing `git clone https://github.com/auliawiguna/goshaka-stringhelper`
- Execute `go mod download && go mod verify`

## Usage in your projects
- Note that you need to include the v in the version tag. Execute `go get github.com/auliawiguna/goshaka-stringhelper@v0.1.1`



# Functions

### After

To get string after given mark

```
After(source, mark string) string
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `source` | `string` | **Required**. Source text |
| `mark` | `string` | **Required**. Mark text |

Usage:

```
After("This is World of Oz", "is World")
```

Result:
` of Oz`
### AfterLast

To get string after last occurence of given mark

```
  AfterLast(source, mark string)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `source` | `string` | **Required**. Source text |
| `mark` | `string` | **Required**. Mark text |

Usage:

```
AfterLast("This is World of Oz and World of Colony and the World of Irony", "World of")
```

Result:
` Irony`

### Before

To get string before given mark

```
  Before(source, mark string)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `source` | `string` | **Required**. Source text |
| `mark` | `string` | **Required**. Mark text |

Usage:

```
Before("This is World of Oz", "is World")
```

Result:
`This `

### BeforeLast

To get string before the last occurence of given mark

```
  BeforeLast(source, mark string)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `source` | `string` | **Required**. Source text |
| `mark` | `string` | **Required**. Mark text |

Usage:

```
BeforeLast("This is World of Oz and World of Colony and the World of Irony", "World of")
```

Result:
`This is World of Oz and World of Colony and the  `


# Running Tests

To run tests, run the following command

```bash
  go test
```